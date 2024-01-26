# Project Config

## init

```
$ go mod init github.com/jyp90/git-fw
```

## install gin 

```
$ go install github.com/gin-gonic/gin@latest
```

## install logr
```
$ go get "github.com/go-logr/logr"
```

## download prometheus using go mod
```
$ go mod download github.com/prometheus-operator/prometheus-operator@v0.57.0
```

## run

```
$ go run main.go | go build -o app && ./app
```