// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3controlmultiregionaccesspointpolicy "github.com/golingon/terraproviders/aws/5.0.1/s3controlmultiregionaccesspointpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3ControlMultiRegionAccessPointPolicy creates a new instance of [S3ControlMultiRegionAccessPointPolicy].
func NewS3ControlMultiRegionAccessPointPolicy(name string, args S3ControlMultiRegionAccessPointPolicyArgs) *S3ControlMultiRegionAccessPointPolicy {
	return &S3ControlMultiRegionAccessPointPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3ControlMultiRegionAccessPointPolicy)(nil)

// S3ControlMultiRegionAccessPointPolicy represents the Terraform resource aws_s3control_multi_region_access_point_policy.
type S3ControlMultiRegionAccessPointPolicy struct {
	Name      string
	Args      S3ControlMultiRegionAccessPointPolicyArgs
	state     *s3ControlMultiRegionAccessPointPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3ControlMultiRegionAccessPointPolicy].
func (smrapp *S3ControlMultiRegionAccessPointPolicy) Type() string {
	return "aws_s3control_multi_region_access_point_policy"
}

// LocalName returns the local name for [S3ControlMultiRegionAccessPointPolicy].
func (smrapp *S3ControlMultiRegionAccessPointPolicy) LocalName() string {
	return smrapp.Name
}

// Configuration returns the configuration (args) for [S3ControlMultiRegionAccessPointPolicy].
func (smrapp *S3ControlMultiRegionAccessPointPolicy) Configuration() interface{} {
	return smrapp.Args
}

// DependOn is used for other resources to depend on [S3ControlMultiRegionAccessPointPolicy].
func (smrapp *S3ControlMultiRegionAccessPointPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(smrapp)
}

// Dependencies returns the list of resources [S3ControlMultiRegionAccessPointPolicy] depends_on.
func (smrapp *S3ControlMultiRegionAccessPointPolicy) Dependencies() terra.Dependencies {
	return smrapp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3ControlMultiRegionAccessPointPolicy].
func (smrapp *S3ControlMultiRegionAccessPointPolicy) LifecycleManagement() *terra.Lifecycle {
	return smrapp.Lifecycle
}

// Attributes returns the attributes for [S3ControlMultiRegionAccessPointPolicy].
func (smrapp *S3ControlMultiRegionAccessPointPolicy) Attributes() s3ControlMultiRegionAccessPointPolicyAttributes {
	return s3ControlMultiRegionAccessPointPolicyAttributes{ref: terra.ReferenceResource(smrapp)}
}

// ImportState imports the given attribute values into [S3ControlMultiRegionAccessPointPolicy]'s state.
func (smrapp *S3ControlMultiRegionAccessPointPolicy) ImportState(av io.Reader) error {
	smrapp.state = &s3ControlMultiRegionAccessPointPolicyState{}
	if err := json.NewDecoder(av).Decode(smrapp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smrapp.Type(), smrapp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3ControlMultiRegionAccessPointPolicy] has state.
func (smrapp *S3ControlMultiRegionAccessPointPolicy) State() (*s3ControlMultiRegionAccessPointPolicyState, bool) {
	return smrapp.state, smrapp.state != nil
}

// StateMust returns the state for [S3ControlMultiRegionAccessPointPolicy]. Panics if the state is nil.
func (smrapp *S3ControlMultiRegionAccessPointPolicy) StateMust() *s3ControlMultiRegionAccessPointPolicyState {
	if smrapp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smrapp.Type(), smrapp.LocalName()))
	}
	return smrapp.state
}

// S3ControlMultiRegionAccessPointPolicyArgs contains the configurations for aws_s3control_multi_region_access_point_policy.
type S3ControlMultiRegionAccessPointPolicyArgs struct {
	// AccountId: string, optional
	AccountId terra.StringValue `hcl:"account_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Details: required
	Details *s3controlmultiregionaccesspointpolicy.Details `hcl:"details,block" validate:"required"`
	// Timeouts: optional
	Timeouts *s3controlmultiregionaccesspointpolicy.Timeouts `hcl:"timeouts,block"`
}
type s3ControlMultiRegionAccessPointPolicyAttributes struct {
	ref terra.Reference
}

// AccountId returns a reference to field account_id of aws_s3control_multi_region_access_point_policy.
func (smrapp s3ControlMultiRegionAccessPointPolicyAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(smrapp.ref.Append("account_id"))
}

// Established returns a reference to field established of aws_s3control_multi_region_access_point_policy.
func (smrapp s3ControlMultiRegionAccessPointPolicyAttributes) Established() terra.StringValue {
	return terra.ReferenceAsString(smrapp.ref.Append("established"))
}

// Id returns a reference to field id of aws_s3control_multi_region_access_point_policy.
func (smrapp s3ControlMultiRegionAccessPointPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(smrapp.ref.Append("id"))
}

// Proposed returns a reference to field proposed of aws_s3control_multi_region_access_point_policy.
func (smrapp s3ControlMultiRegionAccessPointPolicyAttributes) Proposed() terra.StringValue {
	return terra.ReferenceAsString(smrapp.ref.Append("proposed"))
}

func (smrapp s3ControlMultiRegionAccessPointPolicyAttributes) Details() terra.ListValue[s3controlmultiregionaccesspointpolicy.DetailsAttributes] {
	return terra.ReferenceAsList[s3controlmultiregionaccesspointpolicy.DetailsAttributes](smrapp.ref.Append("details"))
}

func (smrapp s3ControlMultiRegionAccessPointPolicyAttributes) Timeouts() s3controlmultiregionaccesspointpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[s3controlmultiregionaccesspointpolicy.TimeoutsAttributes](smrapp.ref.Append("timeouts"))
}

type s3ControlMultiRegionAccessPointPolicyState struct {
	AccountId   string                                               `json:"account_id"`
	Established string                                               `json:"established"`
	Id          string                                               `json:"id"`
	Proposed    string                                               `json:"proposed"`
	Details     []s3controlmultiregionaccesspointpolicy.DetailsState `json:"details"`
	Timeouts    *s3controlmultiregionaccesspointpolicy.TimeoutsState `json:"timeouts"`
}
