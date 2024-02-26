package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

const debug = false
const version = "4.1"

const usages string = `Usage of goshare:
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
`

var (
	rootDir, port, password string

	dontAllowUploads, dontAllowZipping, showStat bool
)

func flagParse() {
	flag.StringVar(&rootDir, "d", ".", "the directory for sharing")
	flag.StringVar(&port, "port", "8001", "port number")
	flag.StringVar(&password, "p", "", "password")
	flag.BoolVar(&showStat, "s", false, "don't show request information. aka silent")
	flag.BoolVar(&dontAllowUploads, "noup", false, "don't allow uploads")
	flag.BoolVar(&dontAllowZipping, "nozip", false, "don't allow zipping")
	v := flag.Bool("version", false, "show version number")
	flag.Usage = func() { fmt.Print(usages) }
	flag.Parse()

	if *v {
		mode := "release"
		if debug {
			mode = "debug"
		}
		fmt.Printf("goshare version v%s %s\n", version, mode)
		os.Exit(0)
	}

	if port == "" {
		log.Fatal("prot number can't be empty")
	}

	if rootDir == "" {
		log.Fatal("directory name can't be empty")
	}
}

func localIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "localhost"
}

func main() {
	flagParse()

	sv := newServer()
	go sv.cleanup()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
	})

	// don't chage the /browse/ ok it will break suff
	http.HandleFunc("/auth", sv.auth)
	http.HandleFunc("/browse/", sv.authBrowse)
	http.HandleFunc("/zip", sv.authZip)
	http.HandleFunc("/downzip/", sv.authDownZip)
	http.HandleFunc("/upload", sv.authUpload)
	http.HandleFunc("/mkdir", sv.authMkdir)

	http.HandleFunc("/static/", sv.authServeStaticFilese)

	if debug {
		fmt.Printf("Running in debug mode version: %s\n", version)
	}
	fmt.Printf("Serving at http://%s:%s\n", localIp(), port)
	if password != "" {
		fmt.Printf("Password is: %s\n\n", password)
	} else {
		fmt.Println()
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("\nwhile serving err: %v\n", err)
		fmt.Printf("Hint: mostlikely the issue is the port is alredy in use\n")
		fmt.Printf("use `--port 8002` to specefy another port\n")
		os.Exit(1)
	}
}
