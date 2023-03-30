// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataConnectLambdaFunctionAssociation(name string, args DataConnectLambdaFunctionAssociationArgs) *DataConnectLambdaFunctionAssociation {
	return &DataConnectLambdaFunctionAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectLambdaFunctionAssociation)(nil)

type DataConnectLambdaFunctionAssociation struct {
	Name string
	Args DataConnectLambdaFunctionAssociationArgs
}

func (clfa *DataConnectLambdaFunctionAssociation) DataSource() string {
	return "aws_connect_lambda_function_association"
}

func (clfa *DataConnectLambdaFunctionAssociation) LocalName() string {
	return clfa.Name
}

func (clfa *DataConnectLambdaFunctionAssociation) Configuration() interface{} {
	return clfa.Args
}

func (clfa *DataConnectLambdaFunctionAssociation) Attributes() dataConnectLambdaFunctionAssociationAttributes {
	return dataConnectLambdaFunctionAssociationAttributes{ref: terra.ReferenceDataResource(clfa)}
}

type DataConnectLambdaFunctionAssociationArgs struct {
	// FunctionArn: string, required
	FunctionArn terra.StringValue `hcl:"function_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
}
type dataConnectLambdaFunctionAssociationAttributes struct {
	ref terra.Reference
}

func (clfa dataConnectLambdaFunctionAssociationAttributes) FunctionArn() terra.StringValue {
	return terra.ReferenceString(clfa.ref.Append("function_arn"))
}

func (clfa dataConnectLambdaFunctionAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(clfa.ref.Append("id"))
}

func (clfa dataConnectLambdaFunctionAssociationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(clfa.ref.Append("instance_id"))
}
