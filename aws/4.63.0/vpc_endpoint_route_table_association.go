// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcEndpointRouteTableAssociation creates a new instance of [VpcEndpointRouteTableAssociation].
func NewVpcEndpointRouteTableAssociation(name string, args VpcEndpointRouteTableAssociationArgs) *VpcEndpointRouteTableAssociation {
	return &VpcEndpointRouteTableAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpointRouteTableAssociation)(nil)

// VpcEndpointRouteTableAssociation represents the Terraform resource aws_vpc_endpoint_route_table_association.
type VpcEndpointRouteTableAssociation struct {
	Name      string
	Args      VpcEndpointRouteTableAssociationArgs
	state     *vpcEndpointRouteTableAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcEndpointRouteTableAssociation].
func (verta *VpcEndpointRouteTableAssociation) Type() string {
	return "aws_vpc_endpoint_route_table_association"
}

// LocalName returns the local name for [VpcEndpointRouteTableAssociation].
func (verta *VpcEndpointRouteTableAssociation) LocalName() string {
	return verta.Name
}

// Configuration returns the configuration (args) for [VpcEndpointRouteTableAssociation].
func (verta *VpcEndpointRouteTableAssociation) Configuration() interface{} {
	return verta.Args
}

// DependOn is used for other resources to depend on [VpcEndpointRouteTableAssociation].
func (verta *VpcEndpointRouteTableAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(verta)
}

// Dependencies returns the list of resources [VpcEndpointRouteTableAssociation] depends_on.
func (verta *VpcEndpointRouteTableAssociation) Dependencies() terra.Dependencies {
	return verta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcEndpointRouteTableAssociation].
func (verta *VpcEndpointRouteTableAssociation) LifecycleManagement() *terra.Lifecycle {
	return verta.Lifecycle
}

// Attributes returns the attributes for [VpcEndpointRouteTableAssociation].
func (verta *VpcEndpointRouteTableAssociation) Attributes() vpcEndpointRouteTableAssociationAttributes {
	return vpcEndpointRouteTableAssociationAttributes{ref: terra.ReferenceResource(verta)}
}

// ImportState imports the given attribute values into [VpcEndpointRouteTableAssociation]'s state.
func (verta *VpcEndpointRouteTableAssociation) ImportState(av io.Reader) error {
	verta.state = &vpcEndpointRouteTableAssociationState{}
	if err := json.NewDecoder(av).Decode(verta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", verta.Type(), verta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcEndpointRouteTableAssociation] has state.
func (verta *VpcEndpointRouteTableAssociation) State() (*vpcEndpointRouteTableAssociationState, bool) {
	return verta.state, verta.state != nil
}

// StateMust returns the state for [VpcEndpointRouteTableAssociation]. Panics if the state is nil.
func (verta *VpcEndpointRouteTableAssociation) StateMust() *vpcEndpointRouteTableAssociationState {
	if verta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", verta.Type(), verta.LocalName()))
	}
	return verta.state
}

// VpcEndpointRouteTableAssociationArgs contains the configurations for aws_vpc_endpoint_route_table_association.
type VpcEndpointRouteTableAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// VpcEndpointId: string, required
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr" validate:"required"`
}
type vpcEndpointRouteTableAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_endpoint_route_table_association.
func (verta vpcEndpointRouteTableAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(verta.ref.Append("id"))
}

// RouteTableId returns a reference to field route_table_id of aws_vpc_endpoint_route_table_association.
func (verta vpcEndpointRouteTableAssociationAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(verta.ref.Append("route_table_id"))
}

// VpcEndpointId returns a reference to field vpc_endpoint_id of aws_vpc_endpoint_route_table_association.
func (verta vpcEndpointRouteTableAssociationAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(verta.ref.Append("vpc_endpoint_id"))
}

type vpcEndpointRouteTableAssociationState struct {
	Id            string `json:"id"`
	RouteTableId  string `json:"route_table_id"`
	VpcEndpointId string `json:"vpc_endpoint_id"`
}
