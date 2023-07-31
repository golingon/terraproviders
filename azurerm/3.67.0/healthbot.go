// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	healthbot "github.com/golingon/terraproviders/azurerm/3.67.0/healthbot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthbot creates a new instance of [Healthbot].
func NewHealthbot(name string, args HealthbotArgs) *Healthbot {
	return &Healthbot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Healthbot)(nil)

// Healthbot represents the Terraform resource azurerm_healthbot.
type Healthbot struct {
	Name      string
	Args      HealthbotArgs
	state     *healthbotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Healthbot].
func (h *Healthbot) Type() string {
	return "azurerm_healthbot"
}

// LocalName returns the local name for [Healthbot].
func (h *Healthbot) LocalName() string {
	return h.Name
}

// Configuration returns the configuration (args) for [Healthbot].
func (h *Healthbot) Configuration() interface{} {
	return h.Args
}

// DependOn is used for other resources to depend on [Healthbot].
func (h *Healthbot) DependOn() terra.Reference {
	return terra.ReferenceResource(h)
}

// Dependencies returns the list of resources [Healthbot] depends_on.
func (h *Healthbot) Dependencies() terra.Dependencies {
	return h.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Healthbot].
func (h *Healthbot) LifecycleManagement() *terra.Lifecycle {
	return h.Lifecycle
}

// Attributes returns the attributes for [Healthbot].
func (h *Healthbot) Attributes() healthbotAttributes {
	return healthbotAttributes{ref: terra.ReferenceResource(h)}
}

// ImportState imports the given attribute values into [Healthbot]'s state.
func (h *Healthbot) ImportState(av io.Reader) error {
	h.state = &healthbotState{}
	if err := json.NewDecoder(av).Decode(h.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", h.Type(), h.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Healthbot] has state.
func (h *Healthbot) State() (*healthbotState, bool) {
	return h.state, h.state != nil
}

// StateMust returns the state for [Healthbot]. Panics if the state is nil.
func (h *Healthbot) StateMust() *healthbotState {
	if h.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", h.Type(), h.LocalName()))
	}
	return h.state
}

// HealthbotArgs contains the configurations for azurerm_healthbot.
type HealthbotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *healthbot.Timeouts `hcl:"timeouts,block"`
}
type healthbotAttributes struct {
	ref terra.Reference
}

// BotManagementPortalUrl returns a reference to field bot_management_portal_url of azurerm_healthbot.
func (h healthbotAttributes) BotManagementPortalUrl() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("bot_management_portal_url"))
}

// Id returns a reference to field id of azurerm_healthbot.
func (h healthbotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_healthbot.
func (h healthbotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_healthbot.
func (h healthbotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_healthbot.
func (h healthbotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_healthbot.
func (h healthbotAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(h.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_healthbot.
func (h healthbotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](h.ref.Append("tags"))
}

func (h healthbotAttributes) Timeouts() healthbot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[healthbot.TimeoutsAttributes](h.ref.Append("timeouts"))
}

type healthbotState struct {
	BotManagementPortalUrl string                   `json:"bot_management_portal_url"`
	Id                     string                   `json:"id"`
	Location               string                   `json:"location"`
	Name                   string                   `json:"name"`
	ResourceGroupName      string                   `json:"resource_group_name"`
	SkuName                string                   `json:"sku_name"`
	Tags                   map[string]string        `json:"tags"`
	Timeouts               *healthbot.TimeoutsState `json:"timeouts"`
}
