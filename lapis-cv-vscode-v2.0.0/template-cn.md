# 段晨曦

> <span class="icon">&#xe60f;</span> `15626542043`&emsp;&emsp;
> <span class="icon">&#xe7ca;</span> `dddcx.d@qq.com`&emsp;&emsp;
> <span class="icon">&#xe600;</span> [github](https://github.com/CharlesDCXX?tab=repositories)

<img class="avatar" src="https://avatars.githubusercontent.com/u/583231?v=4">

## &#xe80c; 教育经历
<div class="entry-title">
    <h5> 2020.09 - 2023.06 &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 中国矿业大学 - 硕士 - 软件工程技术 </h5> 
</div>
<div class="entry-title">    
    <h5> 2014.09 - 2018.06 &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp; 河南科技学院 - 本科 - 数学与应用数学 </h5> 
</div>

## &#xe618; 专业技能
1. 微服务与架构设计：具备微服务架构落地经验，熟悉服务发现、熔断限流、配置中心、链路追踪与日志采集，能够独立设计并优化分布式系统。

2. Golang 开发：精通 Goroutine 并发模型与 Channel 通信机制，熟练掌握 Gin、gRPC 等框架，擅长高性能网络编程与服务端开发。

3. 云原生与运维：熟悉 Docker、Kubernetes、Helm 等工具，具备容器化部署、微服务编排与 CI/CD 实践经验。实际在 阿里云、华为云环境中部署与运维云平台服务，熟悉其网络、存储与监控体系。

4. 数据库与缓存：熟练使用 MySQL / PostgreSQL，具备索引优化、SQL 调优经验；掌握 Redis 高可用与缓存设计模式。

5. 消息与中间件：了解 Kafka等消息队列及其在高并发、解耦与削峰填谷场景下的应用。

6. 计算机基础：扎实的操作系统、计算机网络与数据结构与算法功底，能够结合底层原理进行性能优化与问题排查。

7. 工具链：熟悉 Git、Linux、Prometheus + Grafana 等监控工具，能够搭建监控体系并进行性能追踪与优化。

## &#xe618; 工作经验

<div alt="entry-title">
    <h3>平台开发工程师 - 深圳市海雀科技有限公司</h3> 
    <p>2023.07 - 至今</p>
</div>

作为核心开发成员参与了海雀摄像头物联网平台的全栈开发与架构设计。

- 负责用户侧与设备侧接口的研发与维护，保障数据交互的高效与可靠。
- 构建并优化监控体系以跟踪微服务运行指标，确保云平台的稳定性与高可用性。
- 参与云平台的持续优化与运维改进，提升整体服务质量与用户体验。


## &#xe635; 项目经历

<div class="entry-title">
    <h3>跨域RPC调用服务</h3>
</div>

跨节点调用容易因网络延迟和抖动导致质量问题，因此主导并落地跨域高可用方案：
- 设计并实现多域部署的 API 网关双向代理模式，支持网关之间互相调用，有效提升服务稳定性。

- 实现跨域调用的服务发现机制，保证请求可在多节点间无缝流转，提升系统灵活性与容错能力。

- 简化微服务调用路径，业务方仅需调用本地网关即可完成跨节点无感互通，降低接入复杂度。

- 引入请求队列化机制，优化网关吞吐量，在高并发场景下显著提升系统整体性能。


<div class="entry-title">
    <h3>gRPC封装服务</h3> 
</div>

在使用grpc时，由于grpcAPI无法满足公司微服务一些需求，因此主导在grpc上封装一层适用接口，实现以下功能：

- 客户端直接通过服务名及方法名调用服务，无需IP地址。
- 从etcd发现多个服务端pod地址，对多个服务端进行无感轮询访问。
- 将鉴权封装在RPC调用层，业务端与服务端无感，只需专注业务本身。
- 
