// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iottimeseriesinsightsgen2environment "github.com/golingon/terraproviders/azurerm/3.69.0/iottimeseriesinsightsgen2environment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotTimeSeriesInsightsGen2Environment creates a new instance of [IotTimeSeriesInsightsGen2Environment].
func NewIotTimeSeriesInsightsGen2Environment(name string, args IotTimeSeriesInsightsGen2EnvironmentArgs) *IotTimeSeriesInsightsGen2Environment {
	return &IotTimeSeriesInsightsGen2Environment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotTimeSeriesInsightsGen2Environment)(nil)

// IotTimeSeriesInsightsGen2Environment represents the Terraform resource azurerm_iot_time_series_insights_gen2_environment.
type IotTimeSeriesInsightsGen2Environment struct {
	Name      string
	Args      IotTimeSeriesInsightsGen2EnvironmentArgs
	state     *iotTimeSeriesInsightsGen2EnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotTimeSeriesInsightsGen2Environment].
func (itsige *IotTimeSeriesInsightsGen2Environment) Type() string {
	return "azurerm_iot_time_series_insights_gen2_environment"
}

// LocalName returns the local name for [IotTimeSeriesInsightsGen2Environment].
func (itsige *IotTimeSeriesInsightsGen2Environment) LocalName() string {
	return itsige.Name
}

// Configuration returns the configuration (args) for [IotTimeSeriesInsightsGen2Environment].
func (itsige *IotTimeSeriesInsightsGen2Environment) Configuration() interface{} {
	return itsige.Args
}

// DependOn is used for other resources to depend on [IotTimeSeriesInsightsGen2Environment].
func (itsige *IotTimeSeriesInsightsGen2Environment) DependOn() terra.Reference {
	return terra.ReferenceResource(itsige)
}

// Dependencies returns the list of resources [IotTimeSeriesInsightsGen2Environment] depends_on.
func (itsige *IotTimeSeriesInsightsGen2Environment) Dependencies() terra.Dependencies {
	return itsige.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotTimeSeriesInsightsGen2Environment].
func (itsige *IotTimeSeriesInsightsGen2Environment) LifecycleManagement() *terra.Lifecycle {
	return itsige.Lifecycle
}

// Attributes returns the attributes for [IotTimeSeriesInsightsGen2Environment].
func (itsige *IotTimeSeriesInsightsGen2Environment) Attributes() iotTimeSeriesInsightsGen2EnvironmentAttributes {
	return iotTimeSeriesInsightsGen2EnvironmentAttributes{ref: terra.ReferenceResource(itsige)}
}

// ImportState imports the given attribute values into [IotTimeSeriesInsightsGen2Environment]'s state.
func (itsige *IotTimeSeriesInsightsGen2Environment) ImportState(av io.Reader) error {
	itsige.state = &iotTimeSeriesInsightsGen2EnvironmentState{}
	if err := json.NewDecoder(av).Decode(itsige.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itsige.Type(), itsige.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotTimeSeriesInsightsGen2Environment] has state.
func (itsige *IotTimeSeriesInsightsGen2Environment) State() (*iotTimeSeriesInsightsGen2EnvironmentState, bool) {
	return itsige.state, itsige.state != nil
}

// StateMust returns the state for [IotTimeSeriesInsightsGen2Environment]. Panics if the state is nil.
func (itsige *IotTimeSeriesInsightsGen2Environment) StateMust() *iotTimeSeriesInsightsGen2EnvironmentState {
	if itsige.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itsige.Type(), itsige.LocalName()))
	}
	return itsige.state
}

// IotTimeSeriesInsightsGen2EnvironmentArgs contains the configurations for azurerm_iot_time_series_insights_gen2_environment.
type IotTimeSeriesInsightsGen2EnvironmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdProperties: list of string, required
	IdProperties terra.ListValue[terra.StringValue] `hcl:"id_properties,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WarmStoreDataRetentionTime: string, optional
	WarmStoreDataRetentionTime terra.StringValue `hcl:"warm_store_data_retention_time,attr"`
	// Storage: required
	Storage *iottimeseriesinsightsgen2environment.Storage `hcl:"storage,block" validate:"required"`
	// Timeouts: optional
	Timeouts *iottimeseriesinsightsgen2environment.Timeouts `hcl:"timeouts,block"`
}
type iotTimeSeriesInsightsGen2EnvironmentAttributes struct {
	ref terra.Reference
}

// DataAccessFqdn returns a reference to field data_access_fqdn of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) DataAccessFqdn() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("data_access_fqdn"))
}

// Id returns a reference to field id of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("id"))
}

// IdProperties returns a reference to field id_properties of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) IdProperties() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](itsige.ref.Append("id_properties"))
}

// Location returns a reference to field location of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itsige.ref.Append("tags"))
}

// WarmStoreDataRetentionTime returns a reference to field warm_store_data_retention_time of azurerm_iot_time_series_insights_gen2_environment.
func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) WarmStoreDataRetentionTime() terra.StringValue {
	return terra.ReferenceAsString(itsige.ref.Append("warm_store_data_retention_time"))
}

func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) Storage() terra.ListValue[iottimeseriesinsightsgen2environment.StorageAttributes] {
	return terra.ReferenceAsList[iottimeseriesinsightsgen2environment.StorageAttributes](itsige.ref.Append("storage"))
}

func (itsige iotTimeSeriesInsightsGen2EnvironmentAttributes) Timeouts() iottimeseriesinsightsgen2environment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iottimeseriesinsightsgen2environment.TimeoutsAttributes](itsige.ref.Append("timeouts"))
}

type iotTimeSeriesInsightsGen2EnvironmentState struct {
	DataAccessFqdn             string                                              `json:"data_access_fqdn"`
	Id                         string                                              `json:"id"`
	IdProperties               []string                                            `json:"id_properties"`
	Location                   string                                              `json:"location"`
	Name                       string                                              `json:"name"`
	ResourceGroupName          string                                              `json:"resource_group_name"`
	SkuName                    string                                              `json:"sku_name"`
	Tags                       map[string]string                                   `json:"tags"`
	WarmStoreDataRetentionTime string                                              `json:"warm_store_data_retention_time"`
	Storage                    []iottimeseriesinsightsgen2environment.StorageState `json:"storage"`
	Timeouts                   *iottimeseriesinsightsgen2environment.TimeoutsState `json:"timeouts"`
}
