// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataredshiftserverlessworkgroup "github.com/golingon/terraproviders/aws/5.8.0/dataredshiftserverlessworkgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRedshiftserverlessWorkgroup creates a new instance of [DataRedshiftserverlessWorkgroup].
func NewDataRedshiftserverlessWorkgroup(name string, args DataRedshiftserverlessWorkgroupArgs) *DataRedshiftserverlessWorkgroup {
	return &DataRedshiftserverlessWorkgroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRedshiftserverlessWorkgroup)(nil)

// DataRedshiftserverlessWorkgroup represents the Terraform data resource aws_redshiftserverless_workgroup.
type DataRedshiftserverlessWorkgroup struct {
	Name string
	Args DataRedshiftserverlessWorkgroupArgs
}

// DataSource returns the Terraform object type for [DataRedshiftserverlessWorkgroup].
func (rw *DataRedshiftserverlessWorkgroup) DataSource() string {
	return "aws_redshiftserverless_workgroup"
}

// LocalName returns the local name for [DataRedshiftserverlessWorkgroup].
func (rw *DataRedshiftserverlessWorkgroup) LocalName() string {
	return rw.Name
}

// Configuration returns the configuration (args) for [DataRedshiftserverlessWorkgroup].
func (rw *DataRedshiftserverlessWorkgroup) Configuration() interface{} {
	return rw.Args
}

// Attributes returns the attributes for [DataRedshiftserverlessWorkgroup].
func (rw *DataRedshiftserverlessWorkgroup) Attributes() dataRedshiftserverlessWorkgroupAttributes {
	return dataRedshiftserverlessWorkgroupAttributes{ref: terra.ReferenceDataResource(rw)}
}

// DataRedshiftserverlessWorkgroupArgs contains the configurations for aws_redshiftserverless_workgroup.
type DataRedshiftserverlessWorkgroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// WorkgroupName: string, required
	WorkgroupName terra.StringValue `hcl:"workgroup_name,attr" validate:"required"`
	// Endpoint: min=0
	Endpoint []dataredshiftserverlessworkgroup.Endpoint `hcl:"endpoint,block" validate:"min=0"`
}
type dataRedshiftserverlessWorkgroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("arn"))
}

// EnhancedVpcRouting returns a reference to field enhanced_vpc_routing of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) EnhancedVpcRouting() terra.BoolValue {
	return terra.ReferenceAsBool(rw.ref.Append("enhanced_vpc_routing"))
}

// Id returns a reference to field id of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("id"))
}

// NamespaceName returns a reference to field namespace_name of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("namespace_name"))
}

// PubliclyAccessible returns a reference to field publicly_accessible of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) PubliclyAccessible() terra.BoolValue {
	return terra.ReferenceAsBool(rw.ref.Append("publicly_accessible"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rw.ref.Append("security_group_ids"))
}

// SubnetIds returns a reference to field subnet_ids of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rw.ref.Append("subnet_ids"))
}

// WorkgroupId returns a reference to field workgroup_id of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) WorkgroupId() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("workgroup_id"))
}

// WorkgroupName returns a reference to field workgroup_name of aws_redshiftserverless_workgroup.
func (rw dataRedshiftserverlessWorkgroupAttributes) WorkgroupName() terra.StringValue {
	return terra.ReferenceAsString(rw.ref.Append("workgroup_name"))
}

func (rw dataRedshiftserverlessWorkgroupAttributes) Endpoint() terra.ListValue[dataredshiftserverlessworkgroup.EndpointAttributes] {
	return terra.ReferenceAsList[dataredshiftserverlessworkgroup.EndpointAttributes](rw.ref.Append("endpoint"))
}
