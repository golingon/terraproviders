// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasignersigningprofile "github.com/golingon/terraproviders/aws/5.10.0/datasignersigningprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSignerSigningProfile creates a new instance of [DataSignerSigningProfile].
func NewDataSignerSigningProfile(name string, args DataSignerSigningProfileArgs) *DataSignerSigningProfile {
	return &DataSignerSigningProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSignerSigningProfile)(nil)

// DataSignerSigningProfile represents the Terraform data resource aws_signer_signing_profile.
type DataSignerSigningProfile struct {
	Name string
	Args DataSignerSigningProfileArgs
}

// DataSource returns the Terraform object type for [DataSignerSigningProfile].
func (ssp *DataSignerSigningProfile) DataSource() string {
	return "aws_signer_signing_profile"
}

// LocalName returns the local name for [DataSignerSigningProfile].
func (ssp *DataSignerSigningProfile) LocalName() string {
	return ssp.Name
}

// Configuration returns the configuration (args) for [DataSignerSigningProfile].
func (ssp *DataSignerSigningProfile) Configuration() interface{} {
	return ssp.Args
}

// Attributes returns the attributes for [DataSignerSigningProfile].
func (ssp *DataSignerSigningProfile) Attributes() dataSignerSigningProfileAttributes {
	return dataSignerSigningProfileAttributes{ref: terra.ReferenceDataResource(ssp)}
}

// DataSignerSigningProfileArgs contains the configurations for aws_signer_signing_profile.
type DataSignerSigningProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// RevocationRecord: min=0
	RevocationRecord []datasignersigningprofile.RevocationRecord `hcl:"revocation_record,block" validate:"min=0"`
	// SignatureValidityPeriod: min=0
	SignatureValidityPeriod []datasignersigningprofile.SignatureValidityPeriod `hcl:"signature_validity_period,block" validate:"min=0"`
}
type dataSignerSigningProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("id"))
}

// Name returns a reference to field name of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("name"))
}

// PlatformDisplayName returns a reference to field platform_display_name of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) PlatformDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("platform_display_name"))
}

// PlatformId returns a reference to field platform_id of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) PlatformId() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("platform_id"))
}

// Status returns a reference to field status of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssp.ref.Append("tags"))
}

// Version returns a reference to field version of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("version"))
}

// VersionArn returns a reference to field version_arn of aws_signer_signing_profile.
func (ssp dataSignerSigningProfileAttributes) VersionArn() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("version_arn"))
}

func (ssp dataSignerSigningProfileAttributes) RevocationRecord() terra.ListValue[datasignersigningprofile.RevocationRecordAttributes] {
	return terra.ReferenceAsList[datasignersigningprofile.RevocationRecordAttributes](ssp.ref.Append("revocation_record"))
}

func (ssp dataSignerSigningProfileAttributes) SignatureValidityPeriod() terra.ListValue[datasignersigningprofile.SignatureValidityPeriodAttributes] {
	return terra.ReferenceAsList[datasignersigningprofile.SignatureValidityPeriodAttributes](ssp.ref.Append("signature_validity_period"))
}