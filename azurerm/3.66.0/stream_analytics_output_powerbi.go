// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsoutputpowerbi "github.com/golingon/terraproviders/azurerm/3.66.0/streamanalyticsoutputpowerbi"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsOutputPowerbi creates a new instance of [StreamAnalyticsOutputPowerbi].
func NewStreamAnalyticsOutputPowerbi(name string, args StreamAnalyticsOutputPowerbiArgs) *StreamAnalyticsOutputPowerbi {
	return &StreamAnalyticsOutputPowerbi{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsOutputPowerbi)(nil)

// StreamAnalyticsOutputPowerbi represents the Terraform resource azurerm_stream_analytics_output_powerbi.
type StreamAnalyticsOutputPowerbi struct {
	Name      string
	Args      StreamAnalyticsOutputPowerbiArgs
	state     *streamAnalyticsOutputPowerbiState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsOutputPowerbi].
func (saop *StreamAnalyticsOutputPowerbi) Type() string {
	return "azurerm_stream_analytics_output_powerbi"
}

// LocalName returns the local name for [StreamAnalyticsOutputPowerbi].
func (saop *StreamAnalyticsOutputPowerbi) LocalName() string {
	return saop.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsOutputPowerbi].
func (saop *StreamAnalyticsOutputPowerbi) Configuration() interface{} {
	return saop.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsOutputPowerbi].
func (saop *StreamAnalyticsOutputPowerbi) DependOn() terra.Reference {
	return terra.ReferenceResource(saop)
}

// Dependencies returns the list of resources [StreamAnalyticsOutputPowerbi] depends_on.
func (saop *StreamAnalyticsOutputPowerbi) Dependencies() terra.Dependencies {
	return saop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsOutputPowerbi].
func (saop *StreamAnalyticsOutputPowerbi) LifecycleManagement() *terra.Lifecycle {
	return saop.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsOutputPowerbi].
func (saop *StreamAnalyticsOutputPowerbi) Attributes() streamAnalyticsOutputPowerbiAttributes {
	return streamAnalyticsOutputPowerbiAttributes{ref: terra.ReferenceResource(saop)}
}

// ImportState imports the given attribute values into [StreamAnalyticsOutputPowerbi]'s state.
func (saop *StreamAnalyticsOutputPowerbi) ImportState(av io.Reader) error {
	saop.state = &streamAnalyticsOutputPowerbiState{}
	if err := json.NewDecoder(av).Decode(saop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saop.Type(), saop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsOutputPowerbi] has state.
func (saop *StreamAnalyticsOutputPowerbi) State() (*streamAnalyticsOutputPowerbiState, bool) {
	return saop.state, saop.state != nil
}

// StateMust returns the state for [StreamAnalyticsOutputPowerbi]. Panics if the state is nil.
func (saop *StreamAnalyticsOutputPowerbi) StateMust() *streamAnalyticsOutputPowerbiState {
	if saop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saop.Type(), saop.LocalName()))
	}
	return saop.state
}

// StreamAnalyticsOutputPowerbiArgs contains the configurations for azurerm_stream_analytics_output_powerbi.
type StreamAnalyticsOutputPowerbiArgs struct {
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// GroupId: string, required
	GroupId terra.StringValue `hcl:"group_id,attr" validate:"required"`
	// GroupName: string, required
	GroupName terra.StringValue `hcl:"group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StreamAnalyticsJobId: string, required
	StreamAnalyticsJobId terra.StringValue `hcl:"stream_analytics_job_id,attr" validate:"required"`
	// Table: string, required
	Table terra.StringValue `hcl:"table,attr" validate:"required"`
	// TokenUserDisplayName: string, optional
	TokenUserDisplayName terra.StringValue `hcl:"token_user_display_name,attr"`
	// TokenUserPrincipalName: string, optional
	TokenUserPrincipalName terra.StringValue `hcl:"token_user_principal_name,attr"`
	// Timeouts: optional
	Timeouts *streamanalyticsoutputpowerbi.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsOutputPowerbiAttributes struct {
	ref terra.Reference
}

// Dataset returns a reference to field dataset of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("dataset"))
}

// GroupId returns a reference to field group_id of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("group_id"))
}

// GroupName returns a reference to field group_name of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("group_name"))
}

// Id returns a reference to field id of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("name"))
}

// StreamAnalyticsJobId returns a reference to field stream_analytics_job_id of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) StreamAnalyticsJobId() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("stream_analytics_job_id"))
}

// Table returns a reference to field table of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("table"))
}

// TokenUserDisplayName returns a reference to field token_user_display_name of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) TokenUserDisplayName() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("token_user_display_name"))
}

// TokenUserPrincipalName returns a reference to field token_user_principal_name of azurerm_stream_analytics_output_powerbi.
func (saop streamAnalyticsOutputPowerbiAttributes) TokenUserPrincipalName() terra.StringValue {
	return terra.ReferenceAsString(saop.ref.Append("token_user_principal_name"))
}

func (saop streamAnalyticsOutputPowerbiAttributes) Timeouts() streamanalyticsoutputpowerbi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsoutputpowerbi.TimeoutsAttributes](saop.ref.Append("timeouts"))
}

type streamAnalyticsOutputPowerbiState struct {
	Dataset                string                                      `json:"dataset"`
	GroupId                string                                      `json:"group_id"`
	GroupName              string                                      `json:"group_name"`
	Id                     string                                      `json:"id"`
	Name                   string                                      `json:"name"`
	StreamAnalyticsJobId   string                                      `json:"stream_analytics_job_id"`
	Table                  string                                      `json:"table"`
	TokenUserDisplayName   string                                      `json:"token_user_display_name"`
	TokenUserPrincipalName string                                      `json:"token_user_principal_name"`
	Timeouts               *streamanalyticsoutputpowerbi.TimeoutsState `json:"timeouts"`
}
