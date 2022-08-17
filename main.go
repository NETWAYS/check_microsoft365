package main

import (
	"fmt"
	"github.com/NETWAYS/check_microsoft365/cmd"
)

var (
	version = "1.0.0"
	commit  = "102e653"
	date    = "04.04.2022"
)

func main() {
	cmd.Execute(buildVersion())
}

//goland:noinspection GoBoolExpressions
func buildVersion() (result string) {
	result = version

	if commit != "" {
		result += fmt.Sprintf("\ncommit: %s", commit)
	}

	if date != "" {
		result += fmt.Sprintf("\ndate: %s", date)
	}

	return result
}
