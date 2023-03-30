// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataglueconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PhysicalConnectionRequirements struct{}

type PhysicalConnectionRequirementsAttributes struct {
	ref terra.Reference
}

func (pcr PhysicalConnectionRequirementsAttributes) InternalRef() terra.Reference {
	return pcr.ref
}

func (pcr PhysicalConnectionRequirementsAttributes) InternalWithRef(ref terra.Reference) PhysicalConnectionRequirementsAttributes {
	return PhysicalConnectionRequirementsAttributes{ref: ref}
}

func (pcr PhysicalConnectionRequirementsAttributes) InternalTokens() hclwrite.Tokens {
	return pcr.ref.InternalTokens()
}

func (pcr PhysicalConnectionRequirementsAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceString(pcr.ref.Append("availability_zone"))
}

func (pcr PhysicalConnectionRequirementsAttributes) SecurityGroupIdList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](pcr.ref.Append("security_group_id_list"))
}

func (pcr PhysicalConnectionRequirementsAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(pcr.ref.Append("subnet_id"))
}

type PhysicalConnectionRequirementsState struct {
	AvailabilityZone    string   `json:"availability_zone"`
	SecurityGroupIdList []string `json:"security_group_id_list"`
	SubnetId            string   `json:"subnet_id"`
}
