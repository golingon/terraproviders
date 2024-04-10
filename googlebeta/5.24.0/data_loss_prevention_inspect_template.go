// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	datalosspreventioninspecttemplate "github.com/golingon/terraproviders/googlebeta/5.24.0/datalosspreventioninspecttemplate"
	"io"
)

// NewDataLossPreventionInspectTemplate creates a new instance of [DataLossPreventionInspectTemplate].
func NewDataLossPreventionInspectTemplate(name string, args DataLossPreventionInspectTemplateArgs) *DataLossPreventionInspectTemplate {
	return &DataLossPreventionInspectTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataLossPreventionInspectTemplate)(nil)

// DataLossPreventionInspectTemplate represents the Terraform resource google_data_loss_prevention_inspect_template.
type DataLossPreventionInspectTemplate struct {
	Name      string
	Args      DataLossPreventionInspectTemplateArgs
	state     *dataLossPreventionInspectTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataLossPreventionInspectTemplate].
func (dlpit *DataLossPreventionInspectTemplate) Type() string {
	return "google_data_loss_prevention_inspect_template"
}

// LocalName returns the local name for [DataLossPreventionInspectTemplate].
func (dlpit *DataLossPreventionInspectTemplate) LocalName() string {
	return dlpit.Name
}

// Configuration returns the configuration (args) for [DataLossPreventionInspectTemplate].
func (dlpit *DataLossPreventionInspectTemplate) Configuration() interface{} {
	return dlpit.Args
}

// DependOn is used for other resources to depend on [DataLossPreventionInspectTemplate].
func (dlpit *DataLossPreventionInspectTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(dlpit)
}

// Dependencies returns the list of resources [DataLossPreventionInspectTemplate] depends_on.
func (dlpit *DataLossPreventionInspectTemplate) Dependencies() terra.Dependencies {
	return dlpit.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataLossPreventionInspectTemplate].
func (dlpit *DataLossPreventionInspectTemplate) LifecycleManagement() *terra.Lifecycle {
	return dlpit.Lifecycle
}

// Attributes returns the attributes for [DataLossPreventionInspectTemplate].
func (dlpit *DataLossPreventionInspectTemplate) Attributes() dataLossPreventionInspectTemplateAttributes {
	return dataLossPreventionInspectTemplateAttributes{ref: terra.ReferenceResource(dlpit)}
}

// ImportState imports the given attribute values into [DataLossPreventionInspectTemplate]'s state.
func (dlpit *DataLossPreventionInspectTemplate) ImportState(av io.Reader) error {
	dlpit.state = &dataLossPreventionInspectTemplateState{}
	if err := json.NewDecoder(av).Decode(dlpit.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dlpit.Type(), dlpit.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataLossPreventionInspectTemplate] has state.
func (dlpit *DataLossPreventionInspectTemplate) State() (*dataLossPreventionInspectTemplateState, bool) {
	return dlpit.state, dlpit.state != nil
}

// StateMust returns the state for [DataLossPreventionInspectTemplate]. Panics if the state is nil.
func (dlpit *DataLossPreventionInspectTemplate) StateMust() *dataLossPreventionInspectTemplateState {
	if dlpit.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dlpit.Type(), dlpit.LocalName()))
	}
	return dlpit.state
}

// DataLossPreventionInspectTemplateArgs contains the configurations for google_data_loss_prevention_inspect_template.
type DataLossPreventionInspectTemplateArgs struct {
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
	// InspectConfig: optional
	InspectConfig *datalosspreventioninspecttemplate.InspectConfig `hcl:"inspect_config,block"`
	// Timeouts: optional
	Timeouts *datalosspreventioninspecttemplate.Timeouts `hcl:"timeouts,block"`
}
type dataLossPreventionInspectTemplateAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_data_loss_prevention_inspect_template.
func (dlpit dataLossPreventionInspectTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dlpit.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_loss_prevention_inspect_template.
func (dlpit dataLossPreventionInspectTemplateAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dlpit.ref.Append("display_name"))
}

// Id returns a reference to field id of google_data_loss_prevention_inspect_template.
func (dlpit dataLossPreventionInspectTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dlpit.ref.Append("id"))
}

// Name returns a reference to field name of google_data_loss_prevention_inspect_template.
func (dlpit dataLossPreventionInspectTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dlpit.ref.Append("name"))
}

// Parent returns a reference to field parent of google_data_loss_prevention_inspect_template.
func (dlpit dataLossPreventionInspectTemplateAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dlpit.ref.Append("parent"))
}

// TemplateId returns a reference to field template_id of google_data_loss_prevention_inspect_template.
func (dlpit dataLossPreventionInspectTemplateAttributes) TemplateId() terra.StringValue {
	return terra.ReferenceAsString(dlpit.ref.Append("template_id"))
}

func (dlpit dataLossPreventionInspectTemplateAttributes) InspectConfig() terra.ListValue[datalosspreventioninspecttemplate.InspectConfigAttributes] {
	return terra.ReferenceAsList[datalosspreventioninspecttemplate.InspectConfigAttributes](dlpit.ref.Append("inspect_config"))
}

func (dlpit dataLossPreventionInspectTemplateAttributes) Timeouts() datalosspreventioninspecttemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalosspreventioninspecttemplate.TimeoutsAttributes](dlpit.ref.Append("timeouts"))
}

type dataLossPreventionInspectTemplateState struct {
	Description   string                                                 `json:"description"`
	DisplayName   string                                                 `json:"display_name"`
	Id            string                                                 `json:"id"`
	Name          string                                                 `json:"name"`
	Parent        string                                                 `json:"parent"`
	TemplateId    string                                                 `json:"template_id"`
	InspectConfig []datalosspreventioninspecttemplate.InspectConfigState `json:"inspect_config"`
	Timeouts      *datalosspreventioninspecttemplate.TimeoutsState       `json:"timeouts"`
}
