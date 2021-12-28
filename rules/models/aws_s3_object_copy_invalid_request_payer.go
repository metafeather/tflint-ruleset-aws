// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3ObjectCopyInvalidRequestPayerRule checks the pattern is valid
type AwsS3ObjectCopyInvalidRequestPayerRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3ObjectCopyInvalidRequestPayerRule returns new rule with default attributes
func NewAwsS3ObjectCopyInvalidRequestPayerRule() *AwsS3ObjectCopyInvalidRequestPayerRule {
	return &AwsS3ObjectCopyInvalidRequestPayerRule{
		resourceType:  "aws_s3_object_copy",
		attributeName: "request_payer",
		enum: []string{
			"requester",
		},
	}
}

// Name returns the rule name
func (r *AwsS3ObjectCopyInvalidRequestPayerRule) Name() string {
	return "aws_s3_object_copy_invalid_request_payer"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3ObjectCopyInvalidRequestPayerRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3ObjectCopyInvalidRequestPayerRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3ObjectCopyInvalidRequestPayerRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3ObjectCopyInvalidRequestPayerRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as request_payer`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
