// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datacatalogpolicytag "github.com/golingon/terraproviders/googlebeta/4.72.1/datacatalogpolicytag"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataCatalogPolicyTag creates a new instance of [DataCatalogPolicyTag].
func NewDataCatalogPolicyTag(name string, args DataCatalogPolicyTagArgs) *DataCatalogPolicyTag {
	return &DataCatalogPolicyTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataCatalogPolicyTag)(nil)

// DataCatalogPolicyTag represents the Terraform resource google_data_catalog_policy_tag.
type DataCatalogPolicyTag struct {
	Name      string
	Args      DataCatalogPolicyTagArgs
	state     *dataCatalogPolicyTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataCatalogPolicyTag].
func (dcpt *DataCatalogPolicyTag) Type() string {
	return "google_data_catalog_policy_tag"
}

// LocalName returns the local name for [DataCatalogPolicyTag].
func (dcpt *DataCatalogPolicyTag) LocalName() string {
	return dcpt.Name
}

// Configuration returns the configuration (args) for [DataCatalogPolicyTag].
func (dcpt *DataCatalogPolicyTag) Configuration() interface{} {
	return dcpt.Args
}

// DependOn is used for other resources to depend on [DataCatalogPolicyTag].
func (dcpt *DataCatalogPolicyTag) DependOn() terra.Reference {
	return terra.ReferenceResource(dcpt)
}

// Dependencies returns the list of resources [DataCatalogPolicyTag] depends_on.
func (dcpt *DataCatalogPolicyTag) Dependencies() terra.Dependencies {
	return dcpt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataCatalogPolicyTag].
func (dcpt *DataCatalogPolicyTag) LifecycleManagement() *terra.Lifecycle {
	return dcpt.Lifecycle
}

// Attributes returns the attributes for [DataCatalogPolicyTag].
func (dcpt *DataCatalogPolicyTag) Attributes() dataCatalogPolicyTagAttributes {
	return dataCatalogPolicyTagAttributes{ref: terra.ReferenceResource(dcpt)}
}

// ImportState imports the given attribute values into [DataCatalogPolicyTag]'s state.
func (dcpt *DataCatalogPolicyTag) ImportState(av io.Reader) error {
	dcpt.state = &dataCatalogPolicyTagState{}
	if err := json.NewDecoder(av).Decode(dcpt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcpt.Type(), dcpt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataCatalogPolicyTag] has state.
func (dcpt *DataCatalogPolicyTag) State() (*dataCatalogPolicyTagState, bool) {
	return dcpt.state, dcpt.state != nil
}

// StateMust returns the state for [DataCatalogPolicyTag]. Panics if the state is nil.
func (dcpt *DataCatalogPolicyTag) StateMust() *dataCatalogPolicyTagState {
	if dcpt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcpt.Type(), dcpt.LocalName()))
	}
	return dcpt.state
}

// DataCatalogPolicyTagArgs contains the configurations for google_data_catalog_policy_tag.
type DataCatalogPolicyTagArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ParentPolicyTag: string, optional
	ParentPolicyTag terra.StringValue `hcl:"parent_policy_tag,attr"`
	// Taxonomy: string, required
	Taxonomy terra.StringValue `hcl:"taxonomy,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacatalogpolicytag.Timeouts `hcl:"timeouts,block"`
}
type dataCatalogPolicyTagAttributes struct {
	ref terra.Reference
}

// ChildPolicyTags returns a reference to field child_policy_tags of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) ChildPolicyTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcpt.ref.Append("child_policy_tags"))
}

// Description returns a reference to field description of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dcpt.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcpt.ref.Append("display_name"))
}

// Id returns a reference to field id of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcpt.ref.Append("id"))
}

// Name returns a reference to field name of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcpt.ref.Append("name"))
}

// ParentPolicyTag returns a reference to field parent_policy_tag of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) ParentPolicyTag() terra.StringValue {
	return terra.ReferenceAsString(dcpt.ref.Append("parent_policy_tag"))
}

// Taxonomy returns a reference to field taxonomy of google_data_catalog_policy_tag.
func (dcpt dataCatalogPolicyTagAttributes) Taxonomy() terra.StringValue {
	return terra.ReferenceAsString(dcpt.ref.Append("taxonomy"))
}

func (dcpt dataCatalogPolicyTagAttributes) Timeouts() datacatalogpolicytag.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacatalogpolicytag.TimeoutsAttributes](dcpt.ref.Append("timeouts"))
}

type dataCatalogPolicyTagState struct {
	ChildPolicyTags []string                            `json:"child_policy_tags"`
	Description     string                              `json:"description"`
	DisplayName     string                              `json:"display_name"`
	Id              string                              `json:"id"`
	Name            string                              `json:"name"`
	ParentPolicyTag string                              `json:"parent_policy_tag"`
	Taxonomy        string                              `json:"taxonomy"`
	Timeouts        *datacatalogpolicytag.TimeoutsState `json:"timeouts"`
}
