// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAzureSecretBackend creates a new instance of [AzureSecretBackend].
func NewAzureSecretBackend(name string, args AzureSecretBackendArgs) *AzureSecretBackend {
	return &AzureSecretBackend{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AzureSecretBackend)(nil)

// AzureSecretBackend represents the Terraform resource vault_azure_secret_backend.
type AzureSecretBackend struct {
	Name      string
	Args      AzureSecretBackendArgs
	state     *azureSecretBackendState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AzureSecretBackend].
func (asb *AzureSecretBackend) Type() string {
	return "vault_azure_secret_backend"
}

// LocalName returns the local name for [AzureSecretBackend].
func (asb *AzureSecretBackend) LocalName() string {
	return asb.Name
}

// Configuration returns the configuration (args) for [AzureSecretBackend].
func (asb *AzureSecretBackend) Configuration() interface{} {
	return asb.Args
}

// DependOn is used for other resources to depend on [AzureSecretBackend].
func (asb *AzureSecretBackend) DependOn() terra.Reference {
	return terra.ReferenceResource(asb)
}

// Dependencies returns the list of resources [AzureSecretBackend] depends_on.
func (asb *AzureSecretBackend) Dependencies() terra.Dependencies {
	return asb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AzureSecretBackend].
func (asb *AzureSecretBackend) LifecycleManagement() *terra.Lifecycle {
	return asb.Lifecycle
}

// Attributes returns the attributes for [AzureSecretBackend].
func (asb *AzureSecretBackend) Attributes() azureSecretBackendAttributes {
	return azureSecretBackendAttributes{ref: terra.ReferenceResource(asb)}
}

// ImportState imports the given attribute values into [AzureSecretBackend]'s state.
func (asb *AzureSecretBackend) ImportState(av io.Reader) error {
	asb.state = &azureSecretBackendState{}
	if err := json.NewDecoder(av).Decode(asb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asb.Type(), asb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AzureSecretBackend] has state.
func (asb *AzureSecretBackend) State() (*azureSecretBackendState, bool) {
	return asb.state, asb.state != nil
}

// StateMust returns the state for [AzureSecretBackend]. Panics if the state is nil.
func (asb *AzureSecretBackend) StateMust() *azureSecretBackendState {
	if asb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asb.Type(), asb.LocalName()))
	}
	return asb.state
}

// AzureSecretBackendArgs contains the configurations for vault_azure_secret_backend.
type AzureSecretBackendArgs struct {
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisableRemount: bool, optional
	DisableRemount terra.BoolValue `hcl:"disable_remount,attr"`
	// Environment: string, optional
	Environment terra.StringValue `hcl:"environment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// UseMicrosoftGraphApi: bool, optional
	UseMicrosoftGraphApi terra.BoolValue `hcl:"use_microsoft_graph_api,attr"`
}
type azureSecretBackendAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("client_id"))
}

// ClientSecret returns a reference to field client_secret of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("client_secret"))
}

// Description returns a reference to field description of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("description"))
}

// DisableRemount returns a reference to field disable_remount of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) DisableRemount() terra.BoolValue {
	return terra.ReferenceAsBool(asb.ref.Append("disable_remount"))
}

// Environment returns a reference to field environment of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("environment"))
}

// Id returns a reference to field id of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("id"))
}

// Namespace returns a reference to field namespace of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("namespace"))
}

// Path returns a reference to field path of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("path"))
}

// SubscriptionId returns a reference to field subscription_id of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("subscription_id"))
}

// TenantId returns a reference to field tenant_id of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(asb.ref.Append("tenant_id"))
}

// UseMicrosoftGraphApi returns a reference to field use_microsoft_graph_api of vault_azure_secret_backend.
func (asb azureSecretBackendAttributes) UseMicrosoftGraphApi() terra.BoolValue {
	return terra.ReferenceAsBool(asb.ref.Append("use_microsoft_graph_api"))
}

type azureSecretBackendState struct {
	ClientId             string `json:"client_id"`
	ClientSecret         string `json:"client_secret"`
	Description          string `json:"description"`
	DisableRemount       bool   `json:"disable_remount"`
	Environment          string `json:"environment"`
	Id                   string `json:"id"`
	Namespace            string `json:"namespace"`
	Path                 string `json:"path"`
	SubscriptionId       string `json:"subscription_id"`
	TenantId             string `json:"tenant_id"`
	UseMicrosoftGraphApi bool   `json:"use_microsoft_graph_api"`
}
