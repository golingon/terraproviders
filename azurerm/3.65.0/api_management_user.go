// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementuser "github.com/golingon/terraproviders/azurerm/3.65.0/apimanagementuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementUser creates a new instance of [ApiManagementUser].
func NewApiManagementUser(name string, args ApiManagementUserArgs) *ApiManagementUser {
	return &ApiManagementUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementUser)(nil)

// ApiManagementUser represents the Terraform resource azurerm_api_management_user.
type ApiManagementUser struct {
	Name      string
	Args      ApiManagementUserArgs
	state     *apiManagementUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementUser].
func (amu *ApiManagementUser) Type() string {
	return "azurerm_api_management_user"
}

// LocalName returns the local name for [ApiManagementUser].
func (amu *ApiManagementUser) LocalName() string {
	return amu.Name
}

// Configuration returns the configuration (args) for [ApiManagementUser].
func (amu *ApiManagementUser) Configuration() interface{} {
	return amu.Args
}

// DependOn is used for other resources to depend on [ApiManagementUser].
func (amu *ApiManagementUser) DependOn() terra.Reference {
	return terra.ReferenceResource(amu)
}

// Dependencies returns the list of resources [ApiManagementUser] depends_on.
func (amu *ApiManagementUser) Dependencies() terra.Dependencies {
	return amu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementUser].
func (amu *ApiManagementUser) LifecycleManagement() *terra.Lifecycle {
	return amu.Lifecycle
}

// Attributes returns the attributes for [ApiManagementUser].
func (amu *ApiManagementUser) Attributes() apiManagementUserAttributes {
	return apiManagementUserAttributes{ref: terra.ReferenceResource(amu)}
}

// ImportState imports the given attribute values into [ApiManagementUser]'s state.
func (amu *ApiManagementUser) ImportState(av io.Reader) error {
	amu.state = &apiManagementUserState{}
	if err := json.NewDecoder(av).Decode(amu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amu.Type(), amu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementUser] has state.
func (amu *ApiManagementUser) State() (*apiManagementUserState, bool) {
	return amu.state, amu.state != nil
}

// StateMust returns the state for [ApiManagementUser]. Panics if the state is nil.
func (amu *ApiManagementUser) StateMust() *apiManagementUserState {
	if amu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amu.Type(), amu.LocalName()))
	}
	return amu.state
}

// ApiManagementUserArgs contains the configurations for azurerm_api_management_user.
type ApiManagementUserArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Confirmation: string, optional
	Confirmation terra.StringValue `hcl:"confirmation,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// FirstName: string, required
	FirstName terra.StringValue `hcl:"first_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LastName: string, required
	LastName terra.StringValue `hcl:"last_name,attr" validate:"required"`
	// Note: string, optional
	Note terra.StringValue `hcl:"note,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// UserId: string, required
	UserId terra.StringValue `hcl:"user_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementuser.Timeouts `hcl:"timeouts,block"`
}
type apiManagementUserAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_user.
func (amu apiManagementUserAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("api_management_name"))
}

// Confirmation returns a reference to field confirmation of azurerm_api_management_user.
func (amu apiManagementUserAttributes) Confirmation() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("confirmation"))
}

// Email returns a reference to field email of azurerm_api_management_user.
func (amu apiManagementUserAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("email"))
}

// FirstName returns a reference to field first_name of azurerm_api_management_user.
func (amu apiManagementUserAttributes) FirstName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("first_name"))
}

// Id returns a reference to field id of azurerm_api_management_user.
func (amu apiManagementUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("id"))
}

// LastName returns a reference to field last_name of azurerm_api_management_user.
func (amu apiManagementUserAttributes) LastName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("last_name"))
}

// Note returns a reference to field note of azurerm_api_management_user.
func (amu apiManagementUserAttributes) Note() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("note"))
}

// Password returns a reference to field password of azurerm_api_management_user.
func (amu apiManagementUserAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_user.
func (amu apiManagementUserAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("resource_group_name"))
}

// State returns a reference to field state of azurerm_api_management_user.
func (amu apiManagementUserAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("state"))
}

// UserId returns a reference to field user_id of azurerm_api_management_user.
func (amu apiManagementUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(amu.ref.Append("user_id"))
}

func (amu apiManagementUserAttributes) Timeouts() apimanagementuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementuser.TimeoutsAttributes](amu.ref.Append("timeouts"))
}

type apiManagementUserState struct {
	ApiManagementName string                           `json:"api_management_name"`
	Confirmation      string                           `json:"confirmation"`
	Email             string                           `json:"email"`
	FirstName         string                           `json:"first_name"`
	Id                string                           `json:"id"`
	LastName          string                           `json:"last_name"`
	Note              string                           `json:"note"`
	Password          string                           `json:"password"`
	ResourceGroupName string                           `json:"resource_group_name"`
	State             string                           `json:"state"`
	UserId            string                           `json:"user_id"`
	Timeouts          *apimanagementuser.TimeoutsState `json:"timeouts"`
}
