// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	batchcomputeenvironment "github.com/golingon/terraproviders/aws/5.8.0/batchcomputeenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBatchComputeEnvironment creates a new instance of [BatchComputeEnvironment].
func NewBatchComputeEnvironment(name string, args BatchComputeEnvironmentArgs) *BatchComputeEnvironment {
	return &BatchComputeEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BatchComputeEnvironment)(nil)

// BatchComputeEnvironment represents the Terraform resource aws_batch_compute_environment.
type BatchComputeEnvironment struct {
	Name      string
	Args      BatchComputeEnvironmentArgs
	state     *batchComputeEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BatchComputeEnvironment].
func (bce *BatchComputeEnvironment) Type() string {
	return "aws_batch_compute_environment"
}

// LocalName returns the local name for [BatchComputeEnvironment].
func (bce *BatchComputeEnvironment) LocalName() string {
	return bce.Name
}

// Configuration returns the configuration (args) for [BatchComputeEnvironment].
func (bce *BatchComputeEnvironment) Configuration() interface{} {
	return bce.Args
}

// DependOn is used for other resources to depend on [BatchComputeEnvironment].
func (bce *BatchComputeEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(bce)
}

// Dependencies returns the list of resources [BatchComputeEnvironment] depends_on.
func (bce *BatchComputeEnvironment) Dependencies() terra.Dependencies {
	return bce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BatchComputeEnvironment].
func (bce *BatchComputeEnvironment) LifecycleManagement() *terra.Lifecycle {
	return bce.Lifecycle
}

// Attributes returns the attributes for [BatchComputeEnvironment].
func (bce *BatchComputeEnvironment) Attributes() batchComputeEnvironmentAttributes {
	return batchComputeEnvironmentAttributes{ref: terra.ReferenceResource(bce)}
}

// ImportState imports the given attribute values into [BatchComputeEnvironment]'s state.
func (bce *BatchComputeEnvironment) ImportState(av io.Reader) error {
	bce.state = &batchComputeEnvironmentState{}
	if err := json.NewDecoder(av).Decode(bce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bce.Type(), bce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BatchComputeEnvironment] has state.
func (bce *BatchComputeEnvironment) State() (*batchComputeEnvironmentState, bool) {
	return bce.state, bce.state != nil
}

// StateMust returns the state for [BatchComputeEnvironment]. Panics if the state is nil.
func (bce *BatchComputeEnvironment) StateMust() *batchComputeEnvironmentState {
	if bce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bce.Type(), bce.LocalName()))
	}
	return bce.state
}

// BatchComputeEnvironmentArgs contains the configurations for aws_batch_compute_environment.
type BatchComputeEnvironmentArgs struct {
	// ComputeEnvironmentName: string, optional
	ComputeEnvironmentName terra.StringValue `hcl:"compute_environment_name,attr"`
	// ComputeEnvironmentNamePrefix: string, optional
	ComputeEnvironmentNamePrefix terra.StringValue `hcl:"compute_environment_name_prefix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceRole: string, optional
	ServiceRole terra.StringValue `hcl:"service_role,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// ComputeResources: optional
	ComputeResources *batchcomputeenvironment.ComputeResources `hcl:"compute_resources,block"`
	// EksConfiguration: optional
	EksConfiguration *batchcomputeenvironment.EksConfiguration `hcl:"eks_configuration,block"`
}
type batchComputeEnvironmentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("arn"))
}

// ComputeEnvironmentName returns a reference to field compute_environment_name of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) ComputeEnvironmentName() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("compute_environment_name"))
}

// ComputeEnvironmentNamePrefix returns a reference to field compute_environment_name_prefix of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) ComputeEnvironmentNamePrefix() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("compute_environment_name_prefix"))
}

// EcsClusterArn returns a reference to field ecs_cluster_arn of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) EcsClusterArn() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("ecs_cluster_arn"))
}

// Id returns a reference to field id of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("id"))
}

// ServiceRole returns a reference to field service_role of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("service_role"))
}

// State returns a reference to field state of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("state"))
}

// Status returns a reference to field status of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("status"))
}

// StatusReason returns a reference to field status_reason of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("status_reason"))
}

// Tags returns a reference to field tags of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bce.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bce.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_batch_compute_environment.
func (bce batchComputeEnvironmentAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bce.ref.Append("type"))
}

func (bce batchComputeEnvironmentAttributes) ComputeResources() terra.ListValue[batchcomputeenvironment.ComputeResourcesAttributes] {
	return terra.ReferenceAsList[batchcomputeenvironment.ComputeResourcesAttributes](bce.ref.Append("compute_resources"))
}

func (bce batchComputeEnvironmentAttributes) EksConfiguration() terra.ListValue[batchcomputeenvironment.EksConfigurationAttributes] {
	return terra.ReferenceAsList[batchcomputeenvironment.EksConfigurationAttributes](bce.ref.Append("eks_configuration"))
}

type batchComputeEnvironmentState struct {
	Arn                          string                                          `json:"arn"`
	ComputeEnvironmentName       string                                          `json:"compute_environment_name"`
	ComputeEnvironmentNamePrefix string                                          `json:"compute_environment_name_prefix"`
	EcsClusterArn                string                                          `json:"ecs_cluster_arn"`
	Id                           string                                          `json:"id"`
	ServiceRole                  string                                          `json:"service_role"`
	State                        string                                          `json:"state"`
	Status                       string                                          `json:"status"`
	StatusReason                 string                                          `json:"status_reason"`
	Tags                         map[string]string                               `json:"tags"`
	TagsAll                      map[string]string                               `json:"tags_all"`
	Type                         string                                          `json:"type"`
	ComputeResources             []batchcomputeenvironment.ComputeResourcesState `json:"compute_resources"`
	EksConfiguration             []batchcomputeenvironment.EksConfigurationState `json:"eks_configuration"`
}