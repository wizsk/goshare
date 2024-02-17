# GoShare

It is a file server or sharer over the local network.

It's used for quickly share files from one device to another in the local network.

## Why use this?

- Password protection
- Zip directories to download multiple files or directories.
- Upload files to directories.
- Create directory.
- Super-fast transfers
- ~Stylish~ UI (subjective)


## Intstall

### complile

```bash
go install github.com/wizsk/goshare@latest
goshare --help
```

### or see releases

```bash
# linux
cd /tmp
wget 'https://github.com/wizsk/goshare/releases/latest/download/goshare_Linux_static.tar.gz'
tar xf 'goshare_Linux_static.tar.gz'
sudo mv goshare /usr/local/bin/ # or mv goshare ~/.local/bin/
```

## Usages

```
goshare --help
Usage of goshare:
Share specifed directy to the localnetwork.

OPTIONS:
  -d <directory_name>
        the directory for sharing (default ".")
  -p <password>
        password (default is no password)
  -s
        don't show status, be silent
  --noup
        don't allow uploads or making directories
  --nozip
        don't allow zipping
  --port <port_number>
        port number (default "8001")
  --version
        show version number

EXAMPLES
       goshare -d "fo/bar/bazz" -p "777"
           share "fo/bar/bazz" directory. password would be "777"

```

## Screenshots

### auth

![auth](/demo/img/auth.jpg)

### browse

![browse](/demo/img/browse.jpg)

### upload

![upload](/demo/img/up.jpg)

### mobile

![mobile browse and menu](/demo/img/mobile_browse+menu.jpg)


## Thanks to

- [@mdJoOy](https://github.com/mdJoOy) for testing and contributions.