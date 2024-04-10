// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"github.com/golingon/lingon/pkg/terra"
	dataspannerinstance "github.com/golingon/terraproviders/googlebeta/5.24.0/dataspannerinstance"
)

// NewDataSpannerInstance creates a new instance of [DataSpannerInstance].
func NewDataSpannerInstance(name string, args DataSpannerInstanceArgs) *DataSpannerInstance {
	return &DataSpannerInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSpannerInstance)(nil)

// DataSpannerInstance represents the Terraform data resource google_spanner_instance.
type DataSpannerInstance struct {
	Name string
	Args DataSpannerInstanceArgs
}

// DataSource returns the Terraform object type for [DataSpannerInstance].
func (si *DataSpannerInstance) DataSource() string {
	return "google_spanner_instance"
}

// LocalName returns the local name for [DataSpannerInstance].
func (si *DataSpannerInstance) LocalName() string {
	return si.Name
}

// Configuration returns the configuration (args) for [DataSpannerInstance].
func (si *DataSpannerInstance) Configuration() interface{} {
	return si.Args
}

// Attributes returns the attributes for [DataSpannerInstance].
func (si *DataSpannerInstance) Attributes() dataSpannerInstanceAttributes {
	return dataSpannerInstanceAttributes{ref: terra.ReferenceDataResource(si)}
}

// DataSpannerInstanceArgs contains the configurations for google_spanner_instance.
type DataSpannerInstanceArgs struct {
	// Config: string, optional
	Config terra.StringValue `hcl:"config,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AutoscalingConfig: min=0
	AutoscalingConfig []dataspannerinstance.AutoscalingConfig `hcl:"autoscaling_config,block" validate:"min=0"`
}
type dataSpannerInstanceAttributes struct {
	ref terra.Reference
}

// Config returns a reference to field config of google_spanner_instance.
func (si dataSpannerInstanceAttributes) Config() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("config"))
}

// DisplayName returns a reference to field display_name of google_spanner_instance.
func (si dataSpannerInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("display_name"))
}

// EffectiveLabels returns a reference to field effective_labels of google_spanner_instance.
func (si dataSpannerInstanceAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](si.ref.Append("effective_labels"))
}

// ForceDestroy returns a reference to field force_destroy of google_spanner_instance.
func (si dataSpannerInstanceAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(si.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_spanner_instance.
func (si dataSpannerInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("id"))
}

// Labels returns a reference to field labels of google_spanner_instance.
func (si dataSpannerInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](si.ref.Append("labels"))
}

// Name returns a reference to field name of google_spanner_instance.
func (si dataSpannerInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("name"))
}

// NumNodes returns a reference to field num_nodes of google_spanner_instance.
func (si dataSpannerInstanceAttributes) NumNodes() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("num_nodes"))
}

// ProcessingUnits returns a reference to field processing_units of google_spanner_instance.
func (si dataSpannerInstanceAttributes) ProcessingUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(si.ref.Append("processing_units"))
}

// Project returns a reference to field project of google_spanner_instance.
func (si dataSpannerInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("project"))
}

// State returns a reference to field state of google_spanner_instance.
func (si dataSpannerInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("state"))
}

// TerraformLabels returns a reference to field terraform_labels of google_spanner_instance.
func (si dataSpannerInstanceAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](si.ref.Append("terraform_labels"))
}

func (si dataSpannerInstanceAttributes) AutoscalingConfig() terra.ListValue[dataspannerinstance.AutoscalingConfigAttributes] {
	return terra.ReferenceAsList[dataspannerinstance.AutoscalingConfigAttributes](si.ref.Append("autoscaling_config"))
}
