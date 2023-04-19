// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDiscoveryPrivateDnsNamespace creates a new instance of [ServiceDiscoveryPrivateDnsNamespace].
func NewServiceDiscoveryPrivateDnsNamespace(name string, args ServiceDiscoveryPrivateDnsNamespaceArgs) *ServiceDiscoveryPrivateDnsNamespace {
	return &ServiceDiscoveryPrivateDnsNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDiscoveryPrivateDnsNamespace)(nil)

// ServiceDiscoveryPrivateDnsNamespace represents the Terraform resource aws_service_discovery_private_dns_namespace.
type ServiceDiscoveryPrivateDnsNamespace struct {
	Name      string
	Args      ServiceDiscoveryPrivateDnsNamespaceArgs
	state     *serviceDiscoveryPrivateDnsNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDiscoveryPrivateDnsNamespace].
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Type() string {
	return "aws_service_discovery_private_dns_namespace"
}

// LocalName returns the local name for [ServiceDiscoveryPrivateDnsNamespace].
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) LocalName() string {
	return sdpdn.Name
}

// Configuration returns the configuration (args) for [ServiceDiscoveryPrivateDnsNamespace].
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Configuration() interface{} {
	return sdpdn.Args
}

// DependOn is used for other resources to depend on [ServiceDiscoveryPrivateDnsNamespace].
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) DependOn() terra.Reference {
	return terra.ReferenceResource(sdpdn)
}

// Dependencies returns the list of resources [ServiceDiscoveryPrivateDnsNamespace] depends_on.
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Dependencies() terra.Dependencies {
	return sdpdn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDiscoveryPrivateDnsNamespace].
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) LifecycleManagement() *terra.Lifecycle {
	return sdpdn.Lifecycle
}

// Attributes returns the attributes for [ServiceDiscoveryPrivateDnsNamespace].
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) Attributes() serviceDiscoveryPrivateDnsNamespaceAttributes {
	return serviceDiscoveryPrivateDnsNamespaceAttributes{ref: terra.ReferenceResource(sdpdn)}
}

// ImportState imports the given attribute values into [ServiceDiscoveryPrivateDnsNamespace]'s state.
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) ImportState(av io.Reader) error {
	sdpdn.state = &serviceDiscoveryPrivateDnsNamespaceState{}
	if err := json.NewDecoder(av).Decode(sdpdn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdpdn.Type(), sdpdn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDiscoveryPrivateDnsNamespace] has state.
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) State() (*serviceDiscoveryPrivateDnsNamespaceState, bool) {
	return sdpdn.state, sdpdn.state != nil
}

// StateMust returns the state for [ServiceDiscoveryPrivateDnsNamespace]. Panics if the state is nil.
func (sdpdn *ServiceDiscoveryPrivateDnsNamespace) StateMust() *serviceDiscoveryPrivateDnsNamespaceState {
	if sdpdn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdpdn.Type(), sdpdn.LocalName()))
	}
	return sdpdn.state
}

// ServiceDiscoveryPrivateDnsNamespaceArgs contains the configurations for aws_service_discovery_private_dns_namespace.
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
}
type serviceDiscoveryPrivateDnsNamespaceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sdpdn.ref.Append("arn"))
}

// Description returns a reference to field description of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sdpdn.ref.Append("description"))
}

// HostedZone returns a reference to field hosted_zone of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) HostedZone() terra.StringValue {
	return terra.ReferenceAsString(sdpdn.ref.Append("hosted_zone"))
}

// Id returns a reference to field id of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdpdn.ref.Append("id"))
}

// Name returns a reference to field name of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdpdn.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdpdn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdpdn.ref.Append("tags_all"))
}

// Vpc returns a reference to field vpc of aws_service_discovery_private_dns_namespace.
func (sdpdn serviceDiscoveryPrivateDnsNamespaceAttributes) Vpc() terra.StringValue {
	return terra.ReferenceAsString(sdpdn.ref.Append("vpc"))
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