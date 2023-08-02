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
			Name: "basic test",
			Content: `
resource "aws_route53_record" "invalid" {
    type = "A"
    name = "invalid domain.example.com"
}
resource "aws_route53_record" "valid" {
    type = "A"
    name = "valid-domain.example.com"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewRoute53DomainNameRule(),
					Message: "Invalid name for record: invalid domain.example.com",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 12},
						End:      hcl.Pos{Line: 4, Column: 40},
					},
				},
			},
		},
		{
			Name: "no issue",
			Content: `
resource "aws_route53_record" "underscore" {
    type = "A"
    name = "_valid.example.com"
}
resource "aws_route53_record" "asterisk" {
    type = "A"
    name = "*.example.com"
}
`,
			Expected: helper.Issues{},
		},
		{
			Name: "issue found",
			Content: `
# non-ascii character
# https://codepoints.net/U+2013
resource "aws_route53_record" "en_dash" {
    type = "A"
    name = "invalid–domain.example.com"
}
# Japanese letters
resource "aws_route53_record" "ja" {
    type = "A"
    name = "ドメイン.example.com"
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewRoute53DomainNameRule(),
					Message: "Invalid name for record: invalid–domain.example.com",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 12},
						End:      hcl.Pos{Line: 6, Column: 40},
					},
				},
				{
					Rule:    NewRoute53DomainNameRule(),
					Message: "Invalid name for record: ドメイン.example.com",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 11, Column: 12},
						End:      hcl.Pos{Line: 11, Column: 30},
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
