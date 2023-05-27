// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datassmcontactscontactchannel "github.com/golingon/terraproviders/aws/5.0.1/datassmcontactscontactchannel"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSsmcontactsContactChannel creates a new instance of [DataSsmcontactsContactChannel].
func NewDataSsmcontactsContactChannel(name string, args DataSsmcontactsContactChannelArgs) *DataSsmcontactsContactChannel {
	return &DataSsmcontactsContactChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmcontactsContactChannel)(nil)

// DataSsmcontactsContactChannel represents the Terraform data resource aws_ssmcontacts_contact_channel.
type DataSsmcontactsContactChannel struct {
	Name string
	Args DataSsmcontactsContactChannelArgs
}

// DataSource returns the Terraform object type for [DataSsmcontactsContactChannel].
func (scc *DataSsmcontactsContactChannel) DataSource() string {
	return "aws_ssmcontacts_contact_channel"
}

// LocalName returns the local name for [DataSsmcontactsContactChannel].
func (scc *DataSsmcontactsContactChannel) LocalName() string {
	return scc.Name
}

// Configuration returns the configuration (args) for [DataSsmcontactsContactChannel].
func (scc *DataSsmcontactsContactChannel) Configuration() interface{} {
	return scc.Args
}

// Attributes returns the attributes for [DataSsmcontactsContactChannel].
func (scc *DataSsmcontactsContactChannel) Attributes() dataSsmcontactsContactChannelAttributes {
	return dataSsmcontactsContactChannelAttributes{ref: terra.ReferenceDataResource(scc)}
}

// DataSsmcontactsContactChannelArgs contains the configurations for aws_ssmcontacts_contact_channel.
type DataSsmcontactsContactChannelArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DeliveryAddress: min=0
	DeliveryAddress []datassmcontactscontactchannel.DeliveryAddress `hcl:"delivery_address,block" validate:"min=0"`
}
type dataSsmcontactsContactChannelAttributes struct {
	ref terra.Reference
}

// ActivationStatus returns a reference to field activation_status of aws_ssmcontacts_contact_channel.
func (scc dataSsmcontactsContactChannelAttributes) ActivationStatus() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("activation_status"))
}

// Arn returns a reference to field arn of aws_ssmcontacts_contact_channel.
func (scc dataSsmcontactsContactChannelAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("arn"))
}

// ContactId returns a reference to field contact_id of aws_ssmcontacts_contact_channel.
func (scc dataSsmcontactsContactChannelAttributes) ContactId() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("contact_id"))
}

// Id returns a reference to field id of aws_ssmcontacts_contact_channel.
func (scc dataSsmcontactsContactChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssmcontacts_contact_channel.
func (scc dataSsmcontactsContactChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("name"))
}

// Type returns a reference to field type of aws_ssmcontacts_contact_channel.
func (scc dataSsmcontactsContactChannelAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("type"))
}

func (scc dataSsmcontactsContactChannelAttributes) DeliveryAddress() terra.ListValue[datassmcontactscontactchannel.DeliveryAddressAttributes] {
	return terra.ReferenceAsList[datassmcontactscontactchannel.DeliveryAddressAttributes](scc.ref.Append("delivery_address"))
}
