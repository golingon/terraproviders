// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	stackhcicluster "github.com/golingon/terraproviders/azurerm/3.65.0/stackhcicluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStackHciCluster creates a new instance of [StackHciCluster].
func NewStackHciCluster(name string, args StackHciClusterArgs) *StackHciCluster {
	return &StackHciCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StackHciCluster)(nil)

// StackHciCluster represents the Terraform resource azurerm_stack_hci_cluster.
type StackHciCluster struct {
	Name      string
	Args      StackHciClusterArgs
	state     *stackHciClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StackHciCluster].
func (shc *StackHciCluster) Type() string {
	return "azurerm_stack_hci_cluster"
}

// LocalName returns the local name for [StackHciCluster].
func (shc *StackHciCluster) LocalName() string {
	return shc.Name
}

// Configuration returns the configuration (args) for [StackHciCluster].
func (shc *StackHciCluster) Configuration() interface{} {
	return shc.Args
}

// DependOn is used for other resources to depend on [StackHciCluster].
func (shc *StackHciCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(shc)
}

// Dependencies returns the list of resources [StackHciCluster] depends_on.
func (shc *StackHciCluster) Dependencies() terra.Dependencies {
	return shc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StackHciCluster].
func (shc *StackHciCluster) LifecycleManagement() *terra.Lifecycle {
	return shc.Lifecycle
}

// Attributes returns the attributes for [StackHciCluster].
func (shc *StackHciCluster) Attributes() stackHciClusterAttributes {
	return stackHciClusterAttributes{ref: terra.ReferenceResource(shc)}
}

// ImportState imports the given attribute values into [StackHciCluster]'s state.
func (shc *StackHciCluster) ImportState(av io.Reader) error {
	shc.state = &stackHciClusterState{}
	if err := json.NewDecoder(av).Decode(shc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", shc.Type(), shc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StackHciCluster] has state.
func (shc *StackHciCluster) State() (*stackHciClusterState, bool) {
	return shc.state, shc.state != nil
}

// StateMust returns the state for [StackHciCluster]. Panics if the state is nil.
func (shc *StackHciCluster) StateMust() *stackHciClusterState {
	if shc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", shc.Type(), shc.LocalName()))
	}
	return shc.state
}

// StackHciClusterArgs contains the configurations for azurerm_stack_hci_cluster.
type StackHciClusterArgs struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TenantId: string, optional
	TenantId terra.StringValue `hcl:"tenant_id,attr"`
	// Timeouts: optional
	Timeouts *stackhcicluster.Timeouts `hcl:"timeouts,block"`
}
type stackHciClusterAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("client_id"))
}

// Id returns a reference to field id of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](shc.ref.Append("tags"))
}

// TenantId returns a reference to field tenant_id of azurerm_stack_hci_cluster.
func (shc stackHciClusterAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(shc.ref.Append("tenant_id"))
}

func (shc stackHciClusterAttributes) Timeouts() stackhcicluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[stackhcicluster.TimeoutsAttributes](shc.ref.Append("timeouts"))
}

type stackHciClusterState struct {
	ClientId          string                         `json:"client_id"`
	Id                string                         `json:"id"`
	Location          string                         `json:"location"`
	Name              string                         `json:"name"`
	ResourceGroupName string                         `json:"resource_group_name"`
	Tags              map[string]string              `json:"tags"`
	TenantId          string                         `json:"tenant_id"`
	Timeouts          *stackhcicluster.TimeoutsState `json:"timeouts"`
}
