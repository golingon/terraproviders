// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	secretsmanagersecretrotation "github.com/golingon/terraproviders/aws/5.0.1/secretsmanagersecretrotation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSecretsmanagerSecretRotation creates a new instance of [SecretsmanagerSecretRotation].
func NewSecretsmanagerSecretRotation(name string, args SecretsmanagerSecretRotationArgs) *SecretsmanagerSecretRotation {
	return &SecretsmanagerSecretRotation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretsmanagerSecretRotation)(nil)

// SecretsmanagerSecretRotation represents the Terraform resource aws_secretsmanager_secret_rotation.
type SecretsmanagerSecretRotation struct {
	Name      string
	Args      SecretsmanagerSecretRotationArgs
	state     *secretsmanagerSecretRotationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretsmanagerSecretRotation].
func (ssr *SecretsmanagerSecretRotation) Type() string {
	return "aws_secretsmanager_secret_rotation"
}

// LocalName returns the local name for [SecretsmanagerSecretRotation].
func (ssr *SecretsmanagerSecretRotation) LocalName() string {
	return ssr.Name
}

// Configuration returns the configuration (args) for [SecretsmanagerSecretRotation].
func (ssr *SecretsmanagerSecretRotation) Configuration() interface{} {
	return ssr.Args
}

// DependOn is used for other resources to depend on [SecretsmanagerSecretRotation].
func (ssr *SecretsmanagerSecretRotation) DependOn() terra.Reference {
	return terra.ReferenceResource(ssr)
}

// Dependencies returns the list of resources [SecretsmanagerSecretRotation] depends_on.
func (ssr *SecretsmanagerSecretRotation) Dependencies() terra.Dependencies {
	return ssr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretsmanagerSecretRotation].
func (ssr *SecretsmanagerSecretRotation) LifecycleManagement() *terra.Lifecycle {
	return ssr.Lifecycle
}

// Attributes returns the attributes for [SecretsmanagerSecretRotation].
func (ssr *SecretsmanagerSecretRotation) Attributes() secretsmanagerSecretRotationAttributes {
	return secretsmanagerSecretRotationAttributes{ref: terra.ReferenceResource(ssr)}
}

// ImportState imports the given attribute values into [SecretsmanagerSecretRotation]'s state.
func (ssr *SecretsmanagerSecretRotation) ImportState(av io.Reader) error {
	ssr.state = &secretsmanagerSecretRotationState{}
	if err := json.NewDecoder(av).Decode(ssr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssr.Type(), ssr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretsmanagerSecretRotation] has state.
func (ssr *SecretsmanagerSecretRotation) State() (*secretsmanagerSecretRotationState, bool) {
	return ssr.state, ssr.state != nil
}

// StateMust returns the state for [SecretsmanagerSecretRotation]. Panics if the state is nil.
func (ssr *SecretsmanagerSecretRotation) StateMust() *secretsmanagerSecretRotationState {
	if ssr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssr.Type(), ssr.LocalName()))
	}
	return ssr.state
}

// SecretsmanagerSecretRotationArgs contains the configurations for aws_secretsmanager_secret_rotation.
type SecretsmanagerSecretRotationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RotationLambdaArn: string, required
	RotationLambdaArn terra.StringValue `hcl:"rotation_lambda_arn,attr" validate:"required"`
	// SecretId: string, required
	SecretId terra.StringValue `hcl:"secret_id,attr" validate:"required"`
	// RotationRules: required
	RotationRules *secretsmanagersecretrotation.RotationRules `hcl:"rotation_rules,block" validate:"required"`
}
type secretsmanagerSecretRotationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_secretsmanager_secret_rotation.
func (ssr secretsmanagerSecretRotationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssr.ref.Append("id"))
}

// RotationEnabled returns a reference to field rotation_enabled of aws_secretsmanager_secret_rotation.
func (ssr secretsmanagerSecretRotationAttributes) RotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ssr.ref.Append("rotation_enabled"))
}

// RotationLambdaArn returns a reference to field rotation_lambda_arn of aws_secretsmanager_secret_rotation.
func (ssr secretsmanagerSecretRotationAttributes) RotationLambdaArn() terra.StringValue {
	return terra.ReferenceAsString(ssr.ref.Append("rotation_lambda_arn"))
}

// SecretId returns a reference to field secret_id of aws_secretsmanager_secret_rotation.
func (ssr secretsmanagerSecretRotationAttributes) SecretId() terra.StringValue {
	return terra.ReferenceAsString(ssr.ref.Append("secret_id"))
}

func (ssr secretsmanagerSecretRotationAttributes) RotationRules() terra.ListValue[secretsmanagersecretrotation.RotationRulesAttributes] {
	return terra.ReferenceAsList[secretsmanagersecretrotation.RotationRulesAttributes](ssr.ref.Append("rotation_rules"))
}

type secretsmanagerSecretRotationState struct {
	Id                string                                            `json:"id"`
	RotationEnabled   bool                                              `json:"rotation_enabled"`
	RotationLambdaArn string                                            `json:"rotation_lambda_arn"`
	SecretId          string                                            `json:"secret_id"`
	RotationRules     []secretsmanagersecretrotation.RotationRulesState `json:"rotation_rules"`
}
