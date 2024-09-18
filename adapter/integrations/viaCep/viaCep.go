package viacep

import (
	"encoding/json"
	"fmt"
	"golang-uber-fx/core/domain"
	"net/http"
	"os"
	"time"
)

const (
	tryMaxRead         int = 5
	timeoutHTTPRequest     = 10 * time.Second
)

type IhttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
type ViaClient struct {
	client IhttpClient
	url    string
}

type IviaCep interface {
	GetAdress(cep string) (*domain.ViaCep, error)
}

func NewViaCep() IviaCep {

	return &ViaClient{
		client: &http.Client{
			Timeout: timeoutHTTPRequest,
		},
		url: os.Getenv("URL_VIACEP"),
	}
}

func (a *ViaClient) GetAdress(cep string) (*domain.ViaCep, error) {

	uri := fmt.Sprintf(a.url, cep)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	tryCount := 0

	for {
		tryCount++

		resp, err := a.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if (resp.StatusCode / 100) == 2 {
			response := domain.ViaCep{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {

				return nil, fmt.Errorf("[Viacep-api] Parse Response > Cep [%s] HTTP status [%d] - URI: [%s]", cep, resp.StatusCode, uri)
			}
			return &response, nil

		}
		if tryCount >= tryMaxRead {
			return nil, fmt.Errorf("[Viacep-api] cep [%s] HTTP status [%d] - URI: [%s]", cep, resp.StatusCode, uri)
		}

		fmt.Sprintf("[Viacep-api] Try [%d] ID [%s] HTTP status [%d]", tryCount, cep, resp.StatusCode)
		resp.Body.Close()
		time.Sleep(50 * time.Millisecond)
	}

}
