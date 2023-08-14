// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datawebpubsub "github.com/golingon/terraproviders/azurerm/3.69.0/datawebpubsub"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWebPubsub creates a new instance of [DataWebPubsub].
func NewDataWebPubsub(name string, args DataWebPubsubArgs) *DataWebPubsub {
	return &DataWebPubsub{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWebPubsub)(nil)

// DataWebPubsub represents the Terraform data resource azurerm_web_pubsub.
type DataWebPubsub struct {
	Name string
	Args DataWebPubsubArgs
}

// DataSource returns the Terraform object type for [DataWebPubsub].
func (wp *DataWebPubsub) DataSource() string {
	return "azurerm_web_pubsub"
}

// LocalName returns the local name for [DataWebPubsub].
func (wp *DataWebPubsub) LocalName() string {
	return wp.Name
}

// Configuration returns the configuration (args) for [DataWebPubsub].
func (wp *DataWebPubsub) Configuration() interface{} {
	return wp.Args
}

// Attributes returns the attributes for [DataWebPubsub].
func (wp *DataWebPubsub) Attributes() dataWebPubsubAttributes {
	return dataWebPubsubAttributes{ref: terra.ReferenceDataResource(wp)}
}

// DataWebPubsubArgs contains the configurations for azurerm_web_pubsub.
type DataWebPubsubArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datawebpubsub.Timeouts `hcl:"timeouts,block"`
}
type dataWebPubsubAttributes struct {
	ref terra.Reference
}

// AadAuthEnabled returns a reference to field aad_auth_enabled of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) AadAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("aad_auth_enabled"))
}

// Capacity returns a reference to field capacity of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("capacity"))
}

// ExternalIp returns a reference to field external_ip of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) ExternalIp() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("external_ip"))
}

// Hostname returns a reference to field hostname of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("id"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("primary_access_key"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("primary_connection_string"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("public_network_access_enabled"))
}

// PublicPort returns a reference to field public_port of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) PublicPort() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("public_port"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("secondary_access_key"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("secondary_connection_string"))
}

// ServerPort returns a reference to field server_port of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) ServerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(wp.ref.Append("server_port"))
}

// Sku returns a reference to field sku of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wp.ref.Append("tags"))
}

// TlsClientCertEnabled returns a reference to field tls_client_cert_enabled of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) TlsClientCertEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wp.ref.Append("tls_client_cert_enabled"))
}

// Version returns a reference to field version of azurerm_web_pubsub.
func (wp dataWebPubsubAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(wp.ref.Append("version"))
}

func (wp dataWebPubsubAttributes) Timeouts() datawebpubsub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datawebpubsub.TimeoutsAttributes](wp.ref.Append("timeouts"))
}
