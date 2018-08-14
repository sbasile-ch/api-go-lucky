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

type VarArgs []string

type CmdUrl struct {
	Url  string
	Args VarArgs
}
type CmdUrlList map[string]CmdUrl

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
				ALL:          CmdUrl{},
				COMPANY:      CmdUrl{Url: "/companies"},
				DISQ_OFFICER: CmdUrl{Url: "/disqualified-officers"},
				OFFICERS:     CmdUrl{Url: "/officers"},
			}},
		OFFICER_APPOINTMENTS: UrlCategory{
			Category: "/officers/{{.OfficerId}}/appointments"},
		OFFICER_DISQUALIFIED: UrlCategory{
			Category: "/disqualified-officers",
			List: CmdUrlList{
				CORPORATE: CmdUrl{Url: "/corporate/{{.OfficerId}}", Args: []string{"OfficerId"}},
				NATURAL:   CmdUrl{Url: "/natural/{{.OfficerId}}", Args: []string{"OfficerId"}},
			}},
		COMPANY: UrlCategory{
			Category: "/company/{{.CompanyNum}}",
			List: CmdUrlList{
				PROFILE:        CmdUrl{},
				ROA:            CmdUrl{Url: "/registered-office-address"},
				OFFICERS_LIST:  CmdUrl{Url: "/officers"},
				FILING_HISTORY: CmdUrl{Url: "/filing-history"},
				INSOLVENCY:     CmdUrl{Url: "/insolvency"},
				CHARGES:        CmdUrl{Url: "/charges/"},
				CHARGES_ID:     CmdUrl{Url: "/charges/{{.ChargeId}}", Args: []string{"ChargeId"}},
				UK_ESTABLISHED: CmdUrl{Url: "/uk-establishments"},
			}},
		PSC: UrlCategory{
			Category: "/company/{{.CompanyNum}}/persons-with-significant-control",
			List: CmdUrlList{
				LIST:          CmdUrl{},
				STATEMENTS:    CmdUrl{Url: "-statements"},
				CORPORATE:     CmdUrl{Url: "/corporate-entity/{{.PscId}}", Args: []string{"PscId"}},
				INDIVIDUAL:    CmdUrl{Url: "/individual/{{.PscId}}", Args: []string{"PscId"}},
				LEGAL:         CmdUrl{Url: "/legal-person/{{.PscId}}", Args: []string{"PscId"}},
				STATEMENTS_ID: CmdUrl{Url: "-statements/{{.StatementId}}", Args: []string{"StatementId"}},
				SUPER_SECURE:  CmdUrl{Url: "/super-secure/{{.SuperSecureId}}", Args: []string{"SuperSecureId"}},
			}},
		REGISTERS: UrlCategory{
			Category: "/company/{{.CompanyNum}}/registers"},
		EXEMPTIONS: UrlCategory{
			Category: "/company/{{.CompanyNum}}/exemptions"},
	},
}

/*
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}
*/
type JsonCmd struct {
	CmdName string   `json:"CmdName"`
	Args    []string `json:"Args"`
}

type JsonElem struct {
	CmdCategory string    `json:"CmdCategory"`
	CmdValues   []JsonCmd `json:"CmdValues"`
}

//__________________________________________________
func ExportUrls() (s string) {
	fmt.Printf("--ENTERED-----------[%v]\n", ApiUrls)
	jsonStruct := make([]JsonElem, len(ApiUrls.Urls))

	i := 0
	for k, Val := range ApiUrls.Urls {
		fmt.Printf("--SERIAL-----------[%s]\n", k)

		jsonStruct[i].CmdCategory = k

		Cmds := make([]JsonCmd, len(Val.List))
		j := 0
		for c, v := range Val.List {
			fmt.Printf("--SERIAL----------------------------[%s][%v]\n", c, v)
			Cmds[j].CmdName = c
			Cmds[j].Args = v.Args
			j++
		}
		jsonStruct[i].CmdValues = Cmds
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
		url += val.Url
	}
	t, err := urlTemplate.Parse(url)
	if err != nil {
		log.Printf("Error[%s] parsing URL template [%s]%s] with [%v]", err, host, url, UrlVars)
	}
	var buff bytes.Buffer
	err = t.Execute(&buff, UrlVars)
	if err != nil {
		log.Printf("Error[%s] executing URL template [%s][%s] with [%v]", err, host, url, UrlVars)
	}
	fmt.Printf("executing URL template [%s][%s] with [%v]", host, url, UrlVars)
	//fmt.Printf("-------------[%s]\n", buff.String())
	ApiVars.Host = host
	ApiVars.Url = buff.String()
	return err
}
