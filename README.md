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

<!-- ## Install -->
<!---->
<!-- ```bash -->
<!-- go install github.com/wizsk/goshare@latest -->
<!-- goshare -h -->
<!-- ``` -->

## Intstall

see releases
<!-- ```bash
wget 'https://github.com/wizsk/goshare/releases/latest/download/goshare_Linux_static.tar.gz'
# see realse page for windows
tar xvf 'goshare_Linux_static.tar.gz'
sudo mv goshare /usr/local/bin/ # or mv goshare ~/.local/bin/
``` -->

## usages

```
goshare --help
Usage of goshare:
  -d
      the directory for sharing (default ".")
  -p
      password (default is no password)
  -s
      don't show status, be silent
  --noup
      don't allow uploads or making directories
  --nozip
      don't allow zipping
  --port string
      port number (default "8001")
  --version
      show version number

EXAMPLES
       goshare -d "fo/bar/bazz" -p "777"
           share "fo/bar/bazz" directory. password would be "777"
```

<!-- ## Screenshots -->
<!---->
<!-- ## auth -->
<!---->
<!-- ![auth](/assets/ss/desktop-auth.png) -->
<!---->
<!-- ### Light -->
<!---->
<!-- ![light](/assets/ss/desktop-li.png) -->
<!---->
<!-- ### Dark -->
<!---->
<!-- ![dark](/assets/ss/desktop-da.png) -->
<!---->
<!-- ### Mobile -->
<!---->
<!-- <table> -->
<!--   <tr> -->
<!--     <td> <img src="./assets/ss/m-li.png"  alt="1"></td> -->
<!--     <td><img src="./assets/ss/m-da.png" alt="2"></td> -->
<!--    </tr> -->
<!--   </tr> -->
<!-- </table> -->

## Thanks to

- [@mdJoOy](https://github.com/mdJoOy) for testing and bug contributions.

## TODOS for the rewrite

- [ ] do more testings
- [ ] static files dir upload
- [ ] favicons

## DONE

- [x] work on the frontend styles
    - [x] `/browse/` page
    - [x] `/upload` page
    - [x] `/login` page
- [x] don't allow uploads `--noup`
- [x] show no files found page
- [x] clean tmp directory
- [x] browse files.
- [x] upload files where you currently is
- [x] zip dir && download
- [x] name zip files properly
- [x] temp dir
- [x] auth
- [x] bug with zip cancaletion
- [x] network request stat
