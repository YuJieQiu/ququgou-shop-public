# 用户端小程序 API

### 接口文档

### swagger

- 安装 go-swagger 插件
- 自定义配置插件等模版
- 在 setting.json 文件中配置如下
  ```
  "swagger.tpl": "// @Summary 简单描述\n// @Description 描述\n// @Accept json\n// @Produce json\n// @Param body body 参数 true \"body参数\"\n// @Success 200 {object} Response{data=参数} \"ok\" \"返回信息\"\n// @Failure 400 {object} Response \"错误\" \n// @Failure 401 {object} Response \"错误\"\n// @Failure 500 {object} Response \"错误\"\n// @Router /user/person/login [post]\n" 
  ``` 

