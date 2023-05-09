// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package nginxdeployment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type FrontendPrivate struct {
	// AllocationMethod: string, required
	AllocationMethod terra.StringValue `hcl:"allocation_method,attr" validate:"required"`
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type FrontendPublic struct {
	// IpAddress: list of string, optional
	IpAddress terra.ListValue[terra.StringValue] `hcl:"ip_address,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type LoggingStorageAccount struct {
	// ContainerName: string, optional
	ContainerName terra.StringValue `hcl:"container_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type NetworkInterface struct {
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
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

type FrontendPrivateAttributes struct {
	ref terra.Reference
}

func (fp FrontendPrivateAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FrontendPrivateAttributes) InternalWithRef(ref terra.Reference) FrontendPrivateAttributes {
	return FrontendPrivateAttributes{ref: ref}
}

func (fp FrontendPrivateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FrontendPrivateAttributes) AllocationMethod() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("allocation_method"))
}

func (fp FrontendPrivateAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("ip_address"))
}

func (fp FrontendPrivateAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("subnet_id"))
}

type FrontendPublicAttributes struct {
	ref terra.Reference
}

func (fp FrontendPublicAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FrontendPublicAttributes) InternalWithRef(ref terra.Reference) FrontendPublicAttributes {
	return FrontendPublicAttributes{ref: ref}
}

func (fp FrontendPublicAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FrontendPublicAttributes) IpAddress() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("ip_address"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type LoggingStorageAccountAttributes struct {
	ref terra.Reference
}

func (lsa LoggingStorageAccountAttributes) InternalRef() (terra.Reference, error) {
	return lsa.ref, nil
}

func (lsa LoggingStorageAccountAttributes) InternalWithRef(ref terra.Reference) LoggingStorageAccountAttributes {
	return LoggingStorageAccountAttributes{ref: ref}
}

func (lsa LoggingStorageAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lsa.ref.InternalTokens()
}

func (lsa LoggingStorageAccountAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(lsa.ref.Append("container_name"))
}

func (lsa LoggingStorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsa.ref.Append("name"))
}

type NetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni NetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni NetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAttributes {
	return NetworkInterfaceAttributes{ref: ref}
}

func (ni NetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni NetworkInterfaceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet_id"))
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

type FrontendPrivateState struct {
	AllocationMethod string `json:"allocation_method"`
	IpAddress        string `json:"ip_address"`
	SubnetId         string `json:"subnet_id"`
}

type FrontendPublicState struct {
	IpAddress []string `json:"ip_address"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type LoggingStorageAccountState struct {
	ContainerName string `json:"container_name"`
	Name          string `json:"name"`
}

type NetworkInterfaceState struct {
	SubnetId string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
