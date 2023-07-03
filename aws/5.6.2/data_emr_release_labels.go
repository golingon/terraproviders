// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataemrreleaselabels "github.com/golingon/terraproviders/aws/5.6.2/dataemrreleaselabels"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEmrReleaseLabels creates a new instance of [DataEmrReleaseLabels].
func NewDataEmrReleaseLabels(name string, args DataEmrReleaseLabelsArgs) *DataEmrReleaseLabels {
	return &DataEmrReleaseLabels{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEmrReleaseLabels)(nil)

// DataEmrReleaseLabels represents the Terraform data resource aws_emr_release_labels.
type DataEmrReleaseLabels struct {
	Name string
	Args DataEmrReleaseLabelsArgs
}

// DataSource returns the Terraform object type for [DataEmrReleaseLabels].
func (erl *DataEmrReleaseLabels) DataSource() string {
	return "aws_emr_release_labels"
}

// LocalName returns the local name for [DataEmrReleaseLabels].
func (erl *DataEmrReleaseLabels) LocalName() string {
	return erl.Name
}

// Configuration returns the configuration (args) for [DataEmrReleaseLabels].
func (erl *DataEmrReleaseLabels) Configuration() interface{} {
	return erl.Args
}

// Attributes returns the attributes for [DataEmrReleaseLabels].
func (erl *DataEmrReleaseLabels) Attributes() dataEmrReleaseLabelsAttributes {
	return dataEmrReleaseLabelsAttributes{ref: terra.ReferenceDataResource(erl)}
}

// DataEmrReleaseLabelsArgs contains the configurations for aws_emr_release_labels.
type DataEmrReleaseLabelsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filters: optional
	Filters *dataemrreleaselabels.Filters `hcl:"filters,block"`
}
type dataEmrReleaseLabelsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_emr_release_labels.
func (erl dataEmrReleaseLabelsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erl.ref.Append("id"))
}

// ReleaseLabels returns a reference to field release_labels of aws_emr_release_labels.
func (erl dataEmrReleaseLabelsAttributes) ReleaseLabels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](erl.ref.Append("release_labels"))
}

func (erl dataEmrReleaseLabelsAttributes) Filters() terra.ListValue[dataemrreleaselabels.FiltersAttributes] {
	return terra.ReferenceAsList[dataemrreleaselabels.FiltersAttributes](erl.ref.Append("filters"))
}
