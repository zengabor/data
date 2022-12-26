build:
	go build

test:
	go test ./...

release:
	git tag v1.2.0
	GOPROXY=proxy.golang.org go list -m github.com/zengabor/data@v1.2.0
