// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudcontrolapi_resource

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

// Resource represents the Terraform resource aws_cloudcontrolapi_resource.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCloudcontrolapiResourceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acr *Resource) Type() string {
	return "aws_cloudcontrolapi_resource"
}

// LocalName returns the local name for [Resource].
func (acr *Resource) LocalName() string {
	return acr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acr *Resource) Configuration() interface{} {
	return acr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acr *Resource) Dependencies() terra.Dependencies {
	return acr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acr *Resource) LifecycleManagement() *terra.Lifecycle {
	return acr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acr *Resource) Attributes() awsCloudcontrolapiResourceAttributes {
	return awsCloudcontrolapiResourceAttributes{ref: terra.ReferenceResource(acr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acr *Resource) ImportState(state io.Reader) error {
	acr.state = &awsCloudcontrolapiResourceState{}
	if err := json.NewDecoder(state).Decode(acr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acr.Type(), acr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acr *Resource) State() (*awsCloudcontrolapiResourceState, bool) {
	return acr.state, acr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acr *Resource) StateMust() *awsCloudcontrolapiResourceState {
	if acr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acr.Type(), acr.LocalName()))
	}
	return acr.state
}

// Args contains the configurations for aws_cloudcontrolapi_resource.
type Args struct {
	// DesiredState: string, required
	DesiredState terra.StringValue `hcl:"desired_state,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// Schema: string, optional
	Schema terra.StringValue `hcl:"schema,attr"`
	// TypeName: string, required
	TypeName terra.StringValue `hcl:"type_name,attr" validate:"required"`
	// TypeVersionId: string, optional
	TypeVersionId terra.StringValue `hcl:"type_version_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsCloudcontrolapiResourceAttributes struct {
	ref terra.Reference
}

// DesiredState returns a reference to field desired_state of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) DesiredState() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("desired_state"))
}

// Id returns a reference to field id of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("id"))
}

// Properties returns a reference to field properties of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) Properties() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("properties"))
}

// RoleArn returns a reference to field role_arn of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("role_arn"))
}

// Schema returns a reference to field schema of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("schema"))
}

// TypeName returns a reference to field type_name of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) TypeName() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("type_name"))
}

// TypeVersionId returns a reference to field type_version_id of aws_cloudcontrolapi_resource.
func (acr awsCloudcontrolapiResourceAttributes) TypeVersionId() terra.StringValue {
	return terra.ReferenceAsString(acr.ref.Append("type_version_id"))
}

func (acr awsCloudcontrolapiResourceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acr.ref.Append("timeouts"))
}

type awsCloudcontrolapiResourceState struct {
	DesiredState  string         `json:"desired_state"`
	Id            string         `json:"id"`
	Properties    string         `json:"properties"`
	RoleArn       string         `json:"role_arn"`
	Schema        string         `json:"schema"`
	TypeName      string         `json:"type_name"`
	TypeVersionId string         `json:"type_version_id"`
	Timeouts      *TimeoutsState `json:"timeouts"`
}
