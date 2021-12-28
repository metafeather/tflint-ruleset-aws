// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule checks the pattern is valid
type AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketIntelligentTieringConfigurationInvalidStatusRule returns new rule with default attributes
func NewAwsS3BucketIntelligentTieringConfigurationInvalidStatusRule() *AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule {
	return &AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule{
		resourceType:  "aws_s3_bucket_intelligent_tiering_configuration",
		attributeName: "status",
		enum: []string{
			"Enabled",
			"Disabled",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule) Name() string {
	return "aws_s3_bucket_intelligent_tiering_configuration_invalid_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketIntelligentTieringConfigurationInvalidStatusRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as status`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
