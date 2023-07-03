// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerendpoint "github.com/golingon/terraproviders/aws/5.6.2/sagemakerendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerEndpoint creates a new instance of [SagemakerEndpoint].
func NewSagemakerEndpoint(name string, args SagemakerEndpointArgs) *SagemakerEndpoint {
	return &SagemakerEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerEndpoint)(nil)

// SagemakerEndpoint represents the Terraform resource aws_sagemaker_endpoint.
type SagemakerEndpoint struct {
	Name      string
	Args      SagemakerEndpointArgs
	state     *sagemakerEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerEndpoint].
func (se *SagemakerEndpoint) Type() string {
	return "aws_sagemaker_endpoint"
}

// LocalName returns the local name for [SagemakerEndpoint].
func (se *SagemakerEndpoint) LocalName() string {
	return se.Name
}

// Configuration returns the configuration (args) for [SagemakerEndpoint].
func (se *SagemakerEndpoint) Configuration() interface{} {
	return se.Args
}

// DependOn is used for other resources to depend on [SagemakerEndpoint].
func (se *SagemakerEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(se)
}

// Dependencies returns the list of resources [SagemakerEndpoint] depends_on.
func (se *SagemakerEndpoint) Dependencies() terra.Dependencies {
	return se.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerEndpoint].
func (se *SagemakerEndpoint) LifecycleManagement() *terra.Lifecycle {
	return se.Lifecycle
}

// Attributes returns the attributes for [SagemakerEndpoint].
func (se *SagemakerEndpoint) Attributes() sagemakerEndpointAttributes {
	return sagemakerEndpointAttributes{ref: terra.ReferenceResource(se)}
}

// ImportState imports the given attribute values into [SagemakerEndpoint]'s state.
func (se *SagemakerEndpoint) ImportState(av io.Reader) error {
	se.state = &sagemakerEndpointState{}
	if err := json.NewDecoder(av).Decode(se.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", se.Type(), se.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerEndpoint] has state.
func (se *SagemakerEndpoint) State() (*sagemakerEndpointState, bool) {
	return se.state, se.state != nil
}

// StateMust returns the state for [SagemakerEndpoint]. Panics if the state is nil.
func (se *SagemakerEndpoint) StateMust() *sagemakerEndpointState {
	if se.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", se.Type(), se.LocalName()))
	}
	return se.state
}

// SagemakerEndpointArgs contains the configurations for aws_sagemaker_endpoint.
type SagemakerEndpointArgs struct {
	// EndpointConfigName: string, required
	EndpointConfigName terra.StringValue `hcl:"endpoint_config_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DeploymentConfig: optional
	DeploymentConfig *sagemakerendpoint.DeploymentConfig `hcl:"deployment_config,block"`
}
type sagemakerEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_endpoint.
func (se sagemakerEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("arn"))
}

// EndpointConfigName returns a reference to field endpoint_config_name of aws_sagemaker_endpoint.
func (se sagemakerEndpointAttributes) EndpointConfigName() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("endpoint_config_name"))
}

// Id returns a reference to field id of aws_sagemaker_endpoint.
func (se sagemakerEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("id"))
}

// Name returns a reference to field name of aws_sagemaker_endpoint.
func (se sagemakerEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(se.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_sagemaker_endpoint.
func (se sagemakerEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](se.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_endpoint.
func (se sagemakerEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](se.ref.Append("tags_all"))
}

func (se sagemakerEndpointAttributes) DeploymentConfig() terra.ListValue[sagemakerendpoint.DeploymentConfigAttributes] {
	return terra.ReferenceAsList[sagemakerendpoint.DeploymentConfigAttributes](se.ref.Append("deployment_config"))
}

type sagemakerEndpointState struct {
	Arn                string                                    `json:"arn"`
	EndpointConfigName string                                    `json:"endpoint_config_name"`
	Id                 string                                    `json:"id"`
	Name               string                                    `json:"name"`
	Tags               map[string]string                         `json:"tags"`
	TagsAll            map[string]string                         `json:"tags_all"`
	DeploymentConfig   []sagemakerendpoint.DeploymentConfigState `json:"deployment_config"`
}
