// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataCognitoUserPoolClients(name string, args DataCognitoUserPoolClientsArgs) *DataCognitoUserPoolClients {
	return &DataCognitoUserPoolClients{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCognitoUserPoolClients)(nil)

type DataCognitoUserPoolClients struct {
	Name string
	Args DataCognitoUserPoolClientsArgs
}

func (cupc *DataCognitoUserPoolClients) DataSource() string {
	return "aws_cognito_user_pool_clients"
}

func (cupc *DataCognitoUserPoolClients) LocalName() string {
	return cupc.Name
}

func (cupc *DataCognitoUserPoolClients) Configuration() interface{} {
	return cupc.Args
}

func (cupc *DataCognitoUserPoolClients) Attributes() dataCognitoUserPoolClientsAttributes {
	return dataCognitoUserPoolClientsAttributes{ref: terra.ReferenceDataResource(cupc)}
}

type DataCognitoUserPoolClientsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
}
type dataCognitoUserPoolClientsAttributes struct {
	ref terra.Reference
}

func (cupc dataCognitoUserPoolClientsAttributes) ClientIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](cupc.ref.Append("client_ids"))
}

func (cupc dataCognitoUserPoolClientsAttributes) ClientNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](cupc.ref.Append("client_names"))
}

func (cupc dataCognitoUserPoolClientsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cupc.ref.Append("id"))
}

func (cupc dataCognitoUserPoolClientsAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceString(cupc.ref.Append("user_pool_id"))
}
