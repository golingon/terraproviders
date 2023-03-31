// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computenodegroup "github.com/golingon/terraproviders/google/4.59.0/computenodegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeNodeGroup(name string, args ComputeNodeGroupArgs) *ComputeNodeGroup {
	return &ComputeNodeGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNodeGroup)(nil)

type ComputeNodeGroup struct {
	Name  string
	Args  ComputeNodeGroupArgs
	state *computeNodeGroupState
}

func (cng *ComputeNodeGroup) Type() string {
	return "google_compute_node_group"
}

func (cng *ComputeNodeGroup) LocalName() string {
	return cng.Name
}

func (cng *ComputeNodeGroup) Configuration() interface{} {
	return cng.Args
}

func (cng *ComputeNodeGroup) Attributes() computeNodeGroupAttributes {
	return computeNodeGroupAttributes{ref: terra.ReferenceResource(cng)}
}

func (cng *ComputeNodeGroup) ImportState(av io.Reader) error {
	cng.state = &computeNodeGroupState{}
	if err := json.NewDecoder(av).Decode(cng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cng.Type(), cng.LocalName(), err)
	}
	return nil
}

func (cng *ComputeNodeGroup) State() (*computeNodeGroupState, bool) {
	return cng.state, cng.state != nil
}

func (cng *ComputeNodeGroup) StateMust() *computeNodeGroupState {
	if cng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cng.Type(), cng.LocalName()))
	}
	return cng.state
}

func (cng *ComputeNodeGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cng)
}

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
	// DependsOn contains resources that ComputeNodeGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeNodeGroupAttributes struct {
	ref terra.Reference
}

func (cng computeNodeGroupAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("creation_timestamp"))
}

func (cng computeNodeGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("description"))
}

func (cng computeNodeGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("id"))
}

func (cng computeNodeGroupAttributes) InitialSize() terra.NumberValue {
	return terra.ReferenceNumber(cng.ref.Append("initial_size"))
}

func (cng computeNodeGroupAttributes) MaintenancePolicy() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("maintenance_policy"))
}

func (cng computeNodeGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("name"))
}

func (cng computeNodeGroupAttributes) NodeTemplate() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("node_template"))
}

func (cng computeNodeGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("project"))
}

func (cng computeNodeGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("self_link"))
}

func (cng computeNodeGroupAttributes) Size() terra.NumberValue {
	return terra.ReferenceNumber(cng.ref.Append("size"))
}

func (cng computeNodeGroupAttributes) Zone() terra.StringValue {
	return terra.ReferenceString(cng.ref.Append("zone"))
}

func (cng computeNodeGroupAttributes) AutoscalingPolicy() terra.ListValue[computenodegroup.AutoscalingPolicyAttributes] {
	return terra.ReferenceList[computenodegroup.AutoscalingPolicyAttributes](cng.ref.Append("autoscaling_policy"))
}

func (cng computeNodeGroupAttributes) MaintenanceWindow() terra.ListValue[computenodegroup.MaintenanceWindowAttributes] {
	return terra.ReferenceList[computenodegroup.MaintenanceWindowAttributes](cng.ref.Append("maintenance_window"))
}

func (cng computeNodeGroupAttributes) ShareSettings() terra.ListValue[computenodegroup.ShareSettingsAttributes] {
	return terra.ReferenceList[computenodegroup.ShareSettingsAttributes](cng.ref.Append("share_settings"))
}

func (cng computeNodeGroupAttributes) Timeouts() computenodegroup.TimeoutsAttributes {
	return terra.ReferenceSingle[computenodegroup.TimeoutsAttributes](cng.ref.Append("timeouts"))
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
