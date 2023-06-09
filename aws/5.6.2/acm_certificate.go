// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	acmcertificate "github.com/golingon/terraproviders/aws/5.6.2/acmcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAcmCertificate creates a new instance of [AcmCertificate].
func NewAcmCertificate(name string, args AcmCertificateArgs) *AcmCertificate {
	return &AcmCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AcmCertificate)(nil)

// AcmCertificate represents the Terraform resource aws_acm_certificate.
type AcmCertificate struct {
	Name      string
	Args      AcmCertificateArgs
	state     *acmCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AcmCertificate].
func (ac *AcmCertificate) Type() string {
	return "aws_acm_certificate"
}

// LocalName returns the local name for [AcmCertificate].
func (ac *AcmCertificate) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AcmCertificate].
func (ac *AcmCertificate) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AcmCertificate].
func (ac *AcmCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AcmCertificate] depends_on.
func (ac *AcmCertificate) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AcmCertificate].
func (ac *AcmCertificate) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AcmCertificate].
func (ac *AcmCertificate) Attributes() acmCertificateAttributes {
	return acmCertificateAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AcmCertificate]'s state.
func (ac *AcmCertificate) ImportState(av io.Reader) error {
	ac.state = &acmCertificateState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AcmCertificate] has state.
func (ac *AcmCertificate) State() (*acmCertificateState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AcmCertificate]. Panics if the state is nil.
func (ac *AcmCertificate) StateMust() *acmCertificateState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AcmCertificateArgs contains the configurations for aws_acm_certificate.
type AcmCertificateArgs struct {
	// CertificateAuthorityArn: string, optional
	CertificateAuthorityArn terra.StringValue `hcl:"certificate_authority_arn,attr"`
	// CertificateBody: string, optional
	CertificateBody terra.StringValue `hcl:"certificate_body,attr"`
	// CertificateChain: string, optional
	CertificateChain terra.StringValue `hcl:"certificate_chain,attr"`
	// DomainName: string, optional
	DomainName terra.StringValue `hcl:"domain_name,attr"`
	// EarlyRenewalDuration: string, optional
	EarlyRenewalDuration terra.StringValue `hcl:"early_renewal_duration,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyAlgorithm: string, optional
	KeyAlgorithm terra.StringValue `hcl:"key_algorithm,attr"`
	// PrivateKey: string, optional
	PrivateKey terra.StringValue `hcl:"private_key,attr"`
	// SubjectAlternativeNames: set of string, optional
	SubjectAlternativeNames terra.SetValue[terra.StringValue] `hcl:"subject_alternative_names,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ValidationMethod: string, optional
	ValidationMethod terra.StringValue `hcl:"validation_method,attr"`
	// DomainValidationOptions: min=0
	DomainValidationOptions []acmcertificate.DomainValidationOptions `hcl:"domain_validation_options,block" validate:"min=0"`
	// RenewalSummary: min=0
	RenewalSummary []acmcertificate.RenewalSummary `hcl:"renewal_summary,block" validate:"min=0"`
	// Options: optional
	Options *acmcertificate.Options `hcl:"options,block"`
	// ValidationOption: min=0
	ValidationOption []acmcertificate.ValidationOption `hcl:"validation_option,block" validate:"min=0"`
}
type acmCertificateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_acm_certificate.
func (ac acmCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("arn"))
}

// CertificateAuthorityArn returns a reference to field certificate_authority_arn of aws_acm_certificate.
func (ac acmCertificateAttributes) CertificateAuthorityArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_authority_arn"))
}

// CertificateBody returns a reference to field certificate_body of aws_acm_certificate.
func (ac acmCertificateAttributes) CertificateBody() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_body"))
}

// CertificateChain returns a reference to field certificate_chain of aws_acm_certificate.
func (ac acmCertificateAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_chain"))
}

// DomainName returns a reference to field domain_name of aws_acm_certificate.
func (ac acmCertificateAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("domain_name"))
}

// EarlyRenewalDuration returns a reference to field early_renewal_duration of aws_acm_certificate.
func (ac acmCertificateAttributes) EarlyRenewalDuration() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("early_renewal_duration"))
}

// Id returns a reference to field id of aws_acm_certificate.
func (ac acmCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// KeyAlgorithm returns a reference to field key_algorithm of aws_acm_certificate.
func (ac acmCertificateAttributes) KeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("key_algorithm"))
}

// NotAfter returns a reference to field not_after of aws_acm_certificate.
func (ac acmCertificateAttributes) NotAfter() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("not_after"))
}

// NotBefore returns a reference to field not_before of aws_acm_certificate.
func (ac acmCertificateAttributes) NotBefore() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("not_before"))
}

// PendingRenewal returns a reference to field pending_renewal of aws_acm_certificate.
func (ac acmCertificateAttributes) PendingRenewal() terra.BoolValue {
	return terra.ReferenceAsBool(ac.ref.Append("pending_renewal"))
}

// PrivateKey returns a reference to field private_key of aws_acm_certificate.
func (ac acmCertificateAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("private_key"))
}

// RenewalEligibility returns a reference to field renewal_eligibility of aws_acm_certificate.
func (ac acmCertificateAttributes) RenewalEligibility() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("renewal_eligibility"))
}

// Status returns a reference to field status of aws_acm_certificate.
func (ac acmCertificateAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("status"))
}

// SubjectAlternativeNames returns a reference to field subject_alternative_names of aws_acm_certificate.
func (ac acmCertificateAttributes) SubjectAlternativeNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ac.ref.Append("subject_alternative_names"))
}

// Tags returns a reference to field tags of aws_acm_certificate.
func (ac acmCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_acm_certificate.
func (ac acmCertificateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ac.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_acm_certificate.
func (ac acmCertificateAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("type"))
}

// ValidationEmails returns a reference to field validation_emails of aws_acm_certificate.
func (ac acmCertificateAttributes) ValidationEmails() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ac.ref.Append("validation_emails"))
}

// ValidationMethod returns a reference to field validation_method of aws_acm_certificate.
func (ac acmCertificateAttributes) ValidationMethod() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("validation_method"))
}

func (ac acmCertificateAttributes) DomainValidationOptions() terra.SetValue[acmcertificate.DomainValidationOptionsAttributes] {
	return terra.ReferenceAsSet[acmcertificate.DomainValidationOptionsAttributes](ac.ref.Append("domain_validation_options"))
}

func (ac acmCertificateAttributes) RenewalSummary() terra.ListValue[acmcertificate.RenewalSummaryAttributes] {
	return terra.ReferenceAsList[acmcertificate.RenewalSummaryAttributes](ac.ref.Append("renewal_summary"))
}

func (ac acmCertificateAttributes) Options() terra.ListValue[acmcertificate.OptionsAttributes] {
	return terra.ReferenceAsList[acmcertificate.OptionsAttributes](ac.ref.Append("options"))
}

func (ac acmCertificateAttributes) ValidationOption() terra.SetValue[acmcertificate.ValidationOptionAttributes] {
	return terra.ReferenceAsSet[acmcertificate.ValidationOptionAttributes](ac.ref.Append("validation_option"))
}

type acmCertificateState struct {
	Arn                     string                                        `json:"arn"`
	CertificateAuthorityArn string                                        `json:"certificate_authority_arn"`
	CertificateBody         string                                        `json:"certificate_body"`
	CertificateChain        string                                        `json:"certificate_chain"`
	DomainName              string                                        `json:"domain_name"`
	EarlyRenewalDuration    string                                        `json:"early_renewal_duration"`
	Id                      string                                        `json:"id"`
	KeyAlgorithm            string                                        `json:"key_algorithm"`
	NotAfter                string                                        `json:"not_after"`
	NotBefore               string                                        `json:"not_before"`
	PendingRenewal          bool                                          `json:"pending_renewal"`
	PrivateKey              string                                        `json:"private_key"`
	RenewalEligibility      string                                        `json:"renewal_eligibility"`
	Status                  string                                        `json:"status"`
	SubjectAlternativeNames []string                                      `json:"subject_alternative_names"`
	Tags                    map[string]string                             `json:"tags"`
	TagsAll                 map[string]string                             `json:"tags_all"`
	Type                    string                                        `json:"type"`
	ValidationEmails        []string                                      `json:"validation_emails"`
	ValidationMethod        string                                        `json:"validation_method"`
	DomainValidationOptions []acmcertificate.DomainValidationOptionsState `json:"domain_validation_options"`
	RenewalSummary          []acmcertificate.RenewalSummaryState          `json:"renewal_summary"`
	Options                 []acmcertificate.OptionsState                 `json:"options"`
	ValidationOption        []acmcertificate.ValidationOptionState        `json:"validation_option"`
}
