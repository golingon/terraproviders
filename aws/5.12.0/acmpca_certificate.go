// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	acmpcacertificate "github.com/golingon/terraproviders/aws/5.12.0/acmpcacertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAcmpcaCertificate creates a new instance of [AcmpcaCertificate].
func NewAcmpcaCertificate(name string, args AcmpcaCertificateArgs) *AcmpcaCertificate {
	return &AcmpcaCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AcmpcaCertificate)(nil)

// AcmpcaCertificate represents the Terraform resource aws_acmpca_certificate.
type AcmpcaCertificate struct {
	Name      string
	Args      AcmpcaCertificateArgs
	state     *acmpcaCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AcmpcaCertificate].
func (ac *AcmpcaCertificate) Type() string {
	return "aws_acmpca_certificate"
}

// LocalName returns the local name for [AcmpcaCertificate].
func (ac *AcmpcaCertificate) LocalName() string {
	return ac.Name
}

// Configuration returns the configuration (args) for [AcmpcaCertificate].
func (ac *AcmpcaCertificate) Configuration() interface{} {
	return ac.Args
}

// DependOn is used for other resources to depend on [AcmpcaCertificate].
func (ac *AcmpcaCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(ac)
}

// Dependencies returns the list of resources [AcmpcaCertificate] depends_on.
func (ac *AcmpcaCertificate) Dependencies() terra.Dependencies {
	return ac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AcmpcaCertificate].
func (ac *AcmpcaCertificate) LifecycleManagement() *terra.Lifecycle {
	return ac.Lifecycle
}

// Attributes returns the attributes for [AcmpcaCertificate].
func (ac *AcmpcaCertificate) Attributes() acmpcaCertificateAttributes {
	return acmpcaCertificateAttributes{ref: terra.ReferenceResource(ac)}
}

// ImportState imports the given attribute values into [AcmpcaCertificate]'s state.
func (ac *AcmpcaCertificate) ImportState(av io.Reader) error {
	ac.state = &acmpcaCertificateState{}
	if err := json.NewDecoder(av).Decode(ac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ac.Type(), ac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AcmpcaCertificate] has state.
func (ac *AcmpcaCertificate) State() (*acmpcaCertificateState, bool) {
	return ac.state, ac.state != nil
}

// StateMust returns the state for [AcmpcaCertificate]. Panics if the state is nil.
func (ac *AcmpcaCertificate) StateMust() *acmpcaCertificateState {
	if ac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ac.Type(), ac.LocalName()))
	}
	return ac.state
}

// AcmpcaCertificateArgs contains the configurations for aws_acmpca_certificate.
type AcmpcaCertificateArgs struct {
	// ApiPassthrough: string, optional
	ApiPassthrough terra.StringValue `hcl:"api_passthrough,attr"`
	// CertificateAuthorityArn: string, required
	CertificateAuthorityArn terra.StringValue `hcl:"certificate_authority_arn,attr" validate:"required"`
	// CertificateSigningRequest: string, required
	CertificateSigningRequest terra.StringValue `hcl:"certificate_signing_request,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SigningAlgorithm: string, required
	SigningAlgorithm terra.StringValue `hcl:"signing_algorithm,attr" validate:"required"`
	// TemplateArn: string, optional
	TemplateArn terra.StringValue `hcl:"template_arn,attr"`
	// Validity: required
	Validity *acmpcacertificate.Validity `hcl:"validity,block" validate:"required"`
}
type acmpcaCertificateAttributes struct {
	ref terra.Reference
}

// ApiPassthrough returns a reference to field api_passthrough of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) ApiPassthrough() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("api_passthrough"))
}

// Arn returns a reference to field arn of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("arn"))
}

// Certificate returns a reference to field certificate of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate"))
}

// CertificateAuthorityArn returns a reference to field certificate_authority_arn of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) CertificateAuthorityArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_authority_arn"))
}

// CertificateChain returns a reference to field certificate_chain of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_chain"))
}

// CertificateSigningRequest returns a reference to field certificate_signing_request of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) CertificateSigningRequest() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("certificate_signing_request"))
}

// Id returns a reference to field id of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("id"))
}

// SigningAlgorithm returns a reference to field signing_algorithm of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) SigningAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("signing_algorithm"))
}

// TemplateArn returns a reference to field template_arn of aws_acmpca_certificate.
func (ac acmpcaCertificateAttributes) TemplateArn() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("template_arn"))
}

func (ac acmpcaCertificateAttributes) Validity() terra.ListValue[acmpcacertificate.ValidityAttributes] {
	return terra.ReferenceAsList[acmpcacertificate.ValidityAttributes](ac.ref.Append("validity"))
}

type acmpcaCertificateState struct {
	ApiPassthrough            string                            `json:"api_passthrough"`
	Arn                       string                            `json:"arn"`
	Certificate               string                            `json:"certificate"`
	CertificateAuthorityArn   string                            `json:"certificate_authority_arn"`
	CertificateChain          string                            `json:"certificate_chain"`
	CertificateSigningRequest string                            `json:"certificate_signing_request"`
	Id                        string                            `json:"id"`
	SigningAlgorithm          string                            `json:"signing_algorithm"`
	TemplateArn               string                            `json:"template_arn"`
	Validity                  []acmpcacertificate.ValidityState `json:"validity"`
}
