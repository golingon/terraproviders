// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimesdkvoicevoiceprofiledomain "github.com/golingon/terraproviders/aws/5.7.0/chimesdkvoicevoiceprofiledomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimesdkvoiceVoiceProfileDomain creates a new instance of [ChimesdkvoiceVoiceProfileDomain].
func NewChimesdkvoiceVoiceProfileDomain(name string, args ChimesdkvoiceVoiceProfileDomainArgs) *ChimesdkvoiceVoiceProfileDomain {
	return &ChimesdkvoiceVoiceProfileDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimesdkvoiceVoiceProfileDomain)(nil)

// ChimesdkvoiceVoiceProfileDomain represents the Terraform resource aws_chimesdkvoice_voice_profile_domain.
type ChimesdkvoiceVoiceProfileDomain struct {
	Name      string
	Args      ChimesdkvoiceVoiceProfileDomainArgs
	state     *chimesdkvoiceVoiceProfileDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimesdkvoiceVoiceProfileDomain].
func (cvpd *ChimesdkvoiceVoiceProfileDomain) Type() string {
	return "aws_chimesdkvoice_voice_profile_domain"
}

// LocalName returns the local name for [ChimesdkvoiceVoiceProfileDomain].
func (cvpd *ChimesdkvoiceVoiceProfileDomain) LocalName() string {
	return cvpd.Name
}

// Configuration returns the configuration (args) for [ChimesdkvoiceVoiceProfileDomain].
func (cvpd *ChimesdkvoiceVoiceProfileDomain) Configuration() interface{} {
	return cvpd.Args
}

// DependOn is used for other resources to depend on [ChimesdkvoiceVoiceProfileDomain].
func (cvpd *ChimesdkvoiceVoiceProfileDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cvpd)
}

// Dependencies returns the list of resources [ChimesdkvoiceVoiceProfileDomain] depends_on.
func (cvpd *ChimesdkvoiceVoiceProfileDomain) Dependencies() terra.Dependencies {
	return cvpd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimesdkvoiceVoiceProfileDomain].
func (cvpd *ChimesdkvoiceVoiceProfileDomain) LifecycleManagement() *terra.Lifecycle {
	return cvpd.Lifecycle
}

// Attributes returns the attributes for [ChimesdkvoiceVoiceProfileDomain].
func (cvpd *ChimesdkvoiceVoiceProfileDomain) Attributes() chimesdkvoiceVoiceProfileDomainAttributes {
	return chimesdkvoiceVoiceProfileDomainAttributes{ref: terra.ReferenceResource(cvpd)}
}

// ImportState imports the given attribute values into [ChimesdkvoiceVoiceProfileDomain]'s state.
func (cvpd *ChimesdkvoiceVoiceProfileDomain) ImportState(av io.Reader) error {
	cvpd.state = &chimesdkvoiceVoiceProfileDomainState{}
	if err := json.NewDecoder(av).Decode(cvpd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cvpd.Type(), cvpd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimesdkvoiceVoiceProfileDomain] has state.
func (cvpd *ChimesdkvoiceVoiceProfileDomain) State() (*chimesdkvoiceVoiceProfileDomainState, bool) {
	return cvpd.state, cvpd.state != nil
}

// StateMust returns the state for [ChimesdkvoiceVoiceProfileDomain]. Panics if the state is nil.
func (cvpd *ChimesdkvoiceVoiceProfileDomain) StateMust() *chimesdkvoiceVoiceProfileDomainState {
	if cvpd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cvpd.Type(), cvpd.LocalName()))
	}
	return cvpd.state
}

// ChimesdkvoiceVoiceProfileDomainArgs contains the configurations for aws_chimesdkvoice_voice_profile_domain.
type ChimesdkvoiceVoiceProfileDomainArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ServerSideEncryptionConfiguration: required
	ServerSideEncryptionConfiguration *chimesdkvoicevoiceprofiledomain.ServerSideEncryptionConfiguration `hcl:"server_side_encryption_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *chimesdkvoicevoiceprofiledomain.Timeouts `hcl:"timeouts,block"`
}
type chimesdkvoiceVoiceProfileDomainAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_chimesdkvoice_voice_profile_domain.
func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cvpd.ref.Append("arn"))
}

// Description returns a reference to field description of aws_chimesdkvoice_voice_profile_domain.
func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cvpd.ref.Append("description"))
}

// Id returns a reference to field id of aws_chimesdkvoice_voice_profile_domain.
func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cvpd.ref.Append("id"))
}

// Name returns a reference to field name of aws_chimesdkvoice_voice_profile_domain.
func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cvpd.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_chimesdkvoice_voice_profile_domain.
func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cvpd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_chimesdkvoice_voice_profile_domain.
func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cvpd.ref.Append("tags_all"))
}

func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) ServerSideEncryptionConfiguration() terra.ListValue[chimesdkvoicevoiceprofiledomain.ServerSideEncryptionConfigurationAttributes] {
	return terra.ReferenceAsList[chimesdkvoicevoiceprofiledomain.ServerSideEncryptionConfigurationAttributes](cvpd.ref.Append("server_side_encryption_configuration"))
}

func (cvpd chimesdkvoiceVoiceProfileDomainAttributes) Timeouts() chimesdkvoicevoiceprofiledomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[chimesdkvoicevoiceprofiledomain.TimeoutsAttributes](cvpd.ref.Append("timeouts"))
}

type chimesdkvoiceVoiceProfileDomainState struct {
	Arn                               string                                                                   `json:"arn"`
	Description                       string                                                                   `json:"description"`
	Id                                string                                                                   `json:"id"`
	Name                              string                                                                   `json:"name"`
	Tags                              map[string]string                                                        `json:"tags"`
	TagsAll                           map[string]string                                                        `json:"tags_all"`
	ServerSideEncryptionConfiguration []chimesdkvoicevoiceprofiledomain.ServerSideEncryptionConfigurationState `json:"server_side_encryption_configuration"`
	Timeouts                          *chimesdkvoicevoiceprofiledomain.TimeoutsState                           `json:"timeouts"`
}
