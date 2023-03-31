// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSesDomainIdentity creates a new instance of [SesDomainIdentity].
func NewSesDomainIdentity(name string, args SesDomainIdentityArgs) *SesDomainIdentity {
	return &SesDomainIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SesDomainIdentity)(nil)

// SesDomainIdentity represents the Terraform resource aws_ses_domain_identity.
type SesDomainIdentity struct {
	Name      string
	Args      SesDomainIdentityArgs
	state     *sesDomainIdentityState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SesDomainIdentity].
func (sdi *SesDomainIdentity) Type() string {
	return "aws_ses_domain_identity"
}

// LocalName returns the local name for [SesDomainIdentity].
func (sdi *SesDomainIdentity) LocalName() string {
	return sdi.Name
}

// Configuration returns the configuration (args) for [SesDomainIdentity].
func (sdi *SesDomainIdentity) Configuration() interface{} {
	return sdi.Args
}

// DependOn is used for other resources to depend on [SesDomainIdentity].
func (sdi *SesDomainIdentity) DependOn() terra.Reference {
	return terra.ReferenceResource(sdi)
}

// Dependencies returns the list of resources [SesDomainIdentity] depends_on.
func (sdi *SesDomainIdentity) Dependencies() terra.Dependencies {
	return sdi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SesDomainIdentity].
func (sdi *SesDomainIdentity) LifecycleManagement() *terra.Lifecycle {
	return sdi.Lifecycle
}

// Attributes returns the attributes for [SesDomainIdentity].
func (sdi *SesDomainIdentity) Attributes() sesDomainIdentityAttributes {
	return sesDomainIdentityAttributes{ref: terra.ReferenceResource(sdi)}
}

// ImportState imports the given attribute values into [SesDomainIdentity]'s state.
func (sdi *SesDomainIdentity) ImportState(av io.Reader) error {
	sdi.state = &sesDomainIdentityState{}
	if err := json.NewDecoder(av).Decode(sdi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdi.Type(), sdi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SesDomainIdentity] has state.
func (sdi *SesDomainIdentity) State() (*sesDomainIdentityState, bool) {
	return sdi.state, sdi.state != nil
}

// StateMust returns the state for [SesDomainIdentity]. Panics if the state is nil.
func (sdi *SesDomainIdentity) StateMust() *sesDomainIdentityState {
	if sdi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdi.Type(), sdi.LocalName()))
	}
	return sdi.state
}

// SesDomainIdentityArgs contains the configurations for aws_ses_domain_identity.
type SesDomainIdentityArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type sesDomainIdentityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ses_domain_identity.
func (sdi sesDomainIdentityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("arn"))
}

// Domain returns a reference to field domain of aws_ses_domain_identity.
func (sdi sesDomainIdentityAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("domain"))
}

// Id returns a reference to field id of aws_ses_domain_identity.
func (sdi sesDomainIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("id"))
}

// VerificationToken returns a reference to field verification_token of aws_ses_domain_identity.
func (sdi sesDomainIdentityAttributes) VerificationToken() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("verification_token"))
}

type sesDomainIdentityState struct {
	Arn               string `json:"arn"`
	Domain            string `json:"domain"`
	Id                string `json:"id"`
	VerificationToken string `json:"verification_token"`
}