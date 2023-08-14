// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticBeanstalkApplicationVersion creates a new instance of [ElasticBeanstalkApplicationVersion].
func NewElasticBeanstalkApplicationVersion(name string, args ElasticBeanstalkApplicationVersionArgs) *ElasticBeanstalkApplicationVersion {
	return &ElasticBeanstalkApplicationVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticBeanstalkApplicationVersion)(nil)

// ElasticBeanstalkApplicationVersion represents the Terraform resource aws_elastic_beanstalk_application_version.
type ElasticBeanstalkApplicationVersion struct {
	Name      string
	Args      ElasticBeanstalkApplicationVersionArgs
	state     *elasticBeanstalkApplicationVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticBeanstalkApplicationVersion].
func (ebav *ElasticBeanstalkApplicationVersion) Type() string {
	return "aws_elastic_beanstalk_application_version"
}

// LocalName returns the local name for [ElasticBeanstalkApplicationVersion].
func (ebav *ElasticBeanstalkApplicationVersion) LocalName() string {
	return ebav.Name
}

// Configuration returns the configuration (args) for [ElasticBeanstalkApplicationVersion].
func (ebav *ElasticBeanstalkApplicationVersion) Configuration() interface{} {
	return ebav.Args
}

// DependOn is used for other resources to depend on [ElasticBeanstalkApplicationVersion].
func (ebav *ElasticBeanstalkApplicationVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(ebav)
}

// Dependencies returns the list of resources [ElasticBeanstalkApplicationVersion] depends_on.
func (ebav *ElasticBeanstalkApplicationVersion) Dependencies() terra.Dependencies {
	return ebav.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticBeanstalkApplicationVersion].
func (ebav *ElasticBeanstalkApplicationVersion) LifecycleManagement() *terra.Lifecycle {
	return ebav.Lifecycle
}

// Attributes returns the attributes for [ElasticBeanstalkApplicationVersion].
func (ebav *ElasticBeanstalkApplicationVersion) Attributes() elasticBeanstalkApplicationVersionAttributes {
	return elasticBeanstalkApplicationVersionAttributes{ref: terra.ReferenceResource(ebav)}
}

// ImportState imports the given attribute values into [ElasticBeanstalkApplicationVersion]'s state.
func (ebav *ElasticBeanstalkApplicationVersion) ImportState(av io.Reader) error {
	ebav.state = &elasticBeanstalkApplicationVersionState{}
	if err := json.NewDecoder(av).Decode(ebav.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ebav.Type(), ebav.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticBeanstalkApplicationVersion] has state.
func (ebav *ElasticBeanstalkApplicationVersion) State() (*elasticBeanstalkApplicationVersionState, bool) {
	return ebav.state, ebav.state != nil
}

// StateMust returns the state for [ElasticBeanstalkApplicationVersion]. Panics if the state is nil.
func (ebav *ElasticBeanstalkApplicationVersion) StateMust() *elasticBeanstalkApplicationVersionState {
	if ebav.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ebav.Type(), ebav.LocalName()))
	}
	return ebav.state
}

// ElasticBeanstalkApplicationVersionArgs contains the configurations for aws_elastic_beanstalk_application_version.
type ElasticBeanstalkApplicationVersionArgs struct {
	// Application: string, required
	Application terra.StringValue `hcl:"application,attr" validate:"required"`
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type elasticBeanstalkApplicationVersionAttributes struct {
	ref terra.Reference
}

// Application returns a reference to field application of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Application() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("application"))
}

// Arn returns a reference to field arn of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("arn"))
}

// Bucket returns a reference to field bucket of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("bucket"))
}

// Description returns a reference to field description of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("description"))
}

// ForceDelete returns a reference to field force_delete of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(ebav.ref.Append("force_delete"))
}

// Id returns a reference to field id of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("id"))
}

// Key returns a reference to field key of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("key"))
}

// Name returns a reference to field name of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ebav.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ebav.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elastic_beanstalk_application_version.
func (ebav elasticBeanstalkApplicationVersionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ebav.ref.Append("tags_all"))
}

type elasticBeanstalkApplicationVersionState struct {
	Application string            `json:"application"`
	Arn         string            `json:"arn"`
	Bucket      string            `json:"bucket"`
	Description string            `json:"description"`
	ForceDelete bool              `json:"force_delete"`
	Id          string            `json:"id"`
	Key         string            `json:"key"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
}