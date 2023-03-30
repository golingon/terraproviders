// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMainRouteTableAssociation(name string, args MainRouteTableAssociationArgs) *MainRouteTableAssociation {
	return &MainRouteTableAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MainRouteTableAssociation)(nil)

type MainRouteTableAssociation struct {
	Name  string
	Args  MainRouteTableAssociationArgs
	state *mainRouteTableAssociationState
}

func (mrta *MainRouteTableAssociation) Type() string {
	return "aws_main_route_table_association"
}

func (mrta *MainRouteTableAssociation) LocalName() string {
	return mrta.Name
}

func (mrta *MainRouteTableAssociation) Configuration() interface{} {
	return mrta.Args
}

func (mrta *MainRouteTableAssociation) Attributes() mainRouteTableAssociationAttributes {
	return mainRouteTableAssociationAttributes{ref: terra.ReferenceResource(mrta)}
}

func (mrta *MainRouteTableAssociation) ImportState(av io.Reader) error {
	mrta.state = &mainRouteTableAssociationState{}
	if err := json.NewDecoder(av).Decode(mrta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mrta.Type(), mrta.LocalName(), err)
	}
	return nil
}

func (mrta *MainRouteTableAssociation) State() (*mainRouteTableAssociationState, bool) {
	return mrta.state, mrta.state != nil
}

func (mrta *MainRouteTableAssociation) StateMust() *mainRouteTableAssociationState {
	if mrta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mrta.Type(), mrta.LocalName()))
	}
	return mrta.state
}

func (mrta *MainRouteTableAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(mrta)
}

type MainRouteTableAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// DependsOn contains resources that MainRouteTableAssociation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type mainRouteTableAssociationAttributes struct {
	ref terra.Reference
}

func (mrta mainRouteTableAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mrta.ref.Append("id"))
}

func (mrta mainRouteTableAssociationAttributes) OriginalRouteTableId() terra.StringValue {
	return terra.ReferenceString(mrta.ref.Append("original_route_table_id"))
}

func (mrta mainRouteTableAssociationAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceString(mrta.ref.Append("route_table_id"))
}

func (mrta mainRouteTableAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(mrta.ref.Append("vpc_id"))
}

type mainRouteTableAssociationState struct {
	Id                   string `json:"id"`
	OriginalRouteTableId string `json:"original_route_table_id"`
	RouteTableId         string `json:"route_table_id"`
	VpcId                string `json:"vpc_id"`
}
