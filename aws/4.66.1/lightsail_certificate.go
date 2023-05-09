// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lightsailcertificate "github.com/golingon/terraproviders/aws/4.66.1/lightsailcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailCertificate creates a new instance of [LightsailCertificate].
func NewLightsailCertificate(name string, args LightsailCertificateArgs) *LightsailCertificate {
	return &LightsailCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailCertificate)(nil)

// LightsailCertificate represents the Terraform resource aws_lightsail_certificate.
type LightsailCertificate struct {
	Name      string
	Args      LightsailCertificateArgs
	state     *lightsailCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailCertificate].
func (lc *LightsailCertificate) Type() string {
	return "aws_lightsail_certificate"
}

// LocalName returns the local name for [LightsailCertificate].
func (lc *LightsailCertificate) LocalName() string {
	return lc.Name
}

// Configuration returns the configuration (args) for [LightsailCertificate].
func (lc *LightsailCertificate) Configuration() interface{} {
	return lc.Args
}

// DependOn is used for other resources to depend on [LightsailCertificate].
func (lc *LightsailCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(lc)
}

// Dependencies returns the list of resources [LightsailCertificate] depends_on.
func (lc *LightsailCertificate) Dependencies() terra.Dependencies {
	return lc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailCertificate].
func (lc *LightsailCertificate) LifecycleManagement() *terra.Lifecycle {
	return lc.Lifecycle
}

// Attributes returns the attributes for [LightsailCertificate].
func (lc *LightsailCertificate) Attributes() lightsailCertificateAttributes {
	return lightsailCertificateAttributes{ref: terra.ReferenceResource(lc)}
}

// ImportState imports the given attribute values into [LightsailCertificate]'s state.
func (lc *LightsailCertificate) ImportState(av io.Reader) error {
	lc.state = &lightsailCertificateState{}
	if err := json.NewDecoder(av).Decode(lc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lc.Type(), lc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailCertificate] has state.
func (lc *LightsailCertificate) State() (*lightsailCertificateState, bool) {
	return lc.state, lc.state != nil
}

// StateMust returns the state for [LightsailCertificate]. Panics if the state is nil.
func (lc *LightsailCertificate) StateMust() *lightsailCertificateState {
	if lc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lc.Type(), lc.LocalName()))
	}
	return lc.state
}

// LightsailCertificateArgs contains the configurations for aws_lightsail_certificate.
type LightsailCertificateArgs struct {
	// DomainName: string, optional
	DomainName terra.StringValue `hcl:"domain_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubjectAlternativeNames: set of string, optional
	SubjectAlternativeNames terra.SetValue[terra.StringValue] `hcl:"subject_alternative_names,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DomainValidationOptions: min=0
	DomainValidationOptions []lightsailcertificate.DomainValidationOptions `hcl:"domain_validation_options,block" validate:"min=0"`
}
type lightsailCertificateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("created_at"))
}

// DomainName returns a reference to field domain_name of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("id"))
}

// Name returns a reference to field name of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("name"))
}

// SubjectAlternativeNames returns a reference to field subject_alternative_names of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) SubjectAlternativeNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lc.ref.Append("subject_alternative_names"))
}

// Tags returns a reference to field tags of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lightsail_certificate.
func (lc lightsailCertificateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("tags_all"))
}

func (lc lightsailCertificateAttributes) DomainValidationOptions() terra.SetValue[lightsailcertificate.DomainValidationOptionsAttributes] {
	return terra.ReferenceAsSet[lightsailcertificate.DomainValidationOptionsAttributes](lc.ref.Append("domain_validation_options"))
}

type lightsailCertificateState struct {
	Arn                     string                                              `json:"arn"`
	CreatedAt               string                                              `json:"created_at"`
	DomainName              string                                              `json:"domain_name"`
	Id                      string                                              `json:"id"`
	Name                    string                                              `json:"name"`
	SubjectAlternativeNames []string                                            `json:"subject_alternative_names"`
	Tags                    map[string]string                                   `json:"tags"`
	TagsAll                 map[string]string                                   `json:"tags_all"`
	DomainValidationOptions []lightsailcertificate.DomainValidationOptionsState `json:"domain_validation_options"`
}
