// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_emrcontainers_virtual_cluster

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

// Resource represents the Terraform resource aws_emrcontainers_virtual_cluster.
type Resource struct {
	Name      string
	Args      Args
	state     *awsEmrcontainersVirtualClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aevc *Resource) Type() string {
	return "aws_emrcontainers_virtual_cluster"
}

// LocalName returns the local name for [Resource].
func (aevc *Resource) LocalName() string {
	return aevc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aevc *Resource) Configuration() interface{} {
	return aevc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aevc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aevc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aevc *Resource) Dependencies() terra.Dependencies {
	return aevc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aevc *Resource) LifecycleManagement() *terra.Lifecycle {
	return aevc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aevc *Resource) Attributes() awsEmrcontainersVirtualClusterAttributes {
	return awsEmrcontainersVirtualClusterAttributes{ref: terra.ReferenceResource(aevc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aevc *Resource) ImportState(state io.Reader) error {
	aevc.state = &awsEmrcontainersVirtualClusterState{}
	if err := json.NewDecoder(state).Decode(aevc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aevc.Type(), aevc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aevc *Resource) State() (*awsEmrcontainersVirtualClusterState, bool) {
	return aevc.state, aevc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aevc *Resource) StateMust() *awsEmrcontainersVirtualClusterState {
	if aevc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aevc.Type(), aevc.LocalName()))
	}
	return aevc.state
}

// Args contains the configurations for aws_emrcontainers_virtual_cluster.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ContainerProvider: required
	ContainerProvider *ContainerProvider `hcl:"container_provider,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsEmrcontainersVirtualClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_emrcontainers_virtual_cluster.
func (aevc awsEmrcontainersVirtualClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aevc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_emrcontainers_virtual_cluster.
func (aevc awsEmrcontainersVirtualClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aevc.ref.Append("id"))
}

// Name returns a reference to field name of aws_emrcontainers_virtual_cluster.
func (aevc awsEmrcontainersVirtualClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aevc.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_emrcontainers_virtual_cluster.
func (aevc awsEmrcontainersVirtualClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aevc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_emrcontainers_virtual_cluster.
func (aevc awsEmrcontainersVirtualClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aevc.ref.Append("tags_all"))
}

func (aevc awsEmrcontainersVirtualClusterAttributes) ContainerProvider() terra.ListValue[ContainerProviderAttributes] {
	return terra.ReferenceAsList[ContainerProviderAttributes](aevc.ref.Append("container_provider"))
}

func (aevc awsEmrcontainersVirtualClusterAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aevc.ref.Append("timeouts"))
}

type awsEmrcontainersVirtualClusterState struct {
	Arn               string                   `json:"arn"`
	Id                string                   `json:"id"`
	Name              string                   `json:"name"`
	Tags              map[string]string        `json:"tags"`
	TagsAll           map[string]string        `json:"tags_all"`
	ContainerProvider []ContainerProviderState `json:"container_provider"`
	Timeouts          *TimeoutsState           `json:"timeouts"`
}
