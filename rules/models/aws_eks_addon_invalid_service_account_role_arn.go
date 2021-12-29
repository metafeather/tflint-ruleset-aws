// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsEksAddonInvalidServiceAccountRoleArnRule checks the pattern is valid
type AwsEksAddonInvalidServiceAccountRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsEksAddonInvalidServiceAccountRoleArnRule returns new rule with default attributes
func NewAwsEksAddonInvalidServiceAccountRoleArnRule() *AwsEksAddonInvalidServiceAccountRoleArnRule {
	return &AwsEksAddonInvalidServiceAccountRoleArnRule{
		resourceType:  "aws_eks_addon",
		attributeName: "service_account_role_arn",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsEksAddonInvalidServiceAccountRoleArnRule) Name() string {
	return "aws_eks_addon_invalid_service_account_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEksAddonInvalidServiceAccountRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEksAddonInvalidServiceAccountRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEksAddonInvalidServiceAccountRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEksAddonInvalidServiceAccountRoleArnRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"service_account_role_arn must be 255 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"service_account_role_arn must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
