// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package prometheusworkspace

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LoggingConfiguration struct {
	// LogGroupArn: string, required
	LogGroupArn terra.StringValue `hcl:"log_group_arn,attr" validate:"required"`
}

type LoggingConfigurationAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigurationAttributes) InternalRef() terra.Reference {
	return lc.ref
}

func (lc LoggingConfigurationAttributes) InternalWithRef(ref terra.Reference) LoggingConfigurationAttributes {
	return LoggingConfigurationAttributes{ref: ref}
}

func (lc LoggingConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigurationAttributes) LogGroupArn() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_group_arn"))
}

type LoggingConfigurationState struct {
	LogGroupArn string `json:"log_group_arn"`
}