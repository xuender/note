# nginx

## 安装

```shell
sudo apt install nginx
```

## 端口转发

编辑配置文件

```shell
cd /etc/nginx/sites-available
sudo vi domain.com.conf
```

域名对应的配置

```conf
server {
        listen 80;
        server_name  domain.com;
        location / {  
            proxy_pass   http://127.0.0.1:8080;  
            index  index.html index.htm;  
        }
}
```

启用配置

```shell
sudo ln -s /etc/nginx/sites-available/domain.com.conf /etc/nginx/sites-enabled/
```

测试配置

```shell
sudo service nginx configtest
```

重新加载配置

```shell
sudo service nginx reload
```
