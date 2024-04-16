// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_kubernetes_fleet_update_run

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_kubernetes_fleet_update_run.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermKubernetesFleetUpdateRunState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (akfur *Resource) Type() string {
	return "azurerm_kubernetes_fleet_update_run"
}

// LocalName returns the local name for [Resource].
func (akfur *Resource) LocalName() string {
	return akfur.Name
}

// Configuration returns the configuration (args) for [Resource].
func (akfur *Resource) Configuration() interface{} {
	return akfur.Args
}

// DependOn is used for other resources to depend on [Resource].
func (akfur *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(akfur)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (akfur *Resource) Dependencies() terra.Dependencies {
	return akfur.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (akfur *Resource) LifecycleManagement() *terra.Lifecycle {
	return akfur.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (akfur *Resource) Attributes() azurermKubernetesFleetUpdateRunAttributes {
	return azurermKubernetesFleetUpdateRunAttributes{ref: terra.ReferenceResource(akfur)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (akfur *Resource) ImportState(state io.Reader) error {
	akfur.state = &azurermKubernetesFleetUpdateRunState{}
	if err := json.NewDecoder(state).Decode(akfur.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akfur.Type(), akfur.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (akfur *Resource) State() (*azurermKubernetesFleetUpdateRunState, bool) {
	return akfur.state, akfur.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (akfur *Resource) StateMust() *azurermKubernetesFleetUpdateRunState {
	if akfur.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akfur.Type(), akfur.LocalName()))
	}
	return akfur.state
}

// Args contains the configurations for azurerm_kubernetes_fleet_update_run.
type Args struct {
	// FleetUpdateStrategyId: string, optional
	FleetUpdateStrategyId terra.StringValue `hcl:"fleet_update_strategy_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KubernetesFleetManagerId: string, required
	KubernetesFleetManagerId terra.StringValue `hcl:"kubernetes_fleet_manager_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ManagedClusterUpdate: required
	ManagedClusterUpdate *ManagedClusterUpdate `hcl:"managed_cluster_update,block" validate:"required"`
	// Stage: min=0
	Stage []Stage `hcl:"stage,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermKubernetesFleetUpdateRunAttributes struct {
	ref terra.Reference
}

// FleetUpdateStrategyId returns a reference to field fleet_update_strategy_id of azurerm_kubernetes_fleet_update_run.
func (akfur azurermKubernetesFleetUpdateRunAttributes) FleetUpdateStrategyId() terra.StringValue {
	return terra.ReferenceAsString(akfur.ref.Append("fleet_update_strategy_id"))
}

// Id returns a reference to field id of azurerm_kubernetes_fleet_update_run.
func (akfur azurermKubernetesFleetUpdateRunAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akfur.ref.Append("id"))
}

// KubernetesFleetManagerId returns a reference to field kubernetes_fleet_manager_id of azurerm_kubernetes_fleet_update_run.
func (akfur azurermKubernetesFleetUpdateRunAttributes) KubernetesFleetManagerId() terra.StringValue {
	return terra.ReferenceAsString(akfur.ref.Append("kubernetes_fleet_manager_id"))
}

// Name returns a reference to field name of azurerm_kubernetes_fleet_update_run.
func (akfur azurermKubernetesFleetUpdateRunAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(akfur.ref.Append("name"))
}

func (akfur azurermKubernetesFleetUpdateRunAttributes) ManagedClusterUpdate() terra.ListValue[ManagedClusterUpdateAttributes] {
	return terra.ReferenceAsList[ManagedClusterUpdateAttributes](akfur.ref.Append("managed_cluster_update"))
}

func (akfur azurermKubernetesFleetUpdateRunAttributes) Stage() terra.ListValue[StageAttributes] {
	return terra.ReferenceAsList[StageAttributes](akfur.ref.Append("stage"))
}

func (akfur azurermKubernetesFleetUpdateRunAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](akfur.ref.Append("timeouts"))
}

type azurermKubernetesFleetUpdateRunState struct {
	FleetUpdateStrategyId    string                      `json:"fleet_update_strategy_id"`
	Id                       string                      `json:"id"`
	KubernetesFleetManagerId string                      `json:"kubernetes_fleet_manager_id"`
	Name                     string                      `json:"name"`
	ManagedClusterUpdate     []ManagedClusterUpdateState `json:"managed_cluster_update"`
	Stage                    []StageState                `json:"stage"`
	Timeouts                 *TimeoutsState              `json:"timeouts"`
}
