# GoShare

is a file server or sharer over the local network.

I have made this tool for sharing files from my computer (or home server) to my phone and other devices within the local network.

It consumes minimal resources, making it the perfect choice for running on a home server 24/7.

## About

- Only standard library used
- Upload files to the server.
- Password protection
- Zip directories
- Super-fast transfers
- Stylish UI (Tailwind CSS)
- Cross-platform compatible
- Simultaneous multi-file sharing
- Lightweight & efficient.

## Perks

- for streaming or downlading files form cli `$ mpv http://example.com/link/to/vid?cli=pass`
- Zip files for batch download

## Install

```bash
go install github.com/wizsk/goshare@latest
goshare -h
```

## Or get binnary from releases

```bash
wget 'https://github.com/wizsk/goshare/releases/latest/download/goshare_Linux_static.tar.gz'
# see realse page for windows
tar xvf 'goshare_Linux_static.tar.gz'
sudo mv goshare /usr/local/bin/ # or mv goshare ~/.local/bin/
```

## usages

```bash
Usage of goshare:
  -d string
        direcotry name (default ".")
  -p string
        password
  -port string
        port number (default "8001")
  -s    silence print informating about requests
  -u string
        upload directory (default "uploads")
  -v    prints current version

```

## Screenshots

## auth

![auth](/assets/ss/desktop-auth.png)

### Light

![light](/assets/ss/desktop-li.png)

### Dark

![dark](/assets/ss/desktop-da.png)

### Mobile

<table>
  <tr>
    <td> <img src="./assets/ss/m-li.png"  alt="1"></td>
    <td><img src="./assets/ss/m-da.png" alt="2"></td>
   </tr> 
  </tr>
</table>
