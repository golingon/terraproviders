// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	acmcertificatevalidation "github.com/golingon/terraproviders/aws/5.44.0/acmcertificatevalidation"
	"io"
)

// NewAcmCertificateValidation creates a new instance of [AcmCertificateValidation].
func NewAcmCertificateValidation(name string, args AcmCertificateValidationArgs) *AcmCertificateValidation {
	return &AcmCertificateValidation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AcmCertificateValidation)(nil)

// AcmCertificateValidation represents the Terraform resource aws_acm_certificate_validation.
type AcmCertificateValidation struct {
	Name      string
	Args      AcmCertificateValidationArgs
	state     *acmCertificateValidationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AcmCertificateValidation].
func (acv *AcmCertificateValidation) Type() string {
	return "aws_acm_certificate_validation"
}

// LocalName returns the local name for [AcmCertificateValidation].
func (acv *AcmCertificateValidation) LocalName() string {
	return acv.Name
}

// Configuration returns the configuration (args) for [AcmCertificateValidation].
func (acv *AcmCertificateValidation) Configuration() interface{} {
	return acv.Args
}

// DependOn is used for other resources to depend on [AcmCertificateValidation].
func (acv *AcmCertificateValidation) DependOn() terra.Reference {
	return terra.ReferenceResource(acv)
}

// Dependencies returns the list of resources [AcmCertificateValidation] depends_on.
func (acv *AcmCertificateValidation) Dependencies() terra.Dependencies {
	return acv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AcmCertificateValidation].
func (acv *AcmCertificateValidation) LifecycleManagement() *terra.Lifecycle {
	return acv.Lifecycle
}

// Attributes returns the attributes for [AcmCertificateValidation].
func (acv *AcmCertificateValidation) Attributes() acmCertificateValidationAttributes {
	return acmCertificateValidationAttributes{ref: terra.ReferenceResource(acv)}
}

// ImportState imports the given attribute values into [AcmCertificateValidation]'s state.
func (acv *AcmCertificateValidation) ImportState(av io.Reader) error {
	acv.state = &acmCertificateValidationState{}
	if err := json.NewDecoder(av).Decode(acv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acv.Type(), acv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AcmCertificateValidation] has state.
func (acv *AcmCertificateValidation) State() (*acmCertificateValidationState, bool) {
	return acv.state, acv.state != nil
}

// StateMust returns the state for [AcmCertificateValidation]. Panics if the state is nil.
func (acv *AcmCertificateValidation) StateMust() *acmCertificateValidationState {
	if acv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acv.Type(), acv.LocalName()))
	}
	return acv.state
}

// AcmCertificateValidationArgs contains the configurations for aws_acm_certificate_validation.
type AcmCertificateValidationArgs struct {
	// CertificateArn: string, required
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ValidationRecordFqdns: set of string, optional
	ValidationRecordFqdns terra.SetValue[terra.StringValue] `hcl:"validation_record_fqdns,attr"`
	// Timeouts: optional
	Timeouts *acmcertificatevalidation.Timeouts `hcl:"timeouts,block"`
}
type acmCertificateValidationAttributes struct {
	ref terra.Reference
}

// CertificateArn returns a reference to field certificate_arn of aws_acm_certificate_validation.
func (acv acmCertificateValidationAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("certificate_arn"))
}

// Id returns a reference to field id of aws_acm_certificate_validation.
func (acv acmCertificateValidationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acv.ref.Append("id"))
}

// ValidationRecordFqdns returns a reference to field validation_record_fqdns of aws_acm_certificate_validation.
func (acv acmCertificateValidationAttributes) ValidationRecordFqdns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acv.ref.Append("validation_record_fqdns"))
}

func (acv acmCertificateValidationAttributes) Timeouts() acmcertificatevalidation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[acmcertificatevalidation.TimeoutsAttributes](acv.ref.Append("timeouts"))
}

type acmCertificateValidationState struct {
	CertificateArn        string                                  `json:"certificate_arn"`
	Id                    string                                  `json:"id"`
	ValidationRecordFqdns []string                                `json:"validation_record_fqdns"`
	Timeouts              *acmcertificatevalidation.TimeoutsState `json:"timeouts"`
}
