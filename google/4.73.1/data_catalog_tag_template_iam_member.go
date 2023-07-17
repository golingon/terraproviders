// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datacatalogtagtemplateiammember "github.com/golingon/terraproviders/google/4.73.1/datacatalogtagtemplateiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogTagTemplateIamMember creates a new instance of [DataCatalogTagTemplateIamMember].
func NewDataCatalogTagTemplateIamMember(name string, args DataCatalogTagTemplateIamMemberArgs) *DataCatalogTagTemplateIamMember {
	return &DataCatalogTagTemplateIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTagTemplateIamMember)(nil)

// DataCatalogTagTemplateIamMember represents the Terraform resource google_data_catalog_tag_template_iam_member.
type DataCatalogTagTemplateIamMember struct {
	Name      string
	Args      DataCatalogTagTemplateIamMemberArgs
	state     *dataCatalogTagTemplateIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogTagTemplateIamMember].
func (dcttim *DataCatalogTagTemplateIamMember) Type() string {
	return "google_data_catalog_tag_template_iam_member"
}

// LocalName returns the local name for [DataCatalogTagTemplateIamMember].
func (dcttim *DataCatalogTagTemplateIamMember) LocalName() string {
	return dcttim.Name
}

// Configuration returns the configuration (args) for [DataCatalogTagTemplateIamMember].
func (dcttim *DataCatalogTagTemplateIamMember) Configuration() interface{} {
	return dcttim.Args
}

// DependOn is used for other resources to depend on [DataCatalogTagTemplateIamMember].
func (dcttim *DataCatalogTagTemplateIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dcttim)
}

// Dependencies returns the list of resources [DataCatalogTagTemplateIamMember] depends_on.
func (dcttim *DataCatalogTagTemplateIamMember) Dependencies() terra.Dependencies {
	return dcttim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogTagTemplateIamMember].
func (dcttim *DataCatalogTagTemplateIamMember) LifecycleManagement() *terra.Lifecycle {
	return dcttim.Lifecycle
}

// Attributes returns the attributes for [DataCatalogTagTemplateIamMember].
func (dcttim *DataCatalogTagTemplateIamMember) Attributes() dataCatalogTagTemplateIamMemberAttributes {
	return dataCatalogTagTemplateIamMemberAttributes{ref: terra.ReferenceResource(dcttim)}
}

// ImportState imports the given attribute values into [DataCatalogTagTemplateIamMember]'s state.
func (dcttim *DataCatalogTagTemplateIamMember) ImportState(av io.Reader) error {
	dcttim.state = &dataCatalogTagTemplateIamMemberState{}
	if err := json.NewDecoder(av).Decode(dcttim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcttim.Type(), dcttim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogTagTemplateIamMember] has state.
func (dcttim *DataCatalogTagTemplateIamMember) State() (*dataCatalogTagTemplateIamMemberState, bool) {
	return dcttim.state, dcttim.state != nil
}

// StateMust returns the state for [DataCatalogTagTemplateIamMember]. Panics if the state is nil.
func (dcttim *DataCatalogTagTemplateIamMember) StateMust() *dataCatalogTagTemplateIamMemberState {
	if dcttim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcttim.Type(), dcttim.LocalName()))
	}
	return dcttim.state
}

// DataCatalogTagTemplateIamMemberArgs contains the configurations for google_data_catalog_tag_template_iam_member.
type DataCatalogTagTemplateIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// TagTemplate: string, required
	TagTemplate terra.StringValue `hcl:"tag_template,attr" validate:"required"`
	// Condition: optional
	Condition *datacatalogtagtemplateiammember.Condition `hcl:"condition,block"`
}
type dataCatalogTagTemplateIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("id"))
}

// Member returns a reference to field member of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("member"))
}

// Project returns a reference to field project of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("region"))
}

// Role returns a reference to field role of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("role"))
}

// TagTemplate returns a reference to field tag_template of google_data_catalog_tag_template_iam_member.
func (dcttim dataCatalogTagTemplateIamMemberAttributes) TagTemplate() terra.StringValue {
	return terra.ReferenceAsString(dcttim.ref.Append("tag_template"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Condition() terra.ListValue[datacatalogtagtemplateiammember.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogtagtemplateiammember.ConditionAttributes](dcttim.ref.Append("condition"))
}

type dataCatalogTagTemplateIamMemberState struct {
	Etag        string                                           `json:"etag"`
	Id          string                                           `json:"id"`
	Member      string                                           `json:"member"`
	Project     string                                           `json:"project"`
	Region      string                                           `json:"region"`
	Role        string                                           `json:"role"`
	TagTemplate string                                           `json:"tag_template"`
	Condition   []datacatalogtagtemplateiammember.ConditionState `json:"condition"`
}
