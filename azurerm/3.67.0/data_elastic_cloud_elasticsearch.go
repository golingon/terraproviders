// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataelasticcloudelasticsearch "github.com/golingon/terraproviders/azurerm/3.67.0/dataelasticcloudelasticsearch"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataElasticCloudElasticsearch creates a new instance of [DataElasticCloudElasticsearch].
func NewDataElasticCloudElasticsearch(name string, args DataElasticCloudElasticsearchArgs) *DataElasticCloudElasticsearch {
	return &DataElasticCloudElasticsearch{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataElasticCloudElasticsearch)(nil)

// DataElasticCloudElasticsearch represents the Terraform data resource azurerm_elastic_cloud_elasticsearch.
type DataElasticCloudElasticsearch struct {
	Name string
	Args DataElasticCloudElasticsearchArgs
}

// DataSource returns the Terraform object type for [DataElasticCloudElasticsearch].
func (ece *DataElasticCloudElasticsearch) DataSource() string {
	return "azurerm_elastic_cloud_elasticsearch"
}

// LocalName returns the local name for [DataElasticCloudElasticsearch].
func (ece *DataElasticCloudElasticsearch) LocalName() string {
	return ece.Name
}

// Configuration returns the configuration (args) for [DataElasticCloudElasticsearch].
func (ece *DataElasticCloudElasticsearch) Configuration() interface{} {
	return ece.Args
}

// Attributes returns the attributes for [DataElasticCloudElasticsearch].
func (ece *DataElasticCloudElasticsearch) Attributes() dataElasticCloudElasticsearchAttributes {
	return dataElasticCloudElasticsearchAttributes{ref: terra.ReferenceDataResource(ece)}
}

// DataElasticCloudElasticsearchArgs contains the configurations for azurerm_elastic_cloud_elasticsearch.
type DataElasticCloudElasticsearchArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Logs: min=0
	Logs []dataelasticcloudelasticsearch.Logs `hcl:"logs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataelasticcloudelasticsearch.Timeouts `hcl:"timeouts,block"`
}
type dataElasticCloudElasticsearchAttributes struct {
	ref terra.Reference
}

// ElasticCloudDeploymentId returns a reference to field elastic_cloud_deployment_id of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) ElasticCloudDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_deployment_id"))
}

// ElasticCloudEmailAddress returns a reference to field elastic_cloud_email_address of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) ElasticCloudEmailAddress() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_email_address"))
}

// ElasticCloudSsoDefaultUrl returns a reference to field elastic_cloud_sso_default_url of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) ElasticCloudSsoDefaultUrl() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_sso_default_url"))
}

// ElasticCloudUserId returns a reference to field elastic_cloud_user_id of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) ElasticCloudUserId() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_user_id"))
}

// ElasticsearchServiceUrl returns a reference to field elasticsearch_service_url of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) ElasticsearchServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elasticsearch_service_url"))
}

// Id returns a reference to field id of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("id"))
}

// KibanaServiceUrl returns a reference to field kibana_service_url of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) KibanaServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("kibana_service_url"))
}

// KibanaSsoUri returns a reference to field kibana_sso_uri of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) KibanaSsoUri() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("kibana_sso_uri"))
}

// Location returns a reference to field location of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("location"))
}

// MonitoringEnabled returns a reference to field monitoring_enabled of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) MonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ece.ref.Append("monitoring_enabled"))
}

// Name returns a reference to field name of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_elastic_cloud_elasticsearch.
func (ece dataElasticCloudElasticsearchAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ece.ref.Append("tags"))
}

func (ece dataElasticCloudElasticsearchAttributes) Logs() terra.ListValue[dataelasticcloudelasticsearch.LogsAttributes] {
	return terra.ReferenceAsList[dataelasticcloudelasticsearch.LogsAttributes](ece.ref.Append("logs"))
}

func (ece dataElasticCloudElasticsearchAttributes) Timeouts() dataelasticcloudelasticsearch.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataelasticcloudelasticsearch.TimeoutsAttributes](ece.ref.Append("timeouts"))
}
