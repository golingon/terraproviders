// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataeksnodegroup "github.com/golingon/terraproviders/aws/4.63.0/dataeksnodegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEksNodeGroup creates a new instance of [DataEksNodeGroup].
func NewDataEksNodeGroup(name string, args DataEksNodeGroupArgs) *DataEksNodeGroup {
	return &DataEksNodeGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEksNodeGroup)(nil)

// DataEksNodeGroup represents the Terraform data resource aws_eks_node_group.
type DataEksNodeGroup struct {
	Name string
	Args DataEksNodeGroupArgs
}

// DataSource returns the Terraform object type for [DataEksNodeGroup].
func (eng *DataEksNodeGroup) DataSource() string {
	return "aws_eks_node_group"
}

// LocalName returns the local name for [DataEksNodeGroup].
func (eng *DataEksNodeGroup) LocalName() string {
	return eng.Name
}

// Configuration returns the configuration (args) for [DataEksNodeGroup].
func (eng *DataEksNodeGroup) Configuration() interface{} {
	return eng.Args
}

// Attributes returns the attributes for [DataEksNodeGroup].
func (eng *DataEksNodeGroup) Attributes() dataEksNodeGroupAttributes {
	return dataEksNodeGroupAttributes{ref: terra.ReferenceDataResource(eng)}
}

// DataEksNodeGroupArgs contains the configurations for aws_eks_node_group.
type DataEksNodeGroupArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NodeGroupName: string, required
	NodeGroupName terra.StringValue `hcl:"node_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// RemoteAccess: min=0
	RemoteAccess []dataeksnodegroup.RemoteAccess `hcl:"remote_access,block" validate:"min=0"`
	// Resources: min=0
	Resources []dataeksnodegroup.Resources `hcl:"resources,block" validate:"min=0"`
	// ScalingConfig: min=0
	ScalingConfig []dataeksnodegroup.ScalingConfig `hcl:"scaling_config,block" validate:"min=0"`
	// Taints: min=0
	Taints []dataeksnodegroup.Taints `hcl:"taints,block" validate:"min=0"`
}
type dataEksNodeGroupAttributes struct {
	ref terra.Reference
}

// AmiType returns a reference to field ami_type of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) AmiType() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("ami_type"))
}

// Arn returns a reference to field arn of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("arn"))
}

// CapacityType returns a reference to field capacity_type of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) CapacityType() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("capacity_type"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("cluster_name"))
}

// DiskSize returns a reference to field disk_size of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) DiskSize() terra.NumberValue {
	return terra.ReferenceAsNumber(eng.ref.Append("disk_size"))
}

// Id returns a reference to field id of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("id"))
}

// InstanceTypes returns a reference to field instance_types of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) InstanceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eng.ref.Append("instance_types"))
}

// Labels returns a reference to field labels of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eng.ref.Append("labels"))
}

// NodeGroupName returns a reference to field node_group_name of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) NodeGroupName() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("node_group_name"))
}

// NodeRoleArn returns a reference to field node_role_arn of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) NodeRoleArn() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("node_role_arn"))
}

// ReleaseVersion returns a reference to field release_version of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) ReleaseVersion() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("release_version"))
}

// Status returns a reference to field status of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("status"))
}

// SubnetIds returns a reference to field subnet_ids of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](eng.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eng.ref.Append("tags"))
}

// Version returns a reference to field version of aws_eks_node_group.
func (eng dataEksNodeGroupAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(eng.ref.Append("version"))
}

func (eng dataEksNodeGroupAttributes) RemoteAccess() terra.ListValue[dataeksnodegroup.RemoteAccessAttributes] {
	return terra.ReferenceAsList[dataeksnodegroup.RemoteAccessAttributes](eng.ref.Append("remote_access"))
}

func (eng dataEksNodeGroupAttributes) Resources() terra.ListValue[dataeksnodegroup.ResourcesAttributes] {
	return terra.ReferenceAsList[dataeksnodegroup.ResourcesAttributes](eng.ref.Append("resources"))
}

func (eng dataEksNodeGroupAttributes) ScalingConfig() terra.ListValue[dataeksnodegroup.ScalingConfigAttributes] {
	return terra.ReferenceAsList[dataeksnodegroup.ScalingConfigAttributes](eng.ref.Append("scaling_config"))
}

func (eng dataEksNodeGroupAttributes) Taints() terra.ListValue[dataeksnodegroup.TaintsAttributes] {
	return terra.ReferenceAsList[dataeksnodegroup.TaintsAttributes](eng.ref.Append("taints"))
}
