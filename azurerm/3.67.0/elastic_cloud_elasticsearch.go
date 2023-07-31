// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	elasticcloudelasticsearch "github.com/golingon/terraproviders/azurerm/3.67.0/elasticcloudelasticsearch"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticCloudElasticsearch creates a new instance of [ElasticCloudElasticsearch].
func NewElasticCloudElasticsearch(name string, args ElasticCloudElasticsearchArgs) *ElasticCloudElasticsearch {
	return &ElasticCloudElasticsearch{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticCloudElasticsearch)(nil)

// ElasticCloudElasticsearch represents the Terraform resource azurerm_elastic_cloud_elasticsearch.
type ElasticCloudElasticsearch struct {
	Name      string
	Args      ElasticCloudElasticsearchArgs
	state     *elasticCloudElasticsearchState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticCloudElasticsearch].
func (ece *ElasticCloudElasticsearch) Type() string {
	return "azurerm_elastic_cloud_elasticsearch"
}

// LocalName returns the local name for [ElasticCloudElasticsearch].
func (ece *ElasticCloudElasticsearch) LocalName() string {
	return ece.Name
}

// Configuration returns the configuration (args) for [ElasticCloudElasticsearch].
func (ece *ElasticCloudElasticsearch) Configuration() interface{} {
	return ece.Args
}

// DependOn is used for other resources to depend on [ElasticCloudElasticsearch].
func (ece *ElasticCloudElasticsearch) DependOn() terra.Reference {
	return terra.ReferenceResource(ece)
}

// Dependencies returns the list of resources [ElasticCloudElasticsearch] depends_on.
func (ece *ElasticCloudElasticsearch) Dependencies() terra.Dependencies {
	return ece.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticCloudElasticsearch].
func (ece *ElasticCloudElasticsearch) LifecycleManagement() *terra.Lifecycle {
	return ece.Lifecycle
}

// Attributes returns the attributes for [ElasticCloudElasticsearch].
func (ece *ElasticCloudElasticsearch) Attributes() elasticCloudElasticsearchAttributes {
	return elasticCloudElasticsearchAttributes{ref: terra.ReferenceResource(ece)}
}

// ImportState imports the given attribute values into [ElasticCloudElasticsearch]'s state.
func (ece *ElasticCloudElasticsearch) ImportState(av io.Reader) error {
	ece.state = &elasticCloudElasticsearchState{}
	if err := json.NewDecoder(av).Decode(ece.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ece.Type(), ece.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticCloudElasticsearch] has state.
func (ece *ElasticCloudElasticsearch) State() (*elasticCloudElasticsearchState, bool) {
	return ece.state, ece.state != nil
}

// StateMust returns the state for [ElasticCloudElasticsearch]. Panics if the state is nil.
func (ece *ElasticCloudElasticsearch) StateMust() *elasticCloudElasticsearchState {
	if ece.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ece.Type(), ece.LocalName()))
	}
	return ece.state
}

// ElasticCloudElasticsearchArgs contains the configurations for azurerm_elastic_cloud_elasticsearch.
type ElasticCloudElasticsearchArgs struct {
	// ElasticCloudEmailAddress: string, required
	ElasticCloudEmailAddress terra.StringValue `hcl:"elastic_cloud_email_address,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MonitoringEnabled: bool, optional
	MonitoringEnabled terra.BoolValue `hcl:"monitoring_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Logs: optional
	Logs *elasticcloudelasticsearch.Logs `hcl:"logs,block"`
	// Timeouts: optional
	Timeouts *elasticcloudelasticsearch.Timeouts `hcl:"timeouts,block"`
}
type elasticCloudElasticsearchAttributes struct {
	ref terra.Reference
}

// ElasticCloudDeploymentId returns a reference to field elastic_cloud_deployment_id of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) ElasticCloudDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_deployment_id"))
}

// ElasticCloudEmailAddress returns a reference to field elastic_cloud_email_address of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) ElasticCloudEmailAddress() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_email_address"))
}

// ElasticCloudSsoDefaultUrl returns a reference to field elastic_cloud_sso_default_url of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) ElasticCloudSsoDefaultUrl() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_sso_default_url"))
}

// ElasticCloudUserId returns a reference to field elastic_cloud_user_id of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) ElasticCloudUserId() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elastic_cloud_user_id"))
}

// ElasticsearchServiceUrl returns a reference to field elasticsearch_service_url of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) ElasticsearchServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("elasticsearch_service_url"))
}

// Id returns a reference to field id of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("id"))
}

// KibanaServiceUrl returns a reference to field kibana_service_url of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) KibanaServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("kibana_service_url"))
}

// KibanaSsoUri returns a reference to field kibana_sso_uri of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) KibanaSsoUri() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("kibana_sso_uri"))
}

// Location returns a reference to field location of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("location"))
}

// MonitoringEnabled returns a reference to field monitoring_enabled of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) MonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ece.ref.Append("monitoring_enabled"))
}

// Name returns a reference to field name of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ece.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_elastic_cloud_elasticsearch.
func (ece elasticCloudElasticsearchAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ece.ref.Append("tags"))
}

func (ece elasticCloudElasticsearchAttributes) Logs() terra.ListValue[elasticcloudelasticsearch.LogsAttributes] {
	return terra.ReferenceAsList[elasticcloudelasticsearch.LogsAttributes](ece.ref.Append("logs"))
}

func (ece elasticCloudElasticsearchAttributes) Timeouts() elasticcloudelasticsearch.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticcloudelasticsearch.TimeoutsAttributes](ece.ref.Append("timeouts"))
}

type elasticCloudElasticsearchState struct {
	ElasticCloudDeploymentId  string                                   `json:"elastic_cloud_deployment_id"`
	ElasticCloudEmailAddress  string                                   `json:"elastic_cloud_email_address"`
	ElasticCloudSsoDefaultUrl string                                   `json:"elastic_cloud_sso_default_url"`
	ElasticCloudUserId        string                                   `json:"elastic_cloud_user_id"`
	ElasticsearchServiceUrl   string                                   `json:"elasticsearch_service_url"`
	Id                        string                                   `json:"id"`
	KibanaServiceUrl          string                                   `json:"kibana_service_url"`
	KibanaSsoUri              string                                   `json:"kibana_sso_uri"`
	Location                  string                                   `json:"location"`
	MonitoringEnabled         bool                                     `json:"monitoring_enabled"`
	Name                      string                                   `json:"name"`
	ResourceGroupName         string                                   `json:"resource_group_name"`
	SkuName                   string                                   `json:"sku_name"`
	Tags                      map[string]string                        `json:"tags"`
	Logs                      []elasticcloudelasticsearch.LogsState    `json:"logs"`
	Timeouts                  *elasticcloudelasticsearch.TimeoutsState `json:"timeouts"`
}
