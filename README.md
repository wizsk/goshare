# GoShare

is a file server or sharer over the local network.

I have developed a convenient solution that enables seamless file-sharing between my computer, phone, and other devices within the local network.

## About

- Password protection
- CLI client (goshare-cli)
- Super-fast transfers
- Stylish UI (Tailwind CSS)
- Cross-platform compatible
- Simultaneous multi-file sharing
- Lightweight & efficient.

## Perks

- [cli user interface](https://github.com/wizsk/goshare/blob/main/goshare-cli) made with `bash`
- if passord is set then video or other files can be directly streamed any apps with `$ mpv http://example.com/link/to/vid?cli=pass`

## Install

```bash
go install github.com/wizsk/goshare@latest
goshare -h
```

## Or get binnary from releases

```bash
wget 'https://github.com/wizsk/goshare/releases/latest/download/goshare_linux64.tar.gz'
# see realse page for windows
tar xvf 'goshare_linux64.tar.gz'
sudo mv goshare /usr/local/bin/ # or mv goshare ~/.local/bin/
```

## usages

```bash
Usage of goshare:
  -d string
    	direcotry name (default "." current direcotry)
  -p string
    	password (default is empty)
  -port string
    	port number (default "8001")
  -v	prints current version

```

## cli client

I love to work in cli so made this.

dependencies: `curl`, `wget` and `fzf`

```bash
# sudo dnf install curl wget fzf
wget 'https://github.com/wizsk/goshare/blob/main/goshare-cli'
chmod +x goshare-cli
sudo mv goshare-cli /usr/local/bin/ # or mv goshare-cli ~/.local/bin/
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
