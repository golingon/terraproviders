// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_guardduty_detector

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

// Resource represents the Terraform resource aws_guardduty_detector.
type Resource struct {
	Name      string
	Args      Args
	state     *awsGuarddutyDetectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (agd *Resource) Type() string {
	return "aws_guardduty_detector"
}

// LocalName returns the local name for [Resource].
func (agd *Resource) LocalName() string {
	return agd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (agd *Resource) Configuration() interface{} {
	return agd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (agd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(agd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (agd *Resource) Dependencies() terra.Dependencies {
	return agd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (agd *Resource) LifecycleManagement() *terra.Lifecycle {
	return agd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (agd *Resource) Attributes() awsGuarddutyDetectorAttributes {
	return awsGuarddutyDetectorAttributes{ref: terra.ReferenceResource(agd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (agd *Resource) ImportState(state io.Reader) error {
	agd.state = &awsGuarddutyDetectorState{}
	if err := json.NewDecoder(state).Decode(agd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agd.Type(), agd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (agd *Resource) State() (*awsGuarddutyDetectorState, bool) {
	return agd.state, agd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (agd *Resource) StateMust() *awsGuarddutyDetectorState {
	if agd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agd.Type(), agd.LocalName()))
	}
	return agd.state
}

// Args contains the configurations for aws_guardduty_detector.
type Args struct {
	// Enable: bool, optional
	Enable terra.BoolValue `hcl:"enable,attr"`
	// FindingPublishingFrequency: string, optional
	FindingPublishingFrequency terra.StringValue `hcl:"finding_publishing_frequency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Datasources: optional
	Datasources *Datasources `hcl:"datasources,block"`
}

type awsGuarddutyDetectorAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("account_id"))
}

// Arn returns a reference to field arn of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("arn"))
}

// Enable returns a reference to field enable of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(agd.ref.Append("enable"))
}

// FindingPublishingFrequency returns a reference to field finding_publishing_frequency of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) FindingPublishingFrequency() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("finding_publishing_frequency"))
}

// Id returns a reference to field id of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_guardduty_detector.
func (agd awsGuarddutyDetectorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agd.ref.Append("tags_all"))
}

func (agd awsGuarddutyDetectorAttributes) Datasources() terra.ListValue[DatasourcesAttributes] {
	return terra.ReferenceAsList[DatasourcesAttributes](agd.ref.Append("datasources"))
}

type awsGuarddutyDetectorState struct {
	AccountId                  string             `json:"account_id"`
	Arn                        string             `json:"arn"`
	Enable                     bool               `json:"enable"`
	FindingPublishingFrequency string             `json:"finding_publishing_frequency"`
	Id                         string             `json:"id"`
	Tags                       map[string]string  `json:"tags"`
	TagsAll                    map[string]string  `json:"tags_all"`
	Datasources                []DatasourcesState `json:"datasources"`
}
