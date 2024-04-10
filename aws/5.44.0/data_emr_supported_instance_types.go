// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	dataemrsupportedinstancetypes "github.com/golingon/terraproviders/aws/5.44.0/dataemrsupportedinstancetypes"
)

// NewDataEmrSupportedInstanceTypes creates a new instance of [DataEmrSupportedInstanceTypes].
func NewDataEmrSupportedInstanceTypes(name string, args DataEmrSupportedInstanceTypesArgs) *DataEmrSupportedInstanceTypes {
	return &DataEmrSupportedInstanceTypes{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEmrSupportedInstanceTypes)(nil)

// DataEmrSupportedInstanceTypes represents the Terraform data resource aws_emr_supported_instance_types.
type DataEmrSupportedInstanceTypes struct {
	Name string
	Args DataEmrSupportedInstanceTypesArgs
}

// DataSource returns the Terraform object type for [DataEmrSupportedInstanceTypes].
func (esit *DataEmrSupportedInstanceTypes) DataSource() string {
	return "aws_emr_supported_instance_types"
}

// LocalName returns the local name for [DataEmrSupportedInstanceTypes].
func (esit *DataEmrSupportedInstanceTypes) LocalName() string {
	return esit.Name
}

// Configuration returns the configuration (args) for [DataEmrSupportedInstanceTypes].
func (esit *DataEmrSupportedInstanceTypes) Configuration() interface{} {
	return esit.Args
}

// Attributes returns the attributes for [DataEmrSupportedInstanceTypes].
func (esit *DataEmrSupportedInstanceTypes) Attributes() dataEmrSupportedInstanceTypesAttributes {
	return dataEmrSupportedInstanceTypesAttributes{ref: terra.ReferenceDataResource(esit)}
}

// DataEmrSupportedInstanceTypesArgs contains the configurations for aws_emr_supported_instance_types.
type DataEmrSupportedInstanceTypesArgs struct {
	// ReleaseLabel: string, required
	ReleaseLabel terra.StringValue `hcl:"release_label,attr" validate:"required"`
	// SupportedInstanceTypes: min=0
	SupportedInstanceTypes []dataemrsupportedinstancetypes.SupportedInstanceTypes `hcl:"supported_instance_types,block" validate:"min=0"`
}
type dataEmrSupportedInstanceTypesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_emr_supported_instance_types.
func (esit dataEmrSupportedInstanceTypesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esit.ref.Append("id"))
}

// ReleaseLabel returns a reference to field release_label of aws_emr_supported_instance_types.
func (esit dataEmrSupportedInstanceTypesAttributes) ReleaseLabel() terra.StringValue {
	return terra.ReferenceAsString(esit.ref.Append("release_label"))
}

func (esit dataEmrSupportedInstanceTypesAttributes) SupportedInstanceTypes() terra.ListValue[dataemrsupportedinstancetypes.SupportedInstanceTypesAttributes] {
	return terra.ReferenceAsList[dataemrsupportedinstancetypes.SupportedInstanceTypesAttributes](esit.ref.Append("supported_instance_types"))
}
