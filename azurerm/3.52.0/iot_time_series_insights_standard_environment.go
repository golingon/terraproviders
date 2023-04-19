// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iottimeseriesinsightsstandardenvironment "github.com/golingon/terraproviders/azurerm/3.52.0/iottimeseriesinsightsstandardenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotTimeSeriesInsightsStandardEnvironment creates a new instance of [IotTimeSeriesInsightsStandardEnvironment].
func NewIotTimeSeriesInsightsStandardEnvironment(name string, args IotTimeSeriesInsightsStandardEnvironmentArgs) *IotTimeSeriesInsightsStandardEnvironment {
	return &IotTimeSeriesInsightsStandardEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotTimeSeriesInsightsStandardEnvironment)(nil)

// IotTimeSeriesInsightsStandardEnvironment represents the Terraform resource azurerm_iot_time_series_insights_standard_environment.
type IotTimeSeriesInsightsStandardEnvironment struct {
	Name      string
	Args      IotTimeSeriesInsightsStandardEnvironmentArgs
	state     *iotTimeSeriesInsightsStandardEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotTimeSeriesInsightsStandardEnvironment].
func (itsise *IotTimeSeriesInsightsStandardEnvironment) Type() string {
	return "azurerm_iot_time_series_insights_standard_environment"
}

// LocalName returns the local name for [IotTimeSeriesInsightsStandardEnvironment].
func (itsise *IotTimeSeriesInsightsStandardEnvironment) LocalName() string {
	return itsise.Name
}

// Configuration returns the configuration (args) for [IotTimeSeriesInsightsStandardEnvironment].
func (itsise *IotTimeSeriesInsightsStandardEnvironment) Configuration() interface{} {
	return itsise.Args
}

// DependOn is used for other resources to depend on [IotTimeSeriesInsightsStandardEnvironment].
func (itsise *IotTimeSeriesInsightsStandardEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(itsise)
}

// Dependencies returns the list of resources [IotTimeSeriesInsightsStandardEnvironment] depends_on.
func (itsise *IotTimeSeriesInsightsStandardEnvironment) Dependencies() terra.Dependencies {
	return itsise.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotTimeSeriesInsightsStandardEnvironment].
func (itsise *IotTimeSeriesInsightsStandardEnvironment) LifecycleManagement() *terra.Lifecycle {
	return itsise.Lifecycle
}

// Attributes returns the attributes for [IotTimeSeriesInsightsStandardEnvironment].
func (itsise *IotTimeSeriesInsightsStandardEnvironment) Attributes() iotTimeSeriesInsightsStandardEnvironmentAttributes {
	return iotTimeSeriesInsightsStandardEnvironmentAttributes{ref: terra.ReferenceResource(itsise)}
}

// ImportState imports the given attribute values into [IotTimeSeriesInsightsStandardEnvironment]'s state.
func (itsise *IotTimeSeriesInsightsStandardEnvironment) ImportState(av io.Reader) error {
	itsise.state = &iotTimeSeriesInsightsStandardEnvironmentState{}
	if err := json.NewDecoder(av).Decode(itsise.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itsise.Type(), itsise.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotTimeSeriesInsightsStandardEnvironment] has state.
func (itsise *IotTimeSeriesInsightsStandardEnvironment) State() (*iotTimeSeriesInsightsStandardEnvironmentState, bool) {
	return itsise.state, itsise.state != nil
}

// StateMust returns the state for [IotTimeSeriesInsightsStandardEnvironment]. Panics if the state is nil.
func (itsise *IotTimeSeriesInsightsStandardEnvironment) StateMust() *iotTimeSeriesInsightsStandardEnvironmentState {
	if itsise.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itsise.Type(), itsise.LocalName()))
	}
	return itsise.state
}

// IotTimeSeriesInsightsStandardEnvironmentArgs contains the configurations for azurerm_iot_time_series_insights_standard_environment.
type IotTimeSeriesInsightsStandardEnvironmentArgs struct {
	// DataRetentionTime: string, required
	DataRetentionTime terra.StringValue `hcl:"data_retention_time,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PartitionKey: string, optional
	PartitionKey terra.StringValue `hcl:"partition_key,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// StorageLimitExceededBehavior: string, optional
	StorageLimitExceededBehavior terra.StringValue `hcl:"storage_limit_exceeded_behavior,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *iottimeseriesinsightsstandardenvironment.Timeouts `hcl:"timeouts,block"`
}
type iotTimeSeriesInsightsStandardEnvironmentAttributes struct {
	ref terra.Reference
}

// DataRetentionTime returns a reference to field data_retention_time of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) DataRetentionTime() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("data_retention_time"))
}

// Id returns a reference to field id of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("name"))
}

// PartitionKey returns a reference to field partition_key of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) PartitionKey() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("partition_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("sku_name"))
}

// StorageLimitExceededBehavior returns a reference to field storage_limit_exceeded_behavior of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) StorageLimitExceededBehavior() terra.StringValue {
	return terra.ReferenceAsString(itsise.ref.Append("storage_limit_exceeded_behavior"))
}

// Tags returns a reference to field tags of azurerm_iot_time_series_insights_standard_environment.
func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itsise.ref.Append("tags"))
}

func (itsise iotTimeSeriesInsightsStandardEnvironmentAttributes) Timeouts() iottimeseriesinsightsstandardenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iottimeseriesinsightsstandardenvironment.TimeoutsAttributes](itsise.ref.Append("timeouts"))
}

type iotTimeSeriesInsightsStandardEnvironmentState struct {
	DataRetentionTime            string                                                  `json:"data_retention_time"`
	Id                           string                                                  `json:"id"`
	Location                     string                                                  `json:"location"`
	Name                         string                                                  `json:"name"`
	PartitionKey                 string                                                  `json:"partition_key"`
	ResourceGroupName            string                                                  `json:"resource_group_name"`
	SkuName                      string                                                  `json:"sku_name"`
	StorageLimitExceededBehavior string                                                  `json:"storage_limit_exceeded_behavior"`
	Tags                         map[string]string                                       `json:"tags"`
	Timeouts                     *iottimeseriesinsightsstandardenvironment.TimeoutsState `json:"timeouts"`
}
