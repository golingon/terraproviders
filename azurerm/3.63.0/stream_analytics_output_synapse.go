// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsoutputsynapse "github.com/golingon/terraproviders/azurerm/3.63.0/streamanalyticsoutputsynapse"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsOutputSynapse creates a new instance of [StreamAnalyticsOutputSynapse].
func NewStreamAnalyticsOutputSynapse(name string, args StreamAnalyticsOutputSynapseArgs) *StreamAnalyticsOutputSynapse {
	return &StreamAnalyticsOutputSynapse{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputSynapse)(nil)

// StreamAnalyticsOutputSynapse represents the Terraform resource azurerm_stream_analytics_output_synapse.
type StreamAnalyticsOutputSynapse struct {
	Name      string
	Args      StreamAnalyticsOutputSynapseArgs
	state     *streamAnalyticsOutputSynapseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputSynapse].
func (saos *StreamAnalyticsOutputSynapse) Type() string {
	return "azurerm_stream_analytics_output_synapse"
}

// LocalName returns the local name for [StreamAnalyticsOutputSynapse].
func (saos *StreamAnalyticsOutputSynapse) LocalName() string {
	return saos.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputSynapse].
func (saos *StreamAnalyticsOutputSynapse) Configuration() interface{} {
	return saos.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputSynapse].
func (saos *StreamAnalyticsOutputSynapse) DependOn() terra.Reference {
	return terra.ReferenceResource(saos)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputSynapse] depends_on.
func (saos *StreamAnalyticsOutputSynapse) Dependencies() terra.Dependencies {
	return saos.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputSynapse].
func (saos *StreamAnalyticsOutputSynapse) LifecycleManagement() *terra.Lifecycle {
	return saos.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputSynapse].
func (saos *StreamAnalyticsOutputSynapse) Attributes() streamAnalyticsOutputSynapseAttributes {
	return streamAnalyticsOutputSynapseAttributes{ref: terra.ReferenceResource(saos)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputSynapse]'s state.
func (saos *StreamAnalyticsOutputSynapse) ImportState(av io.Reader) error {
	saos.state = &streamAnalyticsOutputSynapseState{}
	if err := json.NewDecoder(av).Decode(saos.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saos.Type(), saos.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputSynapse] has state.
func (saos *StreamAnalyticsOutputSynapse) State() (*streamAnalyticsOutputSynapseState, bool) {
	return saos.state, saos.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputSynapse]. Panics if the state is nil.
func (saos *StreamAnalyticsOutputSynapse) StateMust() *streamAnalyticsOutputSynapseState {
	if saos.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saos.Type(), saos.LocalName()))
	}
	return saos.state
}

// StreamAnalyticsOutputSynapseArgs contains the configurations for azurerm_stream_analytics_output_synapse.
type StreamAnalyticsOutputSynapseArgs struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
	// User: string, required
	User terra.StringValue `hcl:"user,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputsynapse.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputSynapseAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("database"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("password"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("resource_group_name"))
}

// Server returns a reference to field server of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("server"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("stream_analytics_job_name"))
}

// Table returns a reference to field table of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("table"))
}

// User returns a reference to field user of azurerm_stream_analytics_output_synapse.
func (saos streamAnalyticsOutputSynapseAttributes) User() terra.StringValue {
	return terra.ReferenceAsString(saos.ref.Append("user"))
}

func (saos streamAnalyticsOutputSynapseAttributes) Timeouts() streamanalyticsoutputsynapse.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputsynapse.TimeoutsAttributes](saos.ref.Append("timeouts"))
}

type streamAnalyticsOutputSynapseState struct {
	Database               string                                      `json:"database"`
	Id                     string                                      `json:"id"`
	Name                   string                                      `json:"name"`
	Password               string                                      `json:"password"`
	ResourceGroupName      string                                      `json:"resource_group_name"`
	Server                 string                                      `json:"server"`
	StreamAnalyticsJobName string                                      `json:"stream_analytics_job_name"`
	Table                  string                                      `json:"table"`
	User                   string                                      `json:"user"`
	Timeouts               *streamanalyticsoutputsynapse.TimeoutsState `json:"timeouts"`
}