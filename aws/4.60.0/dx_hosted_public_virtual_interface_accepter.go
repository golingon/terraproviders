// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxhostedpublicvirtualinterfaceaccepter "github.com/golingon/terraproviders/aws/4.60.0/dxhostedpublicvirtualinterfaceaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDxHostedPublicVirtualInterfaceAccepter(name string, args DxHostedPublicVirtualInterfaceAccepterArgs) *DxHostedPublicVirtualInterfaceAccepter {
	return &DxHostedPublicVirtualInterfaceAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedPublicVirtualInterfaceAccepter)(nil)

type DxHostedPublicVirtualInterfaceAccepter struct {
	Name  string
	Args  DxHostedPublicVirtualInterfaceAccepterArgs
	state *dxHostedPublicVirtualInterfaceAccepterState
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Type() string {
	return "aws_dx_hosted_public_virtual_interface_accepter"
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) LocalName() string {
	return dhpvia.Name
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Configuration() interface{} {
	return dhpvia.Args
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Attributes() dxHostedPublicVirtualInterfaceAccepterAttributes {
	return dxHostedPublicVirtualInterfaceAccepterAttributes{ref: terra.ReferenceResource(dhpvia)}
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) ImportState(av io.Reader) error {
	dhpvia.state = &dxHostedPublicVirtualInterfaceAccepterState{}
	if err := json.NewDecoder(av).Decode(dhpvia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhpvia.Type(), dhpvia.LocalName(), err)
	}
	return nil
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) State() (*dxHostedPublicVirtualInterfaceAccepterState, bool) {
	return dhpvia.state, dhpvia.state != nil
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) StateMust() *dxHostedPublicVirtualInterfaceAccepterState {
	if dhpvia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhpvia.Type(), dhpvia.LocalName()))
	}
	return dhpvia.state
}

func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(dhpvia)
}

type DxHostedPublicVirtualInterfaceAccepterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VirtualInterfaceId: string, required
	VirtualInterfaceId terra.StringValue `hcl:"virtual_interface_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dxhostedpublicvirtualinterfaceaccepter.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DxHostedPublicVirtualInterfaceAccepter depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dxHostedPublicVirtualInterfaceAccepterAttributes struct {
	ref terra.Reference
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(dhpvia.ref.Append("arn"))
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dhpvia.ref.Append("id"))
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dhpvia.ref.Append("tags"))
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dhpvia.ref.Append("tags_all"))
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) VirtualInterfaceId() terra.StringValue {
	return terra.ReferenceString(dhpvia.ref.Append("virtual_interface_id"))
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Timeouts() dxhostedpublicvirtualinterfaceaccepter.TimeoutsAttributes {
	return terra.ReferenceSingle[dxhostedpublicvirtualinterfaceaccepter.TimeoutsAttributes](dhpvia.ref.Append("timeouts"))
}

type dxHostedPublicVirtualInterfaceAccepterState struct {
	Arn                string                                                `json:"arn"`
	Id                 string                                                `json:"id"`
	Tags               map[string]string                                     `json:"tags"`
	TagsAll            map[string]string                                     `json:"tags_all"`
	VirtualInterfaceId string                                                `json:"virtual_interface_id"`
	Timeouts           *dxhostedpublicvirtualinterfaceaccepter.TimeoutsState `json:"timeouts"`
}
