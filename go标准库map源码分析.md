﻿# map源码分析(UTF-8 with BOM)

> go version go1.20.12 linux/amd64
>
> map源码路径 go/src/runtime/map.go

## 使用方式

go中 map 类似于 java 的 HashMap，Python的字典(dict)，是一种存储键值对(Key-Value)的数据解构。使用方式和其他语言几乎没有区别。

```go
// 仅声明
m1 := make(map[string]int)
// 声明时初始化
m2 := map[string]string{
    "Sam": "Male",
    "Alice": "Female",
    "Jony": "Male",
}
// 赋值/修改
m1["Tom"] = 18
```

但是在for 循环 map 的时候，map 的key是随机出现的，并没有进行排序

```go
for k, v := range m2 {
    fmt.Println(k, v)
}

Sam Male
Alice Female
Jony Male

Alice Female
Jony Male
Sam Male
```

## Map解释

> A map is just a hash table. The data is arranged into an array of buckets. Each bucket contains up to 8 key/elem pairs. The low-order bits of the hash are used to select a bucket. Each bucket contains a few high-order bits of each hash to distinguish the entries within a single bucket.
>
> If more than 8 keys hash to a bucket, we chain on extra buckets.
>
> When the hashtable grows, we allocate a new array of buckets twice as big. Buckets are incrementally copied from the old bucket array to the new bucket array.
>
> Map只是一个哈希表。 数据被排列到桶数组中。 每个存储桶最多包含 8 个键/元素对。 哈希的低位用于选择存储桶。 每个桶包含每个散列的一些高阶位，以区分单个桶内的条目。
>
> 如果超过 8 个键散列到一个存储桶，我们会链接到额外的存储桶上。
>
> 当哈希表增长时，我们分配一个新的存储桶数组，其大小是原来的两倍。 存储桶从旧存储桶数组增量复制到新存储桶数组。

Go语言的 map 在并发环境中是**不安全**的，即在多个 goroutine 同时读写 map 时会出现竞态条件（race condition）的问题。为了解决这个问题，Go语言提供了 sync 包中的 Map 类型来实现并发安全的 map。

这里主要是看 go 中 map 的**实现方式及扩容方式**。  

## hmap
hmap是Go语言中map的底层实现，占用8个字节，定义在runtime包中。初始化map其实就是在内存中生成了一个hmap实例，初始化时定义的map变量是一个指针，指向了hmap。它包含了以下字段：

```go
type hmap struct {
    // count是map中键值对的数量
    count int
    // B是哈希表的桶的数量。注意，实际的桶数量是2^B，因为B的值是2的整数次幂。
    B int
    // hash0是用于计算哈希值的随机数种子
    hash0 uint32
    // buckets是哈希表的桶的数组
    buckets unsafe.Pointer
    // 前一个存储桶数组大小的一半，仅在增长时非零
    oldbuckets unsafe.Pointer 

    // extra是一些附加的标志位，如是否正在扩容等
    extra *mapextra
}

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// 如果 key 和 elem 都不包含指针并且是内联的，那么我们将存储桶类型标记为不包含指针。 这避免了扫描此类map。
	// 但是，bmap.overflow 是一个指针。 为了使溢出桶保持活动状态，我们将指向所有溢出桶的指针存储在 hmap.extra.overflow 和 hmap.extra.oldoverflow 中。
	// 仅当 key 和 elem 不包含指针时才使用 Overflow 和 oldoverflow。
	// overflow 包含 hmap.buckets 的 overflow 桶。
	// oldoverflow 包含 hmap.oldbuckets 的 overflow 桶。
	// 间接允许在 hiter 中存储指向切片的指针。
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}
```

## bmap

bmap是Go语言中map的桶的底层实现，是哈希表中的一个bucket。每个bucket都是一个连续的内存块，数据结构如下：

```go
type bmap struct {
    // tophash 通常包含此存储桶中每个键的哈希值的顶部字节
    tophash [bucketCnt]uint8
}
```

其中，tophash数组是一个长度为bucketCnt（8）的数组，它的每个元素都是一个字节。当一个键值对被插入到桶中时，它的哈希值会被计算出来，并且将哈希值的前8位存储在tophash数组中，用于快速判断是否有哈希冲突。如果两个键的哈希值的前8位相同，则它们会被认为是相等的，需要进一步比较完整的哈希值和键值。

## 创建map的函数makemap

```go
// makemap 为 make(map[k]v,hint) 实现 Go 映射创建。
// 如果编译器确定可以在堆栈上创建映射或第一个存储桶，则 h and/or bucket 可能为非零。
// 如果h != nil，则可以直接在h中创建map。
// 如果h.buckets != nil，则指向的bucket可以作为第一个bucket。
func makemap(t *maptype, hint int, h *hmap) *hmap {
    mem, overflow := math.MulUintptr(uintptr(hint), t.bucket.size)
    if overflow || mem > maxAlloc {
       hint = 0
    }

    // initialize Hmap
    if h == nil {
       h = new(hmap)
    }
    h.hash0 = fastrand()

    // Find the size parameter B which will hold the requested # of elements.
    // For hint < 0 overLoadFactor returns false since hint < bucketCnt.
    B := uint8(0)
    for overLoadFactor(hint, B) {
       B++
    }
    h.B = B

    // allocate initial hash table
    // if B == 0, the buckets field is allocated lazily later (in mapassign)
    // If hint is large zeroing this memory could take a while.
    if h.B != 0 {
       var nextOverflow *bmap
       h.buckets, nextOverflow = makeBucketArray(t, h.B, nil)
       if nextOverflow != nil {
          h.extra = new(mapextra)
          h.extra.nextOverflow = nextOverflow
       }
    }

    return h
}
```

### 参数说明

- `t *maptype`: 指向 map 的类型信息的指针，包括键和值的类型。
- `hint int`: 提示值，预估 map 将要存储的元素数量，用于优化内存分配。
- `h *hmap`: 指向 map 的结构体的指针，如果为 nil，则函数内部会创建一个新的。

### 函数执行步骤

1. **内存计算**：
   - 计算所需的内存大小，这是根据提示值 `hint` 和每个存储桶（bucket）的大小 `t.bucket.size` 来计算的。
   - 使用 `math.MulUintptr` 函数进行计算，以检测乘法是否溢出。
   - 如果计算结果溢出或所需内存超过最大允许分配量 `maxAlloc`，将 `hint` 设置为 0。
2. **哈希表（Hmap）初始化**：
   - 如果传入的 `h *hmap` 是 nil，函数会创建一个新的 `hmap` 结构体实例。
   - 为新的哈希表分配一个随机的初始哈希种子 `hash0`，使用 `fastrand()` 函数生成。
3. **计算存储桶大小**：
   - 初始化变量 `B`（用于记录存储桶的尺寸级别）为 0。
   - 通过循环调用 `overLoadFactor(hint, B)` 函数来确定最小的 `B`，使得给定的 `hint`（元素数量）不会导致负载因子过高。
   - 负载因子过高会导致哈希表操作效率下降，因为冲突更多，查找速度变慢。
4. **分配初始哈希表**：
   - 如果 `B` 不为 0，即需要立即分配存储桶，调用 `makeBucketArray(t, h.B, nil)` 来创建存储桶数组。
   - 这个函数返回两个值：存储桶数组的起始地址和指向可能的溢出存储桶的指针。
   - 如果有溢出存储桶（即 `nextOverflow` 不为 nil），会创建一个 `mapextra` 结构体来存储额外的溢出信息。
5. **返回**：
   - 函数最后返回初始化好的 `hmap` 结构体指针。

## 定位key

主要就是bmap的实现方式，有header和key数组和value数组,通过高8位确认存在tophash里的值，然后在key数组里进行遍历查询。

[接下来参考](./resource/map详解.html)

## 存储数据的过程

1. 计算key的哈希值。当一个键值对被插入到map中时，程序会首先计算该键的哈希值。Go语言中使用的是MurmurHash3算法，它可以将任意长度的数据快速且均匀地映射为固定长度的哈希值。
2. 定位到对应的桶。将哈希值对桶数量取模，得到对应的桶的索引，也就是buckets数组的下标。
3. 在桶中查找键值对。如果桶中没有键值对，则直接插入新的键值对。如果桶中已经存在键值对，则需要查找并更新该键值对。具体查找过程如下：
   1. 比较键的哈希值和tophash数组的值，如果不相等，则跳过该桶。
   2. 如果哈希值和tophash数组的值相等，则比较键值对的键是否相等，如果不相等，则继续查找该桶中的下一个键值对
   3. 如果键相等，则说明找到了对应的键值对，可以更新该键值对的值。

查找数据/删除数据的过程与存储类似，这里不再赘述。需要注意的是，当从 map 中删除键值对时，并不会立即释放对应的内存空间，而是会将对应的键值对标记为删除状态。当下一次对 map 进行操作时，Go 的垃圾回收机制会扫描标记为删除状态的键值对，并将它们从 map 中永久删除，并回收相应的内存空间。这种删除机制被称为延迟删除（deferred deletion），它可以避免在删除键值对时频繁地分配和释放内存空间，从而提高性能。

## hash冲突

按以上存储数据的过程，当插入一个键值对时，会首先根据键的哈希值计算出对应的桶，然后将键值对存储到这个桶中。如果发生哈希冲突，也就是两个不同的键计算出的哈希值相同，那么这两个键值对就会存储在同一个桶中。

为了解决哈希冲突，Go 采用了链式哈希表来存储键值对。在同一个桶中，可能会存储多个键值对，它们会通过一个链表链接在一起。当需要查找某个键值对时，首先根据键的哈希值计算出对应的桶，然后在这个桶中遍历链表，找到对应的键值对。

## 这里重点关注bmap中的overflow字段

## 扩容

### 扩容条件

负载因子 > 6.5时，也即平均每个bucket存储的键值对达到6.5个。
当溢出桶过多时：
当 B < 15 时，如果overflow的bucket数量超过 2^B。
当 B >= 15 时，overflow的bucket数量超过 2^15。
满足以上条件均会触发扩容机制。

### 扩容方案

等量扩容：

当一个桶中多次删除和增加数据之后，多次的hash冲突可能导致bmap的溢出桶很多，链表长度的增加导致扫描map的时间变长，而且浪费了大量存储空间。等量扩容实际上是针对这种情况对桶中数据做整理，把溢出桶中的数据向链表头搬迁，并删除空出来的overflow链表。这种情况下，元素会发生重排，但不会换桶。

增量扩容：

这种扩容发生在桶数量不够用时。具体步骤如下：

1. 计算新的桶的数量：当键值对数量小于 1024 时，每次扩容增加一倍的桶；当键值对数量大于等于 1024 时，每次扩容增加 25% 的桶。这个策略可以保证 map 在性能和空间利用率之间取得一个平衡；
2. 为新的桶分配内存空间：使用 Go 的内存分配器（memory allocator）为新的桶分配内存空间；
3. 将原来的桶中的键值对重新分配到新的桶中：对于每个桶，遍历桶中存储的所有键值对，计算键的哈希值，并根据新的桶数量计算出键对应的桶的位置；
4. 插入键值对到新的桶中：将键值对插入到新的桶中，形成新的链表，并将新的链表连接到哈希表中；
5. 释放原来的桶占用的内存空间：释放原来的桶占用的内存空间，这里需要注意的是，由于 Go 使用了指针指向键值对，所以在释放内存空间时需要注意先释放键值对，再释放链表，最后再释放桶。
考虑到如果map存储了数以亿计的key-value，一次性搬迁将会造成比较大的延时，Go采用逐步搬迁策略，即每次访问map时都会触发一次搬迁，每次搬迁2个键值对。

增量扩容会导致元素换桶。

遍历map为什么是无序的？
使用 range 多次遍历 map 时输出的 key 和 value 的顺序可能不同。

主要原因有2点：

- map在遍历时，并不是从固定的0号bucket开始遍历的，每次遍历，都会从一个随机值序号的bucket，再从其中随机的cell开始遍历
- map遍历时，是按序遍历bucket，同时按需遍历bucket中和其overflow bucket中的cell。但是map在扩容后，会发生key的搬迁，这造成原来落在一个bucket中的key，搬迁后，有可能会落到其他bucket中了，从这个角度看，遍历map的结果就不可能是按照原来的顺序了
  

map 本身是无序的，且遍历时顺序还会被随机化，如果想顺序遍历 map，需要对 map key 先排序（数组），再按照 key 的顺序遍历 map。
