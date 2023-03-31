// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataServiceAccount(name string, args DataServiceAccountArgs) *DataServiceAccount {
	return &DataServiceAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataServiceAccount)(nil)

type DataServiceAccount struct {
	Name string
	Args DataServiceAccountArgs
}

func (sa *DataServiceAccount) DataSource() string {
	return "google_service_account"
}

func (sa *DataServiceAccount) LocalName() string {
	return sa.Name
}

func (sa *DataServiceAccount) Configuration() interface{} {
	return sa.Args
}

func (sa *DataServiceAccount) Attributes() dataServiceAccountAttributes {
	return dataServiceAccountAttributes{ref: terra.ReferenceDataResource(sa)}
}

type DataServiceAccountArgs struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataServiceAccountAttributes struct {
	ref terra.Reference
}

func (sa dataServiceAccountAttributes) AccountId() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("account_id"))
}

func (sa dataServiceAccountAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("display_name"))
}

func (sa dataServiceAccountAttributes) Email() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("email"))
}

func (sa dataServiceAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("id"))
}

func (sa dataServiceAccountAttributes) Member() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("member"))
}

func (sa dataServiceAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("name"))
}

func (sa dataServiceAccountAttributes) Project() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("project"))
}

func (sa dataServiceAccountAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceString(sa.ref.Append("unique_id"))
}
