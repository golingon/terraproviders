// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsfunctionjavascriptuda "github.com/golingon/terraproviders/azurerm/3.52.0/streamanalyticsfunctionjavascriptuda"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsFunctionJavascriptUda creates a new instance of [StreamAnalyticsFunctionJavascriptUda].
func NewStreamAnalyticsFunctionJavascriptUda(name string, args StreamAnalyticsFunctionJavascriptUdaArgs) *StreamAnalyticsFunctionJavascriptUda {
	return &StreamAnalyticsFunctionJavascriptUda{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsFunctionJavascriptUda)(nil)

// StreamAnalyticsFunctionJavascriptUda represents the Terraform resource azurerm_stream_analytics_function_javascript_uda.
type StreamAnalyticsFunctionJavascriptUda struct {
	Name      string
	Args      StreamAnalyticsFunctionJavascriptUdaArgs
	state     *streamAnalyticsFunctionJavascriptUdaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsFunctionJavascriptUda].
func (safju *StreamAnalyticsFunctionJavascriptUda) Type() string {
	return "azurerm_stream_analytics_function_javascript_uda"
}

// LocalName returns the local name for [StreamAnalyticsFunctionJavascriptUda].
func (safju *StreamAnalyticsFunctionJavascriptUda) LocalName() string {
	return safju.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsFunctionJavascriptUda].
func (safju *StreamAnalyticsFunctionJavascriptUda) Configuration() interface{} {
	return safju.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsFunctionJavascriptUda].
func (safju *StreamAnalyticsFunctionJavascriptUda) DependOn() terra.Reference {
	return terra.ReferenceResource(safju)
}

// Dependencies returns the list of resources [StreamAnalyticsFunctionJavascriptUda] depends_on.
func (safju *StreamAnalyticsFunctionJavascriptUda) Dependencies() terra.Dependencies {
	return safju.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsFunctionJavascriptUda].
func (safju *StreamAnalyticsFunctionJavascriptUda) LifecycleManagement() *terra.Lifecycle {
	return safju.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsFunctionJavascriptUda].
func (safju *StreamAnalyticsFunctionJavascriptUda) Attributes() streamAnalyticsFunctionJavascriptUdaAttributes {
	return streamAnalyticsFunctionJavascriptUdaAttributes{ref: terra.ReferenceResource(safju)}
}

// ImportState imports the given attribute values into [StreamAnalyticsFunctionJavascriptUda]'s state.
func (safju *StreamAnalyticsFunctionJavascriptUda) ImportState(av io.Reader) error {
	safju.state = &streamAnalyticsFunctionJavascriptUdaState{}
	if err := json.NewDecoder(av).Decode(safju.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", safju.Type(), safju.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsFunctionJavascriptUda] has state.
func (safju *StreamAnalyticsFunctionJavascriptUda) State() (*streamAnalyticsFunctionJavascriptUdaState, bool) {
	return safju.state, safju.state != nil
}

// StateMust returns the state for [StreamAnalyticsFunctionJavascriptUda]. Panics if the state is nil.
func (safju *StreamAnalyticsFunctionJavascriptUda) StateMust() *streamAnalyticsFunctionJavascriptUdaState {
	if safju.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", safju.Type(), safju.LocalName()))
	}
	return safju.state
}

// StreamAnalyticsFunctionJavascriptUdaArgs contains the configurations for azurerm_stream_analytics_function_javascript_uda.
type StreamAnalyticsFunctionJavascriptUdaArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Script: string, required
	Script terra.StringValue `hcl:"script,attr" validate:"required"`
	// StreamAnalyticsJobId: string, required
	StreamAnalyticsJobId terra.StringValue `hcl:"stream_analytics_job_id,attr" validate:"required"`
	// Input: min=1
	Input []streamanalyticsfunctionjavascriptuda.Input `hcl:"input,block" validate:"min=1"`
	// Output: required
	Output *streamanalyticsfunctionjavascriptuda.Output `hcl:"output,block" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsfunctionjavascriptuda.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsFunctionJavascriptUdaAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_stream_analytics_function_javascript_uda.
func (safju streamAnalyticsFunctionJavascriptUdaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(safju.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_function_javascript_uda.
func (safju streamAnalyticsFunctionJavascriptUdaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(safju.ref.Append("name"))
}

// Script returns a reference to field script of azurerm_stream_analytics_function_javascript_uda.
func (safju streamAnalyticsFunctionJavascriptUdaAttributes) Script() terra.StringValue {
	return terra.ReferenceAsString(safju.ref.Append("script"))
}

// StreamAnalyticsJobId returns a reference to field stream_analytics_job_id of azurerm_stream_analytics_function_javascript_uda.
func (safju streamAnalyticsFunctionJavascriptUdaAttributes) StreamAnalyticsJobId() terra.StringValue {
	return terra.ReferenceAsString(safju.ref.Append("stream_analytics_job_id"))
}

func (safju streamAnalyticsFunctionJavascriptUdaAttributes) Input() terra.ListValue[streamanalyticsfunctionjavascriptuda.InputAttributes] {
	return terra.ReferenceAsList[streamanalyticsfunctionjavascriptuda.InputAttributes](safju.ref.Append("input"))
}

func (safju streamAnalyticsFunctionJavascriptUdaAttributes) Output() terra.ListValue[streamanalyticsfunctionjavascriptuda.OutputAttributes] {
	return terra.ReferenceAsList[streamanalyticsfunctionjavascriptuda.OutputAttributes](safju.ref.Append("output"))
}

func (safju streamAnalyticsFunctionJavascriptUdaAttributes) Timeouts() streamanalyticsfunctionjavascriptuda.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsfunctionjavascriptuda.TimeoutsAttributes](safju.ref.Append("timeouts"))
}

type streamAnalyticsFunctionJavascriptUdaState struct {
	Id                   string                                              `json:"id"`
	Name                 string                                              `json:"name"`
	Script               string                                              `json:"script"`
	StreamAnalyticsJobId string                                              `json:"stream_analytics_job_id"`
	Input                []streamanalyticsfunctionjavascriptuda.InputState   `json:"input"`
	Output               []streamanalyticsfunctionjavascriptuda.OutputState  `json:"output"`
	Timeouts             *streamanalyticsfunctionjavascriptuda.TimeoutsState `json:"timeouts"`
}
