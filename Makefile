curr:
	go build -ldflags "-s -w" -o goshare

linuxArm64:
	@echo "[+] Building the Linux ARM64 version"
	GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o goshare

winAmd64:
	@echo "[+] Building the Windows version"
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare.exe

