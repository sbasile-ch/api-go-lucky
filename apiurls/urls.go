package apiurls

import (
	"bytes"
	"encoding/json"
	"fmt"
	pkgApi "github.com/sbasile-ch/api-go-lucky/apiclient"
	"log"
	"regexp"
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

type UrlVariables struct {
	CmdCategory   string
	CmdValue      string
	CompanyNum    string
	OfficerId     string
	ChargeId      string
	PscId         string
	StatementId   string
	SuperSecureId string
}

var urlTemplate = template.New("Urls")

type CmdUrlList map[string]string

type UrlCategory struct {
	Category string
	Host     string // if need of different from 'DefaultHost'
	List     CmdUrlList
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
			List: CmdUrlList{
				ALL:          "",
				COMPANY:      "/companies",
				DISQ_OFFICER: "/disqualified-officers",
				OFFICERS:     "/officers",
			}},
		OFFICER_APPOINTMENTS: UrlCategory{
			Category: "/officers/{{.OfficerId}}/appointments"},
		OFFICER_DISQUALIFIED: UrlCategory{
			Category: "/disqualified-officers",
			List: CmdUrlList{
				CORPORATE: "/corporate/{{.OfficerId}}",
				NATURAL:   "/natural/{{.OfficerId}}",
			}},
		COMPANY: UrlCategory{
			Category: "/company/{{.CompanyNum}}",
			List: CmdUrlList{
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
			List: CmdUrlList{
				LIST:          "",
				STATEMENTS:    "-statements",
				CORPORATE:     "/corporate-entity/{{.PscId}}",
				INDIVIDUAL:    "/individual/{{.PscId}}",
				LEGAL:         "/legal-person/{{.PscId}}",
				STATEMENTS_ID: "-statements/{{.StatementId}}",
				SUPER_SECURE:  "/super-secure/{{.SuperSecureId}}",
			}},
		REGISTERS: UrlCategory{
			Category: "/company/{{.CompanyNum}}/registers"},
		EXEMPTIONS: UrlCategory{
			Category: "/company/{{.CompanyNum}}/exemptions"},
	},
}

type JsonCmd struct {
	CmdName string   `json:"CmdName"`
	Args    []string `json:"Args"`
}

type JsonElem struct {
	CmdCategory string    `json:"CmdCategory"`
	CmdValues   []JsonCmd `json:"CmdValues"`
}

//__________________________________________________
// extract all Templates vars (eg. {{.xx}} ...{{.yy}} --> ["xx", ... "yy" ]
func extractArgs(s string) []string {

	arr := make([]string, 0, 5) // start with a capacity 5, (before allocating memory when running out).
	r, _ := regexp.Compile("{{\\.([^}]+)}}")
	mapArray := r.FindAllStringSubmatch(s, -1)
	for _, v := range mapArray {
		arr = append(arr, v[1])
	}
	return arr
}

//__________________________________________________
func ExportUrls() string {
	fmt.Printf("--ENTERED-----------[%v]\n", ApiUrls)
	jsonStruct := make([]JsonElem, len(ApiUrls.Urls))

	i := 0
	for k, Val := range ApiUrls.Urls {
		fmt.Printf("--SERIAL-----------[%s]\n", k)

		jsonStruct[i].CmdCategory = k

		Cmds := make([]JsonCmd, len(Val.List))
		j := 0
		for c, v := range Val.List {
			Cmds[j].CmdName = c
			//Cmds[j].Args = v.Args
			Cmds[j].Args = extractArgs(Val.Category + v)
			fmt.Printf("--SERIAL----------------------------[%s][%s][%v]\n", c, v, Cmds[j].Args)
			j++
		}
		jsonStruct[i].CmdValues = Cmds
		i++
	}

	output, err := json.MarshalIndent(&jsonStruct, "", "\t\t")
	if err != nil {
		log.Printf("Error[%v] marshalling to JSON", err)
		return ""
	}
	return string(output)
}

//__________________________________________________
func GetApiUrl(UrlVars *UrlVariables, ApiVars *pkgApi.ApiVars) error {

	c := ApiUrls.Urls[UrlVars.CmdCategory]
	// get proper Host
	host := c.Host
	if len(host) == 0 {
		host = ApiUrls.DefaultHost
	}

	// build the url
	url := c.Category
	if val, ok := c.List[UrlVars.CmdValue]; ok {
		url += val
	}
	t, err := urlTemplate.Parse(url)
	if err != nil {
		log.Printf("Error[%v] parsing URL template [%s]%s] with [%v]", err, host, url, UrlVars)
	}
	var buff bytes.Buffer
	err = t.Execute(&buff, UrlVars)
	if err != nil {
		log.Printf("Error[%v] executing URL template [%s][%s] with [%v]", err, host, url, UrlVars)
	}
	fmt.Printf("executing URL template [%s][%s] with [%v]", host, url, UrlVars)
	//fmt.Printf("-------------[%s]\n", buff.String())
	ApiVars.Host = host
	ApiVars.Url = buff.String()
	return err
}
