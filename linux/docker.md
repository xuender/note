# docker

## 安装

```shell
sudo apt install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
```

目前 20.04 还没有国内 docker 源，暂时用 18.04 版本。

```shell
sudo add-apt-repository \
  "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
  bionic \
  stable"
```

安装

```shell
sudo apt install docker-ce docker-ce-cli containerd.io
```

测试

```shell
sudo docker run hello-world
```

免sudo运行

```shell
sudo usermod -aG docker ${USER}
sudo reboot
```

## docker-compose

[install](https://docs.docker.com/compose/install)