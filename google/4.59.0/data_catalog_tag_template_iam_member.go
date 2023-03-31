// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datacatalogtagtemplateiammember "github.com/golingon/terraproviders/google/4.59.0/datacatalogtagtemplateiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataCatalogTagTemplateIamMember(name string, args DataCatalogTagTemplateIamMemberArgs) *DataCatalogTagTemplateIamMember {
	return &DataCatalogTagTemplateIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTagTemplateIamMember)(nil)

type DataCatalogTagTemplateIamMember struct {
	Name  string
	Args  DataCatalogTagTemplateIamMemberArgs
	state *dataCatalogTagTemplateIamMemberState
}

func (dcttim *DataCatalogTagTemplateIamMember) Type() string {
	return "google_data_catalog_tag_template_iam_member"
}

func (dcttim *DataCatalogTagTemplateIamMember) LocalName() string {
	return dcttim.Name
}

func (dcttim *DataCatalogTagTemplateIamMember) Configuration() interface{} {
	return dcttim.Args
}

func (dcttim *DataCatalogTagTemplateIamMember) Attributes() dataCatalogTagTemplateIamMemberAttributes {
	return dataCatalogTagTemplateIamMemberAttributes{ref: terra.ReferenceResource(dcttim)}
}

func (dcttim *DataCatalogTagTemplateIamMember) ImportState(av io.Reader) error {
	dcttim.state = &dataCatalogTagTemplateIamMemberState{}
	if err := json.NewDecoder(av).Decode(dcttim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcttim.Type(), dcttim.LocalName(), err)
	}
	return nil
}

func (dcttim *DataCatalogTagTemplateIamMember) State() (*dataCatalogTagTemplateIamMemberState, bool) {
	return dcttim.state, dcttim.state != nil
}

func (dcttim *DataCatalogTagTemplateIamMember) StateMust() *dataCatalogTagTemplateIamMemberState {
	if dcttim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcttim.Type(), dcttim.LocalName()))
	}
	return dcttim.state
}

func (dcttim *DataCatalogTagTemplateIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dcttim)
}

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
	// DependsOn contains resources that DataCatalogTagTemplateIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataCatalogTagTemplateIamMemberAttributes struct {
	ref terra.Reference
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("etag"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("id"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("member"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("project"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("region"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("role"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) TagTemplate() terra.StringValue {
	return terra.ReferenceString(dcttim.ref.Append("tag_template"))
}

func (dcttim dataCatalogTagTemplateIamMemberAttributes) Condition() terra.ListValue[datacatalogtagtemplateiammember.ConditionAttributes] {
	return terra.ReferenceList[datacatalogtagtemplateiammember.ConditionAttributes](dcttim.ref.Append("condition"))
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
