// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataLbHostedZoneId creates a new instance of [DataLbHostedZoneId].
func NewDataLbHostedZoneId(name string, args DataLbHostedZoneIdArgs) *DataLbHostedZoneId {
	return &DataLbHostedZoneId{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbHostedZoneId)(nil)

// DataLbHostedZoneId represents the Terraform data resource aws_lb_hosted_zone_id.
type DataLbHostedZoneId struct {
	Name string
	Args DataLbHostedZoneIdArgs
}

// DataSource returns the Terraform object type for [DataLbHostedZoneId].
func (lhzi *DataLbHostedZoneId) DataSource() string {
	return "aws_lb_hosted_zone_id"
}

// LocalName returns the local name for [DataLbHostedZoneId].
func (lhzi *DataLbHostedZoneId) LocalName() string {
	return lhzi.Name
}

// Configuration returns the configuration (args) for [DataLbHostedZoneId].
func (lhzi *DataLbHostedZoneId) Configuration() interface{} {
	return lhzi.Args
}

// Attributes returns the attributes for [DataLbHostedZoneId].
func (lhzi *DataLbHostedZoneId) Attributes() dataLbHostedZoneIdAttributes {
	return dataLbHostedZoneIdAttributes{ref: terra.ReferenceDataResource(lhzi)}
}

// DataLbHostedZoneIdArgs contains the configurations for aws_lb_hosted_zone_id.
type DataLbHostedZoneIdArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadBalancerType: string, optional
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataLbHostedZoneIdAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_lb_hosted_zone_id.
func (lhzi dataLbHostedZoneIdAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lhzi.ref.Append("id"))
}

// LoadBalancerType returns a reference to field load_balancer_type of aws_lb_hosted_zone_id.
func (lhzi dataLbHostedZoneIdAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(lhzi.ref.Append("load_balancer_type"))
}

// Region returns a reference to field region of aws_lb_hosted_zone_id.
func (lhzi dataLbHostedZoneIdAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(lhzi.ref.Append("region"))
}