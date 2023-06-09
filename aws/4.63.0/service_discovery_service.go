// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicediscoveryservice "github.com/golingon/terraproviders/aws/4.63.0/servicediscoveryservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDiscoveryService creates a new instance of [ServiceDiscoveryService].
func NewServiceDiscoveryService(name string, args ServiceDiscoveryServiceArgs) *ServiceDiscoveryService {
	return &ServiceDiscoveryService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDiscoveryService)(nil)

// ServiceDiscoveryService represents the Terraform resource aws_service_discovery_service.
type ServiceDiscoveryService struct {
	Name      string
	Args      ServiceDiscoveryServiceArgs
	state     *serviceDiscoveryServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDiscoveryService].
func (sds *ServiceDiscoveryService) Type() string {
	return "aws_service_discovery_service"
}

// LocalName returns the local name for [ServiceDiscoveryService].
func (sds *ServiceDiscoveryService) LocalName() string {
	return sds.Name
}

// Configuration returns the configuration (args) for [ServiceDiscoveryService].
func (sds *ServiceDiscoveryService) Configuration() interface{} {
	return sds.Args
}

// DependOn is used for other resources to depend on [ServiceDiscoveryService].
func (sds *ServiceDiscoveryService) DependOn() terra.Reference {
	return terra.ReferenceResource(sds)
}

// Dependencies returns the list of resources [ServiceDiscoveryService] depends_on.
func (sds *ServiceDiscoveryService) Dependencies() terra.Dependencies {
	return sds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDiscoveryService].
func (sds *ServiceDiscoveryService) LifecycleManagement() *terra.Lifecycle {
	return sds.Lifecycle
}

// Attributes returns the attributes for [ServiceDiscoveryService].
func (sds *ServiceDiscoveryService) Attributes() serviceDiscoveryServiceAttributes {
	return serviceDiscoveryServiceAttributes{ref: terra.ReferenceResource(sds)}
}

// ImportState imports the given attribute values into [ServiceDiscoveryService]'s state.
func (sds *ServiceDiscoveryService) ImportState(av io.Reader) error {
	sds.state = &serviceDiscoveryServiceState{}
	if err := json.NewDecoder(av).Decode(sds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sds.Type(), sds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDiscoveryService] has state.
func (sds *ServiceDiscoveryService) State() (*serviceDiscoveryServiceState, bool) {
	return sds.state, sds.state != nil
}

// StateMust returns the state for [ServiceDiscoveryService]. Panics if the state is nil.
func (sds *ServiceDiscoveryService) StateMust() *serviceDiscoveryServiceState {
	if sds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sds.Type(), sds.LocalName()))
	}
	return sds.state
}

// ServiceDiscoveryServiceArgs contains the configurations for aws_service_discovery_service.
type ServiceDiscoveryServiceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, optional
	NamespaceId terra.StringValue `hcl:"namespace_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// DnsConfig: optional
	DnsConfig *servicediscoveryservice.DnsConfig `hcl:"dns_config,block"`
	// HealthCheckConfig: optional
	HealthCheckConfig *servicediscoveryservice.HealthCheckConfig `hcl:"health_check_config,block"`
	// HealthCheckCustomConfig: optional
	HealthCheckCustomConfig *servicediscoveryservice.HealthCheckCustomConfig `hcl:"health_check_custom_config,block"`
}
type serviceDiscoveryServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("arn"))
}

// Description returns a reference to field description of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("description"))
}

// ForceDestroy returns a reference to field force_destroy of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(sds.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("id"))
}

// Name returns a reference to field name of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("namespace_id"))
}

// Tags returns a reference to field tags of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sds.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_service_discovery_service.
func (sds serviceDiscoveryServiceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("type"))
}

func (sds serviceDiscoveryServiceAttributes) DnsConfig() terra.ListValue[servicediscoveryservice.DnsConfigAttributes] {
	return terra.ReferenceAsList[servicediscoveryservice.DnsConfigAttributes](sds.ref.Append("dns_config"))
}

func (sds serviceDiscoveryServiceAttributes) HealthCheckConfig() terra.ListValue[servicediscoveryservice.HealthCheckConfigAttributes] {
	return terra.ReferenceAsList[servicediscoveryservice.HealthCheckConfigAttributes](sds.ref.Append("health_check_config"))
}

func (sds serviceDiscoveryServiceAttributes) HealthCheckCustomConfig() terra.ListValue[servicediscoveryservice.HealthCheckCustomConfigAttributes] {
	return terra.ReferenceAsList[servicediscoveryservice.HealthCheckCustomConfigAttributes](sds.ref.Append("health_check_custom_config"))
}

type serviceDiscoveryServiceState struct {
	Arn                     string                                                 `json:"arn"`
	Description             string                                                 `json:"description"`
	ForceDestroy            bool                                                   `json:"force_destroy"`
	Id                      string                                                 `json:"id"`
	Name                    string                                                 `json:"name"`
	NamespaceId             string                                                 `json:"namespace_id"`
	Tags                    map[string]string                                      `json:"tags"`
	TagsAll                 map[string]string                                      `json:"tags_all"`
	Type                    string                                                 `json:"type"`
	DnsConfig               []servicediscoveryservice.DnsConfigState               `json:"dns_config"`
	HealthCheckConfig       []servicediscoveryservice.HealthCheckConfigState       `json:"health_check_config"`
	HealthCheckCustomConfig []servicediscoveryservice.HealthCheckCustomConfigState `json:"health_check_custom_config"`
}
