// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshvirtualservice "github.com/golingon/terraproviders/aws/5.7.0/appmeshvirtualservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppmeshVirtualService creates a new instance of [AppmeshVirtualService].
func NewAppmeshVirtualService(name string, args AppmeshVirtualServiceArgs) *AppmeshVirtualService {
	return &AppmeshVirtualService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshVirtualService)(nil)

// AppmeshVirtualService represents the Terraform resource aws_appmesh_virtual_service.
type AppmeshVirtualService struct {
	Name      string
	Args      AppmeshVirtualServiceArgs
	state     *appmeshVirtualServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppmeshVirtualService].
func (avs *AppmeshVirtualService) Type() string {
	return "aws_appmesh_virtual_service"
}

// LocalName returns the local name for [AppmeshVirtualService].
func (avs *AppmeshVirtualService) LocalName() string {
	return avs.Name
}

// Configuration returns the configuration (args) for [AppmeshVirtualService].
func (avs *AppmeshVirtualService) Configuration() interface{} {
	return avs.Args
}

// DependOn is used for other resources to depend on [AppmeshVirtualService].
func (avs *AppmeshVirtualService) DependOn() terra.Reference {
	return terra.ReferenceResource(avs)
}

// Dependencies returns the list of resources [AppmeshVirtualService] depends_on.
func (avs *AppmeshVirtualService) Dependencies() terra.Dependencies {
	return avs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppmeshVirtualService].
func (avs *AppmeshVirtualService) LifecycleManagement() *terra.Lifecycle {
	return avs.Lifecycle
}

// Attributes returns the attributes for [AppmeshVirtualService].
func (avs *AppmeshVirtualService) Attributes() appmeshVirtualServiceAttributes {
	return appmeshVirtualServiceAttributes{ref: terra.ReferenceResource(avs)}
}

// ImportState imports the given attribute values into [AppmeshVirtualService]'s state.
func (avs *AppmeshVirtualService) ImportState(av io.Reader) error {
	avs.state = &appmeshVirtualServiceState{}
	if err := json.NewDecoder(av).Decode(avs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avs.Type(), avs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppmeshVirtualService] has state.
func (avs *AppmeshVirtualService) State() (*appmeshVirtualServiceState, bool) {
	return avs.state, avs.state != nil
}

// StateMust returns the state for [AppmeshVirtualService]. Panics if the state is nil.
func (avs *AppmeshVirtualService) StateMust() *appmeshVirtualServiceState {
	if avs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avs.Type(), avs.LocalName()))
	}
	return avs.state
}

// AppmeshVirtualServiceArgs contains the configurations for aws_appmesh_virtual_service.
type AppmeshVirtualServiceArgs struct {
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
	Spec *appmeshvirtualservice.Spec `hcl:"spec,block" validate:"required"`
}
type appmeshVirtualServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avs.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_virtual_service.
func (avs appmeshVirtualServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avs.ref.Append("tags_all"))
}

func (avs appmeshVirtualServiceAttributes) Spec() terra.ListValue[appmeshvirtualservice.SpecAttributes] {
	return terra.ReferenceAsList[appmeshvirtualservice.SpecAttributes](avs.ref.Append("spec"))
}

type appmeshVirtualServiceState struct {
	Arn             string                            `json:"arn"`
	CreatedDate     string                            `json:"created_date"`
	Id              string                            `json:"id"`
	LastUpdatedDate string                            `json:"last_updated_date"`
	MeshName        string                            `json:"mesh_name"`
	MeshOwner       string                            `json:"mesh_owner"`
	Name            string                            `json:"name"`
	ResourceOwner   string                            `json:"resource_owner"`
	Tags            map[string]string                 `json:"tags"`
	TagsAll         map[string]string                 `json:"tags_all"`
	Spec            []appmeshvirtualservice.SpecState `json:"spec"`
}
