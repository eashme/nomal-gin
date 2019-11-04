# Swagger UseAge

+ 使用swagger-gin 与 gin框架进行结合
+ 编译go build -o swag swagger目录/cmd/swag/main.go 文件,生成 swagger命令行工具,将其放入GOBIN目录下
+ 需要使用swag init 命令生成 docs 文件夹下的所有的文件
+ 使用前需要在项目起始位置 导入 import _ "nomal-gin/docs"
+ 添加swagger文档路由 g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
+ 项目运行后访问/swagger/index.html即可访问swagger文档

### swagger文档使用方法
- 项目全局定义
```
// @title 项目名称
// @version 1.0
// @description areas 项目描述
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email st5983@outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8872
// @BasePath /v1
```
> 主要修改有三点
> 1. title : 项目名称
> 2. host: 请求路径
> 3. BasePath: 基础路径 （host+BasePath配合api的路由，组成访问路径，即 host:port/v1/api.path）

- 路由的定义
```
// @Summary 用户接口
// @Description 根据用户xxxx
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData  int  true "姓名"
// @Param age formData  int  true "年龄"
// @Success 200 {object} docs.ResponseOk "success" 
// @Failure 400 {object} docs.ResponseInvalidParam "bad parameters"
// @Failure 500 {object} docs.ResponseServerErr "internal server error"
// @Router /users [get]
``` 
> 1. Summary: 接口接口名称
> 2. Tags: swagger靠这个参数分类接口
> 3. Accept: 通过表单来传递时使用application/x-www-form-urlencoded， 其他为json
> 4. Produce: 一般为json
> 5. Param: 参数依次为：参数名 传递方式 参数类型 是否一定存在 注释
>> 参数类型一般有：int number string (number对应float类型)
>> 传递方式有：path query formData (path即参数在路径中，如 user/1; query为传统传参，如user?id=1; formData为表单数据)
> 6. Success/Failure 成功或者失败返回状态码， 以及返回的example，example在 /docs/swag_model.go中定义。如果返回值是结构体用`{object}`，如果是数组，用`{array}`
> 7. Router: 定义路由，如果传递方式是path，参数用{}包含，如/user/{id};如果是query或者formData，则不用展示参数。第二个参数为restful对应的类别，如get post delete put patch

3.其他
> 1. 在gin中，如果参数在路径中，即path，用Param()来获取，如果是表单数据，用PostForm()或者Bind()一个结构体，如果是传统传参,即query,用Query()来获取
> 2. 每次更改配置文件，需要使用`swag init`来初始化