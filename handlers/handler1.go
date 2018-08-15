package handlers

import (
	"fmt"
	pkgApi "github.com/sbasile-ch/api-go-lucky/apiclient"
	pkgUrl "github.com/sbasile-ch/api-go-lucky/apiurls"
	"html/template"
	"log"
	"net/http"
)

type MainTemplateVars struct {
	JsonResponse string
	JsonCmdsList string
	ApiUrl       string
	CmdCategory  string
	CmdValue     string
	UrlVars      pkgUrl.UrlVariables
}

//__________________________________________________
func GetCompanyCmd(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error[%v] on HTML Template Parse", err)
	} else {

		CmdCategory := r.FormValue("CmdCategory")
		CmdValue := r.FormValue("CmdValue")
		UrlVars := pkgUrl.UrlVariables{
			//CmdCategory:   r.FormValue("CmdCategory"),
			//CmdValue:      r.FormValue("CmdValue"),
			CmdCategory:   CmdCategory,
			CmdValue:      CmdValue,
			CompanyNum:    r.FormValue("CompanyNum"),
			OfficerId:     r.FormValue("OfficerId"),
			ChargeId:      r.FormValue("ChargeId"),
			PscId:         r.FormValue("PscId"),
			StatementId:   r.FormValue("StatementId"),
			SuperSecureId: r.FormValue("SuperSecureId"),
		}
		templVars := MainTemplateVars{JsonCmdsList: pkgUrl.JSONEXPORT, UrlVars: UrlVars}

		fmt.Printf("-----QUI-------[%s][%s][%+v]\n", UrlVars.CmdCategory, UrlVars.CmdValue, UrlVars)
		ApiVars := pkgApi.ApiVars{}
		err = pkgUrl.GetApiUrl(&UrlVars, &ApiVars)
		templVars.ApiUrl = ApiVars.Host + ApiVars.Url
		respTxt, err := pkgApi.GetApiResp(&ApiVars)
		if err != nil {
			log.Printf("Error[%v] calling the API", err)
		} else {
			templVars.JsonResponse = respTxt
			err = t.Execute(w, templVars)
			if err != nil {
				log.Printf("Error[%v] on HTML Template Execute [%v]", err, templVars)
			}
		}
	}
}

//__________________________________________________
func StartPage(w http.ResponseWriter, r *http.Request) {
	templVars := MainTemplateVars{UrlVars: pkgUrl.UrlVariables{CompanyNum: "00006400"}, JsonCmdsList: pkgUrl.JSONEXPORT}

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Printf("Error[%v] parsing Template", err)
	} else {
		err = t.Execute(w, templVars)
		if err != nil {
			log.Printf("Error[%v] executing Template", err)
		}
	}
}
