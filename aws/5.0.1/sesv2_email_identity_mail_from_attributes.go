// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesv2EmailIdentityMailFromAttributes creates a new instance of [Sesv2EmailIdentityMailFromAttributes].
func NewSesv2EmailIdentityMailFromAttributes(name string, args Sesv2EmailIdentityMailFromAttributesArgs) *Sesv2EmailIdentityMailFromAttributes {
	return &Sesv2EmailIdentityMailFromAttributes{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Sesv2EmailIdentityMailFromAttributes)(nil)

// Sesv2EmailIdentityMailFromAttributes represents the Terraform resource aws_sesv2_email_identity_mail_from_attributes.
type Sesv2EmailIdentityMailFromAttributes struct {
	Name      string
	Args      Sesv2EmailIdentityMailFromAttributesArgs
	state     *sesv2EmailIdentityMailFromAttributesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Sesv2EmailIdentityMailFromAttributes].
func (seimfa *Sesv2EmailIdentityMailFromAttributes) Type() string {
	return "aws_sesv2_email_identity_mail_from_attributes"
}

// LocalName returns the local name for [Sesv2EmailIdentityMailFromAttributes].
func (seimfa *Sesv2EmailIdentityMailFromAttributes) LocalName() string {
	return seimfa.Name
}

// Configuration returns the configuration (args) for [Sesv2EmailIdentityMailFromAttributes].
func (seimfa *Sesv2EmailIdentityMailFromAttributes) Configuration() interface{} {
	return seimfa.Args
}

// DependOn is used for other resources to depend on [Sesv2EmailIdentityMailFromAttributes].
func (seimfa *Sesv2EmailIdentityMailFromAttributes) DependOn() terra.Reference {
	return terra.ReferenceResource(seimfa)
}

// Dependencies returns the list of resources [Sesv2EmailIdentityMailFromAttributes] depends_on.
func (seimfa *Sesv2EmailIdentityMailFromAttributes) Dependencies() terra.Dependencies {
	return seimfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Sesv2EmailIdentityMailFromAttributes].
func (seimfa *Sesv2EmailIdentityMailFromAttributes) LifecycleManagement() *terra.Lifecycle {
	return seimfa.Lifecycle
}

// Attributes returns the attributes for [Sesv2EmailIdentityMailFromAttributes].
func (seimfa *Sesv2EmailIdentityMailFromAttributes) Attributes() sesv2EmailIdentityMailFromAttributesAttributes {
	return sesv2EmailIdentityMailFromAttributesAttributes{ref: terra.ReferenceResource(seimfa)}
}

// ImportState imports the given attribute values into [Sesv2EmailIdentityMailFromAttributes]'s state.
func (seimfa *Sesv2EmailIdentityMailFromAttributes) ImportState(av io.Reader) error {
	seimfa.state = &sesv2EmailIdentityMailFromAttributesState{}
	if err := json.NewDecoder(av).Decode(seimfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", seimfa.Type(), seimfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Sesv2EmailIdentityMailFromAttributes] has state.
func (seimfa *Sesv2EmailIdentityMailFromAttributes) State() (*sesv2EmailIdentityMailFromAttributesState, bool) {
	return seimfa.state, seimfa.state != nil
}

// StateMust returns the state for [Sesv2EmailIdentityMailFromAttributes]. Panics if the state is nil.
func (seimfa *Sesv2EmailIdentityMailFromAttributes) StateMust() *sesv2EmailIdentityMailFromAttributesState {
	if seimfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", seimfa.Type(), seimfa.LocalName()))
	}
	return seimfa.state
}

// Sesv2EmailIdentityMailFromAttributesArgs contains the configurations for aws_sesv2_email_identity_mail_from_attributes.
type Sesv2EmailIdentityMailFromAttributesArgs struct {
	// BehaviorOnMxFailure: string, optional
	BehaviorOnMxFailure terra.StringValue `hcl:"behavior_on_mx_failure,attr"`
	// EmailIdentity: string, required
	EmailIdentity terra.StringValue `hcl:"email_identity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MailFromDomain: string, optional
	MailFromDomain terra.StringValue `hcl:"mail_from_domain,attr"`
}
type sesv2EmailIdentityMailFromAttributesAttributes struct {
	ref terra.Reference
}

// BehaviorOnMxFailure returns a reference to field behavior_on_mx_failure of aws_sesv2_email_identity_mail_from_attributes.
func (seimfa sesv2EmailIdentityMailFromAttributesAttributes) BehaviorOnMxFailure() terra.StringValue {
	return terra.ReferenceAsString(seimfa.ref.Append("behavior_on_mx_failure"))
}

// EmailIdentity returns a reference to field email_identity of aws_sesv2_email_identity_mail_from_attributes.
func (seimfa sesv2EmailIdentityMailFromAttributesAttributes) EmailIdentity() terra.StringValue {
	return terra.ReferenceAsString(seimfa.ref.Append("email_identity"))
}

// Id returns a reference to field id of aws_sesv2_email_identity_mail_from_attributes.
func (seimfa sesv2EmailIdentityMailFromAttributesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(seimfa.ref.Append("id"))
}

// MailFromDomain returns a reference to field mail_from_domain of aws_sesv2_email_identity_mail_from_attributes.
func (seimfa sesv2EmailIdentityMailFromAttributesAttributes) MailFromDomain() terra.StringValue {
	return terra.ReferenceAsString(seimfa.ref.Append("mail_from_domain"))
}

type sesv2EmailIdentityMailFromAttributesState struct {
	BehaviorOnMxFailure string `json:"behavior_on_mx_failure"`
	EmailIdentity       string `json:"email_identity"`
	Id                  string `json:"id"`
	MailFromDomain      string `json:"mail_from_domain"`
}
