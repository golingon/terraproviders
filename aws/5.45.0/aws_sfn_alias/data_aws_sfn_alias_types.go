// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_sfn_alias

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataRoutingConfigurationAttributes struct {
	ref terra.Reference
}

func (rc DataRoutingConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc DataRoutingConfigurationAttributes) InternalWithRef(ref terra.Reference) DataRoutingConfigurationAttributes {
	return DataRoutingConfigurationAttributes{ref: ref}
}

func (rc DataRoutingConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc DataRoutingConfigurationAttributes) StateMachineVersionArn() terra.StringValue {
	return terra.ReferenceAsString(rc.ref.Append("state_machine_version_arn"))
}

func (rc DataRoutingConfigurationAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(rc.ref.Append("weight"))
}

type DataRoutingConfigurationState struct {
	StateMachineVersionArn string  `json:"state_machine_version_arn"`
	Weight                 float64 `json:"weight"`
}
