// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_appmesh_virtual_service

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_appmesh_virtual_service.
type Resource struct {
	Name      string
	Args      Args
	state     *awsAppmeshVirtualServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aavs *Resource) Type() string {
	return "aws_appmesh_virtual_service"
}

// LocalName returns the local name for [Resource].
func (aavs *Resource) LocalName() string {
	return aavs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aavs *Resource) Configuration() interface{} {
	return aavs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aavs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aavs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aavs *Resource) Dependencies() terra.Dependencies {
	return aavs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aavs *Resource) LifecycleManagement() *terra.Lifecycle {
	return aavs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aavs *Resource) Attributes() awsAppmeshVirtualServiceAttributes {
	return awsAppmeshVirtualServiceAttributes{ref: terra.ReferenceResource(aavs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aavs *Resource) ImportState(state io.Reader) error {
	aavs.state = &awsAppmeshVirtualServiceState{}
	if err := json.NewDecoder(state).Decode(aavs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aavs.Type(), aavs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aavs *Resource) State() (*awsAppmeshVirtualServiceState, bool) {
	return aavs.state, aavs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aavs *Resource) StateMust() *awsAppmeshVirtualServiceState {
	if aavs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aavs.Type(), aavs.LocalName()))
	}
	return aavs.state
}

// Args contains the configurations for aws_appmesh_virtual_service.
type Args struct {
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
	Spec *Spec `hcl:"spec,block" validate:"required"`
}

type awsAppmeshVirtualServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(aavs.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aavs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_virtual_service.
func (aavs awsAppmeshVirtualServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aavs.ref.Append("tags_all"))
}

func (aavs awsAppmeshVirtualServiceAttributes) Spec() terra.ListValue[SpecAttributes] {
	return terra.ReferenceAsList[SpecAttributes](aavs.ref.Append("spec"))
}

type awsAppmeshVirtualServiceState struct {
	Arn             string            `json:"arn"`
	CreatedDate     string            `json:"created_date"`
	Id              string            `json:"id"`
	LastUpdatedDate string            `json:"last_updated_date"`
	MeshName        string            `json:"mesh_name"`
	MeshOwner       string            `json:"mesh_owner"`
	Name            string            `json:"name"`
	ResourceOwner   string            `json:"resource_owner"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
	Spec            []SpecState       `json:"spec"`
}
