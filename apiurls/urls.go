package apiurls

import (
	"bytes"
	"encoding/json"
	"fmt"
	pkgApi "github.com/sbasile-ch/api-go-lucky/apiclient"
	"log"
	"text/template"
)

const SEARCH string = "SEARCH"
const ALL string = "ALL"
const COMPANY string = "COMPANY"
const DISQ_OFFICER string = "DISQ_OFFICER"
const OFFICERS string = "OFFICERS"
const OFFICER_APPOINTMENTS string = "OFFICER_APPOINTMENTS"
const OFFICER_DISQUALIFIED string = "OFFICER_DISQUALIFIED"
const CORPORATE string = "CORPORATE"
const NATURAL string = "NATURAL"
const PROFILE string = "PROFILE"
const ROA string = "ROA"
const OFFICERS_LIST string = "OFFICERS_LIST"
const FILING_HISTORY string = "FILING_HISTORY"
const INSOLVENCY string = "INSOLVENCY"
const CHARGES string = "CHARGES"
const CHARGES_ID string = "CHARGES_ID"
const UK_ESTABLISHED string = "UK_ESTABLISHED"
const PSC string = "PSC"
const LIST string = "LIST"
const INDIVIDUAL string = "INDIVIDUAL"
const LEGAL string = "LEGAL"
const STATEMENTS string = "STATEMENTS"
const STATEMENTS_ID string = "STATEMENTS_ID"
const SUPER_SECURE string = "SUPER_SECURE"
const REGISTERS string = "REGISTERS"
const EXEMPTIONS string = "EXEMPTIONS"

var JSONEXPORT string = ExportUrls()

type TemplateUrl struct {
	OfficerId     string
	CompanyNum    string
	ChargeId      string
	PscId         string
	StatementId   string
	SuperSecureId string
}

var urlTemplate = template.New("Urls")

type UrlList map[string]string

type UrlCategory struct {
	Category string
	Host     string // if need of different from 'DefaultHost'
	List     UrlList
}

type UrlSet map[string]UrlCategory

type ApiFactory struct {
	DefaultHost string
	Urls        UrlSet
}

var ApiUrls = ApiFactory{
	DefaultHost: "https://api.companieshouse.gov.uk",
	Urls: UrlSet{

		SEARCH: UrlCategory{
			Category: "/search",
			List: UrlList{
				ALL:          "",
				COMPANY:      "/companies",
				DISQ_OFFICER: "/disqualified-officers",
				OFFICERS:     "/officers",
			}},
		OFFICER_APPOINTMENTS: UrlCategory{
			Category: "/officers/{{.OfficerId}}/appointments", List: UrlList{}},
		OFFICER_DISQUALIFIED: UrlCategory{
			Category: "/disqualified-officers",
			List: UrlList{
				CORPORATE: "/corporate/{{.OfficerId}}",
				NATURAL:   "/natural/{{.OfficerId}}",
			}},
		COMPANY: UrlCategory{
			Category: "/company/{{.CompanyNum}}",
			List: UrlList{
				PROFILE:        "",
				ROA:            "/registered-office-address",
				OFFICERS_LIST:  "/officers",
				FILING_HISTORY: "/filing-history",
				INSOLVENCY:     "/insolvency",
				CHARGES:        "/charges/",
				CHARGES_ID:     "/charges/{{.ChargeId}}",
				UK_ESTABLISHED: "/uk-establishments",
			}},
		PSC: UrlCategory{
			Category: "/company/{{.CompanyNum}}/persons-with-significant-control",
			List: UrlList{
				LIST:          "",
				STATEMENTS:    "-statements",
				CORPORATE:     "/corporate-entity/{{.PscId}}",
				INDIVIDUAL:    "/individual/{{.PscId}}",
				LEGAL:         "/legal-person/{{.PscId}}",
				STATEMENTS_ID: "-statements/{{.StatementId}}",
				SUPER_SECURE:  "/super-secure/{{.SuperSecureId}}",
			}},
		REGISTERS: UrlCategory{
			Category: "/company/{{.CompanyNum}}/registers", List: UrlList{}},
		EXEMPTIONS: UrlCategory{
			Category: "/company/{{.CompanyNum}}/exemptions", List: UrlList{}},
	},
}

/*
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}
*/

type JsonElem struct {
	CmdCategory string   `json:"CmdCategory"`
	CmdValues   []string `json:"CmdValues"`
}

//__________________________________________________
func ExportUrls() (s string) {
	fmt.Printf("--ENTERED-----------[%v]\n", ApiUrls)
	jsonStruct := make([]JsonElem, len(ApiUrls.Urls))

	i := 0
	for k, Val := range ApiUrls.Urls {
		fmt.Printf("--SERIAL-----------[%s]\n", k)

		jsonStruct[i].CmdCategory = k

		values := make([]string, len(Val.List))
		j := 0
		for v := range Val.List {
			fmt.Printf("--SERIAL----------------------------[%s]\n", v)
			values[j] = v
			j++
		}
		jsonStruct[i].CmdValues = values
		i++
	}

	output, err := json.MarshalIndent(&jsonStruct, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	s = string(output)
	return
}

//__________________________________________________
func GetApiUrl(cat string, subcat string, param *TemplateUrl, apiParam *pkgApi.ApiParam) error {

	c := ApiUrls.Urls[cat]
	// get proper Host
	host := c.Host
	if len(host) == 0 {
		host = ApiUrls.DefaultHost
	}

	// build the url
	url := c.Category
	if val, ok := c.List[subcat]; ok {
		url += val
	}
	t, err := urlTemplate.Parse(url)
	if err != nil {
		log.Printf("Error[%s] parsing URL template [%s]%s] with [%s][%s][%v]", err, host, url, cat, subcat, param)
	}
	var buff bytes.Buffer
	err = t.Execute(&buff, param)
	if err != nil {
		log.Printf("Error[%s] executing URL template [%s][%s] with [%s][%s][%v]", err, host, url, cat, subcat, param)
	}
	fmt.Printf("executing URL template [%s][%s] with [%s][%s][%v]", host, url, cat, subcat, param)
	fmt.Printf("-------------[%s]\n", buff.String())
	apiParam.Host = host
	apiParam.Url = buff.String()
	return err
}
