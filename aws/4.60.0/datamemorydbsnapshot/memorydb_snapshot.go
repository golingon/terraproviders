// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamemorydbsnapshot

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClusterConfiguration struct{}

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
	return terra.ReferenceAsString(cc.ref.Append("description"))
}

func (cc ClusterConfigurationAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("engine_version"))
}

func (cc ClusterConfigurationAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("maintenance_window"))
}

func (cc ClusterConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

func (cc ClusterConfigurationAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("node_type"))
}

func (cc ClusterConfigurationAttributes) NumShards() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("num_shards"))
}

func (cc ClusterConfigurationAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("parameter_group_name"))
}

func (cc ClusterConfigurationAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("port"))
}

func (cc ClusterConfigurationAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("snapshot_retention_limit"))
}

func (cc ClusterConfigurationAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("snapshot_window"))
}

func (cc ClusterConfigurationAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("subnet_group_name"))
}

func (cc ClusterConfigurationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("topic_arn"))
}

func (cc ClusterConfigurationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("vpc_id"))
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