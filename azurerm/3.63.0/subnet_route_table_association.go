// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subnetroutetableassociation "github.com/golingon/terraproviders/azurerm/3.63.0/subnetroutetableassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubnetRouteTableAssociation creates a new instance of [SubnetRouteTableAssociation].
func NewSubnetRouteTableAssociation(name string, args SubnetRouteTableAssociationArgs) *SubnetRouteTableAssociation {
	return &SubnetRouteTableAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubnetRouteTableAssociation)(nil)

// SubnetRouteTableAssociation represents the Terraform resource azurerm_subnet_route_table_association.
type SubnetRouteTableAssociation struct {
	Name      string
	Args      SubnetRouteTableAssociationArgs
	state     *subnetRouteTableAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubnetRouteTableAssociation].
func (srta *SubnetRouteTableAssociation) Type() string {
	return "azurerm_subnet_route_table_association"
}

// LocalName returns the local name for [SubnetRouteTableAssociation].
func (srta *SubnetRouteTableAssociation) LocalName() string {
	return srta.Name
}

// Configuration returns the configuration (args) for [SubnetRouteTableAssociation].
func (srta *SubnetRouteTableAssociation) Configuration() interface{} {
	return srta.Args
}

// DependOn is used for other resources to depend on [SubnetRouteTableAssociation].
func (srta *SubnetRouteTableAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(srta)
}

// Dependencies returns the list of resources [SubnetRouteTableAssociation] depends_on.
func (srta *SubnetRouteTableAssociation) Dependencies() terra.Dependencies {
	return srta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubnetRouteTableAssociation].
func (srta *SubnetRouteTableAssociation) LifecycleManagement() *terra.Lifecycle {
	return srta.Lifecycle
}

// Attributes returns the attributes for [SubnetRouteTableAssociation].
func (srta *SubnetRouteTableAssociation) Attributes() subnetRouteTableAssociationAttributes {
	return subnetRouteTableAssociationAttributes{ref: terra.ReferenceResource(srta)}
}

// ImportState imports the given attribute values into [SubnetRouteTableAssociation]'s state.
func (srta *SubnetRouteTableAssociation) ImportState(av io.Reader) error {
	srta.state = &subnetRouteTableAssociationState{}
	if err := json.NewDecoder(av).Decode(srta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srta.Type(), srta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubnetRouteTableAssociation] has state.
func (srta *SubnetRouteTableAssociation) State() (*subnetRouteTableAssociationState, bool) {
	return srta.state, srta.state != nil
}

// StateMust returns the state for [SubnetRouteTableAssociation]. Panics if the state is nil.
func (srta *SubnetRouteTableAssociation) StateMust() *subnetRouteTableAssociationState {
	if srta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srta.Type(), srta.LocalName()))
	}
	return srta.state
}

// SubnetRouteTableAssociationArgs contains the configurations for azurerm_subnet_route_table_association.
type SubnetRouteTableAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *subnetroutetableassociation.Timeouts `hcl:"timeouts,block"`
}
type subnetRouteTableAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_subnet_route_table_association.
func (srta subnetRouteTableAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srta.ref.Append("id"))
}

// RouteTableId returns a reference to field route_table_id of azurerm_subnet_route_table_association.
func (srta subnetRouteTableAssociationAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(srta.ref.Append("route_table_id"))
}

// SubnetId returns a reference to field subnet_id of azurerm_subnet_route_table_association.
func (srta subnetRouteTableAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(srta.ref.Append("subnet_id"))
}

func (srta subnetRouteTableAssociationAttributes) Timeouts() subnetroutetableassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subnetroutetableassociation.TimeoutsAttributes](srta.ref.Append("timeouts"))
}

type subnetRouteTableAssociationState struct {
	Id           string                                     `json:"id"`
	RouteTableId string                                     `json:"route_table_id"`
	SubnetId     string                                     `json:"subnet_id"`
	Timeouts     *subnetroutetableassociation.TimeoutsState `json:"timeouts"`
}
