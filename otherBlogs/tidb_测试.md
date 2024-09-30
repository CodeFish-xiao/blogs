# TiDB入门
## 安装TiDB

本文在Mac环境下安装测试集群，不适用于生产环境。


TiDB 是一个分布式系统。最基础的 TiDB 测试集群通常由 2 个 TiDB 实例、3 个 TiKV 实例、3 个 PD 实例和可选的 TiFlash 实例构成。通过 TiUP Playground，可以快速搭建出上述的一套基础测试集群，步骤如下：

1. 下载并安装 TiUP。
```shell
curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
```
安装完成后会提示如下信息：

```shell
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 7238k  100 7238k    0     0  7979k      0 --:--:-- --:--:-- --:--:-- 8015k
WARN: adding root certificate via internet: https://tiup-mirrors.pingcap.com/root.json
You can revoke this by remove /Users/codefish/.tiup/bin/7b8e153f2e2d0928.root.json
Successfully set mirror to https://tiup-mirrors.pingcap.com
Detected shell: zsh
Shell profile:  /Users/codefish/.zshrc
/Users/codefish/.zshrc has been modified to add tiup to PATH
open a new terminal or source /Users/codefish/.zshrc to use it
Installed path: /Users/codefish/.tiup/bin/tiup
===============================================
Have a try:     tiup playground
===============================================
```
2. 声明全局环境变量。
- TiUP 安装完成后会提示 Shell profile 文件的绝对路径。在执行以下 source 命令前，需要将 ${your_shell_profile} 修改为 Shell profile 文件的实际位置。
```shell
source ${your_shell_profile}
```
3. 