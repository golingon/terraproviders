// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	kmscryptokeyiammember "github.com/golingon/terraproviders/googlebeta/4.64.0/kmscryptokeyiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKmsCryptoKeyIamMember creates a new instance of [KmsCryptoKeyIamMember].
func NewKmsCryptoKeyIamMember(name string, args KmsCryptoKeyIamMemberArgs) *KmsCryptoKeyIamMember {
	return &KmsCryptoKeyIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsCryptoKeyIamMember)(nil)

// KmsCryptoKeyIamMember represents the Terraform resource google_kms_crypto_key_iam_member.
type KmsCryptoKeyIamMember struct {
	Name      string
	Args      KmsCryptoKeyIamMemberArgs
	state     *kmsCryptoKeyIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsCryptoKeyIamMember].
func (kckim *KmsCryptoKeyIamMember) Type() string {
	return "google_kms_crypto_key_iam_member"
}

// LocalName returns the local name for [KmsCryptoKeyIamMember].
func (kckim *KmsCryptoKeyIamMember) LocalName() string {
	return kckim.Name
}

// Configuration returns the configuration (args) for [KmsCryptoKeyIamMember].
func (kckim *KmsCryptoKeyIamMember) Configuration() interface{} {
	return kckim.Args
}

// DependOn is used for other resources to depend on [KmsCryptoKeyIamMember].
func (kckim *KmsCryptoKeyIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(kckim)
}

// Dependencies returns the list of resources [KmsCryptoKeyIamMember] depends_on.
func (kckim *KmsCryptoKeyIamMember) Dependencies() terra.Dependencies {
	return kckim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsCryptoKeyIamMember].
func (kckim *KmsCryptoKeyIamMember) LifecycleManagement() *terra.Lifecycle {
	return kckim.Lifecycle
}

// Attributes returns the attributes for [KmsCryptoKeyIamMember].
func (kckim *KmsCryptoKeyIamMember) Attributes() kmsCryptoKeyIamMemberAttributes {
	return kmsCryptoKeyIamMemberAttributes{ref: terra.ReferenceResource(kckim)}
}

// ImportState imports the given attribute values into [KmsCryptoKeyIamMember]'s state.
func (kckim *KmsCryptoKeyIamMember) ImportState(av io.Reader) error {
	kckim.state = &kmsCryptoKeyIamMemberState{}
	if err := json.NewDecoder(av).Decode(kckim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kckim.Type(), kckim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsCryptoKeyIamMember] has state.
func (kckim *KmsCryptoKeyIamMember) State() (*kmsCryptoKeyIamMemberState, bool) {
	return kckim.state, kckim.state != nil
}

// StateMust returns the state for [KmsCryptoKeyIamMember]. Panics if the state is nil.
func (kckim *KmsCryptoKeyIamMember) StateMust() *kmsCryptoKeyIamMemberState {
	if kckim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kckim.Type(), kckim.LocalName()))
	}
	return kckim.state
}

// KmsCryptoKeyIamMemberArgs contains the configurations for google_kms_crypto_key_iam_member.
type KmsCryptoKeyIamMemberArgs struct {
	// CryptoKeyId: string, required
	CryptoKeyId terra.StringValue `hcl:"crypto_key_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *kmscryptokeyiammember.Condition `hcl:"condition,block"`
}
type kmsCryptoKeyIamMemberAttributes struct {
	ref terra.Reference
}

// CryptoKeyId returns a reference to field crypto_key_id of google_kms_crypto_key_iam_member.
func (kckim kmsCryptoKeyIamMemberAttributes) CryptoKeyId() terra.StringValue {
	return terra.ReferenceAsString(kckim.ref.Append("crypto_key_id"))
}

// Etag returns a reference to field etag of google_kms_crypto_key_iam_member.
func (kckim kmsCryptoKeyIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(kckim.ref.Append("etag"))
}

// Id returns a reference to field id of google_kms_crypto_key_iam_member.
func (kckim kmsCryptoKeyIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kckim.ref.Append("id"))
}

// Member returns a reference to field member of google_kms_crypto_key_iam_member.
func (kckim kmsCryptoKeyIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(kckim.ref.Append("member"))
}

// Role returns a reference to field role of google_kms_crypto_key_iam_member.
func (kckim kmsCryptoKeyIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(kckim.ref.Append("role"))
}

func (kckim kmsCryptoKeyIamMemberAttributes) Condition() terra.ListValue[kmscryptokeyiammember.ConditionAttributes] {
	return terra.ReferenceAsList[kmscryptokeyiammember.ConditionAttributes](kckim.ref.Append("condition"))
}

type kmsCryptoKeyIamMemberState struct {
	CryptoKeyId string                                 `json:"crypto_key_id"`
	Etag        string                                 `json:"etag"`
	Id          string                                 `json:"id"`
	Member      string                                 `json:"member"`
	Role        string                                 `json:"role"`
	Condition   []kmscryptokeyiammember.ConditionState `json:"condition"`
}
