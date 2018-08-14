package apiclient

import (
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ApiParam struct {
	Host       string
	Url        string
	UserAgent  string
	httpClient *http.Client
	secTimeout time.Duration
}

var apiKey string = os.Getenv("MY_CH_API")

/*
   RFC7230: example of a GET request
      <http://www.example.org/pub/WWW/> would begin with:

	       GET /pub/WWW/ HTTP/1.1
		   Host: www.example.org


   GET /company/00000006 HTTP/1.1
   Host: api.companieshouse.gov.uk
   Authorization: Basic bXlfYXBpX2tleTo=
*/
//__________________________________________________
func GetApiResp(param *ApiParam) (respTxt string, err error) {
	//var c = &http.Client{Timeout: Timeout: param.secTimeout}
	var c = &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest("GET", param.Host+param.Url, nil)
	if err != nil {
		log.Printf("Error[%s] while forming new HTTP request [%s][%s]", err, param.Host, param.Url)
		return
	}
	//req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", param.UserAgent)
	req.Header.Set("Authorization", apiKey)

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Error[%s] while sending HTTP GET to [%s][%s][%v]", err, param.Host, param.Url, *c)
		return
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error[%s] while processing  HTTP response [%s][%s][%v]", err, param.Host, param.Url, *c)
		return
	}
	fmt.Printf("-------------[%s]\n", string(buff))
	return string(buff), err
}
