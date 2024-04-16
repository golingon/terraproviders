// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_mskconnect_worker_configuration

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_mskconnect_worker_configuration.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (amwc *DataSource) DataSource() string {
	return "aws_mskconnect_worker_configuration"
}

// LocalName returns the local name for [DataSource].
func (amwc *DataSource) LocalName() string {
	return amwc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (amwc *DataSource) Configuration() interface{} {
	return amwc.Args
}

// Attributes returns the attributes for [DataSource].
func (amwc *DataSource) Attributes() dataAwsMskconnectWorkerConfigurationAttributes {
	return dataAwsMskconnectWorkerConfigurationAttributes{ref: terra.ReferenceDataSource(amwc)}
}

// DataArgs contains the configurations for aws_mskconnect_worker_configuration.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type dataAwsMskconnectWorkerConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_mskconnect_worker_configuration.
func (amwc dataAwsMskconnectWorkerConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(amwc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_mskconnect_worker_configuration.
func (amwc dataAwsMskconnectWorkerConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amwc.ref.Append("description"))
}

// Id returns a reference to field id of aws_mskconnect_worker_configuration.
func (amwc dataAwsMskconnectWorkerConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amwc.ref.Append("id"))
}

// LatestRevision returns a reference to field latest_revision of aws_mskconnect_worker_configuration.
func (amwc dataAwsMskconnectWorkerConfigurationAttributes) LatestRevision() terra.NumberValue {
	return terra.ReferenceAsNumber(amwc.ref.Append("latest_revision"))
}

// Name returns a reference to field name of aws_mskconnect_worker_configuration.
func (amwc dataAwsMskconnectWorkerConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amwc.ref.Append("name"))
}

// PropertiesFileContent returns a reference to field properties_file_content of aws_mskconnect_worker_configuration.
func (amwc dataAwsMskconnectWorkerConfigurationAttributes) PropertiesFileContent() terra.StringValue {
	return terra.ReferenceAsString(amwc.ref.Append("properties_file_content"))
}
