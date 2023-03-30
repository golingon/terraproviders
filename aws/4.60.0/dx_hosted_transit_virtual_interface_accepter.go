// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxhostedtransitvirtualinterfaceaccepter "github.com/golingon/terraproviders/aws/4.60.0/dxhostedtransitvirtualinterfaceaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDxHostedTransitVirtualInterfaceAccepter(name string, args DxHostedTransitVirtualInterfaceAccepterArgs) *DxHostedTransitVirtualInterfaceAccepter {
	return &DxHostedTransitVirtualInterfaceAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedTransitVirtualInterfaceAccepter)(nil)

type DxHostedTransitVirtualInterfaceAccepter struct {
	Name  string
	Args  DxHostedTransitVirtualInterfaceAccepterArgs
	state *dxHostedTransitVirtualInterfaceAccepterState
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) Type() string {
	return "aws_dx_hosted_transit_virtual_interface_accepter"
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) LocalName() string {
	return dhtvia.Name
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) Configuration() interface{} {
	return dhtvia.Args
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) Attributes() dxHostedTransitVirtualInterfaceAccepterAttributes {
	return dxHostedTransitVirtualInterfaceAccepterAttributes{ref: terra.ReferenceResource(dhtvia)}
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) ImportState(av io.Reader) error {
	dhtvia.state = &dxHostedTransitVirtualInterfaceAccepterState{}
	if err := json.NewDecoder(av).Decode(dhtvia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhtvia.Type(), dhtvia.LocalName(), err)
	}
	return nil
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) State() (*dxHostedTransitVirtualInterfaceAccepterState, bool) {
	return dhtvia.state, dhtvia.state != nil
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) StateMust() *dxHostedTransitVirtualInterfaceAccepterState {
	if dhtvia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhtvia.Type(), dhtvia.LocalName()))
	}
	return dhtvia.state
}

func (dhtvia *DxHostedTransitVirtualInterfaceAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(dhtvia)
}

type DxHostedTransitVirtualInterfaceAccepterArgs struct {
	// DxGatewayId: string, required
	DxGatewayId terra.StringValue `hcl:"dx_gateway_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VirtualInterfaceId: string, required
	VirtualInterfaceId terra.StringValue `hcl:"virtual_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxhostedtransitvirtualinterfaceaccepter.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DxHostedTransitVirtualInterfaceAccepter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dxHostedTransitVirtualInterfaceAccepterAttributes struct {
	ref terra.Reference
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(dhtvia.ref.Append("arn"))
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) DxGatewayId() terra.StringValue {
	return terra.ReferenceString(dhtvia.ref.Append("dx_gateway_id"))
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dhtvia.ref.Append("id"))
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dhtvia.ref.Append("tags"))
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dhtvia.ref.Append("tags_all"))
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) VirtualInterfaceId() terra.StringValue {
	return terra.ReferenceString(dhtvia.ref.Append("virtual_interface_id"))
}

func (dhtvia dxHostedTransitVirtualInterfaceAccepterAttributes) Timeouts() dxhostedtransitvirtualinterfaceaccepter.TimeoutsAttributes {
	return terra.ReferenceSingle[dxhostedtransitvirtualinterfaceaccepter.TimeoutsAttributes](dhtvia.ref.Append("timeouts"))
}

type dxHostedTransitVirtualInterfaceAccepterState struct {
	Arn                string                                                 `json:"arn"`
	DxGatewayId        string                                                 `json:"dx_gateway_id"`
	Id                 string                                                 `json:"id"`
	Tags               map[string]string                                      `json:"tags"`
	TagsAll            map[string]string                                      `json:"tags_all"`
	VirtualInterfaceId string                                                 `json:"virtual_interface_id"`
	Timeouts           *dxhostedtransitvirtualinterfaceaccepter.TimeoutsState `json:"timeouts"`
}
