// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasubscription "github.com/golingon/terraproviders/azurerm/3.49.0/datasubscription"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSubscription creates a new instance of [DataSubscription].
func NewDataSubscription(name string, args DataSubscriptionArgs) *DataSubscription {
	return &DataSubscription{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSubscription)(nil)

// DataSubscription represents the Terraform data resource azurerm_subscription.
type DataSubscription struct {
	Name string
	Args DataSubscriptionArgs
}

// DataSource returns the Terraform object type for [DataSubscription].
func (s *DataSubscription) DataSource() string {
	return "azurerm_subscription"
}

// LocalName returns the local name for [DataSubscription].
func (s *DataSubscription) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [DataSubscription].
func (s *DataSubscription) Configuration() interface{} {
	return s.Args
}

// Attributes returns the attributes for [DataSubscription].
func (s *DataSubscription) Attributes() dataSubscriptionAttributes {
	return dataSubscriptionAttributes{ref: terra.ReferenceDataResource(s)}
}

// DataSubscriptionArgs contains the configurations for azurerm_subscription.
type DataSubscriptionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubscriptionId: string, optional
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr"`
	// Timeouts: optional
	Timeouts *datasubscription.Timeouts `hcl:"timeouts,block"`
}
type dataSubscriptionAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_subscription.
func (s dataSubscriptionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_subscription.
func (s dataSubscriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// LocationPlacementId returns a reference to field location_placement_id of azurerm_subscription.
func (s dataSubscriptionAttributes) LocationPlacementId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("location_placement_id"))
}

// QuotaId returns a reference to field quota_id of azurerm_subscription.
func (s dataSubscriptionAttributes) QuotaId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("quota_id"))
}

// SpendingLimit returns a reference to field spending_limit of azurerm_subscription.
func (s dataSubscriptionAttributes) SpendingLimit() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("spending_limit"))
}

// State returns a reference to field state of azurerm_subscription.
func (s dataSubscriptionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("state"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_subscription.
func (s dataSubscriptionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("subscription_id"))
}

// Tags returns a reference to field tags of azurerm_subscription.
func (s dataSubscriptionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_subscription.
func (s dataSubscriptionAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("tenant_id"))
}

func (s dataSubscriptionAttributes) Timeouts() datasubscription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasubscription.TimeoutsAttributes](s.ref.Append("timeouts"))
}
