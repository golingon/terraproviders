// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	applicationidentifieruri "github.com/golingon/terraproviders/azuread/2.47.0/applicationidentifieruri"
	"io"
)

// NewApplicationIdentifierUri creates a new instance of [ApplicationIdentifierUri].
func NewApplicationIdentifierUri(name string, args ApplicationIdentifierUriArgs) *ApplicationIdentifierUri {
	return &ApplicationIdentifierUri{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationIdentifierUri)(nil)

// ApplicationIdentifierUri represents the Terraform resource azuread_application_identifier_uri.
type ApplicationIdentifierUri struct {
	Name      string
	Args      ApplicationIdentifierUriArgs
	state     *applicationIdentifierUriState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationIdentifierUri].
func (aiu *ApplicationIdentifierUri) Type() string {
	return "azuread_application_identifier_uri"
}

// LocalName returns the local name for [ApplicationIdentifierUri].
func (aiu *ApplicationIdentifierUri) LocalName() string {
	return aiu.Name
}

// Configuration returns the configuration (args) for [ApplicationIdentifierUri].
func (aiu *ApplicationIdentifierUri) Configuration() interface{} {
	return aiu.Args
}

// DependOn is used for other resources to depend on [ApplicationIdentifierUri].
func (aiu *ApplicationIdentifierUri) DependOn() terra.Reference {
	return terra.ReferenceResource(aiu)
}

// Dependencies returns the list of resources [ApplicationIdentifierUri] depends_on.
func (aiu *ApplicationIdentifierUri) Dependencies() terra.Dependencies {
	return aiu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationIdentifierUri].
func (aiu *ApplicationIdentifierUri) LifecycleManagement() *terra.Lifecycle {
	return aiu.Lifecycle
}

// Attributes returns the attributes for [ApplicationIdentifierUri].
func (aiu *ApplicationIdentifierUri) Attributes() applicationIdentifierUriAttributes {
	return applicationIdentifierUriAttributes{ref: terra.ReferenceResource(aiu)}
}

// ImportState imports the given attribute values into [ApplicationIdentifierUri]'s state.
func (aiu *ApplicationIdentifierUri) ImportState(av io.Reader) error {
	aiu.state = &applicationIdentifierUriState{}
	if err := json.NewDecoder(av).Decode(aiu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aiu.Type(), aiu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationIdentifierUri] has state.
func (aiu *ApplicationIdentifierUri) State() (*applicationIdentifierUriState, bool) {
	return aiu.state, aiu.state != nil
}

// StateMust returns the state for [ApplicationIdentifierUri]. Panics if the state is nil.
func (aiu *ApplicationIdentifierUri) StateMust() *applicationIdentifierUriState {
	if aiu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aiu.Type(), aiu.LocalName()))
	}
	return aiu.state
}

// ApplicationIdentifierUriArgs contains the configurations for azuread_application_identifier_uri.
type ApplicationIdentifierUriArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentifierUri: string, required
	IdentifierUri terra.StringValue `hcl:"identifier_uri,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *applicationidentifieruri.Timeouts `hcl:"timeouts,block"`
}
type applicationIdentifierUriAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of azuread_application_identifier_uri.
func (aiu applicationIdentifierUriAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(aiu.ref.Append("application_id"))
}

// Id returns a reference to field id of azuread_application_identifier_uri.
func (aiu applicationIdentifierUriAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aiu.ref.Append("id"))
}

// IdentifierUri returns a reference to field identifier_uri of azuread_application_identifier_uri.
func (aiu applicationIdentifierUriAttributes) IdentifierUri() terra.StringValue {
	return terra.ReferenceAsString(aiu.ref.Append("identifier_uri"))
}

func (aiu applicationIdentifierUriAttributes) Timeouts() applicationidentifieruri.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationidentifieruri.TimeoutsAttributes](aiu.ref.Append("timeouts"))
}

type applicationIdentifierUriState struct {
	ApplicationId string                                  `json:"application_id"`
	Id            string                                  `json:"id"`
	IdentifierUri string                                  `json:"identifier_uri"`
	Timeouts      *applicationidentifieruri.TimeoutsState `json:"timeouts"`
}
