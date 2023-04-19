// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	datacloudfunctions2function "github.com/golingon/terraproviders/googlebeta/4.62.0/datacloudfunctions2function"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCloudfunctions2Function creates a new instance of [DataCloudfunctions2Function].
func NewDataCloudfunctions2Function(name string, args DataCloudfunctions2FunctionArgs) *DataCloudfunctions2Function {
	return &DataCloudfunctions2Function{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudfunctions2Function)(nil)

// DataCloudfunctions2Function represents the Terraform data resource google_cloudfunctions2_function.
type DataCloudfunctions2Function struct {
	Name string
	Args DataCloudfunctions2FunctionArgs
}

// DataSource returns the Terraform object type for [DataCloudfunctions2Function].
func (cf *DataCloudfunctions2Function) DataSource() string {
	return "google_cloudfunctions2_function"
}

// LocalName returns the local name for [DataCloudfunctions2Function].
func (cf *DataCloudfunctions2Function) LocalName() string {
	return cf.Name
}

// Configuration returns the configuration (args) for [DataCloudfunctions2Function].
func (cf *DataCloudfunctions2Function) Configuration() interface{} {
	return cf.Args
}

// Attributes returns the attributes for [DataCloudfunctions2Function].
func (cf *DataCloudfunctions2Function) Attributes() dataCloudfunctions2FunctionAttributes {
	return dataCloudfunctions2FunctionAttributes{ref: terra.ReferenceDataResource(cf)}
}

// DataCloudfunctions2FunctionArgs contains the configurations for google_cloudfunctions2_function.
type DataCloudfunctions2FunctionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BuildConfig: min=0
	BuildConfig []datacloudfunctions2function.BuildConfig `hcl:"build_config,block" validate:"min=0"`
	// EventTrigger: min=0
	EventTrigger []datacloudfunctions2function.EventTrigger `hcl:"event_trigger,block" validate:"min=0"`
	// ServiceConfig: min=0
	ServiceConfig []datacloudfunctions2function.ServiceConfig `hcl:"service_config,block" validate:"min=0"`
}
type dataCloudfunctions2FunctionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("description"))
}

// Environment returns a reference to field environment of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("environment"))
}

// Id returns a reference to field id of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("id"))
}

// Labels returns a reference to field labels of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("labels"))
}

// Location returns a reference to field location of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("project"))
}

// State returns a reference to field state of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("state"))
}

// UpdateTime returns a reference to field update_time of google_cloudfunctions2_function.
func (cf dataCloudfunctions2FunctionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("update_time"))
}

func (cf dataCloudfunctions2FunctionAttributes) BuildConfig() terra.ListValue[datacloudfunctions2function.BuildConfigAttributes] {
	return terra.ReferenceAsList[datacloudfunctions2function.BuildConfigAttributes](cf.ref.Append("build_config"))
}

func (cf dataCloudfunctions2FunctionAttributes) EventTrigger() terra.ListValue[datacloudfunctions2function.EventTriggerAttributes] {
	return terra.ReferenceAsList[datacloudfunctions2function.EventTriggerAttributes](cf.ref.Append("event_trigger"))
}

func (cf dataCloudfunctions2FunctionAttributes) ServiceConfig() terra.ListValue[datacloudfunctions2function.ServiceConfigAttributes] {
	return terra.ReferenceAsList[datacloudfunctions2function.ServiceConfigAttributes](cf.ref.Append("service_config"))
}
