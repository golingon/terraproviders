// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	evidentlyfeature "github.com/golingon/terraproviders/aws/5.12.0/evidentlyfeature"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEvidentlyFeature creates a new instance of [EvidentlyFeature].
func NewEvidentlyFeature(name string, args EvidentlyFeatureArgs) *EvidentlyFeature {
	return &EvidentlyFeature{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EvidentlyFeature)(nil)

// EvidentlyFeature represents the Terraform resource aws_evidently_feature.
type EvidentlyFeature struct {
	Name      string
	Args      EvidentlyFeatureArgs
	state     *evidentlyFeatureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EvidentlyFeature].
func (ef *EvidentlyFeature) Type() string {
	return "aws_evidently_feature"
}

// LocalName returns the local name for [EvidentlyFeature].
func (ef *EvidentlyFeature) LocalName() string {
	return ef.Name
}

// Configuration returns the configuration (args) for [EvidentlyFeature].
func (ef *EvidentlyFeature) Configuration() interface{} {
	return ef.Args
}

// DependOn is used for other resources to depend on [EvidentlyFeature].
func (ef *EvidentlyFeature) DependOn() terra.Reference {
	return terra.ReferenceResource(ef)
}

// Dependencies returns the list of resources [EvidentlyFeature] depends_on.
func (ef *EvidentlyFeature) Dependencies() terra.Dependencies {
	return ef.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EvidentlyFeature].
func (ef *EvidentlyFeature) LifecycleManagement() *terra.Lifecycle {
	return ef.Lifecycle
}

// Attributes returns the attributes for [EvidentlyFeature].
func (ef *EvidentlyFeature) Attributes() evidentlyFeatureAttributes {
	return evidentlyFeatureAttributes{ref: terra.ReferenceResource(ef)}
}

// ImportState imports the given attribute values into [EvidentlyFeature]'s state.
func (ef *EvidentlyFeature) ImportState(av io.Reader) error {
	ef.state = &evidentlyFeatureState{}
	if err := json.NewDecoder(av).Decode(ef.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ef.Type(), ef.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EvidentlyFeature] has state.
func (ef *EvidentlyFeature) State() (*evidentlyFeatureState, bool) {
	return ef.state, ef.state != nil
}

// StateMust returns the state for [EvidentlyFeature]. Panics if the state is nil.
func (ef *EvidentlyFeature) StateMust() *evidentlyFeatureState {
	if ef.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ef.Type(), ef.LocalName()))
	}
	return ef.state
}

// EvidentlyFeatureArgs contains the configurations for aws_evidently_feature.
type EvidentlyFeatureArgs struct {
	// DefaultVariation: string, optional
	DefaultVariation terra.StringValue `hcl:"default_variation,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EntityOverrides: map of string, optional
	EntityOverrides terra.MapValue[terra.StringValue] `hcl:"entity_overrides,attr"`
	// EvaluationStrategy: string, optional
	EvaluationStrategy terra.StringValue `hcl:"evaluation_strategy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// EvaluationRules: min=0
	EvaluationRules []evidentlyfeature.EvaluationRules `hcl:"evaluation_rules,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *evidentlyfeature.Timeouts `hcl:"timeouts,block"`
	// Variations: min=1,max=5
	Variations []evidentlyfeature.Variations `hcl:"variations,block" validate:"min=1,max=5"`
}
type evidentlyFeatureAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("created_time"))
}

// DefaultVariation returns a reference to field default_variation of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) DefaultVariation() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("default_variation"))
}

// Description returns a reference to field description of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("description"))
}

// EntityOverrides returns a reference to field entity_overrides of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) EntityOverrides() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ef.ref.Append("entity_overrides"))
}

// EvaluationStrategy returns a reference to field evaluation_strategy of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) EvaluationStrategy() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("evaluation_strategy"))
}

// Id returns a reference to field id of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("id"))
}

// LastUpdatedTime returns a reference to field last_updated_time of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) LastUpdatedTime() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("last_updated_time"))
}

// Name returns a reference to field name of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("name"))
}

// Project returns a reference to field project of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("project"))
}

// Status returns a reference to field status of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ef.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ef.ref.Append("tags_all"))
}

// ValueType returns a reference to field value_type of aws_evidently_feature.
func (ef evidentlyFeatureAttributes) ValueType() terra.StringValue {
	return terra.ReferenceAsString(ef.ref.Append("value_type"))
}

func (ef evidentlyFeatureAttributes) EvaluationRules() terra.SetValue[evidentlyfeature.EvaluationRulesAttributes] {
	return terra.ReferenceAsSet[evidentlyfeature.EvaluationRulesAttributes](ef.ref.Append("evaluation_rules"))
}

func (ef evidentlyFeatureAttributes) Timeouts() evidentlyfeature.TimeoutsAttributes {
	return terra.ReferenceAsSingle[evidentlyfeature.TimeoutsAttributes](ef.ref.Append("timeouts"))
}

func (ef evidentlyFeatureAttributes) Variations() terra.SetValue[evidentlyfeature.VariationsAttributes] {
	return terra.ReferenceAsSet[evidentlyfeature.VariationsAttributes](ef.ref.Append("variations"))
}

type evidentlyFeatureState struct {
	Arn                string                                  `json:"arn"`
	CreatedTime        string                                  `json:"created_time"`
	DefaultVariation   string                                  `json:"default_variation"`
	Description        string                                  `json:"description"`
	EntityOverrides    map[string]string                       `json:"entity_overrides"`
	EvaluationStrategy string                                  `json:"evaluation_strategy"`
	Id                 string                                  `json:"id"`
	LastUpdatedTime    string                                  `json:"last_updated_time"`
	Name               string                                  `json:"name"`
	Project            string                                  `json:"project"`
	Status             string                                  `json:"status"`
	Tags               map[string]string                       `json:"tags"`
	TagsAll            map[string]string                       `json:"tags_all"`
	ValueType          string                                  `json:"value_type"`
	EvaluationRules    []evidentlyfeature.EvaluationRulesState `json:"evaluation_rules"`
	Timeouts           *evidentlyfeature.TimeoutsState         `json:"timeouts"`
	Variations         []evidentlyfeature.VariationsState      `json:"variations"`
}
