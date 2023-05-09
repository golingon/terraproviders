// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	webpubsubhub "github.com/golingon/terraproviders/azurerm/3.55.0/webpubsubhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewWebPubsubHub creates a new instance of [WebPubsubHub].
func NewWebPubsubHub(name string, args WebPubsubHubArgs) *WebPubsubHub {
	return &WebPubsubHub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WebPubsubHub)(nil)

// WebPubsubHub represents the Terraform resource azurerm_web_pubsub_hub.
type WebPubsubHub struct {
	Name      string
	Args      WebPubsubHubArgs
	state     *webPubsubHubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WebPubsubHub].
func (wph *WebPubsubHub) Type() string {
	return "azurerm_web_pubsub_hub"
}

// LocalName returns the local name for [WebPubsubHub].
func (wph *WebPubsubHub) LocalName() string {
	return wph.Name
}

// Configuration returns the configuration (args) for [WebPubsubHub].
func (wph *WebPubsubHub) Configuration() interface{} {
	return wph.Args
}

// DependOn is used for other resources to depend on [WebPubsubHub].
func (wph *WebPubsubHub) DependOn() terra.Reference {
	return terra.ReferenceResource(wph)
}

// Dependencies returns the list of resources [WebPubsubHub] depends_on.
func (wph *WebPubsubHub) Dependencies() terra.Dependencies {
	return wph.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WebPubsubHub].
func (wph *WebPubsubHub) LifecycleManagement() *terra.Lifecycle {
	return wph.Lifecycle
}

// Attributes returns the attributes for [WebPubsubHub].
func (wph *WebPubsubHub) Attributes() webPubsubHubAttributes {
	return webPubsubHubAttributes{ref: terra.ReferenceResource(wph)}
}

// ImportState imports the given attribute values into [WebPubsubHub]'s state.
func (wph *WebPubsubHub) ImportState(av io.Reader) error {
	wph.state = &webPubsubHubState{}
	if err := json.NewDecoder(av).Decode(wph.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wph.Type(), wph.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WebPubsubHub] has state.
func (wph *WebPubsubHub) State() (*webPubsubHubState, bool) {
	return wph.state, wph.state != nil
}

// StateMust returns the state for [WebPubsubHub]. Panics if the state is nil.
func (wph *WebPubsubHub) StateMust() *webPubsubHubState {
	if wph.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wph.Type(), wph.LocalName()))
	}
	return wph.state
}

// WebPubsubHubArgs contains the configurations for azurerm_web_pubsub_hub.
type WebPubsubHubArgs struct {
	// AnonymousConnectionsEnabled: bool, optional
	AnonymousConnectionsEnabled terra.BoolValue `hcl:"anonymous_connections_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// WebPubsubId: string, required
	WebPubsubId terra.StringValue `hcl:"web_pubsub_id,attr" validate:"required"`
	// EventHandler: min=0
	EventHandler []webpubsubhub.EventHandler `hcl:"event_handler,block" validate:"min=0"`
	// EventListener: min=0
	EventListener []webpubsubhub.EventListener `hcl:"event_listener,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *webpubsubhub.Timeouts `hcl:"timeouts,block"`
}
type webPubsubHubAttributes struct {
	ref terra.Reference
}

// AnonymousConnectionsEnabled returns a reference to field anonymous_connections_enabled of azurerm_web_pubsub_hub.
func (wph webPubsubHubAttributes) AnonymousConnectionsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(wph.ref.Append("anonymous_connections_enabled"))
}

// Id returns a reference to field id of azurerm_web_pubsub_hub.
func (wph webPubsubHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wph.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_web_pubsub_hub.
func (wph webPubsubHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wph.ref.Append("name"))
}

// WebPubsubId returns a reference to field web_pubsub_id of azurerm_web_pubsub_hub.
func (wph webPubsubHubAttributes) WebPubsubId() terra.StringValue {
	return terra.ReferenceAsString(wph.ref.Append("web_pubsub_id"))
}

func (wph webPubsubHubAttributes) EventHandler() terra.ListValue[webpubsubhub.EventHandlerAttributes] {
	return terra.ReferenceAsList[webpubsubhub.EventHandlerAttributes](wph.ref.Append("event_handler"))
}

func (wph webPubsubHubAttributes) EventListener() terra.ListValue[webpubsubhub.EventListenerAttributes] {
	return terra.ReferenceAsList[webpubsubhub.EventListenerAttributes](wph.ref.Append("event_listener"))
}

func (wph webPubsubHubAttributes) Timeouts() webpubsubhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[webpubsubhub.TimeoutsAttributes](wph.ref.Append("timeouts"))
}

type webPubsubHubState struct {
	AnonymousConnectionsEnabled bool                              `json:"anonymous_connections_enabled"`
	Id                          string                            `json:"id"`
	Name                        string                            `json:"name"`
	WebPubsubId                 string                            `json:"web_pubsub_id"`
	EventHandler                []webpubsubhub.EventHandlerState  `json:"event_handler"`
	EventListener               []webpubsubhub.EventListenerState `json:"event_listener"`
	Timeouts                    *webpubsubhub.TimeoutsState       `json:"timeouts"`
}
