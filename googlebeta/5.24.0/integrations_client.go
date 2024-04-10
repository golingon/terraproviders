// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	integrationsclient "github.com/golingon/terraproviders/googlebeta/5.24.0/integrationsclient"
	"io"
)

// NewIntegrationsClient creates a new instance of [IntegrationsClient].
func NewIntegrationsClient(name string, args IntegrationsClientArgs) *IntegrationsClient {
	return &IntegrationsClient{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IntegrationsClient)(nil)

// IntegrationsClient represents the Terraform resource google_integrations_client.
type IntegrationsClient struct {
	Name      string
	Args      IntegrationsClientArgs
	state     *integrationsClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IntegrationsClient].
func (ic *IntegrationsClient) Type() string {
	return "google_integrations_client"
}

// LocalName returns the local name for [IntegrationsClient].
func (ic *IntegrationsClient) LocalName() string {
	return ic.Name
}

// Configuration returns the configuration (args) for [IntegrationsClient].
func (ic *IntegrationsClient) Configuration() interface{} {
	return ic.Args
}

// DependOn is used for other resources to depend on [IntegrationsClient].
func (ic *IntegrationsClient) DependOn() terra.Reference {
	return terra.ReferenceResource(ic)
}

// Dependencies returns the list of resources [IntegrationsClient] depends_on.
func (ic *IntegrationsClient) Dependencies() terra.Dependencies {
	return ic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IntegrationsClient].
func (ic *IntegrationsClient) LifecycleManagement() *terra.Lifecycle {
	return ic.Lifecycle
}

// Attributes returns the attributes for [IntegrationsClient].
func (ic *IntegrationsClient) Attributes() integrationsClientAttributes {
	return integrationsClientAttributes{ref: terra.ReferenceResource(ic)}
}

// ImportState imports the given attribute values into [IntegrationsClient]'s state.
func (ic *IntegrationsClient) ImportState(av io.Reader) error {
	ic.state = &integrationsClientState{}
	if err := json.NewDecoder(av).Decode(ic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ic.Type(), ic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IntegrationsClient] has state.
func (ic *IntegrationsClient) State() (*integrationsClientState, bool) {
	return ic.state, ic.state != nil
}

// StateMust returns the state for [IntegrationsClient]. Panics if the state is nil.
func (ic *IntegrationsClient) StateMust() *integrationsClientState {
	if ic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ic.Type(), ic.LocalName()))
	}
	return ic.state
}

// IntegrationsClientArgs contains the configurations for google_integrations_client.
type IntegrationsClientArgs struct {
	// CreateSampleWorkflows: bool, optional
	CreateSampleWorkflows terra.BoolValue `hcl:"create_sample_workflows,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// ProvisionGmek: bool, optional
	ProvisionGmek terra.BoolValue `hcl:"provision_gmek,attr"`
	// RunAsServiceAccount: string, optional
	RunAsServiceAccount terra.StringValue `hcl:"run_as_service_account,attr"`
	// CloudKmsConfig: optional
	CloudKmsConfig *integrationsclient.CloudKmsConfig `hcl:"cloud_kms_config,block"`
	// Timeouts: optional
	Timeouts *integrationsclient.Timeouts `hcl:"timeouts,block"`
}
type integrationsClientAttributes struct {
	ref terra.Reference
}

// CreateSampleWorkflows returns a reference to field create_sample_workflows of google_integrations_client.
func (ic integrationsClientAttributes) CreateSampleWorkflows() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("create_sample_workflows"))
}

// Id returns a reference to field id of google_integrations_client.
func (ic integrationsClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

// Location returns a reference to field location of google_integrations_client.
func (ic integrationsClientAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("location"))
}

// Project returns a reference to field project of google_integrations_client.
func (ic integrationsClientAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("project"))
}

// ProvisionGmek returns a reference to field provision_gmek of google_integrations_client.
func (ic integrationsClientAttributes) ProvisionGmek() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("provision_gmek"))
}

// RunAsServiceAccount returns a reference to field run_as_service_account of google_integrations_client.
func (ic integrationsClientAttributes) RunAsServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("run_as_service_account"))
}

func (ic integrationsClientAttributes) CloudKmsConfig() terra.ListValue[integrationsclient.CloudKmsConfigAttributes] {
	return terra.ReferenceAsList[integrationsclient.CloudKmsConfigAttributes](ic.ref.Append("cloud_kms_config"))
}

func (ic integrationsClientAttributes) Timeouts() integrationsclient.TimeoutsAttributes {
	return terra.ReferenceAsSingle[integrationsclient.TimeoutsAttributes](ic.ref.Append("timeouts"))
}

type integrationsClientState struct {
	CreateSampleWorkflows bool                                     `json:"create_sample_workflows"`
	Id                    string                                   `json:"id"`
	Location              string                                   `json:"location"`
	Project               string                                   `json:"project"`
	ProvisionGmek         bool                                     `json:"provision_gmek"`
	RunAsServiceAccount   string                                   `json:"run_as_service_account"`
	CloudKmsConfig        []integrationsclient.CloudKmsConfigState `json:"cloud_kms_config"`
	Timeouts              *integrationsclient.TimeoutsState        `json:"timeouts"`
}
