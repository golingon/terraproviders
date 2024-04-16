// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_neptune_cluster_endpoint

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

// Resource represents the Terraform resource aws_neptune_cluster_endpoint.
type Resource struct {
	Name      string
	Args      Args
	state     *awsNeptuneClusterEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ance *Resource) Type() string {
	return "aws_neptune_cluster_endpoint"
}

// LocalName returns the local name for [Resource].
func (ance *Resource) LocalName() string {
	return ance.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ance *Resource) Configuration() interface{} {
	return ance.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ance *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ance)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ance *Resource) Dependencies() terra.Dependencies {
	return ance.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ance *Resource) LifecycleManagement() *terra.Lifecycle {
	return ance.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ance *Resource) Attributes() awsNeptuneClusterEndpointAttributes {
	return awsNeptuneClusterEndpointAttributes{ref: terra.ReferenceResource(ance)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ance *Resource) ImportState(state io.Reader) error {
	ance.state = &awsNeptuneClusterEndpointState{}
	if err := json.NewDecoder(state).Decode(ance.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ance.Type(), ance.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ance *Resource) State() (*awsNeptuneClusterEndpointState, bool) {
	return ance.state, ance.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ance *Resource) StateMust() *awsNeptuneClusterEndpointState {
	if ance.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ance.Type(), ance.LocalName()))
	}
	return ance.state
}

// Args contains the configurations for aws_neptune_cluster_endpoint.
type Args struct {
	// ClusterEndpointIdentifier: string, required
	ClusterEndpointIdentifier terra.StringValue `hcl:"cluster_endpoint_identifier,attr" validate:"required"`
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// EndpointType: string, required
	EndpointType terra.StringValue `hcl:"endpoint_type,attr" validate:"required"`
	// ExcludedMembers: set of string, optional
	ExcludedMembers terra.SetValue[terra.StringValue] `hcl:"excluded_members,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StaticMembers: set of string, optional
	StaticMembers terra.SetValue[terra.StringValue] `hcl:"static_members,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsNeptuneClusterEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ance.ref.Append("arn"))
}

// ClusterEndpointIdentifier returns a reference to field cluster_endpoint_identifier of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) ClusterEndpointIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ance.ref.Append("cluster_endpoint_identifier"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(ance.ref.Append("cluster_identifier"))
}

// Endpoint returns a reference to field endpoint of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ance.ref.Append("endpoint"))
}

// EndpointType returns a reference to field endpoint_type of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(ance.ref.Append("endpoint_type"))
}

// ExcludedMembers returns a reference to field excluded_members of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) ExcludedMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ance.ref.Append("excluded_members"))
}

// Id returns a reference to field id of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ance.ref.Append("id"))
}

// StaticMembers returns a reference to field static_members of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) StaticMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ance.ref.Append("static_members"))
}

// Tags returns a reference to field tags of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ance.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_neptune_cluster_endpoint.
func (ance awsNeptuneClusterEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ance.ref.Append("tags_all"))
}

type awsNeptuneClusterEndpointState struct {
	Arn                       string            `json:"arn"`
	ClusterEndpointIdentifier string            `json:"cluster_endpoint_identifier"`
	ClusterIdentifier         string            `json:"cluster_identifier"`
	Endpoint                  string            `json:"endpoint"`
	EndpointType              string            `json:"endpoint_type"`
	ExcludedMembers           []string          `json:"excluded_members"`
	Id                        string            `json:"id"`
	StaticMembers             []string          `json:"static_members"`
	Tags                      map[string]string `json:"tags"`
	TagsAll                   map[string]string `json:"tags_all"`
}
