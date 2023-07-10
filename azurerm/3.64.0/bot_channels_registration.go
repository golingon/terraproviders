// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botchannelsregistration "github.com/golingon/terraproviders/azurerm/3.64.0/botchannelsregistration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotChannelsRegistration creates a new instance of [BotChannelsRegistration].
func NewBotChannelsRegistration(name string, args BotChannelsRegistrationArgs) *BotChannelsRegistration {
	return &BotChannelsRegistration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotChannelsRegistration)(nil)

// BotChannelsRegistration represents the Terraform resource azurerm_bot_channels_registration.
type BotChannelsRegistration struct {
	Name      string
	Args      BotChannelsRegistrationArgs
	state     *botChannelsRegistrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotChannelsRegistration].
func (bcr *BotChannelsRegistration) Type() string {
	return "azurerm_bot_channels_registration"
}

// LocalName returns the local name for [BotChannelsRegistration].
func (bcr *BotChannelsRegistration) LocalName() string {
	return bcr.Name
}

// Configuration returns the configuration (args) for [BotChannelsRegistration].
func (bcr *BotChannelsRegistration) Configuration() interface{} {
	return bcr.Args
}

// DependOn is used for other resources to depend on [BotChannelsRegistration].
func (bcr *BotChannelsRegistration) DependOn() terra.Reference {
	return terra.ReferenceResource(bcr)
}

// Dependencies returns the list of resources [BotChannelsRegistration] depends_on.
func (bcr *BotChannelsRegistration) Dependencies() terra.Dependencies {
	return bcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotChannelsRegistration].
func (bcr *BotChannelsRegistration) LifecycleManagement() *terra.Lifecycle {
	return bcr.Lifecycle
}

// Attributes returns the attributes for [BotChannelsRegistration].
func (bcr *BotChannelsRegistration) Attributes() botChannelsRegistrationAttributes {
	return botChannelsRegistrationAttributes{ref: terra.ReferenceResource(bcr)}
}

// ImportState imports the given attribute values into [BotChannelsRegistration]'s state.
func (bcr *BotChannelsRegistration) ImportState(av io.Reader) error {
	bcr.state = &botChannelsRegistrationState{}
	if err := json.NewDecoder(av).Decode(bcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bcr.Type(), bcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotChannelsRegistration] has state.
func (bcr *BotChannelsRegistration) State() (*botChannelsRegistrationState, bool) {
	return bcr.state, bcr.state != nil
}

// StateMust returns the state for [BotChannelsRegistration]. Panics if the state is nil.
func (bcr *BotChannelsRegistration) StateMust() *botChannelsRegistrationState {
	if bcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bcr.Type(), bcr.LocalName()))
	}
	return bcr.state
}

// BotChannelsRegistrationArgs contains the configurations for azurerm_bot_channels_registration.
type BotChannelsRegistrationArgs struct {
	// CmkKeyVaultUrl: string, optional
	CmkKeyVaultUrl terra.StringValue `hcl:"cmk_key_vault_url,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DeveloperAppInsightsApiKey: string, optional
	DeveloperAppInsightsApiKey terra.StringValue `hcl:"developer_app_insights_api_key,attr"`
	// DeveloperAppInsightsApplicationId: string, optional
	DeveloperAppInsightsApplicationId terra.StringValue `hcl:"developer_app_insights_application_id,attr"`
	// DeveloperAppInsightsKey: string, optional
	DeveloperAppInsightsKey terra.StringValue `hcl:"developer_app_insights_key,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Endpoint: string, optional
	Endpoint terra.StringValue `hcl:"endpoint,attr"`
	// IconUrl: string, optional
	IconUrl terra.StringValue `hcl:"icon_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsolatedNetworkEnabled: bool, optional
	IsolatedNetworkEnabled terra.BoolValue `hcl:"isolated_network_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MicrosoftAppId: string, required
	MicrosoftAppId terra.StringValue `hcl:"microsoft_app_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// StreamingEndpointEnabled: bool, optional
	StreamingEndpointEnabled terra.BoolValue `hcl:"streaming_endpoint_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *botchannelsregistration.Timeouts `hcl:"timeouts,block"`
}
type botChannelsRegistrationAttributes struct {
	ref terra.Reference
}

// CmkKeyVaultUrl returns a reference to field cmk_key_vault_url of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) CmkKeyVaultUrl() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("cmk_key_vault_url"))
}

// Description returns a reference to field description of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("description"))
}

// DeveloperAppInsightsApiKey returns a reference to field developer_app_insights_api_key of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) DeveloperAppInsightsApiKey() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("developer_app_insights_api_key"))
}

// DeveloperAppInsightsApplicationId returns a reference to field developer_app_insights_application_id of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) DeveloperAppInsightsApplicationId() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("developer_app_insights_application_id"))
}

// DeveloperAppInsightsKey returns a reference to field developer_app_insights_key of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) DeveloperAppInsightsKey() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("developer_app_insights_key"))
}

// DisplayName returns a reference to field display_name of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("display_name"))
}

// Endpoint returns a reference to field endpoint of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("endpoint"))
}

// IconUrl returns a reference to field icon_url of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) IconUrl() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("icon_url"))
}

// Id returns a reference to field id of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("id"))
}

// IsolatedNetworkEnabled returns a reference to field isolated_network_enabled of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) IsolatedNetworkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bcr.ref.Append("isolated_network_enabled"))
}

// Location returns a reference to field location of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("location"))
}

// MicrosoftAppId returns a reference to field microsoft_app_id of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) MicrosoftAppId() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("microsoft_app_id"))
}

// Name returns a reference to field name of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bcr.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(bcr.ref.Append("sku"))
}

// StreamingEndpointEnabled returns a reference to field streaming_endpoint_enabled of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) StreamingEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bcr.ref.Append("streaming_endpoint_enabled"))
}

// Tags returns a reference to field tags of azurerm_bot_channels_registration.
func (bcr botChannelsRegistrationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bcr.ref.Append("tags"))
}

func (bcr botChannelsRegistrationAttributes) Timeouts() botchannelsregistration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botchannelsregistration.TimeoutsAttributes](bcr.ref.Append("timeouts"))
}

type botChannelsRegistrationState struct {
	CmkKeyVaultUrl                    string                                 `json:"cmk_key_vault_url"`
	Description                       string                                 `json:"description"`
	DeveloperAppInsightsApiKey        string                                 `json:"developer_app_insights_api_key"`
	DeveloperAppInsightsApplicationId string                                 `json:"developer_app_insights_application_id"`
	DeveloperAppInsightsKey           string                                 `json:"developer_app_insights_key"`
	DisplayName                       string                                 `json:"display_name"`
	Endpoint                          string                                 `json:"endpoint"`
	IconUrl                           string                                 `json:"icon_url"`
	Id                                string                                 `json:"id"`
	IsolatedNetworkEnabled            bool                                   `json:"isolated_network_enabled"`
	Location                          string                                 `json:"location"`
	MicrosoftAppId                    string                                 `json:"microsoft_app_id"`
	Name                              string                                 `json:"name"`
	PublicNetworkAccessEnabled        bool                                   `json:"public_network_access_enabled"`
	ResourceGroupName                 string                                 `json:"resource_group_name"`
	Sku                               string                                 `json:"sku"`
	StreamingEndpointEnabled          bool                                   `json:"streaming_endpoint_enabled"`
	Tags                              map[string]string                      `json:"tags"`
	Timeouts                          *botchannelsregistration.TimeoutsState `json:"timeouts"`
}
