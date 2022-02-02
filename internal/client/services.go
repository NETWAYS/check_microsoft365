package client

import (
	"fmt"
	"github.com/NETWAYS/go-check"
	"github.com/NETWAYS/go-check/result"
	"strings"
)

type Services struct {
	Services []*Service
}

func (c *Client) LoadAllServices() (services *Services, err error) {
	servicesInfo, err := c.GraphServiceClient.Admin().ServiceAnnouncement().HealthOverviews().Get(nil)
	if err != nil {
		err = fmt.Errorf("could not fetch services health: %w", err)
		return
	}

	if servicesInfo.IsNil() {
		err = fmt.Errorf("could not fetch services health overview")
		return
	}

	services = &Services{}

	apiServices := servicesInfo.GetValue()
	for _, s := range apiServices {
		cloned := s

		services.Services = append(services.Services, &Service{Service: &cloned})
	}

	return
}

func (c *Client) LoadServices(serviceNames []string) (services *Services, err error) {
	services, err = c.LoadAllServices()
	if err != nil {
		return
	}

	var filteredServices []*Service

	// Iterate over services. If ID from services in service is != nil. If matched -> break
	for _, service := range services.Services {
		for _, name := range serviceNames {
			if name == *service.Service.GetId() || strings.EqualFold(*service.Service.GetService(), name) {
				filteredServices = append(filteredServices, service)
				break
			}
		}
	}

	services.Services = filteredServices

	return
}

func (s *Services) GetStatus(override StateOverride) int {
	var states []int

	for _, service := range s.Services {
		states = append(states, service.GetStatus(override))
	}

	return result.WorstState(states...)
}

func (s *Services) GetOuput(override StateOverride) (output string) {
	for _, service := range s.Services {
		rc := service.GetStatus(override)
		if rc != 0 {
			output += fmt.Sprintf("[%s] %s\n", check.StatusText(rc), service.GetOutput())
		}
	}

	return
}
