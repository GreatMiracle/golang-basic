# golang-basic
docker run --rm -v "C:\Users\admin\Documents\improve-golang\golang-basic\web:/src" -w /src kjconroy/sqlc generate

go run main.go

go mod tidy

cài viper::
go get github.com/spf13/viper

syntax valid request params
https://github.com/go-playground/validator

cài jwt::
go get github.com/golang-jwt/jwt/v5

cài paseto::
go get github.com/o1egl/paseto