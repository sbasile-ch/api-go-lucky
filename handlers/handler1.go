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
	CompanyNum      string
	JsonTextArea    string
	ApiUrl          string
	CmdCategory     string
	CmdValue        string
	OfficerId       string
	ChargeId        string
	PscId           string
	StatementId     string
	SuperSecureId   string
	CommandJsonList string
}

//__________________________________________________
func GetCompanyOfficers(w http.ResponseWriter, r *http.Request) {
	GetCompanyCmd(pkgUrl.COMPANY, pkgUrl.OFFICERS_LIST, w, r)
}

//__________________________________________________
func GetCompanyPscs(w http.ResponseWriter, r *http.Request) {
	GetCompanyCmd(pkgUrl.PSC, pkgUrl.LIST, w, r)
}

//__________________________________________________
func GetCompanyRegisters(w http.ResponseWriter, r *http.Request) {
	GetCompanyCmd(pkgUrl.REGISTERS, "", w, r)
}

//__________________________________________________
func GetCompanyCmd(CmdCategory string, CmdValue string, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error[%s] on HTML Template Parse", err)
	} else {

		if len(CmdCategory) == 0 {
			CmdCategory = r.FormValue("CmdCategory")
			CmdValue = r.FormValue("CmdValue")
		}
		fmt.Printf("-------------[%s][%s]\n", CmdCategory, CmdValue)
		//templVars := MainTemplateVars{CompanyNum: r.FormValue("CompanyNum")}
		templVars := MainTemplateVars{CompanyNum: "00006400"}

		param := pkgUrl.TemplateUrl{CompanyNum: templVars.CompanyNum}
		api := pkgApi.ApiParam{}
		//err = pkgUrl.GetApiUrl(pkgUrl.COMPANY, pkgUrl.ROA, &param, &api)
		err = pkgUrl.GetApiUrl(CmdCategory, CmdValue, &param, &api)
		templVars.ApiUrl = api.Host + api.Url
		respTxt, err := pkgApi.GetApiResp(&api)
		if err != nil {
			log.Printf("Error[%s] calling the API", err)
		} else {
			templVars.JsonTextArea = respTxt
			err = t.Execute(w, templVars)
			if err != nil {
				log.Printf("Error[%s] on HTML Template Execute [%v]", err, templVars)
			}
		}
	}
}

//__________________________________________________
func GetCompanyProfile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error[%s] on HTML Template Parse", err)
	} else {

		templVars := MainTemplateVars{CompanyNum: r.FormValue("CompanyNum"), CommandJsonList: pkgUrl.JSONEXPORT}

		param := pkgUrl.TemplateUrl{CompanyNum: templVars.CompanyNum}
		api := pkgApi.ApiParam{}
		err = pkgUrl.GetApiUrl(pkgUrl.COMPANY, pkgUrl.FILING_HISTORY, &param, &api)
		templVars.ApiUrl = api.Host + api.Url
		respTxt, err := pkgApi.GetApiResp(&api)
		if err != nil {
			log.Printf("Error[%s] calling the API", err)
		} else {
			templVars.JsonTextArea = respTxt
			err = t.Execute(w, templVars)
			if err != nil {
				log.Printf("Error[%s] on HTML Template Execute [%v]", err, templVars)
			}
		}
	}
}

//__________________________________________________
func StartPage(w http.ResponseWriter, r *http.Request) {
	s := pkgUrl.JSONEXPORT
	log.Print("======[%s]====== ", s)
	templVars := MainTemplateVars{CompanyNum: "00006400", CommandJsonList: pkgUrl.JSONEXPORT}

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Print("template parsing error: ", err)
	} else {
		err = t.Execute(w, templVars)
		if err != nil {
			log.Print("template executing error: ", err)
		}
	}
}
