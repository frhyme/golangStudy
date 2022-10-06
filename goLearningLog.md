## 20220914

```sh
go version
# go1.17.5 darwin/amd64
# run go
go run hello.go
go --help
```

## 20220925

```sh
# .go file을 기계어로 변형한다. 아래와 같이 옵션을 주지 않고 작성할 경우에는 현재 시스템 상에서 사용할 수 있는 기계어가 만들어진다.
# hello 라는 실행 파일이 만들어지며, 당연하지만 ./hello 로 바로 실행 가능하다.
go build hello.go
```

- `go run aa.go` 등을 통해 해당 go file을 실행하기 위해서는 해당 파일 내에 package가`package main`으로 작성되어 있어야 한다.

```
package command-line-arguments is not a main package
```
