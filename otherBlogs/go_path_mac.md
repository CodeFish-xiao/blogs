# Mac 设置 GOPATH和GOROOT无效
## 修改.bash_profile
```
vim ~/.bash_profile
```
## 添加如下内容
```
export GOPATH=/Users/xxx/go
export GOROOT=/usr/local/go
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin
```
## 使配置生效
```
source ~/.bash_profile
```
## 查看配置是否生效
```
go env
```
### 测试 go install
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```