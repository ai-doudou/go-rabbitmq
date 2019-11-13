 # lib
 
 ``` 
 go get -u -v  github.com/streadway/amqp
 ```
 
 
 --- 
 
# 如果本机安装过rabbitmq,按照下列步骤停止服务
 
 ``` 
#sudo service rabbitmq-server start

#sudo service rabbitmq-server stop

#sudo service rabbitmq-server status

#禁止开启自启动
sudo systemctl disable rabbitmq-server.service
 
 ```
 ---
 
# Docker install

> 获取镜像

``` 
docker pull rabbitmq:3.7.7-management
```
 
> 启动RabbitMQ
``` 

➜  mkdir -pv rabbitmq/data

➜  rabbitmq pwd
/home/jinchunguang/docker/rabbitmq
➜  rabbitmq tree
.
└── data

1 directory, 0 files
➜  rabbitmq docker run -d --hostname rabbitmq-host --name rabbitmq-srv -p 5672:5672  -p 15672:15672 -p 25672:25672  -v `pwd`/rabbitmq/data:/var/lib/rabbitmq rabbitmq:3.7.7-management
22698826832de91b13e82e27c54d9c5abc58f36471fb0e730333ede938955e65
➜  rabbitmq 
```

默认密码为guest/guest

参数

``` 
说明：

-d 后台运行

--name 容器名；

-p 指定服务运行的端口（5672：应用访问端口 15672：控制台Web端口号 25672:集群需要） 

-v 映射目录或文件

--hostname  主机名
...
```

具体查看

https://hub.docker.com/_/rabbitmq

> 具体操作

查看日志
``` 
 rabbitmq docker logs rabbitmq-srv

  # 忽略
 ...
  ##  ##
  ##  ##      RabbitMQ 3.7.7. Copyright (C) 2007-2018 Pivotal Software, Inc.
  ##########  Licensed under the MPL.  See http://www.rabbitmq.com/
  ######  ##
  ##########  Logs: <stdout>

              Starting broker...
2019-11-13 08:28:15.895 [info] <0.197.0> 
 node           : rabbit@rabbitmq-host
 home dir       : /var/lib/rabbitmq
 config file(s) : /etc/rabbitmq/rabbitmq.conf
 cookie hash    : Yc3Dd6JhyepTOuAgSMQfkg==
 log(s)         : <stdout>
 database dir   : /var/lib/rabbitmq/mnesia/rabbit@rabbitmq-host
 # 忽略
```





docker stop $(docker ps -qa);docker rm $(docker ps -qa);



