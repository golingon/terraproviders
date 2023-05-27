// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datacatalogtagtemplate "github.com/golingon/terraproviders/google/4.66.0/datacatalogtagtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogTagTemplate creates a new instance of [DataCatalogTagTemplate].
func NewDataCatalogTagTemplate(name string, args DataCatalogTagTemplateArgs) *DataCatalogTagTemplate {
	return &DataCatalogTagTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTagTemplate)(nil)

// DataCatalogTagTemplate represents the Terraform resource google_data_catalog_tag_template.
type DataCatalogTagTemplate struct {
	Name      string
	Args      DataCatalogTagTemplateArgs
	state     *dataCatalogTagTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogTagTemplate].
func (dctt *DataCatalogTagTemplate) Type() string {
	return "google_data_catalog_tag_template"
}

// LocalName returns the local name for [DataCatalogTagTemplate].
func (dctt *DataCatalogTagTemplate) LocalName() string {
	return dctt.Name
}

// Configuration returns the configuration (args) for [DataCatalogTagTemplate].
func (dctt *DataCatalogTagTemplate) Configuration() interface{} {
	return dctt.Args
}

// DependOn is used for other resources to depend on [DataCatalogTagTemplate].
func (dctt *DataCatalogTagTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(dctt)
}

// Dependencies returns the list of resources [DataCatalogTagTemplate] depends_on.
func (dctt *DataCatalogTagTemplate) Dependencies() terra.Dependencies {
	return dctt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogTagTemplate].
func (dctt *DataCatalogTagTemplate) LifecycleManagement() *terra.Lifecycle {
	return dctt.Lifecycle
}

// Attributes returns the attributes for [DataCatalogTagTemplate].
func (dctt *DataCatalogTagTemplate) Attributes() dataCatalogTagTemplateAttributes {
	return dataCatalogTagTemplateAttributes{ref: terra.ReferenceResource(dctt)}
}

// ImportState imports the given attribute values into [DataCatalogTagTemplate]'s state.
func (dctt *DataCatalogTagTemplate) ImportState(av io.Reader) error {
	dctt.state = &dataCatalogTagTemplateState{}
	if err := json.NewDecoder(av).Decode(dctt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dctt.Type(), dctt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogTagTemplate] has state.
func (dctt *DataCatalogTagTemplate) State() (*dataCatalogTagTemplateState, bool) {
	return dctt.state, dctt.state != nil
}

// StateMust returns the state for [DataCatalogTagTemplate]. Panics if the state is nil.
func (dctt *DataCatalogTagTemplate) StateMust() *dataCatalogTagTemplateState {
	if dctt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dctt.Type(), dctt.LocalName()))
	}
	return dctt.state
}

// DataCatalogTagTemplateArgs contains the configurations for google_data_catalog_tag_template.
type DataCatalogTagTemplateArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// ForceDelete: bool, optional
	ForceDelete terra.BoolValue `hcl:"force_delete,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// TagTemplateId: string, required
	TagTemplateId terra.StringValue `hcl:"tag_template_id,attr" validate:"required"`
	// Fields: min=1
	Fields []datacatalogtagtemplate.Fields `hcl:"fields,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *datacatalogtagtemplate.Timeouts `hcl:"timeouts,block"`
}
type dataCatalogTagTemplateAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dctt.ref.Append("display_name"))
}

// ForceDelete returns a reference to field force_delete of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) ForceDelete() terra.BoolValue {
	return terra.ReferenceAsBool(dctt.ref.Append("force_delete"))
}

// Id returns a reference to field id of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dctt.ref.Append("id"))
}

// Name returns a reference to field name of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dctt.ref.Append("name"))
}

// Project returns a reference to field project of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dctt.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dctt.ref.Append("region"))
}

// TagTemplateId returns a reference to field tag_template_id of google_data_catalog_tag_template.
func (dctt dataCatalogTagTemplateAttributes) TagTemplateId() terra.StringValue {
	return terra.ReferenceAsString(dctt.ref.Append("tag_template_id"))
}

func (dctt dataCatalogTagTemplateAttributes) Fields() terra.SetValue[datacatalogtagtemplate.FieldsAttributes] {
	return terra.ReferenceAsSet[datacatalogtagtemplate.FieldsAttributes](dctt.ref.Append("fields"))
}

func (dctt dataCatalogTagTemplateAttributes) Timeouts() datacatalogtagtemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacatalogtagtemplate.TimeoutsAttributes](dctt.ref.Append("timeouts"))
}

type dataCatalogTagTemplateState struct {
	DisplayName   string                                `json:"display_name"`
	ForceDelete   bool                                  `json:"force_delete"`
	Id            string                                `json:"id"`
	Name          string                                `json:"name"`
	Project       string                                `json:"project"`
	Region        string                                `json:"region"`
	TagTemplateId string                                `json:"tag_template_id"`
	Fields        []datacatalogtagtemplate.FieldsState  `json:"fields"`
	Timeouts      *datacatalogtagtemplate.TimeoutsState `json:"timeouts"`
}
