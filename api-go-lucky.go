package main

import (
	"fmt"
	pkgHndl "github.com/sbasile-ch/api-go-lucky/handlers"
	"net/http"
	"os/exec"
	"runtime"
)

const (
	ServerIpAddr = "127.0.0.1"
	ServerPort   = 8080
)

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

//__________________________________________________
func main() {
	http.HandleFunc("/index", pkgHndl.StartPage)
	http.HandleFunc("/GetCompanyCmd", pkgHndl.GetCompanyCmd)
	//	http.HandleFunc("/setAlert", setCompanyAlert)

	// serve CSS and JS files locally stored in ./static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	serverAddr := fmt.Sprintf("%s:%d", ServerIpAddr, ServerPort)
	http.ListenAndServe(serverAddr, nil)
	open(serverAddr)
}
