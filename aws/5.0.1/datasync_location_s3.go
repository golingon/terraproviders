// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasynclocations3 "github.com/golingon/terraproviders/aws/5.0.1/datasynclocations3"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatasyncLocationS3 creates a new instance of [DatasyncLocationS3].
func NewDatasyncLocationS3(name string, args DatasyncLocationS3Args) *DatasyncLocationS3 {
	return &DatasyncLocationS3{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncLocationS3)(nil)

// DatasyncLocationS3 represents the Terraform resource aws_datasync_location_s3.
type DatasyncLocationS3 struct {
	Name      string
	Args      DatasyncLocationS3Args
	state     *datasyncLocationS3State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatasyncLocationS3].
func (dls *DatasyncLocationS3) Type() string {
	return "aws_datasync_location_s3"
}

// LocalName returns the local name for [DatasyncLocationS3].
func (dls *DatasyncLocationS3) LocalName() string {
	return dls.Name
}

// Configuration returns the configuration (args) for [DatasyncLocationS3].
func (dls *DatasyncLocationS3) Configuration() interface{} {
	return dls.Args
}

// DependOn is used for other resources to depend on [DatasyncLocationS3].
func (dls *DatasyncLocationS3) DependOn() terra.Reference {
	return terra.ReferenceResource(dls)
}

// Dependencies returns the list of resources [DatasyncLocationS3] depends_on.
func (dls *DatasyncLocationS3) Dependencies() terra.Dependencies {
	return dls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatasyncLocationS3].
func (dls *DatasyncLocationS3) LifecycleManagement() *terra.Lifecycle {
	return dls.Lifecycle
}

// Attributes returns the attributes for [DatasyncLocationS3].
func (dls *DatasyncLocationS3) Attributes() datasyncLocationS3Attributes {
	return datasyncLocationS3Attributes{ref: terra.ReferenceResource(dls)}
}

// ImportState imports the given attribute values into [DatasyncLocationS3]'s state.
func (dls *DatasyncLocationS3) ImportState(av io.Reader) error {
	dls.state = &datasyncLocationS3State{}
	if err := json.NewDecoder(av).Decode(dls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dls.Type(), dls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatasyncLocationS3] has state.
func (dls *DatasyncLocationS3) State() (*datasyncLocationS3State, bool) {
	return dls.state, dls.state != nil
}

// StateMust returns the state for [DatasyncLocationS3]. Panics if the state is nil.
func (dls *DatasyncLocationS3) StateMust() *datasyncLocationS3State {
	if dls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dls.Type(), dls.LocalName()))
	}
	return dls.state
}

// DatasyncLocationS3Args contains the configurations for aws_datasync_location_s3.
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
}
type datasyncLocationS3Attributes struct {
	ref terra.Reference
}

// AgentArns returns a reference to field agent_arns of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) AgentArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dls.ref.Append("agent_arns"))
}

// Arn returns a reference to field arn of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("arn"))
}

// Id returns a reference to field id of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("id"))
}

// S3BucketArn returns a reference to field s3_bucket_arn of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) S3BucketArn() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("s3_bucket_arn"))
}

// S3StorageClass returns a reference to field s3_storage_class of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) S3StorageClass() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("s3_storage_class"))
}

// Subdirectory returns a reference to field subdirectory of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) Subdirectory() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("subdirectory"))
}

// Tags returns a reference to field tags of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dls.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dls.ref.Append("tags_all"))
}

// Uri returns a reference to field uri of aws_datasync_location_s3.
func (dls datasyncLocationS3Attributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(dls.ref.Append("uri"))
}

func (dls datasyncLocationS3Attributes) S3Config() terra.ListValue[datasynclocations3.S3ConfigAttributes] {
	return terra.ReferenceAsList[datasynclocations3.S3ConfigAttributes](dls.ref.Append("s3_config"))
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
