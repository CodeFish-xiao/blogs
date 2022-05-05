## 在Linux环境安装Goalng

## 我的配置

## 安装步骤

1. 下载对应的版本

   打开golang下载官网，选择对应的版本信息复制下载链接：

   ```shell
   # 我喜欢把文件都下载到data文件夹
   cd /data
   # 下载对应安装包
   wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
   ```

2. 解压到对应文件夹

   根据官方意见解压到对应的文件夹

   ```shell
   # 解压到/usr/local文件夹下
   tar -C /usr/local -zxvf  go1.18.1.linux-amd64.tar.gz
   ```

3. 添加到环境变量中

   1. 使用vim编辑环境变量

      ```shell
      vim /etc/profile
      ```

   2. 在最后一行添加

      ```shell
      export GOROOT=/usr/local/go
      export PATH=$PATH:$GOROOT/bin
      ```

   3. 保存后source一下

      ```shell
      source /etc/profile
      ```

4. 检查是否安装成功

   ```shell
   go version
   ```

   出现对应的版本即安装成功

