package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"telecommunication/entities"
	"time"
)

type SMSService interface {
	SendKaleryaSMS(request entities.KaleyraSMSRequest) string
}

type smsService struct {
}

func NewSMSService() SMSService {
	return &smsService{}
}

func (s *smsService) SendKaleryaSMS(request entities.KaleyraSMSRequest) string {
	//service := services.NewSMSService("Kalerya", os.Getenv("KALERYA_BASE_URL"), nil, true)
	baseURL := os.Getenv("KALERYA_BASE_URL")
	token := os.Getenv("KALERYA_TOKEN")
	var params []string
	if request.TemplateID != "" {
		params = append(params, fmt.Sprintf("&template_id=%s", request.TemplateID))
	}
	if request.Webhook != "" {
		params = append(params, fmt.Sprintf("&callback=%s", request.Webhook))
	}
	if request.Schedule != "" {
		params = append(params, fmt.Sprintf("&schedule=%s", request.Schedule))
	}
	serviceURL := fmt.Sprintf("%s/v4/?api_key=%s&method=sms&message=%s&to=%s&sender=%s",
		baseURL,
		token,
		url.QueryEscape(request.Message),
		request.To,
		request.From,
	)
	if len(params) > 0 {
		serviceURL = serviceURL + strings.Join(params, "")
	}
	if len(request.AdditionalParams) > 0 {
		serviceURL = serviceURL + strings.Join(request.AdditionalParams, "")
	}
	client := http.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := client.Get(serviceURL)
	if err != nil {

		return ""
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(bytes)
}
