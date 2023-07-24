// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshroute "github.com/golingon/terraproviders/aws/5.9.0/appmeshroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppmeshRoute creates a new instance of [AppmeshRoute].
func NewAppmeshRoute(name string, args AppmeshRouteArgs) *AppmeshRoute {
	return &AppmeshRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshRoute)(nil)

// AppmeshRoute represents the Terraform resource aws_appmesh_route.
type AppmeshRoute struct {
	Name      string
	Args      AppmeshRouteArgs
	state     *appmeshRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppmeshRoute].
func (ar *AppmeshRoute) Type() string {
	return "aws_appmesh_route"
}

// LocalName returns the local name for [AppmeshRoute].
func (ar *AppmeshRoute) LocalName() string {
	return ar.Name
}

// Configuration returns the configuration (args) for [AppmeshRoute].
func (ar *AppmeshRoute) Configuration() interface{} {
	return ar.Args
}

// DependOn is used for other resources to depend on [AppmeshRoute].
func (ar *AppmeshRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(ar)
}

// Dependencies returns the list of resources [AppmeshRoute] depends_on.
func (ar *AppmeshRoute) Dependencies() terra.Dependencies {
	return ar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppmeshRoute].
func (ar *AppmeshRoute) LifecycleManagement() *terra.Lifecycle {
	return ar.Lifecycle
}

// Attributes returns the attributes for [AppmeshRoute].
func (ar *AppmeshRoute) Attributes() appmeshRouteAttributes {
	return appmeshRouteAttributes{ref: terra.ReferenceResource(ar)}
}

// ImportState imports the given attribute values into [AppmeshRoute]'s state.
func (ar *AppmeshRoute) ImportState(av io.Reader) error {
	ar.state = &appmeshRouteState{}
	if err := json.NewDecoder(av).Decode(ar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ar.Type(), ar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppmeshRoute] has state.
func (ar *AppmeshRoute) State() (*appmeshRouteState, bool) {
	return ar.state, ar.state != nil
}

// StateMust returns the state for [AppmeshRoute]. Panics if the state is nil.
func (ar *AppmeshRoute) StateMust() *appmeshRouteState {
	if ar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ar.Type(), ar.LocalName()))
	}
	return ar.state
}

// AppmeshRouteArgs contains the configurations for aws_appmesh_route.
type AppmeshRouteArgs struct {
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
	// VirtualRouterName: string, required
	VirtualRouterName terra.StringValue `hcl:"virtual_router_name,attr" validate:"required"`
	// Spec: required
	Spec *appmeshroute.Spec `hcl:"spec,block" validate:"required"`
}
type appmeshRouteAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_route.
func (ar appmeshRouteAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_route.
func (ar appmeshRouteAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_route.
func (ar appmeshRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_route.
func (ar appmeshRouteAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_route.
func (ar appmeshRouteAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_route.
func (ar appmeshRouteAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_route.
func (ar appmeshRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_route.
func (ar appmeshRouteAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_route.
func (ar appmeshRouteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ar.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_route.
func (ar appmeshRouteAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ar.ref.Append("tags_all"))
}

// VirtualRouterName returns a reference to field virtual_router_name of aws_appmesh_route.
func (ar appmeshRouteAttributes) VirtualRouterName() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("virtual_router_name"))
}

func (ar appmeshRouteAttributes) Spec() terra.ListValue[appmeshroute.SpecAttributes] {
	return terra.ReferenceAsList[appmeshroute.SpecAttributes](ar.ref.Append("spec"))
}

type appmeshRouteState struct {
	Arn               string                   `json:"arn"`
	CreatedDate       string                   `json:"created_date"`
	Id                string                   `json:"id"`
	LastUpdatedDate   string                   `json:"last_updated_date"`
	MeshName          string                   `json:"mesh_name"`
	MeshOwner         string                   `json:"mesh_owner"`
	Name              string                   `json:"name"`
	ResourceOwner     string                   `json:"resource_owner"`
	Tags              map[string]string        `json:"tags"`
	TagsAll           map[string]string        `json:"tags_all"`
	VirtualRouterName string                   `json:"virtual_router_name"`
	Spec              []appmeshroute.SpecState `json:"spec"`
}
