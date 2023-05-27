// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cognitiveaccount "github.com/golingon/terraproviders/azurerm/3.58.0/cognitiveaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitiveAccount creates a new instance of [CognitiveAccount].
func NewCognitiveAccount(name string, args CognitiveAccountArgs) *CognitiveAccount {
	return &CognitiveAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitiveAccount)(nil)

// CognitiveAccount represents the Terraform resource azurerm_cognitive_account.
type CognitiveAccount struct {
	Name      string
	Args      CognitiveAccountArgs
	state     *cognitiveAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitiveAccount].
func (ca *CognitiveAccount) Type() string {
	return "azurerm_cognitive_account"
}

// LocalName returns the local name for [CognitiveAccount].
func (ca *CognitiveAccount) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [CognitiveAccount].
func (ca *CognitiveAccount) Configuration() interface{} {
	return ca.Args
}

// DependOn is used for other resources to depend on [CognitiveAccount].
func (ca *CognitiveAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(ca)
}

// Dependencies returns the list of resources [CognitiveAccount] depends_on.
func (ca *CognitiveAccount) Dependencies() terra.Dependencies {
	return ca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitiveAccount].
func (ca *CognitiveAccount) LifecycleManagement() *terra.Lifecycle {
	return ca.Lifecycle
}

// Attributes returns the attributes for [CognitiveAccount].
func (ca *CognitiveAccount) Attributes() cognitiveAccountAttributes {
	return cognitiveAccountAttributes{ref: terra.ReferenceResource(ca)}
}

// ImportState imports the given attribute values into [CognitiveAccount]'s state.
func (ca *CognitiveAccount) ImportState(av io.Reader) error {
	ca.state = &cognitiveAccountState{}
	if err := json.NewDecoder(av).Decode(ca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ca.Type(), ca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitiveAccount] has state.
func (ca *CognitiveAccount) State() (*cognitiveAccountState, bool) {
	return ca.state, ca.state != nil
}

// StateMust returns the state for [CognitiveAccount]. Panics if the state is nil.
func (ca *CognitiveAccount) StateMust() *cognitiveAccountState {
	if ca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ca.Type(), ca.LocalName()))
	}
	return ca.state
}

// CognitiveAccountArgs contains the configurations for azurerm_cognitive_account.
type CognitiveAccountArgs struct {
	// CustomQuestionAnsweringSearchServiceId: string, optional
	CustomQuestionAnsweringSearchServiceId terra.StringValue `hcl:"custom_question_answering_search_service_id,attr"`
	// CustomQuestionAnsweringSearchServiceKey: string, optional
	CustomQuestionAnsweringSearchServiceKey terra.StringValue `hcl:"custom_question_answering_search_service_key,attr"`
	// CustomSubdomainName: string, optional
	CustomSubdomainName terra.StringValue `hcl:"custom_subdomain_name,attr"`
	// DynamicThrottlingEnabled: bool, optional
	DynamicThrottlingEnabled terra.BoolValue `hcl:"dynamic_throttling_enabled,attr"`
	// Fqdns: list of string, optional
	Fqdns terra.ListValue[terra.StringValue] `hcl:"fqdns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// LocalAuthEnabled: bool, optional
	LocalAuthEnabled terra.BoolValue `hcl:"local_auth_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MetricsAdvisorAadClientId: string, optional
	MetricsAdvisorAadClientId terra.StringValue `hcl:"metrics_advisor_aad_client_id,attr"`
	// MetricsAdvisorAadTenantId: string, optional
	MetricsAdvisorAadTenantId terra.StringValue `hcl:"metrics_advisor_aad_tenant_id,attr"`
	// MetricsAdvisorSuperUserName: string, optional
	MetricsAdvisorSuperUserName terra.StringValue `hcl:"metrics_advisor_super_user_name,attr"`
	// MetricsAdvisorWebsiteName: string, optional
	MetricsAdvisorWebsiteName terra.StringValue `hcl:"metrics_advisor_website_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OutboundNetworkAccessRestricted: bool, optional
	OutboundNetworkAccessRestricted terra.BoolValue `hcl:"outbound_network_access_restricted,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// QnaRuntimeEndpoint: string, optional
	QnaRuntimeEndpoint terra.StringValue `hcl:"qna_runtime_endpoint,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// CustomerManagedKey: optional
	CustomerManagedKey *cognitiveaccount.CustomerManagedKey `hcl:"customer_managed_key,block"`
	// Identity: optional
	Identity *cognitiveaccount.Identity `hcl:"identity,block"`
	// NetworkAcls: optional
	NetworkAcls *cognitiveaccount.NetworkAcls `hcl:"network_acls,block"`
	// Storage: min=0
	Storage []cognitiveaccount.Storage `hcl:"storage,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *cognitiveaccount.Timeouts `hcl:"timeouts,block"`
}
type cognitiveAccountAttributes struct {
	ref terra.Reference
}

// CustomQuestionAnsweringSearchServiceId returns a reference to field custom_question_answering_search_service_id of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) CustomQuestionAnsweringSearchServiceId() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("custom_question_answering_search_service_id"))
}

// CustomQuestionAnsweringSearchServiceKey returns a reference to field custom_question_answering_search_service_key of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) CustomQuestionAnsweringSearchServiceKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("custom_question_answering_search_service_key"))
}

// CustomSubdomainName returns a reference to field custom_subdomain_name of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) CustomSubdomainName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("custom_subdomain_name"))
}

// DynamicThrottlingEnabled returns a reference to field dynamic_throttling_enabled of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) DynamicThrottlingEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("dynamic_throttling_enabled"))
}

// Endpoint returns a reference to field endpoint of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("endpoint"))
}

// Fqdns returns a reference to field fqdns of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("fqdns"))
}

// Id returns a reference to field id of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// Kind returns a reference to field kind of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("kind"))
}

// LocalAuthEnabled returns a reference to field local_auth_enabled of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) LocalAuthEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("local_auth_enabled"))
}

// Location returns a reference to field location of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("location"))
}

// MetricsAdvisorAadClientId returns a reference to field metrics_advisor_aad_client_id of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) MetricsAdvisorAadClientId() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("metrics_advisor_aad_client_id"))
}

// MetricsAdvisorAadTenantId returns a reference to field metrics_advisor_aad_tenant_id of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) MetricsAdvisorAadTenantId() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("metrics_advisor_aad_tenant_id"))
}

// MetricsAdvisorSuperUserName returns a reference to field metrics_advisor_super_user_name of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) MetricsAdvisorSuperUserName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("metrics_advisor_super_user_name"))
}

// MetricsAdvisorWebsiteName returns a reference to field metrics_advisor_website_name of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) MetricsAdvisorWebsiteName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("metrics_advisor_website_name"))
}

// Name returns a reference to field name of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// OutboundNetworkAccessRestricted returns a reference to field outbound_network_access_restricted of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) OutboundNetworkAccessRestricted() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("outbound_network_access_restricted"))
}

// PrimaryAccessKey returns a reference to field primary_access_key of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) PrimaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("primary_access_key"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ca.ref.Append("public_network_access_enabled"))
}

// QnaRuntimeEndpoint returns a reference to field qna_runtime_endpoint of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) QnaRuntimeEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("qna_runtime_endpoint"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("resource_group_name"))
}

// SecondaryAccessKey returns a reference to field secondary_access_key of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) SecondaryAccessKey() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("secondary_access_key"))
}

// SkuName returns a reference to field sku_name of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_cognitive_account.
func (ca cognitiveAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("tags"))
}

func (ca cognitiveAccountAttributes) CustomerManagedKey() terra.ListValue[cognitiveaccount.CustomerManagedKeyAttributes] {
	return terra.ReferenceAsList[cognitiveaccount.CustomerManagedKeyAttributes](ca.ref.Append("customer_managed_key"))
}

func (ca cognitiveAccountAttributes) Identity() terra.ListValue[cognitiveaccount.IdentityAttributes] {
	return terra.ReferenceAsList[cognitiveaccount.IdentityAttributes](ca.ref.Append("identity"))
}

func (ca cognitiveAccountAttributes) NetworkAcls() terra.ListValue[cognitiveaccount.NetworkAclsAttributes] {
	return terra.ReferenceAsList[cognitiveaccount.NetworkAclsAttributes](ca.ref.Append("network_acls"))
}

func (ca cognitiveAccountAttributes) Storage() terra.ListValue[cognitiveaccount.StorageAttributes] {
	return terra.ReferenceAsList[cognitiveaccount.StorageAttributes](ca.ref.Append("storage"))
}

func (ca cognitiveAccountAttributes) Timeouts() cognitiveaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cognitiveaccount.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

type cognitiveAccountState struct {
	CustomQuestionAnsweringSearchServiceId  string                                     `json:"custom_question_answering_search_service_id"`
	CustomQuestionAnsweringSearchServiceKey string                                     `json:"custom_question_answering_search_service_key"`
	CustomSubdomainName                     string                                     `json:"custom_subdomain_name"`
	DynamicThrottlingEnabled                bool                                       `json:"dynamic_throttling_enabled"`
	Endpoint                                string                                     `json:"endpoint"`
	Fqdns                                   []string                                   `json:"fqdns"`
	Id                                      string                                     `json:"id"`
	Kind                                    string                                     `json:"kind"`
	LocalAuthEnabled                        bool                                       `json:"local_auth_enabled"`
	Location                                string                                     `json:"location"`
	MetricsAdvisorAadClientId               string                                     `json:"metrics_advisor_aad_client_id"`
	MetricsAdvisorAadTenantId               string                                     `json:"metrics_advisor_aad_tenant_id"`
	MetricsAdvisorSuperUserName             string                                     `json:"metrics_advisor_super_user_name"`
	MetricsAdvisorWebsiteName               string                                     `json:"metrics_advisor_website_name"`
	Name                                    string                                     `json:"name"`
	OutboundNetworkAccessRestricted         bool                                       `json:"outbound_network_access_restricted"`
	PrimaryAccessKey                        string                                     `json:"primary_access_key"`
	PublicNetworkAccessEnabled              bool                                       `json:"public_network_access_enabled"`
	QnaRuntimeEndpoint                      string                                     `json:"qna_runtime_endpoint"`
	ResourceGroupName                       string                                     `json:"resource_group_name"`
	SecondaryAccessKey                      string                                     `json:"secondary_access_key"`
	SkuName                                 string                                     `json:"sku_name"`
	Tags                                    map[string]string                          `json:"tags"`
	CustomerManagedKey                      []cognitiveaccount.CustomerManagedKeyState `json:"customer_managed_key"`
	Identity                                []cognitiveaccount.IdentityState           `json:"identity"`
	NetworkAcls                             []cognitiveaccount.NetworkAclsState        `json:"network_acls"`
	Storage                                 []cognitiveaccount.StorageState            `json:"storage"`
	Timeouts                                *cognitiveaccount.TimeoutsState            `json:"timeouts"`
}
