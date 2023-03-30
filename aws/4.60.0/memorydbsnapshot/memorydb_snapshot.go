// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package memorydbsnapshot

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClusterConfiguration struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type ClusterConfigurationAttributes struct {
	ref terra.Reference
}

func (cc ClusterConfigurationAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc ClusterConfigurationAttributes) InternalWithRef(ref terra.Reference) ClusterConfigurationAttributes {
	return ClusterConfigurationAttributes{ref: ref}
}

func (cc ClusterConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc ClusterConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("description"))
}

func (cc ClusterConfigurationAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("engine_version"))
}

func (cc ClusterConfigurationAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("maintenance_window"))
}

func (cc ClusterConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("name"))
}

func (cc ClusterConfigurationAttributes) NodeType() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("node_type"))
}

func (cc ClusterConfigurationAttributes) NumShards() terra.NumberValue {
	return terra.ReferenceNumber(cc.ref.Append("num_shards"))
}

func (cc ClusterConfigurationAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("parameter_group_name"))
}

func (cc ClusterConfigurationAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(cc.ref.Append("port"))
}

func (cc ClusterConfigurationAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceNumber(cc.ref.Append("snapshot_retention_limit"))
}

func (cc ClusterConfigurationAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("snapshot_window"))
}

func (cc ClusterConfigurationAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("subnet_group_name"))
}

func (cc ClusterConfigurationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("topic_arn"))
}

func (cc ClusterConfigurationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(cc.ref.Append("vpc_id"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

type ClusterConfigurationState struct {
	Description            string  `json:"description"`
	EngineVersion          string  `json:"engine_version"`
	MaintenanceWindow      string  `json:"maintenance_window"`
	Name                   string  `json:"name"`
	NodeType               string  `json:"node_type"`
	NumShards              float64 `json:"num_shards"`
	ParameterGroupName     string  `json:"parameter_group_name"`
	Port                   float64 `json:"port"`
	SnapshotRetentionLimit float64 `json:"snapshot_retention_limit"`
	SnapshotWindow         string  `json:"snapshot_window"`
	SubnetGroupName        string  `json:"subnet_group_name"`
	TopicArn               string  `json:"topic_arn"`
	VpcId                  string  `json:"vpc_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
