// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataRoute53DelegationSet(name string, args DataRoute53DelegationSetArgs) *DataRoute53DelegationSet {
	return &DataRoute53DelegationSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute53DelegationSet)(nil)

type DataRoute53DelegationSet struct {
	Name string
	Args DataRoute53DelegationSetArgs
}

func (rds *DataRoute53DelegationSet) DataSource() string {
	return "aws_route53_delegation_set"
}

func (rds *DataRoute53DelegationSet) LocalName() string {
	return rds.Name
}

func (rds *DataRoute53DelegationSet) Configuration() interface{} {
	return rds.Args
}

func (rds *DataRoute53DelegationSet) Attributes() dataRoute53DelegationSetAttributes {
	return dataRoute53DelegationSetAttributes{ref: terra.ReferenceDataResource(rds)}
}

type DataRoute53DelegationSetArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
}
type dataRoute53DelegationSetAttributes struct {
	ref terra.Reference
}

func (rds dataRoute53DelegationSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(rds.ref.Append("arn"))
}

func (rds dataRoute53DelegationSetAttributes) CallerReference() terra.StringValue {
	return terra.ReferenceString(rds.ref.Append("caller_reference"))
}

func (rds dataRoute53DelegationSetAttributes) Id() terra.StringValue {
	return terra.ReferenceString(rds.ref.Append("id"))
}

func (rds dataRoute53DelegationSetAttributes) NameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](rds.ref.Append("name_servers"))
}
