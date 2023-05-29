// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kubernetesfleetmanager "github.com/golingon/terraproviders/azurerm/3.58.0/kubernetesfleetmanager"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKubernetesFleetManager creates a new instance of [KubernetesFleetManager].
func NewKubernetesFleetManager(name string, args KubernetesFleetManagerArgs) *KubernetesFleetManager {
	return &KubernetesFleetManager{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KubernetesFleetManager)(nil)

// KubernetesFleetManager represents the Terraform resource azurerm_kubernetes_fleet_manager.
type KubernetesFleetManager struct {
	Name      string
	Args      KubernetesFleetManagerArgs
	state     *kubernetesFleetManagerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KubernetesFleetManager].
func (kfm *KubernetesFleetManager) Type() string {
	return "azurerm_kubernetes_fleet_manager"
}

// LocalName returns the local name for [KubernetesFleetManager].
func (kfm *KubernetesFleetManager) LocalName() string {
	return kfm.Name
}

// Configuration returns the configuration (args) for [KubernetesFleetManager].
func (kfm *KubernetesFleetManager) Configuration() interface{} {
	return kfm.Args
}

// DependOn is used for other resources to depend on [KubernetesFleetManager].
func (kfm *KubernetesFleetManager) DependOn() terra.Reference {
	return terra.ReferenceResource(kfm)
}

// Dependencies returns the list of resources [KubernetesFleetManager] depends_on.
func (kfm *KubernetesFleetManager) Dependencies() terra.Dependencies {
	return kfm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KubernetesFleetManager].
func (kfm *KubernetesFleetManager) LifecycleManagement() *terra.Lifecycle {
	return kfm.Lifecycle
}

// Attributes returns the attributes for [KubernetesFleetManager].
func (kfm *KubernetesFleetManager) Attributes() kubernetesFleetManagerAttributes {
	return kubernetesFleetManagerAttributes{ref: terra.ReferenceResource(kfm)}
}

// ImportState imports the given attribute values into [KubernetesFleetManager]'s state.
func (kfm *KubernetesFleetManager) ImportState(av io.Reader) error {
	kfm.state = &kubernetesFleetManagerState{}
	if err := json.NewDecoder(av).Decode(kfm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kfm.Type(), kfm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KubernetesFleetManager] has state.
func (kfm *KubernetesFleetManager) State() (*kubernetesFleetManagerState, bool) {
	return kfm.state, kfm.state != nil
}

// StateMust returns the state for [KubernetesFleetManager]. Panics if the state is nil.
func (kfm *KubernetesFleetManager) StateMust() *kubernetesFleetManagerState {
	if kfm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kfm.Type(), kfm.LocalName()))
	}
	return kfm.state
}

// KubernetesFleetManagerArgs contains the configurations for azurerm_kubernetes_fleet_manager.
type KubernetesFleetManagerArgs struct {
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
	// HubProfile: optional
	HubProfile *kubernetesfleetmanager.HubProfile `hcl:"hub_profile,block"`
	// Timeouts: optional
	Timeouts *kubernetesfleetmanager.Timeouts `hcl:"timeouts,block"`
}
type kubernetesFleetManagerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_kubernetes_fleet_manager.
func (kfm kubernetesFleetManagerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kfm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_kubernetes_fleet_manager.
func (kfm kubernetesFleetManagerAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kfm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kubernetes_fleet_manager.
func (kfm kubernetesFleetManagerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kfm.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kubernetes_fleet_manager.
func (kfm kubernetesFleetManagerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kfm.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_kubernetes_fleet_manager.
func (kfm kubernetesFleetManagerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kfm.ref.Append("tags"))
}

func (kfm kubernetesFleetManagerAttributes) HubProfile() terra.ListValue[kubernetesfleetmanager.HubProfileAttributes] {
	return terra.ReferenceAsList[kubernetesfleetmanager.HubProfileAttributes](kfm.ref.Append("hub_profile"))
}

func (kfm kubernetesFleetManagerAttributes) Timeouts() kubernetesfleetmanager.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kubernetesfleetmanager.TimeoutsAttributes](kfm.ref.Append("timeouts"))
}

type kubernetesFleetManagerState struct {
	Id                string                                   `json:"id"`
	Location          string                                   `json:"location"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	Tags              map[string]string                        `json:"tags"`
	HubProfile        []kubernetesfleetmanager.HubProfileState `json:"hub_profile"`
	Timeouts          *kubernetesfleetmanager.TimeoutsState    `json:"timeouts"`
}