// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datalosspreventiondeidentifytemplate "github.com/golingon/terraproviders/google/4.72.1/datalosspreventiondeidentifytemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataLossPreventionDeidentifyTemplate creates a new instance of [DataLossPreventionDeidentifyTemplate].
func NewDataLossPreventionDeidentifyTemplate(name string, args DataLossPreventionDeidentifyTemplateArgs) *DataLossPreventionDeidentifyTemplate {
	return &DataLossPreventionDeidentifyTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataLossPreventionDeidentifyTemplate)(nil)

// DataLossPreventionDeidentifyTemplate represents the Terraform resource google_data_loss_prevention_deidentify_template.
type DataLossPreventionDeidentifyTemplate struct {
	Name      string
	Args      DataLossPreventionDeidentifyTemplateArgs
	state     *dataLossPreventionDeidentifyTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataLossPreventionDeidentifyTemplate].
func (dlpdt *DataLossPreventionDeidentifyTemplate) Type() string {
	return "google_data_loss_prevention_deidentify_template"
}

// LocalName returns the local name for [DataLossPreventionDeidentifyTemplate].
func (dlpdt *DataLossPreventionDeidentifyTemplate) LocalName() string {
	return dlpdt.Name
}

// Configuration returns the configuration (args) for [DataLossPreventionDeidentifyTemplate].
func (dlpdt *DataLossPreventionDeidentifyTemplate) Configuration() interface{} {
	return dlpdt.Args
}

// DependOn is used for other resources to depend on [DataLossPreventionDeidentifyTemplate].
func (dlpdt *DataLossPreventionDeidentifyTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(dlpdt)
}

// Dependencies returns the list of resources [DataLossPreventionDeidentifyTemplate] depends_on.
func (dlpdt *DataLossPreventionDeidentifyTemplate) Dependencies() terra.Dependencies {
	return dlpdt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataLossPreventionDeidentifyTemplate].
func (dlpdt *DataLossPreventionDeidentifyTemplate) LifecycleManagement() *terra.Lifecycle {
	return dlpdt.Lifecycle
}

// Attributes returns the attributes for [DataLossPreventionDeidentifyTemplate].
func (dlpdt *DataLossPreventionDeidentifyTemplate) Attributes() dataLossPreventionDeidentifyTemplateAttributes {
	return dataLossPreventionDeidentifyTemplateAttributes{ref: terra.ReferenceResource(dlpdt)}
}

// ImportState imports the given attribute values into [DataLossPreventionDeidentifyTemplate]'s state.
func (dlpdt *DataLossPreventionDeidentifyTemplate) ImportState(av io.Reader) error {
	dlpdt.state = &dataLossPreventionDeidentifyTemplateState{}
	if err := json.NewDecoder(av).Decode(dlpdt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlpdt.Type(), dlpdt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataLossPreventionDeidentifyTemplate] has state.
func (dlpdt *DataLossPreventionDeidentifyTemplate) State() (*dataLossPreventionDeidentifyTemplateState, bool) {
	return dlpdt.state, dlpdt.state != nil
}

// StateMust returns the state for [DataLossPreventionDeidentifyTemplate]. Panics if the state is nil.
func (dlpdt *DataLossPreventionDeidentifyTemplate) StateMust() *dataLossPreventionDeidentifyTemplateState {
	if dlpdt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlpdt.Type(), dlpdt.LocalName()))
	}
	return dlpdt.state
}

// DataLossPreventionDeidentifyTemplateArgs contains the configurations for google_data_loss_prevention_deidentify_template.
type DataLossPreventionDeidentifyTemplateArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// TemplateId: string, optional
	TemplateId terra.StringValue `hcl:"template_id,attr"`
	// DeidentifyConfig: required
	DeidentifyConfig *datalosspreventiondeidentifytemplate.DeidentifyConfig `hcl:"deidentify_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *datalosspreventiondeidentifytemplate.Timeouts `hcl:"timeouts,block"`
}
type dataLossPreventionDeidentifyTemplateAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("create_time"))
}

// Description returns a reference to field description of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("display_name"))
}

// Id returns a reference to field id of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("id"))
}

// Name returns a reference to field name of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("name"))
}

// Parent returns a reference to field parent of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("parent"))
}

// TemplateId returns a reference to field template_id of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("template_id"))
}

// UpdateTime returns a reference to field update_time of google_data_loss_prevention_deidentify_template.
func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(dlpdt.ref.Append("update_time"))
}

func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) DeidentifyConfig() terra.ListValue[datalosspreventiondeidentifytemplate.DeidentifyConfigAttributes] {
	return terra.ReferenceAsList[datalosspreventiondeidentifytemplate.DeidentifyConfigAttributes](dlpdt.ref.Append("deidentify_config"))
}

func (dlpdt dataLossPreventionDeidentifyTemplateAttributes) Timeouts() datalosspreventiondeidentifytemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalosspreventiondeidentifytemplate.TimeoutsAttributes](dlpdt.ref.Append("timeouts"))
}

type dataLossPreventionDeidentifyTemplateState struct {
	CreateTime       string                                                       `json:"create_time"`
	Description      string                                                       `json:"description"`
	DisplayName      string                                                       `json:"display_name"`
	Id               string                                                       `json:"id"`
	Name             string                                                       `json:"name"`
	Parent           string                                                       `json:"parent"`
	TemplateId       string                                                       `json:"template_id"`
	UpdateTime       string                                                       `json:"update_time"`
	DeidentifyConfig []datalosspreventiondeidentifytemplate.DeidentifyConfigState `json:"deidentify_config"`
	Timeouts         *datalosspreventiondeidentifytemplate.TimeoutsState          `json:"timeouts"`
}
