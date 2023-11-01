# minikebe

## minikube含义

>Minikube是一个用于在本地开发环境中运行Kubernetes集群的工具。它能够在单个节点上启动一个虚拟机，并在其上运行一个简化的Kubernetes集群。通过使用Minikube，开发人员可以在本地进行Kubernetes应用程序的开发和测试，而无需依赖云上的Kubernetes集群。
>Minikube提供了一种快速、简便的方式来体验和学习Kubernetes。它支持在不同操作系统上运行，包括Windows、Mac和Linux。使用Minikube，开发人员可以轻松地在本地创建、启动和管理Kubernetes集群，以及部署和测试应用程序。
>Minikube通过在本地创建一个虚拟机来模拟Kubernetes集群的环境。它会自动安装和配置所需的组件，如Kubernetes控制平面、容器运行时（如Docker）等。一旦Minikube集群启动，开发人员可以使用kubectl命令行工具与集群进行交互，部署和管理应用程序。
>除了基本的Kubernetes功能外，Minikube还提供了其他有用的功能，如自动部署Kubernetes Dashboard、支持存储卷、支持多节点集群等。开发人员可以根据自己的需要配置和扩展Minikube集群。
>总而言之，Minikube是一个方便的工具，可用于在本地开发环境中运行和测试Kubernetes应用程序，以及学习和探索Kubernetes的功能和特性。

## minikube-vm

>Minikube-vm（Minikube Virtual Machine）是Minikube工具使用的虚拟机环境，用于在本地运行和管理Kubernetes集群。
>Minikube-vm基于虚拟化技术（如VirtualBox、VMware等）创建一个虚拟机，该虚拟机充当了一个简化的单节点Kubernetes集群。Minikube-vm会自动在虚拟机中安装和配置必要的组件，包括Kubernetes控制平面、容器运行时（如Docker）等。
>通过Minikube-vm，开发人员可以在本地快速搭建一个Kubernetes集群，以便进行开发、测试和学习。它提供了一种方便的方式来体验和探索Kubernetes的功能和特性，而无需依赖云上的Kubernetes集群。
>Minikube-vm还提供了其他有用的功能，如自动部署Kubernetes Dashboard、支持存储卷、支持多节点集群等。开发人员可以使用Minikube命令行工具与Minikube-vm进行交互，例如启动和停止集群、部署应用程序、扩展集群等操作。
>总而言之，Minikube-vm是Minikube工具使用的虚拟机环境，用于在本地运行和管理Kubernetes集群。它为开发人员提供了一个方便的方式来快速搭建和使用Kubernetes，并进行应用程序的开发和测试。

## minikube 命令速查

`minikube stop` 不会删除任何数据，只是停止 VM 和 k8s 集群。

`minikube delete` 删除所有 minikube 启动后的数据。

`minikube ip` 查看集群和 docker enginer 运行的 IP 地址。

`minikube pause` 暂停当前的资源和 k8s 集群

`minikube status` 查看当前集群状态

[Kubectl所有命令]: https://kubernetes.io/zh-cn/docs/reference/kubectl/cheatsheet/	"Kubectl所有命令"

## 安装minikube for macOS M1

``` bash
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube
```

或者使用brew

``` bash
brew install minikube
// 查看minikube安装在哪里
➜  ~ which minikube
/usr/local/bin/minikube
```

## 运行minikube

```bash
minikube start
```

## 构建容器，复习一下Dockerfile

```go
package main

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "[v1] Hello, Kubernetes!")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":3000", nil)
}
```

```Dockerfile
# Dockerfile
FROM golang:1.16-buster AS builder
RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go env -w GO111MODULE=auto
RUN go build -o main .

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /src/main /main
EXPOSE 3000
ENTRYPOINT ["/main"]
```

>1. FROM golang:1.16-buster AS builder：基础镜像指令，使用 Golang 1.16 版本的官方镜像作为构建阶段的基础镜像，并将其命名为 builder。
>
>2. RUN mkdir /src：在镜像中创建一个 /src 目录。
>
>3. ADD . /src：将当前目录中的所有文件和文件夹复制到镜像中的 /src 目录。
>
>4. WORKDIR /src：将工作目录设置为 /src。
>
>5. RUN go env -w GO111MODULE=auto：设置 Go 环境变量 GO111MODULE 为 auto，以便在构建过程中自动管理依赖。
>
>6. RUN go build -o main .：使用 Go 编译器构建应用程序，并将可执行文件命名为 main。
>
>7. FROM gcr.io/distroless/base-debian10：基础镜像指令，使用 distroless 镜像作为最终运行阶段的基础镜像。
>
>8. WORKDIR /：将工作目录切换到根目录。
>
>9. COPY --from=builder /src/main /main：从构建阶段的 builder 镜像中复制可执行文件 main 到最终镜像的根目录。
>
>10. EXPOSE 3000：声明容器将监听的端口号为 3000。
>
>11. ENTRYPOINT ["/main"]：设置容器启动时的默认命令，即运行 /main 可执行文件。

### 编译容器

```sh
docker build . -t duanchenxi/hellok8s:v1
```

### 运行容器

```sh
docker run -p 3000:3000 --name hellok8s -d guangzhengli/hellok8s:v1
```

## k8s中的pod

在 Kubernetes（K8s）中，Pod 是最小的可调度和管理的单位。它是一个或多个相关容器（应用程序）的组合，它们共享相同的网络和存储资源，并在同一节点上运行。
minikube创建pod并暴露pod的端口号。
container (容器) 的本质是进程，而 pod 是管理这一组进程的资源。
![image](./resource/pod关系.png)
1. 容器组合：Pod 可以包含一个或多个容器。这些容器通常是紧密相关的，并且一起协同工作以提供应用程序或服务的功能。
2. 共享网络命名空间：Pod 中的所有容器共享相同的网络命名空间。这意味着它们可以使用相同的 IP 地址和端口空间，可以通过 `localhost` 或 `127.0.0.1` 直接相互通信。
3. 共享存储卷：Pod 中的容器可以共享相同的存储卷，使它们可以在同一节点上读写相同的数据。这使得容器之间可以共享文件系统或共享存储。
4. 生命周期：Pod 是临时性的实体，它可以被创建、调度、运行和删除。Pod 的生命周期由 Kubernetes 控制器（如 Deployment 或 StatefulSet）管理。
5. 调度和部署：Pod 可以由 Kubernetes 调度到集群的任何可用节点上。调度器根据节点的资源可用性和约束条件来决定将 Pod 分配给哪个节点。
6. 健康检查：Kubernetes 可以通过定义健康检查来监控 Pod 中的容器的状态。如果容器不处于健康状态，Kubernetes 可以自动重启或重新创建 Pod。

Pod 是 Kubernetes 中最基本的部署单元，它提供了一种将相关容器组合在一起、共享资源和网络的方式。Pod 提供了一种逻辑上的抽象，使得容器可以作为一个整体进行管理和调度。

### 创建nginx的pod

```yaml
# nginx.yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
    - name: nginx-container
      image: nginx
```

###  运行pod

```sh
kubectl apply -f nginx.yaml        
# pod/nginx-pod created

kubectl port-forward nginx-pod 4000:80 #当这个命令在终端运行时才可以访问
# Forwarding from 127.0.0.1:4000 -> 80
# Forwarding from [::1]:4000 -> 80
```

## Deployment

Deployment 是一种控制器对象，用于定义和管理应用程序的部署和更新。

Deployment 提供了以下功能：

1. 声明式部署：使用 Deployment，你可以通过声明性的方式定义应用程序的期望状态，而不需要手动管理每个 Pod 的创建和更新。Kubernetes 控制器会根据 Deployment 的定义来创建和维护所需数量的 Pod 副本。
2. 滚动更新：Deployment 允许你以控制的方式进行应用程序的滚动更新。你可以指定要更新的容器镜像的新版本，然后 Deployment 会逐步替换旧的 Pod 副本，以确保应用程序的无缝更新。
3. 回滚功能：如果在滚动更新过程中发生问题，Deployment 具有回滚功能，可以将应用程序回滚到之前的稳定状态。这可以通过简单地将 Deployment 的配置回滚到之前的版本来实现。
4. 自动扩缩容：Deployment 可以根据应用程序的负载情况自动进行水平扩缩容。你可以配置 Deployment 的副本数量，Kubernetes 控制器会根据资源使用情况自动调整副本数量，以满足应用程序的需求。

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hellok8s-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hellok8s
  template:
    metadata:
      labels:
        app: hellok8s
    spec:
      containers:
        - image: guangzhengli/hellok8s:v1
          name: hellok8s-container
```



## k8s中的Replica Sets

在 Kubernetes（K8s）中，Replica Sets（副本集）是一种控制器对象，用于定义和管理一组具有相同配置的 Pod 副本。它可以确保指定数量的 Pod 副本在任何时候都在运行，并且可以自动进行水平扩展或收缩。

以下是一些关键特点和用途：

1. 副本数量控制：Replica Sets 允许你指定希望运行的 Pod 副本数量。Kubernetes 控制器会监控这些副本，并在需要时自动创建或销毁 Pod，以确保副本数量保持在所需的状态。
2. 自愈能力：如果某个 Pod 副本失败或被终止，Replica Sets 会自动创建新的副本来替代它，以确保指定的副本数量保持不变。
3. 水平扩展：通过增加 Replica Sets 中的副本数量，可以实现水平扩展应用程序。Kubernetes 控制器会自动创建新的 Pod 副本，以满足所需的副本数量。
4. 滚动更新：当你需要更新应用程序时，可以通过逐步替换旧的 Pod 副本来进行滚动更新。Replica Sets 可以控制新旧版本的 Pod 副本之间的过渡过程，从而实现无缝的应用程序更新。

要创建和管理一个 Replica Set，你需要使用 Kubernetes 的 YAML 文件来定义它的配置，包括副本数量、选择器和 Pod 模板等信息。然后，使用 `kubectl` 命令将该 YAML 文件部署到 Kubernetes 集群中。

下面是一个示例的 Replica Set 的 YAML 配置文件：

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: my-replicaset
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-container
        image: my-image:latest
```



# 参考文献

>minikube start: https://minikube.sigs.k8s.io/docs/start/
>
>https://github.com/guangzhengli/k8s-tutorials/tree/main#安装-kubectl
>
>https://colobu.com/2022/06/02/setup-a-k8s-cluster-with-minikube/#more