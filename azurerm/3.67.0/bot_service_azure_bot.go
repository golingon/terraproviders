// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botserviceazurebot "github.com/golingon/terraproviders/azurerm/3.67.0/botserviceazurebot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotServiceAzureBot creates a new instance of [BotServiceAzureBot].
func NewBotServiceAzureBot(name string, args BotServiceAzureBotArgs) *BotServiceAzureBot {
	return &BotServiceAzureBot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotServiceAzureBot)(nil)

// BotServiceAzureBot represents the Terraform resource azurerm_bot_service_azure_bot.
type BotServiceAzureBot struct {
	Name      string
	Args      BotServiceAzureBotArgs
	state     *botServiceAzureBotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotServiceAzureBot].
func (bsab *BotServiceAzureBot) Type() string {
	return "azurerm_bot_service_azure_bot"
}

// LocalName returns the local name for [BotServiceAzureBot].
func (bsab *BotServiceAzureBot) LocalName() string {
	return bsab.Name
}

// Configuration returns the configuration (args) for [BotServiceAzureBot].
func (bsab *BotServiceAzureBot) Configuration() interface{} {
	return bsab.Args
}

// DependOn is used for other resources to depend on [BotServiceAzureBot].
func (bsab *BotServiceAzureBot) DependOn() terra.Reference {
	return terra.ReferenceResource(bsab)
}

// Dependencies returns the list of resources [BotServiceAzureBot] depends_on.
func (bsab *BotServiceAzureBot) Dependencies() terra.Dependencies {
	return bsab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotServiceAzureBot].
func (bsab *BotServiceAzureBot) LifecycleManagement() *terra.Lifecycle {
	return bsab.Lifecycle
}

// Attributes returns the attributes for [BotServiceAzureBot].
func (bsab *BotServiceAzureBot) Attributes() botServiceAzureBotAttributes {
	return botServiceAzureBotAttributes{ref: terra.ReferenceResource(bsab)}
}

// ImportState imports the given attribute values into [BotServiceAzureBot]'s state.
func (bsab *BotServiceAzureBot) ImportState(av io.Reader) error {
	bsab.state = &botServiceAzureBotState{}
	if err := json.NewDecoder(av).Decode(bsab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bsab.Type(), bsab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotServiceAzureBot] has state.
func (bsab *BotServiceAzureBot) State() (*botServiceAzureBotState, bool) {
	return bsab.state, bsab.state != nil
}

// StateMust returns the state for [BotServiceAzureBot]. Panics if the state is nil.
func (bsab *BotServiceAzureBot) StateMust() *botServiceAzureBotState {
	if bsab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bsab.Type(), bsab.LocalName()))
	}
	return bsab.state
}

// BotServiceAzureBotArgs contains the configurations for azurerm_bot_service_azure_bot.
type BotServiceAzureBotArgs struct {
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
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LuisAppIds: list of string, optional
	LuisAppIds terra.ListValue[terra.StringValue] `hcl:"luis_app_ids,attr"`
	// LuisKey: string, optional
	LuisKey terra.StringValue `hcl:"luis_key,attr"`
	// MicrosoftAppId: string, required
	MicrosoftAppId terra.StringValue `hcl:"microsoft_app_id,attr" validate:"required"`
	// MicrosoftAppMsiId: string, optional
	MicrosoftAppMsiId terra.StringValue `hcl:"microsoft_app_msi_id,attr"`
	// MicrosoftAppTenantId: string, optional
	MicrosoftAppTenantId terra.StringValue `hcl:"microsoft_app_tenant_id,attr"`
	// MicrosoftAppType: string, optional
	MicrosoftAppType terra.StringValue `hcl:"microsoft_app_type,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// StreamingEndpointEnabled: bool, optional
	StreamingEndpointEnabled terra.BoolValue `hcl:"streaming_endpoint_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *botserviceazurebot.Timeouts `hcl:"timeouts,block"`
}
type botServiceAzureBotAttributes struct {
	ref terra.Reference
}

// DeveloperAppInsightsApiKey returns a reference to field developer_app_insights_api_key of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) DeveloperAppInsightsApiKey() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("developer_app_insights_api_key"))
}

// DeveloperAppInsightsApplicationId returns a reference to field developer_app_insights_application_id of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) DeveloperAppInsightsApplicationId() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("developer_app_insights_application_id"))
}

// DeveloperAppInsightsKey returns a reference to field developer_app_insights_key of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) DeveloperAppInsightsKey() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("developer_app_insights_key"))
}

// DisplayName returns a reference to field display_name of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("display_name"))
}

// Endpoint returns a reference to field endpoint of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("location"))
}

// LuisAppIds returns a reference to field luis_app_ids of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) LuisAppIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bsab.ref.Append("luis_app_ids"))
}

// LuisKey returns a reference to field luis_key of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) LuisKey() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("luis_key"))
}

// MicrosoftAppId returns a reference to field microsoft_app_id of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) MicrosoftAppId() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("microsoft_app_id"))
}

// MicrosoftAppMsiId returns a reference to field microsoft_app_msi_id of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) MicrosoftAppMsiId() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("microsoft_app_msi_id"))
}

// MicrosoftAppTenantId returns a reference to field microsoft_app_tenant_id of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) MicrosoftAppTenantId() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("microsoft_app_tenant_id"))
}

// MicrosoftAppType returns a reference to field microsoft_app_type of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) MicrosoftAppType() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("microsoft_app_type"))
}

// Name returns a reference to field name of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(bsab.ref.Append("sku"))
}

// StreamingEndpointEnabled returns a reference to field streaming_endpoint_enabled of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) StreamingEndpointEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bsab.ref.Append("streaming_endpoint_enabled"))
}

// Tags returns a reference to field tags of azurerm_bot_service_azure_bot.
func (bsab botServiceAzureBotAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bsab.ref.Append("tags"))
}

func (bsab botServiceAzureBotAttributes) Timeouts() botserviceazurebot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botserviceazurebot.TimeoutsAttributes](bsab.ref.Append("timeouts"))
}

type botServiceAzureBotState struct {
	DeveloperAppInsightsApiKey        string                            `json:"developer_app_insights_api_key"`
	DeveloperAppInsightsApplicationId string                            `json:"developer_app_insights_application_id"`
	DeveloperAppInsightsKey           string                            `json:"developer_app_insights_key"`
	DisplayName                       string                            `json:"display_name"`
	Endpoint                          string                            `json:"endpoint"`
	Id                                string                            `json:"id"`
	Location                          string                            `json:"location"`
	LuisAppIds                        []string                          `json:"luis_app_ids"`
	LuisKey                           string                            `json:"luis_key"`
	MicrosoftAppId                    string                            `json:"microsoft_app_id"`
	MicrosoftAppMsiId                 string                            `json:"microsoft_app_msi_id"`
	MicrosoftAppTenantId              string                            `json:"microsoft_app_tenant_id"`
	MicrosoftAppType                  string                            `json:"microsoft_app_type"`
	Name                              string                            `json:"name"`
	ResourceGroupName                 string                            `json:"resource_group_name"`
	Sku                               string                            `json:"sku"`
	StreamingEndpointEnabled          bool                              `json:"streaming_endpoint_enabled"`
	Tags                              map[string]string                 `json:"tags"`
	Timeouts                          *botserviceazurebot.TimeoutsState `json:"timeouts"`
}
