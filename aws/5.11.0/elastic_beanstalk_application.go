// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticbeanstalkapplication "github.com/golingon/terraproviders/aws/5.11.0/elasticbeanstalkapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticBeanstalkApplication creates a new instance of [ElasticBeanstalkApplication].
func NewElasticBeanstalkApplication(name string, args ElasticBeanstalkApplicationArgs) *ElasticBeanstalkApplication {
	return &ElasticBeanstalkApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticBeanstalkApplication)(nil)

// ElasticBeanstalkApplication represents the Terraform resource aws_elastic_beanstalk_application.
type ElasticBeanstalkApplication struct {
	Name      string
	Args      ElasticBeanstalkApplicationArgs
	state     *elasticBeanstalkApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticBeanstalkApplication].
func (eba *ElasticBeanstalkApplication) Type() string {
	return "aws_elastic_beanstalk_application"
}

// LocalName returns the local name for [ElasticBeanstalkApplication].
func (eba *ElasticBeanstalkApplication) LocalName() string {
	return eba.Name
}

// Configuration returns the configuration (args) for [ElasticBeanstalkApplication].
func (eba *ElasticBeanstalkApplication) Configuration() interface{} {
	return eba.Args
}

// DependOn is used for other resources to depend on [ElasticBeanstalkApplication].
func (eba *ElasticBeanstalkApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(eba)
}

// Dependencies returns the list of resources [ElasticBeanstalkApplication] depends_on.
func (eba *ElasticBeanstalkApplication) Dependencies() terra.Dependencies {
	return eba.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticBeanstalkApplication].
func (eba *ElasticBeanstalkApplication) LifecycleManagement() *terra.Lifecycle {
	return eba.Lifecycle
}

// Attributes returns the attributes for [ElasticBeanstalkApplication].
func (eba *ElasticBeanstalkApplication) Attributes() elasticBeanstalkApplicationAttributes {
	return elasticBeanstalkApplicationAttributes{ref: terra.ReferenceResource(eba)}
}

// ImportState imports the given attribute values into [ElasticBeanstalkApplication]'s state.
func (eba *ElasticBeanstalkApplication) ImportState(av io.Reader) error {
	eba.state = &elasticBeanstalkApplicationState{}
	if err := json.NewDecoder(av).Decode(eba.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eba.Type(), eba.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticBeanstalkApplication] has state.
func (eba *ElasticBeanstalkApplication) State() (*elasticBeanstalkApplicationState, bool) {
	return eba.state, eba.state != nil
}

// StateMust returns the state for [ElasticBeanstalkApplication]. Panics if the state is nil.
func (eba *ElasticBeanstalkApplication) StateMust() *elasticBeanstalkApplicationState {
	if eba.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eba.Type(), eba.LocalName()))
	}
	return eba.state
}

// ElasticBeanstalkApplicationArgs contains the configurations for aws_elastic_beanstalk_application.
type ElasticBeanstalkApplicationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AppversionLifecycle: optional
	AppversionLifecycle *elasticbeanstalkapplication.AppversionLifecycle `hcl:"appversion_lifecycle,block"`
}
type elasticBeanstalkApplicationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elastic_beanstalk_application.
func (eba elasticBeanstalkApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("arn"))
}

// Description returns a reference to field description of aws_elastic_beanstalk_application.
func (eba elasticBeanstalkApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("description"))
}

// Id returns a reference to field id of aws_elastic_beanstalk_application.
func (eba elasticBeanstalkApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("id"))
}

// Name returns a reference to field name of aws_elastic_beanstalk_application.
func (eba elasticBeanstalkApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_elastic_beanstalk_application.
func (eba elasticBeanstalkApplicationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eba.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elastic_beanstalk_application.
func (eba elasticBeanstalkApplicationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eba.ref.Append("tags_all"))
}

func (eba elasticBeanstalkApplicationAttributes) AppversionLifecycle() terra.ListValue[elasticbeanstalkapplication.AppversionLifecycleAttributes] {
	return terra.ReferenceAsList[elasticbeanstalkapplication.AppversionLifecycleAttributes](eba.ref.Append("appversion_lifecycle"))
}

type elasticBeanstalkApplicationState struct {
	Arn                 string                                                 `json:"arn"`
	Description         string                                                 `json:"description"`
	Id                  string                                                 `json:"id"`
	Name                string                                                 `json:"name"`
	Tags                map[string]string                                      `json:"tags"`
	TagsAll             map[string]string                                      `json:"tags_all"`
	AppversionLifecycle []elasticbeanstalkapplication.AppversionLifecycleState `json:"appversion_lifecycle"`
}