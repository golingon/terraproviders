// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataappconfigenvironment "github.com/golingon/terraproviders/aws/4.66.1/dataappconfigenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppconfigEnvironment creates a new instance of [DataAppconfigEnvironment].
func NewDataAppconfigEnvironment(name string, args DataAppconfigEnvironmentArgs) *DataAppconfigEnvironment {
	return &DataAppconfigEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppconfigEnvironment)(nil)

// DataAppconfigEnvironment represents the Terraform data resource aws_appconfig_environment.
type DataAppconfigEnvironment struct {
	Name string
	Args DataAppconfigEnvironmentArgs
}

// DataSource returns the Terraform object type for [DataAppconfigEnvironment].
func (ae *DataAppconfigEnvironment) DataSource() string {
	return "aws_appconfig_environment"
}

// LocalName returns the local name for [DataAppconfigEnvironment].
func (ae *DataAppconfigEnvironment) LocalName() string {
	return ae.Name
}

// Configuration returns the configuration (args) for [DataAppconfigEnvironment].
func (ae *DataAppconfigEnvironment) Configuration() interface{} {
	return ae.Args
}

// Attributes returns the attributes for [DataAppconfigEnvironment].
func (ae *DataAppconfigEnvironment) Attributes() dataAppconfigEnvironmentAttributes {
	return dataAppconfigEnvironmentAttributes{ref: terra.ReferenceDataResource(ae)}
}

// DataAppconfigEnvironmentArgs contains the configurations for aws_appconfig_environment.
type DataAppconfigEnvironmentArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// EnvironmentId: string, required
	EnvironmentId terra.StringValue `hcl:"environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Monitor: min=0
	Monitor []dataappconfigenvironment.Monitor `hcl:"monitor,block" validate:"min=0"`
}
type dataAppconfigEnvironmentAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("application_id"))
}

// Arn returns a reference to field arn of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("arn"))
}

// Description returns a reference to field description of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("description"))
}

// EnvironmentId returns a reference to field environment_id of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) EnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("environment_id"))
}

// Id returns a reference to field id of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("id"))
}

// Name returns a reference to field name of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("name"))
}

// State returns a reference to field state of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_appconfig_environment.
func (ae dataAppconfigEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ae.ref.Append("tags"))
}

func (ae dataAppconfigEnvironmentAttributes) Monitor() terra.SetValue[dataappconfigenvironment.MonitorAttributes] {
	return terra.ReferenceAsSet[dataappconfigenvironment.MonitorAttributes](ae.ref.Append("monitor"))
}
