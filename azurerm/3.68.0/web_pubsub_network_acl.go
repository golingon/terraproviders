// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	webpubsubnetworkacl "github.com/golingon/terraproviders/azurerm/3.68.0/webpubsubnetworkacl"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWebPubsubNetworkAcl creates a new instance of [WebPubsubNetworkAcl].
func NewWebPubsubNetworkAcl(name string, args WebPubsubNetworkAclArgs) *WebPubsubNetworkAcl {
	return &WebPubsubNetworkAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WebPubsubNetworkAcl)(nil)

// WebPubsubNetworkAcl represents the Terraform resource azurerm_web_pubsub_network_acl.
type WebPubsubNetworkAcl struct {
	Name      string
	Args      WebPubsubNetworkAclArgs
	state     *webPubsubNetworkAclState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WebPubsubNetworkAcl].
func (wpna *WebPubsubNetworkAcl) Type() string {
	return "azurerm_web_pubsub_network_acl"
}

// LocalName returns the local name for [WebPubsubNetworkAcl].
func (wpna *WebPubsubNetworkAcl) LocalName() string {
	return wpna.Name
}

// Configuration returns the configuration (args) for [WebPubsubNetworkAcl].
func (wpna *WebPubsubNetworkAcl) Configuration() interface{} {
	return wpna.Args
}

// DependOn is used for other resources to depend on [WebPubsubNetworkAcl].
func (wpna *WebPubsubNetworkAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(wpna)
}

// Dependencies returns the list of resources [WebPubsubNetworkAcl] depends_on.
func (wpna *WebPubsubNetworkAcl) Dependencies() terra.Dependencies {
	return wpna.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WebPubsubNetworkAcl].
func (wpna *WebPubsubNetworkAcl) LifecycleManagement() *terra.Lifecycle {
	return wpna.Lifecycle
}

// Attributes returns the attributes for [WebPubsubNetworkAcl].
func (wpna *WebPubsubNetworkAcl) Attributes() webPubsubNetworkAclAttributes {
	return webPubsubNetworkAclAttributes{ref: terra.ReferenceResource(wpna)}
}

// ImportState imports the given attribute values into [WebPubsubNetworkAcl]'s state.
func (wpna *WebPubsubNetworkAcl) ImportState(av io.Reader) error {
	wpna.state = &webPubsubNetworkAclState{}
	if err := json.NewDecoder(av).Decode(wpna.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wpna.Type(), wpna.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WebPubsubNetworkAcl] has state.
func (wpna *WebPubsubNetworkAcl) State() (*webPubsubNetworkAclState, bool) {
	return wpna.state, wpna.state != nil
}

// StateMust returns the state for [WebPubsubNetworkAcl]. Panics if the state is nil.
func (wpna *WebPubsubNetworkAcl) StateMust() *webPubsubNetworkAclState {
	if wpna.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wpna.Type(), wpna.LocalName()))
	}
	return wpna.state
}

// WebPubsubNetworkAclArgs contains the configurations for azurerm_web_pubsub_network_acl.
type WebPubsubNetworkAclArgs struct {
	// DefaultAction: string, optional
	DefaultAction terra.StringValue `hcl:"default_action,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// WebPubsubId: string, required
	WebPubsubId terra.StringValue `hcl:"web_pubsub_id,attr" validate:"required"`
	// PrivateEndpoint: min=0
	PrivateEndpoint []webpubsubnetworkacl.PrivateEndpoint `hcl:"private_endpoint,block" validate:"min=0"`
	// PublicNetwork: required
	PublicNetwork *webpubsubnetworkacl.PublicNetwork `hcl:"public_network,block" validate:"required"`
	// Timeouts: optional
	Timeouts *webpubsubnetworkacl.Timeouts `hcl:"timeouts,block"`
}
type webPubsubNetworkAclAttributes struct {
	ref terra.Reference
}

// DefaultAction returns a reference to field default_action of azurerm_web_pubsub_network_acl.
func (wpna webPubsubNetworkAclAttributes) DefaultAction() terra.StringValue {
	return terra.ReferenceAsString(wpna.ref.Append("default_action"))
}

// Id returns a reference to field id of azurerm_web_pubsub_network_acl.
func (wpna webPubsubNetworkAclAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wpna.ref.Append("id"))
}

// WebPubsubId returns a reference to field web_pubsub_id of azurerm_web_pubsub_network_acl.
func (wpna webPubsubNetworkAclAttributes) WebPubsubId() terra.StringValue {
	return terra.ReferenceAsString(wpna.ref.Append("web_pubsub_id"))
}

func (wpna webPubsubNetworkAclAttributes) PrivateEndpoint() terra.SetValue[webpubsubnetworkacl.PrivateEndpointAttributes] {
	return terra.ReferenceAsSet[webpubsubnetworkacl.PrivateEndpointAttributes](wpna.ref.Append("private_endpoint"))
}

func (wpna webPubsubNetworkAclAttributes) PublicNetwork() terra.ListValue[webpubsubnetworkacl.PublicNetworkAttributes] {
	return terra.ReferenceAsList[webpubsubnetworkacl.PublicNetworkAttributes](wpna.ref.Append("public_network"))
}

func (wpna webPubsubNetworkAclAttributes) Timeouts() webpubsubnetworkacl.TimeoutsAttributes {
	return terra.ReferenceAsSingle[webpubsubnetworkacl.TimeoutsAttributes](wpna.ref.Append("timeouts"))
}

type webPubsubNetworkAclState struct {
	DefaultAction   string                                     `json:"default_action"`
	Id              string                                     `json:"id"`
	WebPubsubId     string                                     `json:"web_pubsub_id"`
	PrivateEndpoint []webpubsubnetworkacl.PrivateEndpointState `json:"private_endpoint"`
	PublicNetwork   []webpubsubnetworkacl.PublicNetworkState   `json:"public_network"`
	Timeouts        *webpubsubnetworkacl.TimeoutsState         `json:"timeouts"`
}
