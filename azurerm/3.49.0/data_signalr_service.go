// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasignalrservice "github.com/golingon/terraproviders/azurerm/3.49.0/datasignalrservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSignalrService creates a new instance of [DataSignalrService].
func NewDataSignalrService(name string, args DataSignalrServiceArgs) *DataSignalrService {
	return &DataSignalrService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSignalrService)(nil)

// DataSignalrService represents the Terraform data resource azurerm_signalr_service.
type DataSignalrService struct {
	Name string
	Args DataSignalrServiceArgs
}

// DataSource returns the Terraform object type for [DataSignalrService].
func (ss *DataSignalrService) DataSource() string {
	return "azurerm_signalr_service"
}

// LocalName returns the local name for [DataSignalrService].
func (ss *DataSignalrService) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [DataSignalrService].
func (ss *DataSignalrService) Configuration() interface{} {
	return ss.Args
}

// Attributes returns the attributes for [DataSignalrService].
func (ss *DataSignalrService) Attributes() dataSignalrServiceAttributes {
	return dataSignalrServiceAttributes{ref: terra.ReferenceDataResource(ss)}
}

// DataSignalrServiceArgs contains the configurations for azurerm_signalr_service.
type DataSignalrServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasignalrservice.Timeouts `hcl:"timeouts,block"`
}
type dataSignalrServiceAttributes struct {
	ref terra.Reference
}

// AadAuthEnabled returns a reference to field aad_auth_enabled of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) AadAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("aad_auth_enabled"))
}

// Hostname returns a reference to field hostname of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) Hostname() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("hostname"))
}

// Id returns a reference to field id of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("ip_address"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("primary_access_key"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("primary_connection_string"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("public_network_access_enabled"))
}

// PublicPort returns a reference to field public_port of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) PublicPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("public_port"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("secondary_access_key"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("secondary_connection_string"))
}

// ServerPort returns a reference to field server_port of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) ServerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("server_port"))
}

// ServerlessConnectionTimeoutInSeconds returns a reference to field serverless_connection_timeout_in_seconds of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) ServerlessConnectionTimeoutInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("serverless_connection_timeout_in_seconds"))
}

// Tags returns a reference to field tags of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

// TlsClientCertEnabled returns a reference to field tls_client_cert_enabled of azurerm_signalr_service.
func (ss dataSignalrServiceAttributes) TlsClientCertEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("tls_client_cert_enabled"))
}

func (ss dataSignalrServiceAttributes) Timeouts() datasignalrservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasignalrservice.TimeoutsAttributes](ss.ref.Append("timeouts"))
}
