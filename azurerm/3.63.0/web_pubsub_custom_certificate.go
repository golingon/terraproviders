// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	webpubsubcustomcertificate "github.com/golingon/terraproviders/azurerm/3.63.0/webpubsubcustomcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWebPubsubCustomCertificate creates a new instance of [WebPubsubCustomCertificate].
func NewWebPubsubCustomCertificate(name string, args WebPubsubCustomCertificateArgs) *WebPubsubCustomCertificate {
	return &WebPubsubCustomCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WebPubsubCustomCertificate)(nil)

// WebPubsubCustomCertificate represents the Terraform resource azurerm_web_pubsub_custom_certificate.
type WebPubsubCustomCertificate struct {
	Name      string
	Args      WebPubsubCustomCertificateArgs
	state     *webPubsubCustomCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WebPubsubCustomCertificate].
func (wpcc *WebPubsubCustomCertificate) Type() string {
	return "azurerm_web_pubsub_custom_certificate"
}

// LocalName returns the local name for [WebPubsubCustomCertificate].
func (wpcc *WebPubsubCustomCertificate) LocalName() string {
	return wpcc.Name
}

// Configuration returns the configuration (args) for [WebPubsubCustomCertificate].
func (wpcc *WebPubsubCustomCertificate) Configuration() interface{} {
	return wpcc.Args
}

// DependOn is used for other resources to depend on [WebPubsubCustomCertificate].
func (wpcc *WebPubsubCustomCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(wpcc)
}

// Dependencies returns the list of resources [WebPubsubCustomCertificate] depends_on.
func (wpcc *WebPubsubCustomCertificate) Dependencies() terra.Dependencies {
	return wpcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WebPubsubCustomCertificate].
func (wpcc *WebPubsubCustomCertificate) LifecycleManagement() *terra.Lifecycle {
	return wpcc.Lifecycle
}

// Attributes returns the attributes for [WebPubsubCustomCertificate].
func (wpcc *WebPubsubCustomCertificate) Attributes() webPubsubCustomCertificateAttributes {
	return webPubsubCustomCertificateAttributes{ref: terra.ReferenceResource(wpcc)}
}

// ImportState imports the given attribute values into [WebPubsubCustomCertificate]'s state.
func (wpcc *WebPubsubCustomCertificate) ImportState(av io.Reader) error {
	wpcc.state = &webPubsubCustomCertificateState{}
	if err := json.NewDecoder(av).Decode(wpcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wpcc.Type(), wpcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WebPubsubCustomCertificate] has state.
func (wpcc *WebPubsubCustomCertificate) State() (*webPubsubCustomCertificateState, bool) {
	return wpcc.state, wpcc.state != nil
}

// StateMust returns the state for [WebPubsubCustomCertificate]. Panics if the state is nil.
func (wpcc *WebPubsubCustomCertificate) StateMust() *webPubsubCustomCertificateState {
	if wpcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wpcc.Type(), wpcc.LocalName()))
	}
	return wpcc.state
}

// WebPubsubCustomCertificateArgs contains the configurations for azurerm_web_pubsub_custom_certificate.
type WebPubsubCustomCertificateArgs struct {
	// CustomCertificateId: string, required
	CustomCertificateId terra.StringValue `hcl:"custom_certificate_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WebPubsubId: string, required
	WebPubsubId terra.StringValue `hcl:"web_pubsub_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *webpubsubcustomcertificate.Timeouts `hcl:"timeouts,block"`
}
type webPubsubCustomCertificateAttributes struct {
	ref terra.Reference
}

// CertificateVersion returns a reference to field certificate_version of azurerm_web_pubsub_custom_certificate.
func (wpcc webPubsubCustomCertificateAttributes) CertificateVersion() terra.StringValue {
	return terra.ReferenceAsString(wpcc.ref.Append("certificate_version"))
}

// CustomCertificateId returns a reference to field custom_certificate_id of azurerm_web_pubsub_custom_certificate.
func (wpcc webPubsubCustomCertificateAttributes) CustomCertificateId() terra.StringValue {
	return terra.ReferenceAsString(wpcc.ref.Append("custom_certificate_id"))
}

// Id returns a reference to field id of azurerm_web_pubsub_custom_certificate.
func (wpcc webPubsubCustomCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wpcc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_web_pubsub_custom_certificate.
func (wpcc webPubsubCustomCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wpcc.ref.Append("name"))
}

// WebPubsubId returns a reference to field web_pubsub_id of azurerm_web_pubsub_custom_certificate.
func (wpcc webPubsubCustomCertificateAttributes) WebPubsubId() terra.StringValue {
	return terra.ReferenceAsString(wpcc.ref.Append("web_pubsub_id"))
}

func (wpcc webPubsubCustomCertificateAttributes) Timeouts() webpubsubcustomcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[webpubsubcustomcertificate.TimeoutsAttributes](wpcc.ref.Append("timeouts"))
}

type webPubsubCustomCertificateState struct {
	CertificateVersion  string                                    `json:"certificate_version"`
	CustomCertificateId string                                    `json:"custom_certificate_id"`
	Id                  string                                    `json:"id"`
	Name                string                                    `json:"name"`
	WebPubsubId         string                                    `json:"web_pubsub_id"`
	Timeouts            *webpubsubcustomcertificate.TimeoutsState `json:"timeouts"`
}
