// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package appsyncdatasource

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DynamodbConfig struct {
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
	// UseCallerCredentials: bool, optional
	UseCallerCredentials terra.BoolValue `hcl:"use_caller_credentials,attr"`
	// Versioned: bool, optional
	Versioned terra.BoolValue `hcl:"versioned,attr"`
	// DeltaSyncConfig: optional
	DeltaSyncConfig *DeltaSyncConfig `hcl:"delta_sync_config,block"`
}

type DeltaSyncConfig struct {
	// BaseTableTtl: number, optional
	BaseTableTtl terra.NumberValue `hcl:"base_table_ttl,attr"`
	// DeltaSyncTableName: string, required
	DeltaSyncTableName terra.StringValue `hcl:"delta_sync_table_name,attr" validate:"required"`
	// DeltaSyncTableTtl: number, optional
	DeltaSyncTableTtl terra.NumberValue `hcl:"delta_sync_table_ttl,attr"`
}

type ElasticsearchConfig struct {
	// Endpoint: string, required
	Endpoint terra.StringValue `hcl:"endpoint,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type EventBridgeConfig struct {
	// EventBusArn: string, required
	EventBusArn terra.StringValue `hcl:"event_bus_arn,attr" validate:"required"`
}

type HttpConfig struct {
	// Endpoint: string, required
	Endpoint terra.StringValue `hcl:"endpoint,attr" validate:"required"`
	// AuthorizationConfig: optional
	AuthorizationConfig *AuthorizationConfig `hcl:"authorization_config,block"`
}

type AuthorizationConfig struct {
	// AuthorizationType: string, optional
	AuthorizationType terra.StringValue `hcl:"authorization_type,attr"`
	// AwsIamConfig: optional
	AwsIamConfig *AwsIamConfig `hcl:"aws_iam_config,block"`
}

type AwsIamConfig struct {
	// SigningRegion: string, optional
	SigningRegion terra.StringValue `hcl:"signing_region,attr"`
	// SigningServiceName: string, optional
	SigningServiceName terra.StringValue `hcl:"signing_service_name,attr"`
}

type LambdaConfig struct {
	// FunctionArn: string, required
	FunctionArn terra.StringValue `hcl:"function_arn,attr" validate:"required"`
}

type RelationalDatabaseConfig struct {
	// SourceType: string, optional
	SourceType terra.StringValue `hcl:"source_type,attr"`
	// HttpEndpointConfig: optional
	HttpEndpointConfig *HttpEndpointConfig `hcl:"http_endpoint_config,block"`
}

type HttpEndpointConfig struct {
	// AwsSecretStoreArn: string, required
	AwsSecretStoreArn terra.StringValue `hcl:"aws_secret_store_arn,attr" validate:"required"`
	// DatabaseName: string, optional
	DatabaseName terra.StringValue `hcl:"database_name,attr"`
	// DbClusterIdentifier: string, required
	DbClusterIdentifier terra.StringValue `hcl:"db_cluster_identifier,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Schema: string, optional
	Schema terra.StringValue `hcl:"schema,attr"`
}

type DynamodbConfigAttributes struct {
	ref terra.Reference
}

func (dc DynamodbConfigAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DynamodbConfigAttributes) InternalWithRef(ref terra.Reference) DynamodbConfigAttributes {
	return DynamodbConfigAttributes{ref: ref}
}

func (dc DynamodbConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DynamodbConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("region"))
}

func (dc DynamodbConfigAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("table_name"))
}

func (dc DynamodbConfigAttributes) UseCallerCredentials() terra.BoolValue {
	return terra.ReferenceAsBool(dc.ref.Append("use_caller_credentials"))
}

func (dc DynamodbConfigAttributes) Versioned() terra.BoolValue {
	return terra.ReferenceAsBool(dc.ref.Append("versioned"))
}

func (dc DynamodbConfigAttributes) DeltaSyncConfig() terra.ListValue[DeltaSyncConfigAttributes] {
	return terra.ReferenceAsList[DeltaSyncConfigAttributes](dc.ref.Append("delta_sync_config"))
}

type DeltaSyncConfigAttributes struct {
	ref terra.Reference
}

func (dsc DeltaSyncConfigAttributes) InternalRef() (terra.Reference, error) {
	return dsc.ref, nil
}

func (dsc DeltaSyncConfigAttributes) InternalWithRef(ref terra.Reference) DeltaSyncConfigAttributes {
	return DeltaSyncConfigAttributes{ref: ref}
}

func (dsc DeltaSyncConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dsc.ref.InternalTokens()
}

func (dsc DeltaSyncConfigAttributes) BaseTableTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(dsc.ref.Append("base_table_ttl"))
}

func (dsc DeltaSyncConfigAttributes) DeltaSyncTableName() terra.StringValue {
	return terra.ReferenceAsString(dsc.ref.Append("delta_sync_table_name"))
}

func (dsc DeltaSyncConfigAttributes) DeltaSyncTableTtl() terra.NumberValue {
	return terra.ReferenceAsNumber(dsc.ref.Append("delta_sync_table_ttl"))
}

type ElasticsearchConfigAttributes struct {
	ref terra.Reference
}

func (ec ElasticsearchConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec ElasticsearchConfigAttributes) InternalWithRef(ref terra.Reference) ElasticsearchConfigAttributes {
	return ElasticsearchConfigAttributes{ref: ref}
}

func (ec ElasticsearchConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec ElasticsearchConfigAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("endpoint"))
}

func (ec ElasticsearchConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("region"))
}

type EventBridgeConfigAttributes struct {
	ref terra.Reference
}

func (ebc EventBridgeConfigAttributes) InternalRef() (terra.Reference, error) {
	return ebc.ref, nil
}

func (ebc EventBridgeConfigAttributes) InternalWithRef(ref terra.Reference) EventBridgeConfigAttributes {
	return EventBridgeConfigAttributes{ref: ref}
}

func (ebc EventBridgeConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ebc.ref.InternalTokens()
}

func (ebc EventBridgeConfigAttributes) EventBusArn() terra.StringValue {
	return terra.ReferenceAsString(ebc.ref.Append("event_bus_arn"))
}

type HttpConfigAttributes struct {
	ref terra.Reference
}

func (hc HttpConfigAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HttpConfigAttributes) InternalWithRef(ref terra.Reference) HttpConfigAttributes {
	return HttpConfigAttributes{ref: ref}
}

func (hc HttpConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HttpConfigAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("endpoint"))
}

func (hc HttpConfigAttributes) AuthorizationConfig() terra.ListValue[AuthorizationConfigAttributes] {
	return terra.ReferenceAsList[AuthorizationConfigAttributes](hc.ref.Append("authorization_config"))
}

type AuthorizationConfigAttributes struct {
	ref terra.Reference
}

func (ac AuthorizationConfigAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AuthorizationConfigAttributes) InternalWithRef(ref terra.Reference) AuthorizationConfigAttributes {
	return AuthorizationConfigAttributes{ref: ref}
}

func (ac AuthorizationConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AuthorizationConfigAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("authorization_type"))
}

func (ac AuthorizationConfigAttributes) AwsIamConfig() terra.ListValue[AwsIamConfigAttributes] {
	return terra.ReferenceAsList[AwsIamConfigAttributes](ac.ref.Append("aws_iam_config"))
}

type AwsIamConfigAttributes struct {
	ref terra.Reference
}

func (aic AwsIamConfigAttributes) InternalRef() (terra.Reference, error) {
	return aic.ref, nil
}

func (aic AwsIamConfigAttributes) InternalWithRef(ref terra.Reference) AwsIamConfigAttributes {
	return AwsIamConfigAttributes{ref: ref}
}

func (aic AwsIamConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aic.ref.InternalTokens()
}

func (aic AwsIamConfigAttributes) SigningRegion() terra.StringValue {
	return terra.ReferenceAsString(aic.ref.Append("signing_region"))
}

func (aic AwsIamConfigAttributes) SigningServiceName() terra.StringValue {
	return terra.ReferenceAsString(aic.ref.Append("signing_service_name"))
}

type LambdaConfigAttributes struct {
	ref terra.Reference
}

func (lc LambdaConfigAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc LambdaConfigAttributes) InternalWithRef(ref terra.Reference) LambdaConfigAttributes {
	return LambdaConfigAttributes{ref: ref}
}

func (lc LambdaConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc LambdaConfigAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("function_arn"))
}

type RelationalDatabaseConfigAttributes struct {
	ref terra.Reference
}

func (rdc RelationalDatabaseConfigAttributes) InternalRef() (terra.Reference, error) {
	return rdc.ref, nil
}

func (rdc RelationalDatabaseConfigAttributes) InternalWithRef(ref terra.Reference) RelationalDatabaseConfigAttributes {
	return RelationalDatabaseConfigAttributes{ref: ref}
}

func (rdc RelationalDatabaseConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rdc.ref.InternalTokens()
}

func (rdc RelationalDatabaseConfigAttributes) SourceType() terra.StringValue {
	return terra.ReferenceAsString(rdc.ref.Append("source_type"))
}

func (rdc RelationalDatabaseConfigAttributes) HttpEndpointConfig() terra.ListValue[HttpEndpointConfigAttributes] {
	return terra.ReferenceAsList[HttpEndpointConfigAttributes](rdc.ref.Append("http_endpoint_config"))
}

type HttpEndpointConfigAttributes struct {
	ref terra.Reference
}

func (hec HttpEndpointConfigAttributes) InternalRef() (terra.Reference, error) {
	return hec.ref, nil
}

func (hec HttpEndpointConfigAttributes) InternalWithRef(ref terra.Reference) HttpEndpointConfigAttributes {
	return HttpEndpointConfigAttributes{ref: ref}
}

func (hec HttpEndpointConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hec.ref.InternalTokens()
}

func (hec HttpEndpointConfigAttributes) AwsSecretStoreArn() terra.StringValue {
	return terra.ReferenceAsString(hec.ref.Append("aws_secret_store_arn"))
}

func (hec HttpEndpointConfigAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(hec.ref.Append("database_name"))
}

func (hec HttpEndpointConfigAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(hec.ref.Append("db_cluster_identifier"))
}

func (hec HttpEndpointConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(hec.ref.Append("region"))
}

func (hec HttpEndpointConfigAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(hec.ref.Append("schema"))
}

type DynamodbConfigState struct {
	Region               string                 `json:"region"`
	TableName            string                 `json:"table_name"`
	UseCallerCredentials bool                   `json:"use_caller_credentials"`
	Versioned            bool                   `json:"versioned"`
	DeltaSyncConfig      []DeltaSyncConfigState `json:"delta_sync_config"`
}

type DeltaSyncConfigState struct {
	BaseTableTtl       float64 `json:"base_table_ttl"`
	DeltaSyncTableName string  `json:"delta_sync_table_name"`
	DeltaSyncTableTtl  float64 `json:"delta_sync_table_ttl"`
}

type ElasticsearchConfigState struct {
	Endpoint string `json:"endpoint"`
	Region   string `json:"region"`
}

type EventBridgeConfigState struct {
	EventBusArn string `json:"event_bus_arn"`
}

type HttpConfigState struct {
	Endpoint            string                     `json:"endpoint"`
	AuthorizationConfig []AuthorizationConfigState `json:"authorization_config"`
}

type AuthorizationConfigState struct {
	AuthorizationType string              `json:"authorization_type"`
	AwsIamConfig      []AwsIamConfigState `json:"aws_iam_config"`
}

type AwsIamConfigState struct {
	SigningRegion      string `json:"signing_region"`
	SigningServiceName string `json:"signing_service_name"`
}

type LambdaConfigState struct {
	FunctionArn string `json:"function_arn"`
}

type RelationalDatabaseConfigState struct {
	SourceType         string                    `json:"source_type"`
	HttpEndpointConfig []HttpEndpointConfigState `json:"http_endpoint_config"`
}

type HttpEndpointConfigState struct {
	AwsSecretStoreArn   string `json:"aws_secret_store_arn"`
	DatabaseName        string `json:"database_name"`
	DbClusterIdentifier string `json:"db_cluster_identifier"`
	Region              string `json:"region"`
	Schema              string `json:"schema"`
}
