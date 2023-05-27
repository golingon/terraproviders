// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vmwarecluster "github.com/golingon/terraproviders/azurerm/3.58.0/vmwarecluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwareCluster creates a new instance of [VmwareCluster].
func NewVmwareCluster(name string, args VmwareClusterArgs) *VmwareCluster {
	return &VmwareCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareCluster)(nil)

// VmwareCluster represents the Terraform resource azurerm_vmware_cluster.
type VmwareCluster struct {
	Name      string
	Args      VmwareClusterArgs
	state     *vmwareClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareCluster].
func (vc *VmwareCluster) Type() string {
	return "azurerm_vmware_cluster"
}

// LocalName returns the local name for [VmwareCluster].
func (vc *VmwareCluster) LocalName() string {
	return vc.Name
}

// Configuration returns the configuration (args) for [VmwareCluster].
func (vc *VmwareCluster) Configuration() interface{} {
	return vc.Args
}

// DependOn is used for other resources to depend on [VmwareCluster].
func (vc *VmwareCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(vc)
}

// Dependencies returns the list of resources [VmwareCluster] depends_on.
func (vc *VmwareCluster) Dependencies() terra.Dependencies {
	return vc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareCluster].
func (vc *VmwareCluster) LifecycleManagement() *terra.Lifecycle {
	return vc.Lifecycle
}

// Attributes returns the attributes for [VmwareCluster].
func (vc *VmwareCluster) Attributes() vmwareClusterAttributes {
	return vmwareClusterAttributes{ref: terra.ReferenceResource(vc)}
}

// ImportState imports the given attribute values into [VmwareCluster]'s state.
func (vc *VmwareCluster) ImportState(av io.Reader) error {
	vc.state = &vmwareClusterState{}
	if err := json.NewDecoder(av).Decode(vc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vc.Type(), vc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareCluster] has state.
func (vc *VmwareCluster) State() (*vmwareClusterState, bool) {
	return vc.state, vc.state != nil
}

// StateMust returns the state for [VmwareCluster]. Panics if the state is nil.
func (vc *VmwareCluster) StateMust() *vmwareClusterState {
	if vc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vc.Type(), vc.LocalName()))
	}
	return vc.state
}

// VmwareClusterArgs contains the configurations for azurerm_vmware_cluster.
type VmwareClusterArgs struct {
	// ClusterNodeCount: number, required
	ClusterNodeCount terra.NumberValue `hcl:"cluster_node_count,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// VmwareCloudId: string, required
	VmwareCloudId terra.StringValue `hcl:"vmware_cloud_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vmwarecluster.Timeouts `hcl:"timeouts,block"`
}
type vmwareClusterAttributes struct {
	ref terra.Reference
}

// ClusterNodeCount returns a reference to field cluster_node_count of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) ClusterNodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("cluster_node_count"))
}

// ClusterNumber returns a reference to field cluster_number of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) ClusterNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(vc.ref.Append("cluster_number"))
}

// Hosts returns a reference to field hosts of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) Hosts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vc.ref.Append("hosts"))
}

// Id returns a reference to field id of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("name"))
}

// SkuName returns a reference to field sku_name of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("sku_name"))
}

// VmwareCloudId returns a reference to field vmware_cloud_id of azurerm_vmware_cluster.
func (vc vmwareClusterAttributes) VmwareCloudId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vmware_cloud_id"))
}

func (vc vmwareClusterAttributes) Timeouts() vmwarecluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwarecluster.TimeoutsAttributes](vc.ref.Append("timeouts"))
}

type vmwareClusterState struct {
	ClusterNodeCount float64                      `json:"cluster_node_count"`
	ClusterNumber    float64                      `json:"cluster_number"`
	Hosts            []string                     `json:"hosts"`
	Id               string                       `json:"id"`
	Name             string                       `json:"name"`
	SkuName          string                       `json:"sku_name"`
	VmwareCloudId    string                       `json:"vmware_cloud_id"`
	Timeouts         *vmwarecluster.TimeoutsState `json:"timeouts"`
}
