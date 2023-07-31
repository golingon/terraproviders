// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcendpointsubnetassociation "github.com/golingon/terraproviders/aws/5.10.0/vpcendpointsubnetassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcEndpointSubnetAssociation creates a new instance of [VpcEndpointSubnetAssociation].
func NewVpcEndpointSubnetAssociation(name string, args VpcEndpointSubnetAssociationArgs) *VpcEndpointSubnetAssociation {
	return &VpcEndpointSubnetAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpointSubnetAssociation)(nil)

// VpcEndpointSubnetAssociation represents the Terraform resource aws_vpc_endpoint_subnet_association.
type VpcEndpointSubnetAssociation struct {
	Name      string
	Args      VpcEndpointSubnetAssociationArgs
	state     *vpcEndpointSubnetAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcEndpointSubnetAssociation].
func (vesa *VpcEndpointSubnetAssociation) Type() string {
	return "aws_vpc_endpoint_subnet_association"
}

// LocalName returns the local name for [VpcEndpointSubnetAssociation].
func (vesa *VpcEndpointSubnetAssociation) LocalName() string {
	return vesa.Name
}

// Configuration returns the configuration (args) for [VpcEndpointSubnetAssociation].
func (vesa *VpcEndpointSubnetAssociation) Configuration() interface{} {
	return vesa.Args
}

// DependOn is used for other resources to depend on [VpcEndpointSubnetAssociation].
func (vesa *VpcEndpointSubnetAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(vesa)
}

// Dependencies returns the list of resources [VpcEndpointSubnetAssociation] depends_on.
func (vesa *VpcEndpointSubnetAssociation) Dependencies() terra.Dependencies {
	return vesa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcEndpointSubnetAssociation].
func (vesa *VpcEndpointSubnetAssociation) LifecycleManagement() *terra.Lifecycle {
	return vesa.Lifecycle
}

// Attributes returns the attributes for [VpcEndpointSubnetAssociation].
func (vesa *VpcEndpointSubnetAssociation) Attributes() vpcEndpointSubnetAssociationAttributes {
	return vpcEndpointSubnetAssociationAttributes{ref: terra.ReferenceResource(vesa)}
}

// ImportState imports the given attribute values into [VpcEndpointSubnetAssociation]'s state.
func (vesa *VpcEndpointSubnetAssociation) ImportState(av io.Reader) error {
	vesa.state = &vpcEndpointSubnetAssociationState{}
	if err := json.NewDecoder(av).Decode(vesa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vesa.Type(), vesa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcEndpointSubnetAssociation] has state.
func (vesa *VpcEndpointSubnetAssociation) State() (*vpcEndpointSubnetAssociationState, bool) {
	return vesa.state, vesa.state != nil
}

// StateMust returns the state for [VpcEndpointSubnetAssociation]. Panics if the state is nil.
func (vesa *VpcEndpointSubnetAssociation) StateMust() *vpcEndpointSubnetAssociationState {
	if vesa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vesa.Type(), vesa.LocalName()))
	}
	return vesa.state
}

// VpcEndpointSubnetAssociationArgs contains the configurations for aws_vpc_endpoint_subnet_association.
type VpcEndpointSubnetAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// VpcEndpointId: string, required
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vpcendpointsubnetassociation.Timeouts `hcl:"timeouts,block"`
}
type vpcEndpointSubnetAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_endpoint_subnet_association.
func (vesa vpcEndpointSubnetAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vesa.ref.Append("id"))
}

// SubnetId returns a reference to field subnet_id of aws_vpc_endpoint_subnet_association.
func (vesa vpcEndpointSubnetAssociationAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(vesa.ref.Append("subnet_id"))
}

// VpcEndpointId returns a reference to field vpc_endpoint_id of aws_vpc_endpoint_subnet_association.
func (vesa vpcEndpointSubnetAssociationAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(vesa.ref.Append("vpc_endpoint_id"))
}

func (vesa vpcEndpointSubnetAssociationAttributes) Timeouts() vpcendpointsubnetassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcendpointsubnetassociation.TimeoutsAttributes](vesa.ref.Append("timeouts"))
}

type vpcEndpointSubnetAssociationState struct {
	Id            string                                      `json:"id"`
	SubnetId      string                                      `json:"subnet_id"`
	VpcEndpointId string                                      `json:"vpc_endpoint_id"`
	Timeouts      *vpcendpointsubnetassociation.TimeoutsState `json:"timeouts"`
}
