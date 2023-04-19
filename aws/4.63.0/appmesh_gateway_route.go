// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshgatewayroute "github.com/golingon/terraproviders/aws/4.63.0/appmeshgatewayroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppmeshGatewayRoute creates a new instance of [AppmeshGatewayRoute].
func NewAppmeshGatewayRoute(name string, args AppmeshGatewayRouteArgs) *AppmeshGatewayRoute {
	return &AppmeshGatewayRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshGatewayRoute)(nil)

// AppmeshGatewayRoute represents the Terraform resource aws_appmesh_gateway_route.
type AppmeshGatewayRoute struct {
	Name      string
	Args      AppmeshGatewayRouteArgs
	state     *appmeshGatewayRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppmeshGatewayRoute].
func (agr *AppmeshGatewayRoute) Type() string {
	return "aws_appmesh_gateway_route"
}

// LocalName returns the local name for [AppmeshGatewayRoute].
func (agr *AppmeshGatewayRoute) LocalName() string {
	return agr.Name
}

// Configuration returns the configuration (args) for [AppmeshGatewayRoute].
func (agr *AppmeshGatewayRoute) Configuration() interface{} {
	return agr.Args
}

// DependOn is used for other resources to depend on [AppmeshGatewayRoute].
func (agr *AppmeshGatewayRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(agr)
}

// Dependencies returns the list of resources [AppmeshGatewayRoute] depends_on.
func (agr *AppmeshGatewayRoute) Dependencies() terra.Dependencies {
	return agr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppmeshGatewayRoute].
func (agr *AppmeshGatewayRoute) LifecycleManagement() *terra.Lifecycle {
	return agr.Lifecycle
}

// Attributes returns the attributes for [AppmeshGatewayRoute].
func (agr *AppmeshGatewayRoute) Attributes() appmeshGatewayRouteAttributes {
	return appmeshGatewayRouteAttributes{ref: terra.ReferenceResource(agr)}
}

// ImportState imports the given attribute values into [AppmeshGatewayRoute]'s state.
func (agr *AppmeshGatewayRoute) ImportState(av io.Reader) error {
	agr.state = &appmeshGatewayRouteState{}
	if err := json.NewDecoder(av).Decode(agr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agr.Type(), agr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppmeshGatewayRoute] has state.
func (agr *AppmeshGatewayRoute) State() (*appmeshGatewayRouteState, bool) {
	return agr.state, agr.state != nil
}

// StateMust returns the state for [AppmeshGatewayRoute]. Panics if the state is nil.
func (agr *AppmeshGatewayRoute) StateMust() *appmeshGatewayRouteState {
	if agr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agr.Type(), agr.LocalName()))
	}
	return agr.state
}

// AppmeshGatewayRouteArgs contains the configurations for aws_appmesh_gateway_route.
type AppmeshGatewayRouteArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MeshName: string, required
	MeshName terra.StringValue `hcl:"mesh_name,attr" validate:"required"`
	// MeshOwner: string, optional
	MeshOwner terra.StringValue `hcl:"mesh_owner,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VirtualGatewayName: string, required
	VirtualGatewayName terra.StringValue `hcl:"virtual_gateway_name,attr" validate:"required"`
	// Spec: required
	Spec *appmeshgatewayroute.Spec `hcl:"spec,block" validate:"required"`
}
type appmeshGatewayRouteAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agr.ref.Append("tags_all"))
}

// VirtualGatewayName returns a reference to field virtual_gateway_name of aws_appmesh_gateway_route.
func (agr appmeshGatewayRouteAttributes) VirtualGatewayName() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("virtual_gateway_name"))
}

func (agr appmeshGatewayRouteAttributes) Spec() terra.ListValue[appmeshgatewayroute.SpecAttributes] {
	return terra.ReferenceAsList[appmeshgatewayroute.SpecAttributes](agr.ref.Append("spec"))
}

type appmeshGatewayRouteState struct {
	Arn                string                          `json:"arn"`
	CreatedDate        string                          `json:"created_date"`
	Id                 string                          `json:"id"`
	LastUpdatedDate    string                          `json:"last_updated_date"`
	MeshName           string                          `json:"mesh_name"`
	MeshOwner          string                          `json:"mesh_owner"`
	Name               string                          `json:"name"`
	ResourceOwner      string                          `json:"resource_owner"`
	Tags               map[string]string               `json:"tags"`
	TagsAll            map[string]string               `json:"tags_all"`
	VirtualGatewayName string                          `json:"virtual_gateway_name"`
	Spec               []appmeshgatewayroute.SpecState `json:"spec"`
}