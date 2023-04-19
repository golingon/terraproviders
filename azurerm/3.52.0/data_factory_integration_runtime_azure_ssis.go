// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactoryintegrationruntimeazuressis "github.com/golingon/terraproviders/azurerm/3.52.0/datafactoryintegrationruntimeazuressis"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryIntegrationRuntimeAzureSsis creates a new instance of [DataFactoryIntegrationRuntimeAzureSsis].
func NewDataFactoryIntegrationRuntimeAzureSsis(name string, args DataFactoryIntegrationRuntimeAzureSsisArgs) *DataFactoryIntegrationRuntimeAzureSsis {
	return &DataFactoryIntegrationRuntimeAzureSsis{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryIntegrationRuntimeAzureSsis)(nil)

// DataFactoryIntegrationRuntimeAzureSsis represents the Terraform resource azurerm_data_factory_integration_runtime_azure_ssis.
type DataFactoryIntegrationRuntimeAzureSsis struct {
	Name      string
	Args      DataFactoryIntegrationRuntimeAzureSsisArgs
	state     *dataFactoryIntegrationRuntimeAzureSsisState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryIntegrationRuntimeAzureSsis].
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) Type() string {
	return "azurerm_data_factory_integration_runtime_azure_ssis"
}

// LocalName returns the local name for [DataFactoryIntegrationRuntimeAzureSsis].
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) LocalName() string {
	return dfiras.Name
}

// Configuration returns the configuration (args) for [DataFactoryIntegrationRuntimeAzureSsis].
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) Configuration() interface{} {
	return dfiras.Args
}

// DependOn is used for other resources to depend on [DataFactoryIntegrationRuntimeAzureSsis].
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) DependOn() terra.Reference {
	return terra.ReferenceResource(dfiras)
}

// Dependencies returns the list of resources [DataFactoryIntegrationRuntimeAzureSsis] depends_on.
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) Dependencies() terra.Dependencies {
	return dfiras.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryIntegrationRuntimeAzureSsis].
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) LifecycleManagement() *terra.Lifecycle {
	return dfiras.Lifecycle
}

// Attributes returns the attributes for [DataFactoryIntegrationRuntimeAzureSsis].
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) Attributes() dataFactoryIntegrationRuntimeAzureSsisAttributes {
	return dataFactoryIntegrationRuntimeAzureSsisAttributes{ref: terra.ReferenceResource(dfiras)}
}

// ImportState imports the given attribute values into [DataFactoryIntegrationRuntimeAzureSsis]'s state.
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) ImportState(av io.Reader) error {
	dfiras.state = &dataFactoryIntegrationRuntimeAzureSsisState{}
	if err := json.NewDecoder(av).Decode(dfiras.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfiras.Type(), dfiras.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryIntegrationRuntimeAzureSsis] has state.
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) State() (*dataFactoryIntegrationRuntimeAzureSsisState, bool) {
	return dfiras.state, dfiras.state != nil
}

// StateMust returns the state for [DataFactoryIntegrationRuntimeAzureSsis]. Panics if the state is nil.
func (dfiras *DataFactoryIntegrationRuntimeAzureSsis) StateMust() *dataFactoryIntegrationRuntimeAzureSsisState {
	if dfiras.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfiras.Type(), dfiras.LocalName()))
	}
	return dfiras.state
}

// DataFactoryIntegrationRuntimeAzureSsisArgs contains the configurations for azurerm_data_factory_integration_runtime_azure_ssis.
type DataFactoryIntegrationRuntimeAzureSsisArgs struct {
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Edition: string, optional
	Edition terra.StringValue `hcl:"edition,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, optional
	LicenseType terra.StringValue `hcl:"license_type,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MaxParallelExecutionsPerNode: number, optional
	MaxParallelExecutionsPerNode terra.NumberValue `hcl:"max_parallel_executions_per_node,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeSize: string, required
	NodeSize terra.StringValue `hcl:"node_size,attr" validate:"required"`
	// NumberOfNodes: number, optional
	NumberOfNodes terra.NumberValue `hcl:"number_of_nodes,attr"`
	// CatalogInfo: optional
	CatalogInfo *datafactoryintegrationruntimeazuressis.CatalogInfo `hcl:"catalog_info,block"`
	// CustomSetupScript: optional
	CustomSetupScript *datafactoryintegrationruntimeazuressis.CustomSetupScript `hcl:"custom_setup_script,block"`
	// ExpressCustomSetup: optional
	ExpressCustomSetup *datafactoryintegrationruntimeazuressis.ExpressCustomSetup `hcl:"express_custom_setup,block"`
	// ExpressVnetIntegration: optional
	ExpressVnetIntegration *datafactoryintegrationruntimeazuressis.ExpressVnetIntegration `hcl:"express_vnet_integration,block"`
	// PackageStore: min=0
	PackageStore []datafactoryintegrationruntimeazuressis.PackageStore `hcl:"package_store,block" validate:"min=0"`
	// Proxy: optional
	Proxy *datafactoryintegrationruntimeazuressis.Proxy `hcl:"proxy,block"`
	// Timeouts: optional
	Timeouts *datafactoryintegrationruntimeazuressis.Timeouts `hcl:"timeouts,block"`
	// VnetIntegration: optional
	VnetIntegration *datafactoryintegrationruntimeazuressis.VnetIntegration `hcl:"vnet_integration,block"`
}
type dataFactoryIntegrationRuntimeAzureSsisAttributes struct {
	ref terra.Reference
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("description"))
}

// Edition returns a reference to field edition of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("edition"))
}

// Id returns a reference to field id of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("id"))
}

// LicenseType returns a reference to field license_type of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("license_type"))
}

// Location returns a reference to field location of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("location"))
}

// MaxParallelExecutionsPerNode returns a reference to field max_parallel_executions_per_node of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) MaxParallelExecutionsPerNode() terra.NumberValue {
	return terra.ReferenceAsNumber(dfiras.ref.Append("max_parallel_executions_per_node"))
}

// Name returns a reference to field name of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("name"))
}

// NodeSize returns a reference to field node_size of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) NodeSize() terra.StringValue {
	return terra.ReferenceAsString(dfiras.ref.Append("node_size"))
}

// NumberOfNodes returns a reference to field number_of_nodes of azurerm_data_factory_integration_runtime_azure_ssis.
func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) NumberOfNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(dfiras.ref.Append("number_of_nodes"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) CatalogInfo() terra.ListValue[datafactoryintegrationruntimeazuressis.CatalogInfoAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.CatalogInfoAttributes](dfiras.ref.Append("catalog_info"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) CustomSetupScript() terra.ListValue[datafactoryintegrationruntimeazuressis.CustomSetupScriptAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.CustomSetupScriptAttributes](dfiras.ref.Append("custom_setup_script"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) ExpressCustomSetup() terra.ListValue[datafactoryintegrationruntimeazuressis.ExpressCustomSetupAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.ExpressCustomSetupAttributes](dfiras.ref.Append("express_custom_setup"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) ExpressVnetIntegration() terra.ListValue[datafactoryintegrationruntimeazuressis.ExpressVnetIntegrationAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.ExpressVnetIntegrationAttributes](dfiras.ref.Append("express_vnet_integration"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) PackageStore() terra.ListValue[datafactoryintegrationruntimeazuressis.PackageStoreAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.PackageStoreAttributes](dfiras.ref.Append("package_store"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Proxy() terra.ListValue[datafactoryintegrationruntimeazuressis.ProxyAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.ProxyAttributes](dfiras.ref.Append("proxy"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) Timeouts() datafactoryintegrationruntimeazuressis.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactoryintegrationruntimeazuressis.TimeoutsAttributes](dfiras.ref.Append("timeouts"))
}

func (dfiras dataFactoryIntegrationRuntimeAzureSsisAttributes) VnetIntegration() terra.ListValue[datafactoryintegrationruntimeazuressis.VnetIntegrationAttributes] {
	return terra.ReferenceAsList[datafactoryintegrationruntimeazuressis.VnetIntegrationAttributes](dfiras.ref.Append("vnet_integration"))
}

type dataFactoryIntegrationRuntimeAzureSsisState struct {
	DataFactoryId                string                                                               `json:"data_factory_id"`
	Description                  string                                                               `json:"description"`
	Edition                      string                                                               `json:"edition"`
	Id                           string                                                               `json:"id"`
	LicenseType                  string                                                               `json:"license_type"`
	Location                     string                                                               `json:"location"`
	MaxParallelExecutionsPerNode float64                                                              `json:"max_parallel_executions_per_node"`
	Name                         string                                                               `json:"name"`
	NodeSize                     string                                                               `json:"node_size"`
	NumberOfNodes                float64                                                              `json:"number_of_nodes"`
	CatalogInfo                  []datafactoryintegrationruntimeazuressis.CatalogInfoState            `json:"catalog_info"`
	CustomSetupScript            []datafactoryintegrationruntimeazuressis.CustomSetupScriptState      `json:"custom_setup_script"`
	ExpressCustomSetup           []datafactoryintegrationruntimeazuressis.ExpressCustomSetupState     `json:"express_custom_setup"`
	ExpressVnetIntegration       []datafactoryintegrationruntimeazuressis.ExpressVnetIntegrationState `json:"express_vnet_integration"`
	PackageStore                 []datafactoryintegrationruntimeazuressis.PackageStoreState           `json:"package_store"`
	Proxy                        []datafactoryintegrationruntimeazuressis.ProxyState                  `json:"proxy"`
	Timeouts                     *datafactoryintegrationruntimeazuressis.TimeoutsState                `json:"timeouts"`
	VnetIntegration              []datafactoryintegrationruntimeazuressis.VnetIntegrationState        `json:"vnet_integration"`
}
