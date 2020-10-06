package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const uuid = "1a1aebea-4e05-45b9-8d11-c4115dbdd4a1"

type ApiRequests struct {
	url      string
	user     string
	password string
}

func NewApiRequests(url, user, password string) *ApiRequests {
	return &ApiRequests{
		url:      url,
		user:     user,
		password: password,
	}
}

func (apiRequests *ApiRequests) RequestServiceStates() ([]Service, error) {
	// token, err := apiRequests.requestToken()
	// if err != nil {
	// 	return nil, fmt.Errorf("token request error:%v", err)
	// }
	services, err := apiRequests.requestServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (apiRequests *ApiRequests) requestToken() (string, error) {
	rawTokenRequest := &TokenRequest{
		User:     apiRequests.user,
		Password: apiRequests.password,
		ID:       uuid,
	}
	tokenRequest, err := json.Marshal(rawTokenRequest)
	if err != nil {
		return "", err
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //TODO: don't skip tls verify
	}
	client := &http.Client{Transport: tr}
	reqPost, err := http.NewRequest("POST", apiRequests.url+"jwt/request-token", bytes.NewBuffer(tokenRequest))
	if err != nil {
		return "", err
	}
	reqPost.Header.Add("Content-Type", "application/json")

	reqGet, err := client.Do(reqPost)
	if err != nil {
		return "", err
	}
	defer reqGet.Body.Close()

	bodyBytesGet, err := ioutil.ReadAll(reqGet.Body)
	if err != nil {
		return "", err
	}

	tokenResponseOkay := &TokenResponseOkay{}
	if err = json.Unmarshal(bodyBytesGet, tokenResponseOkay); err != nil {
		return "", fmt.Errorf("token response error: %v", err)
	}

	if tokenResponseOkay.AccessToken == "" {
		tokenResponseError := &TokenResponseError{}
		err := json.Unmarshal(bodyBytesGet, tokenResponseError)
		if err != nil {
			return "", fmt.Errorf("unmarshal error token response fail: %v", err)
		}
		return "", fmt.Errorf("token response error: %v", tokenResponseError.Error)
	}
	return tokenResponseOkay.AccessToken, nil
}

func (apiRequests *ApiRequests) requestServices() ([]Service, error) {
	// getServicesRequest := &GetServicesRequest{
	// 	ID: uuid,
	// }
	// servicesRequest, err := json.Marshal(getServicesRequest)
	// if err != nil {
	// 	return nil, err
	// }
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //TODO: don't skip tls verify
	}
	client := &http.Client{Transport: tr}
	reqSendGet, err := http.NewRequest("GET", apiRequests.url+"services", nil)
	if err != nil {
		return nil, fmt.Errorf("request services error: %v", err)
	}
	// bearer := "Bearer " + token
	// reqSendGet.Header.Add("Authorization", bearer)
	// reqSendGet.Header.Add("Content-Type", "application/json")

	reqReadGet, err := client.Do(reqSendGet)
	if err != nil {
		return nil, err
	}
	defer reqReadGet.Body.Close()

	bodyBytesGet, err := ioutil.ReadAll(reqReadGet.Body)
	if err != nil {
		return nil, err
	}

	getAllServicesResponse := &GetAllServicesResponse{}
	if err = json.Unmarshal(bodyBytesGet, getAllServicesResponse); err != nil {
		return nil, fmt.Errorf("get all services unmarshal error: %v", err)
	}

	return getAllServicesResponse.AllServices, nil
}
