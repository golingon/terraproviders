// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataeip "github.com/golingon/terraproviders/aws/5.10.0/dataeip"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEip creates a new instance of [DataEip].
func NewDataEip(name string, args DataEipArgs) *DataEip {
	return &DataEip{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEip)(nil)

// DataEip represents the Terraform data resource aws_eip.
type DataEip struct {
	Name string
	Args DataEipArgs
}

// DataSource returns the Terraform object type for [DataEip].
func (e *DataEip) DataSource() string {
	return "aws_eip"
}

// LocalName returns the local name for [DataEip].
func (e *DataEip) LocalName() string {
	return e.Name
}

// Configuration returns the configuration (args) for [DataEip].
func (e *DataEip) Configuration() interface{} {
	return e.Args
}

// Attributes returns the attributes for [DataEip].
func (e *DataEip) Attributes() dataEipAttributes {
	return dataEipAttributes{ref: terra.ReferenceDataResource(e)}
}

// DataEipArgs contains the configurations for aws_eip.
type DataEipArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PublicIp: string, optional
	PublicIp terra.StringValue `hcl:"public_ip,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataeip.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataeip.Timeouts `hcl:"timeouts,block"`
}
type dataEipAttributes struct {
	ref terra.Reference
}

// AssociationId returns a reference to field association_id of aws_eip.
func (e dataEipAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("association_id"))
}

// CarrierIp returns a reference to field carrier_ip of aws_eip.
func (e dataEipAttributes) CarrierIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("carrier_ip"))
}

// CustomerOwnedIp returns a reference to field customer_owned_ip of aws_eip.
func (e dataEipAttributes) CustomerOwnedIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("customer_owned_ip"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_eip.
func (e dataEipAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("customer_owned_ipv4_pool"))
}

// Domain returns a reference to field domain of aws_eip.
func (e dataEipAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("domain"))
}

// Id returns a reference to field id of aws_eip.
func (e dataEipAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_eip.
func (e dataEipAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("instance_id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_eip.
func (e dataEipAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("network_interface_id"))
}

// NetworkInterfaceOwnerId returns a reference to field network_interface_owner_id of aws_eip.
func (e dataEipAttributes) NetworkInterfaceOwnerId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("network_interface_owner_id"))
}

// PrivateDns returns a reference to field private_dns of aws_eip.
func (e dataEipAttributes) PrivateDns() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("private_dns"))
}

// PrivateIp returns a reference to field private_ip of aws_eip.
func (e dataEipAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("private_ip"))
}

// PublicDns returns a reference to field public_dns of aws_eip.
func (e dataEipAttributes) PublicDns() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("public_dns"))
}

// PublicIp returns a reference to field public_ip of aws_eip.
func (e dataEipAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("public_ip"))
}

// PublicIpv4Pool returns a reference to field public_ipv4_pool of aws_eip.
func (e dataEipAttributes) PublicIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("public_ipv4_pool"))
}

// Tags returns a reference to field tags of aws_eip.
func (e dataEipAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("tags"))
}

func (e dataEipAttributes) Filter() terra.SetValue[dataeip.FilterAttributes] {
	return terra.ReferenceAsSet[dataeip.FilterAttributes](e.ref.Append("filter"))
}

func (e dataEipAttributes) Timeouts() dataeip.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeip.TimeoutsAttributes](e.ref.Append("timeouts"))
}
