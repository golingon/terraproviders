// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataElasticBeanstalkHostedZone creates a new instance of [DataElasticBeanstalkHostedZone].
func NewDataElasticBeanstalkHostedZone(name string, args DataElasticBeanstalkHostedZoneArgs) *DataElasticBeanstalkHostedZone {
	return &DataElasticBeanstalkHostedZone{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataElasticBeanstalkHostedZone)(nil)

// DataElasticBeanstalkHostedZone represents the Terraform data resource aws_elastic_beanstalk_hosted_zone.
type DataElasticBeanstalkHostedZone struct {
	Name string
	Args DataElasticBeanstalkHostedZoneArgs
}

// DataSource returns the Terraform object type for [DataElasticBeanstalkHostedZone].
func (ebhz *DataElasticBeanstalkHostedZone) DataSource() string {
	return "aws_elastic_beanstalk_hosted_zone"
}

// LocalName returns the local name for [DataElasticBeanstalkHostedZone].
func (ebhz *DataElasticBeanstalkHostedZone) LocalName() string {
	return ebhz.Name
}

// Configuration returns the configuration (args) for [DataElasticBeanstalkHostedZone].
func (ebhz *DataElasticBeanstalkHostedZone) Configuration() interface{} {
	return ebhz.Args
}

// Attributes returns the attributes for [DataElasticBeanstalkHostedZone].
func (ebhz *DataElasticBeanstalkHostedZone) Attributes() dataElasticBeanstalkHostedZoneAttributes {
	return dataElasticBeanstalkHostedZoneAttributes{ref: terra.ReferenceDataResource(ebhz)}
}

// DataElasticBeanstalkHostedZoneArgs contains the configurations for aws_elastic_beanstalk_hosted_zone.
type DataElasticBeanstalkHostedZoneArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataElasticBeanstalkHostedZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_elastic_beanstalk_hosted_zone.
func (ebhz dataElasticBeanstalkHostedZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ebhz.ref.Append("id"))
}

// Region returns a reference to field region of aws_elastic_beanstalk_hosted_zone.
func (ebhz dataElasticBeanstalkHostedZoneAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ebhz.ref.Append("region"))
}