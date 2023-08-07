// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementnotificationrecipientemail "github.com/golingon/terraproviders/azurerm/3.68.0/apimanagementnotificationrecipientemail"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementNotificationRecipientEmail creates a new instance of [ApiManagementNotificationRecipientEmail].
func NewApiManagementNotificationRecipientEmail(name string, args ApiManagementNotificationRecipientEmailArgs) *ApiManagementNotificationRecipientEmail {
	return &ApiManagementNotificationRecipientEmail{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementNotificationRecipientEmail)(nil)

// ApiManagementNotificationRecipientEmail represents the Terraform resource azurerm_api_management_notification_recipient_email.
type ApiManagementNotificationRecipientEmail struct {
	Name      string
	Args      ApiManagementNotificationRecipientEmailArgs
	state     *apiManagementNotificationRecipientEmailState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementNotificationRecipientEmail].
func (amnre *ApiManagementNotificationRecipientEmail) Type() string {
	return "azurerm_api_management_notification_recipient_email"
}

// LocalName returns the local name for [ApiManagementNotificationRecipientEmail].
func (amnre *ApiManagementNotificationRecipientEmail) LocalName() string {
	return amnre.Name
}

// Configuration returns the configuration (args) for [ApiManagementNotificationRecipientEmail].
func (amnre *ApiManagementNotificationRecipientEmail) Configuration() interface{} {
	return amnre.Args
}

// DependOn is used for other resources to depend on [ApiManagementNotificationRecipientEmail].
func (amnre *ApiManagementNotificationRecipientEmail) DependOn() terra.Reference {
	return terra.ReferenceResource(amnre)
}

// Dependencies returns the list of resources [ApiManagementNotificationRecipientEmail] depends_on.
func (amnre *ApiManagementNotificationRecipientEmail) Dependencies() terra.Dependencies {
	return amnre.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementNotificationRecipientEmail].
func (amnre *ApiManagementNotificationRecipientEmail) LifecycleManagement() *terra.Lifecycle {
	return amnre.Lifecycle
}

// Attributes returns the attributes for [ApiManagementNotificationRecipientEmail].
func (amnre *ApiManagementNotificationRecipientEmail) Attributes() apiManagementNotificationRecipientEmailAttributes {
	return apiManagementNotificationRecipientEmailAttributes{ref: terra.ReferenceResource(amnre)}
}

// ImportState imports the given attribute values into [ApiManagementNotificationRecipientEmail]'s state.
func (amnre *ApiManagementNotificationRecipientEmail) ImportState(av io.Reader) error {
	amnre.state = &apiManagementNotificationRecipientEmailState{}
	if err := json.NewDecoder(av).Decode(amnre.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amnre.Type(), amnre.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementNotificationRecipientEmail] has state.
func (amnre *ApiManagementNotificationRecipientEmail) State() (*apiManagementNotificationRecipientEmailState, bool) {
	return amnre.state, amnre.state != nil
}

// StateMust returns the state for [ApiManagementNotificationRecipientEmail]. Panics if the state is nil.
func (amnre *ApiManagementNotificationRecipientEmail) StateMust() *apiManagementNotificationRecipientEmailState {
	if amnre.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amnre.Type(), amnre.LocalName()))
	}
	return amnre.state
}

// ApiManagementNotificationRecipientEmailArgs contains the configurations for azurerm_api_management_notification_recipient_email.
type ApiManagementNotificationRecipientEmailArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NotificationType: string, required
	NotificationType terra.StringValue `hcl:"notification_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementnotificationrecipientemail.Timeouts `hcl:"timeouts,block"`
}
type apiManagementNotificationRecipientEmailAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_notification_recipient_email.
func (amnre apiManagementNotificationRecipientEmailAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amnre.ref.Append("api_management_id"))
}

// Email returns a reference to field email of azurerm_api_management_notification_recipient_email.
func (amnre apiManagementNotificationRecipientEmailAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(amnre.ref.Append("email"))
}

// Id returns a reference to field id of azurerm_api_management_notification_recipient_email.
func (amnre apiManagementNotificationRecipientEmailAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amnre.ref.Append("id"))
}

// NotificationType returns a reference to field notification_type of azurerm_api_management_notification_recipient_email.
func (amnre apiManagementNotificationRecipientEmailAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceAsString(amnre.ref.Append("notification_type"))
}

func (amnre apiManagementNotificationRecipientEmailAttributes) Timeouts() apimanagementnotificationrecipientemail.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementnotificationrecipientemail.TimeoutsAttributes](amnre.ref.Append("timeouts"))
}

type apiManagementNotificationRecipientEmailState struct {
	ApiManagementId  string                                                 `json:"api_management_id"`
	Email            string                                                 `json:"email"`
	Id               string                                                 `json:"id"`
	NotificationType string                                                 `json:"notification_type"`
	Timeouts         *apimanagementnotificationrecipientemail.TimeoutsState `json:"timeouts"`
}
