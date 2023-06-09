// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	daxcluster "github.com/golingon/terraproviders/aws/4.66.1/daxcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDaxCluster creates a new instance of [DaxCluster].
func NewDaxCluster(name string, args DaxClusterArgs) *DaxCluster {
	return &DaxCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DaxCluster)(nil)

// DaxCluster represents the Terraform resource aws_dax_cluster.
type DaxCluster struct {
	Name      string
	Args      DaxClusterArgs
	state     *daxClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DaxCluster].
func (dc *DaxCluster) Type() string {
	return "aws_dax_cluster"
}

// LocalName returns the local name for [DaxCluster].
func (dc *DaxCluster) LocalName() string {
	return dc.Name
}

// Configuration returns the configuration (args) for [DaxCluster].
func (dc *DaxCluster) Configuration() interface{} {
	return dc.Args
}

// DependOn is used for other resources to depend on [DaxCluster].
func (dc *DaxCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(dc)
}

// Dependencies returns the list of resources [DaxCluster] depends_on.
func (dc *DaxCluster) Dependencies() terra.Dependencies {
	return dc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DaxCluster].
func (dc *DaxCluster) LifecycleManagement() *terra.Lifecycle {
	return dc.Lifecycle
}

// Attributes returns the attributes for [DaxCluster].
func (dc *DaxCluster) Attributes() daxClusterAttributes {
	return daxClusterAttributes{ref: terra.ReferenceResource(dc)}
}

// ImportState imports the given attribute values into [DaxCluster]'s state.
func (dc *DaxCluster) ImportState(av io.Reader) error {
	dc.state = &daxClusterState{}
	if err := json.NewDecoder(av).Decode(dc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dc.Type(), dc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DaxCluster] has state.
func (dc *DaxCluster) State() (*daxClusterState, bool) {
	return dc.state, dc.state != nil
}

// StateMust returns the state for [DaxCluster]. Panics if the state is nil.
func (dc *DaxCluster) StateMust() *daxClusterState {
	if dc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dc.Type(), dc.LocalName()))
	}
	return dc.state
}

// DaxClusterArgs contains the configurations for aws_dax_cluster.
type DaxClusterArgs struct {
	// AvailabilityZones: set of string, optional
	AvailabilityZones terra.SetValue[terra.StringValue] `hcl:"availability_zones,attr"`
	// ClusterEndpointEncryptionType: string, optional
	ClusterEndpointEncryptionType terra.StringValue `hcl:"cluster_endpoint_encryption_type,attr"`
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// IamRoleArn: string, required
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaintenanceWindow: string, optional
	MaintenanceWindow terra.StringValue `hcl:"maintenance_window,attr"`
	// NodeType: string, required
	NodeType terra.StringValue `hcl:"node_type,attr" validate:"required"`
	// NotificationTopicArn: string, optional
	NotificationTopicArn terra.StringValue `hcl:"notification_topic_arn,attr"`
	// ParameterGroupName: string, optional
	ParameterGroupName terra.StringValue `hcl:"parameter_group_name,attr"`
	// ReplicationFactor: number, required
	ReplicationFactor terra.NumberValue `hcl:"replication_factor,attr" validate:"required"`
	// SecurityGroupIds: set of string, optional
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// SubnetGroupName: string, optional
	SubnetGroupName terra.StringValue `hcl:"subnet_group_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Nodes: min=0
	Nodes []daxcluster.Nodes `hcl:"nodes,block" validate:"min=0"`
	// ServerSideEncryption: optional
	ServerSideEncryption *daxcluster.ServerSideEncryption `hcl:"server_side_encryption,block"`
	// Timeouts: optional
	Timeouts *daxcluster.Timeouts `hcl:"timeouts,block"`
}
type daxClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dax_cluster.
func (dc daxClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_dax_cluster.
func (dc daxClusterAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dc.ref.Append("availability_zones"))
}

// ClusterAddress returns a reference to field cluster_address of aws_dax_cluster.
func (dc daxClusterAttributes) ClusterAddress() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("cluster_address"))
}

// ClusterEndpointEncryptionType returns a reference to field cluster_endpoint_encryption_type of aws_dax_cluster.
func (dc daxClusterAttributes) ClusterEndpointEncryptionType() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("cluster_endpoint_encryption_type"))
}

// ClusterName returns a reference to field cluster_name of aws_dax_cluster.
func (dc daxClusterAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("cluster_name"))
}

// ConfigurationEndpoint returns a reference to field configuration_endpoint of aws_dax_cluster.
func (dc daxClusterAttributes) ConfigurationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("configuration_endpoint"))
}

// Description returns a reference to field description of aws_dax_cluster.
func (dc daxClusterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("description"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_dax_cluster.
func (dc daxClusterAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_dax_cluster.
func (dc daxClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("id"))
}

// MaintenanceWindow returns a reference to field maintenance_window of aws_dax_cluster.
func (dc daxClusterAttributes) MaintenanceWindow() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("maintenance_window"))
}

// NodeType returns a reference to field node_type of aws_dax_cluster.
func (dc daxClusterAttributes) NodeType() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("node_type"))
}

// NotificationTopicArn returns a reference to field notification_topic_arn of aws_dax_cluster.
func (dc daxClusterAttributes) NotificationTopicArn() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("notification_topic_arn"))
}

// ParameterGroupName returns a reference to field parameter_group_name of aws_dax_cluster.
func (dc daxClusterAttributes) ParameterGroupName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("parameter_group_name"))
}

// Port returns a reference to field port of aws_dax_cluster.
func (dc daxClusterAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("port"))
}

// ReplicationFactor returns a reference to field replication_factor of aws_dax_cluster.
func (dc daxClusterAttributes) ReplicationFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("replication_factor"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_dax_cluster.
func (dc daxClusterAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dc.ref.Append("security_group_ids"))
}

// SubnetGroupName returns a reference to field subnet_group_name of aws_dax_cluster.
func (dc daxClusterAttributes) SubnetGroupName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("subnet_group_name"))
}

// Tags returns a reference to field tags of aws_dax_cluster.
func (dc daxClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dax_cluster.
func (dc daxClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dc.ref.Append("tags_all"))
}

func (dc daxClusterAttributes) Nodes() terra.ListValue[daxcluster.NodesAttributes] {
	return terra.ReferenceAsList[daxcluster.NodesAttributes](dc.ref.Append("nodes"))
}

func (dc daxClusterAttributes) ServerSideEncryption() terra.ListValue[daxcluster.ServerSideEncryptionAttributes] {
	return terra.ReferenceAsList[daxcluster.ServerSideEncryptionAttributes](dc.ref.Append("server_side_encryption"))
}

func (dc daxClusterAttributes) Timeouts() daxcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[daxcluster.TimeoutsAttributes](dc.ref.Append("timeouts"))
}

type daxClusterState struct {
	Arn                           string                                 `json:"arn"`
	AvailabilityZones             []string                               `json:"availability_zones"`
	ClusterAddress                string                                 `json:"cluster_address"`
	ClusterEndpointEncryptionType string                                 `json:"cluster_endpoint_encryption_type"`
	ClusterName                   string                                 `json:"cluster_name"`
	ConfigurationEndpoint         string                                 `json:"configuration_endpoint"`
	Description                   string                                 `json:"description"`
	IamRoleArn                    string                                 `json:"iam_role_arn"`
	Id                            string                                 `json:"id"`
	MaintenanceWindow             string                                 `json:"maintenance_window"`
	NodeType                      string                                 `json:"node_type"`
	NotificationTopicArn          string                                 `json:"notification_topic_arn"`
	ParameterGroupName            string                                 `json:"parameter_group_name"`
	Port                          float64                                `json:"port"`
	ReplicationFactor             float64                                `json:"replication_factor"`
	SecurityGroupIds              []string                               `json:"security_group_ids"`
	SubnetGroupName               string                                 `json:"subnet_group_name"`
	Tags                          map[string]string                      `json:"tags"`
	TagsAll                       map[string]string                      `json:"tags_all"`
	Nodes                         []daxcluster.NodesState                `json:"nodes"`
	ServerSideEncryption          []daxcluster.ServerSideEncryptionState `json:"server_side_encryption"`
	Timeouts                      *daxcluster.TimeoutsState              `json:"timeouts"`
}
