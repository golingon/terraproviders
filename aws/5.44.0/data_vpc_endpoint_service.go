// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datavpcendpointservice "github.com/golingon/terraproviders/aws/5.44.0/datavpcendpointservice"
)

// NewDataVpcEndpointService creates a new instance of [DataVpcEndpointService].
func NewDataVpcEndpointService(name string, args DataVpcEndpointServiceArgs) *DataVpcEndpointService {
	return &DataVpcEndpointService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcEndpointService)(nil)

// DataVpcEndpointService represents the Terraform data resource aws_vpc_endpoint_service.
type DataVpcEndpointService struct {
	Name string
	Args DataVpcEndpointServiceArgs
}

// DataSource returns the Terraform object type for [DataVpcEndpointService].
func (ves *DataVpcEndpointService) DataSource() string {
	return "aws_vpc_endpoint_service"
}

// LocalName returns the local name for [DataVpcEndpointService].
func (ves *DataVpcEndpointService) LocalName() string {
	return ves.Name
}

// Configuration returns the configuration (args) for [DataVpcEndpointService].
func (ves *DataVpcEndpointService) Configuration() interface{} {
	return ves.Args
}

// Attributes returns the attributes for [DataVpcEndpointService].
func (ves *DataVpcEndpointService) Attributes() dataVpcEndpointServiceAttributes {
	return dataVpcEndpointServiceAttributes{ref: terra.ReferenceDataResource(ves)}
}

// DataVpcEndpointServiceArgs contains the configurations for aws_vpc_endpoint_service.
type DataVpcEndpointServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Service: string, optional
	Service terra.StringValue `hcl:"service,attr"`
	// ServiceName: string, optional
	ServiceName terra.StringValue `hcl:"service_name,attr"`
	// ServiceType: string, optional
	ServiceType terra.StringValue `hcl:"service_type,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpcendpointservice.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcendpointservice.Timeouts `hcl:"timeouts,block"`
}
type dataVpcEndpointServiceAttributes struct {
	ref terra.Reference
}

// AcceptanceRequired returns a reference to field acceptance_required of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) AcceptanceRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ves.ref.Append("acceptance_required"))
}

// Arn returns a reference to field arn of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("availability_zones"))
}

// BaseEndpointDnsNames returns a reference to field base_endpoint_dns_names of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) BaseEndpointDnsNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("base_endpoint_dns_names"))
}

// Id returns a reference to field id of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("id"))
}

// ManagesVpcEndpoints returns a reference to field manages_vpc_endpoints of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) ManagesVpcEndpoints() terra.BoolValue {
	return terra.ReferenceAsBool(ves.ref.Append("manages_vpc_endpoints"))
}

// Owner returns a reference to field owner of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("owner"))
}

// PrivateDnsName returns a reference to field private_dns_name of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) PrivateDnsName() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("private_dns_name"))
}

// Service returns a reference to field service of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("service"))
}

// ServiceId returns a reference to field service_id of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("service_id"))
}

// ServiceName returns a reference to field service_name of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("service_name"))
}

// ServiceType returns a reference to field service_type of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) ServiceType() terra.StringValue {
	return terra.ReferenceAsString(ves.ref.Append("service_type"))
}

// SupportedIpAddressTypes returns a reference to field supported_ip_address_types of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) SupportedIpAddressTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ves.ref.Append("supported_ip_address_types"))
}

// Tags returns a reference to field tags of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ves.ref.Append("tags"))
}

// VpcEndpointPolicySupported returns a reference to field vpc_endpoint_policy_supported of aws_vpc_endpoint_service.
func (ves dataVpcEndpointServiceAttributes) VpcEndpointPolicySupported() terra.BoolValue {
	return terra.ReferenceAsBool(ves.ref.Append("vpc_endpoint_policy_supported"))
}

func (ves dataVpcEndpointServiceAttributes) Filter() terra.SetValue[datavpcendpointservice.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcendpointservice.FilterAttributes](ves.ref.Append("filter"))
}

func (ves dataVpcEndpointServiceAttributes) Timeouts() datavpcendpointservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcendpointservice.TimeoutsAttributes](ves.ref.Append("timeouts"))
}
