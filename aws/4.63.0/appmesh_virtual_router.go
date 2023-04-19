// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshvirtualrouter "github.com/golingon/terraproviders/aws/4.63.0/appmeshvirtualrouter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppmeshVirtualRouter creates a new instance of [AppmeshVirtualRouter].
func NewAppmeshVirtualRouter(name string, args AppmeshVirtualRouterArgs) *AppmeshVirtualRouter {
	return &AppmeshVirtualRouter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshVirtualRouter)(nil)

// AppmeshVirtualRouter represents the Terraform resource aws_appmesh_virtual_router.
type AppmeshVirtualRouter struct {
	Name      string
	Args      AppmeshVirtualRouterArgs
	state     *appmeshVirtualRouterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppmeshVirtualRouter].
func (avr *AppmeshVirtualRouter) Type() string {
	return "aws_appmesh_virtual_router"
}

// LocalName returns the local name for [AppmeshVirtualRouter].
func (avr *AppmeshVirtualRouter) LocalName() string {
	return avr.Name
}

// Configuration returns the configuration (args) for [AppmeshVirtualRouter].
func (avr *AppmeshVirtualRouter) Configuration() interface{} {
	return avr.Args
}

// DependOn is used for other resources to depend on [AppmeshVirtualRouter].
func (avr *AppmeshVirtualRouter) DependOn() terra.Reference {
	return terra.ReferenceResource(avr)
}

// Dependencies returns the list of resources [AppmeshVirtualRouter] depends_on.
func (avr *AppmeshVirtualRouter) Dependencies() terra.Dependencies {
	return avr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppmeshVirtualRouter].
func (avr *AppmeshVirtualRouter) LifecycleManagement() *terra.Lifecycle {
	return avr.Lifecycle
}

// Attributes returns the attributes for [AppmeshVirtualRouter].
func (avr *AppmeshVirtualRouter) Attributes() appmeshVirtualRouterAttributes {
	return appmeshVirtualRouterAttributes{ref: terra.ReferenceResource(avr)}
}

// ImportState imports the given attribute values into [AppmeshVirtualRouter]'s state.
func (avr *AppmeshVirtualRouter) ImportState(av io.Reader) error {
	avr.state = &appmeshVirtualRouterState{}
	if err := json.NewDecoder(av).Decode(avr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avr.Type(), avr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppmeshVirtualRouter] has state.
func (avr *AppmeshVirtualRouter) State() (*appmeshVirtualRouterState, bool) {
	return avr.state, avr.state != nil
}

// StateMust returns the state for [AppmeshVirtualRouter]. Panics if the state is nil.
func (avr *AppmeshVirtualRouter) StateMust() *appmeshVirtualRouterState {
	if avr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avr.Type(), avr.LocalName()))
	}
	return avr.state
}

// AppmeshVirtualRouterArgs contains the configurations for aws_appmesh_virtual_router.
type AppmeshVirtualRouterArgs struct {
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
	// Spec: required
	Spec *appmeshvirtualrouter.Spec `hcl:"spec,block" validate:"required"`
}
type appmeshVirtualRouterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avr.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_virtual_router.
func (avr appmeshVirtualRouterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avr.ref.Append("tags_all"))
}

func (avr appmeshVirtualRouterAttributes) Spec() terra.ListValue[appmeshvirtualrouter.SpecAttributes] {
	return terra.ReferenceAsList[appmeshvirtualrouter.SpecAttributes](avr.ref.Append("spec"))
}

type appmeshVirtualRouterState struct {
	Arn             string                           `json:"arn"`
	CreatedDate     string                           `json:"created_date"`
	Id              string                           `json:"id"`
	LastUpdatedDate string                           `json:"last_updated_date"`
	MeshName        string                           `json:"mesh_name"`
	MeshOwner       string                           `json:"mesh_owner"`
	Name            string                           `json:"name"`
	ResourceOwner   string                           `json:"resource_owner"`
	Tags            map[string]string                `json:"tags"`
	TagsAll         map[string]string                `json:"tags_all"`
	Spec            []appmeshvirtualrouter.SpecState `json:"spec"`
}
