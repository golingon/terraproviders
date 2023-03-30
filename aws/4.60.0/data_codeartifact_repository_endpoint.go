// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataCodeartifactRepositoryEndpoint(name string, args DataCodeartifactRepositoryEndpointArgs) *DataCodeartifactRepositoryEndpoint {
	return &DataCodeartifactRepositoryEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCodeartifactRepositoryEndpoint)(nil)

type DataCodeartifactRepositoryEndpoint struct {
	Name string
	Args DataCodeartifactRepositoryEndpointArgs
}

func (cre *DataCodeartifactRepositoryEndpoint) DataSource() string {
	return "aws_codeartifact_repository_endpoint"
}

func (cre *DataCodeartifactRepositoryEndpoint) LocalName() string {
	return cre.Name
}

func (cre *DataCodeartifactRepositoryEndpoint) Configuration() interface{} {
	return cre.Args
}

func (cre *DataCodeartifactRepositoryEndpoint) Attributes() dataCodeartifactRepositoryEndpointAttributes {
	return dataCodeartifactRepositoryEndpointAttributes{ref: terra.ReferenceDataResource(cre)}
}

type DataCodeartifactRepositoryEndpointArgs struct {
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// DomainOwner: string, optional
	DomainOwner terra.StringValue `hcl:"domain_owner,attr"`
	// Format: string, required
	Format terra.StringValue `hcl:"format,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Repository: string, required
	Repository terra.StringValue `hcl:"repository,attr" validate:"required"`
}
type dataCodeartifactRepositoryEndpointAttributes struct {
	ref terra.Reference
}

func (cre dataCodeartifactRepositoryEndpointAttributes) Domain() terra.StringValue {
	return terra.ReferenceString(cre.ref.Append("domain"))
}

func (cre dataCodeartifactRepositoryEndpointAttributes) DomainOwner() terra.StringValue {
	return terra.ReferenceString(cre.ref.Append("domain_owner"))
}

func (cre dataCodeartifactRepositoryEndpointAttributes) Format() terra.StringValue {
	return terra.ReferenceString(cre.ref.Append("format"))
}

func (cre dataCodeartifactRepositoryEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cre.ref.Append("id"))
}

func (cre dataCodeartifactRepositoryEndpointAttributes) Repository() terra.StringValue {
	return terra.ReferenceString(cre.ref.Append("repository"))
}

func (cre dataCodeartifactRepositoryEndpointAttributes) RepositoryEndpoint() terra.StringValue {
	return terra.ReferenceString(cre.ref.Append("repository_endpoint"))
}
