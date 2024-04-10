// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mediastreamingpolicy "github.com/golingon/terraproviders/azurerm/3.98.0/mediastreamingpolicy"
	"io"
)

// NewMediaStreamingPolicy creates a new instance of [MediaStreamingPolicy].
func NewMediaStreamingPolicy(name string, args MediaStreamingPolicyArgs) *MediaStreamingPolicy {
	return &MediaStreamingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MediaStreamingPolicy)(nil)

// MediaStreamingPolicy represents the Terraform resource azurerm_media_streaming_policy.
type MediaStreamingPolicy struct {
	Name      string
	Args      MediaStreamingPolicyArgs
	state     *mediaStreamingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MediaStreamingPolicy].
func (msp *MediaStreamingPolicy) Type() string {
	return "azurerm_media_streaming_policy"
}

// LocalName returns the local name for [MediaStreamingPolicy].
func (msp *MediaStreamingPolicy) LocalName() string {
	return msp.Name
}

// Configuration returns the configuration (args) for [MediaStreamingPolicy].
func (msp *MediaStreamingPolicy) Configuration() interface{} {
	return msp.Args
}

// DependOn is used for other resources to depend on [MediaStreamingPolicy].
func (msp *MediaStreamingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(msp)
}

// Dependencies returns the list of resources [MediaStreamingPolicy] depends_on.
func (msp *MediaStreamingPolicy) Dependencies() terra.Dependencies {
	return msp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MediaStreamingPolicy].
func (msp *MediaStreamingPolicy) LifecycleManagement() *terra.Lifecycle {
	return msp.Lifecycle
}

// Attributes returns the attributes for [MediaStreamingPolicy].
func (msp *MediaStreamingPolicy) Attributes() mediaStreamingPolicyAttributes {
	return mediaStreamingPolicyAttributes{ref: terra.ReferenceResource(msp)}
}

// ImportState imports the given attribute values into [MediaStreamingPolicy]'s state.
func (msp *MediaStreamingPolicy) ImportState(av io.Reader) error {
	msp.state = &mediaStreamingPolicyState{}
	if err := json.NewDecoder(av).Decode(msp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msp.Type(), msp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MediaStreamingPolicy] has state.
func (msp *MediaStreamingPolicy) State() (*mediaStreamingPolicyState, bool) {
	return msp.state, msp.state != nil
}

// StateMust returns the state for [MediaStreamingPolicy]. Panics if the state is nil.
func (msp *MediaStreamingPolicy) StateMust() *mediaStreamingPolicyState {
	if msp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msp.Type(), msp.LocalName()))
	}
	return msp.state
}

// MediaStreamingPolicyArgs contains the configurations for azurerm_media_streaming_policy.
type MediaStreamingPolicyArgs struct {
	// DefaultContentKeyPolicyName: string, optional
	DefaultContentKeyPolicyName terra.StringValue `hcl:"default_content_key_policy_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MediaServicesAccountName: string, required
	MediaServicesAccountName terra.StringValue `hcl:"media_services_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// CommonEncryptionCbcs: optional
	CommonEncryptionCbcs *mediastreamingpolicy.CommonEncryptionCbcs `hcl:"common_encryption_cbcs,block"`
	// CommonEncryptionCenc: optional
	CommonEncryptionCenc *mediastreamingpolicy.CommonEncryptionCenc `hcl:"common_encryption_cenc,block"`
	// EnvelopeEncryption: optional
	EnvelopeEncryption *mediastreamingpolicy.EnvelopeEncryption `hcl:"envelope_encryption,block"`
	// NoEncryptionEnabledProtocols: optional
	NoEncryptionEnabledProtocols *mediastreamingpolicy.NoEncryptionEnabledProtocols `hcl:"no_encryption_enabled_protocols,block"`
	// Timeouts: optional
	Timeouts *mediastreamingpolicy.Timeouts `hcl:"timeouts,block"`
}
type mediaStreamingPolicyAttributes struct {
	ref terra.Reference
}

// DefaultContentKeyPolicyName returns a reference to field default_content_key_policy_name of azurerm_media_streaming_policy.
func (msp mediaStreamingPolicyAttributes) DefaultContentKeyPolicyName() terra.StringValue {
	return terra.ReferenceAsString(msp.ref.Append("default_content_key_policy_name"))
}

// Id returns a reference to field id of azurerm_media_streaming_policy.
func (msp mediaStreamingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msp.ref.Append("id"))
}

// MediaServicesAccountName returns a reference to field media_services_account_name of azurerm_media_streaming_policy.
func (msp mediaStreamingPolicyAttributes) MediaServicesAccountName() terra.StringValue {
	return terra.ReferenceAsString(msp.ref.Append("media_services_account_name"))
}

// Name returns a reference to field name of azurerm_media_streaming_policy.
func (msp mediaStreamingPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(msp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_media_streaming_policy.
func (msp mediaStreamingPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(msp.ref.Append("resource_group_name"))
}

func (msp mediaStreamingPolicyAttributes) CommonEncryptionCbcs() terra.ListValue[mediastreamingpolicy.CommonEncryptionCbcsAttributes] {
	return terra.ReferenceAsList[mediastreamingpolicy.CommonEncryptionCbcsAttributes](msp.ref.Append("common_encryption_cbcs"))
}

func (msp mediaStreamingPolicyAttributes) CommonEncryptionCenc() terra.ListValue[mediastreamingpolicy.CommonEncryptionCencAttributes] {
	return terra.ReferenceAsList[mediastreamingpolicy.CommonEncryptionCencAttributes](msp.ref.Append("common_encryption_cenc"))
}

func (msp mediaStreamingPolicyAttributes) EnvelopeEncryption() terra.ListValue[mediastreamingpolicy.EnvelopeEncryptionAttributes] {
	return terra.ReferenceAsList[mediastreamingpolicy.EnvelopeEncryptionAttributes](msp.ref.Append("envelope_encryption"))
}

func (msp mediaStreamingPolicyAttributes) NoEncryptionEnabledProtocols() terra.ListValue[mediastreamingpolicy.NoEncryptionEnabledProtocolsAttributes] {
	return terra.ReferenceAsList[mediastreamingpolicy.NoEncryptionEnabledProtocolsAttributes](msp.ref.Append("no_encryption_enabled_protocols"))
}

func (msp mediaStreamingPolicyAttributes) Timeouts() mediastreamingpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mediastreamingpolicy.TimeoutsAttributes](msp.ref.Append("timeouts"))
}

type mediaStreamingPolicyState struct {
	DefaultContentKeyPolicyName  string                                                   `json:"default_content_key_policy_name"`
	Id                           string                                                   `json:"id"`
	MediaServicesAccountName     string                                                   `json:"media_services_account_name"`
	Name                         string                                                   `json:"name"`
	ResourceGroupName            string                                                   `json:"resource_group_name"`
	CommonEncryptionCbcs         []mediastreamingpolicy.CommonEncryptionCbcsState         `json:"common_encryption_cbcs"`
	CommonEncryptionCenc         []mediastreamingpolicy.CommonEncryptionCencState         `json:"common_encryption_cenc"`
	EnvelopeEncryption           []mediastreamingpolicy.EnvelopeEncryptionState           `json:"envelope_encryption"`
	NoEncryptionEnabledProtocols []mediastreamingpolicy.NoEncryptionEnabledProtocolsState `json:"no_encryption_enabled_protocols"`
	Timeouts                     *mediastreamingpolicy.TimeoutsState                      `json:"timeouts"`
}
