build/linux:
	GOARCH=amd64 GOOS=linux go build -trimpath -o ./webgames main.go

build/windows:
	GOARCH=amd64 GOOS=windows go build -trimpath -o ./webgames.exe main.go

run/linux:
	./webgames

compress/windows:
	rm -Rf /tmp/webgames && \
		mkdir -p /tmp/webgames && \
		cp -Rf webgames.exe templates static /tmp/webgames && \
		cd /tmp && zip -r $(PWD)/webgames.zip ./webgames && \
		cd $(PWD) && rm -Rf /tmp/webgames

deps/tidy:
	go mod tidy

deps/sync: deps/tidy
deps/cleanup: deps/tidy

test/unit:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

test/coverage:
	go tool cover -html=coverage.out

test/unit/coverage: test/unit test/coverage