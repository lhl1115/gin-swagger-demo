# gin-swagger-demo
使用gin-swagger在一个项目中生成多个swagger文档

go 版本 >= 1.16


下载demo
```
git clone github.com/lhl1115/gin-swagger-demo
cd gin-swagger-demo
go mod tidy
go run cmd/v1/main.go
or 
go run cmd/control/main.go
```
访问浏览器查看Swagger文档
```
http://localhost:8081/swagger/v1/index.html
http://localhost:8081/swagger/control/index.html
```

生成swagger文档
```
swag init -g  cmd/v1/main.go  -o ./docs/swag  --instanceName v1

swag init -g  ../../cmd/control/main.go  -d internal/control  -o ./docs/swag  --parseDependency --instanceName control
```


