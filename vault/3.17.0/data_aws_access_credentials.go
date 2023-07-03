// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataAwsAccessCredentials creates a new instance of [DataAwsAccessCredentials].
func NewDataAwsAccessCredentials(name string, args DataAwsAccessCredentialsArgs) *DataAwsAccessCredentials {
	return &DataAwsAccessCredentials{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAwsAccessCredentials)(nil)

// DataAwsAccessCredentials represents the Terraform data resource vault_aws_access_credentials.
type DataAwsAccessCredentials struct {
	Name string
	Args DataAwsAccessCredentialsArgs
}

// DataSource returns the Terraform object type for [DataAwsAccessCredentials].
func (aac *DataAwsAccessCredentials) DataSource() string {
	return "vault_aws_access_credentials"
}

// LocalName returns the local name for [DataAwsAccessCredentials].
func (aac *DataAwsAccessCredentials) LocalName() string {
	return aac.Name
}

// Configuration returns the configuration (args) for [DataAwsAccessCredentials].
func (aac *DataAwsAccessCredentials) Configuration() interface{} {
	return aac.Args
}

// Attributes returns the attributes for [DataAwsAccessCredentials].
func (aac *DataAwsAccessCredentials) Attributes() dataAwsAccessCredentialsAttributes {
	return dataAwsAccessCredentialsAttributes{ref: terra.ReferenceDataResource(aac)}
}

// DataAwsAccessCredentialsArgs contains the configurations for vault_aws_access_credentials.
type DataAwsAccessCredentialsArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// Ttl: string, optional
	Ttl terra.StringValue `hcl:"ttl,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}
type dataAwsAccessCredentialsAttributes struct {
	ref terra.Reference
}

// AccessKey returns a reference to field access_key of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) AccessKey() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("access_key"))
}

// Backend returns a reference to field backend of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("backend"))
}

// Id returns a reference to field id of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("id"))
}

// LeaseDuration returns a reference to field lease_duration of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) LeaseDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(aac.ref.Append("lease_duration"))
}

// LeaseId returns a reference to field lease_id of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) LeaseId() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("lease_id"))
}

// LeaseRenewable returns a reference to field lease_renewable of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) LeaseRenewable() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("lease_renewable"))
}

// LeaseStartTime returns a reference to field lease_start_time of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) LeaseStartTime() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("lease_start_time"))
}

// Namespace returns a reference to field namespace of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("namespace"))
}

// Region returns a reference to field region of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("region"))
}

// Role returns a reference to field role of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("role"))
}

// RoleArn returns a reference to field role_arn of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("role_arn"))
}

// SecretKey returns a reference to field secret_key of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) SecretKey() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("secret_key"))
}

// SecurityToken returns a reference to field security_token of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) SecurityToken() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("security_token"))
}

// Ttl returns a reference to field ttl of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Ttl() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("ttl"))
}

// Type returns a reference to field type of vault_aws_access_credentials.
func (aac dataAwsAccessCredentialsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("type"))
}