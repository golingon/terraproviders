// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	signersigningprofile "github.com/golingon/terraproviders/aws/5.6.2/signersigningprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSignerSigningProfile creates a new instance of [SignerSigningProfile].
func NewSignerSigningProfile(name string, args SignerSigningProfileArgs) *SignerSigningProfile {
	return &SignerSigningProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SignerSigningProfile)(nil)

// SignerSigningProfile represents the Terraform resource aws_signer_signing_profile.
type SignerSigningProfile struct {
	Name      string
	Args      SignerSigningProfileArgs
	state     *signerSigningProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SignerSigningProfile].
func (ssp *SignerSigningProfile) Type() string {
	return "aws_signer_signing_profile"
}

// LocalName returns the local name for [SignerSigningProfile].
func (ssp *SignerSigningProfile) LocalName() string {
	return ssp.Name
}

// Configuration returns the configuration (args) for [SignerSigningProfile].
func (ssp *SignerSigningProfile) Configuration() interface{} {
	return ssp.Args
}

// DependOn is used for other resources to depend on [SignerSigningProfile].
func (ssp *SignerSigningProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(ssp)
}

// Dependencies returns the list of resources [SignerSigningProfile] depends_on.
func (ssp *SignerSigningProfile) Dependencies() terra.Dependencies {
	return ssp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SignerSigningProfile].
func (ssp *SignerSigningProfile) LifecycleManagement() *terra.Lifecycle {
	return ssp.Lifecycle
}

// Attributes returns the attributes for [SignerSigningProfile].
func (ssp *SignerSigningProfile) Attributes() signerSigningProfileAttributes {
	return signerSigningProfileAttributes{ref: terra.ReferenceResource(ssp)}
}

// ImportState imports the given attribute values into [SignerSigningProfile]'s state.
func (ssp *SignerSigningProfile) ImportState(av io.Reader) error {
	ssp.state = &signerSigningProfileState{}
	if err := json.NewDecoder(av).Decode(ssp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssp.Type(), ssp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SignerSigningProfile] has state.
func (ssp *SignerSigningProfile) State() (*signerSigningProfileState, bool) {
	return ssp.state, ssp.state != nil
}

// StateMust returns the state for [SignerSigningProfile]. Panics if the state is nil.
func (ssp *SignerSigningProfile) StateMust() *signerSigningProfileState {
	if ssp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssp.Type(), ssp.LocalName()))
	}
	return ssp.state
}

// SignerSigningProfileArgs contains the configurations for aws_signer_signing_profile.
type SignerSigningProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// PlatformId: string, required
	PlatformId terra.StringValue `hcl:"platform_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// RevocationRecord: min=0
	RevocationRecord []signersigningprofile.RevocationRecord `hcl:"revocation_record,block" validate:"min=0"`
	// SignatureValidityPeriod: optional
	SignatureValidityPeriod *signersigningprofile.SignatureValidityPeriod `hcl:"signature_validity_period,block"`
}
type signerSigningProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("arn"))
}

// Id returns a reference to field id of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("id"))
}

// Name returns a reference to field name of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("name_prefix"))
}

// PlatformDisplayName returns a reference to field platform_display_name of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) PlatformDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("platform_display_name"))
}

// PlatformId returns a reference to field platform_id of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) PlatformId() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("platform_id"))
}

// Status returns a reference to field status of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssp.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("version"))
}

// VersionArn returns a reference to field version_arn of aws_signer_signing_profile.
func (ssp signerSigningProfileAttributes) VersionArn() terra.StringValue {
	return terra.ReferenceAsString(ssp.ref.Append("version_arn"))
}

func (ssp signerSigningProfileAttributes) RevocationRecord() terra.ListValue[signersigningprofile.RevocationRecordAttributes] {
	return terra.ReferenceAsList[signersigningprofile.RevocationRecordAttributes](ssp.ref.Append("revocation_record"))
}

func (ssp signerSigningProfileAttributes) SignatureValidityPeriod() terra.ListValue[signersigningprofile.SignatureValidityPeriodAttributes] {
	return terra.ReferenceAsList[signersigningprofile.SignatureValidityPeriodAttributes](ssp.ref.Append("signature_validity_period"))
}

type signerSigningProfileState struct {
	Arn                     string                                              `json:"arn"`
	Id                      string                                              `json:"id"`
	Name                    string                                              `json:"name"`
	NamePrefix              string                                              `json:"name_prefix"`
	PlatformDisplayName     string                                              `json:"platform_display_name"`
	PlatformId              string                                              `json:"platform_id"`
	Status                  string                                              `json:"status"`
	Tags                    map[string]string                                   `json:"tags"`
	TagsAll                 map[string]string                                   `json:"tags_all"`
	Version                 string                                              `json:"version"`
	VersionArn              string                                              `json:"version_arn"`
	RevocationRecord        []signersigningprofile.RevocationRecordState        `json:"revocation_record"`
	SignatureValidityPeriod []signersigningprofile.SignatureValidityPeriodState `json:"signature_validity_period"`
}
