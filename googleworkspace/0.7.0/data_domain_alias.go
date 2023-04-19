// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googleworkspace

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDomainAlias creates a new instance of [DataDomainAlias].
func NewDataDomainAlias(name string, args DataDomainAliasArgs) *DataDomainAlias {
	return &DataDomainAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDomainAlias)(nil)

// DataDomainAlias represents the Terraform data resource googleworkspace_domain_alias.
type DataDomainAlias struct {
	Name string
	Args DataDomainAliasArgs
}

// DataSource returns the Terraform object type for [DataDomainAlias].
func (da *DataDomainAlias) DataSource() string {
	return "googleworkspace_domain_alias"
}

// LocalName returns the local name for [DataDomainAlias].
func (da *DataDomainAlias) LocalName() string {
	return da.Name
}

// Configuration returns the configuration (args) for [DataDomainAlias].
func (da *DataDomainAlias) Configuration() interface{} {
	return da.Args
}

// Attributes returns the attributes for [DataDomainAlias].
func (da *DataDomainAlias) Attributes() dataDomainAliasAttributes {
	return dataDomainAliasAttributes{ref: terra.ReferenceDataResource(da)}
}

// DataDomainAliasArgs contains the configurations for googleworkspace_domain_alias.
type DataDomainAliasArgs struct {
	// DomainAliasName: string, required
	DomainAliasName terra.StringValue `hcl:"domain_alias_name,attr" validate:"required"`
}
type dataDomainAliasAttributes struct {
	ref terra.Reference
}

// CreationTime returns a reference to field creation_time of googleworkspace_domain_alias.
func (da dataDomainAliasAttributes) CreationTime() terra.NumberValue {
	return terra.ReferenceAsNumber(da.ref.Append("creation_time"))
}

// DomainAliasName returns a reference to field domain_alias_name of googleworkspace_domain_alias.
func (da dataDomainAliasAttributes) DomainAliasName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("domain_alias_name"))
}

// Etag returns a reference to field etag of googleworkspace_domain_alias.
func (da dataDomainAliasAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("etag"))
}

// Id returns a reference to field id of googleworkspace_domain_alias.
func (da dataDomainAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("id"))
}

// ParentDomainName returns a reference to field parent_domain_name of googleworkspace_domain_alias.
func (da dataDomainAliasAttributes) ParentDomainName() terra.StringValue {
	return terra.ReferenceAsString(da.ref.Append("parent_domain_name"))
}

// Verified returns a reference to field verified of googleworkspace_domain_alias.
func (da dataDomainAliasAttributes) Verified() terra.BoolValue {
	return terra.ReferenceAsBool(da.ref.Append("verified"))
}