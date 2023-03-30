// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewTimestreamwriteDatabase(name string, args TimestreamwriteDatabaseArgs) *TimestreamwriteDatabase {
	return &TimestreamwriteDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TimestreamwriteDatabase)(nil)

type TimestreamwriteDatabase struct {
	Name  string
	Args  TimestreamwriteDatabaseArgs
	state *timestreamwriteDatabaseState
}

func (td *TimestreamwriteDatabase) Type() string {
	return "aws_timestreamwrite_database"
}

func (td *TimestreamwriteDatabase) LocalName() string {
	return td.Name
}

func (td *TimestreamwriteDatabase) Configuration() interface{} {
	return td.Args
}

func (td *TimestreamwriteDatabase) Attributes() timestreamwriteDatabaseAttributes {
	return timestreamwriteDatabaseAttributes{ref: terra.ReferenceResource(td)}
}

func (td *TimestreamwriteDatabase) ImportState(av io.Reader) error {
	td.state = &timestreamwriteDatabaseState{}
	if err := json.NewDecoder(av).Decode(td.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", td.Type(), td.LocalName(), err)
	}
	return nil
}

func (td *TimestreamwriteDatabase) State() (*timestreamwriteDatabaseState, bool) {
	return td.state, td.state != nil
}

func (td *TimestreamwriteDatabase) StateMust() *timestreamwriteDatabaseState {
	if td.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", td.Type(), td.LocalName()))
	}
	return td.state
}

func (td *TimestreamwriteDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(td)
}

type TimestreamwriteDatabaseArgs struct {
	// DatabaseName: string, required
	DatabaseName terra.StringValue `hcl:"database_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that TimestreamwriteDatabase depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type timestreamwriteDatabaseAttributes struct {
	ref terra.Reference
}

func (td timestreamwriteDatabaseAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(td.ref.Append("arn"))
}

func (td timestreamwriteDatabaseAttributes) DatabaseName() terra.StringValue {
	return terra.ReferenceString(td.ref.Append("database_name"))
}

func (td timestreamwriteDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceString(td.ref.Append("id"))
}

func (td timestreamwriteDatabaseAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(td.ref.Append("kms_key_id"))
}

func (td timestreamwriteDatabaseAttributes) TableCount() terra.NumberValue {
	return terra.ReferenceNumber(td.ref.Append("table_count"))
}

func (td timestreamwriteDatabaseAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](td.ref.Append("tags"))
}

func (td timestreamwriteDatabaseAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](td.ref.Append("tags_all"))
}

type timestreamwriteDatabaseState struct {
	Arn          string            `json:"arn"`
	DatabaseName string            `json:"database_name"`
	Id           string            `json:"id"`
	KmsKeyId     string            `json:"kms_key_id"`
	TableCount   float64           `json:"table_count"`
	Tags         map[string]string `json:"tags"`
	TagsAll      map[string]string `json:"tags_all"`
}
