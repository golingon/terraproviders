// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewServiceDiscoveryPrivateDnsNamespace(name string, args ServiceDiscoveryPrivateDnsNamespaceArgs) *ServiceDiscoveryPrivateDnsNamespace {
	return &ServiceDiscoveryPrivateDnsNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDiscoveryPrivateDnsNamespace)(nil)

type ServiceDiscoveryPrivateDnsNamespace struct {
	Name  string
	Args  ServiceDiscoveryPrivateDnsNamespaceArgs
	state *serviceDiscoveryPrivateDnsNamespaceState
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Type() string {
	return "aws_service_discovery_private_dns_namespace"
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) LocalName() string {
	return sdpdn.Name
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Configuration() interface{} {
	return sdpdn.Args
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Attributes() serviceDiscoveryPrivateDnsNamespaceAttributes {
	return serviceDiscoveryPrivateDnsNamespaceAttributes{ref: terra.ReferenceResource(sdpdn)}
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) ImportState(av io.Reader) error {
	sdpdn.state = &serviceDiscoveryPrivateDnsNamespaceState{}
	if err := json.NewDecoder(av).Decode(sdpdn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdpdn.Type(), sdpdn.LocalName(), err)
	}
	return nil
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) State() (*serviceDiscoveryPrivateDnsNamespaceState, bool) {
	return sdpdn.state, sdpdn.state != nil
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) StateMust() *serviceDiscoveryPrivateDnsNamespaceState {
	if sdpdn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdpdn.Type(), sdpdn.LocalName()))
	}
	return sdpdn.state
}

func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) DependOn() terra.Reference {
	return terra.ReferenceResource(sdpdn)
}

type ServiceDiscoveryPrivateDnsNamespaceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Vpc: string, required
	Vpc terra.StringValue `hcl:"vpc,attr" validate:"required"`
	// DependsOn contains resources that ServiceDiscoveryPrivateDnsNamespace depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type serviceDiscoveryPrivateDnsNamespaceAttributes struct {
	ref terra.Reference
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sdpdn.ref.Append("arn"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Description() terra.StringValue {
	return terra.ReferenceString(sdpdn.ref.Append("description"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) HostedZone() terra.StringValue {
	return terra.ReferenceString(sdpdn.ref.Append("hosted_zone"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sdpdn.ref.Append("id"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sdpdn.ref.Append("name"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sdpdn.ref.Append("tags"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sdpdn.ref.Append("tags_all"))
}

func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Vpc() terra.StringValue {
	return terra.ReferenceString(sdpdn.ref.Append("vpc"))
}

type serviceDiscoveryPrivateDnsNamespaceState struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	HostedZone  string            `json:"hosted_zone"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
	Vpc         string            `json:"vpc"`
}
