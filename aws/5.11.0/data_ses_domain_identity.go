// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSesDomainIdentity creates a new instance of [DataSesDomainIdentity].
func NewDataSesDomainIdentity(name string, args DataSesDomainIdentityArgs) *DataSesDomainIdentity {
	return &DataSesDomainIdentity{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSesDomainIdentity)(nil)

// DataSesDomainIdentity represents the Terraform data resource aws_ses_domain_identity.
type DataSesDomainIdentity struct {
	Name string
	Args DataSesDomainIdentityArgs
}

// DataSource returns the Terraform object type for [DataSesDomainIdentity].
func (sdi *DataSesDomainIdentity) DataSource() string {
	return "aws_ses_domain_identity"
}

// LocalName returns the local name for [DataSesDomainIdentity].
func (sdi *DataSesDomainIdentity) LocalName() string {
	return sdi.Name
}

// Configuration returns the configuration (args) for [DataSesDomainIdentity].
func (sdi *DataSesDomainIdentity) Configuration() interface{} {
	return sdi.Args
}

// Attributes returns the attributes for [DataSesDomainIdentity].
func (sdi *DataSesDomainIdentity) Attributes() dataSesDomainIdentityAttributes {
	return dataSesDomainIdentityAttributes{ref: terra.ReferenceDataResource(sdi)}
}

// DataSesDomainIdentityArgs contains the configurations for aws_ses_domain_identity.
type DataSesDomainIdentityArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataSesDomainIdentityAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ses_domain_identity.
func (sdi dataSesDomainIdentityAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("arn"))
}

// Domain returns a reference to field domain of aws_ses_domain_identity.
func (sdi dataSesDomainIdentityAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("domain"))
}

// Id returns a reference to field id of aws_ses_domain_identity.
func (sdi dataSesDomainIdentityAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("id"))
}

// VerificationToken returns a reference to field verification_token of aws_ses_domain_identity.
func (sdi dataSesDomainIdentityAttributes) VerificationToken() terra.StringValue {
	return terra.ReferenceAsString(sdi.ref.Append("verification_token"))
}