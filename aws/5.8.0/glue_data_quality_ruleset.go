// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	gluedataqualityruleset "github.com/golingon/terraproviders/aws/5.8.0/gluedataqualityruleset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlueDataQualityRuleset creates a new instance of [GlueDataQualityRuleset].
func NewGlueDataQualityRuleset(name string, args GlueDataQualityRulesetArgs) *GlueDataQualityRuleset {
	return &GlueDataQualityRuleset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlueDataQualityRuleset)(nil)

// GlueDataQualityRuleset represents the Terraform resource aws_glue_data_quality_ruleset.
type GlueDataQualityRuleset struct {
	Name      string
	Args      GlueDataQualityRulesetArgs
	state     *glueDataQualityRulesetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlueDataQualityRuleset].
func (gdqr *GlueDataQualityRuleset) Type() string {
	return "aws_glue_data_quality_ruleset"
}

// LocalName returns the local name for [GlueDataQualityRuleset].
func (gdqr *GlueDataQualityRuleset) LocalName() string {
	return gdqr.Name
}

// Configuration returns the configuration (args) for [GlueDataQualityRuleset].
func (gdqr *GlueDataQualityRuleset) Configuration() interface{} {
	return gdqr.Args
}

// DependOn is used for other resources to depend on [GlueDataQualityRuleset].
func (gdqr *GlueDataQualityRuleset) DependOn() terra.Reference {
	return terra.ReferenceResource(gdqr)
}

// Dependencies returns the list of resources [GlueDataQualityRuleset] depends_on.
func (gdqr *GlueDataQualityRuleset) Dependencies() terra.Dependencies {
	return gdqr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlueDataQualityRuleset].
func (gdqr *GlueDataQualityRuleset) LifecycleManagement() *terra.Lifecycle {
	return gdqr.Lifecycle
}

// Attributes returns the attributes for [GlueDataQualityRuleset].
func (gdqr *GlueDataQualityRuleset) Attributes() glueDataQualityRulesetAttributes {
	return glueDataQualityRulesetAttributes{ref: terra.ReferenceResource(gdqr)}
}

// ImportState imports the given attribute values into [GlueDataQualityRuleset]'s state.
func (gdqr *GlueDataQualityRuleset) ImportState(av io.Reader) error {
	gdqr.state = &glueDataQualityRulesetState{}
	if err := json.NewDecoder(av).Decode(gdqr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdqr.Type(), gdqr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlueDataQualityRuleset] has state.
func (gdqr *GlueDataQualityRuleset) State() (*glueDataQualityRulesetState, bool) {
	return gdqr.state, gdqr.state != nil
}

// StateMust returns the state for [GlueDataQualityRuleset]. Panics if the state is nil.
func (gdqr *GlueDataQualityRuleset) StateMust() *glueDataQualityRulesetState {
	if gdqr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdqr.Type(), gdqr.LocalName()))
	}
	return gdqr.state
}

// GlueDataQualityRulesetArgs contains the configurations for aws_glue_data_quality_ruleset.
type GlueDataQualityRulesetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Ruleset: string, required
	Ruleset terra.StringValue `hcl:"ruleset,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetTable: optional
	TargetTable *gluedataqualityruleset.TargetTable `hcl:"target_table,block"`
}
type glueDataQualityRulesetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("arn"))
}

// CreatedOn returns a reference to field created_on of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) CreatedOn() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("created_on"))
}

// Description returns a reference to field description of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("description"))
}

// Id returns a reference to field id of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("id"))
}

// LastModifiedOn returns a reference to field last_modified_on of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) LastModifiedOn() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("last_modified_on"))
}

// Name returns a reference to field name of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("name"))
}

// RecommendationRunId returns a reference to field recommendation_run_id of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) RecommendationRunId() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("recommendation_run_id"))
}

// Ruleset returns a reference to field ruleset of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) Ruleset() terra.StringValue {
	return terra.ReferenceAsString(gdqr.ref.Append("ruleset"))
}

// Tags returns a reference to field tags of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gdqr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_glue_data_quality_ruleset.
func (gdqr glueDataQualityRulesetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gdqr.ref.Append("tags_all"))
}

func (gdqr glueDataQualityRulesetAttributes) TargetTable() terra.ListValue[gluedataqualityruleset.TargetTableAttributes] {
	return terra.ReferenceAsList[gluedataqualityruleset.TargetTableAttributes](gdqr.ref.Append("target_table"))
}

type glueDataQualityRulesetState struct {
	Arn                 string                                    `json:"arn"`
	CreatedOn           string                                    `json:"created_on"`
	Description         string                                    `json:"description"`
	Id                  string                                    `json:"id"`
	LastModifiedOn      string                                    `json:"last_modified_on"`
	Name                string                                    `json:"name"`
	RecommendationRunId string                                    `json:"recommendation_run_id"`
	Ruleset             string                                    `json:"ruleset"`
	Tags                map[string]string                         `json:"tags"`
	TagsAll             map[string]string                         `json:"tags_all"`
	TargetTable         []gluedataqualityruleset.TargetTableState `json:"target_table"`
}
