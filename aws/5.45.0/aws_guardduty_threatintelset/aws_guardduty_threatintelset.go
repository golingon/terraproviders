// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_guardduty_threatintelset

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

// Resource represents the Terraform resource aws_guardduty_threatintelset.
type Resource struct {
	Name      string
	Args      Args
	state     *awsGuarddutyThreatintelsetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (agt *Resource) Type() string {
	return "aws_guardduty_threatintelset"
}

// LocalName returns the local name for [Resource].
func (agt *Resource) LocalName() string {
	return agt.Name
}

// Configuration returns the configuration (args) for [Resource].
func (agt *Resource) Configuration() interface{} {
	return agt.Args
}

// DependOn is used for other resources to depend on [Resource].
func (agt *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(agt)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (agt *Resource) Dependencies() terra.Dependencies {
	return agt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (agt *Resource) LifecycleManagement() *terra.Lifecycle {
	return agt.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (agt *Resource) Attributes() awsGuarddutyThreatintelsetAttributes {
	return awsGuarddutyThreatintelsetAttributes{ref: terra.ReferenceResource(agt)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (agt *Resource) ImportState(state io.Reader) error {
	agt.state = &awsGuarddutyThreatintelsetState{}
	if err := json.NewDecoder(state).Decode(agt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agt.Type(), agt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (agt *Resource) State() (*awsGuarddutyThreatintelsetState, bool) {
	return agt.state, agt.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (agt *Resource) StateMust() *awsGuarddutyThreatintelsetState {
	if agt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agt.Type(), agt.LocalName()))
	}
	return agt.state
}

// Args contains the configurations for aws_guardduty_threatintelset.
type Args struct {
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

type awsGuarddutyThreatintelsetAttributes struct {
	ref terra.Reference
}

// Activate returns a reference to field activate of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Activate() terra.BoolValue {
	return terra.ReferenceAsBool(agt.ref.Append("activate"))
}

// Arn returns a reference to field arn of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(agt.ref.Append("arn"))
}

// DetectorId returns a reference to field detector_id of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(agt.ref.Append("detector_id"))
}

// Format returns a reference to field format of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(agt.ref.Append("format"))
}

// Id returns a reference to field id of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agt.ref.Append("id"))
}

// Location returns a reference to field location of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(agt.ref.Append("location"))
}

// Name returns a reference to field name of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(agt.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_guardduty_threatintelset.
func (agt awsGuarddutyThreatintelsetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agt.ref.Append("tags_all"))
}

type awsGuarddutyThreatintelsetState struct {
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
