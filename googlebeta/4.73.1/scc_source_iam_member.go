// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	sccsourceiammember "github.com/golingon/terraproviders/googlebeta/4.73.1/sccsourceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSccSourceIamMember creates a new instance of [SccSourceIamMember].
func NewSccSourceIamMember(name string, args SccSourceIamMemberArgs) *SccSourceIamMember {
	return &SccSourceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SccSourceIamMember)(nil)

// SccSourceIamMember represents the Terraform resource google_scc_source_iam_member.
type SccSourceIamMember struct {
	Name      string
	Args      SccSourceIamMemberArgs
	state     *sccSourceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SccSourceIamMember].
func (ssim *SccSourceIamMember) Type() string {
	return "google_scc_source_iam_member"
}

// LocalName returns the local name for [SccSourceIamMember].
func (ssim *SccSourceIamMember) LocalName() string {
	return ssim.Name
}

// Configuration returns the configuration (args) for [SccSourceIamMember].
func (ssim *SccSourceIamMember) Configuration() interface{} {
	return ssim.Args
}

// DependOn is used for other resources to depend on [SccSourceIamMember].
func (ssim *SccSourceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(ssim)
}

// Dependencies returns the list of resources [SccSourceIamMember] depends_on.
func (ssim *SccSourceIamMember) Dependencies() terra.Dependencies {
	return ssim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SccSourceIamMember].
func (ssim *SccSourceIamMember) LifecycleManagement() *terra.Lifecycle {
	return ssim.Lifecycle
}

// Attributes returns the attributes for [SccSourceIamMember].
func (ssim *SccSourceIamMember) Attributes() sccSourceIamMemberAttributes {
	return sccSourceIamMemberAttributes{ref: terra.ReferenceResource(ssim)}
}

// ImportState imports the given attribute values into [SccSourceIamMember]'s state.
func (ssim *SccSourceIamMember) ImportState(av io.Reader) error {
	ssim.state = &sccSourceIamMemberState{}
	if err := json.NewDecoder(av).Decode(ssim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssim.Type(), ssim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SccSourceIamMember] has state.
func (ssim *SccSourceIamMember) State() (*sccSourceIamMemberState, bool) {
	return ssim.state, ssim.state != nil
}

// StateMust returns the state for [SccSourceIamMember]. Panics if the state is nil.
func (ssim *SccSourceIamMember) StateMust() *sccSourceIamMemberState {
	if ssim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssim.Type(), ssim.LocalName()))
	}
	return ssim.state
}

// SccSourceIamMemberArgs contains the configurations for google_scc_source_iam_member.
type SccSourceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Organization: string, required
	Organization terra.StringValue `hcl:"organization,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// Condition: optional
	Condition *sccsourceiammember.Condition `hcl:"condition,block"`
}
type sccSourceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_scc_source_iam_member.
func (ssim sccSourceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ssim.ref.Append("etag"))
}

// Id returns a reference to field id of google_scc_source_iam_member.
func (ssim sccSourceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssim.ref.Append("id"))
}

// Member returns a reference to field member of google_scc_source_iam_member.
func (ssim sccSourceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(ssim.ref.Append("member"))
}

// Organization returns a reference to field organization of google_scc_source_iam_member.
func (ssim sccSourceIamMemberAttributes) Organization() terra.StringValue {
	return terra.ReferenceAsString(ssim.ref.Append("organization"))
}

// Role returns a reference to field role of google_scc_source_iam_member.
func (ssim sccSourceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ssim.ref.Append("role"))
}

// Source returns a reference to field source of google_scc_source_iam_member.
func (ssim sccSourceIamMemberAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(ssim.ref.Append("source"))
}

func (ssim sccSourceIamMemberAttributes) Condition() terra.ListValue[sccsourceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[sccsourceiammember.ConditionAttributes](ssim.ref.Append("condition"))
}

type sccSourceIamMemberState struct {
	Etag         string                              `json:"etag"`
	Id           string                              `json:"id"`
	Member       string                              `json:"member"`
	Organization string                              `json:"organization"`
	Role         string                              `json:"role"`
	Source       string                              `json:"source"`
	Condition    []sccsourceiammember.ConditionState `json:"condition"`
}
