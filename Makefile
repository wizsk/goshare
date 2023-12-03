arm:
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o goshare

curr: # current
	go build -ldflags "-s -w" -o goshare
