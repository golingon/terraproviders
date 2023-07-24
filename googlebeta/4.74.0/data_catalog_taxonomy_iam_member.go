// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogtaxonomyiammember "github.com/golingon/terraproviders/googlebeta/4.74.0/datacatalogtaxonomyiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogTaxonomyIamMember creates a new instance of [DataCatalogTaxonomyIamMember].
func NewDataCatalogTaxonomyIamMember(name string, args DataCatalogTaxonomyIamMemberArgs) *DataCatalogTaxonomyIamMember {
	return &DataCatalogTaxonomyIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogTaxonomyIamMember)(nil)

// DataCatalogTaxonomyIamMember represents the Terraform resource google_data_catalog_taxonomy_iam_member.
type DataCatalogTaxonomyIamMember struct {
	Name      string
	Args      DataCatalogTaxonomyIamMemberArgs
	state     *dataCatalogTaxonomyIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogTaxonomyIamMember].
func (dctim *DataCatalogTaxonomyIamMember) Type() string {
	return "google_data_catalog_taxonomy_iam_member"
}

// LocalName returns the local name for [DataCatalogTaxonomyIamMember].
func (dctim *DataCatalogTaxonomyIamMember) LocalName() string {
	return dctim.Name
}

// Configuration returns the configuration (args) for [DataCatalogTaxonomyIamMember].
func (dctim *DataCatalogTaxonomyIamMember) Configuration() interface{} {
	return dctim.Args
}

// DependOn is used for other resources to depend on [DataCatalogTaxonomyIamMember].
func (dctim *DataCatalogTaxonomyIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dctim)
}

// Dependencies returns the list of resources [DataCatalogTaxonomyIamMember] depends_on.
func (dctim *DataCatalogTaxonomyIamMember) Dependencies() terra.Dependencies {
	return dctim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogTaxonomyIamMember].
func (dctim *DataCatalogTaxonomyIamMember) LifecycleManagement() *terra.Lifecycle {
	return dctim.Lifecycle
}

// Attributes returns the attributes for [DataCatalogTaxonomyIamMember].
func (dctim *DataCatalogTaxonomyIamMember) Attributes() dataCatalogTaxonomyIamMemberAttributes {
	return dataCatalogTaxonomyIamMemberAttributes{ref: terra.ReferenceResource(dctim)}
}

// ImportState imports the given attribute values into [DataCatalogTaxonomyIamMember]'s state.
func (dctim *DataCatalogTaxonomyIamMember) ImportState(av io.Reader) error {
	dctim.state = &dataCatalogTaxonomyIamMemberState{}
	if err := json.NewDecoder(av).Decode(dctim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dctim.Type(), dctim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogTaxonomyIamMember] has state.
func (dctim *DataCatalogTaxonomyIamMember) State() (*dataCatalogTaxonomyIamMemberState, bool) {
	return dctim.state, dctim.state != nil
}

// StateMust returns the state for [DataCatalogTaxonomyIamMember]. Panics if the state is nil.
func (dctim *DataCatalogTaxonomyIamMember) StateMust() *dataCatalogTaxonomyIamMemberState {
	if dctim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dctim.Type(), dctim.LocalName()))
	}
	return dctim.state
}

// DataCatalogTaxonomyIamMemberArgs contains the configurations for google_data_catalog_taxonomy_iam_member.
type DataCatalogTaxonomyIamMemberArgs struct {
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
	// Taxonomy: string, required
	Taxonomy terra.StringValue `hcl:"taxonomy,attr" validate:"required"`
	// Condition: optional
	Condition *datacatalogtaxonomyiammember.Condition `hcl:"condition,block"`
}
type dataCatalogTaxonomyIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("id"))
}

// Member returns a reference to field member of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("member"))
}

// Project returns a reference to field project of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("region"))
}

// Role returns a reference to field role of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("role"))
}

// Taxonomy returns a reference to field taxonomy of google_data_catalog_taxonomy_iam_member.
func (dctim dataCatalogTaxonomyIamMemberAttributes) Taxonomy() terra.StringValue {
	return terra.ReferenceAsString(dctim.ref.Append("taxonomy"))
}

func (dctim dataCatalogTaxonomyIamMemberAttributes) Condition() terra.ListValue[datacatalogtaxonomyiammember.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogtaxonomyiammember.ConditionAttributes](dctim.ref.Append("condition"))
}

type dataCatalogTaxonomyIamMemberState struct {
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Region    string                                        `json:"region"`
	Role      string                                        `json:"role"`
	Taxonomy  string                                        `json:"taxonomy"`
	Condition []datacatalogtaxonomyiammember.ConditionState `json:"condition"`
}
