// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataeksnodegroup "github.com/golingon/terraproviders/aws/4.60.0/dataeksnodegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataEksNodeGroup(name string, args DataEksNodeGroupArgs) *DataEksNodeGroup {
	return &DataEksNodeGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEksNodeGroup)(nil)

type DataEksNodeGroup struct {
	Name string
	Args DataEksNodeGroupArgs
}

func (eng *DataEksNodeGroup) DataSource() string {
	return "aws_eks_node_group"
}

func (eng *DataEksNodeGroup) LocalName() string {
	return eng.Name
}

func (eng *DataEksNodeGroup) Configuration() interface{} {
	return eng.Args
}

func (eng *DataEksNodeGroup) Attributes() dataEksNodeGroupAttributes {
	return dataEksNodeGroupAttributes{ref: terra.ReferenceDataResource(eng)}
}

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

func (eng dataEksNodeGroupAttributes) AmiType() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("ami_type"))
}

func (eng dataEksNodeGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("arn"))
}

func (eng dataEksNodeGroupAttributes) CapacityType() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("capacity_type"))
}

func (eng dataEksNodeGroupAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("cluster_name"))
}

func (eng dataEksNodeGroupAttributes) DiskSize() terra.NumberValue {
	return terra.ReferenceNumber(eng.ref.Append("disk_size"))
}

func (eng dataEksNodeGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("id"))
}

func (eng dataEksNodeGroupAttributes) InstanceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](eng.ref.Append("instance_types"))
}

func (eng dataEksNodeGroupAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](eng.ref.Append("labels"))
}

func (eng dataEksNodeGroupAttributes) NodeGroupName() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("node_group_name"))
}

func (eng dataEksNodeGroupAttributes) NodeRoleArn() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("node_role_arn"))
}

func (eng dataEksNodeGroupAttributes) ReleaseVersion() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("release_version"))
}

func (eng dataEksNodeGroupAttributes) Status() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("status"))
}

func (eng dataEksNodeGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](eng.ref.Append("subnet_ids"))
}

func (eng dataEksNodeGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](eng.ref.Append("tags"))
}

func (eng dataEksNodeGroupAttributes) Version() terra.StringValue {
	return terra.ReferenceString(eng.ref.Append("version"))
}

func (eng dataEksNodeGroupAttributes) RemoteAccess() terra.ListValue[dataeksnodegroup.RemoteAccessAttributes] {
	return terra.ReferenceList[dataeksnodegroup.RemoteAccessAttributes](eng.ref.Append("remote_access"))
}

func (eng dataEksNodeGroupAttributes) Resources() terra.ListValue[dataeksnodegroup.ResourcesAttributes] {
	return terra.ReferenceList[dataeksnodegroup.ResourcesAttributes](eng.ref.Append("resources"))
}

func (eng dataEksNodeGroupAttributes) ScalingConfig() terra.ListValue[dataeksnodegroup.ScalingConfigAttributes] {
	return terra.ReferenceList[dataeksnodegroup.ScalingConfigAttributes](eng.ref.Append("scaling_config"))
}

func (eng dataEksNodeGroupAttributes) Taints() terra.ListValue[dataeksnodegroup.TaintsAttributes] {
	return terra.ReferenceList[dataeksnodegroup.TaintsAttributes](eng.ref.Append("taints"))
}
