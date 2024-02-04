build:clean
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	go mod download -x
	GOARCH=amd64 GOOS=linux go build -o nemoexp

clean:
	rm -rf nemoexp

init:
	go mod init

tidy:
	go mod tidy -v
