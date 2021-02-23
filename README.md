# mmwiki-auth-smtp

[mm-wiki](https://github.com/phachon/mm-wiki) SMTP统一认证
![](https://github.com/chennqqi/mmwiki-auth-smtp/raw/main/doc.png)

## roadmap

- 增加自定义用户配置

## 使用与配置	

1. docker启动或者自主编译

2. 默认路径为/smtplogin

   在mm-wiki中【系统】->【配置管理】->【登录认证】
   认证URL http://<IP:port>/smtp/login

   路径也可以通过制定参数path来修改
   端口默认8080

   ** 附加参数需要配置为邮箱后缀 **


