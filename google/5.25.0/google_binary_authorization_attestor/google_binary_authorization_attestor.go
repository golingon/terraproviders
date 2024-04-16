// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_binary_authorization_attestor

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_binary_authorization_attestor.
type Resource struct {
	Name      string
	Args      Args
	state     *googleBinaryAuthorizationAttestorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gbaa *Resource) Type() string {
	return "google_binary_authorization_attestor"
}

// LocalName returns the local name for [Resource].
func (gbaa *Resource) LocalName() string {
	return gbaa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gbaa *Resource) Configuration() interface{} {
	return gbaa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gbaa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gbaa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gbaa *Resource) Dependencies() terra.Dependencies {
	return gbaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gbaa *Resource) LifecycleManagement() *terra.Lifecycle {
	return gbaa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gbaa *Resource) Attributes() googleBinaryAuthorizationAttestorAttributes {
	return googleBinaryAuthorizationAttestorAttributes{ref: terra.ReferenceResource(gbaa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gbaa *Resource) ImportState(state io.Reader) error {
	gbaa.state = &googleBinaryAuthorizationAttestorState{}
	if err := json.NewDecoder(state).Decode(gbaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbaa.Type(), gbaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gbaa *Resource) State() (*googleBinaryAuthorizationAttestorState, bool) {
	return gbaa.state, gbaa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gbaa *Resource) StateMust() *googleBinaryAuthorizationAttestorState {
	if gbaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbaa.Type(), gbaa.LocalName()))
	}
	return gbaa.state
}

// Args contains the configurations for google_binary_authorization_attestor.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AttestationAuthorityNote: required
	AttestationAuthorityNote *AttestationAuthorityNote `hcl:"attestation_authority_note,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleBinaryAuthorizationAttestorAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_binary_authorization_attestor.
func (gbaa googleBinaryAuthorizationAttestorAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gbaa.ref.Append("description"))
}

// Id returns a reference to field id of google_binary_authorization_attestor.
func (gbaa googleBinaryAuthorizationAttestorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbaa.ref.Append("id"))
}

// Name returns a reference to field name of google_binary_authorization_attestor.
func (gbaa googleBinaryAuthorizationAttestorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbaa.ref.Append("name"))
}

// Project returns a reference to field project of google_binary_authorization_attestor.
func (gbaa googleBinaryAuthorizationAttestorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbaa.ref.Append("project"))
}

func (gbaa googleBinaryAuthorizationAttestorAttributes) AttestationAuthorityNote() terra.ListValue[AttestationAuthorityNoteAttributes] {
	return terra.ReferenceAsList[AttestationAuthorityNoteAttributes](gbaa.ref.Append("attestation_authority_note"))
}

func (gbaa googleBinaryAuthorizationAttestorAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gbaa.ref.Append("timeouts"))
}

type googleBinaryAuthorizationAttestorState struct {
	Description              string                          `json:"description"`
	Id                       string                          `json:"id"`
	Name                     string                          `json:"name"`
	Project                  string                          `json:"project"`
	AttestationAuthorityNote []AttestationAuthorityNoteState `json:"attestation_authority_note"`
	Timeouts                 *TimeoutsState                  `json:"timeouts"`
}
