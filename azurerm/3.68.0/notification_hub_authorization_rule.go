// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	notificationhubauthorizationrule "github.com/golingon/terraproviders/azurerm/3.68.0/notificationhubauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotificationHubAuthorizationRule creates a new instance of [NotificationHubAuthorizationRule].
func NewNotificationHubAuthorizationRule(name string, args NotificationHubAuthorizationRuleArgs) *NotificationHubAuthorizationRule {
	return &NotificationHubAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotificationHubAuthorizationRule)(nil)

// NotificationHubAuthorizationRule represents the Terraform resource azurerm_notification_hub_authorization_rule.
type NotificationHubAuthorizationRule struct {
	Name      string
	Args      NotificationHubAuthorizationRuleArgs
	state     *notificationHubAuthorizationRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotificationHubAuthorizationRule].
func (nhar *NotificationHubAuthorizationRule) Type() string {
	return "azurerm_notification_hub_authorization_rule"
}

// LocalName returns the local name for [NotificationHubAuthorizationRule].
func (nhar *NotificationHubAuthorizationRule) LocalName() string {
	return nhar.Name
}

// Configuration returns the configuration (args) for [NotificationHubAuthorizationRule].
func (nhar *NotificationHubAuthorizationRule) Configuration() interface{} {
	return nhar.Args
}

// DependOn is used for other resources to depend on [NotificationHubAuthorizationRule].
func (nhar *NotificationHubAuthorizationRule) DependOn() terra.Reference {
	return terra.ReferenceResource(nhar)
}

// Dependencies returns the list of resources [NotificationHubAuthorizationRule] depends_on.
func (nhar *NotificationHubAuthorizationRule) Dependencies() terra.Dependencies {
	return nhar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotificationHubAuthorizationRule].
func (nhar *NotificationHubAuthorizationRule) LifecycleManagement() *terra.Lifecycle {
	return nhar.Lifecycle
}

// Attributes returns the attributes for [NotificationHubAuthorizationRule].
func (nhar *NotificationHubAuthorizationRule) Attributes() notificationHubAuthorizationRuleAttributes {
	return notificationHubAuthorizationRuleAttributes{ref: terra.ReferenceResource(nhar)}
}

// ImportState imports the given attribute values into [NotificationHubAuthorizationRule]'s state.
func (nhar *NotificationHubAuthorizationRule) ImportState(av io.Reader) error {
	nhar.state = &notificationHubAuthorizationRuleState{}
	if err := json.NewDecoder(av).Decode(nhar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nhar.Type(), nhar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotificationHubAuthorizationRule] has state.
func (nhar *NotificationHubAuthorizationRule) State() (*notificationHubAuthorizationRuleState, bool) {
	return nhar.state, nhar.state != nil
}

// StateMust returns the state for [NotificationHubAuthorizationRule]. Panics if the state is nil.
func (nhar *NotificationHubAuthorizationRule) StateMust() *notificationHubAuthorizationRuleState {
	if nhar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nhar.Type(), nhar.LocalName()))
	}
	return nhar.state
}

// NotificationHubAuthorizationRuleArgs contains the configurations for azurerm_notification_hub_authorization_rule.
type NotificationHubAuthorizationRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// NotificationHubName: string, required
	NotificationHubName terra.StringValue `hcl:"notification_hub_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *notificationhubauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type notificationHubAuthorizationRuleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("id"))
}

// Listen returns a reference to field listen of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceAsBool(nhar.ref.Append("listen"))
}

// Manage returns a reference to field manage of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceAsBool(nhar.ref.Append("manage"))
}

// Name returns a reference to field name of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("namespace_name"))
}

// NotificationHubName returns a reference to field notification_hub_name of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) NotificationHubName() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("notification_hub_name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("primary_access_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(nhar.ref.Append("secondary_access_key"))
}

// Send returns a reference to field send of azurerm_notification_hub_authorization_rule.
func (nhar notificationHubAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceAsBool(nhar.ref.Append("send"))
}

func (nhar notificationHubAuthorizationRuleAttributes) Timeouts() notificationhubauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[notificationhubauthorizationrule.TimeoutsAttributes](nhar.ref.Append("timeouts"))
}

type notificationHubAuthorizationRuleState struct {
	Id                  string                                          `json:"id"`
	Listen              bool                                            `json:"listen"`
	Manage              bool                                            `json:"manage"`
	Name                string                                          `json:"name"`
	NamespaceName       string                                          `json:"namespace_name"`
	NotificationHubName string                                          `json:"notification_hub_name"`
	PrimaryAccessKey    string                                          `json:"primary_access_key"`
	ResourceGroupName   string                                          `json:"resource_group_name"`
	SecondaryAccessKey  string                                          `json:"secondary_access_key"`
	Send                bool                                            `json:"send"`
	Timeouts            *notificationhubauthorizationrule.TimeoutsState `json:"timeouts"`
}
