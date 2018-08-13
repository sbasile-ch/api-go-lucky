package handlers

import (
	"fmt"
	pkgUrl "github.com/sbasile-ch/api-go-lucky/apiurls"
	"html/template"
	"log"
	"net/http"
	"os"
)

type CompanyInfo struct {
	CompanyNum   string
	JsonTextArea string
	OfficerJson  string
}

var ApiKey string = os.Getenv("MY_CH_API")

//__________________________________________________
func GetCompanyProfile(w http.ResponseWriter, r *http.Request) {
	pageVars := CompanyInfo{
		CompanyNum:   "eeeeeeee",
		JsonTextArea: "I'm inside the getCompanyProfile",
		OfficerJson:  "To Evaluate 2",
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error[%s] on HTML Template Parse", err)
	}

	a := r.FormValue("CompanyNum")
	pageVars.CompanyNum = r.FormValue("CompanyNum")
	fmt.Printf("-----received--------[%s]\n", a)
	pageVars.JsonTextArea = r.FormValue("JsonTextArea")

	param := pkgUrl.TemplateUrl{CompanyNum: pageVars.CompanyNum}
	url, err := pkgUrl.GetApiUrl(pkgUrl.COMPANY, pkgUrl.ROA, &param)
	fmt.Printf("-------------[%s]\n", url)
	err = t.Execute(w, pageVars)
	if err != nil {
		log.Printf("Error[%s] on HTML Template Execute [%v]", err, pageVars)
	}
}

//__________________________________________________
func StartPage(w http.ResponseWriter, r *http.Request) {
	pageVars := CompanyInfo{
		CompanyNum:   "10989097",
		JsonTextArea: "I'm inside the startPage",
		OfficerJson:  "To Evaluate 1",
	}

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, pageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
