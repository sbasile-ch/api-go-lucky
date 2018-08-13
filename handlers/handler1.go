package handlers

import (
	"fmt"
	pkgApi "github.com/sbasile-ch/api-go-lucky/apiclient"
	pkgUrl "github.com/sbasile-ch/api-go-lucky/apiurls"
	"html/template"
	"log"
	"net/http"
)

type CompanyInfo struct {
	CompanyNum   string
	JsonTextArea string
	OfficerJson  string
}

//__________________________________________________
func GetCompanyProfile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error[%s] on HTML Template Parse", err)
	}

	pageVars := CompanyInfo{CompanyNum: r.FormValue("CompanyNum")}

	fmt.Printf("-----received--------[%s]\n", pageVars.CompanyNum)

	param := pkgUrl.TemplateUrl{CompanyNum: pageVars.CompanyNum}
	api := pkgApi.ApiParam{}
	//err = pkgUrl.GetApiUrl(pkgUrl.COMPANY, pkgUrl.ROA, &param, &api)
	err = pkgUrl.GetApiUrl(pkgUrl.COMPANY, pkgUrl.FILING_HISTORY, &param, &api)
	fmt.Printf("-------------[%v]\n", &api)
	respTxt, err := pkgApi.GetApiResp(&api)
	if err != nil {
		log.Printf("Error[%s] calling the API", err)
	}
	fmt.Printf("--received-----------[%s]\n", respTxt)

	pageVars.JsonTextArea = respTxt
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
