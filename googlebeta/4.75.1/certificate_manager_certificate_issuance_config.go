// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	certificatemanagercertificateissuanceconfig "github.com/golingon/terraproviders/googlebeta/4.75.1/certificatemanagercertificateissuanceconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCertificateManagerCertificateIssuanceConfig creates a new instance of [CertificateManagerCertificateIssuanceConfig].
func NewCertificateManagerCertificateIssuanceConfig(name string, args CertificateManagerCertificateIssuanceConfigArgs) *CertificateManagerCertificateIssuanceConfig {
	return &CertificateManagerCertificateIssuanceConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertificateManagerCertificateIssuanceConfig)(nil)

// CertificateManagerCertificateIssuanceConfig represents the Terraform resource google_certificate_manager_certificate_issuance_config.
type CertificateManagerCertificateIssuanceConfig struct {
	Name      string
	Args      CertificateManagerCertificateIssuanceConfigArgs
	state     *certificateManagerCertificateIssuanceConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CertificateManagerCertificateIssuanceConfig].
func (cmcic *CertificateManagerCertificateIssuanceConfig) Type() string {
	return "google_certificate_manager_certificate_issuance_config"
}

// LocalName returns the local name for [CertificateManagerCertificateIssuanceConfig].
func (cmcic *CertificateManagerCertificateIssuanceConfig) LocalName() string {
	return cmcic.Name
}

// Configuration returns the configuration (args) for [CertificateManagerCertificateIssuanceConfig].
func (cmcic *CertificateManagerCertificateIssuanceConfig) Configuration() interface{} {
	return cmcic.Args
}

// DependOn is used for other resources to depend on [CertificateManagerCertificateIssuanceConfig].
func (cmcic *CertificateManagerCertificateIssuanceConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(cmcic)
}

// Dependencies returns the list of resources [CertificateManagerCertificateIssuanceConfig] depends_on.
func (cmcic *CertificateManagerCertificateIssuanceConfig) Dependencies() terra.Dependencies {
	return cmcic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CertificateManagerCertificateIssuanceConfig].
func (cmcic *CertificateManagerCertificateIssuanceConfig) LifecycleManagement() *terra.Lifecycle {
	return cmcic.Lifecycle
}

// Attributes returns the attributes for [CertificateManagerCertificateIssuanceConfig].
func (cmcic *CertificateManagerCertificateIssuanceConfig) Attributes() certificateManagerCertificateIssuanceConfigAttributes {
	return certificateManagerCertificateIssuanceConfigAttributes{ref: terra.ReferenceResource(cmcic)}
}

// ImportState imports the given attribute values into [CertificateManagerCertificateIssuanceConfig]'s state.
func (cmcic *CertificateManagerCertificateIssuanceConfig) ImportState(av io.Reader) error {
	cmcic.state = &certificateManagerCertificateIssuanceConfigState{}
	if err := json.NewDecoder(av).Decode(cmcic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmcic.Type(), cmcic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CertificateManagerCertificateIssuanceConfig] has state.
func (cmcic *CertificateManagerCertificateIssuanceConfig) State() (*certificateManagerCertificateIssuanceConfigState, bool) {
	return cmcic.state, cmcic.state != nil
}

// StateMust returns the state for [CertificateManagerCertificateIssuanceConfig]. Panics if the state is nil.
func (cmcic *CertificateManagerCertificateIssuanceConfig) StateMust() *certificateManagerCertificateIssuanceConfigState {
	if cmcic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmcic.Type(), cmcic.LocalName()))
	}
	return cmcic.state
}

// CertificateManagerCertificateIssuanceConfigArgs contains the configurations for google_certificate_manager_certificate_issuance_config.
type CertificateManagerCertificateIssuanceConfigArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyAlgorithm: string, required
	KeyAlgorithm terra.StringValue `hcl:"key_algorithm,attr" validate:"required"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Lifetime: string, required
	Lifetime terra.StringValue `hcl:"lifetime,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RotationWindowPercentage: number, required
	RotationWindowPercentage terra.NumberValue `hcl:"rotation_window_percentage,attr" validate:"required"`
	// CertificateAuthorityConfig: required
	CertificateAuthorityConfig *certificatemanagercertificateissuanceconfig.CertificateAuthorityConfig `hcl:"certificate_authority_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *certificatemanagercertificateissuanceconfig.Timeouts `hcl:"timeouts,block"`
}
type certificateManagerCertificateIssuanceConfigAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("create_time"))
}

// Description returns a reference to field description of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("description"))
}

// Id returns a reference to field id of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("id"))
}

// KeyAlgorithm returns a reference to field key_algorithm of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) KeyAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("key_algorithm"))
}

// Labels returns a reference to field labels of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmcic.ref.Append("labels"))
}

// Lifetime returns a reference to field lifetime of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Lifetime() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("lifetime"))
}

// Location returns a reference to field location of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("location"))
}

// Name returns a reference to field name of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("name"))
}

// Project returns a reference to field project of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("project"))
}

// RotationWindowPercentage returns a reference to field rotation_window_percentage of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) RotationWindowPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(cmcic.ref.Append("rotation_window_percentage"))
}

// UpdateTime returns a reference to field update_time of google_certificate_manager_certificate_issuance_config.
func (cmcic certificateManagerCertificateIssuanceConfigAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cmcic.ref.Append("update_time"))
}

func (cmcic certificateManagerCertificateIssuanceConfigAttributes) CertificateAuthorityConfig() terra.ListValue[certificatemanagercertificateissuanceconfig.CertificateAuthorityConfigAttributes] {
	return terra.ReferenceAsList[certificatemanagercertificateissuanceconfig.CertificateAuthorityConfigAttributes](cmcic.ref.Append("certificate_authority_config"))
}

func (cmcic certificateManagerCertificateIssuanceConfigAttributes) Timeouts() certificatemanagercertificateissuanceconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[certificatemanagercertificateissuanceconfig.TimeoutsAttributes](cmcic.ref.Append("timeouts"))
}

type certificateManagerCertificateIssuanceConfigState struct {
	CreateTime                 string                                                                        `json:"create_time"`
	Description                string                                                                        `json:"description"`
	Id                         string                                                                        `json:"id"`
	KeyAlgorithm               string                                                                        `json:"key_algorithm"`
	Labels                     map[string]string                                                             `json:"labels"`
	Lifetime                   string                                                                        `json:"lifetime"`
	Location                   string                                                                        `json:"location"`
	Name                       string                                                                        `json:"name"`
	Project                    string                                                                        `json:"project"`
	RotationWindowPercentage   float64                                                                       `json:"rotation_window_percentage"`
	UpdateTime                 string                                                                        `json:"update_time"`
	CertificateAuthorityConfig []certificatemanagercertificateissuanceconfig.CertificateAuthorityConfigState `json:"certificate_authority_config"`
	Timeouts                   *certificatemanagercertificateissuanceconfig.TimeoutsState                    `json:"timeouts"`
}
