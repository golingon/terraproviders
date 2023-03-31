// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appconfigenvironment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Monitor struct {
	// AlarmArn: string, required
	AlarmArn terra.StringValue `hcl:"alarm_arn,attr" validate:"required"`
	// AlarmRoleArn: string, optional
	AlarmRoleArn terra.StringValue `hcl:"alarm_role_arn,attr"`
}

type MonitorAttributes struct {
	ref terra.Reference
}

func (m MonitorAttributes) InternalRef() terra.Reference {
	return m.ref
}

func (m MonitorAttributes) InternalWithRef(ref terra.Reference) MonitorAttributes {
	return MonitorAttributes{ref: ref}
}

func (m MonitorAttributes) InternalTokens() hclwrite.Tokens {
	return m.ref.InternalTokens()
}

func (m MonitorAttributes) AlarmArn() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("alarm_arn"))
}

func (m MonitorAttributes) AlarmRoleArn() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("alarm_role_arn"))
}

type MonitorState struct {
	AlarmArn     string `json:"alarm_arn"`
	AlarmRoleArn string `json:"alarm_role_arn"`
}
