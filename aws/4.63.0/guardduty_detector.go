// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	guarddutydetector "github.com/golingon/terraproviders/aws/4.63.0/guarddutydetector"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGuarddutyDetector creates a new instance of [GuarddutyDetector].
func NewGuarddutyDetector(name string, args GuarddutyDetectorArgs) *GuarddutyDetector {
	return &GuarddutyDetector{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GuarddutyDetector)(nil)

// GuarddutyDetector represents the Terraform resource aws_guardduty_detector.
type GuarddutyDetector struct {
	Name      string
	Args      GuarddutyDetectorArgs
	state     *guarddutyDetectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GuarddutyDetector].
func (gd *GuarddutyDetector) Type() string {
	return "aws_guardduty_detector"
}

// LocalName returns the local name for [GuarddutyDetector].
func (gd *GuarddutyDetector) LocalName() string {
	return gd.Name
}

// Configuration returns the configuration (args) for [GuarddutyDetector].
func (gd *GuarddutyDetector) Configuration() interface{} {
	return gd.Args
}

// DependOn is used for other resources to depend on [GuarddutyDetector].
func (gd *GuarddutyDetector) DependOn() terra.Reference {
	return terra.ReferenceResource(gd)
}

// Dependencies returns the list of resources [GuarddutyDetector] depends_on.
func (gd *GuarddutyDetector) Dependencies() terra.Dependencies {
	return gd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GuarddutyDetector].
func (gd *GuarddutyDetector) LifecycleManagement() *terra.Lifecycle {
	return gd.Lifecycle
}

// Attributes returns the attributes for [GuarddutyDetector].
func (gd *GuarddutyDetector) Attributes() guarddutyDetectorAttributes {
	return guarddutyDetectorAttributes{ref: terra.ReferenceResource(gd)}
}

// ImportState imports the given attribute values into [GuarddutyDetector]'s state.
func (gd *GuarddutyDetector) ImportState(av io.Reader) error {
	gd.state = &guarddutyDetectorState{}
	if err := json.NewDecoder(av).Decode(gd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gd.Type(), gd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GuarddutyDetector] has state.
func (gd *GuarddutyDetector) State() (*guarddutyDetectorState, bool) {
	return gd.state, gd.state != nil
}

// StateMust returns the state for [GuarddutyDetector]. Panics if the state is nil.
func (gd *GuarddutyDetector) StateMust() *guarddutyDetectorState {
	if gd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gd.Type(), gd.LocalName()))
	}
	return gd.state
}

// GuarddutyDetectorArgs contains the configurations for aws_guardduty_detector.
type GuarddutyDetectorArgs struct {
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
	Datasources *guarddutydetector.Datasources `hcl:"datasources,block"`
}
type guarddutyDetectorAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("account_id"))
}

// Arn returns a reference to field arn of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("arn"))
}

// Enable returns a reference to field enable of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(gd.ref.Append("enable"))
}

// FindingPublishingFrequency returns a reference to field finding_publishing_frequency of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) FindingPublishingFrequency() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("finding_publishing_frequency"))
}

// Id returns a reference to field id of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gd.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_guardduty_detector.
func (gd guarddutyDetectorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gd.ref.Append("tags_all"))
}

func (gd guarddutyDetectorAttributes) Datasources() terra.ListValue[guarddutydetector.DatasourcesAttributes] {
	return terra.ReferenceAsList[guarddutydetector.DatasourcesAttributes](gd.ref.Append("datasources"))
}

type guarddutyDetectorState struct {
	AccountId                  string                               `json:"account_id"`
	Arn                        string                               `json:"arn"`
	Enable                     bool                                 `json:"enable"`
	FindingPublishingFrequency string                               `json:"finding_publishing_frequency"`
	Id                         string                               `json:"id"`
	Tags                       map[string]string                    `json:"tags"`
	TagsAll                    map[string]string                    `json:"tags_all"`
	Datasources                []guarddutydetector.DatasourcesState `json:"datasources"`
}
