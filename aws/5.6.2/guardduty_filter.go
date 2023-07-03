// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	guarddutyfilter "github.com/golingon/terraproviders/aws/5.6.2/guarddutyfilter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGuarddutyFilter creates a new instance of [GuarddutyFilter].
func NewGuarddutyFilter(name string, args GuarddutyFilterArgs) *GuarddutyFilter {
	return &GuarddutyFilter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GuarddutyFilter)(nil)

// GuarddutyFilter represents the Terraform resource aws_guardduty_filter.
type GuarddutyFilter struct {
	Name      string
	Args      GuarddutyFilterArgs
	state     *guarddutyFilterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GuarddutyFilter].
func (gf *GuarddutyFilter) Type() string {
	return "aws_guardduty_filter"
}

// LocalName returns the local name for [GuarddutyFilter].
func (gf *GuarddutyFilter) LocalName() string {
	return gf.Name
}

// Configuration returns the configuration (args) for [GuarddutyFilter].
func (gf *GuarddutyFilter) Configuration() interface{} {
	return gf.Args
}

// DependOn is used for other resources to depend on [GuarddutyFilter].
func (gf *GuarddutyFilter) DependOn() terra.Reference {
	return terra.ReferenceResource(gf)
}

// Dependencies returns the list of resources [GuarddutyFilter] depends_on.
func (gf *GuarddutyFilter) Dependencies() terra.Dependencies {
	return gf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GuarddutyFilter].
func (gf *GuarddutyFilter) LifecycleManagement() *terra.Lifecycle {
	return gf.Lifecycle
}

// Attributes returns the attributes for [GuarddutyFilter].
func (gf *GuarddutyFilter) Attributes() guarddutyFilterAttributes {
	return guarddutyFilterAttributes{ref: terra.ReferenceResource(gf)}
}

// ImportState imports the given attribute values into [GuarddutyFilter]'s state.
func (gf *GuarddutyFilter) ImportState(av io.Reader) error {
	gf.state = &guarddutyFilterState{}
	if err := json.NewDecoder(av).Decode(gf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gf.Type(), gf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GuarddutyFilter] has state.
func (gf *GuarddutyFilter) State() (*guarddutyFilterState, bool) {
	return gf.state, gf.state != nil
}

// StateMust returns the state for [GuarddutyFilter]. Panics if the state is nil.
func (gf *GuarddutyFilter) StateMust() *guarddutyFilterState {
	if gf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gf.Type(), gf.LocalName()))
	}
	return gf.state
}

// GuarddutyFilterArgs contains the configurations for aws_guardduty_filter.
type GuarddutyFilterArgs struct {
	// Action: string, required
	Action terra.StringValue `hcl:"action,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DetectorId: string, required
	DetectorId terra.StringValue `hcl:"detector_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Rank: number, required
	Rank terra.NumberValue `hcl:"rank,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// FindingCriteria: required
	FindingCriteria *guarddutyfilter.FindingCriteria `hcl:"finding_criteria,block" validate:"required"`
}
type guarddutyFilterAttributes struct {
	ref terra.Reference
}

// Action returns a reference to field action of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Action() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("action"))
}

// Arn returns a reference to field arn of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("arn"))
}

// Description returns a reference to field description of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("description"))
}

// DetectorId returns a reference to field detector_id of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("detector_id"))
}

// Id returns a reference to field id of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("id"))
}

// Name returns a reference to field name of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gf.ref.Append("name"))
}

// Rank returns a reference to field rank of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Rank() terra.NumberValue {
	return terra.ReferenceAsNumber(gf.ref.Append("rank"))
}

// Tags returns a reference to field tags of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gf.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_guardduty_filter.
func (gf guarddutyFilterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gf.ref.Append("tags_all"))
}

func (gf guarddutyFilterAttributes) FindingCriteria() terra.ListValue[guarddutyfilter.FindingCriteriaAttributes] {
	return terra.ReferenceAsList[guarddutyfilter.FindingCriteriaAttributes](gf.ref.Append("finding_criteria"))
}

type guarddutyFilterState struct {
	Action          string                                 `json:"action"`
	Arn             string                                 `json:"arn"`
	Description     string                                 `json:"description"`
	DetectorId      string                                 `json:"detector_id"`
	Id              string                                 `json:"id"`
	Name            string                                 `json:"name"`
	Rank            float64                                `json:"rank"`
	Tags            map[string]string                      `json:"tags"`
	TagsAll         map[string]string                      `json:"tags_all"`
	FindingCriteria []guarddutyfilter.FindingCriteriaState `json:"finding_criteria"`
}
