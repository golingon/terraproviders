// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasecretsmanagersecret "github.com/golingon/terraproviders/aws/4.66.1/datasecretsmanagersecret"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSecretsmanagerSecret creates a new instance of [DataSecretsmanagerSecret].
func NewDataSecretsmanagerSecret(name string, args DataSecretsmanagerSecretArgs) *DataSecretsmanagerSecret {
	return &DataSecretsmanagerSecret{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretsmanagerSecret)(nil)

// DataSecretsmanagerSecret represents the Terraform data resource aws_secretsmanager_secret.
type DataSecretsmanagerSecret struct {
	Name string
	Args DataSecretsmanagerSecretArgs
}

// DataSource returns the Terraform object type for [DataSecretsmanagerSecret].
func (ss *DataSecretsmanagerSecret) DataSource() string {
	return "aws_secretsmanager_secret"
}

// LocalName returns the local name for [DataSecretsmanagerSecret].
func (ss *DataSecretsmanagerSecret) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [DataSecretsmanagerSecret].
func (ss *DataSecretsmanagerSecret) Configuration() interface{} {
	return ss.Args
}

// Attributes returns the attributes for [DataSecretsmanagerSecret].
func (ss *DataSecretsmanagerSecret) Attributes() dataSecretsmanagerSecretAttributes {
	return dataSecretsmanagerSecretAttributes{ref: terra.ReferenceDataResource(ss)}
}

// DataSecretsmanagerSecretArgs contains the configurations for aws_secretsmanager_secret.
type DataSecretsmanagerSecretArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RotationRules: min=0
	RotationRules []datasecretsmanagersecret.RotationRules `hcl:"rotation_rules,block" validate:"min=0"`
}
type dataSecretsmanagerSecretAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("arn"))
}

// Description returns a reference to field description of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("description"))
}

// Id returns a reference to field id of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("kms_key_id"))
}

// Name returns a reference to field name of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("name"))
}

// Policy returns a reference to field policy of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("policy"))
}

// RotationEnabled returns a reference to field rotation_enabled of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) RotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ss.ref.Append("rotation_enabled"))
}

// RotationLambdaArn returns a reference to field rotation_lambda_arn of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) RotationLambdaArn() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("rotation_lambda_arn"))
}

// Tags returns a reference to field tags of aws_secretsmanager_secret.
func (ss dataSecretsmanagerSecretAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

func (ss dataSecretsmanagerSecretAttributes) RotationRules() terra.ListValue[datasecretsmanagersecret.RotationRulesAttributes] {
	return terra.ReferenceAsList[datasecretsmanagersecret.RotationRulesAttributes](ss.ref.Append("rotation_rules"))
}