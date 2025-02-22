# 安装 Dragonfly CDN

本文档阐述如何安装并启动 Dragonfly CDN。

## 部署方式

用下列方式之一部署 CDN：

- 通过 Docker 部署
- 直接在物理机上部署：推荐用于生产用途

## 环境要求

使用 Docker 部署时，以下条件必须满足：

所需软件 | 版本要求
---|---
Git|1.9.1+
Docker|1.12.0+

直接在物理机上部署时，以下条件必须满足：

所需软件 | 版本要求
---|---
Git|1.9.1+
Golang|1.12.x
Nginx|0.8+

## 使用 Docker 部署

### 获取 CDN 镜像

您可以直接从 [DockerHub](https://hub.docker.com/) 获取 CDN 镜像。

1. 获取最新的 CDN 镜像

```sh
docker pull dragonflyoss/cdn
```

或者您可以构建自己的 CDN 镜像

1. 获取 Dragonfly 的源码

```sh
git clone https://github.com/dragonflyoss/Dragonfly2.git
```

2. 打开项目文件夹

```sh
cd Dragonfly2
```

3. 构建 CDN 的 Docker 镜像

```sh
TAG="2.0.0"
make docker-build-cdn D7Y_VERSION=$TAG
```

4. 获取最新的 CDN 镜像 ID

```sh
docker image ls | grep 'cdn' | awk '{print $3}' | head -n1
```

### 启动 cdn

**注意：** 需要使用上述步骤获得的 ID 替换 ${cdnDockerImageId}。

```sh
docker run -d --name cdn --restart=always -p 8001:8001 -p 8003:8003 -v /home/admin/ftp:/home/admin/ftp ${cdnDockerImageId} 
--download-port=8001
```

CDN 部署完成之后，运行以下命令以检查 Nginx 和 **cdn** 是否正在运行，以及 `8001` 和 `8003` 端口是否可用。

```sh
telnet 127.0.0.1 8001
telnet 127.0.0.1 8003
```

## 在物理机上部署

### 获取 CDN 可执行文件

1. 下载 Dragonfly 项目的压缩包。您可以从 [github releases page](https://github.
   com/dragonflyoss/Dragonfly2/releases) 下载一个已发布的最近版本
   
```sh
version=2.0.0
wget https://github.com/dragonflyoss/Dragonfly2/releases/download/v$version/Dragonfly2_$version_linux_amd64.tar.gz
```

2. 解压压缩包

```bash
# Replace `xxx` with the installation directory.
tar -zxf Dragonfly2_2.0.0_linux_amd64.tar.gz -C xxx
```

3. 把 `cdn` 移动到环境变量 `PATH` 下以确保您可以直接使用 `cdn` 命令

或者您可以编译生成自己的 CDN 可执行文件。

1. 获取 Dragonfly 的源码

```sh
git clone https://github.com/dragonflyoss/Dragonfly2.git
```

2. 打开项目文件夹

```sh
cd Dragonfly2
```

3. 编译源码

```sh
make build-cdn && make install-cdn
```

### 启动 cdn

```sh
cdnHomeDir=/home/admin
cdnDownloadPort=8001
cdn --port=8003 --download-port=$cdnDownloadPort
```

### 启动 file server

您可以在满足以下条件的基础上用任何方式启动 file server：

- 必须挂载在前面步骤中已经定义的目录 `${cdnHomeDir}/ftp` 上。
- 必须监听前面步骤中已经定义的 `cdnDownloadPort` 端口。

以 nginx 为例：

1. 将下面的配置添加到 Nginx 配置文件中

```conf
server {
# Must be ${cdnDownloadPort}
listen 8001;
location / {
 # Must be ${cdnHomeDir}/ftp
 root /home/admin/ftp;
}
}
```

2. 启动 Nginx.

```sh
sudo nginx
```

CDN 部署完成之后，运行以下命令以检查 Nginx 和 **cdn** 是否正在运行，以及 `8001` 和 `8003` 端口是否可用。

```sh
telnet 127.0.0.1 8001
telnet 127.0.0.1 8003
```
