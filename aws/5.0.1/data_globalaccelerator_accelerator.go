// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataglobalacceleratoraccelerator "github.com/golingon/terraproviders/aws/5.0.1/dataglobalacceleratoraccelerator"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataGlobalacceleratorAccelerator creates a new instance of [DataGlobalacceleratorAccelerator].
func NewDataGlobalacceleratorAccelerator(name string, args DataGlobalacceleratorAcceleratorArgs) *DataGlobalacceleratorAccelerator {
	return &DataGlobalacceleratorAccelerator{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGlobalacceleratorAccelerator)(nil)

// DataGlobalacceleratorAccelerator represents the Terraform data resource aws_globalaccelerator_accelerator.
type DataGlobalacceleratorAccelerator struct {
	Name string
	Args DataGlobalacceleratorAcceleratorArgs
}

// DataSource returns the Terraform object type for [DataGlobalacceleratorAccelerator].
func (ga *DataGlobalacceleratorAccelerator) DataSource() string {
	return "aws_globalaccelerator_accelerator"
}

// LocalName returns the local name for [DataGlobalacceleratorAccelerator].
func (ga *DataGlobalacceleratorAccelerator) LocalName() string {
	return ga.Name
}

// Configuration returns the configuration (args) for [DataGlobalacceleratorAccelerator].
func (ga *DataGlobalacceleratorAccelerator) Configuration() interface{} {
	return ga.Args
}

// Attributes returns the attributes for [DataGlobalacceleratorAccelerator].
func (ga *DataGlobalacceleratorAccelerator) Attributes() dataGlobalacceleratorAcceleratorAttributes {
	return dataGlobalacceleratorAcceleratorAttributes{ref: terra.ReferenceDataResource(ga)}
}

// DataGlobalacceleratorAcceleratorArgs contains the configurations for aws_globalaccelerator_accelerator.
type DataGlobalacceleratorAcceleratorArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Attributes: min=0
	Attributes []dataglobalacceleratoraccelerator.Attributes `hcl:"attributes,block" validate:"min=0"`
	// IpSets: min=0
	IpSets []dataglobalacceleratoraccelerator.IpSets `hcl:"ip_sets,block" validate:"min=0"`
}
type dataGlobalacceleratorAcceleratorAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("arn"))
}

// DnsName returns a reference to field dns_name of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("dns_name"))
}

// DualStackDnsName returns a reference to field dual_stack_dns_name of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) DualStackDnsName() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("dual_stack_dns_name"))
}

// Enabled returns a reference to field enabled of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ga.ref.Append("enabled"))
}

// HostedZoneId returns a reference to field hosted_zone_id of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) HostedZoneId() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("hosted_zone_id"))
}

// Id returns a reference to field id of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("ip_address_type"))
}

// Name returns a reference to field name of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ga.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_globalaccelerator_accelerator.
func (ga dataGlobalacceleratorAcceleratorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ga.ref.Append("tags"))
}

func (ga dataGlobalacceleratorAcceleratorAttributes) Attributes() terra.ListValue[dataglobalacceleratoraccelerator.AttributesAttributes] {
	return terra.ReferenceAsList[dataglobalacceleratoraccelerator.AttributesAttributes](ga.ref.Append("attributes"))
}

func (ga dataGlobalacceleratorAcceleratorAttributes) IpSets() terra.ListValue[dataglobalacceleratoraccelerator.IpSetsAttributes] {
	return terra.ReferenceAsList[dataglobalacceleratoraccelerator.IpSetsAttributes](ga.ref.Append("ip_sets"))
}
