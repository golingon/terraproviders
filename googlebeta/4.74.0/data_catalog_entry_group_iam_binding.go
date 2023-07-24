// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogentrygroupiambinding "github.com/golingon/terraproviders/googlebeta/4.74.0/datacatalogentrygroupiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogEntryGroupIamBinding creates a new instance of [DataCatalogEntryGroupIamBinding].
func NewDataCatalogEntryGroupIamBinding(name string, args DataCatalogEntryGroupIamBindingArgs) *DataCatalogEntryGroupIamBinding {
	return &DataCatalogEntryGroupIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogEntryGroupIamBinding)(nil)

// DataCatalogEntryGroupIamBinding represents the Terraform resource google_data_catalog_entry_group_iam_binding.
type DataCatalogEntryGroupIamBinding struct {
	Name      string
	Args      DataCatalogEntryGroupIamBindingArgs
	state     *dataCatalogEntryGroupIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogEntryGroupIamBinding].
func (dcegib *DataCatalogEntryGroupIamBinding) Type() string {
	return "google_data_catalog_entry_group_iam_binding"
}

// LocalName returns the local name for [DataCatalogEntryGroupIamBinding].
func (dcegib *DataCatalogEntryGroupIamBinding) LocalName() string {
	return dcegib.Name
}

// Configuration returns the configuration (args) for [DataCatalogEntryGroupIamBinding].
func (dcegib *DataCatalogEntryGroupIamBinding) Configuration() interface{} {
	return dcegib.Args
}

// DependOn is used for other resources to depend on [DataCatalogEntryGroupIamBinding].
func (dcegib *DataCatalogEntryGroupIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(dcegib)
}

// Dependencies returns the list of resources [DataCatalogEntryGroupIamBinding] depends_on.
func (dcegib *DataCatalogEntryGroupIamBinding) Dependencies() terra.Dependencies {
	return dcegib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogEntryGroupIamBinding].
func (dcegib *DataCatalogEntryGroupIamBinding) LifecycleManagement() *terra.Lifecycle {
	return dcegib.Lifecycle
}

// Attributes returns the attributes for [DataCatalogEntryGroupIamBinding].
func (dcegib *DataCatalogEntryGroupIamBinding) Attributes() dataCatalogEntryGroupIamBindingAttributes {
	return dataCatalogEntryGroupIamBindingAttributes{ref: terra.ReferenceResource(dcegib)}
}

// ImportState imports the given attribute values into [DataCatalogEntryGroupIamBinding]'s state.
func (dcegib *DataCatalogEntryGroupIamBinding) ImportState(av io.Reader) error {
	dcegib.state = &dataCatalogEntryGroupIamBindingState{}
	if err := json.NewDecoder(av).Decode(dcegib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcegib.Type(), dcegib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogEntryGroupIamBinding] has state.
func (dcegib *DataCatalogEntryGroupIamBinding) State() (*dataCatalogEntryGroupIamBindingState, bool) {
	return dcegib.state, dcegib.state != nil
}

// StateMust returns the state for [DataCatalogEntryGroupIamBinding]. Panics if the state is nil.
func (dcegib *DataCatalogEntryGroupIamBinding) StateMust() *dataCatalogEntryGroupIamBindingState {
	if dcegib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcegib.Type(), dcegib.LocalName()))
	}
	return dcegib.state
}

// DataCatalogEntryGroupIamBindingArgs contains the configurations for google_data_catalog_entry_group_iam_binding.
type DataCatalogEntryGroupIamBindingArgs struct {
	// EntryGroup: string, required
	EntryGroup terra.StringValue `hcl:"entry_group,attr" validate:"required"`
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
	// Condition: optional
	Condition *datacatalogentrygroupiambinding.Condition `hcl:"condition,block"`
}
type dataCatalogEntryGroupIamBindingAttributes struct {
	ref terra.Reference
}

// EntryGroup returns a reference to field entry_group of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) EntryGroup() terra.StringValue {
	return terra.ReferenceAsString(dcegib.ref.Append("entry_group"))
}

// Etag returns a reference to field etag of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dcegib.ref.Append("etag"))
}

// Id returns a reference to field id of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcegib.ref.Append("id"))
}

// Members returns a reference to field members of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dcegib.ref.Append("members"))
}

// Project returns a reference to field project of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcegib.ref.Append("project"))
}

// Region returns a reference to field region of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(dcegib.ref.Append("region"))
}

// Role returns a reference to field role of google_data_catalog_entry_group_iam_binding.
func (dcegib dataCatalogEntryGroupIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(dcegib.ref.Append("role"))
}

func (dcegib dataCatalogEntryGroupIamBindingAttributes) Condition() terra.ListValue[datacatalogentrygroupiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[datacatalogentrygroupiambinding.ConditionAttributes](dcegib.ref.Append("condition"))
}

type dataCatalogEntryGroupIamBindingState struct {
	EntryGroup string                                           `json:"entry_group"`
	Etag       string                                           `json:"etag"`
	Id         string                                           `json:"id"`
	Members    []string                                         `json:"members"`
	Project    string                                           `json:"project"`
	Region     string                                           `json:"region"`
	Role       string                                           `json:"role"`
	Condition  []datacatalogentrygroupiambinding.ConditionState `json:"condition"`
}
