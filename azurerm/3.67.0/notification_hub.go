// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	notificationhub "github.com/golingon/terraproviders/azurerm/3.67.0/notificationhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotificationHub creates a new instance of [NotificationHub].
func NewNotificationHub(name string, args NotificationHubArgs) *NotificationHub {
	return &NotificationHub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotificationHub)(nil)

// NotificationHub represents the Terraform resource azurerm_notification_hub.
type NotificationHub struct {
	Name      string
	Args      NotificationHubArgs
	state     *notificationHubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotificationHub].
func (nh *NotificationHub) Type() string {
	return "azurerm_notification_hub"
}

// LocalName returns the local name for [NotificationHub].
func (nh *NotificationHub) LocalName() string {
	return nh.Name
}

// Configuration returns the configuration (args) for [NotificationHub].
func (nh *NotificationHub) Configuration() interface{} {
	return nh.Args
}

// DependOn is used for other resources to depend on [NotificationHub].
func (nh *NotificationHub) DependOn() terra.Reference {
	return terra.ReferenceResource(nh)
}

// Dependencies returns the list of resources [NotificationHub] depends_on.
func (nh *NotificationHub) Dependencies() terra.Dependencies {
	return nh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotificationHub].
func (nh *NotificationHub) LifecycleManagement() *terra.Lifecycle {
	return nh.Lifecycle
}

// Attributes returns the attributes for [NotificationHub].
func (nh *NotificationHub) Attributes() notificationHubAttributes {
	return notificationHubAttributes{ref: terra.ReferenceResource(nh)}
}

// ImportState imports the given attribute values into [NotificationHub]'s state.
func (nh *NotificationHub) ImportState(av io.Reader) error {
	nh.state = &notificationHubState{}
	if err := json.NewDecoder(av).Decode(nh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nh.Type(), nh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotificationHub] has state.
func (nh *NotificationHub) State() (*notificationHubState, bool) {
	return nh.state, nh.state != nil
}

// StateMust returns the state for [NotificationHub]. Panics if the state is nil.
func (nh *NotificationHub) StateMust() *notificationHubState {
	if nh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nh.Type(), nh.LocalName()))
	}
	return nh.state
}

// NotificationHubArgs contains the configurations for azurerm_notification_hub.
type NotificationHubArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ApnsCredential: optional
	ApnsCredential *notificationhub.ApnsCredential `hcl:"apns_credential,block"`
	// GcmCredential: optional
	GcmCredential *notificationhub.GcmCredential `hcl:"gcm_credential,block"`
	// Timeouts: optional
	Timeouts *notificationhub.Timeouts `hcl:"timeouts,block"`
}
type notificationHubAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_notification_hub.
func (nh notificationHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_notification_hub.
func (nh notificationHubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_notification_hub.
func (nh notificationHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_notification_hub.
func (nh notificationHubAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("namespace_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_notification_hub.
func (nh notificationHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_notification_hub.
func (nh notificationHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nh.ref.Append("tags"))
}

func (nh notificationHubAttributes) ApnsCredential() terra.ListValue[notificationhub.ApnsCredentialAttributes] {
	return terra.ReferenceAsList[notificationhub.ApnsCredentialAttributes](nh.ref.Append("apns_credential"))
}

func (nh notificationHubAttributes) GcmCredential() terra.ListValue[notificationhub.GcmCredentialAttributes] {
	return terra.ReferenceAsList[notificationhub.GcmCredentialAttributes](nh.ref.Append("gcm_credential"))
}

func (nh notificationHubAttributes) Timeouts() notificationhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[notificationhub.TimeoutsAttributes](nh.ref.Append("timeouts"))
}

type notificationHubState struct {
	Id                string                                `json:"id"`
	Location          string                                `json:"location"`
	Name              string                                `json:"name"`
	NamespaceName     string                                `json:"namespace_name"`
	ResourceGroupName string                                `json:"resource_group_name"`
	Tags              map[string]string                     `json:"tags"`
	ApnsCredential    []notificationhub.ApnsCredentialState `json:"apns_credential"`
	GcmCredential     []notificationhub.GcmCredentialState  `json:"gcm_credential"`
	Timeouts          *notificationhub.TimeoutsState        `json:"timeouts"`
}
