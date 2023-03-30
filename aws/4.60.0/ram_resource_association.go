// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewRamResourceAssociation(name string, args RamResourceAssociationArgs) *RamResourceAssociation {
	return &RamResourceAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RamResourceAssociation)(nil)

type RamResourceAssociation struct {
	Name  string
	Args  RamResourceAssociationArgs
	state *ramResourceAssociationState
}

func (rra *RamResourceAssociation) Type() string {
	return "aws_ram_resource_association"
}

func (rra *RamResourceAssociation) LocalName() string {
	return rra.Name
}

func (rra *RamResourceAssociation) Configuration() interface{} {
	return rra.Args
}

func (rra *RamResourceAssociation) Attributes() ramResourceAssociationAttributes {
	return ramResourceAssociationAttributes{ref: terra.ReferenceResource(rra)}
}

func (rra *RamResourceAssociation) ImportState(av io.Reader) error {
	rra.state = &ramResourceAssociationState{}
	if err := json.NewDecoder(av).Decode(rra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rra.Type(), rra.LocalName(), err)
	}
	return nil
}

func (rra *RamResourceAssociation) State() (*ramResourceAssociationState, bool) {
	return rra.state, rra.state != nil
}

func (rra *RamResourceAssociation) StateMust() *ramResourceAssociationState {
	if rra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rra.Type(), rra.LocalName()))
	}
	return rra.state
}

func (rra *RamResourceAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(rra)
}

type RamResourceAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
	// ResourceShareArn: string, required
	ResourceShareArn terra.StringValue `hcl:"resource_share_arn,attr" validate:"required"`
	// DependsOn contains resources that RamResourceAssociation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ramResourceAssociationAttributes struct {
	ref terra.Reference
}

func (rra ramResourceAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rra.ref.Append("id"))
}

func (rra ramResourceAssociationAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceString(rra.ref.Append("resource_arn"))
}

func (rra ramResourceAssociationAttributes) ResourceShareArn() terra.StringValue {
	return terra.ReferenceString(rra.ref.Append("resource_share_arn"))
}

type ramResourceAssociationState struct {
	Id               string `json:"id"`
	ResourceArn      string `json:"resource_arn"`
	ResourceShareArn string `json:"resource_share_arn"`
}
