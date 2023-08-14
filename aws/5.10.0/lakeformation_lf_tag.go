// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLakeformationLfTag creates a new instance of [LakeformationLfTag].
func NewLakeformationLfTag(name string, args LakeformationLfTagArgs) *LakeformationLfTag {
	return &LakeformationLfTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LakeformationLfTag)(nil)

// LakeformationLfTag represents the Terraform resource aws_lakeformation_lf_tag.
type LakeformationLfTag struct {
	Name      string
	Args      LakeformationLfTagArgs
	state     *lakeformationLfTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LakeformationLfTag].
func (llt *LakeformationLfTag) Type() string {
	return "aws_lakeformation_lf_tag"
}

// LocalName returns the local name for [LakeformationLfTag].
func (llt *LakeformationLfTag) LocalName() string {
	return llt.Name
}

// Configuration returns the configuration (args) for [LakeformationLfTag].
func (llt *LakeformationLfTag) Configuration() interface{} {
	return llt.Args
}

// DependOn is used for other resources to depend on [LakeformationLfTag].
func (llt *LakeformationLfTag) DependOn() terra.Reference {
	return terra.ReferenceResource(llt)
}

// Dependencies returns the list of resources [LakeformationLfTag] depends_on.
func (llt *LakeformationLfTag) Dependencies() terra.Dependencies {
	return llt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LakeformationLfTag].
func (llt *LakeformationLfTag) LifecycleManagement() *terra.Lifecycle {
	return llt.Lifecycle
}

// Attributes returns the attributes for [LakeformationLfTag].
func (llt *LakeformationLfTag) Attributes() lakeformationLfTagAttributes {
	return lakeformationLfTagAttributes{ref: terra.ReferenceResource(llt)}
}

// ImportState imports the given attribute values into [LakeformationLfTag]'s state.
func (llt *LakeformationLfTag) ImportState(av io.Reader) error {
	llt.state = &lakeformationLfTagState{}
	if err := json.NewDecoder(av).Decode(llt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llt.Type(), llt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LakeformationLfTag] has state.
func (llt *LakeformationLfTag) State() (*lakeformationLfTagState, bool) {
	return llt.state, llt.state != nil
}

// StateMust returns the state for [LakeformationLfTag]. Panics if the state is nil.
func (llt *LakeformationLfTag) StateMust() *lakeformationLfTagState {
	if llt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llt.Type(), llt.LocalName()))
	}
	return llt.state
}

// LakeformationLfTagArgs contains the configurations for aws_lakeformation_lf_tag.
type LakeformationLfTagArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}
type lakeformationLfTagAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_lakeformation_lf_tag.
func (llt lakeformationLfTagAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(llt.ref.Append("catalog_id"))
}

// Id returns a reference to field id of aws_lakeformation_lf_tag.
func (llt lakeformationLfTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(llt.ref.Append("id"))
}

// Key returns a reference to field key of aws_lakeformation_lf_tag.
func (llt lakeformationLfTagAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(llt.ref.Append("key"))
}

// Values returns a reference to field values of aws_lakeformation_lf_tag.
func (llt lakeformationLfTagAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](llt.ref.Append("values"))
}

type lakeformationLfTagState struct {
	CatalogId string   `json:"catalog_id"`
	Id        string   `json:"id"`
	Key       string   `json:"key"`
	Values    []string `json:"values"`
}