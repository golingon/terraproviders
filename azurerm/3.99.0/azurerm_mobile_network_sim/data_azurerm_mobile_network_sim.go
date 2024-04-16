// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mobile_network_sim

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_mobile_network_sim.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (amns *DataSource) DataSource() string {
	return "azurerm_mobile_network_sim"
}

// LocalName returns the local name for [DataSource].
func (amns *DataSource) LocalName() string {
	return amns.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (amns *DataSource) Configuration() interface{} {
	return amns.Args
}

// Attributes returns the attributes for [DataSource].
func (amns *DataSource) Attributes() dataAzurermMobileNetworkSimAttributes {
	return dataAzurermMobileNetworkSimAttributes{ref: terra.ReferenceDataSource(amns)}
}

// DataArgs contains the configurations for azurerm_mobile_network_sim.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MobileNetworkSimGroupId: string, required
	MobileNetworkSimGroupId terra.StringValue `hcl:"mobile_network_sim_group_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermMobileNetworkSimAttributes struct {
	ref terra.Reference
}

// DeviceType returns a reference to field device_type of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) DeviceType() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("device_type"))
}

// Id returns a reference to field id of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("id"))
}

// IntegratedCircuitCardIdentifier returns a reference to field integrated_circuit_card_identifier of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) IntegratedCircuitCardIdentifier() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("integrated_circuit_card_identifier"))
}

// InternationalMobileSubscriberIdentity returns a reference to field international_mobile_subscriber_identity of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) InternationalMobileSubscriberIdentity() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("international_mobile_subscriber_identity"))
}

// MobileNetworkSimGroupId returns a reference to field mobile_network_sim_group_id of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) MobileNetworkSimGroupId() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("mobile_network_sim_group_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("name"))
}

// SimPolicyId returns a reference to field sim_policy_id of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) SimPolicyId() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("sim_policy_id"))
}

// SimState returns a reference to field sim_state of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) SimState() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("sim_state"))
}

// VendorKeyFingerprint returns a reference to field vendor_key_fingerprint of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) VendorKeyFingerprint() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("vendor_key_fingerprint"))
}

// VendorName returns a reference to field vendor_name of azurerm_mobile_network_sim.
func (amns dataAzurermMobileNetworkSimAttributes) VendorName() terra.StringValue {
	return terra.ReferenceAsString(amns.ref.Append("vendor_name"))
}

func (amns dataAzurermMobileNetworkSimAttributes) StaticIpConfiguration() terra.ListValue[DataStaticIpConfigurationAttributes] {
	return terra.ReferenceAsList[DataStaticIpConfigurationAttributes](amns.ref.Append("static_ip_configuration"))
}

func (amns dataAzurermMobileNetworkSimAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](amns.ref.Append("timeouts"))
}
