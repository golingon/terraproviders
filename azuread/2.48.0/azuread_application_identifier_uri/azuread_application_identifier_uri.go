// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread_application_identifier_uri

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

// Resource represents the Terraform resource azuread_application_identifier_uri.
type Resource struct {
	Name      string
	Args      Args
	state     *azureadApplicationIdentifierUriState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aaiu *Resource) Type() string {
	return "azuread_application_identifier_uri"
}

// LocalName returns the local name for [Resource].
func (aaiu *Resource) LocalName() string {
	return aaiu.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aaiu *Resource) Configuration() interface{} {
	return aaiu.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aaiu *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aaiu)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aaiu *Resource) Dependencies() terra.Dependencies {
	return aaiu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aaiu *Resource) LifecycleManagement() *terra.Lifecycle {
	return aaiu.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aaiu *Resource) Attributes() azureadApplicationIdentifierUriAttributes {
	return azureadApplicationIdentifierUriAttributes{ref: terra.ReferenceResource(aaiu)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aaiu *Resource) ImportState(state io.Reader) error {
	aaiu.state = &azureadApplicationIdentifierUriState{}
	if err := json.NewDecoder(state).Decode(aaiu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aaiu.Type(), aaiu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aaiu *Resource) State() (*azureadApplicationIdentifierUriState, bool) {
	return aaiu.state, aaiu.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aaiu *Resource) StateMust() *azureadApplicationIdentifierUriState {
	if aaiu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aaiu.Type(), aaiu.LocalName()))
	}
	return aaiu.state
}

// Args contains the configurations for azuread_application_identifier_uri.
type Args struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentifierUri: string, required
	IdentifierUri terra.StringValue `hcl:"identifier_uri,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azureadApplicationIdentifierUriAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of azuread_application_identifier_uri.
func (aaiu azureadApplicationIdentifierUriAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(aaiu.ref.Append("application_id"))
}

// Id returns a reference to field id of azuread_application_identifier_uri.
func (aaiu azureadApplicationIdentifierUriAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaiu.ref.Append("id"))
}

// IdentifierUri returns a reference to field identifier_uri of azuread_application_identifier_uri.
func (aaiu azureadApplicationIdentifierUriAttributes) IdentifierUri() terra.StringValue {
	return terra.ReferenceAsString(aaiu.ref.Append("identifier_uri"))
}

func (aaiu azureadApplicationIdentifierUriAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aaiu.ref.Append("timeouts"))
}

type azureadApplicationIdentifierUriState struct {
	ApplicationId string         `json:"application_id"`
	Id            string         `json:"id"`
	IdentifierUri string         `json:"identifier_uri"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
