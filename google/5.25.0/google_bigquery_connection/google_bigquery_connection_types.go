// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_connection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Aws struct {
	// AwsAccessRole: required
	AccessRole *AwsAccessRole `hcl:"access_role,block" validate:"required"`
}

type AwsAccessRole struct {
	// IamRoleId: string, required
	IamRoleId terra.StringValue `hcl:"iam_role_id,attr" validate:"required"`
}

type Azure struct {
	// CustomerTenantId: string, required
	CustomerTenantId terra.StringValue `hcl:"customer_tenant_id,attr" validate:"required"`
	// FederatedApplicationClientId: string, optional
	FederatedApplicationClientId terra.StringValue `hcl:"federated_application_client_id,attr"`
}

type CloudResource struct{}

type CloudSpanner struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// DatabaseRole: string, optional
	DatabaseRole terra.StringValue `hcl:"database_role,attr"`
	// MaxParallelism: number, optional
	MaxParallelism terra.NumberValue `hcl:"max_parallelism,attr"`
	// UseDataBoost: bool, optional
	UseDataBoost terra.BoolValue `hcl:"use_data_boost,attr"`
	// UseParallelism: bool, optional
	UseParallelism terra.BoolValue `hcl:"use_parallelism,attr"`
	// UseServerlessAnalytics: bool, optional
	UseServerlessAnalytics terra.BoolValue `hcl:"use_serverless_analytics,attr"`
}

type CloudSql struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// CloudSqlCredential: required
	Credential *CloudSqlCredential `hcl:"credential,block" validate:"required"`
}

type CloudSqlCredential struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Spark struct {
	// SparkMetastoreServiceConfig: optional
	MetastoreServiceConfig *SparkMetastoreServiceConfig `hcl:"metastore_service_config,block"`
	// SparkSparkHistoryServerConfig: optional
	SparkHistoryServerConfig *SparkSparkHistoryServerConfig `hcl:"spark_history_server_config,block"`
}

type SparkMetastoreServiceConfig struct {
	// MetastoreService: string, optional
	MetastoreService terra.StringValue `hcl:"metastore_service,attr"`
}

type SparkSparkHistoryServerConfig struct {
	// DataprocCluster: string, optional
	DataprocCluster terra.StringValue `hcl:"dataproc_cluster,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AwsAttributes struct {
	ref terra.Reference
}

func (a AwsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AwsAttributes) InternalWithRef(ref terra.Reference) AwsAttributes {
	return AwsAttributes{ref: ref}
}

func (a AwsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AwsAttributes) AccessRole() terra.ListValue[AwsAccessRoleAttributes] {
	return terra.ReferenceAsList[AwsAccessRoleAttributes](a.ref.Append("access_role"))
}

type AwsAccessRoleAttributes struct {
	ref terra.Reference
}

func (ar AwsAccessRoleAttributes) InternalRef() (terra.Reference, error) {
	return ar.ref, nil
}

func (ar AwsAccessRoleAttributes) InternalWithRef(ref terra.Reference) AwsAccessRoleAttributes {
	return AwsAccessRoleAttributes{ref: ref}
}

func (ar AwsAccessRoleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ar.ref.InternalTokens()
}

func (ar AwsAccessRoleAttributes) IamRoleId() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("iam_role_id"))
}

func (ar AwsAccessRoleAttributes) Identity() terra.StringValue {
	return terra.ReferenceAsString(ar.ref.Append("identity"))
}

type AzureAttributes struct {
	ref terra.Reference
}

func (a AzureAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AzureAttributes) InternalWithRef(ref terra.Reference) AzureAttributes {
	return AzureAttributes{ref: ref}
}

func (a AzureAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AzureAttributes) Application() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("application"))
}

func (a AzureAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("client_id"))
}

func (a AzureAttributes) CustomerTenantId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("customer_tenant_id"))
}

func (a AzureAttributes) FederatedApplicationClientId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("federated_application_client_id"))
}

func (a AzureAttributes) Identity() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("identity"))
}

func (a AzureAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("object_id"))
}

func (a AzureAttributes) RedirectUri() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("redirect_uri"))
}

type CloudResourceAttributes struct {
	ref terra.Reference
}

func (cr CloudResourceAttributes) InternalRef() (terra.Reference, error) {
	return cr.ref, nil
}

func (cr CloudResourceAttributes) InternalWithRef(ref terra.Reference) CloudResourceAttributes {
	return CloudResourceAttributes{ref: ref}
}

func (cr CloudResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cr.ref.InternalTokens()
}

func (cr CloudResourceAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("service_account_id"))
}

type CloudSpannerAttributes struct {
	ref terra.Reference
}

func (cs CloudSpannerAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs CloudSpannerAttributes) InternalWithRef(ref terra.Reference) CloudSpannerAttributes {
	return CloudSpannerAttributes{ref: ref}
}

func (cs CloudSpannerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs CloudSpannerAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("database"))
}

func (cs CloudSpannerAttributes) DatabaseRole() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("database_role"))
}

func (cs CloudSpannerAttributes) MaxParallelism() terra.NumberValue {
	return terra.ReferenceAsNumber(cs.ref.Append("max_parallelism"))
}

func (cs CloudSpannerAttributes) UseDataBoost() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("use_data_boost"))
}

func (cs CloudSpannerAttributes) UseParallelism() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("use_parallelism"))
}

func (cs CloudSpannerAttributes) UseServerlessAnalytics() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("use_serverless_analytics"))
}

type CloudSqlAttributes struct {
	ref terra.Reference
}

func (cs CloudSqlAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs CloudSqlAttributes) InternalWithRef(ref terra.Reference) CloudSqlAttributes {
	return CloudSqlAttributes{ref: ref}
}

func (cs CloudSqlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs CloudSqlAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("database"))
}

func (cs CloudSqlAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("instance_id"))
}

func (cs CloudSqlAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("service_account_id"))
}

func (cs CloudSqlAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("type"))
}

func (cs CloudSqlAttributes) Credential() terra.ListValue[CloudSqlCredentialAttributes] {
	return terra.ReferenceAsList[CloudSqlCredentialAttributes](cs.ref.Append("credential"))
}

type CloudSqlCredentialAttributes struct {
	ref terra.Reference
}

func (c CloudSqlCredentialAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CloudSqlCredentialAttributes) InternalWithRef(ref terra.Reference) CloudSqlCredentialAttributes {
	return CloudSqlCredentialAttributes{ref: ref}
}

func (c CloudSqlCredentialAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CloudSqlCredentialAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("password"))
}

func (c CloudSqlCredentialAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("username"))
}

type SparkAttributes struct {
	ref terra.Reference
}

func (s SparkAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SparkAttributes) InternalWithRef(ref terra.Reference) SparkAttributes {
	return SparkAttributes{ref: ref}
}

func (s SparkAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SparkAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("service_account_id"))
}

func (s SparkAttributes) MetastoreServiceConfig() terra.ListValue[SparkMetastoreServiceConfigAttributes] {
	return terra.ReferenceAsList[SparkMetastoreServiceConfigAttributes](s.ref.Append("metastore_service_config"))
}

func (s SparkAttributes) SparkHistoryServerConfig() terra.ListValue[SparkSparkHistoryServerConfigAttributes] {
	return terra.ReferenceAsList[SparkSparkHistoryServerConfigAttributes](s.ref.Append("spark_history_server_config"))
}

type SparkMetastoreServiceConfigAttributes struct {
	ref terra.Reference
}

func (msc SparkMetastoreServiceConfigAttributes) InternalRef() (terra.Reference, error) {
	return msc.ref, nil
}

func (msc SparkMetastoreServiceConfigAttributes) InternalWithRef(ref terra.Reference) SparkMetastoreServiceConfigAttributes {
	return SparkMetastoreServiceConfigAttributes{ref: ref}
}

func (msc SparkMetastoreServiceConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return msc.ref.InternalTokens()
}

func (msc SparkMetastoreServiceConfigAttributes) MetastoreService() terra.StringValue {
	return terra.ReferenceAsString(msc.ref.Append("metastore_service"))
}

type SparkSparkHistoryServerConfigAttributes struct {
	ref terra.Reference
}

func (shsc SparkSparkHistoryServerConfigAttributes) InternalRef() (terra.Reference, error) {
	return shsc.ref, nil
}

func (shsc SparkSparkHistoryServerConfigAttributes) InternalWithRef(ref terra.Reference) SparkSparkHistoryServerConfigAttributes {
	return SparkSparkHistoryServerConfigAttributes{ref: ref}
}

func (shsc SparkSparkHistoryServerConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return shsc.ref.InternalTokens()
}

func (shsc SparkSparkHistoryServerConfigAttributes) DataprocCluster() terra.StringValue {
	return terra.ReferenceAsString(shsc.ref.Append("dataproc_cluster"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AwsState struct {
	AccessRole []AwsAccessRoleState `json:"access_role"`
}

type AwsAccessRoleState struct {
	IamRoleId string `json:"iam_role_id"`
	Identity  string `json:"identity"`
}

type AzureState struct {
	Application                  string `json:"application"`
	ClientId                     string `json:"client_id"`
	CustomerTenantId             string `json:"customer_tenant_id"`
	FederatedApplicationClientId string `json:"federated_application_client_id"`
	Identity                     string `json:"identity"`
	ObjectId                     string `json:"object_id"`
	RedirectUri                  string `json:"redirect_uri"`
}

type CloudResourceState struct {
	ServiceAccountId string `json:"service_account_id"`
}

type CloudSpannerState struct {
	Database               string  `json:"database"`
	DatabaseRole           string  `json:"database_role"`
	MaxParallelism         float64 `json:"max_parallelism"`
	UseDataBoost           bool    `json:"use_data_boost"`
	UseParallelism         bool    `json:"use_parallelism"`
	UseServerlessAnalytics bool    `json:"use_serverless_analytics"`
}

type CloudSqlState struct {
	Database         string                    `json:"database"`
	InstanceId       string                    `json:"instance_id"`
	ServiceAccountId string                    `json:"service_account_id"`
	Type             string                    `json:"type"`
	Credential       []CloudSqlCredentialState `json:"credential"`
}

type CloudSqlCredentialState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type SparkState struct {
	ServiceAccountId         string                               `json:"service_account_id"`
	MetastoreServiceConfig   []SparkMetastoreServiceConfigState   `json:"metastore_service_config"`
	SparkHistoryServerConfig []SparkSparkHistoryServerConfigState `json:"spark_history_server_config"`
}

type SparkMetastoreServiceConfigState struct {
	MetastoreService string `json:"metastore_service"`
}

type SparkSparkHistoryServerConfigState struct {
	DataprocCluster string `json:"dataproc_cluster"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
