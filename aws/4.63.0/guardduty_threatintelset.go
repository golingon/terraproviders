// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGuarddutyThreatintelset creates a new instance of [GuarddutyThreatintelset].
func NewGuarddutyThreatintelset(name string, args GuarddutyThreatintelsetArgs) *GuarddutyThreatintelset {
	return &GuarddutyThreatintelset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GuarddutyThreatintelset)(nil)

// GuarddutyThreatintelset represents the Terraform resource aws_guardduty_threatintelset.
type GuarddutyThreatintelset struct {
	Name      string
	Args      GuarddutyThreatintelsetArgs
	state     *guarddutyThreatintelsetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GuarddutyThreatintelset].
func (gt *GuarddutyThreatintelset) Type() string {
	return "aws_guardduty_threatintelset"
}

// LocalName returns the local name for [GuarddutyThreatintelset].
func (gt *GuarddutyThreatintelset) LocalName() string {
	return gt.Name
}

// Configuration returns the configuration (args) for [GuarddutyThreatintelset].
func (gt *GuarddutyThreatintelset) Configuration() interface{} {
	return gt.Args
}

// DependOn is used for other resources to depend on [GuarddutyThreatintelset].
func (gt *GuarddutyThreatintelset) DependOn() terra.Reference {
	return terra.ReferenceResource(gt)
}

// Dependencies returns the list of resources [GuarddutyThreatintelset] depends_on.
func (gt *GuarddutyThreatintelset) Dependencies() terra.Dependencies {
	return gt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GuarddutyThreatintelset].
func (gt *GuarddutyThreatintelset) LifecycleManagement() *terra.Lifecycle {
	return gt.Lifecycle
}

// Attributes returns the attributes for [GuarddutyThreatintelset].
func (gt *GuarddutyThreatintelset) Attributes() guarddutyThreatintelsetAttributes {
	return guarddutyThreatintelsetAttributes{ref: terra.ReferenceResource(gt)}
}

// ImportState imports the given attribute values into [GuarddutyThreatintelset]'s state.
func (gt *GuarddutyThreatintelset) ImportState(av io.Reader) error {
	gt.state = &guarddutyThreatintelsetState{}
	if err := json.NewDecoder(av).Decode(gt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gt.Type(), gt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GuarddutyThreatintelset] has state.
func (gt *GuarddutyThreatintelset) State() (*guarddutyThreatintelsetState, bool) {
	return gt.state, gt.state != nil
}

// StateMust returns the state for [GuarddutyThreatintelset]. Panics if the state is nil.
func (gt *GuarddutyThreatintelset) StateMust() *guarddutyThreatintelsetState {
	if gt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gt.Type(), gt.LocalName()))
	}
	return gt.state
}

// GuarddutyThreatintelsetArgs contains the configurations for aws_guardduty_threatintelset.
type GuarddutyThreatintelsetArgs struct {
	// Activate: bool, required
	Activate terra.BoolValue `hcl:"activate,attr" validate:"required"`
	// DetectorId: string, required
	DetectorId terra.StringValue `hcl:"detector_id,attr" validate:"required"`
	// Format: string, required
	Format terra.StringValue `hcl:"format,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type guarddutyThreatintelsetAttributes struct {
	ref terra.Reference
}

// Activate returns a reference to field activate of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Activate() terra.BoolValue {
	return terra.ReferenceAsBool(gt.ref.Append("activate"))
}

// Arn returns a reference to field arn of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("arn"))
}

// DetectorId returns a reference to field detector_id of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("detector_id"))
}

// Format returns a reference to field format of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("format"))
}

// Id returns a reference to field id of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("id"))
}

// Location returns a reference to field location of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("location"))
}

// Name returns a reference to field name of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gt.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_guardduty_threatintelset.
func (gt guarddutyThreatintelsetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gt.ref.Append("tags_all"))
}

type guarddutyThreatintelsetState struct {
	Activate   bool              `json:"activate"`
	Arn        string            `json:"arn"`
	DetectorId string            `json:"detector_id"`
	Format     string            `json:"format"`
	Id         string            `json:"id"`
	Location   string            `json:"location"`
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	TagsAll    map[string]string `json:"tags_all"`
}
