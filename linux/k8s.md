# k8s

Kubernetes

## 安装

### kind

Kubernetes In Docker

```shell
# 安装 kind
go get -u sigs.k8s.io/kind
# 创建集群
kind create cluster
# 自定义集群配置
# kind create cluster --config kindcnf.yaml --name mycluster
# 测试集群是否可用
kind get clusters
# 将本地已经有的镜像加载到kind k8s 中
kind load docker-image nginx nginxz
# 查看节点
kubectl get nodes
# 执行shell
docker exec -ti <nodename> bash
# 查看安装镜像
crictl images
```

### kubectl

kubernetes-cli

```shell
# 安装
curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
mv kubectl ~/bin/
chmod +x ~/bin/kubectl
# 版本
kubectl version
# 集群信息
kubectl cluster-info
# 查看api版本
kubectl api-versions
# 查看 pods
kubectl get pods
# 查看所有的namespace
kubectl get pod --all-namespaces
# 查看副本控制器
kubectl get rc
# 查看所有节点
kubectl get nodes
kubectl get secrets
kubectl get pvc
# 运行的服务
kubectl get services
# 查看部署无状态应用
kubectl get deployment
# 创建副本
kubectl create -f rc-nginx.yaml
# 更新副本
kubectl replace -f rc-nginx.yaml
# 删除
kubectl delete deployment [name]
# 查看配置
kubectl get configmap
# 扩缩容
kubectl scale deployment eis-go --replicas 2
# 详情
kubectl describe pod
# 执行shell
kubectl exec [pid] /bin/sh
```

## java 突然关闭

如果java JAVA_OPT Xms 值和docker容器限制相同，内存满了的时候，java进程会突然关闭
