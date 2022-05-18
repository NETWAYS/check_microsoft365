package cmd

import (
	"fmt"
	"github.com/NETWAYS/check_microsoft365/internal/client"
	"github.com/NETWAYS/go-check"
	"github.com/NETWAYS/go-check/perfdata"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	output    string
	summary   string
	totalCrit int
	totalWarn int
)

var servicehealth = &cobra.Command{
	Use:   "servicehealth",
	Short: "Checks the health status of all subscribed or for on specific service for a tenant",
	Long: `Checks the health status of all subscribed or for on specific service for a tenant

For more information how to register an app see https://docs.microsoft.com/en-us/graph/auth-register-app-v2`,
	Example: " servicehealth -T \"1234-1234-1234\" -c \"0987-6543\" -s \"ngreiuj-wefnu4389IUerfnkefm\" -n \"Exchange online\"",
	Run: func(cmd *cobra.Command, args []string) {
		cliConfig.LoadEnv()

		c := cliConfig.Client()
		err := c.NewGraphServiceClient()
		if err != nil {
			check.ExitError(err)
		}

		override := buildStateOverrides()

		var (
			services *client.Services
			issues   *client.Issues
		)

		if cliConfig.ServiceNames != nil {
			services, err = c.LoadServices(cliConfig.ServiceNames)
		} else {
			services, err = c.LoadAllServices()
		}

		if cliConfig.DisplayIssueMessage {
			issues, err = c.LoadAllIssues(cliConfig.IssueStartTime)
		}

		if err != nil {
			check.ExitError(err)
		}

		summary += fmt.Sprintf("Found %d services - ", len(services.Services))

		for _, service := range services.Services {
			status := service.GetStatus(override)
			if status == 2 {
				totalCrit++
			} else if status == 1 {
				totalWarn++
			}
		}

		rc := services.GetStatus(override)
		output += services.GetOuput(override, cliConfig.ShowAll, issues, cliConfig.DisplayIssueMessage)

		summary += strconv.Itoa(totalCrit) + " Critical - " + " " + strconv.Itoa(totalWarn) + " Warning:\n"

		if rc == 0 && !cliConfig.ShowAll {
			output = "All services are healthy"
		}

		p := perfdata.PerfdataList{
			{Label: "total", Value: len(services.Services)},
			{Label: "warning", Value: totalWarn},
			{Label: "critical", Value: totalCrit},
		}

		check.ExitRaw(rc, summary+output, "|", p.String())
	},
}

func buildStateOverrides() client.StateOverride {
	override := client.StateOverride{}

	for _, o := range cliConfig.StateOverrides {
		pair := strings.SplitN(o, "=", 2)
		state := check.Unknown

		switch strings.ToLower(pair[1]) {
		case "ok":
			state = check.OK
		case "warning":
			state = check.Warning
		case "critical":
			state = check.Critical
		}

		override[pair[0]] = state
	}

	return override
}

func init() {
	rootCmd.AddCommand(servicehealth)

	fs := servicehealth.Flags()
	fs.StringSliceVarP(&cliConfig.ServiceNames, "servicenames", "n", nil,
		"The name of one or more specific service/s")
	fs.StringSliceVar(&cliConfig.StateOverrides, "state-override", nil,
		"States to override (e.g. STATENAME=ok)")
	fs.BoolVar(&cliConfig.ShowAll, "all", false, "Displays all services regardless of the status")
	fs.BoolVarP(&cliConfig.DisplayIssueMessage, "display-message", "M", false,
		"Displays the issue message to the specified service")
	fs.StringVarP(&cliConfig.IssueStartTime, "issue-start-time", "i", "",
		"Displays only issue massages in the period of time given. Possible values are e.G. '1h', '30m'")

	fs.SortFlags = false
	servicehealth.DisableFlagsInUseLine = true
}
