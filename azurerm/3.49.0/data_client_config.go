// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataclientconfig "github.com/golingon/terraproviders/azurerm/3.49.0/dataclientconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataClientConfig creates a new instance of [DataClientConfig].
func NewDataClientConfig(name string, args DataClientConfigArgs) *DataClientConfig {
	return &DataClientConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataClientConfig)(nil)

// DataClientConfig represents the Terraform data resource azurerm_client_config.
type DataClientConfig struct {
	Name string
	Args DataClientConfigArgs
}

// DataSource returns the Terraform object type for [DataClientConfig].
func (cc *DataClientConfig) DataSource() string {
	return "azurerm_client_config"
}

// LocalName returns the local name for [DataClientConfig].
func (cc *DataClientConfig) LocalName() string {
	return cc.Name
}

// Configuration returns the configuration (args) for [DataClientConfig].
func (cc *DataClientConfig) Configuration() interface{} {
	return cc.Args
}

// Attributes returns the attributes for [DataClientConfig].
func (cc *DataClientConfig) Attributes() dataClientConfigAttributes {
	return dataClientConfigAttributes{ref: terra.ReferenceDataResource(cc)}
}

// DataClientConfigArgs contains the configurations for azurerm_client_config.
type DataClientConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *dataclientconfig.Timeouts `hcl:"timeouts,block"`
}
type dataClientConfigAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of azurerm_client_config.
func (cc dataClientConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("client_id"))
}

// Id returns a reference to field id of azurerm_client_config.
func (cc dataClientConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("id"))
}

// ObjectId returns a reference to field object_id of azurerm_client_config.
func (cc dataClientConfigAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("object_id"))
}

// SubscriptionId returns a reference to field subscription_id of azurerm_client_config.
func (cc dataClientConfigAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("subscription_id"))
}

// TenantId returns a reference to field tenant_id of azurerm_client_config.
func (cc dataClientConfigAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("tenant_id"))
}

func (cc dataClientConfigAttributes) Timeouts() dataclientconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataclientconfig.TimeoutsAttributes](cc.ref.Append("timeouts"))
}
