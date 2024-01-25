curr:
	go build -ldflags "-s -w" -o build/goshare

all: curr linuxStatic linuxArm64 winAmd64

linuxStatic:
	@echo "[+] Building the static Linux version"
	@env GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/goshare.static

linuxArm64:
	@echo "[+] Building the Linux ARM64 version"
	@env GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o build/goshare.arm

winAmd64:
	@echo "[+] Building the Windows version"
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/goshare.exe

