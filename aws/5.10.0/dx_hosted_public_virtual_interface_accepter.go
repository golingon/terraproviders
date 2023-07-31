// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dxhostedpublicvirtualinterfaceaccepter "github.com/golingon/terraproviders/aws/5.10.0/dxhostedpublicvirtualinterfaceaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDxHostedPublicVirtualInterfaceAccepter creates a new instance of [DxHostedPublicVirtualInterfaceAccepter].
func NewDxHostedPublicVirtualInterfaceAccepter(name string, args DxHostedPublicVirtualInterfaceAccepterArgs) *DxHostedPublicVirtualInterfaceAccepter {
	return &DxHostedPublicVirtualInterfaceAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DxHostedPublicVirtualInterfaceAccepter)(nil)

// DxHostedPublicVirtualInterfaceAccepter represents the Terraform resource aws_dx_hosted_public_virtual_interface_accepter.
type DxHostedPublicVirtualInterfaceAccepter struct {
	Name      string
	Args      DxHostedPublicVirtualInterfaceAccepterArgs
	state     *dxHostedPublicVirtualInterfaceAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DxHostedPublicVirtualInterfaceAccepter].
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Type() string {
	return "aws_dx_hosted_public_virtual_interface_accepter"
}

// LocalName returns the local name for [DxHostedPublicVirtualInterfaceAccepter].
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) LocalName() string {
	return dhpvia.Name
}

// Configuration returns the configuration (args) for [DxHostedPublicVirtualInterfaceAccepter].
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Configuration() interface{} {
	return dhpvia.Args
}

// DependOn is used for other resources to depend on [DxHostedPublicVirtualInterfaceAccepter].
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(dhpvia)
}

// Dependencies returns the list of resources [DxHostedPublicVirtualInterfaceAccepter] depends_on.
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Dependencies() terra.Dependencies {
	return dhpvia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DxHostedPublicVirtualInterfaceAccepter].
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) LifecycleManagement() *terra.Lifecycle {
	return dhpvia.Lifecycle
}

// Attributes returns the attributes for [DxHostedPublicVirtualInterfaceAccepter].
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) Attributes() dxHostedPublicVirtualInterfaceAccepterAttributes {
	return dxHostedPublicVirtualInterfaceAccepterAttributes{ref: terra.ReferenceResource(dhpvia)}
}

// ImportState imports the given attribute values into [DxHostedPublicVirtualInterfaceAccepter]'s state.
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) ImportState(av io.Reader) error {
	dhpvia.state = &dxHostedPublicVirtualInterfaceAccepterState{}
	if err := json.NewDecoder(av).Decode(dhpvia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dhpvia.Type(), dhpvia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DxHostedPublicVirtualInterfaceAccepter] has state.
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) State() (*dxHostedPublicVirtualInterfaceAccepterState, bool) {
	return dhpvia.state, dhpvia.state != nil
}

// StateMust returns the state for [DxHostedPublicVirtualInterfaceAccepter]. Panics if the state is nil.
func (dhpvia *DxHostedPublicVirtualInterfaceAccepter) StateMust() *dxHostedPublicVirtualInterfaceAccepterState {
	if dhpvia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dhpvia.Type(), dhpvia.LocalName()))
	}
	return dhpvia.state
}

// DxHostedPublicVirtualInterfaceAccepterArgs contains the configurations for aws_dx_hosted_public_virtual_interface_accepter.
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
}
type dxHostedPublicVirtualInterfaceAccepterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_dx_hosted_public_virtual_interface_accepter.
func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dhpvia.ref.Append("arn"))
}

// Id returns a reference to field id of aws_dx_hosted_public_virtual_interface_accepter.
func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dhpvia.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_dx_hosted_public_virtual_interface_accepter.
func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dhpvia.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dx_hosted_public_virtual_interface_accepter.
func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dhpvia.ref.Append("tags_all"))
}

// VirtualInterfaceId returns a reference to field virtual_interface_id of aws_dx_hosted_public_virtual_interface_accepter.
func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) VirtualInterfaceId() terra.StringValue {
	return terra.ReferenceAsString(dhpvia.ref.Append("virtual_interface_id"))
}

func (dhpvia dxHostedPublicVirtualInterfaceAccepterAttributes) Timeouts() dxhostedpublicvirtualinterfaceaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dxhostedpublicvirtualinterfaceaccepter.TimeoutsAttributes](dhpvia.ref.Append("timeouts"))
}

type dxHostedPublicVirtualInterfaceAccepterState struct {
	Arn                string                                                `json:"arn"`
	Id                 string                                                `json:"id"`
	Tags               map[string]string                                     `json:"tags"`
	TagsAll            map[string]string                                     `json:"tags_all"`
	VirtualInterfaceId string                                                `json:"virtual_interface_id"`
	Timeouts           *dxhostedpublicvirtualinterfaceaccepter.TimeoutsState `json:"timeouts"`
}
