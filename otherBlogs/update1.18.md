# 升级1.18遇坑记录（遇见一个更新一个）
今天正式发布了golang 1.18 版本，大家可以直接通过 https://go.dev/dl/ 获得最新版本。
虽然在代码层面不会有问题，但是自己环境搭建可以会遇到一些`小坑坑`
## 1.1：Goland debug 报错 ：
    Debugging programs compiled with go version go1.18 darwin/amd64 is not supported. Use go sdk for darwin/arm64.

###  1.1.1：原因：

旧版本的 https://github.com/go-delve/delve 是没有支持golang 1.18 debug的。在项目的Changelog可见：在1.8版本才支持的1.18，所以在本地升级完版本后需要升级该插件。
### 1.1.2：解决方法：
~~~
git clone https://github.com/go-delve/delve.git
cd delve/cmd/dlv/ 
go build 
go install
~~~

一般安装目录会在你的go安装目录上，或者是你的GOPATH/bin中，然后在Goland中点击: Help → Edit Custom Properties...

~~~
dlv.path=/usr/local/bin/dlv
~~~
保存重启，解决step over(F8) 直接运行DEBUG.

### 1.1.3 2022.03.17修正
不需要第二步添加该命令，即可以修复
~~~
dlv.path=/usr/local/bin/dlv
~~~