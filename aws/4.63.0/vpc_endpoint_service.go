// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcendpointservice "github.com/golingon/terraproviders/aws/4.63.0/vpcendpointservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcEndpointService creates a new instance of [VpcEndpointService].
func NewVpcEndpointService(name string, args VpcEndpointServiceArgs) *VpcEndpointService {
	return &VpcEndpointService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcEndpointService)(nil)

// VpcEndpointService represents the Terraform resource aws_vpc_endpoint_service.
type VpcEndpointService struct {
	Name      string
	Args      VpcEndpointServiceArgs
	state     *vpcEndpointServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcEndpointService].
func (ves *VpcEndpointService) Type() string {
	return "aws_vpc_endpoint_service"
}

// LocalName returns the local name for [VpcEndpointService].
func (ves *VpcEndpointService) LocalName() string {
	return ves.Name
}

// Configuration returns the configuration (args) for [VpcEndpointService].
func (ves *VpcEndpointService) Configuration() interface{} {
	return ves.Args
}

// DependOn is used for other resources to depend on [VpcEndpointService].
func (ves *VpcEndpointService) DependOn() terra.Reference {
	return terra.ReferenceResource(ves)
}

// Dependencies returns the list of resources [VpcEndpointService] depends_on.
func (ves *VpcEndpointService) Dependencies() terra.Dependencies {
	return ves.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcEndpointService].
func (ves *VpcEndpointService) LifecycleManagement() *terra.Lifecycle {
	return ves.Lifecycle
}

// Attributes returns the attributes for [VpcEndpointService].
func (ves *VpcEndpointService) Attributes() vpcEndpointServiceAttributes {
	return vpcEndpointServiceAttributes{ref: terra.ReferenceResource(ves)}
}

// ImportState imports the given attribute values into [VpcEndpointService]'s state.
func (ves *VpcEndpointService) ImportState(av io.Reader) error {
	ves.state = &vpcEndpointServiceState{}
	if err := json.NewDecoder(av).Decode(ves.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ves.Type(), ves.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcEndpointService] has state.
func (ves *VpcEndpointService) State() (*vpcEndpointServiceState, bool) {
	return ves.state, ves.state != nil
}

// StateMust returns the state for [VpcEndpointService]. Panics if the state is nil.
func (ves *VpcEndpointService) StateMust() *vpcEndpointServiceState {
	if ves.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ves.Type(), ves.LocalName()))
	}
	return ves.state
}

// VpcEndpointServiceArgs contains the configurations for aws_vpc_endpoint_service.
type VpcEndpointServiceArgs struct {
	// AcceptanceRequired: bool, required
	AcceptanceRequired terra.BoolValue `hcl:"acceptance_required,attr" validate:"required"`
	// AllowedPrincipals: set of string, optional
	AllowedPrincipals terra.SetValue[terra.StringValue] `hcl:"allowed_principals,attr"`
	// GatewayLoadBalancerArns: set of string, optional
	GatewayLoadBalancerArns terra.SetValue[terra.StringValue] `hcl:"gateway_load_balancer_arns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkLoadBalancerArns: set of string, optional
	NetworkLoadBalancerArns terra.SetValue[terra.StringValue] `hcl:"network_load_balancer_arns,attr"`
	// PrivateDnsName: string, optional
	PrivateDnsName terra.StringValue `hcl:"private_dns_name,attr"`
	// SupportedIpAddressTypes: set of string, optional
	SupportedIpAddressTypes terra.SetValue[terra.StringValue] `hcl:"supported_ip_address_types,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// PrivateDnsNameConfiguration: min=0
	PrivateDnsNameConfiguration []vpcendpointservice.PrivateDnsNameConfiguration `hcl:"private_dns_name_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *vpcendpointservice.Timeouts `hcl:"timeouts,block"`
}
type vpcEndpointServiceAttributes struct {
	ref terra.Reference
}

// AcceptanceRequired returns a reference to field acceptance_required of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) AcceptanceRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ves.ref.Append("acceptance_required"))
}

// AllowedPrincipals returns a reference to field allowed_principals of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) AllowedPrincipals() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("allowed_principals"))
}

// Arn returns a reference to field arn of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("availability_zones"))
}

// BaseEndpointDnsNames returns a reference to field base_endpoint_dns_names of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) BaseEndpointDnsNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("base_endpoint_dns_names"))
}

// GatewayLoadBalancerArns returns a reference to field gateway_load_balancer_arns of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) GatewayLoadBalancerArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("gateway_load_balancer_arns"))
}

// Id returns a reference to field id of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("id"))
}

// ManagesVpcEndpoints returns a reference to field manages_vpc_endpoints of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) ManagesVpcEndpoints() terra.BoolValue {
	return terra.ReferenceAsBool(ves.ref.Append("manages_vpc_endpoints"))
}

// NetworkLoadBalancerArns returns a reference to field network_load_balancer_arns of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) NetworkLoadBalancerArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("network_load_balancer_arns"))
}

// PrivateDnsName returns a reference to field private_dns_name of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) PrivateDnsName() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("private_dns_name"))
}

// ServiceName returns a reference to field service_name of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("service_name"))
}

// ServiceType returns a reference to field service_type of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) ServiceType() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("service_type"))
}

// State returns a reference to field state of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("state"))
}

// SupportedIpAddressTypes returns a reference to field supported_ip_address_types of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) SupportedIpAddressTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("supported_ip_address_types"))
}

// Tags returns a reference to field tags of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ves.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_endpoint_service.
func (ves vpcEndpointServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ves.ref.Append("tags_all"))
}

func (ves vpcEndpointServiceAttributes) PrivateDnsNameConfiguration() terra.ListValue[vpcendpointservice.PrivateDnsNameConfigurationAttributes] {
	return terra.ReferenceAsList[vpcendpointservice.PrivateDnsNameConfigurationAttributes](ves.ref.Append("private_dns_name_configuration"))
}

func (ves vpcEndpointServiceAttributes) Timeouts() vpcendpointservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcendpointservice.TimeoutsAttributes](ves.ref.Append("timeouts"))
}

type vpcEndpointServiceState struct {
	AcceptanceRequired          bool                                                  `json:"acceptance_required"`
	AllowedPrincipals           []string                                              `json:"allowed_principals"`
	Arn                         string                                                `json:"arn"`
	AvailabilityZones           []string                                              `json:"availability_zones"`
	BaseEndpointDnsNames        []string                                              `json:"base_endpoint_dns_names"`
	GatewayLoadBalancerArns     []string                                              `json:"gateway_load_balancer_arns"`
	Id                          string                                                `json:"id"`
	ManagesVpcEndpoints         bool                                                  `json:"manages_vpc_endpoints"`
	NetworkLoadBalancerArns     []string                                              `json:"network_load_balancer_arns"`
	PrivateDnsName              string                                                `json:"private_dns_name"`
	ServiceName                 string                                                `json:"service_name"`
	ServiceType                 string                                                `json:"service_type"`
	State                       string                                                `json:"state"`
	SupportedIpAddressTypes     []string                                              `json:"supported_ip_address_types"`
	Tags                        map[string]string                                     `json:"tags"`
	TagsAll                     map[string]string                                     `json:"tags_all"`
	PrivateDnsNameConfiguration []vpcendpointservice.PrivateDnsNameConfigurationState `json:"private_dns_name_configuration"`
	Timeouts                    *vpcendpointservice.TimeoutsState                     `json:"timeouts"`
}
