// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datasesv2configurationset "github.com/golingon/terraproviders/aws/5.6.2/datasesv2configurationset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSesv2ConfigurationSet creates a new instance of [DataSesv2ConfigurationSet].
func NewDataSesv2ConfigurationSet(name string, args DataSesv2ConfigurationSetArgs) *DataSesv2ConfigurationSet {
	return &DataSesv2ConfigurationSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSesv2ConfigurationSet)(nil)

// DataSesv2ConfigurationSet represents the Terraform data resource aws_sesv2_configuration_set.
type DataSesv2ConfigurationSet struct {
	Name string
	Args DataSesv2ConfigurationSetArgs
}

// DataSource returns the Terraform object type for [DataSesv2ConfigurationSet].
func (scs *DataSesv2ConfigurationSet) DataSource() string {
	return "aws_sesv2_configuration_set"
}

// LocalName returns the local name for [DataSesv2ConfigurationSet].
func (scs *DataSesv2ConfigurationSet) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [DataSesv2ConfigurationSet].
func (scs *DataSesv2ConfigurationSet) Configuration() interface{} {
	return scs.Args
}

// Attributes returns the attributes for [DataSesv2ConfigurationSet].
func (scs *DataSesv2ConfigurationSet) Attributes() dataSesv2ConfigurationSetAttributes {
	return dataSesv2ConfigurationSetAttributes{ref: terra.ReferenceDataResource(scs)}
}

// DataSesv2ConfigurationSetArgs contains the configurations for aws_sesv2_configuration_set.
type DataSesv2ConfigurationSetArgs struct {
	// ConfigurationSetName: string, required
	ConfigurationSetName terra.StringValue `hcl:"configuration_set_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DeliveryOptions: min=0
	DeliveryOptions []datasesv2configurationset.DeliveryOptions `hcl:"delivery_options,block" validate:"min=0"`
	// ReputationOptions: min=0
	ReputationOptions []datasesv2configurationset.ReputationOptions `hcl:"reputation_options,block" validate:"min=0"`
	// SendingOptions: min=0
	SendingOptions []datasesv2configurationset.SendingOptions `hcl:"sending_options,block" validate:"min=0"`
	// SuppressionOptions: min=0
	SuppressionOptions []datasesv2configurationset.SuppressionOptions `hcl:"suppression_options,block" validate:"min=0"`
	// TrackingOptions: min=0
	TrackingOptions []datasesv2configurationset.TrackingOptions `hcl:"tracking_options,block" validate:"min=0"`
	// VdmOptions: min=0
	VdmOptions []datasesv2configurationset.VdmOptions `hcl:"vdm_options,block" validate:"min=0"`
}
type dataSesv2ConfigurationSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sesv2_configuration_set.
func (scs dataSesv2ConfigurationSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("arn"))
}

// ConfigurationSetName returns a reference to field configuration_set_name of aws_sesv2_configuration_set.
func (scs dataSesv2ConfigurationSetAttributes) ConfigurationSetName() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("configuration_set_name"))
}

// Id returns a reference to field id of aws_sesv2_configuration_set.
func (scs dataSesv2ConfigurationSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_sesv2_configuration_set.
func (scs dataSesv2ConfigurationSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("tags"))
}

func (scs dataSesv2ConfigurationSetAttributes) DeliveryOptions() terra.ListValue[datasesv2configurationset.DeliveryOptionsAttributes] {
	return terra.ReferenceAsList[datasesv2configurationset.DeliveryOptionsAttributes](scs.ref.Append("delivery_options"))
}

func (scs dataSesv2ConfigurationSetAttributes) ReputationOptions() terra.ListValue[datasesv2configurationset.ReputationOptionsAttributes] {
	return terra.ReferenceAsList[datasesv2configurationset.ReputationOptionsAttributes](scs.ref.Append("reputation_options"))
}

func (scs dataSesv2ConfigurationSetAttributes) SendingOptions() terra.ListValue[datasesv2configurationset.SendingOptionsAttributes] {
	return terra.ReferenceAsList[datasesv2configurationset.SendingOptionsAttributes](scs.ref.Append("sending_options"))
}

func (scs dataSesv2ConfigurationSetAttributes) SuppressionOptions() terra.ListValue[datasesv2configurationset.SuppressionOptionsAttributes] {
	return terra.ReferenceAsList[datasesv2configurationset.SuppressionOptionsAttributes](scs.ref.Append("suppression_options"))
}

func (scs dataSesv2ConfigurationSetAttributes) TrackingOptions() terra.ListValue[datasesv2configurationset.TrackingOptionsAttributes] {
	return terra.ReferenceAsList[datasesv2configurationset.TrackingOptionsAttributes](scs.ref.Append("tracking_options"))
}

func (scs dataSesv2ConfigurationSetAttributes) VdmOptions() terra.ListValue[datasesv2configurationset.VdmOptionsAttributes] {
	return terra.ReferenceAsList[datasesv2configurationset.VdmOptionsAttributes](scs.ref.Append("vdm_options"))
}
