// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasecretsmanagersecretrotation "github.com/golingon/terraproviders/aws/5.7.0/datasecretsmanagersecretrotation"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSecretsmanagerSecretRotation creates a new instance of [DataSecretsmanagerSecretRotation].
func NewDataSecretsmanagerSecretRotation(name string, args DataSecretsmanagerSecretRotationArgs) *DataSecretsmanagerSecretRotation {
	return &DataSecretsmanagerSecretRotation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretsmanagerSecretRotation)(nil)

// DataSecretsmanagerSecretRotation represents the Terraform data resource aws_secretsmanager_secret_rotation.
type DataSecretsmanagerSecretRotation struct {
	Name string
	Args DataSecretsmanagerSecretRotationArgs
}

// DataSource returns the Terraform object type for [DataSecretsmanagerSecretRotation].
func (ssr *DataSecretsmanagerSecretRotation) DataSource() string {
	return "aws_secretsmanager_secret_rotation"
}

// LocalName returns the local name for [DataSecretsmanagerSecretRotation].
func (ssr *DataSecretsmanagerSecretRotation) LocalName() string {
	return ssr.Name
}

// Configuration returns the configuration (args) for [DataSecretsmanagerSecretRotation].
func (ssr *DataSecretsmanagerSecretRotation) Configuration() interface{} {
	return ssr.Args
}

// Attributes returns the attributes for [DataSecretsmanagerSecretRotation].
func (ssr *DataSecretsmanagerSecretRotation) Attributes() dataSecretsmanagerSecretRotationAttributes {
	return dataSecretsmanagerSecretRotationAttributes{ref: terra.ReferenceDataResource(ssr)}
}

// DataSecretsmanagerSecretRotationArgs contains the configurations for aws_secretsmanager_secret_rotation.
type DataSecretsmanagerSecretRotationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// RotationRules: min=0
	RotationRules []datasecretsmanagersecretrotation.RotationRules `hcl:"rotation_rules,block" validate:"min=0"`
}
type dataSecretsmanagerSecretRotationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_secretsmanager_secret_rotation.
func (ssr dataSecretsmanagerSecretRotationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssr.ref.Append("id"))
}

// RotationEnabled returns a reference to field rotation_enabled of aws_secretsmanager_secret_rotation.
func (ssr dataSecretsmanagerSecretRotationAttributes) RotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssr.ref.Append("rotation_enabled"))
}

// RotationLambdaArn returns a reference to field rotation_lambda_arn of aws_secretsmanager_secret_rotation.
func (ssr dataSecretsmanagerSecretRotationAttributes) RotationLambdaArn() terra.StringValue {
	return terra.ReferenceAsString(ssr.ref.Append("rotation_lambda_arn"))
}

// SecretId returns a reference to field secret_id of aws_secretsmanager_secret_rotation.
func (ssr dataSecretsmanagerSecretRotationAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(ssr.ref.Append("secret_id"))
}

func (ssr dataSecretsmanagerSecretRotationAttributes) RotationRules() terra.ListValue[datasecretsmanagersecretrotation.RotationRulesAttributes] {
	return terra.ReferenceAsList[datasecretsmanagersecretrotation.RotationRulesAttributes](ssr.ref.Append("rotation_rules"))
}
