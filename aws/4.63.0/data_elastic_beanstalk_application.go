// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataelasticbeanstalkapplication "github.com/golingon/terraproviders/aws/4.63.0/dataelasticbeanstalkapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataElasticBeanstalkApplication creates a new instance of [DataElasticBeanstalkApplication].
func NewDataElasticBeanstalkApplication(name string, args DataElasticBeanstalkApplicationArgs) *DataElasticBeanstalkApplication {
	return &DataElasticBeanstalkApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataElasticBeanstalkApplication)(nil)

// DataElasticBeanstalkApplication represents the Terraform data resource aws_elastic_beanstalk_application.
type DataElasticBeanstalkApplication struct {
	Name string
	Args DataElasticBeanstalkApplicationArgs
}

// DataSource returns the Terraform object type for [DataElasticBeanstalkApplication].
func (eba *DataElasticBeanstalkApplication) DataSource() string {
	return "aws_elastic_beanstalk_application"
}

// LocalName returns the local name for [DataElasticBeanstalkApplication].
func (eba *DataElasticBeanstalkApplication) LocalName() string {
	return eba.Name
}

// Configuration returns the configuration (args) for [DataElasticBeanstalkApplication].
func (eba *DataElasticBeanstalkApplication) Configuration() interface{} {
	return eba.Args
}

// Attributes returns the attributes for [DataElasticBeanstalkApplication].
func (eba *DataElasticBeanstalkApplication) Attributes() dataElasticBeanstalkApplicationAttributes {
	return dataElasticBeanstalkApplicationAttributes{ref: terra.ReferenceDataResource(eba)}
}

// DataElasticBeanstalkApplicationArgs contains the configurations for aws_elastic_beanstalk_application.
type DataElasticBeanstalkApplicationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// AppversionLifecycle: min=0
	AppversionLifecycle []dataelasticbeanstalkapplication.AppversionLifecycle `hcl:"appversion_lifecycle,block" validate:"min=0"`
}
type dataElasticBeanstalkApplicationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elastic_beanstalk_application.
func (eba dataElasticBeanstalkApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("arn"))
}

// Description returns a reference to field description of aws_elastic_beanstalk_application.
func (eba dataElasticBeanstalkApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("description"))
}

// Id returns a reference to field id of aws_elastic_beanstalk_application.
func (eba dataElasticBeanstalkApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("id"))
}

// Name returns a reference to field name of aws_elastic_beanstalk_application.
func (eba dataElasticBeanstalkApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(eba.ref.Append("name"))
}

func (eba dataElasticBeanstalkApplicationAttributes) AppversionLifecycle() terra.ListValue[dataelasticbeanstalkapplication.AppversionLifecycleAttributes] {
	return terra.ReferenceAsList[dataelasticbeanstalkapplication.AppversionLifecycleAttributes](eba.ref.Append("appversion_lifecycle"))
}
