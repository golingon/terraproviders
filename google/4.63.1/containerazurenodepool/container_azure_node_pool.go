// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package containerazurenodepool

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Autoscaling struct {
	// MaxNodeCount: number, required
	MaxNodeCount terra.NumberValue `hcl:"max_node_count,attr" validate:"required"`
	// MinNodeCount: number, required
	MinNodeCount terra.NumberValue `hcl:"min_node_count,attr" validate:"required"`
}

type Config struct {
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VmSize: string, optional
	VmSize terra.StringValue `hcl:"vm_size,attr"`
	// ProxyConfig: optional
	ProxyConfig *ProxyConfig `hcl:"proxy_config,block"`
	// RootVolume: optional
	RootVolume *RootVolume `hcl:"root_volume,block"`
	// SshConfig: required
	SshConfig *SshConfig `hcl:"ssh_config,block" validate:"required"`
}

type ProxyConfig struct {
	// ResourceGroupId: string, required
	ResourceGroupId terra.StringValue `hcl:"resource_group_id,attr" validate:"required"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
}

type RootVolume struct {
	// SizeGib: number, optional
	SizeGib terra.NumberValue `hcl:"size_gib,attr"`
}

type SshConfig struct {
	// AuthorizedKey: string, required
	AuthorizedKey terra.StringValue `hcl:"authorized_key,attr" validate:"required"`
}

type MaxPodsConstraint struct {
	// MaxPodsPerNode: number, required
	MaxPodsPerNode terra.NumberValue `hcl:"max_pods_per_node,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AutoscalingAttributes struct {
	ref terra.Reference
}

func (a AutoscalingAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AutoscalingAttributes) InternalWithRef(ref terra.Reference) AutoscalingAttributes {
	return AutoscalingAttributes{ref: ref}
}

func (a AutoscalingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AutoscalingAttributes) MaxNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("max_node_count"))
}

func (a AutoscalingAttributes) MinNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("min_node_count"))
}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("tags"))
}

func (c ConfigAttributes) VmSize() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("vm_size"))
}

func (c ConfigAttributes) ProxyConfig() terra.ListValue[ProxyConfigAttributes] {
	return terra.ReferenceAsList[ProxyConfigAttributes](c.ref.Append("proxy_config"))
}

func (c ConfigAttributes) RootVolume() terra.ListValue[RootVolumeAttributes] {
	return terra.ReferenceAsList[RootVolumeAttributes](c.ref.Append("root_volume"))
}

func (c ConfigAttributes) SshConfig() terra.ListValue[SshConfigAttributes] {
	return terra.ReferenceAsList[SshConfigAttributes](c.ref.Append("ssh_config"))
}

type ProxyConfigAttributes struct {
	ref terra.Reference
}

func (pc ProxyConfigAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ProxyConfigAttributes) InternalWithRef(ref terra.Reference) ProxyConfigAttributes {
	return ProxyConfigAttributes{ref: ref}
}

func (pc ProxyConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ProxyConfigAttributes) ResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("resource_group_id"))
}

func (pc ProxyConfigAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("secret_id"))
}

type RootVolumeAttributes struct {
	ref terra.Reference
}

func (rv RootVolumeAttributes) InternalRef() (terra.Reference, error) {
	return rv.ref, nil
}

func (rv RootVolumeAttributes) InternalWithRef(ref terra.Reference) RootVolumeAttributes {
	return RootVolumeAttributes{ref: ref}
}

func (rv RootVolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rv.ref.InternalTokens()
}

func (rv RootVolumeAttributes) SizeGib() terra.NumberValue {
	return terra.ReferenceAsNumber(rv.ref.Append("size_gib"))
}

type SshConfigAttributes struct {
	ref terra.Reference
}

func (sc SshConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc SshConfigAttributes) InternalWithRef(ref terra.Reference) SshConfigAttributes {
	return SshConfigAttributes{ref: ref}
}

func (sc SshConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc SshConfigAttributes) AuthorizedKey() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("authorized_key"))
}

type MaxPodsConstraintAttributes struct {
	ref terra.Reference
}

func (mpc MaxPodsConstraintAttributes) InternalRef() (terra.Reference, error) {
	return mpc.ref, nil
}

func (mpc MaxPodsConstraintAttributes) InternalWithRef(ref terra.Reference) MaxPodsConstraintAttributes {
	return MaxPodsConstraintAttributes{ref: ref}
}

func (mpc MaxPodsConstraintAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mpc.ref.InternalTokens()
}

func (mpc MaxPodsConstraintAttributes) MaxPodsPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(mpc.ref.Append("max_pods_per_node"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AutoscalingState struct {
	MaxNodeCount float64 `json:"max_node_count"`
	MinNodeCount float64 `json:"min_node_count"`
}

type ConfigState struct {
	Tags        map[string]string  `json:"tags"`
	VmSize      string             `json:"vm_size"`
	ProxyConfig []ProxyConfigState `json:"proxy_config"`
	RootVolume  []RootVolumeState  `json:"root_volume"`
	SshConfig   []SshConfigState   `json:"ssh_config"`
}

type ProxyConfigState struct {
	ResourceGroupId string `json:"resource_group_id"`
	SecretId        string `json:"secret_id"`
}

type RootVolumeState struct {
	SizeGib float64 `json:"size_gib"`
}

type SshConfigState struct {
	AuthorizedKey string `json:"authorized_key"`
}

type MaxPodsConstraintState struct {
	MaxPodsPerNode float64 `json:"max_pods_per_node"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
