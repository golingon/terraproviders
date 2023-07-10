// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datacustomergateway "github.com/golingon/terraproviders/aws/5.7.0/datacustomergateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCustomerGateway creates a new instance of [DataCustomerGateway].
func NewDataCustomerGateway(name string, args DataCustomerGatewayArgs) *DataCustomerGateway {
	return &DataCustomerGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCustomerGateway)(nil)

// DataCustomerGateway represents the Terraform data resource aws_customer_gateway.
type DataCustomerGateway struct {
	Name string
	Args DataCustomerGatewayArgs
}

// DataSource returns the Terraform object type for [DataCustomerGateway].
func (cg *DataCustomerGateway) DataSource() string {
	return "aws_customer_gateway"
}

// LocalName returns the local name for [DataCustomerGateway].
func (cg *DataCustomerGateway) LocalName() string {
	return cg.Name
}

// Configuration returns the configuration (args) for [DataCustomerGateway].
func (cg *DataCustomerGateway) Configuration() interface{} {
	return cg.Args
}

// Attributes returns the attributes for [DataCustomerGateway].
func (cg *DataCustomerGateway) Attributes() dataCustomerGatewayAttributes {
	return dataCustomerGatewayAttributes{ref: terra.ReferenceDataResource(cg)}
}

// DataCustomerGatewayArgs contains the configurations for aws_customer_gateway.
type DataCustomerGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datacustomergateway.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacustomergateway.Timeouts `hcl:"timeouts,block"`
}
type dataCustomerGatewayAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("arn"))
}

// BgpAsn returns a reference to field bgp_asn of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) BgpAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(cg.ref.Append("bgp_asn"))
}

// CertificateArn returns a reference to field certificate_arn of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("certificate_arn"))
}

// DeviceName returns a reference to field device_name of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("device_name"))
}

// Id returns a reference to field id of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("ip_address"))
}

// Tags returns a reference to field tags of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cg.ref.Append("tags"))
}

// Type returns a reference to field type of aws_customer_gateway.
func (cg dataCustomerGatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("type"))
}

func (cg dataCustomerGatewayAttributes) Filter() terra.SetValue[datacustomergateway.FilterAttributes] {
	return terra.ReferenceAsSet[datacustomergateway.FilterAttributes](cg.ref.Append("filter"))
}

func (cg dataCustomerGatewayAttributes) Timeouts() datacustomergateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacustomergateway.TimeoutsAttributes](cg.ref.Append("timeouts"))
}
