// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicebussubscription "github.com/golingon/terraproviders/azurerm/3.58.0/servicebussubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicebusSubscription creates a new instance of [ServicebusSubscription].
func NewServicebusSubscription(name string, args ServicebusSubscriptionArgs) *ServicebusSubscription {
	return &ServicebusSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicebusSubscription)(nil)

// ServicebusSubscription represents the Terraform resource azurerm_servicebus_subscription.
type ServicebusSubscription struct {
	Name      string
	Args      ServicebusSubscriptionArgs
	state     *servicebusSubscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicebusSubscription].
func (ss *ServicebusSubscription) Type() string {
	return "azurerm_servicebus_subscription"
}

// LocalName returns the local name for [ServicebusSubscription].
func (ss *ServicebusSubscription) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [ServicebusSubscription].
func (ss *ServicebusSubscription) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [ServicebusSubscription].
func (ss *ServicebusSubscription) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [ServicebusSubscription] depends_on.
func (ss *ServicebusSubscription) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicebusSubscription].
func (ss *ServicebusSubscription) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [ServicebusSubscription].
func (ss *ServicebusSubscription) Attributes() servicebusSubscriptionAttributes {
	return servicebusSubscriptionAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [ServicebusSubscription]'s state.
func (ss *ServicebusSubscription) ImportState(av io.Reader) error {
	ss.state = &servicebusSubscriptionState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicebusSubscription] has state.
func (ss *ServicebusSubscription) State() (*servicebusSubscriptionState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [ServicebusSubscription]. Panics if the state is nil.
func (ss *ServicebusSubscription) StateMust() *servicebusSubscriptionState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// ServicebusSubscriptionArgs contains the configurations for azurerm_servicebus_subscription.
type ServicebusSubscriptionArgs struct {
	// AutoDeleteOnIdle: string, optional
	AutoDeleteOnIdle terra.StringValue `hcl:"auto_delete_on_idle,attr"`
	// ClientScopedSubscriptionEnabled: bool, optional
	ClientScopedSubscriptionEnabled terra.BoolValue `hcl:"client_scoped_subscription_enabled,attr"`
	// DeadLetteringOnFilterEvaluationError: bool, optional
	DeadLetteringOnFilterEvaluationError terra.BoolValue `hcl:"dead_lettering_on_filter_evaluation_error,attr"`
	// DeadLetteringOnMessageExpiration: bool, optional
	DeadLetteringOnMessageExpiration terra.BoolValue `hcl:"dead_lettering_on_message_expiration,attr"`
	// DefaultMessageTtl: string, optional
	DefaultMessageTtl terra.StringValue `hcl:"default_message_ttl,attr"`
	// EnableBatchedOperations: bool, optional
	EnableBatchedOperations terra.BoolValue `hcl:"enable_batched_operations,attr"`
	// ForwardDeadLetteredMessagesTo: string, optional
	ForwardDeadLetteredMessagesTo terra.StringValue `hcl:"forward_dead_lettered_messages_to,attr"`
	// ForwardTo: string, optional
	ForwardTo terra.StringValue `hcl:"forward_to,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LockDuration: string, optional
	LockDuration terra.StringValue `hcl:"lock_duration,attr"`
	// MaxDeliveryCount: number, required
	MaxDeliveryCount terra.NumberValue `hcl:"max_delivery_count,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequiresSession: bool, optional
	RequiresSession terra.BoolValue `hcl:"requires_session,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// TopicId: string, required
	TopicId terra.StringValue `hcl:"topic_id,attr" validate:"required"`
	// ClientScopedSubscription: optional
	ClientScopedSubscription *servicebussubscription.ClientScopedSubscription `hcl:"client_scoped_subscription,block"`
	// Timeouts: optional
	Timeouts *servicebussubscription.Timeouts `hcl:"timeouts,block"`
}
type servicebusSubscriptionAttributes struct {
	ref terra.Reference
}

// AutoDeleteOnIdle returns a reference to field auto_delete_on_idle of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) AutoDeleteOnIdle() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("auto_delete_on_idle"))
}

// ClientScopedSubscriptionEnabled returns a reference to field client_scoped_subscription_enabled of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) ClientScopedSubscriptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("client_scoped_subscription_enabled"))
}

// DeadLetteringOnFilterEvaluationError returns a reference to field dead_lettering_on_filter_evaluation_error of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) DeadLetteringOnFilterEvaluationError() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("dead_lettering_on_filter_evaluation_error"))
}

// DeadLetteringOnMessageExpiration returns a reference to field dead_lettering_on_message_expiration of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) DeadLetteringOnMessageExpiration() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("dead_lettering_on_message_expiration"))
}

// DefaultMessageTtl returns a reference to field default_message_ttl of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) DefaultMessageTtl() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("default_message_ttl"))
}

// EnableBatchedOperations returns a reference to field enable_batched_operations of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) EnableBatchedOperations() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("enable_batched_operations"))
}

// ForwardDeadLetteredMessagesTo returns a reference to field forward_dead_lettered_messages_to of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) ForwardDeadLetteredMessagesTo() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("forward_dead_lettered_messages_to"))
}

// ForwardTo returns a reference to field forward_to of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) ForwardTo() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("forward_to"))
}

// Id returns a reference to field id of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// LockDuration returns a reference to field lock_duration of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) LockDuration() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("lock_duration"))
}

// MaxDeliveryCount returns a reference to field max_delivery_count of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) MaxDeliveryCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("max_delivery_count"))
}

// Name returns a reference to field name of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// RequiresSession returns a reference to field requires_session of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) RequiresSession() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("requires_session"))
}

// Status returns a reference to field status of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("status"))
}

// TopicId returns a reference to field topic_id of azurerm_servicebus_subscription.
func (ss servicebusSubscriptionAttributes) TopicId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("topic_id"))
}

func (ss servicebusSubscriptionAttributes) ClientScopedSubscription() terra.ListValue[servicebussubscription.ClientScopedSubscriptionAttributes] {
	return terra.ReferenceAsList[servicebussubscription.ClientScopedSubscriptionAttributes](ss.ref.Append("client_scoped_subscription"))
}

func (ss servicebusSubscriptionAttributes) Timeouts() servicebussubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicebussubscription.TimeoutsAttributes](ss.ref.Append("timeouts"))
}

type servicebusSubscriptionState struct {
	AutoDeleteOnIdle                     string                                                 `json:"auto_delete_on_idle"`
	ClientScopedSubscriptionEnabled      bool                                                   `json:"client_scoped_subscription_enabled"`
	DeadLetteringOnFilterEvaluationError bool                                                   `json:"dead_lettering_on_filter_evaluation_error"`
	DeadLetteringOnMessageExpiration     bool                                                   `json:"dead_lettering_on_message_expiration"`
	DefaultMessageTtl                    string                                                 `json:"default_message_ttl"`
	EnableBatchedOperations              bool                                                   `json:"enable_batched_operations"`
	ForwardDeadLetteredMessagesTo        string                                                 `json:"forward_dead_lettered_messages_to"`
	ForwardTo                            string                                                 `json:"forward_to"`
	Id                                   string                                                 `json:"id"`
	LockDuration                         string                                                 `json:"lock_duration"`
	MaxDeliveryCount                     float64                                                `json:"max_delivery_count"`
	Name                                 string                                                 `json:"name"`
	RequiresSession                      bool                                                   `json:"requires_session"`
	Status                               string                                                 `json:"status"`
	TopicId                              string                                                 `json:"topic_id"`
	ClientScopedSubscription             []servicebussubscription.ClientScopedSubscriptionState `json:"client_scoped_subscription"`
	Timeouts                             *servicebussubscription.TimeoutsState                  `json:"timeouts"`
}
