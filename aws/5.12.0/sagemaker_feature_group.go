// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerfeaturegroup "github.com/golingon/terraproviders/aws/5.12.0/sagemakerfeaturegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerFeatureGroup creates a new instance of [SagemakerFeatureGroup].
func NewSagemakerFeatureGroup(name string, args SagemakerFeatureGroupArgs) *SagemakerFeatureGroup {
	return &SagemakerFeatureGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerFeatureGroup)(nil)

// SagemakerFeatureGroup represents the Terraform resource aws_sagemaker_feature_group.
type SagemakerFeatureGroup struct {
	Name      string
	Args      SagemakerFeatureGroupArgs
	state     *sagemakerFeatureGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerFeatureGroup].
func (sfg *SagemakerFeatureGroup) Type() string {
	return "aws_sagemaker_feature_group"
}

// LocalName returns the local name for [SagemakerFeatureGroup].
func (sfg *SagemakerFeatureGroup) LocalName() string {
	return sfg.Name
}

// Configuration returns the configuration (args) for [SagemakerFeatureGroup].
func (sfg *SagemakerFeatureGroup) Configuration() interface{} {
	return sfg.Args
}

// DependOn is used for other resources to depend on [SagemakerFeatureGroup].
func (sfg *SagemakerFeatureGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(sfg)
}

// Dependencies returns the list of resources [SagemakerFeatureGroup] depends_on.
func (sfg *SagemakerFeatureGroup) Dependencies() terra.Dependencies {
	return sfg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerFeatureGroup].
func (sfg *SagemakerFeatureGroup) LifecycleManagement() *terra.Lifecycle {
	return sfg.Lifecycle
}

// Attributes returns the attributes for [SagemakerFeatureGroup].
func (sfg *SagemakerFeatureGroup) Attributes() sagemakerFeatureGroupAttributes {
	return sagemakerFeatureGroupAttributes{ref: terra.ReferenceResource(sfg)}
}

// ImportState imports the given attribute values into [SagemakerFeatureGroup]'s state.
func (sfg *SagemakerFeatureGroup) ImportState(av io.Reader) error {
	sfg.state = &sagemakerFeatureGroupState{}
	if err := json.NewDecoder(av).Decode(sfg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfg.Type(), sfg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerFeatureGroup] has state.
func (sfg *SagemakerFeatureGroup) State() (*sagemakerFeatureGroupState, bool) {
	return sfg.state, sfg.state != nil
}

// StateMust returns the state for [SagemakerFeatureGroup]. Panics if the state is nil.
func (sfg *SagemakerFeatureGroup) StateMust() *sagemakerFeatureGroupState {
	if sfg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfg.Type(), sfg.LocalName()))
	}
	return sfg.state
}

// SagemakerFeatureGroupArgs contains the configurations for aws_sagemaker_feature_group.
type SagemakerFeatureGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EventTimeFeatureName: string, required
	EventTimeFeatureName terra.StringValue `hcl:"event_time_feature_name,attr" validate:"required"`
	// FeatureGroupName: string, required
	FeatureGroupName terra.StringValue `hcl:"feature_group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RecordIdentifierFeatureName: string, required
	RecordIdentifierFeatureName terra.StringValue `hcl:"record_identifier_feature_name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// FeatureDefinition: min=1,max=2500
	FeatureDefinition []sagemakerfeaturegroup.FeatureDefinition `hcl:"feature_definition,block" validate:"min=1,max=2500"`
	// OfflineStoreConfig: optional
	OfflineStoreConfig *sagemakerfeaturegroup.OfflineStoreConfig `hcl:"offline_store_config,block"`
	// OnlineStoreConfig: optional
	OnlineStoreConfig *sagemakerfeaturegroup.OnlineStoreConfig `hcl:"online_store_config,block"`
}
type sagemakerFeatureGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("description"))
}

// EventTimeFeatureName returns a reference to field event_time_feature_name of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) EventTimeFeatureName() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("event_time_feature_name"))
}

// FeatureGroupName returns a reference to field feature_group_name of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) FeatureGroupName() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("feature_group_name"))
}

// Id returns a reference to field id of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("id"))
}

// RecordIdentifierFeatureName returns a reference to field record_identifier_feature_name of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) RecordIdentifierFeatureName() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("record_identifier_feature_name"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sfg.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_feature_group.
func (sfg sagemakerFeatureGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfg.ref.Append("tags_all"))
}

func (sfg sagemakerFeatureGroupAttributes) FeatureDefinition() terra.ListValue[sagemakerfeaturegroup.FeatureDefinitionAttributes] {
	return terra.ReferenceAsList[sagemakerfeaturegroup.FeatureDefinitionAttributes](sfg.ref.Append("feature_definition"))
}

func (sfg sagemakerFeatureGroupAttributes) OfflineStoreConfig() terra.ListValue[sagemakerfeaturegroup.OfflineStoreConfigAttributes] {
	return terra.ReferenceAsList[sagemakerfeaturegroup.OfflineStoreConfigAttributes](sfg.ref.Append("offline_store_config"))
}

func (sfg sagemakerFeatureGroupAttributes) OnlineStoreConfig() terra.ListValue[sagemakerfeaturegroup.OnlineStoreConfigAttributes] {
	return terra.ReferenceAsList[sagemakerfeaturegroup.OnlineStoreConfigAttributes](sfg.ref.Append("online_store_config"))
}

type sagemakerFeatureGroupState struct {
	Arn                         string                                          `json:"arn"`
	Description                 string                                          `json:"description"`
	EventTimeFeatureName        string                                          `json:"event_time_feature_name"`
	FeatureGroupName            string                                          `json:"feature_group_name"`
	Id                          string                                          `json:"id"`
	RecordIdentifierFeatureName string                                          `json:"record_identifier_feature_name"`
	RoleArn                     string                                          `json:"role_arn"`
	Tags                        map[string]string                               `json:"tags"`
	TagsAll                     map[string]string                               `json:"tags_all"`
	FeatureDefinition           []sagemakerfeaturegroup.FeatureDefinitionState  `json:"feature_definition"`
	OfflineStoreConfig          []sagemakerfeaturegroup.OfflineStoreConfigState `json:"offline_store_config"`
	OnlineStoreConfig           []sagemakerfeaturegroup.OnlineStoreConfigState  `json:"online_store_config"`
}
