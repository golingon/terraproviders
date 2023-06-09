// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudformationtype

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type LoggingConfig struct {
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
	// LogRoleArn: string, required
	LogRoleArn terra.StringValue `hcl:"log_role_arn,attr" validate:"required"`
}

type LoggingConfigAttributes struct {
	ref terra.Reference
}

func (lc LoggingConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LoggingConfigAttributes) InternalWithRef(ref terra.Reference) LoggingConfigAttributes {
	return LoggingConfigAttributes{ref: ref}
}

func (lc LoggingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LoggingConfigAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_group_name"))
}

func (lc LoggingConfigAttributes) LogRoleArn() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_role_arn"))
}

type LoggingConfigState struct {
	LogGroupName string `json:"log_group_name"`
	LogRoleArn   string `json:"log_role_arn"`
}
