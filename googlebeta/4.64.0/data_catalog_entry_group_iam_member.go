// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogentrygroupiammember "github.com/golingon/terraproviders/googlebeta/4.64.0/datacatalogentrygroupiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogEntryGroupIamMember creates a new instance of [DataCatalogEntryGroupIamMember].
func NewDataCatalogEntryGroupIamMember(name string, args DataCatalogEntryGroupIamMemberArgs) *DataCatalogEntryGroupIamMember {
	return &DataCatalogEntryGroupIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogEntryGroupIamMember)(nil)

// DataCatalogEntryGroupIamMember represents the Terraform resource google_data_catalog_entry_group_iam_member.
type DataCatalogEntryGroupIamMember struct {
	Name      string
	Args      DataCatalogEntryGroupIamMemberArgs
	state     *dataCatalogEntryGroupIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogEntryGroupIamMember].
func (dcegim *DataCatalogEntryGroupIamMember) Type() string {
	return "google_data_catalog_entry_group_iam_member"
}

// LocalName returns the local name for [DataCatalogEntryGroupIamMember].
func (dcegim *DataCatalogEntryGroupIamMember) LocalName() string {
	return dcegim.Name
}

// Configuration returns the configuration (args) for [DataCatalogEntryGroupIamMember].
func (dcegim *DataCatalogEntryGroupIamMember) Configuration() interface{} {
	return dcegim.Args
}

// DependOn is used for other resources to depend on [DataCatalogEntryGroupIamMember].
func (dcegim *DataCatalogEntryGroupIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(dcegim)
}

// Dependencies returns the list of resources [DataCatalogEntryGroupIamMember] depends_on.
func (dcegim *DataCatalogEntryGroupIamMember) Dependencies() terra.Dependencies {
	return dcegim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogEntryGroupIamMember].
func (dcegim *DataCatalogEntryGroupIamMember) LifecycleManagement() *terra.Lifecycle {
	return dcegim.Lifecycle
}

// Attributes returns the attributes for [DataCatalogEntryGroupIamMember].
func (dcegim *DataCatalogEntryGroupIamMember) Attributes() dataCatalogEntryGroupIamMemberAttributes {
	return dataCatalogEntryGroupIamMemberAttributes{ref: terra.ReferenceResource(dcegim)}
}

// ImportState imports the given attribute values into [DataCatalogEntryGroupIamMember]'s state.
func (dcegim *DataCatalogEntryGroupIamMember) ImportState(av io.Reader) error {
	dcegim.state = &dataCatalogEntryGroupIamMemberState{}
	if err := json.NewDecoder(av).Decode(dcegim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcegim.Type(), dcegim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogEntryGroupIamMember] has state.
func (dcegim *DataCatalogEntryGroupIamMember) State() (*dataCatalogEntryGroupIamMemberState, bool) {
	return dcegim.state, dcegim.state != nil
}

// StateMust returns the state for [DataCatalogEntryGroupIamMember]. Panics if the state is nil.
func (dcegim *DataCatalogEntryGroupIamMember) StateMust() *dataCatalogEntryGroupIamMemberState {
	if dcegim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcegim.Type(), dcegim.LocalName()))
	}
	return dcegim.state
}

// DataCatalogEntryGroupIamMemberArgs contains the configurations for google_data_catalog_entry_group_iam_member.
type DataCatalogEntryGroupIamMemberArgs struct {
	// EntryGroup: string, required
	EntryGroup terra.StringValue `hcl:"entry_group,attr" validate:"required"`
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
	// Condition: optional
	Condition *datacatalogentrygroupiammember.Condition `hcl:"condition,block"`
}
type dataCatalogEntryGroupIamMemberAttributes struct {
	ref terra.Reference
}

// EntryGroup returns a reference to field entry_group of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) EntryGroup() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("entry_group"))
}

// Etag returns a reference to field etag of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("id"))
}

// Member returns a reference to field member of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("member"))
}

// Project returns a reference to field project of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("region"))
}

// Role returns a reference to field role of google_data_catalog_entry_group_iam_member.
func (dcegim dataCatalogEntryGroupIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcegim.ref.Append("role"))
}

func (dcegim dataCatalogEntryGroupIamMemberAttributes) Condition() terra.ListValue[datacatalogentrygroupiammember.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogentrygroupiammember.ConditionAttributes](dcegim.ref.Append("condition"))
}

type dataCatalogEntryGroupIamMemberState struct {
	EntryGroup string                                          `json:"entry_group"`
	Etag       string                                          `json:"etag"`
	Id         string                                          `json:"id"`
	Member     string                                          `json:"member"`
	Project    string                                          `json:"project"`
	Region     string                                          `json:"region"`
	Role       string                                          `json:"role"`
	Condition  []datacatalogentrygroupiammember.ConditionState `json:"condition"`
}
