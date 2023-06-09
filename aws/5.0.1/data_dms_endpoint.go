// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadmsendpoint "github.com/golingon/terraproviders/aws/5.0.1/datadmsendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDmsEndpoint creates a new instance of [DataDmsEndpoint].
func NewDataDmsEndpoint(name string, args DataDmsEndpointArgs) *DataDmsEndpoint {
	return &DataDmsEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDmsEndpoint)(nil)

// DataDmsEndpoint represents the Terraform data resource aws_dms_endpoint.
type DataDmsEndpoint struct {
	Name string
	Args DataDmsEndpointArgs
}

// DataSource returns the Terraform object type for [DataDmsEndpoint].
func (de *DataDmsEndpoint) DataSource() string {
	return "aws_dms_endpoint"
}

// LocalName returns the local name for [DataDmsEndpoint].
func (de *DataDmsEndpoint) LocalName() string {
	return de.Name
}

// Configuration returns the configuration (args) for [DataDmsEndpoint].
func (de *DataDmsEndpoint) Configuration() interface{} {
	return de.Args
}

// Attributes returns the attributes for [DataDmsEndpoint].
func (de *DataDmsEndpoint) Attributes() dataDmsEndpointAttributes {
	return dataDmsEndpointAttributes{ref: terra.ReferenceDataResource(de)}
}

// DataDmsEndpointArgs contains the configurations for aws_dms_endpoint.
type DataDmsEndpointArgs struct {
	// EndpointId: string, required
	EndpointId terra.StringValue `hcl:"endpoint_id,attr" validate:"required"`
	// ExtraConnectionAttributes: string, optional
	ExtraConnectionAttributes terra.StringValue `hcl:"extra_connection_attributes,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// KinesisSettings: min=0
	KinesisSettings []datadmsendpoint.KinesisSettings `hcl:"kinesis_settings,block" validate:"min=0"`
	// RedisSettings: min=0
	RedisSettings []datadmsendpoint.RedisSettings `hcl:"redis_settings,block" validate:"min=0"`
	// RedshiftSettings: min=0
	RedshiftSettings []datadmsendpoint.RedshiftSettings `hcl:"redshift_settings,block" validate:"min=0"`
	// S3Settings: min=0
	S3Settings []datadmsendpoint.S3Settings `hcl:"s3_settings,block" validate:"min=0"`
	// ElasticsearchSettings: min=0
	ElasticsearchSettings []datadmsendpoint.ElasticsearchSettings `hcl:"elasticsearch_settings,block" validate:"min=0"`
	// KafkaSettings: min=0
	KafkaSettings []datadmsendpoint.KafkaSettings `hcl:"kafka_settings,block" validate:"min=0"`
	// MongodbSettings: min=0
	MongodbSettings []datadmsendpoint.MongodbSettings `hcl:"mongodb_settings,block" validate:"min=0"`
}
type dataDmsEndpointAttributes struct {
	ref terra.Reference
}

// CertificateArn returns a reference to field certificate_arn of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("certificate_arn"))
}

// DatabaseName returns a reference to field database_name of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("database_name"))
}

// EndpointArn returns a reference to field endpoint_arn of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) EndpointArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("endpoint_arn"))
}

// EndpointId returns a reference to field endpoint_id of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) EndpointId() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("endpoint_id"))
}

// EndpointType returns a reference to field endpoint_type of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("endpoint_type"))
}

// EngineName returns a reference to field engine_name of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) EngineName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("engine_name"))
}

// ExtraConnectionAttributes returns a reference to field extra_connection_attributes of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) ExtraConnectionAttributes() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("extra_connection_attributes"))
}

// Id returns a reference to field id of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("kms_key_arn"))
}

// Password returns a reference to field password of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("password"))
}

// Port returns a reference to field port of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(de.ref.Append("port"))
}

// SecretsManagerAccessRoleArn returns a reference to field secrets_manager_access_role_arn of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) SecretsManagerAccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("secrets_manager_access_role_arn"))
}

// SecretsManagerArn returns a reference to field secrets_manager_arn of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) SecretsManagerArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("secrets_manager_arn"))
}

// ServerName returns a reference to field server_name of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("server_name"))
}

// ServiceAccessRole returns a reference to field service_access_role of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) ServiceAccessRole() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("service_access_role"))
}

// SslMode returns a reference to field ssl_mode of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) SslMode() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("ssl_mode"))
}

// Tags returns a reference to field tags of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](de.ref.Append("tags"))
}

// Username returns a reference to field username of aws_dms_endpoint.
func (de dataDmsEndpointAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("username"))
}

func (de dataDmsEndpointAttributes) KinesisSettings() terra.ListValue[datadmsendpoint.KinesisSettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.KinesisSettingsAttributes](de.ref.Append("kinesis_settings"))
}

func (de dataDmsEndpointAttributes) RedisSettings() terra.ListValue[datadmsendpoint.RedisSettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.RedisSettingsAttributes](de.ref.Append("redis_settings"))
}

func (de dataDmsEndpointAttributes) RedshiftSettings() terra.ListValue[datadmsendpoint.RedshiftSettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.RedshiftSettingsAttributes](de.ref.Append("redshift_settings"))
}

func (de dataDmsEndpointAttributes) S3Settings() terra.ListValue[datadmsendpoint.S3SettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.S3SettingsAttributes](de.ref.Append("s3_settings"))
}

func (de dataDmsEndpointAttributes) ElasticsearchSettings() terra.ListValue[datadmsendpoint.ElasticsearchSettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.ElasticsearchSettingsAttributes](de.ref.Append("elasticsearch_settings"))
}

func (de dataDmsEndpointAttributes) KafkaSettings() terra.ListValue[datadmsendpoint.KafkaSettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.KafkaSettingsAttributes](de.ref.Append("kafka_settings"))
}

func (de dataDmsEndpointAttributes) MongodbSettings() terra.ListValue[datadmsendpoint.MongodbSettingsAttributes] {
	return terra.ReferenceAsList[datadmsendpoint.MongodbSettingsAttributes](de.ref.Append("mongodb_settings"))
}
