// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataIamSessionContext(name string, args DataIamSessionContextArgs) *DataIamSessionContext {
	return &DataIamSessionContext{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamSessionContext)(nil)

type DataIamSessionContext struct {
	Name string
	Args DataIamSessionContextArgs
}

func (isc *DataIamSessionContext) DataSource() string {
	return "aws_iam_session_context"
}

func (isc *DataIamSessionContext) LocalName() string {
	return isc.Name
}

func (isc *DataIamSessionContext) Configuration() interface{} {
	return isc.Args
}

func (isc *DataIamSessionContext) Attributes() dataIamSessionContextAttributes {
	return dataIamSessionContextAttributes{ref: terra.ReferenceDataResource(isc)}
}

type DataIamSessionContextArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataIamSessionContextAttributes struct {
	ref terra.Reference
}

func (isc dataIamSessionContextAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(isc.ref.Append("arn"))
}

func (isc dataIamSessionContextAttributes) Id() terra.StringValue {
	return terra.ReferenceString(isc.ref.Append("id"))
}

func (isc dataIamSessionContextAttributes) IssuerArn() terra.StringValue {
	return terra.ReferenceString(isc.ref.Append("issuer_arn"))
}

func (isc dataIamSessionContextAttributes) IssuerId() terra.StringValue {
	return terra.ReferenceString(isc.ref.Append("issuer_id"))
}

func (isc dataIamSessionContextAttributes) IssuerName() terra.StringValue {
	return terra.ReferenceString(isc.ref.Append("issuer_name"))
}

func (isc dataIamSessionContextAttributes) SessionName() terra.StringValue {
	return terra.ReferenceString(isc.ref.Append("session_name"))
}
