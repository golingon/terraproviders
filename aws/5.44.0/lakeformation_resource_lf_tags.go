// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	lakeformationresourcelftags "github.com/golingon/terraproviders/aws/5.44.0/lakeformationresourcelftags"
	"io"
)

// NewLakeformationResourceLfTags creates a new instance of [LakeformationResourceLfTags].
func NewLakeformationResourceLfTags(name string, args LakeformationResourceLfTagsArgs) *LakeformationResourceLfTags {
	return &LakeformationResourceLfTags{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LakeformationResourceLfTags)(nil)

// LakeformationResourceLfTags represents the Terraform resource aws_lakeformation_resource_lf_tags.
type LakeformationResourceLfTags struct {
	Name      string
	Args      LakeformationResourceLfTagsArgs
	state     *lakeformationResourceLfTagsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LakeformationResourceLfTags].
func (lrlt *LakeformationResourceLfTags) Type() string {
	return "aws_lakeformation_resource_lf_tags"
}

// LocalName returns the local name for [LakeformationResourceLfTags].
func (lrlt *LakeformationResourceLfTags) LocalName() string {
	return lrlt.Name
}

// Configuration returns the configuration (args) for [LakeformationResourceLfTags].
func (lrlt *LakeformationResourceLfTags) Configuration() interface{} {
	return lrlt.Args
}

// DependOn is used for other resources to depend on [LakeformationResourceLfTags].
func (lrlt *LakeformationResourceLfTags) DependOn() terra.Reference {
	return terra.ReferenceResource(lrlt)
}

// Dependencies returns the list of resources [LakeformationResourceLfTags] depends_on.
func (lrlt *LakeformationResourceLfTags) Dependencies() terra.Dependencies {
	return lrlt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LakeformationResourceLfTags].
func (lrlt *LakeformationResourceLfTags) LifecycleManagement() *terra.Lifecycle {
	return lrlt.Lifecycle
}

// Attributes returns the attributes for [LakeformationResourceLfTags].
func (lrlt *LakeformationResourceLfTags) Attributes() lakeformationResourceLfTagsAttributes {
	return lakeformationResourceLfTagsAttributes{ref: terra.ReferenceResource(lrlt)}
}

// ImportState imports the given attribute values into [LakeformationResourceLfTags]'s state.
func (lrlt *LakeformationResourceLfTags) ImportState(av io.Reader) error {
	lrlt.state = &lakeformationResourceLfTagsState{}
	if err := json.NewDecoder(av).Decode(lrlt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lrlt.Type(), lrlt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LakeformationResourceLfTags] has state.
func (lrlt *LakeformationResourceLfTags) State() (*lakeformationResourceLfTagsState, bool) {
	return lrlt.state, lrlt.state != nil
}

// StateMust returns the state for [LakeformationResourceLfTags]. Panics if the state is nil.
func (lrlt *LakeformationResourceLfTags) StateMust() *lakeformationResourceLfTagsState {
	if lrlt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lrlt.Type(), lrlt.LocalName()))
	}
	return lrlt.state
}

// LakeformationResourceLfTagsArgs contains the configurations for aws_lakeformation_resource_lf_tags.
type LakeformationResourceLfTagsArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Database: optional
	Database *lakeformationresourcelftags.Database `hcl:"database,block"`
	// LfTag: min=1
	LfTag []lakeformationresourcelftags.LfTag `hcl:"lf_tag,block" validate:"min=1"`
	// Table: optional
	Table *lakeformationresourcelftags.Table `hcl:"table,block"`
	// TableWithColumns: optional
	TableWithColumns *lakeformationresourcelftags.TableWithColumns `hcl:"table_with_columns,block"`
	// Timeouts: optional
	Timeouts *lakeformationresourcelftags.Timeouts `hcl:"timeouts,block"`
}
type lakeformationResourceLfTagsAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_lakeformation_resource_lf_tags.
func (lrlt lakeformationResourceLfTagsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(lrlt.ref.Append("catalog_id"))
}

// Id returns a reference to field id of aws_lakeformation_resource_lf_tags.
func (lrlt lakeformationResourceLfTagsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lrlt.ref.Append("id"))
}

func (lrlt lakeformationResourceLfTagsAttributes) Database() terra.ListValue[lakeformationresourcelftags.DatabaseAttributes] {
	return terra.ReferenceAsList[lakeformationresourcelftags.DatabaseAttributes](lrlt.ref.Append("database"))
}

func (lrlt lakeformationResourceLfTagsAttributes) LfTag() terra.SetValue[lakeformationresourcelftags.LfTagAttributes] {
	return terra.ReferenceAsSet[lakeformationresourcelftags.LfTagAttributes](lrlt.ref.Append("lf_tag"))
}

func (lrlt lakeformationResourceLfTagsAttributes) Table() terra.ListValue[lakeformationresourcelftags.TableAttributes] {
	return terra.ReferenceAsList[lakeformationresourcelftags.TableAttributes](lrlt.ref.Append("table"))
}

func (lrlt lakeformationResourceLfTagsAttributes) TableWithColumns() terra.ListValue[lakeformationresourcelftags.TableWithColumnsAttributes] {
	return terra.ReferenceAsList[lakeformationresourcelftags.TableWithColumnsAttributes](lrlt.ref.Append("table_with_columns"))
}

func (lrlt lakeformationResourceLfTagsAttributes) Timeouts() lakeformationresourcelftags.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lakeformationresourcelftags.TimeoutsAttributes](lrlt.ref.Append("timeouts"))
}

type lakeformationResourceLfTagsState struct {
	CatalogId        string                                              `json:"catalog_id"`
	Id               string                                              `json:"id"`
	Database         []lakeformationresourcelftags.DatabaseState         `json:"database"`
	LfTag            []lakeformationresourcelftags.LfTagState            `json:"lf_tag"`
	Table            []lakeformationresourcelftags.TableState            `json:"table"`
	TableWithColumns []lakeformationresourcelftags.TableWithColumnsState `json:"table_with_columns"`
	Timeouts         *lakeformationresourcelftags.TimeoutsState          `json:"timeouts"`
}
