// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	botwebapp "github.com/golingon/terraproviders/azurerm/3.65.0/botwebapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBotWebApp creates a new instance of [BotWebApp].
func NewBotWebApp(name string, args BotWebAppArgs) *BotWebApp {
	return &BotWebApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BotWebApp)(nil)

// BotWebApp represents the Terraform resource azurerm_bot_web_app.
type BotWebApp struct {
	Name      string
	Args      BotWebAppArgs
	state     *botWebAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BotWebApp].
func (bwa *BotWebApp) Type() string {
	return "azurerm_bot_web_app"
}

// LocalName returns the local name for [BotWebApp].
func (bwa *BotWebApp) LocalName() string {
	return bwa.Name
}

// Configuration returns the configuration (args) for [BotWebApp].
func (bwa *BotWebApp) Configuration() interface{} {
	return bwa.Args
}

// DependOn is used for other resources to depend on [BotWebApp].
func (bwa *BotWebApp) DependOn() terra.Reference {
	return terra.ReferenceResource(bwa)
}

// Dependencies returns the list of resources [BotWebApp] depends_on.
func (bwa *BotWebApp) Dependencies() terra.Dependencies {
	return bwa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BotWebApp].
func (bwa *BotWebApp) LifecycleManagement() *terra.Lifecycle {
	return bwa.Lifecycle
}

// Attributes returns the attributes for [BotWebApp].
func (bwa *BotWebApp) Attributes() botWebAppAttributes {
	return botWebAppAttributes{ref: terra.ReferenceResource(bwa)}
}

// ImportState imports the given attribute values into [BotWebApp]'s state.
func (bwa *BotWebApp) ImportState(av io.Reader) error {
	bwa.state = &botWebAppState{}
	if err := json.NewDecoder(av).Decode(bwa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bwa.Type(), bwa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BotWebApp] has state.
func (bwa *BotWebApp) State() (*botWebAppState, bool) {
	return bwa.state, bwa.state != nil
}

// StateMust returns the state for [BotWebApp]. Panics if the state is nil.
func (bwa *BotWebApp) StateMust() *botWebAppState {
	if bwa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bwa.Type(), bwa.LocalName()))
	}
	return bwa.state
}

// BotWebAppArgs contains the configurations for azurerm_bot_web_app.
type BotWebAppArgs struct {
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
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *botwebapp.Timeouts `hcl:"timeouts,block"`
}
type botWebAppAttributes struct {
	ref terra.Reference
}

// DeveloperAppInsightsApiKey returns a reference to field developer_app_insights_api_key of azurerm_bot_web_app.
func (bwa botWebAppAttributes) DeveloperAppInsightsApiKey() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("developer_app_insights_api_key"))
}

// DeveloperAppInsightsApplicationId returns a reference to field developer_app_insights_application_id of azurerm_bot_web_app.
func (bwa botWebAppAttributes) DeveloperAppInsightsApplicationId() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("developer_app_insights_application_id"))
}

// DeveloperAppInsightsKey returns a reference to field developer_app_insights_key of azurerm_bot_web_app.
func (bwa botWebAppAttributes) DeveloperAppInsightsKey() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("developer_app_insights_key"))
}

// DisplayName returns a reference to field display_name of azurerm_bot_web_app.
func (bwa botWebAppAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("display_name"))
}

// Endpoint returns a reference to field endpoint of azurerm_bot_web_app.
func (bwa botWebAppAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("endpoint"))
}

// Id returns a reference to field id of azurerm_bot_web_app.
func (bwa botWebAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_bot_web_app.
func (bwa botWebAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("location"))
}

// LuisAppIds returns a reference to field luis_app_ids of azurerm_bot_web_app.
func (bwa botWebAppAttributes) LuisAppIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bwa.ref.Append("luis_app_ids"))
}

// LuisKey returns a reference to field luis_key of azurerm_bot_web_app.
func (bwa botWebAppAttributes) LuisKey() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("luis_key"))
}

// MicrosoftAppId returns a reference to field microsoft_app_id of azurerm_bot_web_app.
func (bwa botWebAppAttributes) MicrosoftAppId() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("microsoft_app_id"))
}

// Name returns a reference to field name of azurerm_bot_web_app.
func (bwa botWebAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_bot_web_app.
func (bwa botWebAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_bot_web_app.
func (bwa botWebAppAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(bwa.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_bot_web_app.
func (bwa botWebAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bwa.ref.Append("tags"))
}

func (bwa botWebAppAttributes) Timeouts() botwebapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[botwebapp.TimeoutsAttributes](bwa.ref.Append("timeouts"))
}

type botWebAppState struct {
	DeveloperAppInsightsApiKey        string                   `json:"developer_app_insights_api_key"`
	DeveloperAppInsightsApplicationId string                   `json:"developer_app_insights_application_id"`
	DeveloperAppInsightsKey           string                   `json:"developer_app_insights_key"`
	DisplayName                       string                   `json:"display_name"`
	Endpoint                          string                   `json:"endpoint"`
	Id                                string                   `json:"id"`
	Location                          string                   `json:"location"`
	LuisAppIds                        []string                 `json:"luis_app_ids"`
	LuisKey                           string                   `json:"luis_key"`
	MicrosoftAppId                    string                   `json:"microsoft_app_id"`
	Name                              string                   `json:"name"`
	ResourceGroupName                 string                   `json:"resource_group_name"`
	Sku                               string                   `json:"sku"`
	Tags                              map[string]string        `json:"tags"`
	Timeouts                          *botwebapp.TimeoutsState `json:"timeouts"`
}