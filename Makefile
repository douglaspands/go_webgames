build-linux:
    GOARCH=amd64 GOOS=linux go build -trimpath -o ./dist/webgames main.go

build-windows:
    GOARCH=amd64 GOOS=windows go build -trimpath -o ./dist/webgames.exe main.go
