// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasynclocations3 "github.com/golingon/terraproviders/aws/4.60.0/datasynclocations3"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDatasyncLocationS3(name string, args DatasyncLocationS3Args) *DatasyncLocationS3 {
	return &DatasyncLocationS3{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationS3)(nil)

type DatasyncLocationS3 struct {
	Name  string
	Args  DatasyncLocationS3Args
	state *datasyncLocationS3State
}

func (dls *DatasyncLocationS3) Type() string {
	return "aws_datasync_location_s3"
}

func (dls *DatasyncLocationS3) LocalName() string {
	return dls.Name
}

func (dls *DatasyncLocationS3) Configuration() interface{} {
	return dls.Args
}

func (dls *DatasyncLocationS3) Attributes() datasyncLocationS3Attributes {
	return datasyncLocationS3Attributes{ref: terra.ReferenceResource(dls)}
}

func (dls *DatasyncLocationS3) ImportState(av io.Reader) error {
	dls.state = &datasyncLocationS3State{}
	if err := json.NewDecoder(av).Decode(dls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dls.Type(), dls.LocalName(), err)
	}
	return nil
}

func (dls *DatasyncLocationS3) State() (*datasyncLocationS3State, bool) {
	return dls.state, dls.state != nil
}

func (dls *DatasyncLocationS3) StateMust() *datasyncLocationS3State {
	if dls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dls.Type(), dls.LocalName()))
	}
	return dls.state
}

func (dls *DatasyncLocationS3) DependOn() terra.Reference {
	return terra.ReferenceResource(dls)
}

type DatasyncLocationS3Args struct {
	// AgentArns: set of string, optional
	AgentArns terra.SetValue[terra.StringValue] `hcl:"agent_arns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// S3BucketArn: string, required
	S3BucketArn terra.StringValue `hcl:"s3_bucket_arn,attr" validate:"required"`
	// S3StorageClass: string, optional
	S3StorageClass terra.StringValue `hcl:"s3_storage_class,attr"`
	// Subdirectory: string, required
	Subdirectory terra.StringValue `hcl:"subdirectory,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// S3Config: required
	S3Config *datasynclocations3.S3Config `hcl:"s3_config,block" validate:"required"`
	// DependsOn contains resources that DatasyncLocationS3 depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type datasyncLocationS3Attributes struct {
	ref terra.Reference
}

func (dls datasyncLocationS3Attributes) AgentArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dls.ref.Append("agent_arns"))
}

func (dls datasyncLocationS3Attributes) Arn() terra.StringValue {
	return terra.ReferenceString(dls.ref.Append("arn"))
}

func (dls datasyncLocationS3Attributes) Id() terra.StringValue {
	return terra.ReferenceString(dls.ref.Append("id"))
}

func (dls datasyncLocationS3Attributes) S3BucketArn() terra.StringValue {
	return terra.ReferenceString(dls.ref.Append("s3_bucket_arn"))
}

func (dls datasyncLocationS3Attributes) S3StorageClass() terra.StringValue {
	return terra.ReferenceString(dls.ref.Append("s3_storage_class"))
}

func (dls datasyncLocationS3Attributes) Subdirectory() terra.StringValue {
	return terra.ReferenceString(dls.ref.Append("subdirectory"))
}

func (dls datasyncLocationS3Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dls.ref.Append("tags"))
}

func (dls datasyncLocationS3Attributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dls.ref.Append("tags_all"))
}

func (dls datasyncLocationS3Attributes) Uri() terra.StringValue {
	return terra.ReferenceString(dls.ref.Append("uri"))
}

func (dls datasyncLocationS3Attributes) S3Config() terra.ListValue[datasynclocations3.S3ConfigAttributes] {
	return terra.ReferenceList[datasynclocations3.S3ConfigAttributes](dls.ref.Append("s3_config"))
}

type datasyncLocationS3State struct {
	AgentArns      []string                           `json:"agent_arns"`
	Arn            string                             `json:"arn"`
	Id             string                             `json:"id"`
	S3BucketArn    string                             `json:"s3_bucket_arn"`
	S3StorageClass string                             `json:"s3_storage_class"`
	Subdirectory   string                             `json:"subdirectory"`
	Tags           map[string]string                  `json:"tags"`
	TagsAll        map[string]string                  `json:"tags_all"`
	Uri            string                             `json:"uri"`
	S3Config       []datasynclocations3.S3ConfigState `json:"s3_config"`
}
