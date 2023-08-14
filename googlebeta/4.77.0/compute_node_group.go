// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computenodegroup "github.com/golingon/terraproviders/googlebeta/4.77.0/computenodegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNodeGroup creates a new instance of [ComputeNodeGroup].
func NewComputeNodeGroup(name string, args ComputeNodeGroupArgs) *ComputeNodeGroup {
	return &ComputeNodeGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNodeGroup)(nil)

// ComputeNodeGroup represents the Terraform resource google_compute_node_group.
type ComputeNodeGroup struct {
	Name      string
	Args      ComputeNodeGroupArgs
	state     *computeNodeGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNodeGroup].
func (cng *ComputeNodeGroup) Type() string {
	return "google_compute_node_group"
}

// LocalName returns the local name for [ComputeNodeGroup].
func (cng *ComputeNodeGroup) LocalName() string {
	return cng.Name
}

// Configuration returns the configuration (args) for [ComputeNodeGroup].
func (cng *ComputeNodeGroup) Configuration() interface{} {
	return cng.Args
}

// DependOn is used for other resources to depend on [ComputeNodeGroup].
func (cng *ComputeNodeGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cng)
}

// Dependencies returns the list of resources [ComputeNodeGroup] depends_on.
func (cng *ComputeNodeGroup) Dependencies() terra.Dependencies {
	return cng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNodeGroup].
func (cng *ComputeNodeGroup) LifecycleManagement() *terra.Lifecycle {
	return cng.Lifecycle
}

// Attributes returns the attributes for [ComputeNodeGroup].
func (cng *ComputeNodeGroup) Attributes() computeNodeGroupAttributes {
	return computeNodeGroupAttributes{ref: terra.ReferenceResource(cng)}
}

// ImportState imports the given attribute values into [ComputeNodeGroup]'s state.
func (cng *ComputeNodeGroup) ImportState(av io.Reader) error {
	cng.state = &computeNodeGroupState{}
	if err := json.NewDecoder(av).Decode(cng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cng.Type(), cng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNodeGroup] has state.
func (cng *ComputeNodeGroup) State() (*computeNodeGroupState, bool) {
	return cng.state, cng.state != nil
}

// StateMust returns the state for [ComputeNodeGroup]. Panics if the state is nil.
func (cng *ComputeNodeGroup) StateMust() *computeNodeGroupState {
	if cng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cng.Type(), cng.LocalName()))
	}
	return cng.state
}

// ComputeNodeGroupArgs contains the configurations for google_compute_node_group.
type ComputeNodeGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InitialSize: number, optional
	InitialSize terra.NumberValue `hcl:"initial_size,attr"`
	// MaintenancePolicy: string, optional
	MaintenancePolicy terra.StringValue `hcl:"maintenance_policy,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NodeTemplate: string, required
	NodeTemplate terra.StringValue `hcl:"node_template,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Size: number, optional
	Size terra.NumberValue `hcl:"size,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// AutoscalingPolicy: optional
	AutoscalingPolicy *computenodegroup.AutoscalingPolicy `hcl:"autoscaling_policy,block"`
	// MaintenanceWindow: optional
	MaintenanceWindow *computenodegroup.MaintenanceWindow `hcl:"maintenance_window,block"`
	// ShareSettings: optional
	ShareSettings *computenodegroup.ShareSettings `hcl:"share_settings,block"`
	// Timeouts: optional
	Timeouts *computenodegroup.Timeouts `hcl:"timeouts,block"`
}
type computeNodeGroupAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_node_group.
func (cng computeNodeGroupAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_node_group.
func (cng computeNodeGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_node_group.
func (cng computeNodeGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("id"))
}

// InitialSize returns a reference to field initial_size of google_compute_node_group.
func (cng computeNodeGroupAttributes) InitialSize() terra.NumberValue {
	return terra.ReferenceAsNumber(cng.ref.Append("initial_size"))
}

// MaintenancePolicy returns a reference to field maintenance_policy of google_compute_node_group.
func (cng computeNodeGroupAttributes) MaintenancePolicy() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("maintenance_policy"))
}

// Name returns a reference to field name of google_compute_node_group.
func (cng computeNodeGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("name"))
}

// NodeTemplate returns a reference to field node_template of google_compute_node_group.
func (cng computeNodeGroupAttributes) NodeTemplate() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("node_template"))
}

// Project returns a reference to field project of google_compute_node_group.
func (cng computeNodeGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_node_group.
func (cng computeNodeGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("self_link"))
}

// Size returns a reference to field size of google_compute_node_group.
func (cng computeNodeGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(cng.ref.Append("size"))
}

// Zone returns a reference to field zone of google_compute_node_group.
func (cng computeNodeGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(cng.ref.Append("zone"))
}

func (cng computeNodeGroupAttributes) AutoscalingPolicy() terra.ListValue[computenodegroup.AutoscalingPolicyAttributes] {
	return terra.ReferenceAsList[computenodegroup.AutoscalingPolicyAttributes](cng.ref.Append("autoscaling_policy"))
}

func (cng computeNodeGroupAttributes) MaintenanceWindow() terra.ListValue[computenodegroup.MaintenanceWindowAttributes] {
	return terra.ReferenceAsList[computenodegroup.MaintenanceWindowAttributes](cng.ref.Append("maintenance_window"))
}

func (cng computeNodeGroupAttributes) ShareSettings() terra.ListValue[computenodegroup.ShareSettingsAttributes] {
	return terra.ReferenceAsList[computenodegroup.ShareSettingsAttributes](cng.ref.Append("share_settings"))
}

func (cng computeNodeGroupAttributes) Timeouts() computenodegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenodegroup.TimeoutsAttributes](cng.ref.Append("timeouts"))
}

type computeNodeGroupState struct {
	CreationTimestamp string                                    `json:"creation_timestamp"`
	Description       string                                    `json:"description"`
	Id                string                                    `json:"id"`
	InitialSize       float64                                   `json:"initial_size"`
	MaintenancePolicy string                                    `json:"maintenance_policy"`
	Name              string                                    `json:"name"`
	NodeTemplate      string                                    `json:"node_template"`
	Project           string                                    `json:"project"`
	SelfLink          string                                    `json:"self_link"`
	Size              float64                                   `json:"size"`
	Zone              string                                    `json:"zone"`
	AutoscalingPolicy []computenodegroup.AutoscalingPolicyState `json:"autoscaling_policy"`
	MaintenanceWindow []computenodegroup.MaintenanceWindowState `json:"maintenance_window"`
	ShareSettings     []computenodegroup.ShareSettingsState     `json:"share_settings"`
	Timeouts          *computenodegroup.TimeoutsState           `json:"timeouts"`
}
