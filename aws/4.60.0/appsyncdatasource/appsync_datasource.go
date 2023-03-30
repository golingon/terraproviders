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

func (dc DynamodbConfigAttributes) InternalRef() terra.Reference {
	return dc.ref
}

func (dc DynamodbConfigAttributes) InternalWithRef(ref terra.Reference) DynamodbConfigAttributes {
	return DynamodbConfigAttributes{ref: ref}
}

func (dc DynamodbConfigAttributes) InternalTokens() hclwrite.Tokens {
	return dc.ref.InternalTokens()
}

func (dc DynamodbConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("region"))
}

func (dc DynamodbConfigAttributes) TableName() terra.StringValue {
	return terra.ReferenceString(dc.ref.Append("table_name"))
}

func (dc DynamodbConfigAttributes) UseCallerCredentials() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("use_caller_credentials"))
}

func (dc DynamodbConfigAttributes) Versioned() terra.BoolValue {
	return terra.ReferenceBool(dc.ref.Append("versioned"))
}

func (dc DynamodbConfigAttributes) DeltaSyncConfig() terra.ListValue[DeltaSyncConfigAttributes] {
	return terra.ReferenceList[DeltaSyncConfigAttributes](dc.ref.Append("delta_sync_config"))
}

type DeltaSyncConfigAttributes struct {
	ref terra.Reference
}

func (dsc DeltaSyncConfigAttributes) InternalRef() terra.Reference {
	return dsc.ref
}

func (dsc DeltaSyncConfigAttributes) InternalWithRef(ref terra.Reference) DeltaSyncConfigAttributes {
	return DeltaSyncConfigAttributes{ref: ref}
}

func (dsc DeltaSyncConfigAttributes) InternalTokens() hclwrite.Tokens {
	return dsc.ref.InternalTokens()
}

func (dsc DeltaSyncConfigAttributes) BaseTableTtl() terra.NumberValue {
	return terra.ReferenceNumber(dsc.ref.Append("base_table_ttl"))
}

func (dsc DeltaSyncConfigAttributes) DeltaSyncTableName() terra.StringValue {
	return terra.ReferenceString(dsc.ref.Append("delta_sync_table_name"))
}

func (dsc DeltaSyncConfigAttributes) DeltaSyncTableTtl() terra.NumberValue {
	return terra.ReferenceNumber(dsc.ref.Append("delta_sync_table_ttl"))
}

type ElasticsearchConfigAttributes struct {
	ref terra.Reference
}

func (ec ElasticsearchConfigAttributes) InternalRef() terra.Reference {
	return ec.ref
}

func (ec ElasticsearchConfigAttributes) InternalWithRef(ref terra.Reference) ElasticsearchConfigAttributes {
	return ElasticsearchConfigAttributes{ref: ref}
}

func (ec ElasticsearchConfigAttributes) InternalTokens() hclwrite.Tokens {
	return ec.ref.InternalTokens()
}

func (ec ElasticsearchConfigAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("endpoint"))
}

func (ec ElasticsearchConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceString(ec.ref.Append("region"))
}

type EventBridgeConfigAttributes struct {
	ref terra.Reference
}

func (ebc EventBridgeConfigAttributes) InternalRef() terra.Reference {
	return ebc.ref
}

func (ebc EventBridgeConfigAttributes) InternalWithRef(ref terra.Reference) EventBridgeConfigAttributes {
	return EventBridgeConfigAttributes{ref: ref}
}

func (ebc EventBridgeConfigAttributes) InternalTokens() hclwrite.Tokens {
	return ebc.ref.InternalTokens()
}

func (ebc EventBridgeConfigAttributes) EventBusArn() terra.StringValue {
	return terra.ReferenceString(ebc.ref.Append("event_bus_arn"))
}

type HttpConfigAttributes struct {
	ref terra.Reference
}

func (hc HttpConfigAttributes) InternalRef() terra.Reference {
	return hc.ref
}

func (hc HttpConfigAttributes) InternalWithRef(ref terra.Reference) HttpConfigAttributes {
	return HttpConfigAttributes{ref: ref}
}

func (hc HttpConfigAttributes) InternalTokens() hclwrite.Tokens {
	return hc.ref.InternalTokens()
}

func (hc HttpConfigAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(hc.ref.Append("endpoint"))
}

func (hc HttpConfigAttributes) AuthorizationConfig() terra.ListValue[AuthorizationConfigAttributes] {
	return terra.ReferenceList[AuthorizationConfigAttributes](hc.ref.Append("authorization_config"))
}

type AuthorizationConfigAttributes struct {
	ref terra.Reference
}

func (ac AuthorizationConfigAttributes) InternalRef() terra.Reference {
	return ac.ref
}

func (ac AuthorizationConfigAttributes) InternalWithRef(ref terra.Reference) AuthorizationConfigAttributes {
	return AuthorizationConfigAttributes{ref: ref}
}

func (ac AuthorizationConfigAttributes) InternalTokens() hclwrite.Tokens {
	return ac.ref.InternalTokens()
}

func (ac AuthorizationConfigAttributes) AuthorizationType() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("authorization_type"))
}

func (ac AuthorizationConfigAttributes) AwsIamConfig() terra.ListValue[AwsIamConfigAttributes] {
	return terra.ReferenceList[AwsIamConfigAttributes](ac.ref.Append("aws_iam_config"))
}

type AwsIamConfigAttributes struct {
	ref terra.Reference
}

func (aic AwsIamConfigAttributes) InternalRef() terra.Reference {
	return aic.ref
}

func (aic AwsIamConfigAttributes) InternalWithRef(ref terra.Reference) AwsIamConfigAttributes {
	return AwsIamConfigAttributes{ref: ref}
}

func (aic AwsIamConfigAttributes) InternalTokens() hclwrite.Tokens {
	return aic.ref.InternalTokens()
}

func (aic AwsIamConfigAttributes) SigningRegion() terra.StringValue {
	return terra.ReferenceString(aic.ref.Append("signing_region"))
}

func (aic AwsIamConfigAttributes) SigningServiceName() terra.StringValue {
	return terra.ReferenceString(aic.ref.Append("signing_service_name"))
}

type LambdaConfigAttributes struct {
	ref terra.Reference
}

func (lc LambdaConfigAttributes) InternalRef() terra.Reference {
	return lc.ref
}

func (lc LambdaConfigAttributes) InternalWithRef(ref terra.Reference) LambdaConfigAttributes {
	return LambdaConfigAttributes{ref: ref}
}

func (lc LambdaConfigAttributes) InternalTokens() hclwrite.Tokens {
	return lc.ref.InternalTokens()
}

func (lc LambdaConfigAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("function_arn"))
}

type RelationalDatabaseConfigAttributes struct {
	ref terra.Reference
}

func (rdc RelationalDatabaseConfigAttributes) InternalRef() terra.Reference {
	return rdc.ref
}

func (rdc RelationalDatabaseConfigAttributes) InternalWithRef(ref terra.Reference) RelationalDatabaseConfigAttributes {
	return RelationalDatabaseConfigAttributes{ref: ref}
}

func (rdc RelationalDatabaseConfigAttributes) InternalTokens() hclwrite.Tokens {
	return rdc.ref.InternalTokens()
}

func (rdc RelationalDatabaseConfigAttributes) SourceType() terra.StringValue {
	return terra.ReferenceString(rdc.ref.Append("source_type"))
}

func (rdc RelationalDatabaseConfigAttributes) HttpEndpointConfig() terra.ListValue[HttpEndpointConfigAttributes] {
	return terra.ReferenceList[HttpEndpointConfigAttributes](rdc.ref.Append("http_endpoint_config"))
}

type HttpEndpointConfigAttributes struct {
	ref terra.Reference
}

func (hec HttpEndpointConfigAttributes) InternalRef() terra.Reference {
	return hec.ref
}

func (hec HttpEndpointConfigAttributes) InternalWithRef(ref terra.Reference) HttpEndpointConfigAttributes {
	return HttpEndpointConfigAttributes{ref: ref}
}

func (hec HttpEndpointConfigAttributes) InternalTokens() hclwrite.Tokens {
	return hec.ref.InternalTokens()
}

func (hec HttpEndpointConfigAttributes) AwsSecretStoreArn() terra.StringValue {
	return terra.ReferenceString(hec.ref.Append("aws_secret_store_arn"))
}

func (hec HttpEndpointConfigAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(hec.ref.Append("database_name"))
}

func (hec HttpEndpointConfigAttributes) DbClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(hec.ref.Append("db_cluster_identifier"))
}

func (hec HttpEndpointConfigAttributes) Region() terra.StringValue {
	return terra.ReferenceString(hec.ref.Append("region"))
}

func (hec HttpEndpointConfigAttributes) Schema() terra.StringValue {
	return terra.ReferenceString(hec.ref.Append("schema"))
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
