// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNeptuneClusterEndpoint(name string, args NeptuneClusterEndpointArgs) *NeptuneClusterEndpoint {
	return &NeptuneClusterEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NeptuneClusterEndpoint)(nil)

type NeptuneClusterEndpoint struct {
	Name  string
	Args  NeptuneClusterEndpointArgs
	state *neptuneClusterEndpointState
}

func (nce *NeptuneClusterEndpoint) Type() string {
	return "aws_neptune_cluster_endpoint"
}

func (nce *NeptuneClusterEndpoint) LocalName() string {
	return nce.Name
}

func (nce *NeptuneClusterEndpoint) Configuration() interface{} {
	return nce.Args
}

func (nce *NeptuneClusterEndpoint) Attributes() neptuneClusterEndpointAttributes {
	return neptuneClusterEndpointAttributes{ref: terra.ReferenceResource(nce)}
}

func (nce *NeptuneClusterEndpoint) ImportState(av io.Reader) error {
	nce.state = &neptuneClusterEndpointState{}
	if err := json.NewDecoder(av).Decode(nce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nce.Type(), nce.LocalName(), err)
	}
	return nil
}

func (nce *NeptuneClusterEndpoint) State() (*neptuneClusterEndpointState, bool) {
	return nce.state, nce.state != nil
}

func (nce *NeptuneClusterEndpoint) StateMust() *neptuneClusterEndpointState {
	if nce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nce.Type(), nce.LocalName()))
	}
	return nce.state
}

func (nce *NeptuneClusterEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(nce)
}

type NeptuneClusterEndpointArgs struct {
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
	// DependsOn contains resources that NeptuneClusterEndpoint depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type neptuneClusterEndpointAttributes struct {
	ref terra.Reference
}

func (nce neptuneClusterEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(nce.ref.Append("arn"))
}

func (nce neptuneClusterEndpointAttributes) ClusterEndpointIdentifier() terra.StringValue {
	return terra.ReferenceString(nce.ref.Append("cluster_endpoint_identifier"))
}

func (nce neptuneClusterEndpointAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(nce.ref.Append("cluster_identifier"))
}

func (nce neptuneClusterEndpointAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(nce.ref.Append("endpoint"))
}

func (nce neptuneClusterEndpointAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceString(nce.ref.Append("endpoint_type"))
}

func (nce neptuneClusterEndpointAttributes) ExcludedMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](nce.ref.Append("excluded_members"))
}

func (nce neptuneClusterEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceString(nce.ref.Append("id"))
}

func (nce neptuneClusterEndpointAttributes) StaticMembers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](nce.ref.Append("static_members"))
}

func (nce neptuneClusterEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](nce.ref.Append("tags"))
}

func (nce neptuneClusterEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](nce.ref.Append("tags_all"))
}

type neptuneClusterEndpointState struct {
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
