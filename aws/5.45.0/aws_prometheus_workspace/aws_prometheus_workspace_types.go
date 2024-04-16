// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_prometheus_workspace

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type LoggingConfiguration struct {
	// LogGroupArn: string, required
	LogGroupArn terra.StringValue `hcl:"log_group_arn,attr" validate:"required"`
}

type LoggingConfigurationAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LoggingConfigurationAttributes) InternalWithRef(ref terra.Reference) LoggingConfigurationAttributes {
	return LoggingConfigurationAttributes{ref: ref}
}

func (lc LoggingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigurationAttributes) LogGroupArn() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_group_arn"))
}

type LoggingConfigurationState struct {
	LogGroupArn string `json:"log_group_arn"`
}
