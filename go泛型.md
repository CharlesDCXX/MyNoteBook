# 泛型使用场景

1. 当写几乎一样的函数的时候

    ```go
    func containsUint8(needle uint8, haystack []uint8) bool {
    for _, v := range haystack {
        if v == needle {
            return true
        }
    }
    return false
    }

    func containsInt(needle int, haystack []int) bool {
    for _, v := range haystack {
        if v == needle {
            return true
        }
    }
    return false
    }
    ```

    使用泛型

    ```go
    func contains[E constraints.Ordered](needle E, haystack []E) bool {
    for _, v := range haystack {
        if v == needle {
            return true
        }
    }
    return false
    }
    ```

2. 关于集合类型：（切片、映射、数组）

3. 数据结构

