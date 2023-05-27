// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsreferenceinputmssql "github.com/golingon/terraproviders/azurerm/3.58.0/streamanalyticsreferenceinputmssql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsReferenceInputMssql creates a new instance of [StreamAnalyticsReferenceInputMssql].
func NewStreamAnalyticsReferenceInputMssql(name string, args StreamAnalyticsReferenceInputMssqlArgs) *StreamAnalyticsReferenceInputMssql {
	return &StreamAnalyticsReferenceInputMssql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsReferenceInputMssql)(nil)

// StreamAnalyticsReferenceInputMssql represents the Terraform resource azurerm_stream_analytics_reference_input_mssql.
type StreamAnalyticsReferenceInputMssql struct {
	Name      string
	Args      StreamAnalyticsReferenceInputMssqlArgs
	state     *streamAnalyticsReferenceInputMssqlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsReferenceInputMssql].
func (sarim *StreamAnalyticsReferenceInputMssql) Type() string {
	return "azurerm_stream_analytics_reference_input_mssql"
}

// LocalName returns the local name for [StreamAnalyticsReferenceInputMssql].
func (sarim *StreamAnalyticsReferenceInputMssql) LocalName() string {
	return sarim.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsReferenceInputMssql].
func (sarim *StreamAnalyticsReferenceInputMssql) Configuration() interface{} {
	return sarim.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsReferenceInputMssql].
func (sarim *StreamAnalyticsReferenceInputMssql) DependOn() terra.Reference {
	return terra.ReferenceResource(sarim)
}

// Dependencies returns the list of resources [StreamAnalyticsReferenceInputMssql] depends_on.
func (sarim *StreamAnalyticsReferenceInputMssql) Dependencies() terra.Dependencies {
	return sarim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsReferenceInputMssql].
func (sarim *StreamAnalyticsReferenceInputMssql) LifecycleManagement() *terra.Lifecycle {
	return sarim.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsReferenceInputMssql].
func (sarim *StreamAnalyticsReferenceInputMssql) Attributes() streamAnalyticsReferenceInputMssqlAttributes {
	return streamAnalyticsReferenceInputMssqlAttributes{ref: terra.ReferenceResource(sarim)}
}

// ImportState imports the given attribute values into [StreamAnalyticsReferenceInputMssql]'s state.
func (sarim *StreamAnalyticsReferenceInputMssql) ImportState(av io.Reader) error {
	sarim.state = &streamAnalyticsReferenceInputMssqlState{}
	if err := json.NewDecoder(av).Decode(sarim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sarim.Type(), sarim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsReferenceInputMssql] has state.
func (sarim *StreamAnalyticsReferenceInputMssql) State() (*streamAnalyticsReferenceInputMssqlState, bool) {
	return sarim.state, sarim.state != nil
}

// StateMust returns the state for [StreamAnalyticsReferenceInputMssql]. Panics if the state is nil.
func (sarim *StreamAnalyticsReferenceInputMssql) StateMust() *streamAnalyticsReferenceInputMssqlState {
	if sarim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sarim.Type(), sarim.LocalName()))
	}
	return sarim.state
}

// StreamAnalyticsReferenceInputMssqlArgs contains the configurations for azurerm_stream_analytics_reference_input_mssql.
type StreamAnalyticsReferenceInputMssqlArgs struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// DeltaSnapshotQuery: string, optional
	DeltaSnapshotQuery terra.StringValue `hcl:"delta_snapshot_query,attr"`
	// FullSnapshotQuery: string, required
	FullSnapshotQuery terra.StringValue `hcl:"full_snapshot_query,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// RefreshIntervalDuration: string, optional
	RefreshIntervalDuration terra.StringValue `hcl:"refresh_interval_duration,attr"`
	// RefreshType: string, required
	RefreshType terra.StringValue `hcl:"refresh_type,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// StreamAnalyticsJobName: string, required
	StreamAnalyticsJobName terra.StringValue `hcl:"stream_analytics_job_name,attr" validate:"required"`
	// Table: string, optional
	Table terra.StringValue `hcl:"table,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsreferenceinputmssql.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsReferenceInputMssqlAttributes struct {
	ref terra.Reference
}

// Database returns a reference to field database of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("database"))
}

// DeltaSnapshotQuery returns a reference to field delta_snapshot_query of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) DeltaSnapshotQuery() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("delta_snapshot_query"))
}

// FullSnapshotQuery returns a reference to field full_snapshot_query of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) FullSnapshotQuery() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("full_snapshot_query"))
}

// Id returns a reference to field id of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("name"))
}

// Password returns a reference to field password of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("password"))
}

// RefreshIntervalDuration returns a reference to field refresh_interval_duration of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) RefreshIntervalDuration() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("refresh_interval_duration"))
}

// RefreshType returns a reference to field refresh_type of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) RefreshType() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("refresh_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("resource_group_name"))
}

// Server returns a reference to field server of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("server"))
}

// StreamAnalyticsJobName returns a reference to field stream_analytics_job_name of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) StreamAnalyticsJobName() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("stream_analytics_job_name"))
}

// Table returns a reference to field table of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Table() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("table"))
}

// Username returns a reference to field username of azurerm_stream_analytics_reference_input_mssql.
func (sarim streamAnalyticsReferenceInputMssqlAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(sarim.ref.Append("username"))
}

func (sarim streamAnalyticsReferenceInputMssqlAttributes) Timeouts() streamanalyticsreferenceinputmssql.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsreferenceinputmssql.TimeoutsAttributes](sarim.ref.Append("timeouts"))
}

type streamAnalyticsReferenceInputMssqlState struct {
	Database                string                                            `json:"database"`
	DeltaSnapshotQuery      string                                            `json:"delta_snapshot_query"`
	FullSnapshotQuery       string                                            `json:"full_snapshot_query"`
	Id                      string                                            `json:"id"`
	Name                    string                                            `json:"name"`
	Password                string                                            `json:"password"`
	RefreshIntervalDuration string                                            `json:"refresh_interval_duration"`
	RefreshType             string                                            `json:"refresh_type"`
	ResourceGroupName       string                                            `json:"resource_group_name"`
	Server                  string                                            `json:"server"`
	StreamAnalyticsJobName  string                                            `json:"stream_analytics_job_name"`
	Table                   string                                            `json:"table"`
	Username                string                                            `json:"username"`
	Timeouts                *streamanalyticsreferenceinputmssql.TimeoutsState `json:"timeouts"`
}
