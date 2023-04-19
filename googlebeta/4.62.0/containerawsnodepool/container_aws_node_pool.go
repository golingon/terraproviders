// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package containerawsnodepool

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
	// IamInstanceProfile: string, required
	IamInstanceProfile terra.StringValue `hcl:"iam_instance_profile,attr" validate:"required"`
	// ImageType: string, optional
	ImageType terra.StringValue `hcl:"image_type,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// SecurityGroupIds: list of string, optional
	SecurityGroupIds terra.ListValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AutoscalingMetricsCollection: optional
	AutoscalingMetricsCollection *AutoscalingMetricsCollection `hcl:"autoscaling_metrics_collection,block"`
	// ConfigEncryption: required
	ConfigEncryption *ConfigEncryption `hcl:"config_encryption,block" validate:"required"`
	// InstancePlacement: optional
	InstancePlacement *InstancePlacement `hcl:"instance_placement,block"`
	// ProxyConfig: optional
	ProxyConfig *ProxyConfig `hcl:"proxy_config,block"`
	// RootVolume: optional
	RootVolume *RootVolume `hcl:"root_volume,block"`
	// SshConfig: optional
	SshConfig *SshConfig `hcl:"ssh_config,block"`
	// Taints: min=0
	Taints []Taints `hcl:"taints,block" validate:"min=0"`
}

type AutoscalingMetricsCollection struct {
	// Granularity: string, required
	Granularity terra.StringValue `hcl:"granularity,attr" validate:"required"`
	// Metrics: list of string, optional
	Metrics terra.ListValue[terra.StringValue] `hcl:"metrics,attr"`
}

type ConfigEncryption struct {
	// KmsKeyArn: string, required
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr" validate:"required"`
}

type InstancePlacement struct {
	// Tenancy: string, optional
	Tenancy terra.StringValue `hcl:"tenancy,attr"`
}

type ProxyConfig struct {
	// SecretArn: string, required
	SecretArn terra.StringValue `hcl:"secret_arn,attr" validate:"required"`
	// SecretVersion: string, required
	SecretVersion terra.StringValue `hcl:"secret_version,attr" validate:"required"`
}

type RootVolume struct {
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// SizeGib: number, optional
	SizeGib terra.NumberValue `hcl:"size_gib,attr"`
	// VolumeType: string, optional
	VolumeType terra.StringValue `hcl:"volume_type,attr"`
}

type SshConfig struct {
	// Ec2KeyPair: string, required
	Ec2KeyPair terra.StringValue `hcl:"ec2_key_pair,attr" validate:"required"`
}

type Taints struct {
	// Effect: string, required
	Effect terra.StringValue `hcl:"effect,attr" validate:"required"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

func (c ConfigAttributes) IamInstanceProfile() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("iam_instance_profile"))
}

func (c ConfigAttributes) ImageType() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("image_type"))
}

func (c ConfigAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("instance_type"))
}

func (c ConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("labels"))
}

func (c ConfigAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("security_group_ids"))
}

func (c ConfigAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("tags"))
}

func (c ConfigAttributes) AutoscalingMetricsCollection() terra.ListValue[AutoscalingMetricsCollectionAttributes] {
	return terra.ReferenceAsList[AutoscalingMetricsCollectionAttributes](c.ref.Append("autoscaling_metrics_collection"))
}

func (c ConfigAttributes) ConfigEncryption() terra.ListValue[ConfigEncryptionAttributes] {
	return terra.ReferenceAsList[ConfigEncryptionAttributes](c.ref.Append("config_encryption"))
}

func (c ConfigAttributes) InstancePlacement() terra.ListValue[InstancePlacementAttributes] {
	return terra.ReferenceAsList[InstancePlacementAttributes](c.ref.Append("instance_placement"))
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

func (c ConfigAttributes) Taints() terra.ListValue[TaintsAttributes] {
	return terra.ReferenceAsList[TaintsAttributes](c.ref.Append("taints"))
}

type AutoscalingMetricsCollectionAttributes struct {
	ref terra.Reference
}

func (amc AutoscalingMetricsCollectionAttributes) InternalRef() (terra.Reference, error) {
	return amc.ref, nil
}

func (amc AutoscalingMetricsCollectionAttributes) InternalWithRef(ref terra.Reference) AutoscalingMetricsCollectionAttributes {
	return AutoscalingMetricsCollectionAttributes{ref: ref}
}

func (amc AutoscalingMetricsCollectionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return amc.ref.InternalTokens()
}

func (amc AutoscalingMetricsCollectionAttributes) Granularity() terra.StringValue {
	return terra.ReferenceAsString(amc.ref.Append("granularity"))
}

func (amc AutoscalingMetricsCollectionAttributes) Metrics() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](amc.ref.Append("metrics"))
}

type ConfigEncryptionAttributes struct {
	ref terra.Reference
}

func (ce ConfigEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return ce.ref, nil
}

func (ce ConfigEncryptionAttributes) InternalWithRef(ref terra.Reference) ConfigEncryptionAttributes {
	return ConfigEncryptionAttributes{ref: ref}
}

func (ce ConfigEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ce.ref.InternalTokens()
}

func (ce ConfigEncryptionAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("kms_key_arn"))
}

type InstancePlacementAttributes struct {
	ref terra.Reference
}

func (ip InstancePlacementAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip InstancePlacementAttributes) InternalWithRef(ref terra.Reference) InstancePlacementAttributes {
	return InstancePlacementAttributes{ref: ref}
}

func (ip InstancePlacementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip InstancePlacementAttributes) Tenancy() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("tenancy"))
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

func (pc ProxyConfigAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("secret_arn"))
}

func (pc ProxyConfigAttributes) SecretVersion() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("secret_version"))
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

func (rv RootVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(rv.ref.Append("iops"))
}

func (rv RootVolumeAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(rv.ref.Append("kms_key_arn"))
}

func (rv RootVolumeAttributes) SizeGib() terra.NumberValue {
	return terra.ReferenceAsNumber(rv.ref.Append("size_gib"))
}

func (rv RootVolumeAttributes) VolumeType() terra.StringValue {
	return terra.ReferenceAsString(rv.ref.Append("volume_type"))
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

func (sc SshConfigAttributes) Ec2KeyPair() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("ec2_key_pair"))
}

type TaintsAttributes struct {
	ref terra.Reference
}

func (t TaintsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TaintsAttributes) InternalWithRef(ref terra.Reference) TaintsAttributes {
	return TaintsAttributes{ref: ref}
}

func (t TaintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TaintsAttributes) Effect() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("effect"))
}

func (t TaintsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TaintsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("value"))
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
	IamInstanceProfile           string                              `json:"iam_instance_profile"`
	ImageType                    string                              `json:"image_type"`
	InstanceType                 string                              `json:"instance_type"`
	Labels                       map[string]string                   `json:"labels"`
	SecurityGroupIds             []string                            `json:"security_group_ids"`
	Tags                         map[string]string                   `json:"tags"`
	AutoscalingMetricsCollection []AutoscalingMetricsCollectionState `json:"autoscaling_metrics_collection"`
	ConfigEncryption             []ConfigEncryptionState             `json:"config_encryption"`
	InstancePlacement            []InstancePlacementState            `json:"instance_placement"`
	ProxyConfig                  []ProxyConfigState                  `json:"proxy_config"`
	RootVolume                   []RootVolumeState                   `json:"root_volume"`
	SshConfig                    []SshConfigState                    `json:"ssh_config"`
	Taints                       []TaintsState                       `json:"taints"`
}

type AutoscalingMetricsCollectionState struct {
	Granularity string   `json:"granularity"`
	Metrics     []string `json:"metrics"`
}

type ConfigEncryptionState struct {
	KmsKeyArn string `json:"kms_key_arn"`
}

type InstancePlacementState struct {
	Tenancy string `json:"tenancy"`
}

type ProxyConfigState struct {
	SecretArn     string `json:"secret_arn"`
	SecretVersion string `json:"secret_version"`
}

type RootVolumeState struct {
	Iops       float64 `json:"iops"`
	KmsKeyArn  string  `json:"kms_key_arn"`
	SizeGib    float64 `json:"size_gib"`
	VolumeType string  `json:"volume_type"`
}

type SshConfigState struct {
	Ec2KeyPair string `json:"ec2_key_pair"`
}

type TaintsState struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

type MaxPodsConstraintState struct {
	MaxPodsPerNode float64 `json:"max_pods_per_node"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}