// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSecretsSyncGcpDestination creates a new instance of [SecretsSyncGcpDestination].
func NewSecretsSyncGcpDestination(name string, args SecretsSyncGcpDestinationArgs) *SecretsSyncGcpDestination {
	return &SecretsSyncGcpDestination{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecretsSyncGcpDestination)(nil)

// SecretsSyncGcpDestination represents the Terraform resource vault_secrets_sync_gcp_destination.
type SecretsSyncGcpDestination struct {
	Name      string
	Args      SecretsSyncGcpDestinationArgs
	state     *secretsSyncGcpDestinationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecretsSyncGcpDestination].
func (ssgd *SecretsSyncGcpDestination) Type() string {
	return "vault_secrets_sync_gcp_destination"
}

// LocalName returns the local name for [SecretsSyncGcpDestination].
func (ssgd *SecretsSyncGcpDestination) LocalName() string {
	return ssgd.Name
}

// Configuration returns the configuration (args) for [SecretsSyncGcpDestination].
func (ssgd *SecretsSyncGcpDestination) Configuration() interface{} {
	return ssgd.Args
}

// DependOn is used for other resources to depend on [SecretsSyncGcpDestination].
func (ssgd *SecretsSyncGcpDestination) DependOn() terra.Reference {
	return terra.ReferenceResource(ssgd)
}

// Dependencies returns the list of resources [SecretsSyncGcpDestination] depends_on.
func (ssgd *SecretsSyncGcpDestination) Dependencies() terra.Dependencies {
	return ssgd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecretsSyncGcpDestination].
func (ssgd *SecretsSyncGcpDestination) LifecycleManagement() *terra.Lifecycle {
	return ssgd.Lifecycle
}

// Attributes returns the attributes for [SecretsSyncGcpDestination].
func (ssgd *SecretsSyncGcpDestination) Attributes() secretsSyncGcpDestinationAttributes {
	return secretsSyncGcpDestinationAttributes{ref: terra.ReferenceResource(ssgd)}
}

// ImportState imports the given attribute values into [SecretsSyncGcpDestination]'s state.
func (ssgd *SecretsSyncGcpDestination) ImportState(av io.Reader) error {
	ssgd.state = &secretsSyncGcpDestinationState{}
	if err := json.NewDecoder(av).Decode(ssgd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssgd.Type(), ssgd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecretsSyncGcpDestination] has state.
func (ssgd *SecretsSyncGcpDestination) State() (*secretsSyncGcpDestinationState, bool) {
	return ssgd.state, ssgd.state != nil
}

// StateMust returns the state for [SecretsSyncGcpDestination]. Panics if the state is nil.
func (ssgd *SecretsSyncGcpDestination) StateMust() *secretsSyncGcpDestinationState {
	if ssgd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssgd.Type(), ssgd.LocalName()))
	}
	return ssgd.state
}

// SecretsSyncGcpDestinationArgs contains the configurations for vault_secrets_sync_gcp_destination.
type SecretsSyncGcpDestinationArgs struct {
	// Credentials: string, optional
	Credentials terra.StringValue `hcl:"credentials,attr"`
	// CustomTags: map of string, optional
	CustomTags terra.MapValue[terra.StringValue] `hcl:"custom_tags,attr"`
	// Granularity: string, optional
	Granularity terra.StringValue `hcl:"granularity,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// ProjectId: string, optional
	ProjectId terra.StringValue `hcl:"project_id,attr"`
	// SecretNameTemplate: string, optional
	SecretNameTemplate terra.StringValue `hcl:"secret_name_template,attr"`
}
type secretsSyncGcpDestinationAttributes struct {
	ref terra.Reference
}

// Credentials returns a reference to field credentials of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) Credentials() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("credentials"))
}

// CustomTags returns a reference to field custom_tags of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) CustomTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssgd.ref.Append("custom_tags"))
}

// Granularity returns a reference to field granularity of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) Granularity() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("granularity"))
}

// Id returns a reference to field id of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("id"))
}

// Name returns a reference to field name of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("namespace"))
}

// ProjectId returns a reference to field project_id of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("project_id"))
}

// SecretNameTemplate returns a reference to field secret_name_template of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) SecretNameTemplate() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("secret_name_template"))
}

// Type returns a reference to field type of vault_secrets_sync_gcp_destination.
func (ssgd secretsSyncGcpDestinationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ssgd.ref.Append("type"))
}

type secretsSyncGcpDestinationState struct {
	Credentials        string            `json:"credentials"`
	CustomTags         map[string]string `json:"custom_tags"`
	Granularity        string            `json:"granularity"`
	Id                 string            `json:"id"`
	Name               string            `json:"name"`
	Namespace          string            `json:"namespace"`
	ProjectId          string            `json:"project_id"`
	SecretNameTemplate string            `json:"secret_name_template"`
	Type               string            `json:"type"`
}
