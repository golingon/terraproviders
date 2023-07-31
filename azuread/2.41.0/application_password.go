// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	applicationpassword "github.com/golingon/terraproviders/azuread/2.41.0/applicationpassword"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationPassword creates a new instance of [ApplicationPassword].
func NewApplicationPassword(name string, args ApplicationPasswordArgs) *ApplicationPassword {
	return &ApplicationPassword{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationPassword)(nil)

// ApplicationPassword represents the Terraform resource azuread_application_password.
type ApplicationPassword struct {
	Name      string
	Args      ApplicationPasswordArgs
	state     *applicationPasswordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationPassword].
func (ap *ApplicationPassword) Type() string {
	return "azuread_application_password"
}

// LocalName returns the local name for [ApplicationPassword].
func (ap *ApplicationPassword) LocalName() string {
	return ap.Name
}

// Configuration returns the configuration (args) for [ApplicationPassword].
func (ap *ApplicationPassword) Configuration() interface{} {
	return ap.Args
}

// DependOn is used for other resources to depend on [ApplicationPassword].
func (ap *ApplicationPassword) DependOn() terra.Reference {
	return terra.ReferenceResource(ap)
}

// Dependencies returns the list of resources [ApplicationPassword] depends_on.
func (ap *ApplicationPassword) Dependencies() terra.Dependencies {
	return ap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationPassword].
func (ap *ApplicationPassword) LifecycleManagement() *terra.Lifecycle {
	return ap.Lifecycle
}

// Attributes returns the attributes for [ApplicationPassword].
func (ap *ApplicationPassword) Attributes() applicationPasswordAttributes {
	return applicationPasswordAttributes{ref: terra.ReferenceResource(ap)}
}

// ImportState imports the given attribute values into [ApplicationPassword]'s state.
func (ap *ApplicationPassword) ImportState(av io.Reader) error {
	ap.state = &applicationPasswordState{}
	if err := json.NewDecoder(av).Decode(ap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ap.Type(), ap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationPassword] has state.
func (ap *ApplicationPassword) State() (*applicationPasswordState, bool) {
	return ap.state, ap.state != nil
}

// StateMust returns the state for [ApplicationPassword]. Panics if the state is nil.
func (ap *ApplicationPassword) StateMust() *applicationPasswordState {
	if ap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ap.Type(), ap.LocalName()))
	}
	return ap.state
}

// ApplicationPasswordArgs contains the configurations for azuread_application_password.
type ApplicationPasswordArgs struct {
	// ApplicationObjectId: string, required
	ApplicationObjectId terra.StringValue `hcl:"application_object_id,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// EndDate: string, optional
	EndDate terra.StringValue `hcl:"end_date,attr"`
	// EndDateRelative: string, optional
	EndDateRelative terra.StringValue `hcl:"end_date_relative,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RotateWhenChanged: map of string, optional
	RotateWhenChanged terra.MapValue[terra.StringValue] `hcl:"rotate_when_changed,attr"`
	// StartDate: string, optional
	StartDate terra.StringValue `hcl:"start_date,attr"`
	// Timeouts: optional
	Timeouts *applicationpassword.Timeouts `hcl:"timeouts,block"`
}
type applicationPasswordAttributes struct {
	ref terra.Reference
}

// ApplicationObjectId returns a reference to field application_object_id of azuread_application_password.
func (ap applicationPasswordAttributes) ApplicationObjectId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("application_object_id"))
}

// DisplayName returns a reference to field display_name of azuread_application_password.
func (ap applicationPasswordAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("display_name"))
}

// EndDate returns a reference to field end_date of azuread_application_password.
func (ap applicationPasswordAttributes) EndDate() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("end_date"))
}

// EndDateRelative returns a reference to field end_date_relative of azuread_application_password.
func (ap applicationPasswordAttributes) EndDateRelative() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("end_date_relative"))
}

// Id returns a reference to field id of azuread_application_password.
func (ap applicationPasswordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("id"))
}

// KeyId returns a reference to field key_id of azuread_application_password.
func (ap applicationPasswordAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("key_id"))
}

// RotateWhenChanged returns a reference to field rotate_when_changed of azuread_application_password.
func (ap applicationPasswordAttributes) RotateWhenChanged() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ap.ref.Append("rotate_when_changed"))
}

// StartDate returns a reference to field start_date of azuread_application_password.
func (ap applicationPasswordAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("start_date"))
}

// Value returns a reference to field value of azuread_application_password.
func (ap applicationPasswordAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("value"))
}

func (ap applicationPasswordAttributes) Timeouts() applicationpassword.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationpassword.TimeoutsAttributes](ap.ref.Append("timeouts"))
}

type applicationPasswordState struct {
	ApplicationObjectId string                             `json:"application_object_id"`
	DisplayName         string                             `json:"display_name"`
	EndDate             string                             `json:"end_date"`
	EndDateRelative     string                             `json:"end_date_relative"`
	Id                  string                             `json:"id"`
	KeyId               string                             `json:"key_id"`
	RotateWhenChanged   map[string]string                  `json:"rotate_when_changed"`
	StartDate           string                             `json:"start_date"`
	Value               string                             `json:"value"`
	Timeouts            *applicationpassword.TimeoutsState `json:"timeouts"`
}
