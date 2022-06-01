package client

import (
	"fmt"
	"github.com/NETWAYS/go-check"
	"github.com/NETWAYS/go-check/result"
	"github.com/microsoftgraph/msgraph-sdk-go/admin/serviceannouncement/issues"
	"strings"
)

type Services struct {
	Services []*Service
}

type Issues struct {
	Issues []*Issue
}

func (c *Client) LoadAllIssues() (issuesStruct *Issues, err error) {
	var offset int32 = 0

	serviceIssues, err := c.GraphServiceClient.Admin().ServiceAnnouncement().Issues().Get(&issues.IssuesRequestBuilderGetOptions{
		Q: &issues.IssuesRequestBuilderGetQueryParameters{
			// Skip the first n items
			Skip: &offset,
		},
	})
	if err != nil {
		err = fmt.Errorf("could not fetch service issues: %w", err)
		return
	}

	issuesStruct = &Issues{}

	issueVals := serviceIssues.GetValue()

	for _, issue := range issueVals {
		if *issue.GetIsResolved() {
			continue
		}

		cloned := issue

		issuesStruct.Issues = append(issuesStruct.Issues, &Issue{Issue: &cloned})
	}

	return
}

func (c *Client) LoadAllServices() (servicesStruct *Services, err error) {
	servicesInfo, err := c.GraphServiceClient.Admin().ServiceAnnouncement().HealthOverviews().Get(nil)
	if err != nil {
		err = fmt.Errorf("could not fetch services health: %w", err)
		return
	}

	if servicesInfo.IsNil() {
		err = fmt.Errorf("could not fetch services health overview")
		return
	}

	servicesStruct = &Services{}

	apiServices := servicesInfo.GetValue()
	for _, s := range apiServices {
		cloned := s

		servicesStruct.Services = append(servicesStruct.Services, &Service{Service: &cloned})
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

func (s *Services) GetOutput(override StateOverride, all bool, issues *Issues, displMsg bool) (output string) {
	if displMsg {
		for _, service := range s.Services {
			rc := service.GetStatus(override)
			if all {
				output += fmt.Sprintf("[%s] %s\n", check.StatusText(rc), service.GetOutput(issues, displMsg))
			} else if rc != 0 {
				output += fmt.Sprintf("[%s] %s\n", check.StatusText(rc), service.GetOutput(issues, displMsg))
			}
		}
	} else {
		for _, service := range s.Services {
			rc := service.GetStatus(override)
			if all {
				output += fmt.Sprintf("[%s] %s\n", check.StatusText(rc), service.GetOutput(nil, displMsg))
			} else if rc != 0 {
				output += fmt.Sprintf("[%s] %s\n", check.StatusText(rc), service.GetOutput(nil, displMsg))
			}
		}
	}

	return
}
