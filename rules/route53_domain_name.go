package rules

import (
	"fmt"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"regexp"
)

type Route53DomainNameRule struct {
	tflint.DefaultRule
}

func NewRoute53DomainNameRule() *Route53DomainNameRule {
	return &Route53DomainNameRule{}
}

// Name returns the rule name
func (r *Route53DomainNameRule) Name() string {
	return "route53_domain_name"
}

// Enabled returns whether the rule is enabled by default
func (r *Route53DomainNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *Route53DomainNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *Route53DomainNameRule) Link() string {
	return "https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html#domain-name-format-hosted-zones"
}

func (r *Route53DomainNameRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("aws_route53_record", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "name"},
		},
	}, nil)
	if err != nil {
		return err
	}

	// Put a log that can be output with `TFLINT_LOG=debug`
	logger.Debug(fmt.Sprintf("Get %d records", len(resources.Blocks)))

	// TODO: Update re
	nameRegexp := regexp.MustCompile(`^[a-z0-9-_.*]+$`)

	for _, resource := range resources.Blocks {
		nameAttribute, _ := resource.Body.Attributes["name"]
		err := runner.EvaluateExpr(nameAttribute.Expr, func(name string) error {
			if nameRegexp.MatchString(name) {
				// OK
				return nil
			}
			return runner.EmitIssue(
				r,
				fmt.Sprintf("Invalid name for record: %s", name),
				nameAttribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}

	}

	return nil
}
