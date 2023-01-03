package main

import (
	"fmt"
	"github.com/NETWAYS/check_microsoft365/cmd"
)

var (
	// These get filled at build time with the proper vaules
	version = "development"
	commit  = "HEAD"
	date    = "latest"
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
