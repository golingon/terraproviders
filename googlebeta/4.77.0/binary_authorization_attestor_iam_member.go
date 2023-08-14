// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	binaryauthorizationattestoriammember "github.com/golingon/terraproviders/googlebeta/4.77.0/binaryauthorizationattestoriammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBinaryAuthorizationAttestorIamMember creates a new instance of [BinaryAuthorizationAttestorIamMember].
func NewBinaryAuthorizationAttestorIamMember(name string, args BinaryAuthorizationAttestorIamMemberArgs) *BinaryAuthorizationAttestorIamMember {
	return &BinaryAuthorizationAttestorIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BinaryAuthorizationAttestorIamMember)(nil)

// BinaryAuthorizationAttestorIamMember represents the Terraform resource google_binary_authorization_attestor_iam_member.
type BinaryAuthorizationAttestorIamMember struct {
	Name      string
	Args      BinaryAuthorizationAttestorIamMemberArgs
	state     *binaryAuthorizationAttestorIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BinaryAuthorizationAttestorIamMember].
func (baaim *BinaryAuthorizationAttestorIamMember) Type() string {
	return "google_binary_authorization_attestor_iam_member"
}

// LocalName returns the local name for [BinaryAuthorizationAttestorIamMember].
func (baaim *BinaryAuthorizationAttestorIamMember) LocalName() string {
	return baaim.Name
}

// Configuration returns the configuration (args) for [BinaryAuthorizationAttestorIamMember].
func (baaim *BinaryAuthorizationAttestorIamMember) Configuration() interface{} {
	return baaim.Args
}

// DependOn is used for other resources to depend on [BinaryAuthorizationAttestorIamMember].
func (baaim *BinaryAuthorizationAttestorIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(baaim)
}

// Dependencies returns the list of resources [BinaryAuthorizationAttestorIamMember] depends_on.
func (baaim *BinaryAuthorizationAttestorIamMember) Dependencies() terra.Dependencies {
	return baaim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BinaryAuthorizationAttestorIamMember].
func (baaim *BinaryAuthorizationAttestorIamMember) LifecycleManagement() *terra.Lifecycle {
	return baaim.Lifecycle
}

// Attributes returns the attributes for [BinaryAuthorizationAttestorIamMember].
func (baaim *BinaryAuthorizationAttestorIamMember) Attributes() binaryAuthorizationAttestorIamMemberAttributes {
	return binaryAuthorizationAttestorIamMemberAttributes{ref: terra.ReferenceResource(baaim)}
}

// ImportState imports the given attribute values into [BinaryAuthorizationAttestorIamMember]'s state.
func (baaim *BinaryAuthorizationAttestorIamMember) ImportState(av io.Reader) error {
	baaim.state = &binaryAuthorizationAttestorIamMemberState{}
	if err := json.NewDecoder(av).Decode(baaim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baaim.Type(), baaim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BinaryAuthorizationAttestorIamMember] has state.
func (baaim *BinaryAuthorizationAttestorIamMember) State() (*binaryAuthorizationAttestorIamMemberState, bool) {
	return baaim.state, baaim.state != nil
}

// StateMust returns the state for [BinaryAuthorizationAttestorIamMember]. Panics if the state is nil.
func (baaim *BinaryAuthorizationAttestorIamMember) StateMust() *binaryAuthorizationAttestorIamMemberState {
	if baaim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baaim.Type(), baaim.LocalName()))
	}
	return baaim.state
}

// BinaryAuthorizationAttestorIamMemberArgs contains the configurations for google_binary_authorization_attestor_iam_member.
type BinaryAuthorizationAttestorIamMemberArgs struct {
	// Attestor: string, required
	Attestor terra.StringValue `hcl:"attestor,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *binaryauthorizationattestoriammember.Condition `hcl:"condition,block"`
}
type binaryAuthorizationAttestorIamMemberAttributes struct {
	ref terra.Reference
}

// Attestor returns a reference to field attestor of google_binary_authorization_attestor_iam_member.
func (baaim binaryAuthorizationAttestorIamMemberAttributes) Attestor() terra.StringValue {
	return terra.ReferenceAsString(baaim.ref.Append("attestor"))
}

// Etag returns a reference to field etag of google_binary_authorization_attestor_iam_member.
func (baaim binaryAuthorizationAttestorIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baaim.ref.Append("etag"))
}

// Id returns a reference to field id of google_binary_authorization_attestor_iam_member.
func (baaim binaryAuthorizationAttestorIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baaim.ref.Append("id"))
}

// Member returns a reference to field member of google_binary_authorization_attestor_iam_member.
func (baaim binaryAuthorizationAttestorIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(baaim.ref.Append("member"))
}

// Project returns a reference to field project of google_binary_authorization_attestor_iam_member.
func (baaim binaryAuthorizationAttestorIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(baaim.ref.Append("project"))
}

// Role returns a reference to field role of google_binary_authorization_attestor_iam_member.
func (baaim binaryAuthorizationAttestorIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(baaim.ref.Append("role"))
}

func (baaim binaryAuthorizationAttestorIamMemberAttributes) Condition() terra.ListValue[binaryauthorizationattestoriammember.ConditionAttributes] {
	return terra.ReferenceAsList[binaryauthorizationattestoriammember.ConditionAttributes](baaim.ref.Append("condition"))
}

type binaryAuthorizationAttestorIamMemberState struct {
	Attestor  string                                                `json:"attestor"`
	Etag      string                                                `json:"etag"`
	Id        string                                                `json:"id"`
	Member    string                                                `json:"member"`
	Project   string                                                `json:"project"`
	Role      string                                                `json:"role"`
	Condition []binaryauthorizationattestoriammember.ConditionState `json:"condition"`
}
