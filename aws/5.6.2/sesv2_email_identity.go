// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sesv2emailidentity "github.com/golingon/terraproviders/aws/5.6.2/sesv2emailidentity"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesv2EmailIdentity creates a new instance of [Sesv2EmailIdentity].
func NewSesv2EmailIdentity(name string, args Sesv2EmailIdentityArgs) *Sesv2EmailIdentity {
	return &Sesv2EmailIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Sesv2EmailIdentity)(nil)

// Sesv2EmailIdentity represents the Terraform resource aws_sesv2_email_identity.
type Sesv2EmailIdentity struct {
	Name      string
	Args      Sesv2EmailIdentityArgs
	state     *sesv2EmailIdentityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Sesv2EmailIdentity].
func (sei *Sesv2EmailIdentity) Type() string {
	return "aws_sesv2_email_identity"
}

// LocalName returns the local name for [Sesv2EmailIdentity].
func (sei *Sesv2EmailIdentity) LocalName() string {
	return sei.Name
}

// Configuration returns the configuration (args) for [Sesv2EmailIdentity].
func (sei *Sesv2EmailIdentity) Configuration() interface{} {
	return sei.Args
}

// DependOn is used for other resources to depend on [Sesv2EmailIdentity].
func (sei *Sesv2EmailIdentity) DependOn() terra.Reference {
	return terra.ReferenceResource(sei)
}

// Dependencies returns the list of resources [Sesv2EmailIdentity] depends_on.
func (sei *Sesv2EmailIdentity) Dependencies() terra.Dependencies {
	return sei.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Sesv2EmailIdentity].
func (sei *Sesv2EmailIdentity) LifecycleManagement() *terra.Lifecycle {
	return sei.Lifecycle
}

// Attributes returns the attributes for [Sesv2EmailIdentity].
func (sei *Sesv2EmailIdentity) Attributes() sesv2EmailIdentityAttributes {
	return sesv2EmailIdentityAttributes{ref: terra.ReferenceResource(sei)}
}

// ImportState imports the given attribute values into [Sesv2EmailIdentity]'s state.
func (sei *Sesv2EmailIdentity) ImportState(av io.Reader) error {
	sei.state = &sesv2EmailIdentityState{}
	if err := json.NewDecoder(av).Decode(sei.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sei.Type(), sei.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Sesv2EmailIdentity] has state.
func (sei *Sesv2EmailIdentity) State() (*sesv2EmailIdentityState, bool) {
	return sei.state, sei.state != nil
}

// StateMust returns the state for [Sesv2EmailIdentity]. Panics if the state is nil.
func (sei *Sesv2EmailIdentity) StateMust() *sesv2EmailIdentityState {
	if sei.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sei.Type(), sei.LocalName()))
	}
	return sei.state
}

// Sesv2EmailIdentityArgs contains the configurations for aws_sesv2_email_identity.
type Sesv2EmailIdentityArgs struct {
	// ConfigurationSetName: string, optional
	ConfigurationSetName terra.StringValue `hcl:"configuration_set_name,attr"`
	// EmailIdentity: string, required
	EmailIdentity terra.StringValue `hcl:"email_identity,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DkimSigningAttributes: optional
	DkimSigningAttributes *sesv2emailidentity.DkimSigningAttributes `hcl:"dkim_signing_attributes,block"`
}
type sesv2EmailIdentityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("arn"))
}

// ConfigurationSetName returns a reference to field configuration_set_name of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) ConfigurationSetName() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("configuration_set_name"))
}

// EmailIdentity returns a reference to field email_identity of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) EmailIdentity() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("email_identity"))
}

// Id returns a reference to field id of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("id"))
}

// IdentityType returns a reference to field identity_type of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceAsString(sei.ref.Append("identity_type"))
}

// Tags returns a reference to field tags of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sei.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sei.ref.Append("tags_all"))
}

// VerifiedForSendingStatus returns a reference to field verified_for_sending_status of aws_sesv2_email_identity.
func (sei sesv2EmailIdentityAttributes) VerifiedForSendingStatus() terra.BoolValue {
	return terra.ReferenceAsBool(sei.ref.Append("verified_for_sending_status"))
}

func (sei sesv2EmailIdentityAttributes) DkimSigningAttributes() terra.ListValue[sesv2emailidentity.DkimSigningAttributesAttributes] {
	return terra.ReferenceAsList[sesv2emailidentity.DkimSigningAttributesAttributes](sei.ref.Append("dkim_signing_attributes"))
}

type sesv2EmailIdentityState struct {
	Arn                      string                                          `json:"arn"`
	ConfigurationSetName     string                                          `json:"configuration_set_name"`
	EmailIdentity            string                                          `json:"email_identity"`
	Id                       string                                          `json:"id"`
	IdentityType             string                                          `json:"identity_type"`
	Tags                     map[string]string                               `json:"tags"`
	TagsAll                  map[string]string                               `json:"tags_all"`
	VerifiedForSendingStatus bool                                            `json:"verified_for_sending_status"`
	DkimSigningAttributes    []sesv2emailidentity.DkimSigningAttributesState `json:"dkim_signing_attributes"`
}
