// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package quicksightdatasource

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Credentials struct {
	// CopySourceArn: string, optional
	CopySourceArn terra.StringValue `hcl:"copy_source_arn,attr"`
	// CredentialPair: optional
	CredentialPair *CredentialPair `hcl:"credential_pair,block"`
}

type CredentialPair struct {
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
}

type Parameters struct {
	// AmazonElasticsearch: optional
	AmazonElasticsearch *AmazonElasticsearch `hcl:"amazon_elasticsearch,block"`
	// Athena: optional
	Athena *Athena `hcl:"athena,block"`
	// Aurora: optional
	Aurora *Aurora `hcl:"aurora,block"`
	// AuroraPostgresql: optional
	AuroraPostgresql *AuroraPostgresql `hcl:"aurora_postgresql,block"`
	// AwsIotAnalytics: optional
	AwsIotAnalytics *AwsIotAnalytics `hcl:"aws_iot_analytics,block"`
	// Jira: optional
	Jira *Jira `hcl:"jira,block"`
	// MariaDb: optional
	MariaDb *MariaDb `hcl:"maria_db,block"`
	// Mysql: optional
	Mysql *Mysql `hcl:"mysql,block"`
	// Oracle: optional
	Oracle *Oracle `hcl:"oracle,block"`
	// Postgresql: optional
	Postgresql *Postgresql `hcl:"postgresql,block"`
	// Presto: optional
	Presto *Presto `hcl:"presto,block"`
	// Rds: optional
	Rds *Rds `hcl:"rds,block"`
	// Redshift: optional
	Redshift *Redshift `hcl:"redshift,block"`
	// S3: optional
	S3 *S3 `hcl:"s3,block"`
	// ServiceNow: optional
	ServiceNow *ServiceNow `hcl:"service_now,block"`
	// Snowflake: optional
	Snowflake *Snowflake `hcl:"snowflake,block"`
	// Spark: optional
	Spark *Spark `hcl:"spark,block"`
	// SqlServer: optional
	SqlServer *SqlServer `hcl:"sql_server,block"`
	// Teradata: optional
	Teradata *Teradata `hcl:"teradata,block"`
	// Twitter: optional
	Twitter *Twitter `hcl:"twitter,block"`
}

type AmazonElasticsearch struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
}

type Athena struct {
	// WorkGroup: string, optional
	WorkGroup terra.StringValue `hcl:"work_group,attr"`
}

type Aurora struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type AuroraPostgresql struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type AwsIotAnalytics struct {
	// DataSetName: string, required
	DataSetName terra.StringValue `hcl:"data_set_name,attr" validate:"required"`
}

type Jira struct {
	// SiteBaseUrl: string, required
	SiteBaseUrl terra.StringValue `hcl:"site_base_url,attr" validate:"required"`
}

type MariaDb struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Mysql struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Oracle struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Postgresql struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Presto struct {
	// Catalog: string, required
	Catalog terra.StringValue `hcl:"catalog,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Rds struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
}

type Redshift struct {
	// ClusterId: string, optional
	ClusterId terra.StringValue `hcl:"cluster_id,attr"`
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, optional
	Host terra.StringValue `hcl:"host,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
}

type S3 struct {
	// ManifestFileLocation: required
	ManifestFileLocation *ManifestFileLocation `hcl:"manifest_file_location,block" validate:"required"`
}

type ManifestFileLocation struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
}

type ServiceNow struct {
	// SiteBaseUrl: string, required
	SiteBaseUrl terra.StringValue `hcl:"site_base_url,attr" validate:"required"`
}

type Snowflake struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Warehouse: string, required
	Warehouse terra.StringValue `hcl:"warehouse,attr" validate:"required"`
}

type Spark struct {
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type SqlServer struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Teradata struct {
	// Database: string, required
	Database terra.StringValue `hcl:"database,attr" validate:"required"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type Twitter struct {
	// MaxRows: number, required
	MaxRows terra.NumberValue `hcl:"max_rows,attr" validate:"required"`
	// Query: string, required
	Query terra.StringValue `hcl:"query,attr" validate:"required"`
}

type Permission struct {
	// Actions: set of string, required
	Actions terra.SetValue[terra.StringValue] `hcl:"actions,attr" validate:"required"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
}

type SslProperties struct {
	// DisableSsl: bool, required
	DisableSsl terra.BoolValue `hcl:"disable_ssl,attr" validate:"required"`
}

type VpcConnectionProperties struct {
	// VpcConnectionArn: string, required
	VpcConnectionArn terra.StringValue `hcl:"vpc_connection_arn,attr" validate:"required"`
}

type CredentialsAttributes struct {
	ref terra.Reference
}

func (c CredentialsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CredentialsAttributes) InternalWithRef(ref terra.Reference) CredentialsAttributes {
	return CredentialsAttributes{ref: ref}
}

func (c CredentialsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CredentialsAttributes) CopySourceArn() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("copy_source_arn"))
}

func (c CredentialsAttributes) CredentialPair() terra.ListValue[CredentialPairAttributes] {
	return terra.ReferenceAsList[CredentialPairAttributes](c.ref.Append("credential_pair"))
}

type CredentialPairAttributes struct {
	ref terra.Reference
}

func (cp CredentialPairAttributes) InternalRef() (terra.Reference, error) {
	return cp.ref, nil
}

func (cp CredentialPairAttributes) InternalWithRef(ref terra.Reference) CredentialPairAttributes {
	return CredentialPairAttributes{ref: ref}
}

func (cp CredentialPairAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cp.ref.InternalTokens()
}

func (cp CredentialPairAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("password"))
}

func (cp CredentialPairAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(cp.ref.Append("username"))
}

type ParametersAttributes struct {
	ref terra.Reference
}

func (p ParametersAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ParametersAttributes) InternalWithRef(ref terra.Reference) ParametersAttributes {
	return ParametersAttributes{ref: ref}
}

func (p ParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ParametersAttributes) AmazonElasticsearch() terra.ListValue[AmazonElasticsearchAttributes] {
	return terra.ReferenceAsList[AmazonElasticsearchAttributes](p.ref.Append("amazon_elasticsearch"))
}

func (p ParametersAttributes) Athena() terra.ListValue[AthenaAttributes] {
	return terra.ReferenceAsList[AthenaAttributes](p.ref.Append("athena"))
}

func (p ParametersAttributes) Aurora() terra.ListValue[AuroraAttributes] {
	return terra.ReferenceAsList[AuroraAttributes](p.ref.Append("aurora"))
}

func (p ParametersAttributes) AuroraPostgresql() terra.ListValue[AuroraPostgresqlAttributes] {
	return terra.ReferenceAsList[AuroraPostgresqlAttributes](p.ref.Append("aurora_postgresql"))
}

func (p ParametersAttributes) AwsIotAnalytics() terra.ListValue[AwsIotAnalyticsAttributes] {
	return terra.ReferenceAsList[AwsIotAnalyticsAttributes](p.ref.Append("aws_iot_analytics"))
}

func (p ParametersAttributes) Jira() terra.ListValue[JiraAttributes] {
	return terra.ReferenceAsList[JiraAttributes](p.ref.Append("jira"))
}

func (p ParametersAttributes) MariaDb() terra.ListValue[MariaDbAttributes] {
	return terra.ReferenceAsList[MariaDbAttributes](p.ref.Append("maria_db"))
}

func (p ParametersAttributes) Mysql() terra.ListValue[MysqlAttributes] {
	return terra.ReferenceAsList[MysqlAttributes](p.ref.Append("mysql"))
}

func (p ParametersAttributes) Oracle() terra.ListValue[OracleAttributes] {
	return terra.ReferenceAsList[OracleAttributes](p.ref.Append("oracle"))
}

func (p ParametersAttributes) Postgresql() terra.ListValue[PostgresqlAttributes] {
	return terra.ReferenceAsList[PostgresqlAttributes](p.ref.Append("postgresql"))
}

func (p ParametersAttributes) Presto() terra.ListValue[PrestoAttributes] {
	return terra.ReferenceAsList[PrestoAttributes](p.ref.Append("presto"))
}

func (p ParametersAttributes) Rds() terra.ListValue[RdsAttributes] {
	return terra.ReferenceAsList[RdsAttributes](p.ref.Append("rds"))
}

func (p ParametersAttributes) Redshift() terra.ListValue[RedshiftAttributes] {
	return terra.ReferenceAsList[RedshiftAttributes](p.ref.Append("redshift"))
}

func (p ParametersAttributes) S3() terra.ListValue[S3Attributes] {
	return terra.ReferenceAsList[S3Attributes](p.ref.Append("s3"))
}

func (p ParametersAttributes) ServiceNow() terra.ListValue[ServiceNowAttributes] {
	return terra.ReferenceAsList[ServiceNowAttributes](p.ref.Append("service_now"))
}

func (p ParametersAttributes) Snowflake() terra.ListValue[SnowflakeAttributes] {
	return terra.ReferenceAsList[SnowflakeAttributes](p.ref.Append("snowflake"))
}

func (p ParametersAttributes) Spark() terra.ListValue[SparkAttributes] {
	return terra.ReferenceAsList[SparkAttributes](p.ref.Append("spark"))
}

func (p ParametersAttributes) SqlServer() terra.ListValue[SqlServerAttributes] {
	return terra.ReferenceAsList[SqlServerAttributes](p.ref.Append("sql_server"))
}

func (p ParametersAttributes) Teradata() terra.ListValue[TeradataAttributes] {
	return terra.ReferenceAsList[TeradataAttributes](p.ref.Append("teradata"))
}

func (p ParametersAttributes) Twitter() terra.ListValue[TwitterAttributes] {
	return terra.ReferenceAsList[TwitterAttributes](p.ref.Append("twitter"))
}

type AmazonElasticsearchAttributes struct {
	ref terra.Reference
}

func (ae AmazonElasticsearchAttributes) InternalRef() (terra.Reference, error) {
	return ae.ref, nil
}

func (ae AmazonElasticsearchAttributes) InternalWithRef(ref terra.Reference) AmazonElasticsearchAttributes {
	return AmazonElasticsearchAttributes{ref: ref}
}

func (ae AmazonElasticsearchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ae.ref.InternalTokens()
}

func (ae AmazonElasticsearchAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("domain"))
}

type AthenaAttributes struct {
	ref terra.Reference
}

func (a AthenaAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AthenaAttributes) InternalWithRef(ref terra.Reference) AthenaAttributes {
	return AthenaAttributes{ref: ref}
}

func (a AthenaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AthenaAttributes) WorkGroup() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("work_group"))
}

type AuroraAttributes struct {
	ref terra.Reference
}

func (a AuroraAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AuroraAttributes) InternalWithRef(ref terra.Reference) AuroraAttributes {
	return AuroraAttributes{ref: ref}
}

func (a AuroraAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AuroraAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("database"))
}

func (a AuroraAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("host"))
}

func (a AuroraAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("port"))
}

type AuroraPostgresqlAttributes struct {
	ref terra.Reference
}

func (ap AuroraPostgresqlAttributes) InternalRef() (terra.Reference, error) {
	return ap.ref, nil
}

func (ap AuroraPostgresqlAttributes) InternalWithRef(ref terra.Reference) AuroraPostgresqlAttributes {
	return AuroraPostgresqlAttributes{ref: ref}
}

func (ap AuroraPostgresqlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ap.ref.InternalTokens()
}

func (ap AuroraPostgresqlAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("database"))
}

func (ap AuroraPostgresqlAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ap.ref.Append("host"))
}

func (ap AuroraPostgresqlAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ap.ref.Append("port"))
}

type AwsIotAnalyticsAttributes struct {
	ref terra.Reference
}

func (aia AwsIotAnalyticsAttributes) InternalRef() (terra.Reference, error) {
	return aia.ref, nil
}

func (aia AwsIotAnalyticsAttributes) InternalWithRef(ref terra.Reference) AwsIotAnalyticsAttributes {
	return AwsIotAnalyticsAttributes{ref: ref}
}

func (aia AwsIotAnalyticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aia.ref.InternalTokens()
}

func (aia AwsIotAnalyticsAttributes) DataSetName() terra.StringValue {
	return terra.ReferenceAsString(aia.ref.Append("data_set_name"))
}

type JiraAttributes struct {
	ref terra.Reference
}

func (j JiraAttributes) InternalRef() (terra.Reference, error) {
	return j.ref, nil
}

func (j JiraAttributes) InternalWithRef(ref terra.Reference) JiraAttributes {
	return JiraAttributes{ref: ref}
}

func (j JiraAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return j.ref.InternalTokens()
}

func (j JiraAttributes) SiteBaseUrl() terra.StringValue {
	return terra.ReferenceAsString(j.ref.Append("site_base_url"))
}

type MariaDbAttributes struct {
	ref terra.Reference
}

func (md MariaDbAttributes) InternalRef() (terra.Reference, error) {
	return md.ref, nil
}

func (md MariaDbAttributes) InternalWithRef(ref terra.Reference) MariaDbAttributes {
	return MariaDbAttributes{ref: ref}
}

func (md MariaDbAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return md.ref.InternalTokens()
}

func (md MariaDbAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("database"))
}

func (md MariaDbAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(md.ref.Append("host"))
}

func (md MariaDbAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(md.ref.Append("port"))
}

type MysqlAttributes struct {
	ref terra.Reference
}

func (m MysqlAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MysqlAttributes) InternalWithRef(ref terra.Reference) MysqlAttributes {
	return MysqlAttributes{ref: ref}
}

func (m MysqlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MysqlAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("database"))
}

func (m MysqlAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("host"))
}

func (m MysqlAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(m.ref.Append("port"))
}

type OracleAttributes struct {
	ref terra.Reference
}

func (o OracleAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OracleAttributes) InternalWithRef(ref terra.Reference) OracleAttributes {
	return OracleAttributes{ref: ref}
}

func (o OracleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OracleAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("database"))
}

func (o OracleAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("host"))
}

func (o OracleAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(o.ref.Append("port"))
}

type PostgresqlAttributes struct {
	ref terra.Reference
}

func (p PostgresqlAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PostgresqlAttributes) InternalWithRef(ref terra.Reference) PostgresqlAttributes {
	return PostgresqlAttributes{ref: ref}
}

func (p PostgresqlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PostgresqlAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("database"))
}

func (p PostgresqlAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("host"))
}

func (p PostgresqlAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("port"))
}

type PrestoAttributes struct {
	ref terra.Reference
}

func (p PrestoAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PrestoAttributes) InternalWithRef(ref terra.Reference) PrestoAttributes {
	return PrestoAttributes{ref: ref}
}

func (p PrestoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PrestoAttributes) Catalog() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("catalog"))
}

func (p PrestoAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("host"))
}

func (p PrestoAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("port"))
}

type RdsAttributes struct {
	ref terra.Reference
}

func (r RdsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RdsAttributes) InternalWithRef(ref terra.Reference) RdsAttributes {
	return RdsAttributes{ref: ref}
}

func (r RdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RdsAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("database"))
}

func (r RdsAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("instance_id"))
}

type RedshiftAttributes struct {
	ref terra.Reference
}

func (r RedshiftAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RedshiftAttributes) InternalWithRef(ref terra.Reference) RedshiftAttributes {
	return RedshiftAttributes{ref: ref}
}

func (r RedshiftAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RedshiftAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("cluster_id"))
}

func (r RedshiftAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("database"))
}

func (r RedshiftAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("host"))
}

func (r RedshiftAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("port"))
}

type S3Attributes struct {
	ref terra.Reference
}

func (s S3Attributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s S3Attributes) InternalWithRef(ref terra.Reference) S3Attributes {
	return S3Attributes{ref: ref}
}

func (s S3Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s S3Attributes) ManifestFileLocation() terra.ListValue[ManifestFileLocationAttributes] {
	return terra.ReferenceAsList[ManifestFileLocationAttributes](s.ref.Append("manifest_file_location"))
}

type ManifestFileLocationAttributes struct {
	ref terra.Reference
}

func (mfl ManifestFileLocationAttributes) InternalRef() (terra.Reference, error) {
	return mfl.ref, nil
}

func (mfl ManifestFileLocationAttributes) InternalWithRef(ref terra.Reference) ManifestFileLocationAttributes {
	return ManifestFileLocationAttributes{ref: ref}
}

func (mfl ManifestFileLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mfl.ref.InternalTokens()
}

func (mfl ManifestFileLocationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(mfl.ref.Append("bucket"))
}

func (mfl ManifestFileLocationAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(mfl.ref.Append("key"))
}

type ServiceNowAttributes struct {
	ref terra.Reference
}

func (sn ServiceNowAttributes) InternalRef() (terra.Reference, error) {
	return sn.ref, nil
}

func (sn ServiceNowAttributes) InternalWithRef(ref terra.Reference) ServiceNowAttributes {
	return ServiceNowAttributes{ref: ref}
}

func (sn ServiceNowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sn.ref.InternalTokens()
}

func (sn ServiceNowAttributes) SiteBaseUrl() terra.StringValue {
	return terra.ReferenceAsString(sn.ref.Append("site_base_url"))
}

type SnowflakeAttributes struct {
	ref terra.Reference
}

func (s SnowflakeAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SnowflakeAttributes) InternalWithRef(ref terra.Reference) SnowflakeAttributes {
	return SnowflakeAttributes{ref: ref}
}

func (s SnowflakeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SnowflakeAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("database"))
}

func (s SnowflakeAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("host"))
}

func (s SnowflakeAttributes) Warehouse() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("warehouse"))
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

func (s SparkAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("host"))
}

func (s SparkAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("port"))
}

type SqlServerAttributes struct {
	ref terra.Reference
}

func (ss SqlServerAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SqlServerAttributes) InternalWithRef(ref terra.Reference) SqlServerAttributes {
	return SqlServerAttributes{ref: ref}
}

func (ss SqlServerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SqlServerAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("database"))
}

func (ss SqlServerAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("host"))
}

func (ss SqlServerAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("port"))
}

type TeradataAttributes struct {
	ref terra.Reference
}

func (t TeradataAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TeradataAttributes) InternalWithRef(ref terra.Reference) TeradataAttributes {
	return TeradataAttributes{ref: ref}
}

func (t TeradataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TeradataAttributes) Database() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("database"))
}

func (t TeradataAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("host"))
}

func (t TeradataAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("port"))
}

type TwitterAttributes struct {
	ref terra.Reference
}

func (t TwitterAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TwitterAttributes) InternalWithRef(ref terra.Reference) TwitterAttributes {
	return TwitterAttributes{ref: ref}
}

func (t TwitterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TwitterAttributes) MaxRows() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("max_rows"))
}

func (t TwitterAttributes) Query() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("query"))
}

type PermissionAttributes struct {
	ref terra.Reference
}

func (p PermissionAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PermissionAttributes) InternalWithRef(ref terra.Reference) PermissionAttributes {
	return PermissionAttributes{ref: ref}
}

func (p PermissionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PermissionAttributes) Actions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](p.ref.Append("actions"))
}

func (p PermissionAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("principal"))
}

type SslPropertiesAttributes struct {
	ref terra.Reference
}

func (sp SslPropertiesAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp SslPropertiesAttributes) InternalWithRef(ref terra.Reference) SslPropertiesAttributes {
	return SslPropertiesAttributes{ref: ref}
}

func (sp SslPropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp SslPropertiesAttributes) DisableSsl() terra.BoolValue {
	return terra.ReferenceAsBool(sp.ref.Append("disable_ssl"))
}

type VpcConnectionPropertiesAttributes struct {
	ref terra.Reference
}

func (vcp VpcConnectionPropertiesAttributes) InternalRef() (terra.Reference, error) {
	return vcp.ref, nil
}

func (vcp VpcConnectionPropertiesAttributes) InternalWithRef(ref terra.Reference) VpcConnectionPropertiesAttributes {
	return VpcConnectionPropertiesAttributes{ref: ref}
}

func (vcp VpcConnectionPropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vcp.ref.InternalTokens()
}

func (vcp VpcConnectionPropertiesAttributes) VpcConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(vcp.ref.Append("vpc_connection_arn"))
}

type CredentialsState struct {
	CopySourceArn  string                `json:"copy_source_arn"`
	CredentialPair []CredentialPairState `json:"credential_pair"`
}

type CredentialPairState struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type ParametersState struct {
	AmazonElasticsearch []AmazonElasticsearchState `json:"amazon_elasticsearch"`
	Athena              []AthenaState              `json:"athena"`
	Aurora              []AuroraState              `json:"aurora"`
	AuroraPostgresql    []AuroraPostgresqlState    `json:"aurora_postgresql"`
	AwsIotAnalytics     []AwsIotAnalyticsState     `json:"aws_iot_analytics"`
	Jira                []JiraState                `json:"jira"`
	MariaDb             []MariaDbState             `json:"maria_db"`
	Mysql               []MysqlState               `json:"mysql"`
	Oracle              []OracleState              `json:"oracle"`
	Postgresql          []PostgresqlState          `json:"postgresql"`
	Presto              []PrestoState              `json:"presto"`
	Rds                 []RdsState                 `json:"rds"`
	Redshift            []RedshiftState            `json:"redshift"`
	S3                  []S3State                  `json:"s3"`
	ServiceNow          []ServiceNowState          `json:"service_now"`
	Snowflake           []SnowflakeState           `json:"snowflake"`
	Spark               []SparkState               `json:"spark"`
	SqlServer           []SqlServerState           `json:"sql_server"`
	Teradata            []TeradataState            `json:"teradata"`
	Twitter             []TwitterState             `json:"twitter"`
}

type AmazonElasticsearchState struct {
	Domain string `json:"domain"`
}

type AthenaState struct {
	WorkGroup string `json:"work_group"`
}

type AuroraState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type AuroraPostgresqlState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type AwsIotAnalyticsState struct {
	DataSetName string `json:"data_set_name"`
}

type JiraState struct {
	SiteBaseUrl string `json:"site_base_url"`
}

type MariaDbState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type MysqlState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type OracleState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type PostgresqlState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type PrestoState struct {
	Catalog string  `json:"catalog"`
	Host    string  `json:"host"`
	Port    float64 `json:"port"`
}

type RdsState struct {
	Database   string `json:"database"`
	InstanceId string `json:"instance_id"`
}

type RedshiftState struct {
	ClusterId string  `json:"cluster_id"`
	Database  string  `json:"database"`
	Host      string  `json:"host"`
	Port      float64 `json:"port"`
}

type S3State struct {
	ManifestFileLocation []ManifestFileLocationState `json:"manifest_file_location"`
}

type ManifestFileLocationState struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
}

type ServiceNowState struct {
	SiteBaseUrl string `json:"site_base_url"`
}

type SnowflakeState struct {
	Database  string `json:"database"`
	Host      string `json:"host"`
	Warehouse string `json:"warehouse"`
}

type SparkState struct {
	Host string  `json:"host"`
	Port float64 `json:"port"`
}

type SqlServerState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type TeradataState struct {
	Database string  `json:"database"`
	Host     string  `json:"host"`
	Port     float64 `json:"port"`
}

type TwitterState struct {
	MaxRows float64 `json:"max_rows"`
	Query   string  `json:"query"`
}

type PermissionState struct {
	Actions   []string `json:"actions"`
	Principal string   `json:"principal"`
}

type SslPropertiesState struct {
	DisableSsl bool `json:"disable_ssl"`
}

type VpcConnectionPropertiesState struct {
	VpcConnectionArn string `json:"vpc_connection_arn"`
}