// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_subnet_route_table_association

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

// Resource represents the Terraform resource azurerm_subnet_route_table_association.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSubnetRouteTableAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asrta *Resource) Type() string {
	return "azurerm_subnet_route_table_association"
}

// LocalName returns the local name for [Resource].
func (asrta *Resource) LocalName() string {
	return asrta.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asrta *Resource) Configuration() interface{} {
	return asrta.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asrta *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asrta)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asrta *Resource) Dependencies() terra.Dependencies {
	return asrta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asrta *Resource) LifecycleManagement() *terra.Lifecycle {
	return asrta.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asrta *Resource) Attributes() azurermSubnetRouteTableAssociationAttributes {
	return azurermSubnetRouteTableAssociationAttributes{ref: terra.ReferenceResource(asrta)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asrta *Resource) ImportState(state io.Reader) error {
	asrta.state = &azurermSubnetRouteTableAssociationState{}
	if err := json.NewDecoder(state).Decode(asrta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asrta.Type(), asrta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asrta *Resource) State() (*azurermSubnetRouteTableAssociationState, bool) {
	return asrta.state, asrta.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asrta *Resource) StateMust() *azurermSubnetRouteTableAssociationState {
	if asrta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asrta.Type(), asrta.LocalName()))
	}
	return asrta.state
}

// Args contains the configurations for azurerm_subnet_route_table_association.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSubnetRouteTableAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_subnet_route_table_association.
func (asrta azurermSubnetRouteTableAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asrta.ref.Append("id"))
}

// RouteTableId returns a reference to field route_table_id of azurerm_subnet_route_table_association.
func (asrta azurermSubnetRouteTableAssociationAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(asrta.ref.Append("route_table_id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_subnet_route_table_association.
func (asrta azurermSubnetRouteTableAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(asrta.ref.Append("subnet_id"))
}

func (asrta azurermSubnetRouteTableAssociationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asrta.ref.Append("timeouts"))
}

type azurermSubnetRouteTableAssociationState struct {
	Id           string         `json:"id"`
	RouteTableId string         `json:"route_table_id"`
	SubnetId     string         `json:"subnet_id"`
	Timeouts     *TimeoutsState `json:"timeouts"`
}
