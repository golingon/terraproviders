// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestvirtualnetwork "github.com/golingon/terraproviders/azurerm/3.65.0/devtestvirtualnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDevTestVirtualNetwork creates a new instance of [DevTestVirtualNetwork].
func NewDevTestVirtualNetwork(name string, args DevTestVirtualNetworkArgs) *DevTestVirtualNetwork {
	return &DevTestVirtualNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestVirtualNetwork)(nil)

// DevTestVirtualNetwork represents the Terraform resource azurerm_dev_test_virtual_network.
type DevTestVirtualNetwork struct {
	Name      string
	Args      DevTestVirtualNetworkArgs
	state     *devTestVirtualNetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DevTestVirtualNetwork].
func (dtvn *DevTestVirtualNetwork) Type() string {
	return "azurerm_dev_test_virtual_network"
}

// LocalName returns the local name for [DevTestVirtualNetwork].
func (dtvn *DevTestVirtualNetwork) LocalName() string {
	return dtvn.Name
}

// Configuration returns the configuration (args) for [DevTestVirtualNetwork].
func (dtvn *DevTestVirtualNetwork) Configuration() interface{} {
	return dtvn.Args
}

// DependOn is used for other resources to depend on [DevTestVirtualNetwork].
func (dtvn *DevTestVirtualNetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(dtvn)
}

// Dependencies returns the list of resources [DevTestVirtualNetwork] depends_on.
func (dtvn *DevTestVirtualNetwork) Dependencies() terra.Dependencies {
	return dtvn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DevTestVirtualNetwork].
func (dtvn *DevTestVirtualNetwork) LifecycleManagement() *terra.Lifecycle {
	return dtvn.Lifecycle
}

// Attributes returns the attributes for [DevTestVirtualNetwork].
func (dtvn *DevTestVirtualNetwork) Attributes() devTestVirtualNetworkAttributes {
	return devTestVirtualNetworkAttributes{ref: terra.ReferenceResource(dtvn)}
}

// ImportState imports the given attribute values into [DevTestVirtualNetwork]'s state.
func (dtvn *DevTestVirtualNetwork) ImportState(av io.Reader) error {
	dtvn.state = &devTestVirtualNetworkState{}
	if err := json.NewDecoder(av).Decode(dtvn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtvn.Type(), dtvn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DevTestVirtualNetwork] has state.
func (dtvn *DevTestVirtualNetwork) State() (*devTestVirtualNetworkState, bool) {
	return dtvn.state, dtvn.state != nil
}

// StateMust returns the state for [DevTestVirtualNetwork]. Panics if the state is nil.
func (dtvn *DevTestVirtualNetwork) StateMust() *devTestVirtualNetworkState {
	if dtvn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtvn.Type(), dtvn.LocalName()))
	}
	return dtvn.state
}

// DevTestVirtualNetworkArgs contains the configurations for azurerm_dev_test_virtual_network.
type DevTestVirtualNetworkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabName: string, required
	LabName terra.StringValue `hcl:"lab_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Subnet: optional
	Subnet *devtestvirtualnetwork.Subnet `hcl:"subnet,block"`
	// Timeouts: optional
	Timeouts *devtestvirtualnetwork.Timeouts `hcl:"timeouts,block"`
}
type devTestVirtualNetworkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dtvn.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtvn.ref.Append("id"))
}

// LabName returns a reference to field lab_name of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) LabName() terra.StringValue {
	return terra.ReferenceAsString(dtvn.ref.Append("lab_name"))
}

// Name returns a reference to field name of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtvn.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtvn.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtvn.ref.Append("tags"))
}

// UniqueIdentifier returns a reference to field unique_identifier of azurerm_dev_test_virtual_network.
func (dtvn devTestVirtualNetworkAttributes) UniqueIdentifier() terra.StringValue {
	return terra.ReferenceAsString(dtvn.ref.Append("unique_identifier"))
}

func (dtvn devTestVirtualNetworkAttributes) Subnet() terra.ListValue[devtestvirtualnetwork.SubnetAttributes] {
	return terra.ReferenceAsList[devtestvirtualnetwork.SubnetAttributes](dtvn.ref.Append("subnet"))
}

func (dtvn devTestVirtualNetworkAttributes) Timeouts() devtestvirtualnetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[devtestvirtualnetwork.TimeoutsAttributes](dtvn.ref.Append("timeouts"))
}

type devTestVirtualNetworkState struct {
	Description       string                               `json:"description"`
	Id                string                               `json:"id"`
	LabName           string                               `json:"lab_name"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	UniqueIdentifier  string                               `json:"unique_identifier"`
	Subnet            []devtestvirtualnetwork.SubnetState  `json:"subnet"`
	Timeouts          *devtestvirtualnetwork.TimeoutsState `json:"timeouts"`
}
