[TOC]
## 部署相关
### 部署chat
```shell
cd /www/wwwroot/chatgpt-wechat  #进入项目目录
git pull #拉取最新代码
vi ./service/chat/api/etc/chat-api.yaml #修改配置文件，如果配置无变动，跳过这一步
sudo docker-compose build #打包
sudo docker-compose down #关闭服务
sudo docker-compose up -d #启用服务
docker logs --tail 500 容器名 #查看接口日志
```

### docker
清理磁盘 `docker system prune`

## 框架相关
### 生成model
```
cd service/user/model
goctl model mysql ddl -src user.sql -dir . -c
```

### 生成api
```
cd chat/service/chat/api
goctl api go -api chat.api -dir .
```

## git相关
### 将 fork 的仓库合并到自己的仓库
#### 在本地电脑上，将自己仓库的代码克隆到本地：
```
git clone git@github.com:<your-username>/<a>.git
git@github.com:chy4pro/chatgpt-wechat.git

cd <a>
git remote add upstream git://github.com/<original-author>/<b>.git
git remote add upstream git@github.com:whyiyhw/chatgpt-wechat.git


git fetch upstream
git merge upstream/main


git push origin main

```


### 本地创建新分支：

1. 在本地的main分支上，使用以下命令创建一个名为f-scale2的新分支：
```
git checkout -b f-scale2
```
2. 确认当前处于新创建的分支中：
```
git branch
```
3. 在远程仓库上，使用以下命令创建一个名为f-scale2的新分支：
```
git push origin f-scale2
```
4. 确认新的远程分支已经创建：
```
git branch -r
```
5. 将本地新分支与远程新分支关联：
```
git push --set-upstream origin f-scale2
```

## 需要开放的端口
### chat
- 8888
- 8887

### redis
- 46379

### mysql
- 43306

### script
- 9997

### gva-web
- 8080
## gva-server
- 8886
## 启动nginx
`systemctl start nginx`

## 启动squid代理
sudo systemctl restart squid

