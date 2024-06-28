# gin-swagger-demo

gin-swagger create multiple api documents in one project

go version >= 1.16

download demo
```
git clone https://github.com/lhl1115/gin-swagger-demo.git
cd gin-swagger-demo
go mod tidy
go run cmd/v1/main.go
or 
go run cmd/control/main.go
```
view swagger documents in browser
```
http://localhost:8081/swagger/v1/index.html
http://localhost:8081/swagger/control/index.html
```
create swagger documents
```
swag init -g  cmd/v1/main.go  -o ./docs/swag  --instanceName v1

swag init -g  ../../cmd/control/main.go  -d internal/control  -o ./docs/swag  --parseDependency --instanceName control
```


