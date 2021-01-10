# 1 启动
## 1.1 后端 
```
$ cd aping-blog/blog-back
$ ./run.sh
```

## 1.2 前端
```
$ cd aping-blog/blog-front
$ yarn install
$ yarn start
```
# 2 HTTPS
## 2.1 SSL 证书
```
# 安装Let'sEncrypt的certbot工具
$ sudo yum install certbot

# 停止Nginx服务
$ sudo killall nginx

# 申请证书
$ sudo certbot certonly --standalone -n -m your-email@example.com --$ agree-tos -d YourAwesome.Domain

# 证书自动续期
$ yum install python-certbot-nginx -y
$ sed -i "s/renew/renew --nginx/g" /usr/lib/systemd/system/certbot-$ renew.service
$ systemctl daemon-reload
$ systemctl start certbot-renew.service
```
## 2.2 nginx 配置
```
server {
    listen 443 ssl default_server;
    server_name YourAwesome.Domain;
    server_tokens off;

    keepalive_timeout 5;

    ssl_certificate "/etc/letsencrypt/live/YourAwesome.Domain/fullchain.pem";
    ssl_certificate_key "/etc/letsencrypt/live/YourAwesome.Domain/privkey.pem";
    ssl_session_cache shared:SSL:1m;
    ssl_session_timeout  10m;
    ssl_protocols TLSv1.2;
    ssl_ciphers ECDHE-RSA-AES256-GCM-SHA384:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;

    root xxx;
    index index.php index.html;

    access_log logs/blog.log combinediox;
    error_log logs/blog.error.log;

    location ~* ^/back/(.*) {
        rewrite ^/back(.*)$ $1 break;

        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";

        proxy_pass   http://127.0.0.1:9000;
    }

    location / {
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";

        proxy_pass   http://127.0.0.1:8080;
    }
}

server {
    listen       80;
    server_name  YourAwesome.Domain;
    return 301 https://$host$request_uri;
}
```


