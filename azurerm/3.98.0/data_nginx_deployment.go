// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datanginxdeployment "github.com/golingon/terraproviders/azurerm/3.98.0/datanginxdeployment"
)

// NewDataNginxDeployment creates a new instance of [DataNginxDeployment].
func NewDataNginxDeployment(name string, args DataNginxDeploymentArgs) *DataNginxDeployment {
	return &DataNginxDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNginxDeployment)(nil)

// DataNginxDeployment represents the Terraform data resource azurerm_nginx_deployment.
type DataNginxDeployment struct {
	Name string
	Args DataNginxDeploymentArgs
}

// DataSource returns the Terraform object type for [DataNginxDeployment].
func (nd *DataNginxDeployment) DataSource() string {
	return "azurerm_nginx_deployment"
}

// LocalName returns the local name for [DataNginxDeployment].
func (nd *DataNginxDeployment) LocalName() string {
	return nd.Name
}

// Configuration returns the configuration (args) for [DataNginxDeployment].
func (nd *DataNginxDeployment) Configuration() interface{} {
	return nd.Args
}

// Attributes returns the attributes for [DataNginxDeployment].
func (nd *DataNginxDeployment) Attributes() dataNginxDeploymentAttributes {
	return dataNginxDeploymentAttributes{ref: terra.ReferenceDataResource(nd)}
}

// DataNginxDeploymentArgs contains the configurations for azurerm_nginx_deployment.
type DataNginxDeploymentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AutoScaleProfile: min=0
	AutoScaleProfile []datanginxdeployment.AutoScaleProfile `hcl:"auto_scale_profile,block" validate:"min=0"`
	// FrontendPrivate: min=0
	FrontendPrivate []datanginxdeployment.FrontendPrivate `hcl:"frontend_private,block" validate:"min=0"`
	// FrontendPublic: min=0
	FrontendPublic []datanginxdeployment.FrontendPublic `hcl:"frontend_public,block" validate:"min=0"`
	// Identity: min=0
	Identity []datanginxdeployment.Identity `hcl:"identity,block" validate:"min=0"`
	// LoggingStorageAccount: min=0
	LoggingStorageAccount []datanginxdeployment.LoggingStorageAccount `hcl:"logging_storage_account,block" validate:"min=0"`
	// NetworkInterface: min=0
	NetworkInterface []datanginxdeployment.NetworkInterface `hcl:"network_interface,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanginxdeployment.Timeouts `hcl:"timeouts,block"`
}
type dataNginxDeploymentAttributes struct {
	ref terra.Reference
}

// AutomaticUpgradeChannel returns a reference to field automatic_upgrade_channel of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) AutomaticUpgradeChannel() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("automatic_upgrade_channel"))
}

// Capacity returns a reference to field capacity of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(nd.ref.Append("capacity"))
}

// DiagnoseSupportEnabled returns a reference to field diagnose_support_enabled of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) DiagnoseSupportEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(nd.ref.Append("diagnose_support_enabled"))
}

// Email returns a reference to field email of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("email"))
}

// Id returns a reference to field id of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("ip_address"))
}

// Location returns a reference to field location of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("location"))
}

// ManagedResourceGroup returns a reference to field managed_resource_group of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) ManagedResourceGroup() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("managed_resource_group"))
}

// Name returns a reference to field name of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("name"))
}

// NginxVersion returns a reference to field nginx_version of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) NginxVersion() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("nginx_version"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(nd.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_nginx_deployment.
func (nd dataNginxDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nd.ref.Append("tags"))
}

func (nd dataNginxDeploymentAttributes) AutoScaleProfile() terra.ListValue[datanginxdeployment.AutoScaleProfileAttributes] {
	return terra.ReferenceAsList[datanginxdeployment.AutoScaleProfileAttributes](nd.ref.Append("auto_scale_profile"))
}

func (nd dataNginxDeploymentAttributes) FrontendPrivate() terra.ListValue[datanginxdeployment.FrontendPrivateAttributes] {
	return terra.ReferenceAsList[datanginxdeployment.FrontendPrivateAttributes](nd.ref.Append("frontend_private"))
}

func (nd dataNginxDeploymentAttributes) FrontendPublic() terra.ListValue[datanginxdeployment.FrontendPublicAttributes] {
	return terra.ReferenceAsList[datanginxdeployment.FrontendPublicAttributes](nd.ref.Append("frontend_public"))
}

func (nd dataNginxDeploymentAttributes) Identity() terra.ListValue[datanginxdeployment.IdentityAttributes] {
	return terra.ReferenceAsList[datanginxdeployment.IdentityAttributes](nd.ref.Append("identity"))
}

func (nd dataNginxDeploymentAttributes) LoggingStorageAccount() terra.ListValue[datanginxdeployment.LoggingStorageAccountAttributes] {
	return terra.ReferenceAsList[datanginxdeployment.LoggingStorageAccountAttributes](nd.ref.Append("logging_storage_account"))
}

func (nd dataNginxDeploymentAttributes) NetworkInterface() terra.ListValue[datanginxdeployment.NetworkInterfaceAttributes] {
	return terra.ReferenceAsList[datanginxdeployment.NetworkInterfaceAttributes](nd.ref.Append("network_interface"))
}

func (nd dataNginxDeploymentAttributes) Timeouts() datanginxdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanginxdeployment.TimeoutsAttributes](nd.ref.Append("timeouts"))
}
