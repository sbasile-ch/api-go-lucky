package main

import (
	pkgHndl "github.com/sbasile-ch/api-go-lucky/handlers"
	"net/http"
)

//__________________________________________________
func main() {
	http.HandleFunc("/getCompany", pkgHndl.GetCompanyProfile)
	http.HandleFunc("/index", pkgHndl.StartPage)
	/*
		http.HandleFunc("/getOfficer", getCompanyOfficer)
		http.HandleFunc("/getPsc", getCompanyPsc)
		http.HandleFunc("/getRegiters", getCompanyRegisters)
		http.HandleFunc("/setAlert", setCompanyAlert)
	*/

	// serve CSS and JS files locally stored in ./static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe("127.0.0.1:8080", nil)
}
