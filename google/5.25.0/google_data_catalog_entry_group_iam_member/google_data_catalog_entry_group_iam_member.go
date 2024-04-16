// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_data_catalog_entry_group_iam_member

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_data_catalog_entry_group_iam_member.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDataCatalogEntryGroupIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdcegim *Resource) Type() string {
	return "google_data_catalog_entry_group_iam_member"
}

// LocalName returns the local name for [Resource].
func (gdcegim *Resource) LocalName() string {
	return gdcegim.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdcegim *Resource) Configuration() interface{} {
	return gdcegim.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdcegim *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdcegim)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdcegim *Resource) Dependencies() terra.Dependencies {
	return gdcegim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdcegim *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdcegim.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdcegim *Resource) Attributes() googleDataCatalogEntryGroupIamMemberAttributes {
	return googleDataCatalogEntryGroupIamMemberAttributes{ref: terra.ReferenceResource(gdcegim)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdcegim *Resource) ImportState(state io.Reader) error {
	gdcegim.state = &googleDataCatalogEntryGroupIamMemberState{}
	if err := json.NewDecoder(state).Decode(gdcegim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdcegim.Type(), gdcegim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdcegim *Resource) State() (*googleDataCatalogEntryGroupIamMemberState, bool) {
	return gdcegim.state, gdcegim.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdcegim *Resource) StateMust() *googleDataCatalogEntryGroupIamMemberState {
	if gdcegim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdcegim.Type(), gdcegim.LocalName()))
	}
	return gdcegim.state
}

// Args contains the configurations for google_data_catalog_entry_group_iam_member.
type Args struct {
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
	Condition *Condition `hcl:"condition,block"`
}

type googleDataCatalogEntryGroupIamMemberAttributes struct {
	ref terra.Reference
}

// EntryGroup returns a reference to field entry_group of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) EntryGroup() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("entry_group"))
}

// Etag returns a reference to field etag of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("id"))
}

// Member returns a reference to field member of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("member"))
}

// Project returns a reference to field project of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("region"))
}

// Role returns a reference to field role of google_data_catalog_entry_group_iam_member.
func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gdcegim.ref.Append("role"))
}

func (gdcegim googleDataCatalogEntryGroupIamMemberAttributes) Condition() terra.ListValue[ConditionAttributes] {
	return terra.ReferenceAsList[ConditionAttributes](gdcegim.ref.Append("condition"))
}

type googleDataCatalogEntryGroupIamMemberState struct {
	EntryGroup string           `json:"entry_group"`
	Etag       string           `json:"etag"`
	Id         string           `json:"id"`
	Member     string           `json:"member"`
	Project    string           `json:"project"`
	Region     string           `json:"region"`
	Role       string           `json:"role"`
	Condition  []ConditionState `json:"condition"`
}
