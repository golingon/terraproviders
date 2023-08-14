// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	vmwareenginecluster "github.com/golingon/terraproviders/googlebeta/4.77.0/vmwareenginecluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwareengineCluster creates a new instance of [VmwareengineCluster].
func NewVmwareengineCluster(name string, args VmwareengineClusterArgs) *VmwareengineCluster {
	return &VmwareengineCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareengineCluster)(nil)

// VmwareengineCluster represents the Terraform resource google_vmwareengine_cluster.
type VmwareengineCluster struct {
	Name      string
	Args      VmwareengineClusterArgs
	state     *vmwareengineClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareengineCluster].
func (vc *VmwareengineCluster) Type() string {
	return "google_vmwareengine_cluster"
}

// LocalName returns the local name for [VmwareengineCluster].
func (vc *VmwareengineCluster) LocalName() string {
	return vc.Name
}

// Configuration returns the configuration (args) for [VmwareengineCluster].
func (vc *VmwareengineCluster) Configuration() interface{} {
	return vc.Args
}

// DependOn is used for other resources to depend on [VmwareengineCluster].
func (vc *VmwareengineCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(vc)
}

// Dependencies returns the list of resources [VmwareengineCluster] depends_on.
func (vc *VmwareengineCluster) Dependencies() terra.Dependencies {
	return vc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareengineCluster].
func (vc *VmwareengineCluster) LifecycleManagement() *terra.Lifecycle {
	return vc.Lifecycle
}

// Attributes returns the attributes for [VmwareengineCluster].
func (vc *VmwareengineCluster) Attributes() vmwareengineClusterAttributes {
	return vmwareengineClusterAttributes{ref: terra.ReferenceResource(vc)}
}

// ImportState imports the given attribute values into [VmwareengineCluster]'s state.
func (vc *VmwareengineCluster) ImportState(av io.Reader) error {
	vc.state = &vmwareengineClusterState{}
	if err := json.NewDecoder(av).Decode(vc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vc.Type(), vc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareengineCluster] has state.
func (vc *VmwareengineCluster) State() (*vmwareengineClusterState, bool) {
	return vc.state, vc.state != nil
}

// StateMust returns the state for [VmwareengineCluster]. Panics if the state is nil.
func (vc *VmwareengineCluster) StateMust() *vmwareengineClusterState {
	if vc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vc.Type(), vc.LocalName()))
	}
	return vc.state
}

// VmwareengineClusterArgs contains the configurations for google_vmwareengine_cluster.
type VmwareengineClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// NodeTypeConfigs: min=0
	NodeTypeConfigs []vmwareenginecluster.NodeTypeConfigs `hcl:"node_type_configs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *vmwareenginecluster.Timeouts `hcl:"timeouts,block"`
}
type vmwareengineClusterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_vmwareengine_cluster.
func (vc vmwareengineClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("id"))
}

// Management returns a reference to field management of google_vmwareengine_cluster.
func (vc vmwareengineClusterAttributes) Management() terra.BoolValue {
	return terra.ReferenceAsBool(vc.ref.Append("management"))
}

// Name returns a reference to field name of google_vmwareengine_cluster.
func (vc vmwareengineClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("name"))
}

// Parent returns a reference to field parent of google_vmwareengine_cluster.
func (vc vmwareengineClusterAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("parent"))
}

// State returns a reference to field state of google_vmwareengine_cluster.
func (vc vmwareengineClusterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("state"))
}

// Uid returns a reference to field uid of google_vmwareengine_cluster.
func (vc vmwareengineClusterAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("uid"))
}

func (vc vmwareengineClusterAttributes) NodeTypeConfigs() terra.SetValue[vmwareenginecluster.NodeTypeConfigsAttributes] {
	return terra.ReferenceAsSet[vmwareenginecluster.NodeTypeConfigsAttributes](vc.ref.Append("node_type_configs"))
}

func (vc vmwareengineClusterAttributes) Timeouts() vmwareenginecluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwareenginecluster.TimeoutsAttributes](vc.ref.Append("timeouts"))
}

type vmwareengineClusterState struct {
	Id              string                                     `json:"id"`
	Management      bool                                       `json:"management"`
	Name            string                                     `json:"name"`
	Parent          string                                     `json:"parent"`
	State           string                                     `json:"state"`
	Uid             string                                     `json:"uid"`
	NodeTypeConfigs []vmwareenginecluster.NodeTypeConfigsState `json:"node_type_configs"`
	Timeouts        *vmwareenginecluster.TimeoutsState         `json:"timeouts"`
}
