// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscription "github.com/golingon/terraproviders/azurerm/3.55.0/subscription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubscription creates a new instance of [Subscription].
func NewSubscription(name string, args SubscriptionArgs) *Subscription {
	return &Subscription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Subscription)(nil)

// Subscription represents the Terraform resource azurerm_subscription.
type Subscription struct {
	Name      string
	Args      SubscriptionArgs
	state     *subscriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Subscription].
func (s *Subscription) Type() string {
	return "azurerm_subscription"
}

// LocalName returns the local name for [Subscription].
func (s *Subscription) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [Subscription].
func (s *Subscription) Configuration() interface{} {
	return s.Args
}

// DependOn is used for other resources to depend on [Subscription].
func (s *Subscription) DependOn() terra.Reference {
	return terra.ReferenceResource(s)
}

// Dependencies returns the list of resources [Subscription] depends_on.
func (s *Subscription) Dependencies() terra.Dependencies {
	return s.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Subscription].
func (s *Subscription) LifecycleManagement() *terra.Lifecycle {
	return s.Lifecycle
}

// Attributes returns the attributes for [Subscription].
func (s *Subscription) Attributes() subscriptionAttributes {
	return subscriptionAttributes{ref: terra.ReferenceResource(s)}
}

// ImportState imports the given attribute values into [Subscription]'s state.
func (s *Subscription) ImportState(av io.Reader) error {
	s.state = &subscriptionState{}
	if err := json.NewDecoder(av).Decode(s.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", s.Type(), s.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Subscription] has state.
func (s *Subscription) State() (*subscriptionState, bool) {
	return s.state, s.state != nil
}

// StateMust returns the state for [Subscription]. Panics if the state is nil.
func (s *Subscription) StateMust() *subscriptionState {
	if s.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", s.Type(), s.LocalName()))
	}
	return s.state
}

// SubscriptionArgs contains the configurations for azurerm_subscription.
type SubscriptionArgs struct {
	// Alias: string, optional
	Alias terra.StringValue `hcl:"alias,attr"`
	// BillingScopeId: string, optional
	BillingScopeId terra.StringValue `hcl:"billing_scope_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// SubscriptionName: string, required
	SubscriptionName terra.StringValue `hcl:"subscription_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Workload: string, optional
	Workload terra.StringValue `hcl:"workload,attr"`
	// Timeouts: optional
	Timeouts *subscription.Timeouts `hcl:"timeouts,block"`
}
type subscriptionAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of azurerm_subscription.
func (s subscriptionAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("alias"))
}

// BillingScopeId returns a reference to field billing_scope_id of azurerm_subscription.
func (s subscriptionAttributes) BillingScopeId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("billing_scope_id"))
}

// Id returns a reference to field id of azurerm_subscription.
func (s subscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_subscription.
func (s subscriptionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("subscription_id"))
}

// SubscriptionName returns a reference to field subscription_name of azurerm_subscription.
func (s subscriptionAttributes) SubscriptionName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("subscription_name"))
}

// Tags returns a reference to field tags of azurerm_subscription.
func (s subscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_subscription.
func (s subscriptionAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tenant_id"))
}

// Workload returns a reference to field workload of azurerm_subscription.
func (s subscriptionAttributes) Workload() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("workload"))
}

func (s subscriptionAttributes) Timeouts() subscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subscription.TimeoutsAttributes](s.ref.Append("timeouts"))
}

type subscriptionState struct {
	Alias            string                      `json:"alias"`
	BillingScopeId   string                      `json:"billing_scope_id"`
	Id               string                      `json:"id"`
	SubscriptionId   string                      `json:"subscription_id"`
	SubscriptionName string                      `json:"subscription_name"`
	Tags             map[string]string           `json:"tags"`
	TenantId         string                      `json:"tenant_id"`
	Workload         string                      `json:"workload"`
	Timeouts         *subscription.TimeoutsState `json:"timeouts"`
}
