// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_redshiftserverless_namespace

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_redshiftserverless_namespace.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (arn *DataSource) DataSource() string {
	return "aws_redshiftserverless_namespace"
}

// LocalName returns the local name for [DataSource].
func (arn *DataSource) LocalName() string {
	return arn.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (arn *DataSource) Configuration() interface{} {
	return arn.Args
}

// Attributes returns the attributes for [DataSource].
func (arn *DataSource) Attributes() dataAwsRedshiftserverlessNamespaceAttributes {
	return dataAwsRedshiftserverlessNamespaceAttributes{ref: terra.ReferenceDataSource(arn)}
}

// DataArgs contains the configurations for aws_redshiftserverless_namespace.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
}

type dataAwsRedshiftserverlessNamespaceAttributes struct {
	ref terra.Reference
}

// AdminUsername returns a reference to field admin_username of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) AdminUsername() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("admin_username"))
}

// Arn returns a reference to field arn of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("arn"))
}

// DbName returns a reference to field db_name of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) DbName() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("db_name"))
}

// DefaultIamRoleArn returns a reference to field default_iam_role_arn of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) DefaultIamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("default_iam_role_arn"))
}

// IamRoles returns a reference to field iam_roles of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) IamRoles() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](arn.ref.Append("iam_roles"))
}

// Id returns a reference to field id of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("kms_key_id"))
}

// LogExports returns a reference to field log_exports of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) LogExports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](arn.ref.Append("log_exports"))
}

// NamespaceId returns a reference to field namespace_id of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("namespace_id"))
}

// NamespaceName returns a reference to field namespace_name of aws_redshiftserverless_namespace.
func (arn dataAwsRedshiftserverlessNamespaceAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(arn.ref.Append("namespace_name"))
}
