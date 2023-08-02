package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_Route53DomainName(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "aws_route53_record" "invalid" {
    type = "A"
    name = "不適切なドメイン.example.com"
}
resource "aws_route53_record" "valid" {
    type = "A"
    name = "valid-domain.example.com"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewRoute53DomainNameRule(),
					Message: "Invalid name for record: 不適切なドメイン.example.com",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 12},
						End:      hcl.Pos{Line: 4, Column: 34},
					},
				},
			},
		},
	}

	rule := NewRoute53DomainNameRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
