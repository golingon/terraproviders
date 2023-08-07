// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	monitoringgroup "github.com/golingon/terraproviders/google/4.76.0/monitoringgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitoringGroup creates a new instance of [MonitoringGroup].
func NewMonitoringGroup(name string, args MonitoringGroupArgs) *MonitoringGroup {
	return &MonitoringGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitoringGroup)(nil)

// MonitoringGroup represents the Terraform resource google_monitoring_group.
type MonitoringGroup struct {
	Name      string
	Args      MonitoringGroupArgs
	state     *monitoringGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitoringGroup].
func (mg *MonitoringGroup) Type() string {
	return "google_monitoring_group"
}

// LocalName returns the local name for [MonitoringGroup].
func (mg *MonitoringGroup) LocalName() string {
	return mg.Name
}

// Configuration returns the configuration (args) for [MonitoringGroup].
func (mg *MonitoringGroup) Configuration() interface{} {
	return mg.Args
}

// DependOn is used for other resources to depend on [MonitoringGroup].
func (mg *MonitoringGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mg)
}

// Dependencies returns the list of resources [MonitoringGroup] depends_on.
func (mg *MonitoringGroup) Dependencies() terra.Dependencies {
	return mg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitoringGroup].
func (mg *MonitoringGroup) LifecycleManagement() *terra.Lifecycle {
	return mg.Lifecycle
}

// Attributes returns the attributes for [MonitoringGroup].
func (mg *MonitoringGroup) Attributes() monitoringGroupAttributes {
	return monitoringGroupAttributes{ref: terra.ReferenceResource(mg)}
}

// ImportState imports the given attribute values into [MonitoringGroup]'s state.
func (mg *MonitoringGroup) ImportState(av io.Reader) error {
	mg.state = &monitoringGroupState{}
	if err := json.NewDecoder(av).Decode(mg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mg.Type(), mg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitoringGroup] has state.
func (mg *MonitoringGroup) State() (*monitoringGroupState, bool) {
	return mg.state, mg.state != nil
}

// StateMust returns the state for [MonitoringGroup]. Panics if the state is nil.
func (mg *MonitoringGroup) StateMust() *monitoringGroupState {
	if mg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mg.Type(), mg.LocalName()))
	}
	return mg.state
}

// MonitoringGroupArgs contains the configurations for google_monitoring_group.
type MonitoringGroupArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Filter: string, required
	Filter terra.StringValue `hcl:"filter,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsCluster: bool, optional
	IsCluster terra.BoolValue `hcl:"is_cluster,attr"`
	// ParentName: string, optional
	ParentName terra.StringValue `hcl:"parent_name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *monitoringgroup.Timeouts `hcl:"timeouts,block"`
}
type monitoringGroupAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_monitoring_group.
func (mg monitoringGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("display_name"))
}

// Filter returns a reference to field filter of google_monitoring_group.
func (mg monitoringGroupAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("filter"))
}

// Id returns a reference to field id of google_monitoring_group.
func (mg monitoringGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("id"))
}

// IsCluster returns a reference to field is_cluster of google_monitoring_group.
func (mg monitoringGroupAttributes) IsCluster() terra.BoolValue {
	return terra.ReferenceAsBool(mg.ref.Append("is_cluster"))
}

// Name returns a reference to field name of google_monitoring_group.
func (mg monitoringGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("name"))
}

// ParentName returns a reference to field parent_name of google_monitoring_group.
func (mg monitoringGroupAttributes) ParentName() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("parent_name"))
}

// Project returns a reference to field project of google_monitoring_group.
func (mg monitoringGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mg.ref.Append("project"))
}

func (mg monitoringGroupAttributes) Timeouts() monitoringgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoringgroup.TimeoutsAttributes](mg.ref.Append("timeouts"))
}

type monitoringGroupState struct {
	DisplayName string                         `json:"display_name"`
	Filter      string                         `json:"filter"`
	Id          string                         `json:"id"`
	IsCluster   bool                           `json:"is_cluster"`
	Name        string                         `json:"name"`
	ParentName  string                         `json:"parent_name"`
	Project     string                         `json:"project"`
	Timeouts    *monitoringgroup.TimeoutsState `json:"timeouts"`
}
