// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMainRouteTableAssociation creates a new instance of [MainRouteTableAssociation].
func NewMainRouteTableAssociation(name string, args MainRouteTableAssociationArgs) *MainRouteTableAssociation {
	return &MainRouteTableAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MainRouteTableAssociation)(nil)

// MainRouteTableAssociation represents the Terraform resource aws_main_route_table_association.
type MainRouteTableAssociation struct {
	Name      string
	Args      MainRouteTableAssociationArgs
	state     *mainRouteTableAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MainRouteTableAssociation].
func (mrta *MainRouteTableAssociation) Type() string {
	return "aws_main_route_table_association"
}

// LocalName returns the local name for [MainRouteTableAssociation].
func (mrta *MainRouteTableAssociation) LocalName() string {
	return mrta.Name
}

// Configuration returns the configuration (args) for [MainRouteTableAssociation].
func (mrta *MainRouteTableAssociation) Configuration() interface{} {
	return mrta.Args
}

// DependOn is used for other resources to depend on [MainRouteTableAssociation].
func (mrta *MainRouteTableAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(mrta)
}

// Dependencies returns the list of resources [MainRouteTableAssociation] depends_on.
func (mrta *MainRouteTableAssociation) Dependencies() terra.Dependencies {
	return mrta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MainRouteTableAssociation].
func (mrta *MainRouteTableAssociation) LifecycleManagement() *terra.Lifecycle {
	return mrta.Lifecycle
}

// Attributes returns the attributes for [MainRouteTableAssociation].
func (mrta *MainRouteTableAssociation) Attributes() mainRouteTableAssociationAttributes {
	return mainRouteTableAssociationAttributes{ref: terra.ReferenceResource(mrta)}
}

// ImportState imports the given attribute values into [MainRouteTableAssociation]'s state.
func (mrta *MainRouteTableAssociation) ImportState(av io.Reader) error {
	mrta.state = &mainRouteTableAssociationState{}
	if err := json.NewDecoder(av).Decode(mrta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mrta.Type(), mrta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MainRouteTableAssociation] has state.
func (mrta *MainRouteTableAssociation) State() (*mainRouteTableAssociationState, bool) {
	return mrta.state, mrta.state != nil
}

// StateMust returns the state for [MainRouteTableAssociation]. Panics if the state is nil.
func (mrta *MainRouteTableAssociation) StateMust() *mainRouteTableAssociationState {
	if mrta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mrta.Type(), mrta.LocalName()))
	}
	return mrta.state
}

// MainRouteTableAssociationArgs contains the configurations for aws_main_route_table_association.
type MainRouteTableAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RouteTableId: string, required
	RouteTableId terra.StringValue `hcl:"route_table_id,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}
type mainRouteTableAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_main_route_table_association.
func (mrta mainRouteTableAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mrta.ref.Append("id"))
}

// OriginalRouteTableId returns a reference to field original_route_table_id of aws_main_route_table_association.
func (mrta mainRouteTableAssociationAttributes) OriginalRouteTableId() terra.StringValue {
	return terra.ReferenceAsString(mrta.ref.Append("original_route_table_id"))
}

// RouteTableId returns a reference to field route_table_id of aws_main_route_table_association.
func (mrta mainRouteTableAssociationAttributes) RouteTableId() terra.StringValue {
	return terra.ReferenceAsString(mrta.ref.Append("route_table_id"))
}

// VpcId returns a reference to field vpc_id of aws_main_route_table_association.
func (mrta mainRouteTableAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(mrta.ref.Append("vpc_id"))
}

type mainRouteTableAssociationState struct {
	Id                   string `json:"id"`
	OriginalRouteTableId string `json:"original_route_table_id"`
	RouteTableId         string `json:"route_table_id"`
	VpcId                string `json:"vpc_id"`
}
