package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strings"
	"time"
)

type EnvHttp struct {
	DebugClient		bool		`envconfig:"DEBUG_CLIENT" default:"true"`
	Timeout			string		`envconfig:"TIMEOUT" default:"60s"`
	RetryBad		int			`envconfig:"RETRY_BAD" default:"1"`
}

var (

	envHttp		EnvHttp

)


func init() {

	err := envconfig.Process("HTTP", &envHttp)
	if err != nil {
		fmt.Println("Failed to get HTTP env:", err)
	}

}

// HTTPGet func
func HTTPGet(url string, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(envHttp.DebugClient)
	timeout, _ := time.ParseDuration(envHttp.Timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Get(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Timeout(timeout).
		Retry(envHttp.RetryBad, time.Second, http.StatusInternalServerError).
		End()


	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPost func
func HTTPPost(url string, jsondata interface{}) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(envHttp.DebugClient)
	timeout, _ := time.ParseDuration(envHttp.Timeout)
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(envHttp.RetryBad, time.Second, http.StatusInternalServerError).
		End()

	if errs != nil {
		return nil, errs[0]
	}
	return []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(envHttp.DebugClient)
	timeout, _ := time.ParseDuration(envHttp.Timeout)
	//_ := errors.New("Connection Problem")
	if strings.HasPrefix(url, "https") {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(envHttp.RetryBad, time.Second, http.StatusInternalServerError).
		End()

	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}
