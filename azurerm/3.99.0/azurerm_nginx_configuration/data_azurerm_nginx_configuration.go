// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_nginx_configuration

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_nginx_configuration.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (anc *DataSource) DataSource() string {
	return "azurerm_nginx_configuration"
}

// LocalName returns the local name for [DataSource].
func (anc *DataSource) LocalName() string {
	return anc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (anc *DataSource) Configuration() interface{} {
	return anc.Args
}

// Attributes returns the attributes for [DataSource].
func (anc *DataSource) Attributes() dataAzurermNginxConfigurationAttributes {
	return dataAzurermNginxConfigurationAttributes{ref: terra.ReferenceDataSource(anc)}
}

// DataArgs contains the configurations for azurerm_nginx_configuration.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NginxDeploymentId: string, required
	NginxDeploymentId terra.StringValue `hcl:"nginx_deployment_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermNginxConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_nginx_configuration.
func (anc dataAzurermNginxConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("id"))
}

// NginxDeploymentId returns a reference to field nginx_deployment_id of azurerm_nginx_configuration.
func (anc dataAzurermNginxConfigurationAttributes) NginxDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("nginx_deployment_id"))
}

// PackageData returns a reference to field package_data of azurerm_nginx_configuration.
func (anc dataAzurermNginxConfigurationAttributes) PackageData() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("package_data"))
}

// RootFile returns a reference to field root_file of azurerm_nginx_configuration.
func (anc dataAzurermNginxConfigurationAttributes) RootFile() terra.StringValue {
	return terra.ReferenceAsString(anc.ref.Append("root_file"))
}

func (anc dataAzurermNginxConfigurationAttributes) ConfigFile() terra.SetValue[DataConfigFileAttributes] {
	return terra.ReferenceAsSet[DataConfigFileAttributes](anc.ref.Append("config_file"))
}

func (anc dataAzurermNginxConfigurationAttributes) ProtectedFile() terra.SetValue[DataProtectedFileAttributes] {
	return terra.ReferenceAsSet[DataProtectedFileAttributes](anc.ref.Append("protected_file"))
}

func (anc dataAzurermNginxConfigurationAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](anc.ref.Append("timeouts"))
}
