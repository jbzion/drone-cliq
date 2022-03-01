
build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -v -a -o release/linux/amd64/server