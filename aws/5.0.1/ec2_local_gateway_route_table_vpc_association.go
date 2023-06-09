// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2LocalGatewayRouteTableVpcAssociation creates a new instance of [Ec2LocalGatewayRouteTableVpcAssociation].
func NewEc2LocalGatewayRouteTableVpcAssociation(name string, args Ec2LocalGatewayRouteTableVpcAssociationArgs) *Ec2LocalGatewayRouteTableVpcAssociation {
	return &Ec2LocalGatewayRouteTableVpcAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2LocalGatewayRouteTableVpcAssociation)(nil)

// Ec2LocalGatewayRouteTableVpcAssociation represents the Terraform resource aws_ec2_local_gateway_route_table_vpc_association.
type Ec2LocalGatewayRouteTableVpcAssociation struct {
	Name      string
	Args      Ec2LocalGatewayRouteTableVpcAssociationArgs
	state     *ec2LocalGatewayRouteTableVpcAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2LocalGatewayRouteTableVpcAssociation].
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) Type() string {
	return "aws_ec2_local_gateway_route_table_vpc_association"
}

// LocalName returns the local name for [Ec2LocalGatewayRouteTableVpcAssociation].
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) LocalName() string {
	return elgrtva.Name
}

// Configuration returns the configuration (args) for [Ec2LocalGatewayRouteTableVpcAssociation].
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) Configuration() interface{} {
	return elgrtva.Args
}

// DependOn is used for other resources to depend on [Ec2LocalGatewayRouteTableVpcAssociation].
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(elgrtva)
}

// Dependencies returns the list of resources [Ec2LocalGatewayRouteTableVpcAssociation] depends_on.
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) Dependencies() terra.Dependencies {
	return elgrtva.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2LocalGatewayRouteTableVpcAssociation].
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) LifecycleManagement() *terra.Lifecycle {
	return elgrtva.Lifecycle
}

// Attributes returns the attributes for [Ec2LocalGatewayRouteTableVpcAssociation].
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) Attributes() ec2LocalGatewayRouteTableVpcAssociationAttributes {
	return ec2LocalGatewayRouteTableVpcAssociationAttributes{ref: terra.ReferenceResource(elgrtva)}
}

// ImportState imports the given attribute values into [Ec2LocalGatewayRouteTableVpcAssociation]'s state.
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) ImportState(av io.Reader) error {
	elgrtva.state = &ec2LocalGatewayRouteTableVpcAssociationState{}
	if err := json.NewDecoder(av).Decode(elgrtva.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", elgrtva.Type(), elgrtva.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2LocalGatewayRouteTableVpcAssociation] has state.
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) State() (*ec2LocalGatewayRouteTableVpcAssociationState, bool) {
	return elgrtva.state, elgrtva.state != nil
}

// StateMust returns the state for [Ec2LocalGatewayRouteTableVpcAssociation]. Panics if the state is nil.
func (elgrtva *Ec2LocalGatewayRouteTableVpcAssociation) StateMust() *ec2LocalGatewayRouteTableVpcAssociationState {
	if elgrtva.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", elgrtva.Type(), elgrtva.LocalName()))
	}
	return elgrtva.state
}

// Ec2LocalGatewayRouteTableVpcAssociationArgs contains the configurations for aws_ec2_local_gateway_route_table_vpc_association.
type Ec2LocalGatewayRouteTableVpcAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocalGatewayRouteTableId: string, required
	LocalGatewayRouteTableId terra.StringValue `hcl:"local_gateway_route_table_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}
type ec2LocalGatewayRouteTableVpcAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway_route_table_vpc_association.
func (elgrtva ec2LocalGatewayRouteTableVpcAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elgrtva.ref.Append("id"))
}

// LocalGatewayId returns a reference to field local_gateway_id of aws_ec2_local_gateway_route_table_vpc_association.
func (elgrtva ec2LocalGatewayRouteTableVpcAssociationAttributes) LocalGatewayId() terra.StringValue {
	return terra.ReferenceAsString(elgrtva.ref.Append("local_gateway_id"))
}

// LocalGatewayRouteTableId returns a reference to field local_gateway_route_table_id of aws_ec2_local_gateway_route_table_vpc_association.
func (elgrtva ec2LocalGatewayRouteTableVpcAssociationAttributes) LocalGatewayRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(elgrtva.ref.Append("local_gateway_route_table_id"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway_route_table_vpc_association.
func (elgrtva ec2LocalGatewayRouteTableVpcAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgrtva.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_local_gateway_route_table_vpc_association.
func (elgrtva ec2LocalGatewayRouteTableVpcAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elgrtva.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_ec2_local_gateway_route_table_vpc_association.
func (elgrtva ec2LocalGatewayRouteTableVpcAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(elgrtva.ref.Append("vpc_id"))
}

type ec2LocalGatewayRouteTableVpcAssociationState struct {
	Id                       string            `json:"id"`
	LocalGatewayId           string            `json:"local_gateway_id"`
	LocalGatewayRouteTableId string            `json:"local_gateway_route_table_id"`
	Tags                     map[string]string `json:"tags"`
	TagsAll                  map[string]string `json:"tags_all"`
	VpcId                    string            `json:"vpc_id"`
}
