package insightly

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

// Instance stores Instance from Service
//
type Instance struct {
	InstanceName             string  `json:"INSTANCE_NAME"`
	InstanceSubdomain        *string `json:"INSTANCE_SUBDOMAIN"`
	PlanName                 string  `json:"PLAN_NAME"`
	NewUserExperienceEnabled bool    `json:"NEW_USER_EXPERIENCE_ENABLED"`
}

// GetInstance returns all instance
//

func (service *Service) GetInstance() (*Instance, *http.Response, *errortools.Error) {
	instance := Instance{}
	requestConfig := go_http.RequestConfig{
		URL:           service.url("Instance"),
		ResponseModel: &instance,
	}

	_, response, e := service.get(&requestConfig)
	if e != nil {
		return nil, response, e
	}

	return &instance, response, nil
}
