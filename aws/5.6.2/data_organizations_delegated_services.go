// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataorganizationsdelegatedservices "github.com/golingon/terraproviders/aws/5.6.2/dataorganizationsdelegatedservices"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOrganizationsDelegatedServices creates a new instance of [DataOrganizationsDelegatedServices].
func NewDataOrganizationsDelegatedServices(name string, args DataOrganizationsDelegatedServicesArgs) *DataOrganizationsDelegatedServices {
	return &DataOrganizationsDelegatedServices{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationsDelegatedServices)(nil)

// DataOrganizationsDelegatedServices represents the Terraform data resource aws_organizations_delegated_services.
type DataOrganizationsDelegatedServices struct {
	Name string
	Args DataOrganizationsDelegatedServicesArgs
}

// DataSource returns the Terraform object type for [DataOrganizationsDelegatedServices].
func (ods *DataOrganizationsDelegatedServices) DataSource() string {
	return "aws_organizations_delegated_services"
}

// LocalName returns the local name for [DataOrganizationsDelegatedServices].
func (ods *DataOrganizationsDelegatedServices) LocalName() string {
	return ods.Name
}

// Configuration returns the configuration (args) for [DataOrganizationsDelegatedServices].
func (ods *DataOrganizationsDelegatedServices) Configuration() interface{} {
	return ods.Args
}

// Attributes returns the attributes for [DataOrganizationsDelegatedServices].
func (ods *DataOrganizationsDelegatedServices) Attributes() dataOrganizationsDelegatedServicesAttributes {
	return dataOrganizationsDelegatedServicesAttributes{ref: terra.ReferenceDataResource(ods)}
}

// DataOrganizationsDelegatedServicesArgs contains the configurations for aws_organizations_delegated_services.
type DataOrganizationsDelegatedServicesArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DelegatedServices: min=0
	DelegatedServices []dataorganizationsdelegatedservices.DelegatedServices `hcl:"delegated_services,block" validate:"min=0"`
}
type dataOrganizationsDelegatedServicesAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_organizations_delegated_services.
func (ods dataOrganizationsDelegatedServicesAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(ods.ref.Append("account_id"))
}

// Id returns a reference to field id of aws_organizations_delegated_services.
func (ods dataOrganizationsDelegatedServicesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ods.ref.Append("id"))
}

func (ods dataOrganizationsDelegatedServicesAttributes) DelegatedServices() terra.SetValue[dataorganizationsdelegatedservices.DelegatedServicesAttributes] {
	return terra.ReferenceAsSet[dataorganizationsdelegatedservices.DelegatedServicesAttributes](ods.ref.Append("delegated_services"))
}
