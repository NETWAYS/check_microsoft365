package client

import (
	"fmt"
	"github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type Service struct {
	Service *graph.ServiceHealth
}

//func (c *Client) ServiceHealth(serviceName string) (service *Service, err error) {
//	serviceInfo, err := c.GraphServiceClient.Admin().ServiceAnnouncement().HealthOverviewsById(serviceName).Get(nil)
//	if err != nil {
//		err = fmt.Errorf("could not fetch service health: %w", err)
//		return nil, err
//	}
//
//	if serviceInfo.IsNil() {
//		err := fmt.Errorf("could not fetch service health info")
//		check.ExitError(err)
//	}
//
//	service = &Service{Service: serviceInfo}
//
//	return
//}

func (s *Service) GetStatus(override StateOverride) int {
	stateName := s.Service.GetStatus().String()

	// Check if state is overridden
	if state, ok := override[stateName]; ok {
		return state
	} else {
		return LeveltoState(stateName)
	}
}

func (s *Service) GetOutput() (output string) {
	service := s.Service

	output = fmt.Sprintf("%s: %s %s", *service.GetId(), *service.GetService(), service.GetStatus().String())

	return
}
