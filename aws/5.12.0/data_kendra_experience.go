// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datakendraexperience "github.com/golingon/terraproviders/aws/5.12.0/datakendraexperience"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataKendraExperience creates a new instance of [DataKendraExperience].
func NewDataKendraExperience(name string, args DataKendraExperienceArgs) *DataKendraExperience {
	return &DataKendraExperience{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKendraExperience)(nil)

// DataKendraExperience represents the Terraform data resource aws_kendra_experience.
type DataKendraExperience struct {
	Name string
	Args DataKendraExperienceArgs
}

// DataSource returns the Terraform object type for [DataKendraExperience].
func (ke *DataKendraExperience) DataSource() string {
	return "aws_kendra_experience"
}

// LocalName returns the local name for [DataKendraExperience].
func (ke *DataKendraExperience) LocalName() string {
	return ke.Name
}

// Configuration returns the configuration (args) for [DataKendraExperience].
func (ke *DataKendraExperience) Configuration() interface{} {
	return ke.Args
}

// Attributes returns the attributes for [DataKendraExperience].
func (ke *DataKendraExperience) Attributes() dataKendraExperienceAttributes {
	return dataKendraExperienceAttributes{ref: terra.ReferenceDataResource(ke)}
}

// DataKendraExperienceArgs contains the configurations for aws_kendra_experience.
type DataKendraExperienceArgs struct {
	// ExperienceId: string, required
	ExperienceId terra.StringValue `hcl:"experience_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexId: string, required
	IndexId terra.StringValue `hcl:"index_id,attr" validate:"required"`
	// Configuration: min=0
	Configuration []datakendraexperience.Configuration `hcl:"configuration,block" validate:"min=0"`
	// Endpoints: min=0
	Endpoints []datakendraexperience.Endpoints `hcl:"endpoints,block" validate:"min=0"`
}
type dataKendraExperienceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("arn"))
}

// CreatedAt returns a reference to field created_at of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("description"))
}

// ErrorMessage returns a reference to field error_message of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) ErrorMessage() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("error_message"))
}

// ExperienceId returns a reference to field experience_id of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) ExperienceId() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("experience_id"))
}

// Id returns a reference to field id of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("index_id"))
}

// Name returns a reference to field name of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("status"))
}

// UpdatedAt returns a reference to field updated_at of aws_kendra_experience.
func (ke dataKendraExperienceAttributes) UpdatedAt() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("updated_at"))
}

func (ke dataKendraExperienceAttributes) Configuration() terra.ListValue[datakendraexperience.ConfigurationAttributes] {
	return terra.ReferenceAsList[datakendraexperience.ConfigurationAttributes](ke.ref.Append("configuration"))
}

func (ke dataKendraExperienceAttributes) Endpoints() terra.SetValue[datakendraexperience.EndpointsAttributes] {
	return terra.ReferenceAsSet[datakendraexperience.EndpointsAttributes](ke.ref.Append("endpoints"))
}