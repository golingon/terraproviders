// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	applicationredirecturis "github.com/golingon/terraproviders/azuread/2.47.0/applicationredirecturis"
	"io"
)

// NewApplicationRedirectUris creates a new instance of [ApplicationRedirectUris].
func NewApplicationRedirectUris(name string, args ApplicationRedirectUrisArgs) *ApplicationRedirectUris {
	return &ApplicationRedirectUris{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationRedirectUris)(nil)

// ApplicationRedirectUris represents the Terraform resource azuread_application_redirect_uris.
type ApplicationRedirectUris struct {
	Name      string
	Args      ApplicationRedirectUrisArgs
	state     *applicationRedirectUrisState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationRedirectUris].
func (aru *ApplicationRedirectUris) Type() string {
	return "azuread_application_redirect_uris"
}

// LocalName returns the local name for [ApplicationRedirectUris].
func (aru *ApplicationRedirectUris) LocalName() string {
	return aru.Name
}

// Configuration returns the configuration (args) for [ApplicationRedirectUris].
func (aru *ApplicationRedirectUris) Configuration() interface{} {
	return aru.Args
}

// DependOn is used for other resources to depend on [ApplicationRedirectUris].
func (aru *ApplicationRedirectUris) DependOn() terra.Reference {
	return terra.ReferenceResource(aru)
}

// Dependencies returns the list of resources [ApplicationRedirectUris] depends_on.
func (aru *ApplicationRedirectUris) Dependencies() terra.Dependencies {
	return aru.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationRedirectUris].
func (aru *ApplicationRedirectUris) LifecycleManagement() *terra.Lifecycle {
	return aru.Lifecycle
}

// Attributes returns the attributes for [ApplicationRedirectUris].
func (aru *ApplicationRedirectUris) Attributes() applicationRedirectUrisAttributes {
	return applicationRedirectUrisAttributes{ref: terra.ReferenceResource(aru)}
}

// ImportState imports the given attribute values into [ApplicationRedirectUris]'s state.
func (aru *ApplicationRedirectUris) ImportState(av io.Reader) error {
	aru.state = &applicationRedirectUrisState{}
	if err := json.NewDecoder(av).Decode(aru.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aru.Type(), aru.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationRedirectUris] has state.
func (aru *ApplicationRedirectUris) State() (*applicationRedirectUrisState, bool) {
	return aru.state, aru.state != nil
}

// StateMust returns the state for [ApplicationRedirectUris]. Panics if the state is nil.
func (aru *ApplicationRedirectUris) StateMust() *applicationRedirectUrisState {
	if aru.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aru.Type(), aru.LocalName()))
	}
	return aru.state
}

// ApplicationRedirectUrisArgs contains the configurations for azuread_application_redirect_uris.
type ApplicationRedirectUrisArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RedirectUris: set of string, required
	RedirectUris terra.SetValue[terra.StringValue] `hcl:"redirect_uris,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *applicationredirecturis.Timeouts `hcl:"timeouts,block"`
}
type applicationRedirectUrisAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of azuread_application_redirect_uris.
func (aru applicationRedirectUrisAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(aru.ref.Append("application_id"))
}

// Id returns a reference to field id of azuread_application_redirect_uris.
func (aru applicationRedirectUrisAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aru.ref.Append("id"))
}

// RedirectUris returns a reference to field redirect_uris of azuread_application_redirect_uris.
func (aru applicationRedirectUrisAttributes) RedirectUris() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aru.ref.Append("redirect_uris"))
}

// Type returns a reference to field type of azuread_application_redirect_uris.
func (aru applicationRedirectUrisAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aru.ref.Append("type"))
}

func (aru applicationRedirectUrisAttributes) Timeouts() applicationredirecturis.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationredirecturis.TimeoutsAttributes](aru.ref.Append("timeouts"))
}

type applicationRedirectUrisState struct {
	ApplicationId string                                 `json:"application_id"`
	Id            string                                 `json:"id"`
	RedirectUris  []string                               `json:"redirect_uris"`
	Type          string                                 `json:"type"`
	Timeouts      *applicationredirecturis.TimeoutsState `json:"timeouts"`
}
