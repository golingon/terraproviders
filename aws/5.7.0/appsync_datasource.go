// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appsyncdatasource "github.com/golingon/terraproviders/aws/5.7.0/appsyncdatasource"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppsyncDatasource creates a new instance of [AppsyncDatasource].
func NewAppsyncDatasource(name string, args AppsyncDatasourceArgs) *AppsyncDatasource {
	return &AppsyncDatasource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppsyncDatasource)(nil)

// AppsyncDatasource represents the Terraform resource aws_appsync_datasource.
type AppsyncDatasource struct {
	Name      string
	Args      AppsyncDatasourceArgs
	state     *appsyncDatasourceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppsyncDatasource].
func (ad *AppsyncDatasource) Type() string {
	return "aws_appsync_datasource"
}

// LocalName returns the local name for [AppsyncDatasource].
func (ad *AppsyncDatasource) LocalName() string {
	return ad.Name
}

// Configuration returns the configuration (args) for [AppsyncDatasource].
func (ad *AppsyncDatasource) Configuration() interface{} {
	return ad.Args
}

// DependOn is used for other resources to depend on [AppsyncDatasource].
func (ad *AppsyncDatasource) DependOn() terra.Reference {
	return terra.ReferenceResource(ad)
}

// Dependencies returns the list of resources [AppsyncDatasource] depends_on.
func (ad *AppsyncDatasource) Dependencies() terra.Dependencies {
	return ad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppsyncDatasource].
func (ad *AppsyncDatasource) LifecycleManagement() *terra.Lifecycle {
	return ad.Lifecycle
}

// Attributes returns the attributes for [AppsyncDatasource].
func (ad *AppsyncDatasource) Attributes() appsyncDatasourceAttributes {
	return appsyncDatasourceAttributes{ref: terra.ReferenceResource(ad)}
}

// ImportState imports the given attribute values into [AppsyncDatasource]'s state.
func (ad *AppsyncDatasource) ImportState(av io.Reader) error {
	ad.state = &appsyncDatasourceState{}
	if err := json.NewDecoder(av).Decode(ad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ad.Type(), ad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppsyncDatasource] has state.
func (ad *AppsyncDatasource) State() (*appsyncDatasourceState, bool) {
	return ad.state, ad.state != nil
}

// StateMust returns the state for [AppsyncDatasource]. Panics if the state is nil.
func (ad *AppsyncDatasource) StateMust() *appsyncDatasourceState {
	if ad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ad.Type(), ad.LocalName()))
	}
	return ad.state
}

// AppsyncDatasourceArgs contains the configurations for aws_appsync_datasource.
type AppsyncDatasourceArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceRoleArn: string, optional
	ServiceRoleArn terra.StringValue `hcl:"service_role_arn,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// DynamodbConfig: optional
	DynamodbConfig *appsyncdatasource.DynamodbConfig `hcl:"dynamodb_config,block"`
	// ElasticsearchConfig: optional
	ElasticsearchConfig *appsyncdatasource.ElasticsearchConfig `hcl:"elasticsearch_config,block"`
	// EventBridgeConfig: optional
	EventBridgeConfig *appsyncdatasource.EventBridgeConfig `hcl:"event_bridge_config,block"`
	// HttpConfig: optional
	HttpConfig *appsyncdatasource.HttpConfig `hcl:"http_config,block"`
	// LambdaConfig: optional
	LambdaConfig *appsyncdatasource.LambdaConfig `hcl:"lambda_config,block"`
	// OpensearchserviceConfig: optional
	OpensearchserviceConfig *appsyncdatasource.OpensearchserviceConfig `hcl:"opensearchservice_config,block"`
	// RelationalDatabaseConfig: optional
	RelationalDatabaseConfig *appsyncdatasource.RelationalDatabaseConfig `hcl:"relational_database_config,block"`
}
type appsyncDatasourceAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("api_id"))
}

// Arn returns a reference to field arn of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("arn"))
}

// Description returns a reference to field description of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("description"))
}

// Id returns a reference to field id of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("id"))
}

// Name returns a reference to field name of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("name"))
}

// ServiceRoleArn returns a reference to field service_role_arn of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) ServiceRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("service_role_arn"))
}

// Type returns a reference to field type of aws_appsync_datasource.
func (ad appsyncDatasourceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("type"))
}

func (ad appsyncDatasourceAttributes) DynamodbConfig() terra.ListValue[appsyncdatasource.DynamodbConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.DynamodbConfigAttributes](ad.ref.Append("dynamodb_config"))
}

func (ad appsyncDatasourceAttributes) ElasticsearchConfig() terra.ListValue[appsyncdatasource.ElasticsearchConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.ElasticsearchConfigAttributes](ad.ref.Append("elasticsearch_config"))
}

func (ad appsyncDatasourceAttributes) EventBridgeConfig() terra.ListValue[appsyncdatasource.EventBridgeConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.EventBridgeConfigAttributes](ad.ref.Append("event_bridge_config"))
}

func (ad appsyncDatasourceAttributes) HttpConfig() terra.ListValue[appsyncdatasource.HttpConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.HttpConfigAttributes](ad.ref.Append("http_config"))
}

func (ad appsyncDatasourceAttributes) LambdaConfig() terra.ListValue[appsyncdatasource.LambdaConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.LambdaConfigAttributes](ad.ref.Append("lambda_config"))
}

func (ad appsyncDatasourceAttributes) OpensearchserviceConfig() terra.ListValue[appsyncdatasource.OpensearchserviceConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.OpensearchserviceConfigAttributes](ad.ref.Append("opensearchservice_config"))
}

func (ad appsyncDatasourceAttributes) RelationalDatabaseConfig() terra.ListValue[appsyncdatasource.RelationalDatabaseConfigAttributes] {
	return terra.ReferenceAsList[appsyncdatasource.RelationalDatabaseConfigAttributes](ad.ref.Append("relational_database_config"))
}

type appsyncDatasourceState struct {
	ApiId                    string                                            `json:"api_id"`
	Arn                      string                                            `json:"arn"`
	Description              string                                            `json:"description"`
	Id                       string                                            `json:"id"`
	Name                     string                                            `json:"name"`
	ServiceRoleArn           string                                            `json:"service_role_arn"`
	Type                     string                                            `json:"type"`
	DynamodbConfig           []appsyncdatasource.DynamodbConfigState           `json:"dynamodb_config"`
	ElasticsearchConfig      []appsyncdatasource.ElasticsearchConfigState      `json:"elasticsearch_config"`
	EventBridgeConfig        []appsyncdatasource.EventBridgeConfigState        `json:"event_bridge_config"`
	HttpConfig               []appsyncdatasource.HttpConfigState               `json:"http_config"`
	LambdaConfig             []appsyncdatasource.LambdaConfigState             `json:"lambda_config"`
	OpensearchserviceConfig  []appsyncdatasource.OpensearchserviceConfigState  `json:"opensearchservice_config"`
	RelationalDatabaseConfig []appsyncdatasource.RelationalDatabaseConfigState `json:"relational_database_config"`
}