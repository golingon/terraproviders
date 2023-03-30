// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	quicksightdatasource "github.com/golingon/terraproviders/aws/4.60.0/quicksightdatasource"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewQuicksightDataSource(name string, args QuicksightDataSourceArgs) *QuicksightDataSource {
	return &QuicksightDataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightDataSource)(nil)

type QuicksightDataSource struct {
	Name  string
	Args  QuicksightDataSourceArgs
	state *quicksightDataSourceState
}

func (qds *QuicksightDataSource) Type() string {
	return "aws_quicksight_data_source"
}

func (qds *QuicksightDataSource) LocalName() string {
	return qds.Name
}

func (qds *QuicksightDataSource) Configuration() interface{} {
	return qds.Args
}

func (qds *QuicksightDataSource) Attributes() quicksightDataSourceAttributes {
	return quicksightDataSourceAttributes{ref: terra.ReferenceResource(qds)}
}

func (qds *QuicksightDataSource) ImportState(av io.Reader) error {
	qds.state = &quicksightDataSourceState{}
	if err := json.NewDecoder(av).Decode(qds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qds.Type(), qds.LocalName(), err)
	}
	return nil
}

func (qds *QuicksightDataSource) State() (*quicksightDataSourceState, bool) {
	return qds.state, qds.state != nil
}

func (qds *QuicksightDataSource) StateMust() *quicksightDataSourceState {
	if qds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qds.Type(), qds.LocalName()))
	}
	return qds.state
}

func (qds *QuicksightDataSource) DependOn() terra.Reference {
	return terra.ReferenceResource(qds)
}

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
	// DependsOn contains resources that QuicksightDataSource depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type quicksightDataSourceAttributes struct {
	ref terra.Reference
}

func (qds quicksightDataSourceAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(qds.ref.Append("arn"))
}

func (qds quicksightDataSourceAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceString(qds.ref.Append("aws_account_id"))
}

func (qds quicksightDataSourceAttributes) DataSourceId() terra.StringValue {
	return terra.ReferenceString(qds.ref.Append("data_source_id"))
}

func (qds quicksightDataSourceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(qds.ref.Append("id"))
}

func (qds quicksightDataSourceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(qds.ref.Append("name"))
}

func (qds quicksightDataSourceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](qds.ref.Append("tags"))
}

func (qds quicksightDataSourceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](qds.ref.Append("tags_all"))
}

func (qds quicksightDataSourceAttributes) Type() terra.StringValue {
	return terra.ReferenceString(qds.ref.Append("type"))
}

func (qds quicksightDataSourceAttributes) Credentials() terra.ListValue[quicksightdatasource.CredentialsAttributes] {
	return terra.ReferenceList[quicksightdatasource.CredentialsAttributes](qds.ref.Append("credentials"))
}

func (qds quicksightDataSourceAttributes) Parameters() terra.ListValue[quicksightdatasource.ParametersAttributes] {
	return terra.ReferenceList[quicksightdatasource.ParametersAttributes](qds.ref.Append("parameters"))
}

func (qds quicksightDataSourceAttributes) Permission() terra.SetValue[quicksightdatasource.PermissionAttributes] {
	return terra.ReferenceSet[quicksightdatasource.PermissionAttributes](qds.ref.Append("permission"))
}

func (qds quicksightDataSourceAttributes) SslProperties() terra.ListValue[quicksightdatasource.SslPropertiesAttributes] {
	return terra.ReferenceList[quicksightdatasource.SslPropertiesAttributes](qds.ref.Append("ssl_properties"))
}

func (qds quicksightDataSourceAttributes) VpcConnectionProperties() terra.ListValue[quicksightdatasource.VpcConnectionPropertiesAttributes] {
	return terra.ReferenceList[quicksightdatasource.VpcConnectionPropertiesAttributes](qds.ref.Append("vpc_connection_properties"))
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
