// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	signalrservicecustomcertificate "github.com/golingon/terraproviders/azurerm/3.98.0/signalrservicecustomcertificate"
	"io"
)

// NewSignalrServiceCustomCertificate creates a new instance of [SignalrServiceCustomCertificate].
func NewSignalrServiceCustomCertificate(name string, args SignalrServiceCustomCertificateArgs) *SignalrServiceCustomCertificate {
	return &SignalrServiceCustomCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SignalrServiceCustomCertificate)(nil)

// SignalrServiceCustomCertificate represents the Terraform resource azurerm_signalr_service_custom_certificate.
type SignalrServiceCustomCertificate struct {
	Name      string
	Args      SignalrServiceCustomCertificateArgs
	state     *signalrServiceCustomCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SignalrServiceCustomCertificate].
func (sscc *SignalrServiceCustomCertificate) Type() string {
	return "azurerm_signalr_service_custom_certificate"
}

// LocalName returns the local name for [SignalrServiceCustomCertificate].
func (sscc *SignalrServiceCustomCertificate) LocalName() string {
	return sscc.Name
}

// Configuration returns the configuration (args) for [SignalrServiceCustomCertificate].
func (sscc *SignalrServiceCustomCertificate) Configuration() interface{} {
	return sscc.Args
}

// DependOn is used for other resources to depend on [SignalrServiceCustomCertificate].
func (sscc *SignalrServiceCustomCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(sscc)
}

// Dependencies returns the list of resources [SignalrServiceCustomCertificate] depends_on.
func (sscc *SignalrServiceCustomCertificate) Dependencies() terra.Dependencies {
	return sscc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SignalrServiceCustomCertificate].
func (sscc *SignalrServiceCustomCertificate) LifecycleManagement() *terra.Lifecycle {
	return sscc.Lifecycle
}

// Attributes returns the attributes for [SignalrServiceCustomCertificate].
func (sscc *SignalrServiceCustomCertificate) Attributes() signalrServiceCustomCertificateAttributes {
	return signalrServiceCustomCertificateAttributes{ref: terra.ReferenceResource(sscc)}
}

// ImportState imports the given attribute values into [SignalrServiceCustomCertificate]'s state.
func (sscc *SignalrServiceCustomCertificate) ImportState(av io.Reader) error {
	sscc.state = &signalrServiceCustomCertificateState{}
	if err := json.NewDecoder(av).Decode(sscc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sscc.Type(), sscc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SignalrServiceCustomCertificate] has state.
func (sscc *SignalrServiceCustomCertificate) State() (*signalrServiceCustomCertificateState, bool) {
	return sscc.state, sscc.state != nil
}

// StateMust returns the state for [SignalrServiceCustomCertificate]. Panics if the state is nil.
func (sscc *SignalrServiceCustomCertificate) StateMust() *signalrServiceCustomCertificateState {
	if sscc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sscc.Type(), sscc.LocalName()))
	}
	return sscc.state
}

// SignalrServiceCustomCertificateArgs contains the configurations for azurerm_signalr_service_custom_certificate.
type SignalrServiceCustomCertificateArgs struct {
	// CustomCertificateId: string, required
	CustomCertificateId terra.StringValue `hcl:"custom_certificate_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SignalrServiceId: string, required
	SignalrServiceId terra.StringValue `hcl:"signalr_service_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *signalrservicecustomcertificate.Timeouts `hcl:"timeouts,block"`
}
type signalrServiceCustomCertificateAttributes struct {
	ref terra.Reference
}

// CertificateVersion returns a reference to field certificate_version of azurerm_signalr_service_custom_certificate.
func (sscc signalrServiceCustomCertificateAttributes) CertificateVersion() terra.StringValue {
	return terra.ReferenceAsString(sscc.ref.Append("certificate_version"))
}

// CustomCertificateId returns a reference to field custom_certificate_id of azurerm_signalr_service_custom_certificate.
func (sscc signalrServiceCustomCertificateAttributes) CustomCertificateId() terra.StringValue {
	return terra.ReferenceAsString(sscc.ref.Append("custom_certificate_id"))
}

// Id returns a reference to field id of azurerm_signalr_service_custom_certificate.
func (sscc signalrServiceCustomCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sscc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_signalr_service_custom_certificate.
func (sscc signalrServiceCustomCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sscc.ref.Append("name"))
}

// SignalrServiceId returns a reference to field signalr_service_id of azurerm_signalr_service_custom_certificate.
func (sscc signalrServiceCustomCertificateAttributes) SignalrServiceId() terra.StringValue {
	return terra.ReferenceAsString(sscc.ref.Append("signalr_service_id"))
}

func (sscc signalrServiceCustomCertificateAttributes) Timeouts() signalrservicecustomcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[signalrservicecustomcertificate.TimeoutsAttributes](sscc.ref.Append("timeouts"))
}

type signalrServiceCustomCertificateState struct {
	CertificateVersion  string                                         `json:"certificate_version"`
	CustomCertificateId string                                         `json:"custom_certificate_id"`
	Id                  string                                         `json:"id"`
	Name                string                                         `json:"name"`
	SignalrServiceId    string                                         `json:"signalr_service_id"`
	Timeouts            *signalrservicecustomcertificate.TimeoutsState `json:"timeouts"`
}
