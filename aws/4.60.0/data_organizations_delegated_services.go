// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataorganizationsdelegatedservices "github.com/golingon/terraproviders/aws/4.60.0/dataorganizationsdelegatedservices"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataOrganizationsDelegatedServices(name string, args DataOrganizationsDelegatedServicesArgs) *DataOrganizationsDelegatedServices {
	return &DataOrganizationsDelegatedServices{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationsDelegatedServices)(nil)

type DataOrganizationsDelegatedServices struct {
	Name string
	Args DataOrganizationsDelegatedServicesArgs
}

func (ods *DataOrganizationsDelegatedServices) DataSource() string {
	return "aws_organizations_delegated_services"
}

func (ods *DataOrganizationsDelegatedServices) LocalName() string {
	return ods.Name
}

func (ods *DataOrganizationsDelegatedServices) Configuration() interface{} {
	return ods.Args
}

func (ods *DataOrganizationsDelegatedServices) Attributes() dataOrganizationsDelegatedServicesAttributes {
	return dataOrganizationsDelegatedServicesAttributes{ref: terra.ReferenceDataResource(ods)}
}

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

func (ods dataOrganizationsDelegatedServicesAttributes) AccountId() terra.StringValue {
	return terra.ReferenceString(ods.ref.Append("account_id"))
}

func (ods dataOrganizationsDelegatedServicesAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ods.ref.Append("id"))
}

func (ods dataOrganizationsDelegatedServicesAttributes) DelegatedServices() terra.SetValue[dataorganizationsdelegatedservices.DelegatedServicesAttributes] {
	return terra.ReferenceSet[dataorganizationsdelegatedservices.DelegatedServicesAttributes](ods.ref.Append("delegated_services"))
}
