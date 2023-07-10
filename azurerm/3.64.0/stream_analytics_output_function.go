// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsoutputfunction "github.com/golingon/terraproviders/azurerm/3.64.0/streamanalyticsoutputfunction"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsOutputFunction creates a new instance of [StreamAnalyticsOutputFunction].
func NewStreamAnalyticsOutputFunction(name string, args StreamAnalyticsOutputFunctionArgs) *StreamAnalyticsOutputFunction {
	return &StreamAnalyticsOutputFunction{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputFunction)(nil)

// StreamAnalyticsOutputFunction represents the Terraform resource azurerm_stream_analytics_output_function.
type StreamAnalyticsOutputFunction struct {
	Name      string
	Args      StreamAnalyticsOutputFunctionArgs
	state     *streamAnalyticsOutputFunctionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputFunction].
func (saof *StreamAnalyticsOutputFunction) Type() string {
	return "azurerm_stream_analytics_output_function"
}

// LocalName returns the local name for [StreamAnalyticsOutputFunction].
func (saof *StreamAnalyticsOutputFunction) LocalName() string {
	return saof.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputFunction].
func (saof *StreamAnalyticsOutputFunction) Configuration() interface{} {
	return saof.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputFunction].
func (saof *StreamAnalyticsOutputFunction) DependOn() terra.Reference {
	return terra.ReferenceResource(saof)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputFunction] depends_on.
func (saof *StreamAnalyticsOutputFunction) Dependencies() terra.Dependencies {
	return saof.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputFunction].
func (saof *StreamAnalyticsOutputFunction) LifecycleManagement() *terra.Lifecycle {
	return saof.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputFunction].
func (saof *StreamAnalyticsOutputFunction) Attributes() streamAnalyticsOutputFunctionAttributes {
	return streamAnalyticsOutputFunctionAttributes{ref: terra.ReferenceResource(saof)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputFunction]'s state.
func (saof *StreamAnalyticsOutputFunction) ImportState(av io.Reader) error {
	saof.state = &streamAnalyticsOutputFunctionState{}
	if err := json.NewDecoder(av).Decode(saof.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saof.Type(), saof.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputFunction] has state.
func (saof *StreamAnalyticsOutputFunction) State() (*streamAnalyticsOutputFunctionState, bool) {
	return saof.state, saof.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputFunction]. Panics if the state is nil.
func (saof *StreamAnalyticsOutputFunction) StateMust() *streamAnalyticsOutputFunctionState {
	if saof.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saof.Type(), saof.LocalName()))
	}
	return saof.state
}

// StreamAnalyticsOutputFunctionArgs contains the configurations for azurerm_stream_analytics_output_function.
type StreamAnalyticsOutputFunctionArgs struct {
	// ApiKey: string, required
	ApiKey terra.StringValue `hcl:"api_key,attr" validate:"required"`
	// BatchMaxCount: number, optional
	BatchMaxCount terra.NumberValue `hcl:"batch_max_count,attr"`
	// BatchMaxInBytes: number, optional
	BatchMaxInBytes terra.NumberValue `hcl:"batch_max_in_bytes,attr"`
	// FunctionApp: string, required
	FunctionApp terra.StringValue `hcl:"function_app,attr" validate:"required"`
	// FunctionName: string, required
	FunctionName terra.StringValue `hcl:"function_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputfunction.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputFunctionAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("api_key"))
}

// BatchMaxCount returns a reference to field batch_max_count of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) BatchMaxCount() terra.NumberValue {
	return terra.ReferenceAsNumber(saof.ref.Append("batch_max_count"))
}

// BatchMaxInBytes returns a reference to field batch_max_in_bytes of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) BatchMaxInBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(saof.ref.Append("batch_max_in_bytes"))
}

// FunctionApp returns a reference to field function_app of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) FunctionApp() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("function_app"))
}

// FunctionName returns a reference to field function_name of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) FunctionName() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("function_name"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("resource_group_name"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_function.
func (saof streamAnalyticsOutputFunctionAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(saof.ref.Append("stream_analytics_job_name"))
}

func (saof streamAnalyticsOutputFunctionAttributes) Timeouts() streamanalyticsoutputfunction.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputfunction.TimeoutsAttributes](saof.ref.Append("timeouts"))
}

type streamAnalyticsOutputFunctionState struct {
	ApiKey                 string                                       `json:"api_key"`
	BatchMaxCount          float64                                      `json:"batch_max_count"`
	BatchMaxInBytes        float64                                      `json:"batch_max_in_bytes"`
	FunctionApp            string                                       `json:"function_app"`
	FunctionName           string                                       `json:"function_name"`
	Id                     string                                       `json:"id"`
	Name                   string                                       `json:"name"`
	ResourceGroupName      string                                       `json:"resource_group_name"`
	StreamAnalyticsJobName string                                       `json:"stream_analytics_job_name"`
	Timeouts               *streamanalyticsoutputfunction.TimeoutsState `json:"timeouts"`
}
