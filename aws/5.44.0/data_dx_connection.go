// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataDxConnection creates a new instance of [DataDxConnection].
func NewDataDxConnection(name string, args DataDxConnectionArgs) *DataDxConnection {
	return &DataDxConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDxConnection)(nil)

// DataDxConnection represents the Terraform data resource aws_dx_connection.
type DataDxConnection struct {
	Name string
	Args DataDxConnectionArgs
}

// DataSource returns the Terraform object type for [DataDxConnection].
func (dc *DataDxConnection) DataSource() string {
	return "aws_dx_connection"
}

// LocalName returns the local name for [DataDxConnection].
func (dc *DataDxConnection) LocalName() string {
	return dc.Name
}

// Configuration returns the configuration (args) for [DataDxConnection].
func (dc *DataDxConnection) Configuration() interface{} {
	return dc.Args
}

// Attributes returns the attributes for [DataDxConnection].
func (dc *DataDxConnection) Attributes() dataDxConnectionAttributes {
	return dataDxConnectionAttributes{ref: terra.ReferenceDataResource(dc)}
}

// DataDxConnectionArgs contains the configurations for aws_dx_connection.
type DataDxConnectionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}
type dataDxConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dx_connection.
func (dc dataDxConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("arn"))
}

// AwsDevice returns a reference to field aws_device of aws_dx_connection.
func (dc dataDxConnectionAttributes) AwsDevice() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("aws_device"))
}

// Bandwidth returns a reference to field bandwidth of aws_dx_connection.
func (dc dataDxConnectionAttributes) Bandwidth() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("bandwidth"))
}

// Id returns a reference to field id of aws_dx_connection.
func (dc dataDxConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("id"))
}

// Location returns a reference to field location of aws_dx_connection.
func (dc dataDxConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("location"))
}

// Name returns a reference to field name of aws_dx_connection.
func (dc dataDxConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("name"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_dx_connection.
func (dc dataDxConnectionAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("owner_account_id"))
}

// PartnerName returns a reference to field partner_name of aws_dx_connection.
func (dc dataDxConnectionAttributes) PartnerName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("partner_name"))
}

// ProviderName returns a reference to field provider_name of aws_dx_connection.
func (dc dataDxConnectionAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("provider_name"))
}

// Tags returns a reference to field tags of aws_dx_connection.
func (dc dataDxConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dc.ref.Append("tags"))
}

// VlanId returns a reference to field vlan_id of aws_dx_connection.
func (dc dataDxConnectionAttributes) VlanId() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("vlan_id"))
}
