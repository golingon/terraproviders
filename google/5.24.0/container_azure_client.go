// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	containerazureclient "github.com/golingon/terraproviders/google/5.24.0/containerazureclient"
	"io"
)

// NewContainerAzureClient creates a new instance of [ContainerAzureClient].
func NewContainerAzureClient(name string, args ContainerAzureClientArgs) *ContainerAzureClient {
	return &ContainerAzureClient{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAzureClient)(nil)

// ContainerAzureClient represents the Terraform resource google_container_azure_client.
type ContainerAzureClient struct {
	Name      string
	Args      ContainerAzureClientArgs
	state     *containerAzureClientState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAzureClient].
func (cac *ContainerAzureClient) Type() string {
	return "google_container_azure_client"
}

// LocalName returns the local name for [ContainerAzureClient].
func (cac *ContainerAzureClient) LocalName() string {
	return cac.Name
}

// Configuration returns the configuration (args) for [ContainerAzureClient].
func (cac *ContainerAzureClient) Configuration() interface{} {
	return cac.Args
}

// DependOn is used for other resources to depend on [ContainerAzureClient].
func (cac *ContainerAzureClient) DependOn() terra.Reference {
	return terra.ReferenceResource(cac)
}

// Dependencies returns the list of resources [ContainerAzureClient] depends_on.
func (cac *ContainerAzureClient) Dependencies() terra.Dependencies {
	return cac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAzureClient].
func (cac *ContainerAzureClient) LifecycleManagement() *terra.Lifecycle {
	return cac.Lifecycle
}

// Attributes returns the attributes for [ContainerAzureClient].
func (cac *ContainerAzureClient) Attributes() containerAzureClientAttributes {
	return containerAzureClientAttributes{ref: terra.ReferenceResource(cac)}
}

// ImportState imports the given attribute values into [ContainerAzureClient]'s state.
func (cac *ContainerAzureClient) ImportState(av io.Reader) error {
	cac.state = &containerAzureClientState{}
	if err := json.NewDecoder(av).Decode(cac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cac.Type(), cac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAzureClient] has state.
func (cac *ContainerAzureClient) State() (*containerAzureClientState, bool) {
	return cac.state, cac.state != nil
}

// StateMust returns the state for [ContainerAzureClient]. Panics if the state is nil.
func (cac *ContainerAzureClient) StateMust() *containerAzureClientState {
	if cac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cac.Type(), cac.LocalName()))
	}
	return cac.state
}

// ContainerAzureClientArgs contains the configurations for google_container_azure_client.
type ContainerAzureClientArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TenantId: string, required
	TenantId terra.StringValue `hcl:"tenant_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *containerazureclient.Timeouts `hcl:"timeouts,block"`
}
type containerAzureClientAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of google_container_azure_client.
func (cac containerAzureClientAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("application_id"))
}

// Certificate returns a reference to field certificate of google_container_azure_client.
func (cac containerAzureClientAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("certificate"))
}

// CreateTime returns a reference to field create_time of google_container_azure_client.
func (cac containerAzureClientAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("create_time"))
}

// Id returns a reference to field id of google_container_azure_client.
func (cac containerAzureClientAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("id"))
}

// Location returns a reference to field location of google_container_azure_client.
func (cac containerAzureClientAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("location"))
}

// Name returns a reference to field name of google_container_azure_client.
func (cac containerAzureClientAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("name"))
}

// Project returns a reference to field project of google_container_azure_client.
func (cac containerAzureClientAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("project"))
}

// TenantId returns a reference to field tenant_id of google_container_azure_client.
func (cac containerAzureClientAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("tenant_id"))
}

// Uid returns a reference to field uid of google_container_azure_client.
func (cac containerAzureClientAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(cac.ref.Append("uid"))
}

func (cac containerAzureClientAttributes) Timeouts() containerazureclient.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containerazureclient.TimeoutsAttributes](cac.ref.Append("timeouts"))
}

type containerAzureClientState struct {
	ApplicationId string                              `json:"application_id"`
	Certificate   string                              `json:"certificate"`
	CreateTime    string                              `json:"create_time"`
	Id            string                              `json:"id"`
	Location      string                              `json:"location"`
	Name          string                              `json:"name"`
	Project       string                              `json:"project"`
	TenantId      string                              `json:"tenant_id"`
	Uid           string                              `json:"uid"`
	Timeouts      *containerazureclient.TimeoutsState `json:"timeouts"`
}
