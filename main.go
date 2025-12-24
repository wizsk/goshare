package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/skip2/go-qrcode"
	"github.com/wizsk/goshare/tls"
)

const (
	debug        = false
	version      = "4.4"
	defaultPortS = 8080
	defaultPortE = 8099
	authPostPath = "/authp"
)

var usages string = `Usage of goshare:
Share specifed directy to the localnetwork.

OPTIONS:
  -d <directory_name>
        the directory for sharing (default ".")
  -p <password>
        password (default is no password)
  -s
        don't show status, be silent
  --noqr
        don't show qrcode
  --noup
        don't allow uploads or making directories
  --nozip
        don't allow zipping
  --port <port_number>
        port number (default range "` + fmt.Sprintf("%d-%d", defaultPortS, defaultPortE) + `")
  --nohttps
        don't use https aka tls
  --version
        show version number

EXAMPLES
       goshare -d "fo/bar/bazz" -p "777"
           share "fo/bar/bazz" directory. password would be "777"
`

var (
	rootDir, port, password string

	dontAllowUploads, dontAllowZipping, dontShowStat, dontShowQr, noHttps bool
)

func flagParse() {
	flag.StringVar(&rootDir, "d", ".", "the directory for sharing")
	flag.StringVar(&port, "port", "", "port number")
	flag.StringVar(&password, "p", "", "password")
	flag.BoolVar(&dontShowStat, "s", false, "don't show request information. aka silent")
	flag.BoolVar(&dontShowQr, "noqr", false, "don't show qr")
	flag.BoolVar(&dontAllowUploads, "noup", false, "don't allow uploads")
	flag.BoolVar(&dontAllowZipping, "nozip", false, "don't allow zipping")
	flag.BoolVar(&noHttps, "nohttps", false, "use https")
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
	fmt.Println("Starting server")

	sv := newServer()
	sv.showStat = false
	go sv.cleanup()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/browse/", http.StatusMovedPermanently)
	})

	// don't chage the /browse/ ok it will break suff
	http.HandleFunc("/auth", sv.auth)
	// work aoround
	http.HandleFunc(authPostPath, func(w http.ResponseWriter, r *http.Request) {
		r.Method = http.MethodPost
		sv.auth(w, r)
	})
	http.HandleFunc("/browse/", sv.authBrowse)
	http.HandleFunc("/zip", sv.authZip)
	http.HandleFunc("/downzip/", sv.authDownZip)
	http.HandleFunc("/upload", sv.authUpload)
	http.HandleFunc("/mkdir", sv.authMkdir)

	http.HandleFunc("/static/", sv.authServeStaticFilese)

	if debug {
		fmt.Printf("Running in debug mode version: %s\n", version)
	}

	if password != "" {
		fmt.Printf("Password is: %s\n", password)
	}

	lIP := localIp()
	var tp *tls.Provider
	var err error

	if !noHttps {
		if tp, err = tls.New(); err != nil {
			log.Fatal("Could not initiate tls certs", err)
		} else if err = tp.Ensure(); err != nil {
			log.Fatal("Could not ensure tls certs", err)
		}
	}

	var p string
	if port != "" {
		p = port
	} else {
		p = findFreePort(defaultPortS, defaultPortE)
		if p == "" {
			log.Fatal("Could not find a free port")
		}
	}

	go func() {
		var err error
		if noHttps {
			err = http.ListenAndServe(":"+p, nil)
		} else {
			err = http.ListenAndServeTLS(":"+p, tp.CertFile, tp.KeyFile, nil)
		}

		if err != nil {
			fmt.Printf("\nwhile serving err: %v\n", err)
			fmt.Printf("Hint: most likely the issue is, the port is alredy in use\n")
			fmt.Printf("use `--port 8002` to specefy another port\n")
			os.Exit(1)
		}
	}()

	s := "s"
	if noHttps {
		s = ""
	}

	m := fmt.Sprintf("http%s://%s:%s", s, lIP, p)
	if password != "" {
		m += "?password=" + url.QueryEscape(password)
	}

	if !noHttps {
		fmt.Println("INFO: as the certificate is self signed please ignore browser warning by clicking on advaced and continue")
		fmt.Println("INFO: do not forget to use https://")
	}
	fmt.Printf("\n%s\n\n", m)
	if !dontShowQr {
		var err error
		q, err := qrcode.New(m, qrcode.Medium)
		if err == nil {
			fmt.Println(q.ToSmallString(false))
		}
	}
	fmt.Println()
	sv.showStat = !dontShowStat

	select {}
}
