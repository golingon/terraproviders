// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagerdeployment "github.com/golingon/terraproviders/azurerm/3.63.0/networkmanagerdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerDeployment creates a new instance of [NetworkManagerDeployment].
func NewNetworkManagerDeployment(name string, args NetworkManagerDeploymentArgs) *NetworkManagerDeployment {
	return &NetworkManagerDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerDeployment)(nil)

// NetworkManagerDeployment represents the Terraform resource azurerm_network_manager_deployment.
type NetworkManagerDeployment struct {
	Name      string
	Args      NetworkManagerDeploymentArgs
	state     *networkManagerDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerDeployment].
func (nmd *NetworkManagerDeployment) Type() string {
	return "azurerm_network_manager_deployment"
}

// LocalName returns the local name for [NetworkManagerDeployment].
func (nmd *NetworkManagerDeployment) LocalName() string {
	return nmd.Name
}

// Configuration returns the configuration (args) for [NetworkManagerDeployment].
func (nmd *NetworkManagerDeployment) Configuration() interface{} {
	return nmd.Args
}

// DependOn is used for other resources to depend on [NetworkManagerDeployment].
func (nmd *NetworkManagerDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(nmd)
}

// Dependencies returns the list of resources [NetworkManagerDeployment] depends_on.
func (nmd *NetworkManagerDeployment) Dependencies() terra.Dependencies {
	return nmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerDeployment].
func (nmd *NetworkManagerDeployment) LifecycleManagement() *terra.Lifecycle {
	return nmd.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerDeployment].
func (nmd *NetworkManagerDeployment) Attributes() networkManagerDeploymentAttributes {
	return networkManagerDeploymentAttributes{ref: terra.ReferenceResource(nmd)}
}

// ImportState imports the given attribute values into [NetworkManagerDeployment]'s state.
func (nmd *NetworkManagerDeployment) ImportState(av io.Reader) error {
	nmd.state = &networkManagerDeploymentState{}
	if err := json.NewDecoder(av).Decode(nmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmd.Type(), nmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerDeployment] has state.
func (nmd *NetworkManagerDeployment) State() (*networkManagerDeploymentState, bool) {
	return nmd.state, nmd.state != nil
}

// StateMust returns the state for [NetworkManagerDeployment]. Panics if the state is nil.
func (nmd *NetworkManagerDeployment) StateMust() *networkManagerDeploymentState {
	if nmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmd.Type(), nmd.LocalName()))
	}
	return nmd.state
}

// NetworkManagerDeploymentArgs contains the configurations for azurerm_network_manager_deployment.
type NetworkManagerDeploymentArgs struct {
	// ConfigurationIds: list of string, required
	ConfigurationIds terra.ListValue[terra.StringValue] `hcl:"configuration_ids,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// ScopeAccess: string, required
	ScopeAccess terra.StringValue `hcl:"scope_access,attr" validate:"required"`
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
	// Timeouts: optional
	Timeouts *networkmanagerdeployment.Timeouts `hcl:"timeouts,block"`
}
type networkManagerDeploymentAttributes struct {
	ref terra.Reference
}

// ConfigurationIds returns a reference to field configuration_ids of azurerm_network_manager_deployment.
func (nmd networkManagerDeploymentAttributes) ConfigurationIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nmd.ref.Append("configuration_ids"))
}

// Id returns a reference to field id of azurerm_network_manager_deployment.
func (nmd networkManagerDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_manager_deployment.
func (nmd networkManagerDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nmd.ref.Append("location"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_deployment.
func (nmd networkManagerDeploymentAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(nmd.ref.Append("network_manager_id"))
}

// ScopeAccess returns a reference to field scope_access of azurerm_network_manager_deployment.
func (nmd networkManagerDeploymentAttributes) ScopeAccess() terra.StringValue {
	return terra.ReferenceAsString(nmd.ref.Append("scope_access"))
}

// Triggers returns a reference to field triggers of azurerm_network_manager_deployment.
func (nmd networkManagerDeploymentAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nmd.ref.Append("triggers"))
}

func (nmd networkManagerDeploymentAttributes) Timeouts() networkmanagerdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerdeployment.TimeoutsAttributes](nmd.ref.Append("timeouts"))
}

type networkManagerDeploymentState struct {
	ConfigurationIds []string                                `json:"configuration_ids"`
	Id               string                                  `json:"id"`
	Location         string                                  `json:"location"`
	NetworkManagerId string                                  `json:"network_manager_id"`
	ScopeAccess      string                                  `json:"scope_access"`
	Triggers         map[string]string                       `json:"triggers"`
	Timeouts         *networkmanagerdeployment.TimeoutsState `json:"timeouts"`
}
