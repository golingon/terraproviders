// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataservicediscoveryservice "github.com/golingon/terraproviders/aws/5.6.2/dataservicediscoveryservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataServiceDiscoveryService creates a new instance of [DataServiceDiscoveryService].
func NewDataServiceDiscoveryService(name string, args DataServiceDiscoveryServiceArgs) *DataServiceDiscoveryService {
	return &DataServiceDiscoveryService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServiceDiscoveryService)(nil)

// DataServiceDiscoveryService represents the Terraform data resource aws_service_discovery_service.
type DataServiceDiscoveryService struct {
	Name string
	Args DataServiceDiscoveryServiceArgs
}

// DataSource returns the Terraform object type for [DataServiceDiscoveryService].
func (sds *DataServiceDiscoveryService) DataSource() string {
	return "aws_service_discovery_service"
}

// LocalName returns the local name for [DataServiceDiscoveryService].
func (sds *DataServiceDiscoveryService) LocalName() string {
	return sds.Name
}

// Configuration returns the configuration (args) for [DataServiceDiscoveryService].
func (sds *DataServiceDiscoveryService) Configuration() interface{} {
	return sds.Args
}

// Attributes returns the attributes for [DataServiceDiscoveryService].
func (sds *DataServiceDiscoveryService) Attributes() dataServiceDiscoveryServiceAttributes {
	return dataServiceDiscoveryServiceAttributes{ref: terra.ReferenceDataResource(sds)}
}

// DataServiceDiscoveryServiceArgs contains the configurations for aws_service_discovery_service.
type DataServiceDiscoveryServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceId: string, required
	NamespaceId terra.StringValue `hcl:"namespace_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DnsConfig: min=0
	DnsConfig []dataservicediscoveryservice.DnsConfig `hcl:"dns_config,block" validate:"min=0"`
	// HealthCheckConfig: min=0
	HealthCheckConfig []dataservicediscoveryservice.HealthCheckConfig `hcl:"health_check_config,block" validate:"min=0"`
	// HealthCheckCustomConfig: min=0
	HealthCheckCustomConfig []dataservicediscoveryservice.HealthCheckCustomConfig `hcl:"health_check_custom_config,block" validate:"min=0"`
}
type dataServiceDiscoveryServiceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("arn"))
}

// Description returns a reference to field description of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("description"))
}

// Id returns a reference to field id of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("id"))
}

// Name returns a reference to field name of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("name"))
}

// NamespaceId returns a reference to field namespace_id of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(sds.ref.Append("namespace_id"))
}

// Tags returns a reference to field tags of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_service_discovery_service.
func (sds dataServiceDiscoveryServiceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sds.ref.Append("tags_all"))
}

func (sds dataServiceDiscoveryServiceAttributes) DnsConfig() terra.ListValue[dataservicediscoveryservice.DnsConfigAttributes] {
	return terra.ReferenceAsList[dataservicediscoveryservice.DnsConfigAttributes](sds.ref.Append("dns_config"))
}

func (sds dataServiceDiscoveryServiceAttributes) HealthCheckConfig() terra.ListValue[dataservicediscoveryservice.HealthCheckConfigAttributes] {
	return terra.ReferenceAsList[dataservicediscoveryservice.HealthCheckConfigAttributes](sds.ref.Append("health_check_config"))
}

func (sds dataServiceDiscoveryServiceAttributes) HealthCheckCustomConfig() terra.ListValue[dataservicediscoveryservice.HealthCheckCustomConfigAttributes] {
	return terra.ReferenceAsList[dataservicediscoveryservice.HealthCheckCustomConfigAttributes](sds.ref.Append("health_check_custom_config"))
}
