install

go get -u github.com/wesovilabs/beyond

on project root

CGO_ENABLED=0 beyond --verbose --work run examples/http/router/echo/aop/main.go 