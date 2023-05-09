// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudbuildbitbucketserverconfig "github.com/golingon/terraproviders/google/4.64.0/cloudbuildbitbucketserverconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildBitbucketServerConfig creates a new instance of [CloudbuildBitbucketServerConfig].
func NewCloudbuildBitbucketServerConfig(name string, args CloudbuildBitbucketServerConfigArgs) *CloudbuildBitbucketServerConfig {
	return &CloudbuildBitbucketServerConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudbuildBitbucketServerConfig)(nil)

// CloudbuildBitbucketServerConfig represents the Terraform resource google_cloudbuild_bitbucket_server_config.
type CloudbuildBitbucketServerConfig struct {
	Name      string
	Args      CloudbuildBitbucketServerConfigArgs
	state     *cloudbuildBitbucketServerConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudbuildBitbucketServerConfig].
func (cbsc *CloudbuildBitbucketServerConfig) Type() string {
	return "google_cloudbuild_bitbucket_server_config"
}

// LocalName returns the local name for [CloudbuildBitbucketServerConfig].
func (cbsc *CloudbuildBitbucketServerConfig) LocalName() string {
	return cbsc.Name
}

// Configuration returns the configuration (args) for [CloudbuildBitbucketServerConfig].
func (cbsc *CloudbuildBitbucketServerConfig) Configuration() interface{} {
	return cbsc.Args
}

// DependOn is used for other resources to depend on [CloudbuildBitbucketServerConfig].
func (cbsc *CloudbuildBitbucketServerConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(cbsc)
}

// Dependencies returns the list of resources [CloudbuildBitbucketServerConfig] depends_on.
func (cbsc *CloudbuildBitbucketServerConfig) Dependencies() terra.Dependencies {
	return cbsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudbuildBitbucketServerConfig].
func (cbsc *CloudbuildBitbucketServerConfig) LifecycleManagement() *terra.Lifecycle {
	return cbsc.Lifecycle
}

// Attributes returns the attributes for [CloudbuildBitbucketServerConfig].
func (cbsc *CloudbuildBitbucketServerConfig) Attributes() cloudbuildBitbucketServerConfigAttributes {
	return cloudbuildBitbucketServerConfigAttributes{ref: terra.ReferenceResource(cbsc)}
}

// ImportState imports the given attribute values into [CloudbuildBitbucketServerConfig]'s state.
func (cbsc *CloudbuildBitbucketServerConfig) ImportState(av io.Reader) error {
	cbsc.state = &cloudbuildBitbucketServerConfigState{}
	if err := json.NewDecoder(av).Decode(cbsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbsc.Type(), cbsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudbuildBitbucketServerConfig] has state.
func (cbsc *CloudbuildBitbucketServerConfig) State() (*cloudbuildBitbucketServerConfigState, bool) {
	return cbsc.state, cbsc.state != nil
}

// StateMust returns the state for [CloudbuildBitbucketServerConfig]. Panics if the state is nil.
func (cbsc *CloudbuildBitbucketServerConfig) StateMust() *cloudbuildBitbucketServerConfigState {
	if cbsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbsc.Type(), cbsc.LocalName()))
	}
	return cbsc.state
}

// CloudbuildBitbucketServerConfigArgs contains the configurations for google_cloudbuild_bitbucket_server_config.
type CloudbuildBitbucketServerConfigArgs struct {
	// ApiKey: string, required
	ApiKey terra.StringValue `hcl:"api_key,attr" validate:"required"`
	// ConfigId: string, required
	ConfigId terra.StringValue `hcl:"config_id,attr" validate:"required"`
	// HostUri: string, required
	HostUri terra.StringValue `hcl:"host_uri,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PeeredNetwork: string, optional
	PeeredNetwork terra.StringValue `hcl:"peered_network,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SslCa: string, optional
	SslCa terra.StringValue `hcl:"ssl_ca,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// ConnectedRepositories: min=0
	ConnectedRepositories []cloudbuildbitbucketserverconfig.ConnectedRepositories `hcl:"connected_repositories,block" validate:"min=0"`
	// Secrets: required
	Secrets *cloudbuildbitbucketserverconfig.Secrets `hcl:"secrets,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudbuildbitbucketserverconfig.Timeouts `hcl:"timeouts,block"`
}
type cloudbuildBitbucketServerConfigAttributes struct {
	ref terra.Reference
}

// ApiKey returns a reference to field api_key of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("api_key"))
}

// ConfigId returns a reference to field config_id of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) ConfigId() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("config_id"))
}

// HostUri returns a reference to field host_uri of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) HostUri() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("host_uri"))
}

// Id returns a reference to field id of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("name"))
}

// PeeredNetwork returns a reference to field peered_network of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) PeeredNetwork() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("peered_network"))
}

// Project returns a reference to field project of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("project"))
}

// SslCa returns a reference to field ssl_ca of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) SslCa() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("ssl_ca"))
}

// Username returns a reference to field username of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("username"))
}

// WebhookKey returns a reference to field webhook_key of google_cloudbuild_bitbucket_server_config.
func (cbsc cloudbuildBitbucketServerConfigAttributes) WebhookKey() terra.StringValue {
	return terra.ReferenceAsString(cbsc.ref.Append("webhook_key"))
}

func (cbsc cloudbuildBitbucketServerConfigAttributes) ConnectedRepositories() terra.SetValue[cloudbuildbitbucketserverconfig.ConnectedRepositoriesAttributes] {
	return terra.ReferenceAsSet[cloudbuildbitbucketserverconfig.ConnectedRepositoriesAttributes](cbsc.ref.Append("connected_repositories"))
}

func (cbsc cloudbuildBitbucketServerConfigAttributes) Secrets() terra.ListValue[cloudbuildbitbucketserverconfig.SecretsAttributes] {
	return terra.ReferenceAsList[cloudbuildbitbucketserverconfig.SecretsAttributes](cbsc.ref.Append("secrets"))
}

func (cbsc cloudbuildBitbucketServerConfigAttributes) Timeouts() cloudbuildbitbucketserverconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudbuildbitbucketserverconfig.TimeoutsAttributes](cbsc.ref.Append("timeouts"))
}

type cloudbuildBitbucketServerConfigState struct {
	ApiKey                string                                                       `json:"api_key"`
	ConfigId              string                                                       `json:"config_id"`
	HostUri               string                                                       `json:"host_uri"`
	Id                    string                                                       `json:"id"`
	Location              string                                                       `json:"location"`
	Name                  string                                                       `json:"name"`
	PeeredNetwork         string                                                       `json:"peered_network"`
	Project               string                                                       `json:"project"`
	SslCa                 string                                                       `json:"ssl_ca"`
	Username              string                                                       `json:"username"`
	WebhookKey            string                                                       `json:"webhook_key"`
	ConnectedRepositories []cloudbuildbitbucketserverconfig.ConnectedRepositoriesState `json:"connected_repositories"`
	Secrets               []cloudbuildbitbucketserverconfig.SecretsState               `json:"secrets"`
	Timeouts              *cloudbuildbitbucketserverconfig.TimeoutsState               `json:"timeouts"`
}
