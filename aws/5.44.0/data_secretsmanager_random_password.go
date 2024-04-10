// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataSecretsmanagerRandomPassword creates a new instance of [DataSecretsmanagerRandomPassword].
func NewDataSecretsmanagerRandomPassword(name string, args DataSecretsmanagerRandomPasswordArgs) *DataSecretsmanagerRandomPassword {
	return &DataSecretsmanagerRandomPassword{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSecretsmanagerRandomPassword)(nil)

// DataSecretsmanagerRandomPassword represents the Terraform data resource aws_secretsmanager_random_password.
type DataSecretsmanagerRandomPassword struct {
	Name string
	Args DataSecretsmanagerRandomPasswordArgs
}

// DataSource returns the Terraform object type for [DataSecretsmanagerRandomPassword].
func (srp *DataSecretsmanagerRandomPassword) DataSource() string {
	return "aws_secretsmanager_random_password"
}

// LocalName returns the local name for [DataSecretsmanagerRandomPassword].
func (srp *DataSecretsmanagerRandomPassword) LocalName() string {
	return srp.Name
}

// Configuration returns the configuration (args) for [DataSecretsmanagerRandomPassword].
func (srp *DataSecretsmanagerRandomPassword) Configuration() interface{} {
	return srp.Args
}

// Attributes returns the attributes for [DataSecretsmanagerRandomPassword].
func (srp *DataSecretsmanagerRandomPassword) Attributes() dataSecretsmanagerRandomPasswordAttributes {
	return dataSecretsmanagerRandomPasswordAttributes{ref: terra.ReferenceDataResource(srp)}
}

// DataSecretsmanagerRandomPasswordArgs contains the configurations for aws_secretsmanager_random_password.
type DataSecretsmanagerRandomPasswordArgs struct {
	// ExcludeCharacters: string, optional
	ExcludeCharacters terra.StringValue `hcl:"exclude_characters,attr"`
	// ExcludeLowercase: bool, optional
	ExcludeLowercase terra.BoolValue `hcl:"exclude_lowercase,attr"`
	// ExcludeNumbers: bool, optional
	ExcludeNumbers terra.BoolValue `hcl:"exclude_numbers,attr"`
	// ExcludePunctuation: bool, optional
	ExcludePunctuation terra.BoolValue `hcl:"exclude_punctuation,attr"`
	// ExcludeUppercase: bool, optional
	ExcludeUppercase terra.BoolValue `hcl:"exclude_uppercase,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeSpace: bool, optional
	IncludeSpace terra.BoolValue `hcl:"include_space,attr"`
	// PasswordLength: number, optional
	PasswordLength terra.NumberValue `hcl:"password_length,attr"`
	// RequireEachIncludedType: bool, optional
	RequireEachIncludedType terra.BoolValue `hcl:"require_each_included_type,attr"`
}
type dataSecretsmanagerRandomPasswordAttributes struct {
	ref terra.Reference
}

// ExcludeCharacters returns a reference to field exclude_characters of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) ExcludeCharacters() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("exclude_characters"))
}

// ExcludeLowercase returns a reference to field exclude_lowercase of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) ExcludeLowercase() terra.BoolValue {
	return terra.ReferenceAsBool(srp.ref.Append("exclude_lowercase"))
}

// ExcludeNumbers returns a reference to field exclude_numbers of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) ExcludeNumbers() terra.BoolValue {
	return terra.ReferenceAsBool(srp.ref.Append("exclude_numbers"))
}

// ExcludePunctuation returns a reference to field exclude_punctuation of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) ExcludePunctuation() terra.BoolValue {
	return terra.ReferenceAsBool(srp.ref.Append("exclude_punctuation"))
}

// ExcludeUppercase returns a reference to field exclude_uppercase of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) ExcludeUppercase() terra.BoolValue {
	return terra.ReferenceAsBool(srp.ref.Append("exclude_uppercase"))
}

// Id returns a reference to field id of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("id"))
}

// IncludeSpace returns a reference to field include_space of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) IncludeSpace() terra.BoolValue {
	return terra.ReferenceAsBool(srp.ref.Append("include_space"))
}

// PasswordLength returns a reference to field password_length of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) PasswordLength() terra.NumberValue {
	return terra.ReferenceAsNumber(srp.ref.Append("password_length"))
}

// RandomPassword returns a reference to field random_password of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) RandomPassword() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("random_password"))
}

// RequireEachIncludedType returns a reference to field require_each_included_type of aws_secretsmanager_random_password.
func (srp dataSecretsmanagerRandomPasswordAttributes) RequireEachIncludedType() terra.BoolValue {
	return terra.ReferenceAsBool(srp.ref.Append("require_each_included_type"))
}
