// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package glueconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PhysicalConnectionRequirements struct {
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// SecurityGroupIdList: set of string, optional
	SecurityGroupIdList terra.SetValue[terra.StringValue] `hcl:"security_group_id_list,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
}

type PhysicalConnectionRequirementsAttributes struct {
	ref terra.Reference
}

func (pcr PhysicalConnectionRequirementsAttributes) InternalRef() (terra.Reference, error) {
	return pcr.ref, nil
}

func (pcr PhysicalConnectionRequirementsAttributes) InternalWithRef(ref terra.Reference) PhysicalConnectionRequirementsAttributes {
	return PhysicalConnectionRequirementsAttributes{ref: ref}
}

func (pcr PhysicalConnectionRequirementsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pcr.ref.InternalTokens()
}

func (pcr PhysicalConnectionRequirementsAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(pcr.ref.Append("availability_zone"))
}

func (pcr PhysicalConnectionRequirementsAttributes) SecurityGroupIdList() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pcr.ref.Append("security_group_id_list"))
}

func (pcr PhysicalConnectionRequirementsAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(pcr.ref.Append("subnet_id"))
}

type PhysicalConnectionRequirementsState struct {
	AvailabilityZone    string   `json:"availability_zone"`
	SecurityGroupIdList []string `json:"security_group_id_list"`
	SubnetId            string   `json:"subnet_id"`
}
