curr:
	go build -ldflags "-s -w" -o goshare

arm:
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o goshare

