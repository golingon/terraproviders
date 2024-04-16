// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_batch_job_definition

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

// Resource represents the Terraform resource aws_batch_job_definition.
type Resource struct {
	Name      string
	Args      Args
	state     *awsBatchJobDefinitionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (abjd *Resource) Type() string {
	return "aws_batch_job_definition"
}

// LocalName returns the local name for [Resource].
func (abjd *Resource) LocalName() string {
	return abjd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (abjd *Resource) Configuration() interface{} {
	return abjd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (abjd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(abjd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (abjd *Resource) Dependencies() terra.Dependencies {
	return abjd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (abjd *Resource) LifecycleManagement() *terra.Lifecycle {
	return abjd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (abjd *Resource) Attributes() awsBatchJobDefinitionAttributes {
	return awsBatchJobDefinitionAttributes{ref: terra.ReferenceResource(abjd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (abjd *Resource) ImportState(state io.Reader) error {
	abjd.state = &awsBatchJobDefinitionState{}
	if err := json.NewDecoder(state).Decode(abjd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", abjd.Type(), abjd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (abjd *Resource) State() (*awsBatchJobDefinitionState, bool) {
	return abjd.state, abjd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (abjd *Resource) StateMust() *awsBatchJobDefinitionState {
	if abjd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", abjd.Type(), abjd.LocalName()))
	}
	return abjd.state
}

// Args contains the configurations for aws_batch_job_definition.
type Args struct {
	// ContainerProperties: string, optional
	ContainerProperties terra.StringValue `hcl:"container_properties,attr"`
	// DeregisterOnNewRevision: bool, optional
	DeregisterOnNewRevision terra.BoolValue `hcl:"deregister_on_new_revision,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeProperties: string, optional
	NodeProperties terra.StringValue `hcl:"node_properties,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// PlatformCapabilities: set of string, optional
	PlatformCapabilities terra.SetValue[terra.StringValue] `hcl:"platform_capabilities,attr"`
	// PropagateTags: bool, optional
	PropagateTags terra.BoolValue `hcl:"propagate_tags,attr"`
	// SchedulingPriority: number, optional
	SchedulingPriority terra.NumberValue `hcl:"scheduling_priority,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// EksProperties: optional
	EksProperties *EksProperties `hcl:"eks_properties,block"`
	// RetryStrategy: optional
	RetryStrategy *RetryStrategy `hcl:"retry_strategy,block"`
	// Timeout: optional
	Timeout *Timeout `hcl:"timeout,block"`
}

type awsBatchJobDefinitionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("arn"))
}

// ArnPrefix returns a reference to field arn_prefix of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) ArnPrefix() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("arn_prefix"))
}

// ContainerProperties returns a reference to field container_properties of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) ContainerProperties() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("container_properties"))
}

// DeregisterOnNewRevision returns a reference to field deregister_on_new_revision of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) DeregisterOnNewRevision() terra.BoolValue {
	return terra.ReferenceAsBool(abjd.ref.Append("deregister_on_new_revision"))
}

// Id returns a reference to field id of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("id"))
}

// Name returns a reference to field name of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("name"))
}

// NodeProperties returns a reference to field node_properties of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) NodeProperties() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("node_properties"))
}

// Parameters returns a reference to field parameters of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abjd.ref.Append("parameters"))
}

// PlatformCapabilities returns a reference to field platform_capabilities of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) PlatformCapabilities() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](abjd.ref.Append("platform_capabilities"))
}

// PropagateTags returns a reference to field propagate_tags of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) PropagateTags() terra.BoolValue {
	return terra.ReferenceAsBool(abjd.ref.Append("propagate_tags"))
}

// Revision returns a reference to field revision of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Revision() terra.NumberValue {
	return terra.ReferenceAsNumber(abjd.ref.Append("revision"))
}

// SchedulingPriority returns a reference to field scheduling_priority of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) SchedulingPriority() terra.NumberValue {
	return terra.ReferenceAsNumber(abjd.ref.Append("scheduling_priority"))
}

// Tags returns a reference to field tags of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abjd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abjd.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_batch_job_definition.
func (abjd awsBatchJobDefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(abjd.ref.Append("type"))
}

func (abjd awsBatchJobDefinitionAttributes) EksProperties() terra.ListValue[EksPropertiesAttributes] {
	return terra.ReferenceAsList[EksPropertiesAttributes](abjd.ref.Append("eks_properties"))
}

func (abjd awsBatchJobDefinitionAttributes) RetryStrategy() terra.ListValue[RetryStrategyAttributes] {
	return terra.ReferenceAsList[RetryStrategyAttributes](abjd.ref.Append("retry_strategy"))
}

func (abjd awsBatchJobDefinitionAttributes) Timeout() terra.ListValue[TimeoutAttributes] {
	return terra.ReferenceAsList[TimeoutAttributes](abjd.ref.Append("timeout"))
}

type awsBatchJobDefinitionState struct {
	Arn                     string               `json:"arn"`
	ArnPrefix               string               `json:"arn_prefix"`
	ContainerProperties     string               `json:"container_properties"`
	DeregisterOnNewRevision bool                 `json:"deregister_on_new_revision"`
	Id                      string               `json:"id"`
	Name                    string               `json:"name"`
	NodeProperties          string               `json:"node_properties"`
	Parameters              map[string]string    `json:"parameters"`
	PlatformCapabilities    []string             `json:"platform_capabilities"`
	PropagateTags           bool                 `json:"propagate_tags"`
	Revision                float64              `json:"revision"`
	SchedulingPriority      float64              `json:"scheduling_priority"`
	Tags                    map[string]string    `json:"tags"`
	TagsAll                 map[string]string    `json:"tags_all"`
	Type                    string               `json:"type"`
	EksProperties           []EksPropertiesState `json:"eks_properties"`
	RetryStrategy           []RetryStrategyState `json:"retry_strategy"`
	Timeout                 []TimeoutState       `json:"timeout"`
}
