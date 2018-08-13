package apiurls

import (
	"bytes"
	"fmt"
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

type TemplateUrl struct {
	OfficerId     string
	CompanyNum    string
	ChargeId      string
	PscId         string
	StatementId   string
	SuperSecureId string
}

var urlTemplate = template.New("Urls")

type SubCategoryList map[string]string

type Category struct {
	name    string
	subList SubCategoryList
}

type CategoryList map[string]Category

var Urls = CategoryList{

	SEARCH: Category{
		name: "https://api.companieshouse.gov.uk/search",
		subList: SubCategoryList{
			ALL:          "",
			COMPANY:      "/companies",
			DISQ_OFFICER: "/disqualified-officers",
			OFFICERS:     "/officers",
		}},
	OFFICER_APPOINTMENTS: Category{
		name: "https://api.companieshouse.gov.uk/officers/{{.OfficerId}}/appointments", subList: SubCategoryList{}},
	OFFICER_DISQUALIFIED: Category{
		name: "https://api.companieshouse.gov.uk/disqualified-officers",
		subList: SubCategoryList{
			CORPORATE: "/corporate/{{.OfficerId}}",
			NATURAL:   "/natural/{{.OfficerId}}",
		}},
	COMPANY: Category{
		name: "https://api.companieshouse.gov.uk/company/{{.CompanyNum}}",
		subList: SubCategoryList{
			PROFILE:        "",
			ROA:            "/registered-office-address",
			OFFICERS_LIST:  "/officers",
			FILING_HISTORY: "/filing-history",
			INSOLVENCY:     "/insolvency",
			CHARGES:        "/charges/",
			CHARGES_ID:     "/charges/{{.ChargeId}}",
			UK_ESTABLISHED: "/uk-establishments",
		}},
	PSC: Category{
		name: "https://api.companieshouse.gov.uk/company/{{.CompanyNum}}/persons-with-significant-control",
		subList: SubCategoryList{
			LIST:          "",
			STATEMENTS:    "-statements",
			CORPORATE:     "/corporate-entity/{{.PscId}}",
			INDIVIDUAL:    "/individual/{{.PscId}}",
			LEGAL:         "/legal-person/{{.PscId}}",
			STATEMENTS_ID: "-statements/{{.StatementId}}",
			SUPER_SECURE:  "/super-secure/{{.SuperSecureId}}",
		}},
	REGISTERS: Category{
		name: "https://api.companieshouse.gov.uk/company/{{.CompanyNum}}/registers", subList: SubCategoryList{}},
	EXEMPTIONS: Category{
		name: "https://api.companieshouse.gov.uk/company/{{.CompanyNum}}/exemptions", subList: SubCategoryList{}},
}

/*
type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}
*/

//__________________________________________________
func GetApiUrl(cat string, subcat string, param *TemplateUrl) (string, error) {
	c := Urls[cat]
	url := c.name
	if val, ok := c.subList[subcat]; ok {
		url += val
	}
	t, err := urlTemplate.Parse(url)
	if err != nil {
		log.Printf("Error[%s] parsing URL template [%s] with [%s][%s][%v]", err, url, cat, subcat, param)
	}
	var buff bytes.Buffer
	err = t.Execute(&buff, param)
	if err != nil {
		log.Printf("Error[%s] executing URL template [%s] with [%s][%s][%v]", err, url, cat, subcat, param)
	}
	fmt.Printf("executing URL template [%s] with [%s][%s][%v]", url, cat, subcat, param)
	fmt.Printf("-------------[%s]\n", buff.String())
	return buff.String(), err
}
