// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	accesscontextmanagergcpuseraccessbinding "github.com/golingon/terraproviders/google/4.73.1/accesscontextmanagergcpuseraccessbinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerGcpUserAccessBinding creates a new instance of [AccessContextManagerGcpUserAccessBinding].
func NewAccessContextManagerGcpUserAccessBinding(name string, args AccessContextManagerGcpUserAccessBindingArgs) *AccessContextManagerGcpUserAccessBinding {
	return &AccessContextManagerGcpUserAccessBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerGcpUserAccessBinding)(nil)

// AccessContextManagerGcpUserAccessBinding represents the Terraform resource google_access_context_manager_gcp_user_access_binding.
type AccessContextManagerGcpUserAccessBinding struct {
	Name      string
	Args      AccessContextManagerGcpUserAccessBindingArgs
	state     *accessContextManagerGcpUserAccessBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerGcpUserAccessBinding].
func (acmguab *AccessContextManagerGcpUserAccessBinding) Type() string {
	return "google_access_context_manager_gcp_user_access_binding"
}

// LocalName returns the local name for [AccessContextManagerGcpUserAccessBinding].
func (acmguab *AccessContextManagerGcpUserAccessBinding) LocalName() string {
	return acmguab.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerGcpUserAccessBinding].
func (acmguab *AccessContextManagerGcpUserAccessBinding) Configuration() interface{} {
	return acmguab.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerGcpUserAccessBinding].
func (acmguab *AccessContextManagerGcpUserAccessBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(acmguab)
}

// Dependencies returns the list of resources [AccessContextManagerGcpUserAccessBinding] depends_on.
func (acmguab *AccessContextManagerGcpUserAccessBinding) Dependencies() terra.Dependencies {
	return acmguab.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerGcpUserAccessBinding].
func (acmguab *AccessContextManagerGcpUserAccessBinding) LifecycleManagement() *terra.Lifecycle {
	return acmguab.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerGcpUserAccessBinding].
func (acmguab *AccessContextManagerGcpUserAccessBinding) Attributes() accessContextManagerGcpUserAccessBindingAttributes {
	return accessContextManagerGcpUserAccessBindingAttributes{ref: terra.ReferenceResource(acmguab)}
}

// ImportState imports the given attribute values into [AccessContextManagerGcpUserAccessBinding]'s state.
func (acmguab *AccessContextManagerGcpUserAccessBinding) ImportState(av io.Reader) error {
	acmguab.state = &accessContextManagerGcpUserAccessBindingState{}
	if err := json.NewDecoder(av).Decode(acmguab.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmguab.Type(), acmguab.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerGcpUserAccessBinding] has state.
func (acmguab *AccessContextManagerGcpUserAccessBinding) State() (*accessContextManagerGcpUserAccessBindingState, bool) {
	return acmguab.state, acmguab.state != nil
}

// StateMust returns the state for [AccessContextManagerGcpUserAccessBinding]. Panics if the state is nil.
func (acmguab *AccessContextManagerGcpUserAccessBinding) StateMust() *accessContextManagerGcpUserAccessBindingState {
	if acmguab.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmguab.Type(), acmguab.LocalName()))
	}
	return acmguab.state
}

// AccessContextManagerGcpUserAccessBindingArgs contains the configurations for google_access_context_manager_gcp_user_access_binding.
type AccessContextManagerGcpUserAccessBindingArgs struct {
	// AccessLevels: list of string, required
	AccessLevels terra.ListValue[terra.StringValue] `hcl:"access_levels,attr" validate:"required"`
	// GroupKey: string, required
	GroupKey terra.StringValue `hcl:"group_key,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrganizationId: string, required
	OrganizationId terra.StringValue `hcl:"organization_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *accesscontextmanagergcpuseraccessbinding.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerGcpUserAccessBindingAttributes struct {
	ref terra.Reference
}

// AccessLevels returns a reference to field access_levels of google_access_context_manager_gcp_user_access_binding.
func (acmguab accessContextManagerGcpUserAccessBindingAttributes) AccessLevels() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](acmguab.ref.Append("access_levels"))
}

// GroupKey returns a reference to field group_key of google_access_context_manager_gcp_user_access_binding.
func (acmguab accessContextManagerGcpUserAccessBindingAttributes) GroupKey() terra.StringValue {
	return terra.ReferenceAsString(acmguab.ref.Append("group_key"))
}

// Id returns a reference to field id of google_access_context_manager_gcp_user_access_binding.
func (acmguab accessContextManagerGcpUserAccessBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmguab.ref.Append("id"))
}

// Name returns a reference to field name of google_access_context_manager_gcp_user_access_binding.
func (acmguab accessContextManagerGcpUserAccessBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmguab.ref.Append("name"))
}

// OrganizationId returns a reference to field organization_id of google_access_context_manager_gcp_user_access_binding.
func (acmguab accessContextManagerGcpUserAccessBindingAttributes) OrganizationId() terra.StringValue {
	return terra.ReferenceAsString(acmguab.ref.Append("organization_id"))
}

func (acmguab accessContextManagerGcpUserAccessBindingAttributes) Timeouts() accesscontextmanagergcpuseraccessbinding.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagergcpuseraccessbinding.TimeoutsAttributes](acmguab.ref.Append("timeouts"))
}

type accessContextManagerGcpUserAccessBindingState struct {
	AccessLevels   []string                                                `json:"access_levels"`
	GroupKey       string                                                  `json:"group_key"`
	Id             string                                                  `json:"id"`
	Name           string                                                  `json:"name"`
	OrganizationId string                                                  `json:"organization_id"`
	Timeouts       *accesscontextmanagergcpuseraccessbinding.TimeoutsState `json:"timeouts"`
}
