// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigtable_instance

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Cluster struct {
	// ClusterId: string, required
	ClusterId terra.StringValue `hcl:"cluster_id,attr" validate:"required"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// NumNodes: number, optional
	NumNodes terra.NumberValue `hcl:"num_nodes,attr"`
	// StorageType: string, optional
	StorageType terra.StringValue `hcl:"storage_type,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// ClusterAutoscalingConfig: optional
	AutoscalingConfig *ClusterAutoscalingConfig `hcl:"autoscaling_config,block"`
}

type ClusterAutoscalingConfig struct {
	// CpuTarget: number, required
	CpuTarget terra.NumberValue `hcl:"cpu_target,attr" validate:"required"`
	// MaxNodes: number, required
	MaxNodes terra.NumberValue `hcl:"max_nodes,attr" validate:"required"`
	// MinNodes: number, required
	MinNodes terra.NumberValue `hcl:"min_nodes,attr" validate:"required"`
	// StorageTarget: number, optional
	StorageTarget terra.NumberValue `hcl:"storage_target,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ClusterAttributes struct {
	ref terra.Reference
}

func (c ClusterAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ClusterAttributes) InternalWithRef(ref terra.Reference) ClusterAttributes {
	return ClusterAttributes{ref: ref}
}

func (c ClusterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("cluster_id"))
}

func (c ClusterAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("kms_key_name"))
}

func (c ClusterAttributes) NumNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(c.ref.Append("num_nodes"))
}

func (c ClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("state"))
}

func (c ClusterAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("storage_type"))
}

func (c ClusterAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("zone"))
}

func (c ClusterAttributes) AutoscalingConfig() terra.ListValue[ClusterAutoscalingConfigAttributes] {
	return terra.ReferenceAsList[ClusterAutoscalingConfigAttributes](c.ref.Append("autoscaling_config"))
}

type ClusterAutoscalingConfigAttributes struct {
	ref terra.Reference
}

func (ac ClusterAutoscalingConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac ClusterAutoscalingConfigAttributes) InternalWithRef(ref terra.Reference) ClusterAutoscalingConfigAttributes {
	return ClusterAutoscalingConfigAttributes{ref: ref}
}

func (ac ClusterAutoscalingConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac ClusterAutoscalingConfigAttributes) CpuTarget() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("cpu_target"))
}

func (ac ClusterAutoscalingConfigAttributes) MaxNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("max_nodes"))
}

func (ac ClusterAutoscalingConfigAttributes) MinNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("min_nodes"))
}

func (ac ClusterAutoscalingConfigAttributes) StorageTarget() terra.NumberValue {
	return terra.ReferenceAsNumber(ac.ref.Append("storage_target"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ClusterState struct {
	ClusterId         string                          `json:"cluster_id"`
	KmsKeyName        string                          `json:"kms_key_name"`
	NumNodes          float64                         `json:"num_nodes"`
	State             string                          `json:"state"`
	StorageType       string                          `json:"storage_type"`
	Zone              string                          `json:"zone"`
	AutoscalingConfig []ClusterAutoscalingConfigState `json:"autoscaling_config"`
}

type ClusterAutoscalingConfigState struct {
	CpuTarget     float64 `json:"cpu_target"`
	MaxNodes      float64 `json:"max_nodes"`
	MinNodes      float64 `json:"min_nodes"`
	StorageTarget float64 `json:"storage_target"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
