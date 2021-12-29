// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsConnectContactFlowInvalidTypeRule checks the pattern is valid
type AwsConnectContactFlowInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsConnectContactFlowInvalidTypeRule returns new rule with default attributes
func NewAwsConnectContactFlowInvalidTypeRule() *AwsConnectContactFlowInvalidTypeRule {
	return &AwsConnectContactFlowInvalidTypeRule{
		resourceType:  "aws_connect_contact_flow",
		attributeName: "type",
		enum: []string{
			"CONTACT_FLOW",
			"CUSTOMER_QUEUE",
			"CUSTOMER_HOLD",
			"CUSTOMER_WHISPER",
			"AGENT_HOLD",
			"AGENT_WHISPER",
			"OUTBOUND_WHISPER",
			"AGENT_TRANSFER",
			"QUEUE_TRANSFER",
		},
	}
}

// Name returns the rule name
func (r *AwsConnectContactFlowInvalidTypeRule) Name() string {
	return "aws_connect_contact_flow_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConnectContactFlowInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConnectContactFlowInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConnectContactFlowInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConnectContactFlowInvalidTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
