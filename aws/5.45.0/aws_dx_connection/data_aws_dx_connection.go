// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_dx_connection

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_dx_connection.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (adc *DataSource) DataSource() string {
	return "aws_dx_connection"
}

// LocalName returns the local name for [DataSource].
func (adc *DataSource) LocalName() string {
	return adc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (adc *DataSource) Configuration() interface{} {
	return adc.Args
}

// Attributes returns the attributes for [DataSource].
func (adc *DataSource) Attributes() dataAwsDxConnectionAttributes {
	return dataAwsDxConnectionAttributes{ref: terra.ReferenceDataSource(adc)}
}

// DataArgs contains the configurations for aws_dx_connection.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type dataAwsDxConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("aws_device"))
}

// Bandwidth returns a reference to field bandwidth of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) Bandwidth() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("bandwidth"))
}

// Id returns a reference to field id of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("id"))
}

// Location returns a reference to field location of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("location"))
}

// Name returns a reference to field name of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("owner_account_id"))
}

// PartnerName returns a reference to field partner_name of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) PartnerName() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("partner_name"))
}

// ProviderName returns a reference to field provider_name of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("provider_name"))
}

// Tags returns a reference to field tags of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adc.ref.Append("tags"))
}

// VlanId returns a reference to field vlan_id of aws_dx_connection.
func (adc dataAwsDxConnectionAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(adc.ref.Append("vlan_id"))
}
