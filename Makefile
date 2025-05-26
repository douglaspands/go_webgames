build/linux:
    GOARCH=amd64 GOOS=linux go build -trimpath -o ./webgames main.go

build/windows:
    GOARCH=amd64 GOOS=windows go build -trimpath -o ./webgames.exe main.go

