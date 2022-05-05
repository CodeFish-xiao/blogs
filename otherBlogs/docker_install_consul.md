# Docker 安装 Consul 做注册中心以及配置中心

## 前言

- 随着公司服务增多，以及每个服务至少有开发环境，测试环境，正式环境，所有的服务的配置文件需要在服务数*3的数量，管理难度上升。

- 之前的配置是每个服务默认读取yaml文件,然后通过环境变量进行覆盖配置，每次修改都需要重新部署一次环境变量
- 因为目前SAE环境限制，所以只能通过镜像方式部署应用。

## 目的

- 在sae服务中搭建自己的配置中心，完成所有服务的统一管理

## 安装

### 本地安装

1. 拉取镜像

`docker pull consul:latest`

2. 启动参数解释

```
–net=host docker参数, 使得docker容器越过了net namespace的隔离，免去手动指定端口映射的步骤
-server consul支持以server或client的模式运行, server是服务发现模块的核心, client主要用于转发请求
-advertise 将本机私有IP传递到consul
-retry-join 指定要加入的consul节点地址，失败后会重试, 可多次指定不同的地址
-client 指定consul绑定在哪个client地址上，这个地址可提供HTTP、DNS、RPC等服务，默认是>127.0.0.1
-bind 绑定服务器的ip地址；该地址用来在集群内部的通讯，集群内的所有节点到地址必须是可达的，>默认是0.0.0.0
allow_stale 设置为true则表明可从consul集群的任一server节点获取dns信息, false则表明每次请求都会>经过consul的server leader
-bootstrap-expect 数据中心中预期的服务器数。指定后，Consul将等待指定数量的服务器可用，然后>启动群集。允许自动选举leader，但不能与传统-bootstrap标志一起使用, 需要在server模式下运行。
-data-dir 数据存放的位置，用于持久化保存集群状态
-node 群集中此节点的名称，这在群集中必须是唯一的，默认情况下是节点的主机名。
-config-dir 指定配置文件，当这个目录下有 .json 结尾的文件就会被加载，详细可参考https://www.consul.io/docs/agent/options.html#configuration_files
-enable-script-checks 检查服务是否处于活动状态，类似开启心跳
-datacenter 数据中心名称
-ui 开启ui界面
-join 指定ip, 加入到已有的集群中
```

```
Usage:
  -advertise value
    	Sets the advertise address to use.
  -advertise-wan value
    	Sets address to advertise on WAN instead of -advertise address.
  -allow-write-http-from value
    	Only allow write endpoint calls from given network. CIDR format, can be specified multiple times.
  -alt-domain value
    	Alternate domain to use for DNS interface.
  -bind value
    	Sets the bind address for cluster communication.
  -bootstrap
    	Sets server to bootstrap mode.
  -bootstrap-expect value
    	Sets server to expect bootstrap mode.
  -check_output_max_size value
    	Sets the maximum output size for checks on this agent
  -client value
    	Sets the address to bind for client access. This includes RPC, DNS, HTTP, HTTPS and gRPC (if configured).
  -config-dir value
    	Path to a directory to read configuration files from. This will read every file ending in '.json' as configuration in this directory in alphabetical order. Can be specified multiple times.
  -config-file value
    	Path to a file in JSON or HCL format with a matching file extension. Can be specified multiple times.
  -config-format string
    	Config files are in this format irrespective of their extension. Must be 'hcl' or 'json'
  -data-dir value
    	Path to a data directory to store agent state.
  -datacenter value
    	Datacenter of the agent.
  -default-query-time value
    	the amount of time a blocking query will wait before Consul will force a response. This value can be overridden by the 'wait' query parameter.
  -dev
    	Starts the agent in development mode.
  -disable-host-node-id
    	Setting this to true will prevent Consul from using information from the host to generate a node ID, and will cause Consul to generate a random node ID instead.
  -disable-keyring-file
    	Disables the backing up of the keyring to a file.
  -dns-port value
    	DNS port to use.
  -domain value
    	Domain to use for DNS interface.
  -enable-local-script-checks
    	Enables health check scripts from configuration file.
  -enable-script-checks
    	Enables health check scripts.
  -encrypt value
    	Provides the gossip encryption key.
  -grpc-port value
    	Sets the gRPC API port to listen on (currently needed for Envoy xDS only).
  -hcl value
    	hcl config fragment. Can be specified multiple times.
  -http-port value
    	Sets the HTTP API port to listen on.
  -https-port value
    	Sets the HTTPS API port to listen on.
  -join value
    	Address of an agent to join at start time. Can be specified multiple times.
  -join-wan value
    	Address of an agent to join -wan at start time. Can be specified multiple times.
  -log-file value
    	Path to the file the logs get written to
  -log-json
    	Output logs in JSON format.
  -log-level value
    	Log level of the agent.
  -log-rotate-bytes value
    	Maximum number of bytes that should be written to a log file
  -log-rotate-duration value
    	Time after which log rotation needs to be performed
  -log-rotate-max-files value
    	Maximum number of log file archives to keep
  -max-query-time value
    	the maximum amount of time a blocking query can wait before Consul will force a response. Consul applies jitter to the wait time. The jittered time will be capped to MaxQueryTime.
  -node value
    	Name of this node. Must be unique in the cluster.
  -node-id value
    	A unique ID for this node across space and time. Defaults to a randomly-generated ID that persists in the data-dir.
  -node-meta key:value
    	An arbitrary metadata key/value pair for this node, of the format key:value. Can be specified multiple times.
  -non-voting-server
    	(Enterprise-only) DEPRECATED: -read-replica should be used instead
  -pid-file value
    	Path to file to store agent PID.
  -primary-gateway value
    	Address of a mesh gateway in the primary datacenter to use to bootstrap WAN federation at start time with retries enabled. Can be specified multiple times.
  -protocol value
    	Sets the protocol version. Defaults to latest.
  -raft-protocol value
    	Sets the Raft protocol version. Defaults to latest.
  -read-replica
    	(Enterprise-only) This flag is used to make the server not participate in the Raft quorum, and have it only receive the data replication stream. This can be used to add read scalability to a cluster in cases where a high volume of reads to servers are needed.
  -recursor value
    	Address of an upstream DNS server. Can be specified multiple times.
  -rejoin
    	Ignores a previous leave and attempts to rejoin the cluster.
  -retry-interval value
    	Time to wait between join attempts.
  -retry-interval-wan value
    	Time to wait between join -wan attempts.
  -retry-join value
    	Address of an agent to join at start time with retries enabled. Can be specified multiple times.
  -retry-join-wan value
    	Address of an agent to join -wan at start time with retries enabled. Can be specified multiple times.
  -retry-max value
    	Maximum number of join attempts. Defaults to 0, which will retry indefinitely.
  -retry-max-wan value
    	Maximum number of join -wan attempts. Defaults to 0, which will retry indefinitely.
  -segment value
    	(Enterprise-only) Sets the network segment to join.
  -serf-lan-allowed-cidrs value
    	Networks (eg: 192.168.1.0/24) allowed for Serf LAN. Can be specified multiple times.
  -serf-lan-bind value
    	Address to bind Serf LAN listeners to.
  -serf-lan-port value
    	Sets the Serf LAN port to listen on.
  -serf-wan-allowed-cidrs value
    	Networks (eg: 192.168.1.0/24) allowed for Serf WAN (other datacenters). Can be specified multiple times.
  -serf-wan-bind value
    	Address to bind Serf WAN listeners to.
  -serf-wan-port value
    	Sets the Serf WAN port to listen on.
  -server
    	Switches agent to server mode.
  -server-port value
    	Sets the server port to listen on.
  -syslog
    	Enables logging to syslog.
  -ui
    	Enables the built-in static web UI server.
  -ui-content-path value
    	Sets the external UI path to a string. Defaults to: /ui/ 
  -ui-dir value
    	Path to directory containing the web UI resources.
```



3. 启动单节点

`docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul agent -server -bootstrap-expect=1 -ui -bind=0.0.0.0 -client=0.0.0.0`

4. 建立集群

   1. 查看单节点ip
   2. 开启第二节点

   3. 开启第三节点

5. 查看集群地址



## 参考

- [consul文档](https://www.consul.io/docs)
- 

