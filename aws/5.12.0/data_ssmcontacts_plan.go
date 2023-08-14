// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datassmcontactsplan "github.com/golingon/terraproviders/aws/5.12.0/datassmcontactsplan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSsmcontactsPlan creates a new instance of [DataSsmcontactsPlan].
func NewDataSsmcontactsPlan(name string, args DataSsmcontactsPlanArgs) *DataSsmcontactsPlan {
	return &DataSsmcontactsPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmcontactsPlan)(nil)

// DataSsmcontactsPlan represents the Terraform data resource aws_ssmcontacts_plan.
type DataSsmcontactsPlan struct {
	Name string
	Args DataSsmcontactsPlanArgs
}

// DataSource returns the Terraform object type for [DataSsmcontactsPlan].
func (sp *DataSsmcontactsPlan) DataSource() string {
	return "aws_ssmcontacts_plan"
}

// LocalName returns the local name for [DataSsmcontactsPlan].
func (sp *DataSsmcontactsPlan) LocalName() string {
	return sp.Name
}

// Configuration returns the configuration (args) for [DataSsmcontactsPlan].
func (sp *DataSsmcontactsPlan) Configuration() interface{} {
	return sp.Args
}

// Attributes returns the attributes for [DataSsmcontactsPlan].
func (sp *DataSsmcontactsPlan) Attributes() dataSsmcontactsPlanAttributes {
	return dataSsmcontactsPlanAttributes{ref: terra.ReferenceDataResource(sp)}
}

// DataSsmcontactsPlanArgs contains the configurations for aws_ssmcontacts_plan.
type DataSsmcontactsPlanArgs struct {
	// ContactId: string, required
	ContactId terra.StringValue `hcl:"contact_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Stage: min=0
	Stage []datassmcontactsplan.Stage `hcl:"stage,block" validate:"min=0"`
}
type dataSsmcontactsPlanAttributes struct {
	ref terra.Reference
}

// ContactId returns a reference to field contact_id of aws_ssmcontacts_plan.
func (sp dataSsmcontactsPlanAttributes) ContactId() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("contact_id"))
}

// Id returns a reference to field id of aws_ssmcontacts_plan.
func (sp dataSsmcontactsPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("id"))
}

func (sp dataSsmcontactsPlanAttributes) Stage() terra.ListValue[datassmcontactsplan.StageAttributes] {
	return terra.ReferenceAsList[datassmcontactsplan.StageAttributes](sp.ref.Append("stage"))
}
