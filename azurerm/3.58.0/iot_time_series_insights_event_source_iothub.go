// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iottimeseriesinsightseventsourceiothub "github.com/golingon/terraproviders/azurerm/3.58.0/iottimeseriesinsightseventsourceiothub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotTimeSeriesInsightsEventSourceIothub creates a new instance of [IotTimeSeriesInsightsEventSourceIothub].
func NewIotTimeSeriesInsightsEventSourceIothub(name string, args IotTimeSeriesInsightsEventSourceIothubArgs) *IotTimeSeriesInsightsEventSourceIothub {
	return &IotTimeSeriesInsightsEventSourceIothub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotTimeSeriesInsightsEventSourceIothub)(nil)

// IotTimeSeriesInsightsEventSourceIothub represents the Terraform resource azurerm_iot_time_series_insights_event_source_iothub.
type IotTimeSeriesInsightsEventSourceIothub struct {
	Name      string
	Args      IotTimeSeriesInsightsEventSourceIothubArgs
	state     *iotTimeSeriesInsightsEventSourceIothubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotTimeSeriesInsightsEventSourceIothub].
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) Type() string {
	return "azurerm_iot_time_series_insights_event_source_iothub"
}

// LocalName returns the local name for [IotTimeSeriesInsightsEventSourceIothub].
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) LocalName() string {
	return itsiesi.Name
}

// Configuration returns the configuration (args) for [IotTimeSeriesInsightsEventSourceIothub].
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) Configuration() interface{} {
	return itsiesi.Args
}

// DependOn is used for other resources to depend on [IotTimeSeriesInsightsEventSourceIothub].
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) DependOn() terra.Reference {
	return terra.ReferenceResource(itsiesi)
}

// Dependencies returns the list of resources [IotTimeSeriesInsightsEventSourceIothub] depends_on.
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) Dependencies() terra.Dependencies {
	return itsiesi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotTimeSeriesInsightsEventSourceIothub].
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) LifecycleManagement() *terra.Lifecycle {
	return itsiesi.Lifecycle
}

// Attributes returns the attributes for [IotTimeSeriesInsightsEventSourceIothub].
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) Attributes() iotTimeSeriesInsightsEventSourceIothubAttributes {
	return iotTimeSeriesInsightsEventSourceIothubAttributes{ref: terra.ReferenceResource(itsiesi)}
}

// ImportState imports the given attribute values into [IotTimeSeriesInsightsEventSourceIothub]'s state.
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) ImportState(av io.Reader) error {
	itsiesi.state = &iotTimeSeriesInsightsEventSourceIothubState{}
	if err := json.NewDecoder(av).Decode(itsiesi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itsiesi.Type(), itsiesi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotTimeSeriesInsightsEventSourceIothub] has state.
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) State() (*iotTimeSeriesInsightsEventSourceIothubState, bool) {
	return itsiesi.state, itsiesi.state != nil
}

// StateMust returns the state for [IotTimeSeriesInsightsEventSourceIothub]. Panics if the state is nil.
func (itsiesi *IotTimeSeriesInsightsEventSourceIothub) StateMust() *iotTimeSeriesInsightsEventSourceIothubState {
	if itsiesi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itsiesi.Type(), itsiesi.LocalName()))
	}
	return itsiesi.state
}

// IotTimeSeriesInsightsEventSourceIothubArgs contains the configurations for azurerm_iot_time_series_insights_event_source_iothub.
type IotTimeSeriesInsightsEventSourceIothubArgs struct {
	// ConsumerGroupName: string, required
	ConsumerGroupName terra.StringValue `hcl:"consumer_group_name,attr" validate:"required"`
	// EnvironmentId: string, required
	EnvironmentId terra.StringValue `hcl:"environment_id,attr" validate:"required"`
	// EventSourceResourceId: string, required
	EventSourceResourceId terra.StringValue `hcl:"event_source_resource_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubName: string, required
	IothubName terra.StringValue `hcl:"iothub_name,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SharedAccessKey: string, required
	SharedAccessKey terra.StringValue `hcl:"shared_access_key,attr" validate:"required"`
	// SharedAccessKeyName: string, required
	SharedAccessKeyName terra.StringValue `hcl:"shared_access_key_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TimestampPropertyName: string, optional
	TimestampPropertyName terra.StringValue `hcl:"timestamp_property_name,attr"`
	// Timeouts: optional
	Timeouts *iottimeseriesinsightseventsourceiothub.Timeouts `hcl:"timeouts,block"`
}
type iotTimeSeriesInsightsEventSourceIothubAttributes struct {
	ref terra.Reference
}

// ConsumerGroupName returns a reference to field consumer_group_name of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) ConsumerGroupName() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("consumer_group_name"))
}

// EnvironmentId returns a reference to field environment_id of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) EnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("environment_id"))
}

// EventSourceResourceId returns a reference to field event_source_resource_id of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) EventSourceResourceId() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("event_source_resource_id"))
}

// Id returns a reference to field id of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("iothub_name"))
}

// Location returns a reference to field location of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("name"))
}

// SharedAccessKey returns a reference to field shared_access_key of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) SharedAccessKey() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("shared_access_key"))
}

// SharedAccessKeyName returns a reference to field shared_access_key_name of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) SharedAccessKeyName() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("shared_access_key_name"))
}

// Tags returns a reference to field tags of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itsiesi.ref.Append("tags"))
}

// TimestampPropertyName returns a reference to field timestamp_property_name of azurerm_iot_time_series_insights_event_source_iothub.
func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) TimestampPropertyName() terra.StringValue {
	return terra.ReferenceAsString(itsiesi.ref.Append("timestamp_property_name"))
}

func (itsiesi iotTimeSeriesInsightsEventSourceIothubAttributes) Timeouts() iottimeseriesinsightseventsourceiothub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iottimeseriesinsightseventsourceiothub.TimeoutsAttributes](itsiesi.ref.Append("timeouts"))
}

type iotTimeSeriesInsightsEventSourceIothubState struct {
	ConsumerGroupName     string                                                `json:"consumer_group_name"`
	EnvironmentId         string                                                `json:"environment_id"`
	EventSourceResourceId string                                                `json:"event_source_resource_id"`
	Id                    string                                                `json:"id"`
	IothubName            string                                                `json:"iothub_name"`
	Location              string                                                `json:"location"`
	Name                  string                                                `json:"name"`
	SharedAccessKey       string                                                `json:"shared_access_key"`
	SharedAccessKeyName   string                                                `json:"shared_access_key_name"`
	Tags                  map[string]string                                     `json:"tags"`
	TimestampPropertyName string                                                `json:"timestamp_property_name"`
	Timeouts              *iottimeseriesinsightseventsourceiothub.TimeoutsState `json:"timeouts"`
}
