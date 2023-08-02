package main

import (
	"github.com/10sr/tflint-ruleset-domainname/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "domainname",
			Version: "0.0.3",
			Rules: []tflint.Rule{
				rules.NewRoute53DomainNameRule(),
			},
		},
	})
}
