// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLakeformationLfTag(name string, args LakeformationLfTagArgs) *LakeformationLfTag {
	return &LakeformationLfTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LakeformationLfTag)(nil)

type LakeformationLfTag struct {
	Name  string
	Args  LakeformationLfTagArgs
	state *lakeformationLfTagState
}

func (llt *LakeformationLfTag) Type() string {
	return "aws_lakeformation_lf_tag"
}

func (llt *LakeformationLfTag) LocalName() string {
	return llt.Name
}

func (llt *LakeformationLfTag) Configuration() interface{} {
	return llt.Args
}

func (llt *LakeformationLfTag) Attributes() lakeformationLfTagAttributes {
	return lakeformationLfTagAttributes{ref: terra.ReferenceResource(llt)}
}

func (llt *LakeformationLfTag) ImportState(av io.Reader) error {
	llt.state = &lakeformationLfTagState{}
	if err := json.NewDecoder(av).Decode(llt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llt.Type(), llt.LocalName(), err)
	}
	return nil
}

func (llt *LakeformationLfTag) State() (*lakeformationLfTagState, bool) {
	return llt.state, llt.state != nil
}

func (llt *LakeformationLfTag) StateMust() *lakeformationLfTagState {
	if llt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llt.Type(), llt.LocalName()))
	}
	return llt.state
}

func (llt *LakeformationLfTag) DependOn() terra.Reference {
	return terra.ReferenceResource(llt)
}

type LakeformationLfTagArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
	// DependsOn contains resources that LakeformationLfTag depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lakeformationLfTagAttributes struct {
	ref terra.Reference
}

func (llt lakeformationLfTagAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceString(llt.ref.Append("catalog_id"))
}

func (llt lakeformationLfTagAttributes) Id() terra.StringValue {
	return terra.ReferenceString(llt.ref.Append("id"))
}

func (llt lakeformationLfTagAttributes) Key() terra.StringValue {
	return terra.ReferenceString(llt.ref.Append("key"))
}

func (llt lakeformationLfTagAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](llt.ref.Append("values"))
}

type lakeformationLfTagState struct {
	CatalogId string   `json:"catalog_id"`
	Id        string   `json:"id"`
	Key       string   `json:"key"`
	Values    []string `json:"values"`
}
