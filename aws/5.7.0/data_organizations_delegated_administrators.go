// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataorganizationsdelegatedadministrators "github.com/golingon/terraproviders/aws/5.7.0/dataorganizationsdelegatedadministrators"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOrganizationsDelegatedAdministrators creates a new instance of [DataOrganizationsDelegatedAdministrators].
func NewDataOrganizationsDelegatedAdministrators(name string, args DataOrganizationsDelegatedAdministratorsArgs) *DataOrganizationsDelegatedAdministrators {
	return &DataOrganizationsDelegatedAdministrators{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationsDelegatedAdministrators)(nil)

// DataOrganizationsDelegatedAdministrators represents the Terraform data resource aws_organizations_delegated_administrators.
type DataOrganizationsDelegatedAdministrators struct {
	Name string
	Args DataOrganizationsDelegatedAdministratorsArgs
}

// DataSource returns the Terraform object type for [DataOrganizationsDelegatedAdministrators].
func (oda *DataOrganizationsDelegatedAdministrators) DataSource() string {
	return "aws_organizations_delegated_administrators"
}

// LocalName returns the local name for [DataOrganizationsDelegatedAdministrators].
func (oda *DataOrganizationsDelegatedAdministrators) LocalName() string {
	return oda.Name
}

// Configuration returns the configuration (args) for [DataOrganizationsDelegatedAdministrators].
func (oda *DataOrganizationsDelegatedAdministrators) Configuration() interface{} {
	return oda.Args
}

// Attributes returns the attributes for [DataOrganizationsDelegatedAdministrators].
func (oda *DataOrganizationsDelegatedAdministrators) Attributes() dataOrganizationsDelegatedAdministratorsAttributes {
	return dataOrganizationsDelegatedAdministratorsAttributes{ref: terra.ReferenceDataResource(oda)}
}

// DataOrganizationsDelegatedAdministratorsArgs contains the configurations for aws_organizations_delegated_administrators.
type DataOrganizationsDelegatedAdministratorsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServicePrincipal: string, optional
	ServicePrincipal terra.StringValue `hcl:"service_principal,attr"`
	// DelegatedAdministrators: min=0
	DelegatedAdministrators []dataorganizationsdelegatedadministrators.DelegatedAdministrators `hcl:"delegated_administrators,block" validate:"min=0"`
}
type dataOrganizationsDelegatedAdministratorsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_organizations_delegated_administrators.
func (oda dataOrganizationsDelegatedAdministratorsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oda.ref.Append("id"))
}

// ServicePrincipal returns a reference to field service_principal of aws_organizations_delegated_administrators.
func (oda dataOrganizationsDelegatedAdministratorsAttributes) ServicePrincipal() terra.StringValue {
	return terra.ReferenceAsString(oda.ref.Append("service_principal"))
}

func (oda dataOrganizationsDelegatedAdministratorsAttributes) DelegatedAdministrators() terra.SetValue[dataorganizationsdelegatedadministrators.DelegatedAdministratorsAttributes] {
	return terra.ReferenceAsSet[dataorganizationsdelegatedadministrators.DelegatedAdministratorsAttributes](oda.ref.Append("delegated_administrators"))
}
