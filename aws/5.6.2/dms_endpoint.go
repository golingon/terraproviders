// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	dmsendpoint "github.com/golingon/terraproviders/aws/5.6.2/dmsendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDmsEndpoint creates a new instance of [DmsEndpoint].
func NewDmsEndpoint(name string, args DmsEndpointArgs) *DmsEndpoint {
	return &DmsEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DmsEndpoint)(nil)

// DmsEndpoint represents the Terraform resource aws_dms_endpoint.
type DmsEndpoint struct {
	Name      string
	Args      DmsEndpointArgs
	state     *dmsEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DmsEndpoint].
func (de *DmsEndpoint) Type() string {
	return "aws_dms_endpoint"
}

// LocalName returns the local name for [DmsEndpoint].
func (de *DmsEndpoint) LocalName() string {
	return de.Name
}

// Configuration returns the configuration (args) for [DmsEndpoint].
func (de *DmsEndpoint) Configuration() interface{} {
	return de.Args
}

// DependOn is used for other resources to depend on [DmsEndpoint].
func (de *DmsEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(de)
}

// Dependencies returns the list of resources [DmsEndpoint] depends_on.
func (de *DmsEndpoint) Dependencies() terra.Dependencies {
	return de.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DmsEndpoint].
func (de *DmsEndpoint) LifecycleManagement() *terra.Lifecycle {
	return de.Lifecycle
}

// Attributes returns the attributes for [DmsEndpoint].
func (de *DmsEndpoint) Attributes() dmsEndpointAttributes {
	return dmsEndpointAttributes{ref: terra.ReferenceResource(de)}
}

// ImportState imports the given attribute values into [DmsEndpoint]'s state.
func (de *DmsEndpoint) ImportState(av io.Reader) error {
	de.state = &dmsEndpointState{}
	if err := json.NewDecoder(av).Decode(de.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", de.Type(), de.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DmsEndpoint] has state.
func (de *DmsEndpoint) State() (*dmsEndpointState, bool) {
	return de.state, de.state != nil
}

// StateMust returns the state for [DmsEndpoint]. Panics if the state is nil.
func (de *DmsEndpoint) StateMust() *dmsEndpointState {
	if de.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", de.Type(), de.LocalName()))
	}
	return de.state
}

// DmsEndpointArgs contains the configurations for aws_dms_endpoint.
type DmsEndpointArgs struct {
	// CertificateArn: string, optional
	CertificateArn terra.StringValue `hcl:"certificate_arn,attr"`
	// DatabaseName: string, optional
	DatabaseName terra.StringValue `hcl:"database_name,attr"`
	// EndpointId: string, required
	EndpointId terra.StringValue `hcl:"endpoint_id,attr" validate:"required"`
	// EndpointType: string, required
	EndpointType terra.StringValue `hcl:"endpoint_type,attr" validate:"required"`
	// EngineName: string, required
	EngineName terra.StringValue `hcl:"engine_name,attr" validate:"required"`
	// ExtraConnectionAttributes: string, optional
	ExtraConnectionAttributes terra.StringValue `hcl:"extra_connection_attributes,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// SecretsManagerAccessRoleArn: string, optional
	SecretsManagerAccessRoleArn terra.StringValue `hcl:"secrets_manager_access_role_arn,attr"`
	// SecretsManagerArn: string, optional
	SecretsManagerArn terra.StringValue `hcl:"secrets_manager_arn,attr"`
	// ServerName: string, optional
	ServerName terra.StringValue `hcl:"server_name,attr"`
	// ServiceAccessRole: string, optional
	ServiceAccessRole terra.StringValue `hcl:"service_access_role,attr"`
	// SslMode: string, optional
	SslMode terra.StringValue `hcl:"ssl_mode,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Username: string, optional
	Username terra.StringValue `hcl:"username,attr"`
	// ElasticsearchSettings: optional
	ElasticsearchSettings *dmsendpoint.ElasticsearchSettings `hcl:"elasticsearch_settings,block"`
	// KafkaSettings: optional
	KafkaSettings *dmsendpoint.KafkaSettings `hcl:"kafka_settings,block"`
	// KinesisSettings: optional
	KinesisSettings *dmsendpoint.KinesisSettings `hcl:"kinesis_settings,block"`
	// MongodbSettings: optional
	MongodbSettings *dmsendpoint.MongodbSettings `hcl:"mongodb_settings,block"`
	// RedisSettings: optional
	RedisSettings *dmsendpoint.RedisSettings `hcl:"redis_settings,block"`
	// RedshiftSettings: optional
	RedshiftSettings *dmsendpoint.RedshiftSettings `hcl:"redshift_settings,block"`
	// S3Settings: optional
	S3Settings *dmsendpoint.S3Settings `hcl:"s3_settings,block"`
	// Timeouts: optional
	Timeouts *dmsendpoint.Timeouts `hcl:"timeouts,block"`
}
type dmsEndpointAttributes struct {
	ref terra.Reference
}

// CertificateArn returns a reference to field certificate_arn of aws_dms_endpoint.
func (de dmsEndpointAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("certificate_arn"))
}

// DatabaseName returns a reference to field database_name of aws_dms_endpoint.
func (de dmsEndpointAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("database_name"))
}

// EndpointArn returns a reference to field endpoint_arn of aws_dms_endpoint.
func (de dmsEndpointAttributes) EndpointArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("endpoint_arn"))
}

// EndpointId returns a reference to field endpoint_id of aws_dms_endpoint.
func (de dmsEndpointAttributes) EndpointId() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("endpoint_id"))
}

// EndpointType returns a reference to field endpoint_type of aws_dms_endpoint.
func (de dmsEndpointAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("endpoint_type"))
}

// EngineName returns a reference to field engine_name of aws_dms_endpoint.
func (de dmsEndpointAttributes) EngineName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("engine_name"))
}

// ExtraConnectionAttributes returns a reference to field extra_connection_attributes of aws_dms_endpoint.
func (de dmsEndpointAttributes) ExtraConnectionAttributes() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("extra_connection_attributes"))
}

// Id returns a reference to field id of aws_dms_endpoint.
func (de dmsEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_dms_endpoint.
func (de dmsEndpointAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("kms_key_arn"))
}

// Password returns a reference to field password of aws_dms_endpoint.
func (de dmsEndpointAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("password"))
}

// Port returns a reference to field port of aws_dms_endpoint.
func (de dmsEndpointAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(de.ref.Append("port"))
}

// SecretsManagerAccessRoleArn returns a reference to field secrets_manager_access_role_arn of aws_dms_endpoint.
func (de dmsEndpointAttributes) SecretsManagerAccessRoleArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("secrets_manager_access_role_arn"))
}

// SecretsManagerArn returns a reference to field secrets_manager_arn of aws_dms_endpoint.
func (de dmsEndpointAttributes) SecretsManagerArn() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("secrets_manager_arn"))
}

// ServerName returns a reference to field server_name of aws_dms_endpoint.
func (de dmsEndpointAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("server_name"))
}

// ServiceAccessRole returns a reference to field service_access_role of aws_dms_endpoint.
func (de dmsEndpointAttributes) ServiceAccessRole() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("service_access_role"))
}

// SslMode returns a reference to field ssl_mode of aws_dms_endpoint.
func (de dmsEndpointAttributes) SslMode() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("ssl_mode"))
}

// Tags returns a reference to field tags of aws_dms_endpoint.
func (de dmsEndpointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](de.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_endpoint.
func (de dmsEndpointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](de.ref.Append("tags_all"))
}

// Username returns a reference to field username of aws_dms_endpoint.
func (de dmsEndpointAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(de.ref.Append("username"))
}

func (de dmsEndpointAttributes) ElasticsearchSettings() terra.ListValue[dmsendpoint.ElasticsearchSettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.ElasticsearchSettingsAttributes](de.ref.Append("elasticsearch_settings"))
}

func (de dmsEndpointAttributes) KafkaSettings() terra.ListValue[dmsendpoint.KafkaSettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.KafkaSettingsAttributes](de.ref.Append("kafka_settings"))
}

func (de dmsEndpointAttributes) KinesisSettings() terra.ListValue[dmsendpoint.KinesisSettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.KinesisSettingsAttributes](de.ref.Append("kinesis_settings"))
}

func (de dmsEndpointAttributes) MongodbSettings() terra.ListValue[dmsendpoint.MongodbSettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.MongodbSettingsAttributes](de.ref.Append("mongodb_settings"))
}

func (de dmsEndpointAttributes) RedisSettings() terra.ListValue[dmsendpoint.RedisSettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.RedisSettingsAttributes](de.ref.Append("redis_settings"))
}

func (de dmsEndpointAttributes) RedshiftSettings() terra.ListValue[dmsendpoint.RedshiftSettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.RedshiftSettingsAttributes](de.ref.Append("redshift_settings"))
}

func (de dmsEndpointAttributes) S3Settings() terra.ListValue[dmsendpoint.S3SettingsAttributes] {
	return terra.ReferenceAsList[dmsendpoint.S3SettingsAttributes](de.ref.Append("s3_settings"))
}

func (de dmsEndpointAttributes) Timeouts() dmsendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dmsendpoint.TimeoutsAttributes](de.ref.Append("timeouts"))
}

type dmsEndpointState struct {
	CertificateArn              string                                   `json:"certificate_arn"`
	DatabaseName                string                                   `json:"database_name"`
	EndpointArn                 string                                   `json:"endpoint_arn"`
	EndpointId                  string                                   `json:"endpoint_id"`
	EndpointType                string                                   `json:"endpoint_type"`
	EngineName                  string                                   `json:"engine_name"`
	ExtraConnectionAttributes   string                                   `json:"extra_connection_attributes"`
	Id                          string                                   `json:"id"`
	KmsKeyArn                   string                                   `json:"kms_key_arn"`
	Password                    string                                   `json:"password"`
	Port                        float64                                  `json:"port"`
	SecretsManagerAccessRoleArn string                                   `json:"secrets_manager_access_role_arn"`
	SecretsManagerArn           string                                   `json:"secrets_manager_arn"`
	ServerName                  string                                   `json:"server_name"`
	ServiceAccessRole           string                                   `json:"service_access_role"`
	SslMode                     string                                   `json:"ssl_mode"`
	Tags                        map[string]string                        `json:"tags"`
	TagsAll                     map[string]string                        `json:"tags_all"`
	Username                    string                                   `json:"username"`
	ElasticsearchSettings       []dmsendpoint.ElasticsearchSettingsState `json:"elasticsearch_settings"`
	KafkaSettings               []dmsendpoint.KafkaSettingsState         `json:"kafka_settings"`
	KinesisSettings             []dmsendpoint.KinesisSettingsState       `json:"kinesis_settings"`
	MongodbSettings             []dmsendpoint.MongodbSettingsState       `json:"mongodb_settings"`
	RedisSettings               []dmsendpoint.RedisSettingsState         `json:"redis_settings"`
	RedshiftSettings            []dmsendpoint.RedshiftSettingsState      `json:"redshift_settings"`
	S3Settings                  []dmsendpoint.S3SettingsState            `json:"s3_settings"`
	Timeouts                    *dmsendpoint.TimeoutsState               `json:"timeouts"`
}
