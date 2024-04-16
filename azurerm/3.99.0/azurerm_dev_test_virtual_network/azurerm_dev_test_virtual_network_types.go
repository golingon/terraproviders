// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_dev_test_virtual_network

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Subnet struct {
	// UseInVirtualMachineCreation: string, optional
	UseInVirtualMachineCreation terra.StringValue `hcl:"use_in_virtual_machine_creation,attr"`
	// UsePublicIpAddress: string, optional
	UsePublicIpAddress terra.StringValue `hcl:"use_public_ip_address,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SubnetAttributes struct {
	ref terra.Reference
}

func (s SubnetAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SubnetAttributes) InternalWithRef(ref terra.Reference) SubnetAttributes {
	return SubnetAttributes{ref: ref}
}

func (s SubnetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SubnetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SubnetAttributes) UseInVirtualMachineCreation() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("use_in_virtual_machine_creation"))
}

func (s SubnetAttributes) UsePublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("use_public_ip_address"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type SubnetState struct {
	Name                        string `json:"name"`
	UseInVirtualMachineCreation string `json:"use_in_virtual_machine_creation"`
	UsePublicIpAddress          string `json:"use_public_ip_address"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
