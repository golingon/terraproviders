// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNatGateway creates a new instance of [NatGateway].
func NewNatGateway(name string, args NatGatewayArgs) *NatGateway {
	return &NatGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NatGateway)(nil)

// NatGateway represents the Terraform resource aws_nat_gateway.
type NatGateway struct {
	Name      string
	Args      NatGatewayArgs
	state     *natGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NatGateway].
func (ng *NatGateway) Type() string {
	return "aws_nat_gateway"
}

// LocalName returns the local name for [NatGateway].
func (ng *NatGateway) LocalName() string {
	return ng.Name
}

// Configuration returns the configuration (args) for [NatGateway].
func (ng *NatGateway) Configuration() interface{} {
	return ng.Args
}

// DependOn is used for other resources to depend on [NatGateway].
func (ng *NatGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(ng)
}

// Dependencies returns the list of resources [NatGateway] depends_on.
func (ng *NatGateway) Dependencies() terra.Dependencies {
	return ng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NatGateway].
func (ng *NatGateway) LifecycleManagement() *terra.Lifecycle {
	return ng.Lifecycle
}

// Attributes returns the attributes for [NatGateway].
func (ng *NatGateway) Attributes() natGatewayAttributes {
	return natGatewayAttributes{ref: terra.ReferenceResource(ng)}
}

// ImportState imports the given attribute values into [NatGateway]'s state.
func (ng *NatGateway) ImportState(av io.Reader) error {
	ng.state = &natGatewayState{}
	if err := json.NewDecoder(av).Decode(ng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ng.Type(), ng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NatGateway] has state.
func (ng *NatGateway) State() (*natGatewayState, bool) {
	return ng.state, ng.state != nil
}

// StateMust returns the state for [NatGateway]. Panics if the state is nil.
func (ng *NatGateway) StateMust() *natGatewayState {
	if ng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ng.Type(), ng.LocalName()))
	}
	return ng.state
}

// NatGatewayArgs contains the configurations for aws_nat_gateway.
type NatGatewayArgs struct {
	// AllocationId: string, optional
	AllocationId terra.StringValue `hcl:"allocation_id,attr"`
	// ConnectivityType: string, optional
	ConnectivityType terra.StringValue `hcl:"connectivity_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrivateIp: string, optional
	PrivateIp terra.StringValue `hcl:"private_ip,attr"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type natGatewayAttributes struct {
	ref terra.Reference
}

// AllocationId returns a reference to field allocation_id of aws_nat_gateway.
func (ng natGatewayAttributes) AllocationId() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("allocation_id"))
}

// AssociationId returns a reference to field association_id of aws_nat_gateway.
func (ng natGatewayAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("association_id"))
}

// ConnectivityType returns a reference to field connectivity_type of aws_nat_gateway.
func (ng natGatewayAttributes) ConnectivityType() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("connectivity_type"))
}

// Id returns a reference to field id of aws_nat_gateway.
func (ng natGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("id"))
}

// NetworkInterfaceId returns a reference to field network_interface_id of aws_nat_gateway.
func (ng natGatewayAttributes) NetworkInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("network_interface_id"))
}

// PrivateIp returns a reference to field private_ip of aws_nat_gateway.
func (ng natGatewayAttributes) PrivateIp() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("private_ip"))
}

// PublicIp returns a reference to field public_ip of aws_nat_gateway.
func (ng natGatewayAttributes) PublicIp() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("public_ip"))
}

// SubnetId returns a reference to field subnet_id of aws_nat_gateway.
func (ng natGatewayAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_nat_gateway.
func (ng natGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ng.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_nat_gateway.
func (ng natGatewayAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ng.ref.Append("tags_all"))
}

type natGatewayState struct {
	AllocationId       string            `json:"allocation_id"`
	AssociationId      string            `json:"association_id"`
	ConnectivityType   string            `json:"connectivity_type"`
	Id                 string            `json:"id"`
	NetworkInterfaceId string            `json:"network_interface_id"`
	PrivateIp          string            `json:"private_ip"`
	PublicIp           string            `json:"public_ip"`
	SubnetId           string            `json:"subnet_id"`
	Tags               map[string]string `json:"tags"`
	TagsAll            map[string]string `json:"tags_all"`
}
