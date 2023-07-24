// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasesv2emailidentity "github.com/golingon/terraproviders/aws/5.9.0/datasesv2emailidentity"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSesv2EmailIdentity creates a new instance of [DataSesv2EmailIdentity].
func NewDataSesv2EmailIdentity(name string, args DataSesv2EmailIdentityArgs) *DataSesv2EmailIdentity {
	return &DataSesv2EmailIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSesv2EmailIdentity)(nil)

// DataSesv2EmailIdentity represents the Terraform data resource aws_sesv2_email_identity.
type DataSesv2EmailIdentity struct {
	Name string
	Args DataSesv2EmailIdentityArgs
}

// DataSource returns the Terraform object type for [DataSesv2EmailIdentity].
func (sei *DataSesv2EmailIdentity) DataSource() string {
	return "aws_sesv2_email_identity"
}

// LocalName returns the local name for [DataSesv2EmailIdentity].
func (sei *DataSesv2EmailIdentity) LocalName() string {
	return sei.Name
}

// Configuration returns the configuration (args) for [DataSesv2EmailIdentity].
func (sei *DataSesv2EmailIdentity) Configuration() interface{} {
	return sei.Args
}

// Attributes returns the attributes for [DataSesv2EmailIdentity].
func (sei *DataSesv2EmailIdentity) Attributes() dataSesv2EmailIdentityAttributes {
	return dataSesv2EmailIdentityAttributes{ref: terra.ReferenceDataResource(sei)}
}

// DataSesv2EmailIdentityArgs contains the configurations for aws_sesv2_email_identity.
type DataSesv2EmailIdentityArgs struct {
	// EmailIdentity: string, required
	EmailIdentity terra.StringValue `hcl:"email_identity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DkimSigningAttributes: min=0
	DkimSigningAttributes []datasesv2emailidentity.DkimSigningAttributes `hcl:"dkim_signing_attributes,block" validate:"min=0"`
}
type dataSesv2EmailIdentityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("arn"))
}

// ConfigurationSetName returns a reference to field configuration_set_name of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) ConfigurationSetName() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("configuration_set_name"))
}

// EmailIdentity returns a reference to field email_identity of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) EmailIdentity() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("email_identity"))
}

// Id returns a reference to field id of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("id"))
}

// IdentityType returns a reference to field identity_type of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("identity_type"))
}

// Tags returns a reference to field tags of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sei.ref.Append("tags"))
}

// VerifiedForSendingStatus returns a reference to field verified_for_sending_status of aws_sesv2_email_identity.
func (sei dataSesv2EmailIdentityAttributes) VerifiedForSendingStatus() terra.BoolValue {
	return terra.ReferenceAsBool(sei.ref.Append("verified_for_sending_status"))
}

func (sei dataSesv2EmailIdentityAttributes) DkimSigningAttributes() terra.ListValue[datasesv2emailidentity.DkimSigningAttributesAttributes] {
	return terra.ReferenceAsList[datasesv2emailidentity.DkimSigningAttributesAttributes](sei.ref.Append("dkim_signing_attributes"))
}
