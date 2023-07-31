// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshvirtualnode "github.com/golingon/terraproviders/aws/5.10.0/appmeshvirtualnode"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppmeshVirtualNode creates a new instance of [AppmeshVirtualNode].
func NewAppmeshVirtualNode(name string, args AppmeshVirtualNodeArgs) *AppmeshVirtualNode {
	return &AppmeshVirtualNode{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshVirtualNode)(nil)

// AppmeshVirtualNode represents the Terraform resource aws_appmesh_virtual_node.
type AppmeshVirtualNode struct {
	Name      string
	Args      AppmeshVirtualNodeArgs
	state     *appmeshVirtualNodeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppmeshVirtualNode].
func (avn *AppmeshVirtualNode) Type() string {
	return "aws_appmesh_virtual_node"
}

// LocalName returns the local name for [AppmeshVirtualNode].
func (avn *AppmeshVirtualNode) LocalName() string {
	return avn.Name
}

// Configuration returns the configuration (args) for [AppmeshVirtualNode].
func (avn *AppmeshVirtualNode) Configuration() interface{} {
	return avn.Args
}

// DependOn is used for other resources to depend on [AppmeshVirtualNode].
func (avn *AppmeshVirtualNode) DependOn() terra.Reference {
	return terra.ReferenceResource(avn)
}

// Dependencies returns the list of resources [AppmeshVirtualNode] depends_on.
func (avn *AppmeshVirtualNode) Dependencies() terra.Dependencies {
	return avn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppmeshVirtualNode].
func (avn *AppmeshVirtualNode) LifecycleManagement() *terra.Lifecycle {
	return avn.Lifecycle
}

// Attributes returns the attributes for [AppmeshVirtualNode].
func (avn *AppmeshVirtualNode) Attributes() appmeshVirtualNodeAttributes {
	return appmeshVirtualNodeAttributes{ref: terra.ReferenceResource(avn)}
}

// ImportState imports the given attribute values into [AppmeshVirtualNode]'s state.
func (avn *AppmeshVirtualNode) ImportState(av io.Reader) error {
	avn.state = &appmeshVirtualNodeState{}
	if err := json.NewDecoder(av).Decode(avn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avn.Type(), avn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppmeshVirtualNode] has state.
func (avn *AppmeshVirtualNode) State() (*appmeshVirtualNodeState, bool) {
	return avn.state, avn.state != nil
}

// StateMust returns the state for [AppmeshVirtualNode]. Panics if the state is nil.
func (avn *AppmeshVirtualNode) StateMust() *appmeshVirtualNodeState {
	if avn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avn.Type(), avn.LocalName()))
	}
	return avn.state
}

// AppmeshVirtualNodeArgs contains the configurations for aws_appmesh_virtual_node.
type AppmeshVirtualNodeArgs struct {
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
	Spec *appmeshvirtualnode.Spec `hcl:"spec,block" validate:"required"`
}
type appmeshVirtualNodeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avn.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_virtual_node.
func (avn appmeshVirtualNodeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avn.ref.Append("tags_all"))
}

func (avn appmeshVirtualNodeAttributes) Spec() terra.ListValue[appmeshvirtualnode.SpecAttributes] {
	return terra.ReferenceAsList[appmeshvirtualnode.SpecAttributes](avn.ref.Append("spec"))
}

type appmeshVirtualNodeState struct {
	Arn             string                         `json:"arn"`
	CreatedDate     string                         `json:"created_date"`
	Id              string                         `json:"id"`
	LastUpdatedDate string                         `json:"last_updated_date"`
	MeshName        string                         `json:"mesh_name"`
	MeshOwner       string                         `json:"mesh_owner"`
	Name            string                         `json:"name"`
	ResourceOwner   string                         `json:"resource_owner"`
	Tags            map[string]string              `json:"tags"`
	TagsAll         map[string]string              `json:"tags_all"`
	Spec            []appmeshvirtualnode.SpecState `json:"spec"`
}
