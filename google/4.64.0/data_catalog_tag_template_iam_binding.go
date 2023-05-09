// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datacatalogtagtemplateiambinding "github.com/golingon/terraproviders/google/4.64.0/datacatalogtagtemplateiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogTagTemplateIamBinding creates a new instance of [DataCatalogTagTemplateIamBinding].
func NewDataCatalogTagTemplateIamBinding(name string, args DataCatalogTagTemplateIamBindingArgs) *DataCatalogTagTemplateIamBinding {
	return &DataCatalogTagTemplateIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTagTemplateIamBinding)(nil)

// DataCatalogTagTemplateIamBinding represents the Terraform resource google_data_catalog_tag_template_iam_binding.
type DataCatalogTagTemplateIamBinding struct {
	Name      string
	Args      DataCatalogTagTemplateIamBindingArgs
	state     *dataCatalogTagTemplateIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogTagTemplateIamBinding].
func (dcttib *DataCatalogTagTemplateIamBinding) Type() string {
	return "google_data_catalog_tag_template_iam_binding"
}

// LocalName returns the local name for [DataCatalogTagTemplateIamBinding].
func (dcttib *DataCatalogTagTemplateIamBinding) LocalName() string {
	return dcttib.Name
}

// Configuration returns the configuration (args) for [DataCatalogTagTemplateIamBinding].
func (dcttib *DataCatalogTagTemplateIamBinding) Configuration() interface{} {
	return dcttib.Args
}

// DependOn is used for other resources to depend on [DataCatalogTagTemplateIamBinding].
func (dcttib *DataCatalogTagTemplateIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dcttib)
}

// Dependencies returns the list of resources [DataCatalogTagTemplateIamBinding] depends_on.
func (dcttib *DataCatalogTagTemplateIamBinding) Dependencies() terra.Dependencies {
	return dcttib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogTagTemplateIamBinding].
func (dcttib *DataCatalogTagTemplateIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dcttib.Lifecycle
}

// Attributes returns the attributes for [DataCatalogTagTemplateIamBinding].
func (dcttib *DataCatalogTagTemplateIamBinding) Attributes() dataCatalogTagTemplateIamBindingAttributes {
	return dataCatalogTagTemplateIamBindingAttributes{ref: terra.ReferenceResource(dcttib)}
}

// ImportState imports the given attribute values into [DataCatalogTagTemplateIamBinding]'s state.
func (dcttib *DataCatalogTagTemplateIamBinding) ImportState(av io.Reader) error {
	dcttib.state = &dataCatalogTagTemplateIamBindingState{}
	if err := json.NewDecoder(av).Decode(dcttib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcttib.Type(), dcttib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogTagTemplateIamBinding] has state.
func (dcttib *DataCatalogTagTemplateIamBinding) State() (*dataCatalogTagTemplateIamBindingState, bool) {
	return dcttib.state, dcttib.state != nil
}

// StateMust returns the state for [DataCatalogTagTemplateIamBinding]. Panics if the state is nil.
func (dcttib *DataCatalogTagTemplateIamBinding) StateMust() *dataCatalogTagTemplateIamBindingState {
	if dcttib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcttib.Type(), dcttib.LocalName()))
	}
	return dcttib.state
}

// DataCatalogTagTemplateIamBindingArgs contains the configurations for google_data_catalog_tag_template_iam_binding.
type DataCatalogTagTemplateIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TagTemplate: string, required
	TagTemplate terra.StringValue `hcl:"tag_template,attr" validate:"required"`
	// Condition: optional
	Condition *datacatalogtagtemplateiambinding.Condition `hcl:"condition,block"`
}
type dataCatalogTagTemplateIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcttib.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcttib.ref.Append("id"))
}

// Members returns a reference to field members of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dcttib.ref.Append("members"))
}

// Project returns a reference to field project of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcttib.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcttib.ref.Append("region"))
}

// Role returns a reference to field role of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcttib.ref.Append("role"))
}

// TagTemplate returns a reference to field tag_template of google_data_catalog_tag_template_iam_binding.
func (dcttib dataCatalogTagTemplateIamBindingAttributes) TagTemplate() terra.StringValue {
	return terra.ReferenceAsString(dcttib.ref.Append("tag_template"))
}

func (dcttib dataCatalogTagTemplateIamBindingAttributes) Condition() terra.ListValue[datacatalogtagtemplateiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogtagtemplateiambinding.ConditionAttributes](dcttib.ref.Append("condition"))
}

type dataCatalogTagTemplateIamBindingState struct {
	Etag        string                                            `json:"etag"`
	Id          string                                            `json:"id"`
	Members     []string                                          `json:"members"`
	Project     string                                            `json:"project"`
	Region      string                                            `json:"region"`
	Role        string                                            `json:"role"`
	TagTemplate string                                            `json:"tag_template"`
	Condition   []datacatalogtagtemplateiambinding.ConditionState `json:"condition"`
}
