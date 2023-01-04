package v3

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

var version = "v3"

func NewAMapServiceV3(key string) amaphelper {
	svc := &service{
		key: key,
	}
	svc.client = resty.New()
	svc.client.SetAllowGetMethodPayload(true)
	svc.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	svc.client.SetTimeout(10 * time.Second)
	svc.client.SetBaseURL("https://restapi.amap.com/" + version)
	return svc
}

type service struct {
	client *resty.Client
	key    string
}

func (svc *service) Regeo(longitude float64, latitude float64) (*RegeoResponse, error) {
	// https://restapi.amap.com/v3/geocode/regeo
	params := map[string]string{
		"key":        svc.key,
		"location":   fmt.Sprintf("%0.6f,%0.6f", longitude, latitude),
		"roadlevel":  "1",
		"extensions": "all",
	}
	result := &RegeoResponse{}
	response, err := svc.client.SetQueryParams(params).R().SetResult(&result).Get("geocode/regeo")
	if err != nil {
		return nil, err
	}
	if response.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf(response.Status())
	}
	return result, nil
}
func (svc *service) Geo() (float64, float64) {
	return 0, 0
}
