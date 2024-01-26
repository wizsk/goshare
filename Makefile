# .PHONY: all clean

all: setup curr linuxStatic linuxArm64 winAmd64 clean

setup:
	sed -i 's/const debug = !false/const debug = false/' main.go

clean:
	sed -i 's/const debug = false/const debug = !false/' main.go

curr:
	go build -ldflags "-s -w" -o build/goshare

linuxStatic:
	@echo "[+] Building the static Linux version"
	@env GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/goshare.static

linuxArm64:
	@echo "[+] Building the Linux ARM64 version"
	@env GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o build/goshare.arm

winAmd64:
	@echo "[+] Building the Windows version"
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/goshare.exe

# This target is always executed, even if there are errors in previous targets.
.PHONY: always
always: clean
