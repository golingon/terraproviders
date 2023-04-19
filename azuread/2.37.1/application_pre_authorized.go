// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	applicationpreauthorized "github.com/golingon/terraproviders/azuread/2.37.1/applicationpreauthorized"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationPreAuthorized creates a new instance of [ApplicationPreAuthorized].
func NewApplicationPreAuthorized(name string, args ApplicationPreAuthorizedArgs) *ApplicationPreAuthorized {
	return &ApplicationPreAuthorized{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationPreAuthorized)(nil)

// ApplicationPreAuthorized represents the Terraform resource azuread_application_pre_authorized.
type ApplicationPreAuthorized struct {
	Name      string
	Args      ApplicationPreAuthorizedArgs
	state     *applicationPreAuthorizedState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationPreAuthorized].
func (apa *ApplicationPreAuthorized) Type() string {
	return "azuread_application_pre_authorized"
}

// LocalName returns the local name for [ApplicationPreAuthorized].
func (apa *ApplicationPreAuthorized) LocalName() string {
	return apa.Name
}

// Configuration returns the configuration (args) for [ApplicationPreAuthorized].
func (apa *ApplicationPreAuthorized) Configuration() interface{} {
	return apa.Args
}

// DependOn is used for other resources to depend on [ApplicationPreAuthorized].
func (apa *ApplicationPreAuthorized) DependOn() terra.Reference {
	return terra.ReferenceResource(apa)
}

// Dependencies returns the list of resources [ApplicationPreAuthorized] depends_on.
func (apa *ApplicationPreAuthorized) Dependencies() terra.Dependencies {
	return apa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationPreAuthorized].
func (apa *ApplicationPreAuthorized) LifecycleManagement() *terra.Lifecycle {
	return apa.Lifecycle
}

// Attributes returns the attributes for [ApplicationPreAuthorized].
func (apa *ApplicationPreAuthorized) Attributes() applicationPreAuthorizedAttributes {
	return applicationPreAuthorizedAttributes{ref: terra.ReferenceResource(apa)}
}

// ImportState imports the given attribute values into [ApplicationPreAuthorized]'s state.
func (apa *ApplicationPreAuthorized) ImportState(av io.Reader) error {
	apa.state = &applicationPreAuthorizedState{}
	if err := json.NewDecoder(av).Decode(apa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apa.Type(), apa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationPreAuthorized] has state.
func (apa *ApplicationPreAuthorized) State() (*applicationPreAuthorizedState, bool) {
	return apa.state, apa.state != nil
}

// StateMust returns the state for [ApplicationPreAuthorized]. Panics if the state is nil.
func (apa *ApplicationPreAuthorized) StateMust() *applicationPreAuthorizedState {
	if apa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apa.Type(), apa.LocalName()))
	}
	return apa.state
}

// ApplicationPreAuthorizedArgs contains the configurations for azuread_application_pre_authorized.
type ApplicationPreAuthorizedArgs struct {
	// ApplicationObjectId: string, required
	ApplicationObjectId terra.StringValue `hcl:"application_object_id,attr" validate:"required"`
	// AuthorizedAppId: string, required
	AuthorizedAppId terra.StringValue `hcl:"authorized_app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PermissionIds: set of string, required
	PermissionIds terra.SetValue[terra.StringValue] `hcl:"permission_ids,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *applicationpreauthorized.Timeouts `hcl:"timeouts,block"`
}
type applicationPreAuthorizedAttributes struct {
	ref terra.Reference
}

// ApplicationObjectId returns a reference to field application_object_id of azuread_application_pre_authorized.
func (apa applicationPreAuthorizedAttributes) ApplicationObjectId() terra.StringValue {
	return terra.ReferenceAsString(apa.ref.Append("application_object_id"))
}

// AuthorizedAppId returns a reference to field authorized_app_id of azuread_application_pre_authorized.
func (apa applicationPreAuthorizedAttributes) AuthorizedAppId() terra.StringValue {
	return terra.ReferenceAsString(apa.ref.Append("authorized_app_id"))
}

// Id returns a reference to field id of azuread_application_pre_authorized.
func (apa applicationPreAuthorizedAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apa.ref.Append("id"))
}

// PermissionIds returns a reference to field permission_ids of azuread_application_pre_authorized.
func (apa applicationPreAuthorizedAttributes) PermissionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](apa.ref.Append("permission_ids"))
}

func (apa applicationPreAuthorizedAttributes) Timeouts() applicationpreauthorized.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationpreauthorized.TimeoutsAttributes](apa.ref.Append("timeouts"))
}

type applicationPreAuthorizedState struct {
	ApplicationObjectId string                                  `json:"application_object_id"`
	AuthorizedAppId     string                                  `json:"authorized_app_id"`
	Id                  string                                  `json:"id"`
	PermissionIds       []string                                `json:"permission_ids"`
	Timeouts            *applicationpreauthorized.TimeoutsState `json:"timeouts"`
}
