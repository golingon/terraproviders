// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataSecretsmanagerSecretVersion creates a new instance of [DataSecretsmanagerSecretVersion].
func NewDataSecretsmanagerSecretVersion(name string, args DataSecretsmanagerSecretVersionArgs) *DataSecretsmanagerSecretVersion {
	return &DataSecretsmanagerSecretVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretsmanagerSecretVersion)(nil)

// DataSecretsmanagerSecretVersion represents the Terraform data resource aws_secretsmanager_secret_version.
type DataSecretsmanagerSecretVersion struct {
	Name string
	Args DataSecretsmanagerSecretVersionArgs
}

// DataSource returns the Terraform object type for [DataSecretsmanagerSecretVersion].
func (ssv *DataSecretsmanagerSecretVersion) DataSource() string {
	return "aws_secretsmanager_secret_version"
}

// LocalName returns the local name for [DataSecretsmanagerSecretVersion].
func (ssv *DataSecretsmanagerSecretVersion) LocalName() string {
	return ssv.Name
}

// Configuration returns the configuration (args) for [DataSecretsmanagerSecretVersion].
func (ssv *DataSecretsmanagerSecretVersion) Configuration() interface{} {
	return ssv.Args
}

// Attributes returns the attributes for [DataSecretsmanagerSecretVersion].
func (ssv *DataSecretsmanagerSecretVersion) Attributes() dataSecretsmanagerSecretVersionAttributes {
	return dataSecretsmanagerSecretVersionAttributes{ref: terra.ReferenceDataResource(ssv)}
}

// DataSecretsmanagerSecretVersionArgs contains the configurations for aws_secretsmanager_secret_version.
type DataSecretsmanagerSecretVersionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// VersionId: string, optional
	VersionId terra.StringValue `hcl:"version_id,attr"`
	// VersionStage: string, optional
	VersionStage terra.StringValue `hcl:"version_stage,attr"`
}
type dataSecretsmanagerSecretVersionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("id"))
}

// SecretBinary returns a reference to field secret_binary of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) SecretBinary() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("secret_binary"))
}

// SecretId returns a reference to field secret_id of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("secret_id"))
}

// SecretString returns a reference to field secret_string of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) SecretString() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("secret_string"))
}

// VersionId returns a reference to field version_id of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("version_id"))
}

// VersionStage returns a reference to field version_stage of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) VersionStage() terra.StringValue {
	return terra.ReferenceAsString(ssv.ref.Append("version_stage"))
}

// VersionStages returns a reference to field version_stages of aws_secretsmanager_secret_version.
func (ssv dataSecretsmanagerSecretVersionAttributes) VersionStages() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ssv.ref.Append("version_stages"))
}
