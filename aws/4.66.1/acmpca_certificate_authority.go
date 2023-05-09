// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	acmpcacertificateauthority "github.com/golingon/terraproviders/aws/4.66.1/acmpcacertificateauthority"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAcmpcaCertificateAuthority creates a new instance of [AcmpcaCertificateAuthority].
func NewAcmpcaCertificateAuthority(name string, args AcmpcaCertificateAuthorityArgs) *AcmpcaCertificateAuthority {
	return &AcmpcaCertificateAuthority{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AcmpcaCertificateAuthority)(nil)

// AcmpcaCertificateAuthority represents the Terraform resource aws_acmpca_certificate_authority.
type AcmpcaCertificateAuthority struct {
	Name      string
	Args      AcmpcaCertificateAuthorityArgs
	state     *acmpcaCertificateAuthorityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AcmpcaCertificateAuthority].
func (aca *AcmpcaCertificateAuthority) Type() string {
	return "aws_acmpca_certificate_authority"
}

// LocalName returns the local name for [AcmpcaCertificateAuthority].
func (aca *AcmpcaCertificateAuthority) LocalName() string {
	return aca.Name
}

// Configuration returns the configuration (args) for [AcmpcaCertificateAuthority].
func (aca *AcmpcaCertificateAuthority) Configuration() interface{} {
	return aca.Args
}

// DependOn is used for other resources to depend on [AcmpcaCertificateAuthority].
func (aca *AcmpcaCertificateAuthority) DependOn() terra.Reference {
	return terra.ReferenceResource(aca)
}

// Dependencies returns the list of resources [AcmpcaCertificateAuthority] depends_on.
func (aca *AcmpcaCertificateAuthority) Dependencies() terra.Dependencies {
	return aca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AcmpcaCertificateAuthority].
func (aca *AcmpcaCertificateAuthority) LifecycleManagement() *terra.Lifecycle {
	return aca.Lifecycle
}

// Attributes returns the attributes for [AcmpcaCertificateAuthority].
func (aca *AcmpcaCertificateAuthority) Attributes() acmpcaCertificateAuthorityAttributes {
	return acmpcaCertificateAuthorityAttributes{ref: terra.ReferenceResource(aca)}
}

// ImportState imports the given attribute values into [AcmpcaCertificateAuthority]'s state.
func (aca *AcmpcaCertificateAuthority) ImportState(av io.Reader) error {
	aca.state = &acmpcaCertificateAuthorityState{}
	if err := json.NewDecoder(av).Decode(aca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aca.Type(), aca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AcmpcaCertificateAuthority] has state.
func (aca *AcmpcaCertificateAuthority) State() (*acmpcaCertificateAuthorityState, bool) {
	return aca.state, aca.state != nil
}

// StateMust returns the state for [AcmpcaCertificateAuthority]. Panics if the state is nil.
func (aca *AcmpcaCertificateAuthority) StateMust() *acmpcaCertificateAuthorityState {
	if aca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aca.Type(), aca.LocalName()))
	}
	return aca.state
}

// AcmpcaCertificateAuthorityArgs contains the configurations for aws_acmpca_certificate_authority.
type AcmpcaCertificateAuthorityArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PermanentDeletionTimeInDays: number, optional
	PermanentDeletionTimeInDays terra.NumberValue `hcl:"permanent_deletion_time_in_days,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// UsageMode: string, optional
	UsageMode terra.StringValue `hcl:"usage_mode,attr"`
	// CertificateAuthorityConfiguration: required
	CertificateAuthorityConfiguration *acmpcacertificateauthority.CertificateAuthorityConfiguration `hcl:"certificate_authority_configuration,block" validate:"required"`
	// RevocationConfiguration: optional
	RevocationConfiguration *acmpcacertificateauthority.RevocationConfiguration `hcl:"revocation_configuration,block"`
	// Timeouts: optional
	Timeouts *acmpcacertificateauthority.Timeouts `hcl:"timeouts,block"`
}
type acmpcaCertificateAuthorityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("arn"))
}

// Certificate returns a reference to field certificate of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("certificate"))
}

// CertificateChain returns a reference to field certificate_chain of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("certificate_chain"))
}

// CertificateSigningRequest returns a reference to field certificate_signing_request of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) CertificateSigningRequest() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("certificate_signing_request"))
}

// Enabled returns a reference to field enabled of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(aca.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("id"))
}

// NotAfter returns a reference to field not_after of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) NotAfter() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("not_after"))
}

// NotBefore returns a reference to field not_before of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("not_before"))
}

// PermanentDeletionTimeInDays returns a reference to field permanent_deletion_time_in_days of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) PermanentDeletionTimeInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(aca.ref.Append("permanent_deletion_time_in_days"))
}

// Serial returns a reference to field serial of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Serial() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("serial"))
}

// Status returns a reference to field status of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aca.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aca.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("type"))
}

// UsageMode returns a reference to field usage_mode of aws_acmpca_certificate_authority.
func (aca acmpcaCertificateAuthorityAttributes) UsageMode() terra.StringValue {
	return terra.ReferenceAsString(aca.ref.Append("usage_mode"))
}

func (aca acmpcaCertificateAuthorityAttributes) CertificateAuthorityConfiguration() terra.ListValue[acmpcacertificateauthority.CertificateAuthorityConfigurationAttributes] {
	return terra.ReferenceAsList[acmpcacertificateauthority.CertificateAuthorityConfigurationAttributes](aca.ref.Append("certificate_authority_configuration"))
}

func (aca acmpcaCertificateAuthorityAttributes) RevocationConfiguration() terra.ListValue[acmpcacertificateauthority.RevocationConfigurationAttributes] {
	return terra.ReferenceAsList[acmpcacertificateauthority.RevocationConfigurationAttributes](aca.ref.Append("revocation_configuration"))
}

func (aca acmpcaCertificateAuthorityAttributes) Timeouts() acmpcacertificateauthority.TimeoutsAttributes {
	return terra.ReferenceAsSingle[acmpcacertificateauthority.TimeoutsAttributes](aca.ref.Append("timeouts"))
}

type acmpcaCertificateAuthorityState struct {
	Arn                               string                                                              `json:"arn"`
	Certificate                       string                                                              `json:"certificate"`
	CertificateChain                  string                                                              `json:"certificate_chain"`
	CertificateSigningRequest         string                                                              `json:"certificate_signing_request"`
	Enabled                           bool                                                                `json:"enabled"`
	Id                                string                                                              `json:"id"`
	NotAfter                          string                                                              `json:"not_after"`
	NotBefore                         string                                                              `json:"not_before"`
	PermanentDeletionTimeInDays       float64                                                             `json:"permanent_deletion_time_in_days"`
	Serial                            string                                                              `json:"serial"`
	Status                            string                                                              `json:"status"`
	Tags                              map[string]string                                                   `json:"tags"`
	TagsAll                           map[string]string                                                   `json:"tags_all"`
	Type                              string                                                              `json:"type"`
	UsageMode                         string                                                              `json:"usage_mode"`
	CertificateAuthorityConfiguration []acmpcacertificateauthority.CertificateAuthorityConfigurationState `json:"certificate_authority_configuration"`
	RevocationConfiguration           []acmpcacertificateauthority.RevocationConfigurationState           `json:"revocation_configuration"`
	Timeouts                          *acmpcacertificateauthority.TimeoutsState                           `json:"timeouts"`
}
