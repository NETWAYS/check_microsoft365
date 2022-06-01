package client

import (
	"fmt"
	"github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

type Service struct {
	Service *graph.ServiceHealth
}

type Issue struct {
	Issue *graph.ServiceHealthIssue
}

func (s *Service) GetStatus(override StateOverride) int {
	stateName := s.Service.GetStatus().String()

	// Check if state is overridden
	if state, ok := override[stateName]; ok {
		return state
	} else {
		return LeveltoState(stateName)
	}
}

func (s *Service) GetOutput(issues *Issues, displMsg bool) (output string) {
	service := s.Service

	output = fmt.Sprintf("%s: %s %s", *service.GetId(), *service.GetService(), service.GetStatus().String())

	if displMsg {
		for idx, issue := range issues.Issues {
			if *service.GetService() == *issue.Issue.GetService() {
				output += fmt.Sprintf("\n  \\_[%d] %s: %s", idx, *issue.Issue.GetStartDateTime(), *issue.Issue.GetImpactDescription())
			}
		}
	}

	return
}
