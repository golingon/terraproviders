// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kustocluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type OptimizedAutoScale struct {
	// MaximumInstances: number, required
	MaximumInstances terra.NumberValue `hcl:"maximum_instances,attr" validate:"required"`
	// MinimumInstances: number, required
	MinimumInstances terra.NumberValue `hcl:"minimum_instances,attr" validate:"required"`
}

type Sku struct {
	// Capacity: number, optional
	Capacity terra.NumberValue `hcl:"capacity,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type VirtualNetworkConfiguration struct {
	// DataManagementPublicIpId: string, required
	DataManagementPublicIpId terra.StringValue `hcl:"data_management_public_ip_id,attr" validate:"required"`
	// EnginePublicIpId: string, required
	EnginePublicIpId terra.StringValue `hcl:"engine_public_ip_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
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

type OptimizedAutoScaleAttributes struct {
	ref terra.Reference
}

func (oas OptimizedAutoScaleAttributes) InternalRef() (terra.Reference, error) {
	return oas.ref, nil
}

func (oas OptimizedAutoScaleAttributes) InternalWithRef(ref terra.Reference) OptimizedAutoScaleAttributes {
	return OptimizedAutoScaleAttributes{ref: ref}
}

func (oas OptimizedAutoScaleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oas.ref.InternalTokens()
}

func (oas OptimizedAutoScaleAttributes) MaximumInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(oas.ref.Append("maximum_instances"))
}

func (oas OptimizedAutoScaleAttributes) MinimumInstances() terra.NumberValue {
	return terra.ReferenceAsNumber(oas.ref.Append("minimum_instances"))
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s SkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
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

type VirtualNetworkConfigurationAttributes struct {
	ref terra.Reference
}

func (vnc VirtualNetworkConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return vnc.ref, nil
}

func (vnc VirtualNetworkConfigurationAttributes) InternalWithRef(ref terra.Reference) VirtualNetworkConfigurationAttributes {
	return VirtualNetworkConfigurationAttributes{ref: ref}
}

func (vnc VirtualNetworkConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vnc.ref.InternalTokens()
}

func (vnc VirtualNetworkConfigurationAttributes) DataManagementPublicIpId() terra.StringValue {
	return terra.ReferenceAsString(vnc.ref.Append("data_management_public_ip_id"))
}

func (vnc VirtualNetworkConfigurationAttributes) EnginePublicIpId() terra.StringValue {
	return terra.ReferenceAsString(vnc.ref.Append("engine_public_ip_id"))
}

func (vnc VirtualNetworkConfigurationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(vnc.ref.Append("subnet_id"))
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type OptimizedAutoScaleState struct {
	MaximumInstances float64 `json:"maximum_instances"`
	MinimumInstances float64 `json:"minimum_instances"`
}

type SkuState struct {
	Capacity float64 `json:"capacity"`
	Name     string  `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type VirtualNetworkConfigurationState struct {
	DataManagementPublicIpId string `json:"data_management_public_ip_id"`
	EnginePublicIpId         string `json:"engine_public_ip_id"`
	SubnetId                 string `json:"subnet_id"`
}
