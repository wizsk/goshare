PACKAGE_NAME = goshare
VERSION = 4.0
ARCHITECTURE = amd64
EXEC_DIR = build/
MAINTAINER = https://github.com/wizsk
DESCRIPTION =  file share // server written in golang

# all: setup curr linuxStatic linuxArm64 winAmd64 clean

# release:
# 	sed -i 's/const debug = !false/const debug = false/' main.go
#
# debug:
# 	sed -i 's/const debug = false/const debug = !false/' main.go

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


all: build
check: format vet test

build: clean update tidyup format vet test
	@echo
	@echo "[+] Version: $(VERSION)"
	@echo

	@mkdir -p $(EXEC_DIR)

	@echo "[+] Building the Linux version"
	@env GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare

	@echo "[+] Packaging the Linux version"
	@tar -czvf $(EXEC_DIR)goshare_Linux_x86_64.tar.gz -C $(EXEC_DIR) goshare > /dev/null
	# @sha256sum $(EXEC_DIR)goshare_Linux.tar.gz

	@echo "[+] Building the Linux ARM version"
	@env GOARCH=arm64 GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare

	@echo "[+] Packaging the Linux ARM version"
	@tar -czvf $(EXEC_DIR)goshare_Linux_aarch64.tar.gz -C $(EXEC_DIR) goshare > /dev/null
	# @sha256sum $(EXEC_DIR)goshare_Linux.tar.gz

	# @sha256sum $(EXEC_DIR)goshare_Linux.tar.gz > $(EXEC_DIR)goshare_Linux_sha256sum.txt

	# @echo
	# @echo "[+] Building the static Linux version"
	# @CGO_ENABLED=0 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare

	# @echo "[+] Packaging the static Linux version"
	# @tar -czvf $(EXEC_DIR)goshare_Linux_static.tar.gz -C $(EXEC_DIR) goshare > /dev/null
	# @sha256sum $(EXEC_DIR)goshare_Linux_static.tar.gz
	# @sha256sum $(EXEC_DIR)goshare_Linux_static.tar.gz > $(EXEC_DIR)goshare_Linux_static_sha256sum.txt

	# @echo
	# @echo "[+] Building the Debian package"
	# @cp $(EXECUTABLE_PATH) $(TARGET_EXECUTABLE_PATH)
	#
	# @echo "[+] Creating the Debian control file"
	# @echo "Package: $(PACKAGE_NAME)" > $(CONTROL_FILE)
	# @echo "Version: $(VERSION)" >> $(CONTROL_FILE)
	# @echo "Section: custom" >> $(CONTROL_FILE)
	# @echo "Priority: optional" >> $(CONTROL_FILE)
	# @echo "Architecture: amd64" >> $(CONTROL_FILE)
	# @echo "Essential: no" >> $(CONTROL_FILE)
	# @echo "Maintainer: $(MAINTAINER)" >> $(CONTROL_FILE)
	# @echo "Description: $(DESCRIPTION)" >> $(CONTROL_FILE)
	#
	# @echo "[+] Running dpkg-deb build"
	# @dpkg-deb --build $(DEB_PACKAGE_DIR)
	#
	# @echo "[+] Renaming the Debian package"
	# @mv $(DEB_PACKAGE_DIR).deb $(EXEC_DIR)/$(PACKAGE)

	@echo "[+] Removing the static Linux binary"
	@rm $(EXEC_DIR)goshare


	@echo
	@echo "[+] Building the Windows version"
	@env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare.exe

	@echo "[+] Packaging the Windows version"
	@zip -j $(EXEC_DIR)goshare_Windows.zip $(EXEC_DIR)goshare.exe > /dev/null
	@sha256sum  $(EXEC_DIR)goshare_Windows.zip
	# @sha256sum  $(EXEC_DIR)goshare_Windows.zip > $(EXEC_DIR)goshare_Windows_sha256sum.txt

	@echo "[+] Removing the Windows binary"
	@rm $(EXEC_DIR)goshare.exe

	@echo
	@echo "[+] Building the MacOS version"
	@env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare

	@echo "[+] Packaging the MacOS version"
	@tar -czvf $(EXEC_DIR)goshare_MacOS.tar.gz -C $(EXEC_DIR) goshare > /dev/null
	# @sha256sum $(EXEC_DIR)goshare_MacOS.tar.gz
	# @sha256sum $(EXEC_DIR)goshare_MacOS.tar.gz > $(EXEC_DIR)goshare_MacOS_sha256sum.txt


	@echo "[+] Removing the MacOS binary"
	@rm $(EXEC_DIR)goshare

	@echo
	@echo "[+] Building the MacOS ARM version"
	@env GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare

	@echo "[+] Packaging the MacOS ARM version"
	@tar -czvf $(EXEC_DIR)goshare_MacOS_ARM.tar.gz -C $(EXEC_DIR) goshare > /dev/null
	# @sha256sum $(EXEC_DIR)goshare_MacOS_ARM.tar.gz
	# @sha256sum $(EXEC_DIR)goshare_MacOS_ARM.tar.gz > $(EXEC_DIR)goshare_MacOS_ARM_sha256sum.txt

	@echo "[+] Removing the MacOS ARM binary"
	@rm $(EXEC_DIR)goshare

	@echo
	@echo "[+] Building the FreeBSD version"
	@env GOOS=freebsd GOARCH=amd64 go build -ldflags "-s -w" -o $(EXEC_DIR)goshare

	@echo "[+] Packaging the FreeBSD AMD64 version"
	@tar -czvf $(EXEC_DIR)goshare_FreeBSD.tar.gz -C $(EXEC_DIR) goshare > /dev/null
	# @sha256sum $(EXEC_DIR)goshare_FreeBSD.tar.gz
	# @sha256sum $(EXEC_DIR)goshare_FreeBSD.tar.gz > $(EXEC_DIR)goshare_FreeBSD_sha256sum.txt

	@echo "[+] Removing the FreeBSD binary"
	@rm $(EXEC_DIR)goshare

	@echo
	@echo "[+] Done"

update:
	@echo "[+] Updating Go dependencies"
	@go get -u
	@echo "[+] Done"

clean:
	@echo "[+] Cleaning files"
	@rm -rf $(EXEC_DIR)
	@echo "[+] Done"

format:
	@echo "[+] Formatting files"
	@gofmt -w *.go

vet:
	@echo "[+] Running Go vet"
	@go vet

test:
	@echo "[+] Running tests"
	@go test

tidyup:
	@echo "[+] Running go mod tidy"
	@go get -u ./...
	@go mod tidy

# This target is always executed, even if there are errors in previous targets.
.PHONY: always
always: debug
