# gin-swagger-demo
[简体中文](README.md)

Generate multiple Swagger documents in a single project using [gin-swagger](https://github.com/swaggo/gin-swagger)

Go version >= 1.16

## Requirements and Results
In a gin project with multiple versions, such as v1 and v2, or when there are subprojects, generating only one document can lead to confusion and poor readability. Therefore, we generate a separate document for each version or subproject and access them through different routes.

![img.png](img/img.png)
![img_1.png](img/img_1.png)

## How to use the demo
Download the demo:
```bash
git clone https://github.com/lhl1115/gin-swagger-demo.git
cd gin-swagger-demo
go mod tidy
go run cmd/v1/main.go
or
go run cmd/control/main.go
or
go run --tags "doc" ./cmd/v2
```

Access the Swagger documents in your browser:

```bash
http://localhost:8081/swagger/v1/index.html
http://localhost:8081/swagger/v2/index.html
http://localhost:8081/swagger/control/index.html
```

Generate Swagger documents:

```bash
// Generate all documents
swag init -g cmd/v1/main.go -o ./docs/swag --instanceName v1

// Generate v2 documents
swag init -g ../../cmd/v2/main.go -d internal/v2 -o ./docs/swag --parseDependency --instanceName v2

// Generate control documents
swag init -g ../../cmd/control/main.go -d internal/control -o ./docs/swag --parseDependency --instanceName control
```

## Ignoring swag during packaging

In version v2, we added +build doc to swag.go, so we can ignore the swag documents during packaging to reduce the package size.

```bash
// Package size approximately 11M without swag documents
go build -o v2.exe ./cmd/v2

// Package size approximately 26M with swag documents
go build -o v2_doc.exe -tags "doc" ./cmd/v2
```