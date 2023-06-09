// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datainternetgateway "github.com/golingon/terraproviders/aws/4.66.1/datainternetgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataInternetGateway creates a new instance of [DataInternetGateway].
func NewDataInternetGateway(name string, args DataInternetGatewayArgs) *DataInternetGateway {
	return &DataInternetGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataInternetGateway)(nil)

// DataInternetGateway represents the Terraform data resource aws_internet_gateway.
type DataInternetGateway struct {
	Name string
	Args DataInternetGatewayArgs
}

// DataSource returns the Terraform object type for [DataInternetGateway].
func (ig *DataInternetGateway) DataSource() string {
	return "aws_internet_gateway"
}

// LocalName returns the local name for [DataInternetGateway].
func (ig *DataInternetGateway) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [DataInternetGateway].
func (ig *DataInternetGateway) Configuration() interface{} {
	return ig.Args
}

// Attributes returns the attributes for [DataInternetGateway].
func (ig *DataInternetGateway) Attributes() dataInternetGatewayAttributes {
	return dataInternetGatewayAttributes{ref: terra.ReferenceDataResource(ig)}
}

// DataInternetGatewayArgs contains the configurations for aws_internet_gateway.
type DataInternetGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternetGatewayId: string, optional
	InternetGatewayId terra.StringValue `hcl:"internet_gateway_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Attachments: min=0
	Attachments []datainternetgateway.Attachments `hcl:"attachments,block" validate:"min=0"`
	// Filter: min=0
	Filter []datainternetgateway.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datainternetgateway.Timeouts `hcl:"timeouts,block"`
}
type dataInternetGatewayAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_internet_gateway.
func (ig dataInternetGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("arn"))
}

// Id returns a reference to field id of aws_internet_gateway.
func (ig dataInternetGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// InternetGatewayId returns a reference to field internet_gateway_id of aws_internet_gateway.
func (ig dataInternetGatewayAttributes) InternetGatewayId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("internet_gateway_id"))
}

// OwnerId returns a reference to field owner_id of aws_internet_gateway.
func (ig dataInternetGatewayAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_internet_gateway.
func (ig dataInternetGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("tags"))
}

func (ig dataInternetGatewayAttributes) Attachments() terra.ListValue[datainternetgateway.AttachmentsAttributes] {
	return terra.ReferenceAsList[datainternetgateway.AttachmentsAttributes](ig.ref.Append("attachments"))
}

func (ig dataInternetGatewayAttributes) Filter() terra.SetValue[datainternetgateway.FilterAttributes] {
	return terra.ReferenceAsSet[datainternetgateway.FilterAttributes](ig.ref.Append("filter"))
}

func (ig dataInternetGatewayAttributes) Timeouts() datainternetgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datainternetgateway.TimeoutsAttributes](ig.ref.Append("timeouts"))
}
