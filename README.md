# go-gin-blog-eddycjy

## 引言

无意中发现了这本书，[Go 语言编程之旅](https://golang2.eddycjy.com/)，
这本书中的 [示例代码](https://github.com/go-programming-tour-book) 地址也找到了。

书的作者 煎鱼 的博客地址是 [煎鱼博客](https://eddycjy.com/) ，
博客的 [Github地址](https://github.com/eddycjy/blog) ，
作者的 [Github地址](https://github.com/eddycjy)。

里面看到了 [Gin框架示例](https://github.com/eddycjy/go-gin-example)，
相应的一系列 [博文](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)，
里面说是 [Gin实践](https://segmentfault.com/a/1190000013297625) 的连载，在 segmentfault 网站上，
另外在作者博客中也记录了这个系列，不过内容看着不一样， [gin](https://eddycjy.com/tags/gin/)，
看着内容是作者博客 [Go语言入门系列](https://eddycjy.com/go-categories/) 中的一部分。
对比看了下，segmentfault 上那个是2018年写的，好多已经过时了，就不要看了。

作者还有 [Go 语言进阶之旅](https://golang1.eddycjy.com/)。

这是跟着作者学习的一个示例项目。

## 数据库

创建库：
```
CREATE DATABASE
IF
	NOT EXISTS blog_service DEFAULT CHARACTER 
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
```

标签表：
```
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';
```

文章表：
```
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '文章简述',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面图片地址',
  `content` longtext COMMENT '文章内容',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0 为禁用、1 为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';
```

文章标签关联表：
```
CREATE TABLE `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT '文章 ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '标签 ID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';
```

权限表：
```
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='认证管理';

INSERT INTO `blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES (1, 'abc', 'go-gin-blog', 0, 'author', 0, '', 0, 0);
```

## 命令

### 获取包

基础包：
> go get -u github.com/gin-gonic/gin@v1.9.1

配置处理：
> go get -u github.com/spf13/viper@v1.4.0

数据操作：
> go get -u github.com/jinzhu/gorm@v1.9.12

日志操作：
> go get -u gopkg.in/natefinch/lumberjack.v2

Swagger 相关联的库：
> go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
>
> go get -u github.com/swaggo/gin-swagger@v1.2.0 
>
> go get -u github.com/swaggo/files
>
> go get -u github.com/alecthomas/template

参数校验：
> go get -u github.com/go-playground/validator/v10

JWT处理：
> go get -u github.com/dgrijalva/jwt-go@v3.2.0

邮件操作：
> go get -u gopkg.in/gomail.v2

令牌桶：
> go get -u github.com/juju/ratelimit@v1.0.1

OpenTracing API和Jaeger Client：
> go get -u github.com/opentracing/opentracing-go@v1.1.0
> 
> go get -u github.com/uber/jaeger-client-go@v2.22.1


### tag标签

新增一条：
> curl -X POST http://127.0.0.1:8000/api/v1/tags -F name=Go -F created_by=one -F state=1

获取指定：
> curl -X GET http://127.0.0.1:8000/api/v1/tags/1

更新指定：
> curl -X PUT http://127.0.0.1:8000/api/v1/tags/1 -F name=PHP -F modified_by=two -F state=0

删除指定：
> curl -X DELETE  http://127.0.0.1:8000/api/v1/tags/1

获取列表（执行时总是报错，原来是win10中curl包的问题，浏览器访问没问题）：
> curl -X GET http://127.0.0.1:8000/api/v1/tags?name=Go&state=1&page=2&page_size=2

### Swagger文档

生成文档：
> swag init

### JWT认证

获取token：
> curl -X POST http://127.0.0.1:8000/auth  -F app_key=abc  -F app_secret=go-gin-blog

访问验证：
> curl -X GET http://127.0.0.1:8000/api/v1/tags?token={token}

### 其他

应用配置后运行命令，如：
> go run main.go -port=8000 -mode=release -config=configs/



