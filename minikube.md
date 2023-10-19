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

## minikube

minikube创建pod并暴露pod的端口号
container (容器) 的本质是进程，而 pod 是管理这一组进程的资源。
![image](./resource/pod关系.png)

