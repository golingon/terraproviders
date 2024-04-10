// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewRedshiftUsageLimit creates a new instance of [RedshiftUsageLimit].
func NewRedshiftUsageLimit(name string, args RedshiftUsageLimitArgs) *RedshiftUsageLimit {
	return &RedshiftUsageLimit{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftUsageLimit)(nil)

// RedshiftUsageLimit represents the Terraform resource aws_redshift_usage_limit.
type RedshiftUsageLimit struct {
	Name      string
	Args      RedshiftUsageLimitArgs
	state     *redshiftUsageLimitState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftUsageLimit].
func (rul *RedshiftUsageLimit) Type() string {
	return "aws_redshift_usage_limit"
}

// LocalName returns the local name for [RedshiftUsageLimit].
func (rul *RedshiftUsageLimit) LocalName() string {
	return rul.Name
}

// Configuration returns the configuration (args) for [RedshiftUsageLimit].
func (rul *RedshiftUsageLimit) Configuration() interface{} {
	return rul.Args
}

// DependOn is used for other resources to depend on [RedshiftUsageLimit].
func (rul *RedshiftUsageLimit) DependOn() terra.Reference {
	return terra.ReferenceResource(rul)
}

// Dependencies returns the list of resources [RedshiftUsageLimit] depends_on.
func (rul *RedshiftUsageLimit) Dependencies() terra.Dependencies {
	return rul.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftUsageLimit].
func (rul *RedshiftUsageLimit) LifecycleManagement() *terra.Lifecycle {
	return rul.Lifecycle
}

// Attributes returns the attributes for [RedshiftUsageLimit].
func (rul *RedshiftUsageLimit) Attributes() redshiftUsageLimitAttributes {
	return redshiftUsageLimitAttributes{ref: terra.ReferenceResource(rul)}
}

// ImportState imports the given attribute values into [RedshiftUsageLimit]'s state.
func (rul *RedshiftUsageLimit) ImportState(av io.Reader) error {
	rul.state = &redshiftUsageLimitState{}
	if err := json.NewDecoder(av).Decode(rul.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rul.Type(), rul.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftUsageLimit] has state.
func (rul *RedshiftUsageLimit) State() (*redshiftUsageLimitState, bool) {
	return rul.state, rul.state != nil
}

// StateMust returns the state for [RedshiftUsageLimit]. Panics if the state is nil.
func (rul *RedshiftUsageLimit) StateMust() *redshiftUsageLimitState {
	if rul.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rul.Type(), rul.LocalName()))
	}
	return rul.state
}

// RedshiftUsageLimitArgs contains the configurations for aws_redshift_usage_limit.
type RedshiftUsageLimitArgs struct {
	// Amount: number, required
	Amount terra.NumberValue `hcl:"amount,attr" validate:"required"`
	// BreachAction: string, optional
	BreachAction terra.StringValue `hcl:"breach_action,attr"`
	// ClusterIdentifier: string, required
	ClusterIdentifier terra.StringValue `hcl:"cluster_identifier,attr" validate:"required"`
	// FeatureType: string, required
	FeatureType terra.StringValue `hcl:"feature_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LimitType: string, required
	LimitType terra.StringValue `hcl:"limit_type,attr" validate:"required"`
	// Period: string, optional
	Period terra.StringValue `hcl:"period,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type redshiftUsageLimitAttributes struct {
	ref terra.Reference
}

// Amount returns a reference to field amount of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) Amount() terra.NumberValue {
	return terra.ReferenceAsNumber(rul.ref.Append("amount"))
}

// Arn returns a reference to field arn of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("arn"))
}

// BreachAction returns a reference to field breach_action of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) BreachAction() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("breach_action"))
}

// ClusterIdentifier returns a reference to field cluster_identifier of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) ClusterIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("cluster_identifier"))
}

// FeatureType returns a reference to field feature_type of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) FeatureType() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("feature_type"))
}

// Id returns a reference to field id of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("id"))
}

// LimitType returns a reference to field limit_type of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) LimitType() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("limit_type"))
}

// Period returns a reference to field period of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) Period() terra.StringValue {
	return terra.ReferenceAsString(rul.ref.Append("period"))
}

// Tags returns a reference to field tags of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rul.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_usage_limit.
func (rul redshiftUsageLimitAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rul.ref.Append("tags_all"))
}

type redshiftUsageLimitState struct {
	Amount            float64           `json:"amount"`
	Arn               string            `json:"arn"`
	BreachAction      string            `json:"breach_action"`
	ClusterIdentifier string            `json:"cluster_identifier"`
	FeatureType       string            `json:"feature_type"`
	Id                string            `json:"id"`
	LimitType         string            `json:"limit_type"`
	Period            string            `json:"period"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
}
