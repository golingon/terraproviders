// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVpcEndpointConnectionAccepter(name string, args VpcEndpointConnectionAccepterArgs) *VpcEndpointConnectionAccepter {
	return &VpcEndpointConnectionAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpointConnectionAccepter)(nil)

type VpcEndpointConnectionAccepter struct {
	Name  string
	Args  VpcEndpointConnectionAccepterArgs
	state *vpcEndpointConnectionAccepterState
}

func (veca *VpcEndpointConnectionAccepter) Type() string {
	return "aws_vpc_endpoint_connection_accepter"
}

func (veca *VpcEndpointConnectionAccepter) LocalName() string {
	return veca.Name
}

func (veca *VpcEndpointConnectionAccepter) Configuration() interface{} {
	return veca.Args
}

func (veca *VpcEndpointConnectionAccepter) Attributes() vpcEndpointConnectionAccepterAttributes {
	return vpcEndpointConnectionAccepterAttributes{ref: terra.ReferenceResource(veca)}
}

func (veca *VpcEndpointConnectionAccepter) ImportState(av io.Reader) error {
	veca.state = &vpcEndpointConnectionAccepterState{}
	if err := json.NewDecoder(av).Decode(veca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", veca.Type(), veca.LocalName(), err)
	}
	return nil
}

func (veca *VpcEndpointConnectionAccepter) State() (*vpcEndpointConnectionAccepterState, bool) {
	return veca.state, veca.state != nil
}

func (veca *VpcEndpointConnectionAccepter) StateMust() *vpcEndpointConnectionAccepterState {
	if veca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", veca.Type(), veca.LocalName()))
	}
	return veca.state
}

func (veca *VpcEndpointConnectionAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(veca)
}

type VpcEndpointConnectionAccepterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VpcEndpointId: string, required
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr" validate:"required"`
	// VpcEndpointServiceId: string, required
	VpcEndpointServiceId terra.StringValue `hcl:"vpc_endpoint_service_id,attr" validate:"required"`
	// DependsOn contains resources that VpcEndpointConnectionAccepter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type vpcEndpointConnectionAccepterAttributes struct {
	ref terra.Reference
}

func (veca vpcEndpointConnectionAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(veca.ref.Append("id"))
}

func (veca vpcEndpointConnectionAccepterAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceString(veca.ref.Append("vpc_endpoint_id"))
}

func (veca vpcEndpointConnectionAccepterAttributes) VpcEndpointServiceId() terra.StringValue {
	return terra.ReferenceString(veca.ref.Append("vpc_endpoint_service_id"))
}

func (veca vpcEndpointConnectionAccepterAttributes) VpcEndpointState() terra.StringValue {
	return terra.ReferenceString(veca.ref.Append("vpc_endpoint_state"))
}

type vpcEndpointConnectionAccepterState struct {
	Id                   string `json:"id"`
	VpcEndpointId        string `json:"vpc_endpoint_id"`
	VpcEndpointServiceId string `json:"vpc_endpoint_service_id"`
	VpcEndpointState     string `json:"vpc_endpoint_state"`
}
