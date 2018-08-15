package apiclient

import (
	"fmt"
	"github.com/sbasile-ch/api-go-lucky/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type ApiVars struct {
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
func GetApiResp(param *ApiVars) (respTxt string, err error) {

	cfg, err := config.Get()
	if err != nil {
		log.Printf("Error[%v] getting config", err)
	}
	var c = &http.Client{Timeout: time.Duration(cfg.HttpReqTimeout) * time.Second}
	req, err := http.NewRequest("GET", param.Host+param.Url, nil)
	if err != nil {
		log.Printf("Error[%v] while forming new HTTP request [%s][%s]", err, param.Host, param.Url)
		return
	}
	req.Header.Set("User-Agent", param.UserAgent)
	req.Header.Set("Authorization", apiKey)

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("Error[%v] while sending HTTP GET to [%s][%s][%v]", err, param.Host, param.Url, *c)
		return
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error[%v] while processing  HTTP response [%s][%s][%v]", err, param.Host, param.Url, *c)
		return
	}
	fmt.Printf("-------------[%s]\n", string(buff))
	return string(buff), err
}
