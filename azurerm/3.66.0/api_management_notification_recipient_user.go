// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementnotificationrecipientuser "github.com/golingon/terraproviders/azurerm/3.66.0/apimanagementnotificationrecipientuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementNotificationRecipientUser creates a new instance of [ApiManagementNotificationRecipientUser].
func NewApiManagementNotificationRecipientUser(name string, args ApiManagementNotificationRecipientUserArgs) *ApiManagementNotificationRecipientUser {
	return &ApiManagementNotificationRecipientUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementNotificationRecipientUser)(nil)

// ApiManagementNotificationRecipientUser represents the Terraform resource azurerm_api_management_notification_recipient_user.
type ApiManagementNotificationRecipientUser struct {
	Name      string
	Args      ApiManagementNotificationRecipientUserArgs
	state     *apiManagementNotificationRecipientUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementNotificationRecipientUser].
func (amnru *ApiManagementNotificationRecipientUser) Type() string {
	return "azurerm_api_management_notification_recipient_user"
}

// LocalName returns the local name for [ApiManagementNotificationRecipientUser].
func (amnru *ApiManagementNotificationRecipientUser) LocalName() string {
	return amnru.Name
}

// Configuration returns the configuration (args) for [ApiManagementNotificationRecipientUser].
func (amnru *ApiManagementNotificationRecipientUser) Configuration() interface{} {
	return amnru.Args
}

// DependOn is used for other resources to depend on [ApiManagementNotificationRecipientUser].
func (amnru *ApiManagementNotificationRecipientUser) DependOn() terra.Reference {
	return terra.ReferenceResource(amnru)
}

// Dependencies returns the list of resources [ApiManagementNotificationRecipientUser] depends_on.
func (amnru *ApiManagementNotificationRecipientUser) Dependencies() terra.Dependencies {
	return amnru.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementNotificationRecipientUser].
func (amnru *ApiManagementNotificationRecipientUser) LifecycleManagement() *terra.Lifecycle {
	return amnru.Lifecycle
}

// Attributes returns the attributes for [ApiManagementNotificationRecipientUser].
func (amnru *ApiManagementNotificationRecipientUser) Attributes() apiManagementNotificationRecipientUserAttributes {
	return apiManagementNotificationRecipientUserAttributes{ref: terra.ReferenceResource(amnru)}
}

// ImportState imports the given attribute values into [ApiManagementNotificationRecipientUser]'s state.
func (amnru *ApiManagementNotificationRecipientUser) ImportState(av io.Reader) error {
	amnru.state = &apiManagementNotificationRecipientUserState{}
	if err := json.NewDecoder(av).Decode(amnru.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amnru.Type(), amnru.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementNotificationRecipientUser] has state.
func (amnru *ApiManagementNotificationRecipientUser) State() (*apiManagementNotificationRecipientUserState, bool) {
	return amnru.state, amnru.state != nil
}

// StateMust returns the state for [ApiManagementNotificationRecipientUser]. Panics if the state is nil.
func (amnru *ApiManagementNotificationRecipientUser) StateMust() *apiManagementNotificationRecipientUserState {
	if amnru.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amnru.Type(), amnru.LocalName()))
	}
	return amnru.state
}

// ApiManagementNotificationRecipientUserArgs contains the configurations for azurerm_api_management_notification_recipient_user.
type ApiManagementNotificationRecipientUserArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationType: string, required
	NotificationType terra.StringValue `hcl:"notification_type,attr" validate:"required"`
	// UserId: string, required
	UserId terra.StringValue `hcl:"user_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementnotificationrecipientuser.Timeouts `hcl:"timeouts,block"`
}
type apiManagementNotificationRecipientUserAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_notification_recipient_user.
func (amnru apiManagementNotificationRecipientUserAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amnru.ref.Append("api_management_id"))
}

// Id returns a reference to field id of azurerm_api_management_notification_recipient_user.
func (amnru apiManagementNotificationRecipientUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amnru.ref.Append("id"))
}

// NotificationType returns a reference to field notification_type of azurerm_api_management_notification_recipient_user.
func (amnru apiManagementNotificationRecipientUserAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceAsString(amnru.ref.Append("notification_type"))
}

// UserId returns a reference to field user_id of azurerm_api_management_notification_recipient_user.
func (amnru apiManagementNotificationRecipientUserAttributes) UserId() terra.StringValue {
	return terra.ReferenceAsString(amnru.ref.Append("user_id"))
}

func (amnru apiManagementNotificationRecipientUserAttributes) Timeouts() apimanagementnotificationrecipientuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementnotificationrecipientuser.TimeoutsAttributes](amnru.ref.Append("timeouts"))
}

type apiManagementNotificationRecipientUserState struct {
	ApiManagementId  string                                                `json:"api_management_id"`
	Id               string                                                `json:"id"`
	NotificationType string                                                `json:"notification_type"`
	UserId           string                                                `json:"user_id"`
	Timeouts         *apimanagementnotificationrecipientuser.TimeoutsState `json:"timeouts"`
}
