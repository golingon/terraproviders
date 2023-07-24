// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	eksfargateprofile "github.com/golingon/terraproviders/aws/5.9.0/eksfargateprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEksFargateProfile creates a new instance of [EksFargateProfile].
func NewEksFargateProfile(name string, args EksFargateProfileArgs) *EksFargateProfile {
	return &EksFargateProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EksFargateProfile)(nil)

// EksFargateProfile represents the Terraform resource aws_eks_fargate_profile.
type EksFargateProfile struct {
	Name      string
	Args      EksFargateProfileArgs
	state     *eksFargateProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EksFargateProfile].
func (efp *EksFargateProfile) Type() string {
	return "aws_eks_fargate_profile"
}

// LocalName returns the local name for [EksFargateProfile].
func (efp *EksFargateProfile) LocalName() string {
	return efp.Name
}

// Configuration returns the configuration (args) for [EksFargateProfile].
func (efp *EksFargateProfile) Configuration() interface{} {
	return efp.Args
}

// DependOn is used for other resources to depend on [EksFargateProfile].
func (efp *EksFargateProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(efp)
}

// Dependencies returns the list of resources [EksFargateProfile] depends_on.
func (efp *EksFargateProfile) Dependencies() terra.Dependencies {
	return efp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EksFargateProfile].
func (efp *EksFargateProfile) LifecycleManagement() *terra.Lifecycle {
	return efp.Lifecycle
}

// Attributes returns the attributes for [EksFargateProfile].
func (efp *EksFargateProfile) Attributes() eksFargateProfileAttributes {
	return eksFargateProfileAttributes{ref: terra.ReferenceResource(efp)}
}

// ImportState imports the given attribute values into [EksFargateProfile]'s state.
func (efp *EksFargateProfile) ImportState(av io.Reader) error {
	efp.state = &eksFargateProfileState{}
	if err := json.NewDecoder(av).Decode(efp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", efp.Type(), efp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EksFargateProfile] has state.
func (efp *EksFargateProfile) State() (*eksFargateProfileState, bool) {
	return efp.state, efp.state != nil
}

// StateMust returns the state for [EksFargateProfile]. Panics if the state is nil.
func (efp *EksFargateProfile) StateMust() *eksFargateProfileState {
	if efp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", efp.Type(), efp.LocalName()))
	}
	return efp.state
}

// EksFargateProfileArgs contains the configurations for aws_eks_fargate_profile.
type EksFargateProfileArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// FargateProfileName: string, required
	FargateProfileName terra.StringValue `hcl:"fargate_profile_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PodExecutionRoleArn: string, required
	PodExecutionRoleArn terra.StringValue `hcl:"pod_execution_role_arn,attr" validate:"required"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Selector: min=1
	Selector []eksfargateprofile.Selector `hcl:"selector,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *eksfargateprofile.Timeouts `hcl:"timeouts,block"`
}
type eksFargateProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(efp.ref.Append("arn"))
}

// ClusterName returns a reference to field cluster_name of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(efp.ref.Append("cluster_name"))
}

// FargateProfileName returns a reference to field fargate_profile_name of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) FargateProfileName() terra.StringValue {
	return terra.ReferenceAsString(efp.ref.Append("fargate_profile_name"))
}

// Id returns a reference to field id of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(efp.ref.Append("id"))
}

// PodExecutionRoleArn returns a reference to field pod_execution_role_arn of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) PodExecutionRoleArn() terra.StringValue {
	return terra.ReferenceAsString(efp.ref.Append("pod_execution_role_arn"))
}

// Status returns a reference to field status of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(efp.ref.Append("status"))
}

// SubnetIds returns a reference to field subnet_ids of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](efp.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](efp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_eks_fargate_profile.
func (efp eksFargateProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](efp.ref.Append("tags_all"))
}

func (efp eksFargateProfileAttributes) Selector() terra.SetValue[eksfargateprofile.SelectorAttributes] {
	return terra.ReferenceAsSet[eksfargateprofile.SelectorAttributes](efp.ref.Append("selector"))
}

func (efp eksFargateProfileAttributes) Timeouts() eksfargateprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[eksfargateprofile.TimeoutsAttributes](efp.ref.Append("timeouts"))
}

type eksFargateProfileState struct {
	Arn                 string                            `json:"arn"`
	ClusterName         string                            `json:"cluster_name"`
	FargateProfileName  string                            `json:"fargate_profile_name"`
	Id                  string                            `json:"id"`
	PodExecutionRoleArn string                            `json:"pod_execution_role_arn"`
	Status              string                            `json:"status"`
	SubnetIds           []string                          `json:"subnet_ids"`
	Tags                map[string]string                 `json:"tags"`
	TagsAll             map[string]string                 `json:"tags_all"`
	Selector            []eksfargateprofile.SelectorState `json:"selector"`
	Timeouts            *eksfargateprofile.TimeoutsState  `json:"timeouts"`
}
