// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_memorydb_snapshot

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataClusterConfigurationAttributes struct {
	ref terra.Reference
}

func (cc DataClusterConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc DataClusterConfigurationAttributes) InternalWithRef(ref terra.Reference) DataClusterConfigurationAttributes {
	return DataClusterConfigurationAttributes{ref: ref}
}

func (cc DataClusterConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc DataClusterConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("description"))
}

func (cc DataClusterConfigurationAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("engine_version"))
}

func (cc DataClusterConfigurationAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("maintenance_window"))
}

func (cc DataClusterConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("name"))
}

func (cc DataClusterConfigurationAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("node_type"))
}

func (cc DataClusterConfigurationAttributes) NumShards() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("num_shards"))
}

func (cc DataClusterConfigurationAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("parameter_group_name"))
}

func (cc DataClusterConfigurationAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("port"))
}

func (cc DataClusterConfigurationAttributes) SnapshotRetentionLimit() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("snapshot_retention_limit"))
}

func (cc DataClusterConfigurationAttributes) SnapshotWindow() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("snapshot_window"))
}

func (cc DataClusterConfigurationAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("subnet_group_name"))
}

func (cc DataClusterConfigurationAttributes) TopicArn() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("topic_arn"))
}

func (cc DataClusterConfigurationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("vpc_id"))
}

type DataClusterConfigurationState struct {
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
