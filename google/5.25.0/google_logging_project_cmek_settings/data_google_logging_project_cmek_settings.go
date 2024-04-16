// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_logging_project_cmek_settings

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_logging_project_cmek_settings.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (glpcs *DataSource) DataSource() string {
	return "google_logging_project_cmek_settings"
}

// LocalName returns the local name for [DataSource].
func (glpcs *DataSource) LocalName() string {
	return glpcs.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (glpcs *DataSource) Configuration() interface{} {
	return glpcs.Args
}

// Attributes returns the attributes for [DataSource].
func (glpcs *DataSource) Attributes() dataGoogleLoggingProjectCmekSettingsAttributes {
	return dataGoogleLoggingProjectCmekSettingsAttributes{ref: terra.ReferenceDataSource(glpcs)}
}

// DataArgs contains the configurations for google_logging_project_cmek_settings.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}

type dataGoogleLoggingProjectCmekSettingsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_logging_project_cmek_settings.
func (glpcs dataGoogleLoggingProjectCmekSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(glpcs.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_logging_project_cmek_settings.
func (glpcs dataGoogleLoggingProjectCmekSettingsAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(glpcs.ref.Append("kms_key_name"))
}

// KmsKeyVersionName returns a reference to field kms_key_version_name of google_logging_project_cmek_settings.
func (glpcs dataGoogleLoggingProjectCmekSettingsAttributes) KmsKeyVersionName() terra.StringValue {
	return terra.ReferenceAsString(glpcs.ref.Append("kms_key_version_name"))
}

// Name returns a reference to field name of google_logging_project_cmek_settings.
func (glpcs dataGoogleLoggingProjectCmekSettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(glpcs.ref.Append("name"))
}

// Project returns a reference to field project of google_logging_project_cmek_settings.
func (glpcs dataGoogleLoggingProjectCmekSettingsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(glpcs.ref.Append("project"))
}

// ServiceAccountId returns a reference to field service_account_id of google_logging_project_cmek_settings.
func (glpcs dataGoogleLoggingProjectCmekSettingsAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceAsString(glpcs.ref.Append("service_account_id"))
}
