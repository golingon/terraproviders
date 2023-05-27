// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	tpunode "github.com/golingon/terraproviders/google/4.66.0/tpunode"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTpuNode creates a new instance of [TpuNode].
func NewTpuNode(name string, args TpuNodeArgs) *TpuNode {
	return &TpuNode{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TpuNode)(nil)

// TpuNode represents the Terraform resource google_tpu_node.
type TpuNode struct {
	Name      string
	Args      TpuNodeArgs
	state     *tpuNodeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TpuNode].
func (tn *TpuNode) Type() string {
	return "google_tpu_node"
}

// LocalName returns the local name for [TpuNode].
func (tn *TpuNode) LocalName() string {
	return tn.Name
}

// Configuration returns the configuration (args) for [TpuNode].
func (tn *TpuNode) Configuration() interface{} {
	return tn.Args
}

// DependOn is used for other resources to depend on [TpuNode].
func (tn *TpuNode) DependOn() terra.Reference {
	return terra.ReferenceResource(tn)
}

// Dependencies returns the list of resources [TpuNode] depends_on.
func (tn *TpuNode) Dependencies() terra.Dependencies {
	return tn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TpuNode].
func (tn *TpuNode) LifecycleManagement() *terra.Lifecycle {
	return tn.Lifecycle
}

// Attributes returns the attributes for [TpuNode].
func (tn *TpuNode) Attributes() tpuNodeAttributes {
	return tpuNodeAttributes{ref: terra.ReferenceResource(tn)}
}

// ImportState imports the given attribute values into [TpuNode]'s state.
func (tn *TpuNode) ImportState(av io.Reader) error {
	tn.state = &tpuNodeState{}
	if err := json.NewDecoder(av).Decode(tn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tn.Type(), tn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TpuNode] has state.
func (tn *TpuNode) State() (*tpuNodeState, bool) {
	return tn.state, tn.state != nil
}

// StateMust returns the state for [TpuNode]. Panics if the state is nil.
func (tn *TpuNode) StateMust() *tpuNodeState {
	if tn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tn.Type(), tn.LocalName()))
	}
	return tn.state
}

// TpuNodeArgs contains the configurations for google_tpu_node.
type TpuNodeArgs struct {
	// AcceleratorType: string, required
	AcceleratorType terra.StringValue `hcl:"accelerator_type,attr" validate:"required"`
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TensorflowVersion: string, required
	TensorflowVersion terra.StringValue `hcl:"tensorflow_version,attr" validate:"required"`
	// UseServiceNetworking: bool, optional
	UseServiceNetworking terra.BoolValue `hcl:"use_service_networking,attr"`
	// Zone: string, optional
	Zone terra.StringValue `hcl:"zone,attr"`
	// NetworkEndpoints: min=0
	NetworkEndpoints []tpunode.NetworkEndpoints `hcl:"network_endpoints,block" validate:"min=0"`
	// SchedulingConfig: optional
	SchedulingConfig *tpunode.SchedulingConfig `hcl:"scheduling_config,block"`
	// Timeouts: optional
	Timeouts *tpunode.Timeouts `hcl:"timeouts,block"`
}
type tpuNodeAttributes struct {
	ref terra.Reference
}

// AcceleratorType returns a reference to field accelerator_type of google_tpu_node.
func (tn tpuNodeAttributes) AcceleratorType() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("accelerator_type"))
}

// CidrBlock returns a reference to field cidr_block of google_tpu_node.
func (tn tpuNodeAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("cidr_block"))
}

// Description returns a reference to field description of google_tpu_node.
func (tn tpuNodeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("description"))
}

// Id returns a reference to field id of google_tpu_node.
func (tn tpuNodeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("id"))
}

// Labels returns a reference to field labels of google_tpu_node.
func (tn tpuNodeAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tn.ref.Append("labels"))
}

// Name returns a reference to field name of google_tpu_node.
func (tn tpuNodeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("name"))
}

// Network returns a reference to field network of google_tpu_node.
func (tn tpuNodeAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("network"))
}

// Project returns a reference to field project of google_tpu_node.
func (tn tpuNodeAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("project"))
}

// ServiceAccount returns a reference to field service_account of google_tpu_node.
func (tn tpuNodeAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("service_account"))
}

// TensorflowVersion returns a reference to field tensorflow_version of google_tpu_node.
func (tn tpuNodeAttributes) TensorflowVersion() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("tensorflow_version"))
}

// UseServiceNetworking returns a reference to field use_service_networking of google_tpu_node.
func (tn tpuNodeAttributes) UseServiceNetworking() terra.BoolValue {
	return terra.ReferenceAsBool(tn.ref.Append("use_service_networking"))
}

// Zone returns a reference to field zone of google_tpu_node.
func (tn tpuNodeAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(tn.ref.Append("zone"))
}

func (tn tpuNodeAttributes) NetworkEndpoints() terra.ListValue[tpunode.NetworkEndpointsAttributes] {
	return terra.ReferenceAsList[tpunode.NetworkEndpointsAttributes](tn.ref.Append("network_endpoints"))
}

func (tn tpuNodeAttributes) SchedulingConfig() terra.ListValue[tpunode.SchedulingConfigAttributes] {
	return terra.ReferenceAsList[tpunode.SchedulingConfigAttributes](tn.ref.Append("scheduling_config"))
}

func (tn tpuNodeAttributes) Timeouts() tpunode.TimeoutsAttributes {
	return terra.ReferenceAsSingle[tpunode.TimeoutsAttributes](tn.ref.Append("timeouts"))
}

type tpuNodeState struct {
	AcceleratorType      string                          `json:"accelerator_type"`
	CidrBlock            string                          `json:"cidr_block"`
	Description          string                          `json:"description"`
	Id                   string                          `json:"id"`
	Labels               map[string]string               `json:"labels"`
	Name                 string                          `json:"name"`
	Network              string                          `json:"network"`
	Project              string                          `json:"project"`
	ServiceAccount       string                          `json:"service_account"`
	TensorflowVersion    string                          `json:"tensorflow_version"`
	UseServiceNetworking bool                            `json:"use_service_networking"`
	Zone                 string                          `json:"zone"`
	NetworkEndpoints     []tpunode.NetworkEndpointsState `json:"network_endpoints"`
	SchedulingConfig     []tpunode.SchedulingConfigState `json:"scheduling_config"`
	Timeouts             *tpunode.TimeoutsState          `json:"timeouts"`
}
