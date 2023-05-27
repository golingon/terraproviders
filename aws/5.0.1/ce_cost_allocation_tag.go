// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCeCostAllocationTag creates a new instance of [CeCostAllocationTag].
func NewCeCostAllocationTag(name string, args CeCostAllocationTagArgs) *CeCostAllocationTag {
	return &CeCostAllocationTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CeCostAllocationTag)(nil)

// CeCostAllocationTag represents the Terraform resource aws_ce_cost_allocation_tag.
type CeCostAllocationTag struct {
	Name      string
	Args      CeCostAllocationTagArgs
	state     *ceCostAllocationTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CeCostAllocationTag].
func (ccat *CeCostAllocationTag) Type() string {
	return "aws_ce_cost_allocation_tag"
}

// LocalName returns the local name for [CeCostAllocationTag].
func (ccat *CeCostAllocationTag) LocalName() string {
	return ccat.Name
}

// Configuration returns the configuration (args) for [CeCostAllocationTag].
func (ccat *CeCostAllocationTag) Configuration() interface{} {
	return ccat.Args
}

// DependOn is used for other resources to depend on [CeCostAllocationTag].
func (ccat *CeCostAllocationTag) DependOn() terra.Reference {
	return terra.ReferenceResource(ccat)
}

// Dependencies returns the list of resources [CeCostAllocationTag] depends_on.
func (ccat *CeCostAllocationTag) Dependencies() terra.Dependencies {
	return ccat.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CeCostAllocationTag].
func (ccat *CeCostAllocationTag) LifecycleManagement() *terra.Lifecycle {
	return ccat.Lifecycle
}

// Attributes returns the attributes for [CeCostAllocationTag].
func (ccat *CeCostAllocationTag) Attributes() ceCostAllocationTagAttributes {
	return ceCostAllocationTagAttributes{ref: terra.ReferenceResource(ccat)}
}

// ImportState imports the given attribute values into [CeCostAllocationTag]'s state.
func (ccat *CeCostAllocationTag) ImportState(av io.Reader) error {
	ccat.state = &ceCostAllocationTagState{}
	if err := json.NewDecoder(av).Decode(ccat.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccat.Type(), ccat.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CeCostAllocationTag] has state.
func (ccat *CeCostAllocationTag) State() (*ceCostAllocationTagState, bool) {
	return ccat.state, ccat.state != nil
}

// StateMust returns the state for [CeCostAllocationTag]. Panics if the state is nil.
func (ccat *CeCostAllocationTag) StateMust() *ceCostAllocationTagState {
	if ccat.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccat.Type(), ccat.LocalName()))
	}
	return ccat.state
}

// CeCostAllocationTagArgs contains the configurations for aws_ce_cost_allocation_tag.
type CeCostAllocationTagArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Status: string, required
	Status terra.StringValue `hcl:"status,attr" validate:"required"`
	// TagKey: string, required
	TagKey terra.StringValue `hcl:"tag_key,attr" validate:"required"`
}
type ceCostAllocationTagAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ce_cost_allocation_tag.
func (ccat ceCostAllocationTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("id"))
}

// Status returns a reference to field status of aws_ce_cost_allocation_tag.
func (ccat ceCostAllocationTagAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("status"))
}

// TagKey returns a reference to field tag_key of aws_ce_cost_allocation_tag.
func (ccat ceCostAllocationTagAttributes) TagKey() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("tag_key"))
}

// Type returns a reference to field type of aws_ce_cost_allocation_tag.
func (ccat ceCostAllocationTagAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ccat.ref.Append("type"))
}

type ceCostAllocationTagState struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	TagKey string `json:"tag_key"`
	Type   string `json:"type"`
}
