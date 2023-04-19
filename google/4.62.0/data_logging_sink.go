// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataloggingsink "github.com/golingon/terraproviders/google/4.62.0/dataloggingsink"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLoggingSink creates a new instance of [DataLoggingSink].
func NewDataLoggingSink(name string, args DataLoggingSinkArgs) *DataLoggingSink {
	return &DataLoggingSink{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLoggingSink)(nil)

// DataLoggingSink represents the Terraform data resource google_logging_sink.
type DataLoggingSink struct {
	Name string
	Args DataLoggingSinkArgs
}

// DataSource returns the Terraform object type for [DataLoggingSink].
func (ls *DataLoggingSink) DataSource() string {
	return "google_logging_sink"
}

// LocalName returns the local name for [DataLoggingSink].
func (ls *DataLoggingSink) LocalName() string {
	return ls.Name
}

// Configuration returns the configuration (args) for [DataLoggingSink].
func (ls *DataLoggingSink) Configuration() interface{} {
	return ls.Args
}

// Attributes returns the attributes for [DataLoggingSink].
func (ls *DataLoggingSink) Attributes() dataLoggingSinkAttributes {
	return dataLoggingSinkAttributes{ref: terra.ReferenceDataResource(ls)}
}

// DataLoggingSinkArgs contains the configurations for google_logging_sink.
type DataLoggingSinkArgs struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// BigqueryOptions: min=0
	BigqueryOptions []dataloggingsink.BigqueryOptions `hcl:"bigquery_options,block" validate:"min=0"`
	// Exclusions: min=0
	Exclusions []dataloggingsink.Exclusions `hcl:"exclusions,block" validate:"min=0"`
}
type dataLoggingSinkAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_logging_sink.
func (ls dataLoggingSinkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("description"))
}

// Destination returns a reference to field destination of google_logging_sink.
func (ls dataLoggingSinkAttributes) Destination() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("destination"))
}

// Disabled returns a reference to field disabled of google_logging_sink.
func (ls dataLoggingSinkAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(ls.ref.Append("disabled"))
}

// Filter returns a reference to field filter of google_logging_sink.
func (ls dataLoggingSinkAttributes) Filter() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("filter"))
}

// Id returns a reference to field id of google_logging_sink.
func (ls dataLoggingSinkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("id"))
}

// Name returns a reference to field name of google_logging_sink.
func (ls dataLoggingSinkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("name"))
}

// WriterIdentity returns a reference to field writer_identity of google_logging_sink.
func (ls dataLoggingSinkAttributes) WriterIdentity() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("writer_identity"))
}

func (ls dataLoggingSinkAttributes) BigqueryOptions() terra.ListValue[dataloggingsink.BigqueryOptionsAttributes] {
	return terra.ReferenceAsList[dataloggingsink.BigqueryOptionsAttributes](ls.ref.Append("bigquery_options"))
}

func (ls dataLoggingSinkAttributes) Exclusions() terra.ListValue[dataloggingsink.ExclusionsAttributes] {
	return terra.ReferenceAsList[dataloggingsink.ExclusionsAttributes](ls.ref.Append("exclusions"))
}
