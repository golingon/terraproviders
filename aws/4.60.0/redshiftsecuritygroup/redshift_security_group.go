// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package redshiftsecuritygroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Ingress struct {
	// Cidr: string, optional
	Cidr terra.StringValue `hcl:"cidr,attr"`
	// SecurityGroupName: string, optional
	SecurityGroupName terra.StringValue `hcl:"security_group_name,attr"`
	// SecurityGroupOwnerId: string, optional
	SecurityGroupOwnerId terra.StringValue `hcl:"security_group_owner_id,attr"`
}

type IngressAttributes struct {
	ref terra.Reference
}

func (i IngressAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IngressAttributes) InternalWithRef(ref terra.Reference) IngressAttributes {
	return IngressAttributes{ref: ref}
}

func (i IngressAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IngressAttributes) Cidr() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("cidr"))
}

func (i IngressAttributes) SecurityGroupName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("security_group_name"))
}

func (i IngressAttributes) SecurityGroupOwnerId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("security_group_owner_id"))
}

type IngressState struct {
	Cidr                 string `json:"cidr"`
	SecurityGroupName    string `json:"security_group_name"`
	SecurityGroupOwnerId string `json:"security_group_owner_id"`
}
