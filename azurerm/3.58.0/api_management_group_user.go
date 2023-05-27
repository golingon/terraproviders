// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementgroupuser "github.com/golingon/terraproviders/azurerm/3.58.0/apimanagementgroupuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementGroupUser creates a new instance of [ApiManagementGroupUser].
func NewApiManagementGroupUser(name string, args ApiManagementGroupUserArgs) *ApiManagementGroupUser {
	return &ApiManagementGroupUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementGroupUser)(nil)

// ApiManagementGroupUser represents the Terraform resource azurerm_api_management_group_user.
type ApiManagementGroupUser struct {
	Name      string
	Args      ApiManagementGroupUserArgs
	state     *apiManagementGroupUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementGroupUser].
func (amgu *ApiManagementGroupUser) Type() string {
	return "azurerm_api_management_group_user"
}

// LocalName returns the local name for [ApiManagementGroupUser].
func (amgu *ApiManagementGroupUser) LocalName() string {
	return amgu.Name
}

// Configuration returns the configuration (args) for [ApiManagementGroupUser].
func (amgu *ApiManagementGroupUser) Configuration() interface{} {
	return amgu.Args
}

// DependOn is used for other resources to depend on [ApiManagementGroupUser].
func (amgu *ApiManagementGroupUser) DependOn() terra.Reference {
	return terra.ReferenceResource(amgu)
}

// Dependencies returns the list of resources [ApiManagementGroupUser] depends_on.
func (amgu *ApiManagementGroupUser) Dependencies() terra.Dependencies {
	return amgu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementGroupUser].
func (amgu *ApiManagementGroupUser) LifecycleManagement() *terra.Lifecycle {
	return amgu.Lifecycle
}

// Attributes returns the attributes for [ApiManagementGroupUser].
func (amgu *ApiManagementGroupUser) Attributes() apiManagementGroupUserAttributes {
	return apiManagementGroupUserAttributes{ref: terra.ReferenceResource(amgu)}
}

// ImportState imports the given attribute values into [ApiManagementGroupUser]'s state.
func (amgu *ApiManagementGroupUser) ImportState(av io.Reader) error {
	amgu.state = &apiManagementGroupUserState{}
	if err := json.NewDecoder(av).Decode(amgu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amgu.Type(), amgu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementGroupUser] has state.
func (amgu *ApiManagementGroupUser) State() (*apiManagementGroupUserState, bool) {
	return amgu.state, amgu.state != nil
}

// StateMust returns the state for [ApiManagementGroupUser]. Panics if the state is nil.
func (amgu *ApiManagementGroupUser) StateMust() *apiManagementGroupUserState {
	if amgu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amgu.Type(), amgu.LocalName()))
	}
	return amgu.state
}

// ApiManagementGroupUserArgs contains the configurations for azurerm_api_management_group_user.
type ApiManagementGroupUserArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// UserId: string, required
	UserId terra.StringValue `hcl:"user_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementgroupuser.Timeouts `hcl:"timeouts,block"`
}
type apiManagementGroupUserAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_group_user.
func (amgu apiManagementGroupUserAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amgu.ref.Append("api_management_name"))
}

// GroupName returns a reference to field group_name of azurerm_api_management_group_user.
func (amgu apiManagementGroupUserAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(amgu.ref.Append("group_name"))
}

// Id returns a reference to field id of azurerm_api_management_group_user.
func (amgu apiManagementGroupUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amgu.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_group_user.
func (amgu apiManagementGroupUserAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amgu.ref.Append("resource_group_name"))
}

// UserId returns a reference to field user_id of azurerm_api_management_group_user.
func (amgu apiManagementGroupUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(amgu.ref.Append("user_id"))
}

func (amgu apiManagementGroupUserAttributes) Timeouts() apimanagementgroupuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementgroupuser.TimeoutsAttributes](amgu.ref.Append("timeouts"))
}

type apiManagementGroupUserState struct {
	ApiManagementName string                                `json:"api_management_name"`
	GroupName         string                                `json:"group_name"`
	Id                string                                `json:"id"`
	ResourceGroupName string                                `json:"resource_group_name"`
	UserId            string                                `json:"user_id"`
	Timeouts          *apimanagementgroupuser.TimeoutsState `json:"timeouts"`
}
