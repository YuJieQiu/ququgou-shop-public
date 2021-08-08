# 资源存储模块

####  描述
*该模块和资源上传相关，有上传本地和对接第三方平台(七牛云)* 

-----
#### 订单相关表
-  __resources__  
-  __resource_config__ 
    
-----
#### 功能:
    > 文件资源上传
    上传类型有Image、Video、Andio  
    
#### 配置文件信息:
    ` uploadconfigforimage:
       path: 路径
       type: 类型 1本地、2七牛云
       hostaddress: host
       limitsize: 限制传文件大小
       limittypes: 限制上传文件类型 [".jpg",".png",".gif",".jpeg"]
       qiniuconfigmodel: 七牛云相关配置
         accesskey: 
         secretkey: `
         
