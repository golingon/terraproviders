// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_eks_node_group

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_eks_node_group.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aeng *DataSource) DataSource() string {
	return "aws_eks_node_group"
}

// LocalName returns the local name for [DataSource].
func (aeng *DataSource) LocalName() string {
	return aeng.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aeng *DataSource) Configuration() interface{} {
	return aeng.Args
}

// Attributes returns the attributes for [DataSource].
func (aeng *DataSource) Attributes() dataAwsEksNodeGroupAttributes {
	return dataAwsEksNodeGroupAttributes{ref: terra.ReferenceDataSource(aeng)}
}

// DataArgs contains the configurations for aws_eks_node_group.
type DataArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NodeGroupName: string, required
	NodeGroupName terra.StringValue `hcl:"node_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsEksNodeGroupAttributes struct {
	ref terra.Reference
}

// AmiType returns a reference to field ami_type of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) AmiType() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("ami_type"))
}

// Arn returns a reference to field arn of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("arn"))
}

// CapacityType returns a reference to field capacity_type of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) CapacityType() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("capacity_type"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("cluster_name"))
}

// DiskSize returns a reference to field disk_size of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) DiskSize() terra.NumberValue {
	return terra.ReferenceAsNumber(aeng.ref.Append("disk_size"))
}

// Id returns a reference to field id of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("id"))
}

// InstanceTypes returns a reference to field instance_types of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) InstanceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aeng.ref.Append("instance_types"))
}

// Labels returns a reference to field labels of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeng.ref.Append("labels"))
}

// NodeGroupName returns a reference to field node_group_name of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) NodeGroupName() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("node_group_name"))
}

// NodeRoleArn returns a reference to field node_role_arn of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) NodeRoleArn() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("node_role_arn"))
}

// ReleaseVersion returns a reference to field release_version of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) ReleaseVersion() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("release_version"))
}

// Status returns a reference to field status of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("status"))
}

// SubnetIds returns a reference to field subnet_ids of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aeng.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aeng.ref.Append("tags"))
}

// Version returns a reference to field version of aws_eks_node_group.
func (aeng dataAwsEksNodeGroupAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aeng.ref.Append("version"))
}

func (aeng dataAwsEksNodeGroupAttributes) LaunchTemplate() terra.ListValue[DataLaunchTemplateAttributes] {
	return terra.ReferenceAsList[DataLaunchTemplateAttributes](aeng.ref.Append("launch_template"))
}

func (aeng dataAwsEksNodeGroupAttributes) RemoteAccess() terra.ListValue[DataRemoteAccessAttributes] {
	return terra.ReferenceAsList[DataRemoteAccessAttributes](aeng.ref.Append("remote_access"))
}

func (aeng dataAwsEksNodeGroupAttributes) Resources() terra.ListValue[DataResourcesAttributes] {
	return terra.ReferenceAsList[DataResourcesAttributes](aeng.ref.Append("resources"))
}

func (aeng dataAwsEksNodeGroupAttributes) ScalingConfig() terra.ListValue[DataScalingConfigAttributes] {
	return terra.ReferenceAsList[DataScalingConfigAttributes](aeng.ref.Append("scaling_config"))
}

func (aeng dataAwsEksNodeGroupAttributes) Taints() terra.ListValue[DataTaintsAttributes] {
	return terra.ReferenceAsList[DataTaintsAttributes](aeng.ref.Append("taints"))
}
