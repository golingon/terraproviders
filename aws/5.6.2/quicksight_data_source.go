// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	quicksightdatasource "github.com/golingon/terraproviders/aws/5.6.2/quicksightdatasource"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQuicksightDataSource creates a new instance of [QuicksightDataSource].
func NewQuicksightDataSource(name string, args QuicksightDataSourceArgs) *QuicksightDataSource {
	return &QuicksightDataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightDataSource)(nil)

// QuicksightDataSource represents the Terraform resource aws_quicksight_data_source.
type QuicksightDataSource struct {
	Name      string
	Args      QuicksightDataSourceArgs
	state     *quicksightDataSourceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightDataSource].
func (qds *QuicksightDataSource) Type() string {
	return "aws_quicksight_data_source"
}

// LocalName returns the local name for [QuicksightDataSource].
func (qds *QuicksightDataSource) LocalName() string {
	return qds.Name
}

// Configuration returns the configuration (args) for [QuicksightDataSource].
func (qds *QuicksightDataSource) Configuration() interface{} {
	return qds.Args
}

// DependOn is used for other resources to depend on [QuicksightDataSource].
func (qds *QuicksightDataSource) DependOn() terra.Reference {
	return terra.ReferenceResource(qds)
}

// Dependencies returns the list of resources [QuicksightDataSource] depends_on.
func (qds *QuicksightDataSource) Dependencies() terra.Dependencies {
	return qds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightDataSource].
func (qds *QuicksightDataSource) LifecycleManagement() *terra.Lifecycle {
	return qds.Lifecycle
}

// Attributes returns the attributes for [QuicksightDataSource].
func (qds *QuicksightDataSource) Attributes() quicksightDataSourceAttributes {
	return quicksightDataSourceAttributes{ref: terra.ReferenceResource(qds)}
}

// ImportState imports the given attribute values into [QuicksightDataSource]'s state.
func (qds *QuicksightDataSource) ImportState(av io.Reader) error {
	qds.state = &quicksightDataSourceState{}
	if err := json.NewDecoder(av).Decode(qds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qds.Type(), qds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightDataSource] has state.
func (qds *QuicksightDataSource) State() (*quicksightDataSourceState, bool) {
	return qds.state, qds.state != nil
}

// StateMust returns the state for [QuicksightDataSource]. Panics if the state is nil.
func (qds *QuicksightDataSource) StateMust() *quicksightDataSourceState {
	if qds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qds.Type(), qds.LocalName()))
	}
	return qds.state
}

// QuicksightDataSourceArgs contains the configurations for aws_quicksight_data_source.
type QuicksightDataSourceArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// DataSourceId: string, required
	DataSourceId terra.StringValue `hcl:"data_source_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Credentials: optional
	Credentials *quicksightdatasource.Credentials `hcl:"credentials,block"`
	// Parameters: required
	Parameters *quicksightdatasource.Parameters `hcl:"parameters,block" validate:"required"`
	// Permission: min=0,max=64
	Permission []quicksightdatasource.Permission `hcl:"permission,block" validate:"min=0,max=64"`
	// SslProperties: optional
	SslProperties *quicksightdatasource.SslProperties `hcl:"ssl_properties,block"`
	// VpcConnectionProperties: optional
	VpcConnectionProperties *quicksightdatasource.VpcConnectionProperties `hcl:"vpc_connection_properties,block"`
}
type quicksightDataSourceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("aws_account_id"))
}

// DataSourceId returns a reference to field data_source_id of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("data_source_id"))
}

// Id returns a reference to field id of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("id"))
}

// Name returns a reference to field name of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qds.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qds.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_quicksight_data_source.
func (qds quicksightDataSourceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(qds.ref.Append("type"))
}

func (qds quicksightDataSourceAttributes) Credentials() terra.ListValue[quicksightdatasource.CredentialsAttributes] {
	return terra.ReferenceAsList[quicksightdatasource.CredentialsAttributes](qds.ref.Append("credentials"))
}

func (qds quicksightDataSourceAttributes) Parameters() terra.ListValue[quicksightdatasource.ParametersAttributes] {
	return terra.ReferenceAsList[quicksightdatasource.ParametersAttributes](qds.ref.Append("parameters"))
}

func (qds quicksightDataSourceAttributes) Permission() terra.SetValue[quicksightdatasource.PermissionAttributes] {
	return terra.ReferenceAsSet[quicksightdatasource.PermissionAttributes](qds.ref.Append("permission"))
}

func (qds quicksightDataSourceAttributes) SslProperties() terra.ListValue[quicksightdatasource.SslPropertiesAttributes] {
	return terra.ReferenceAsList[quicksightdatasource.SslPropertiesAttributes](qds.ref.Append("ssl_properties"))
}

func (qds quicksightDataSourceAttributes) VpcConnectionProperties() terra.ListValue[quicksightdatasource.VpcConnectionPropertiesAttributes] {
	return terra.ReferenceAsList[quicksightdatasource.VpcConnectionPropertiesAttributes](qds.ref.Append("vpc_connection_properties"))
}

type quicksightDataSourceState struct {
	Arn                     string                                              `json:"arn"`
	AwsAccountId            string                                              `json:"aws_account_id"`
	DataSourceId            string                                              `json:"data_source_id"`
	Id                      string                                              `json:"id"`
	Name                    string                                              `json:"name"`
	Tags                    map[string]string                                   `json:"tags"`
	TagsAll                 map[string]string                                   `json:"tags_all"`
	Type                    string                                              `json:"type"`
	Credentials             []quicksightdatasource.CredentialsState             `json:"credentials"`
	Parameters              []quicksightdatasource.ParametersState              `json:"parameters"`
	Permission              []quicksightdatasource.PermissionState              `json:"permission"`
	SslProperties           []quicksightdatasource.SslPropertiesState           `json:"ssl_properties"`
	VpcConnectionProperties []quicksightdatasource.VpcConnectionPropertiesState `json:"vpc_connection_properties"`
}
