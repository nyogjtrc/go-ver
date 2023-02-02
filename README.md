# go-ver
application version in golang


go build command example
```
$ go build -v -ldflags "-X github.com/nyogjtrc/go-ver.Version=$(git describe --tags --dirty --match 'v*' || echo 'dev') \
    -X github.com/nyogjtrc/go-ver.BuildTime=$(date +%Y/%m/%dT%H:%M:%S) \
    -X github.com/nyogjtrc/go-ver.Commit=$(git rev-parse --short HEAD) ./main.go"
```
