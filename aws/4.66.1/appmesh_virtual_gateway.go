// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appmeshvirtualgateway "github.com/golingon/terraproviders/aws/4.66.1/appmeshvirtualgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppmeshVirtualGateway creates a new instance of [AppmeshVirtualGateway].
func NewAppmeshVirtualGateway(name string, args AppmeshVirtualGatewayArgs) *AppmeshVirtualGateway {
	return &AppmeshVirtualGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppmeshVirtualGateway)(nil)

// AppmeshVirtualGateway represents the Terraform resource aws_appmesh_virtual_gateway.
type AppmeshVirtualGateway struct {
	Name      string
	Args      AppmeshVirtualGatewayArgs
	state     *appmeshVirtualGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppmeshVirtualGateway].
func (avg *AppmeshVirtualGateway) Type() string {
	return "aws_appmesh_virtual_gateway"
}

// LocalName returns the local name for [AppmeshVirtualGateway].
func (avg *AppmeshVirtualGateway) LocalName() string {
	return avg.Name
}

// Configuration returns the configuration (args) for [AppmeshVirtualGateway].
func (avg *AppmeshVirtualGateway) Configuration() interface{} {
	return avg.Args
}

// DependOn is used for other resources to depend on [AppmeshVirtualGateway].
func (avg *AppmeshVirtualGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(avg)
}

// Dependencies returns the list of resources [AppmeshVirtualGateway] depends_on.
func (avg *AppmeshVirtualGateway) Dependencies() terra.Dependencies {
	return avg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppmeshVirtualGateway].
func (avg *AppmeshVirtualGateway) LifecycleManagement() *terra.Lifecycle {
	return avg.Lifecycle
}

// Attributes returns the attributes for [AppmeshVirtualGateway].
func (avg *AppmeshVirtualGateway) Attributes() appmeshVirtualGatewayAttributes {
	return appmeshVirtualGatewayAttributes{ref: terra.ReferenceResource(avg)}
}

// ImportState imports the given attribute values into [AppmeshVirtualGateway]'s state.
func (avg *AppmeshVirtualGateway) ImportState(av io.Reader) error {
	avg.state = &appmeshVirtualGatewayState{}
	if err := json.NewDecoder(av).Decode(avg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avg.Type(), avg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppmeshVirtualGateway] has state.
func (avg *AppmeshVirtualGateway) State() (*appmeshVirtualGatewayState, bool) {
	return avg.state, avg.state != nil
}

// StateMust returns the state for [AppmeshVirtualGateway]. Panics if the state is nil.
func (avg *AppmeshVirtualGateway) StateMust() *appmeshVirtualGatewayState {
	if avg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avg.Type(), avg.LocalName()))
	}
	return avg.state
}

// AppmeshVirtualGatewayArgs contains the configurations for aws_appmesh_virtual_gateway.
type AppmeshVirtualGatewayArgs struct {
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
	Spec *appmeshvirtualgateway.Spec `hcl:"spec,block" validate:"required"`
}
type appmeshVirtualGatewayAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(avg.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appmesh_virtual_gateway.
func (avg appmeshVirtualGatewayAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avg.ref.Append("tags_all"))
}

func (avg appmeshVirtualGatewayAttributes) Spec() terra.ListValue[appmeshvirtualgateway.SpecAttributes] {
	return terra.ReferenceAsList[appmeshvirtualgateway.SpecAttributes](avg.ref.Append("spec"))
}

type appmeshVirtualGatewayState struct {
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
	Spec            []appmeshvirtualgateway.SpecState `json:"spec"`
}
