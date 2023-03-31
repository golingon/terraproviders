// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mediacontentkeypolicy "github.com/golingon/terraproviders/azurerm/3.49.0/mediacontentkeypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMediaContentKeyPolicy(name string, args MediaContentKeyPolicyArgs) *MediaContentKeyPolicy {
	return &MediaContentKeyPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaContentKeyPolicy)(nil)

type MediaContentKeyPolicy struct {
	Name  string
	Args  MediaContentKeyPolicyArgs
	state *mediaContentKeyPolicyState
}

func (mckp *MediaContentKeyPolicy) Type() string {
	return "azurerm_media_content_key_policy"
}

func (mckp *MediaContentKeyPolicy) LocalName() string {
	return mckp.Name
}

func (mckp *MediaContentKeyPolicy) Configuration() interface{} {
	return mckp.Args
}

func (mckp *MediaContentKeyPolicy) Attributes() mediaContentKeyPolicyAttributes {
	return mediaContentKeyPolicyAttributes{ref: terra.ReferenceResource(mckp)}
}

func (mckp *MediaContentKeyPolicy) ImportState(av io.Reader) error {
	mckp.state = &mediaContentKeyPolicyState{}
	if err := json.NewDecoder(av).Decode(mckp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mckp.Type(), mckp.LocalName(), err)
	}
	return nil
}

func (mckp *MediaContentKeyPolicy) State() (*mediaContentKeyPolicyState, bool) {
	return mckp.state, mckp.state != nil
}

func (mckp *MediaContentKeyPolicy) StateMust() *mediaContentKeyPolicyState {
	if mckp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mckp.Type(), mckp.LocalName()))
	}
	return mckp.state
}

func (mckp *MediaContentKeyPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(mckp)
}

type MediaContentKeyPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MediaServicesAccountName: string, required
	MediaServicesAccountName terra.StringValue `hcl:"media_services_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// PolicyOption: min=1
	PolicyOption []mediacontentkeypolicy.PolicyOption `hcl:"policy_option,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *mediacontentkeypolicy.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that MediaContentKeyPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type mediaContentKeyPolicyAttributes struct {
	ref terra.Reference
}

func (mckp mediaContentKeyPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceString(mckp.ref.Append("description"))
}

func (mckp mediaContentKeyPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mckp.ref.Append("id"))
}

func (mckp mediaContentKeyPolicyAttributes) MediaServicesAccountName() terra.StringValue {
	return terra.ReferenceString(mckp.ref.Append("media_services_account_name"))
}

func (mckp mediaContentKeyPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mckp.ref.Append("name"))
}

func (mckp mediaContentKeyPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mckp.ref.Append("resource_group_name"))
}

func (mckp mediaContentKeyPolicyAttributes) PolicyOption() terra.SetValue[mediacontentkeypolicy.PolicyOptionAttributes] {
	return terra.ReferenceSet[mediacontentkeypolicy.PolicyOptionAttributes](mckp.ref.Append("policy_option"))
}

func (mckp mediaContentKeyPolicyAttributes) Timeouts() mediacontentkeypolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[mediacontentkeypolicy.TimeoutsAttributes](mckp.ref.Append("timeouts"))
}

type mediaContentKeyPolicyState struct {
	Description              string                                    `json:"description"`
	Id                       string                                    `json:"id"`
	MediaServicesAccountName string                                    `json:"media_services_account_name"`
	Name                     string                                    `json:"name"`
	ResourceGroupName        string                                    `json:"resource_group_name"`
	PolicyOption             []mediacontentkeypolicy.PolicyOptionState `json:"policy_option"`
	Timeouts                 *mediacontentkeypolicy.TimeoutsState      `json:"timeouts"`
}
