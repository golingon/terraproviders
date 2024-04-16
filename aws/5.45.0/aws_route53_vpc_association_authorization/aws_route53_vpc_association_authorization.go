// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53_vpc_association_authorization

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

// Resource represents the Terraform resource aws_route53_vpc_association_authorization.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRoute53VpcAssociationAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arvaa *Resource) Type() string {
	return "aws_route53_vpc_association_authorization"
}

// LocalName returns the local name for [Resource].
func (arvaa *Resource) LocalName() string {
	return arvaa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arvaa *Resource) Configuration() interface{} {
	return arvaa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arvaa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arvaa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arvaa *Resource) Dependencies() terra.Dependencies {
	return arvaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arvaa *Resource) LifecycleManagement() *terra.Lifecycle {
	return arvaa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arvaa *Resource) Attributes() awsRoute53VpcAssociationAuthorizationAttributes {
	return awsRoute53VpcAssociationAuthorizationAttributes{ref: terra.ReferenceResource(arvaa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arvaa *Resource) ImportState(state io.Reader) error {
	arvaa.state = &awsRoute53VpcAssociationAuthorizationState{}
	if err := json.NewDecoder(state).Decode(arvaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arvaa.Type(), arvaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arvaa *Resource) State() (*awsRoute53VpcAssociationAuthorizationState, bool) {
	return arvaa.state, arvaa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arvaa *Resource) StateMust() *awsRoute53VpcAssociationAuthorizationState {
	if arvaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arvaa.Type(), arvaa.LocalName()))
	}
	return arvaa.state
}

// Args contains the configurations for aws_route53_vpc_association_authorization.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// VpcRegion: string, optional
	VpcRegion terra.StringValue `hcl:"vpc_region,attr"`
	// ZoneId: string, required
	ZoneId terra.StringValue `hcl:"zone_id,attr" validate:"required"`
}

type awsRoute53VpcAssociationAuthorizationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_route53_vpc_association_authorization.
func (arvaa awsRoute53VpcAssociationAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arvaa.ref.Append("id"))
}

// VpcId returns a reference to field vpc_id of aws_route53_vpc_association_authorization.
func (arvaa awsRoute53VpcAssociationAuthorizationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(arvaa.ref.Append("vpc_id"))
}

// VpcRegion returns a reference to field vpc_region of aws_route53_vpc_association_authorization.
func (arvaa awsRoute53VpcAssociationAuthorizationAttributes) VpcRegion() terra.StringValue {
	return terra.ReferenceAsString(arvaa.ref.Append("vpc_region"))
}

// ZoneId returns a reference to field zone_id of aws_route53_vpc_association_authorization.
func (arvaa awsRoute53VpcAssociationAuthorizationAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(arvaa.ref.Append("zone_id"))
}

type awsRoute53VpcAssociationAuthorizationState struct {
	Id        string `json:"id"`
	VpcId     string `json:"vpc_id"`
	VpcRegion string `json:"vpc_region"`
	ZoneId    string `json:"zone_id"`
}
