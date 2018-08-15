package main

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/sbasile-ch/api-go-lucky/config"
	pkgHndl "github.com/sbasile-ch/api-go-lucky/handlers"
	"log"
	"net/http"
	"os"
)

//__________________________________________________
func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Printf("Error[%v] getting config", err)
		os.Exit(1)
	}

	http.HandleFunc("/index", pkgHndl.StartPage)
	http.HandleFunc("/GetCompanyCmd", pkgHndl.GetCompanyCmd)
	//	http.HandleFunc("/setAlert", setCompanyAlert)

	// serve CSS and JS files locally stored in ./static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	serverAddress := fmt.Sprintf("%s:%d", cfg.ServerIpAddr, cfg.ServerPort)

	//prepare to launch the browser automatically
	url := fmt.Sprintf("http://%s/index", serverAddress)
	browser.OpenURL(url)

	http.ListenAndServe(serverAddress, nil)
}
