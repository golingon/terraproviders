// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	identityplatformtenant "github.com/golingon/terraproviders/googlebeta/4.76.0/identityplatformtenant"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIdentityPlatformTenant creates a new instance of [IdentityPlatformTenant].
func NewIdentityPlatformTenant(name string, args IdentityPlatformTenantArgs) *IdentityPlatformTenant {
	return &IdentityPlatformTenant{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformTenant)(nil)

// IdentityPlatformTenant represents the Terraform resource google_identity_platform_tenant.
type IdentityPlatformTenant struct {
	Name      string
	Args      IdentityPlatformTenantArgs
	state     *identityPlatformTenantState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IdentityPlatformTenant].
func (ipt *IdentityPlatformTenant) Type() string {
	return "google_identity_platform_tenant"
}

// LocalName returns the local name for [IdentityPlatformTenant].
func (ipt *IdentityPlatformTenant) LocalName() string {
	return ipt.Name
}

// Configuration returns the configuration (args) for [IdentityPlatformTenant].
func (ipt *IdentityPlatformTenant) Configuration() interface{} {
	return ipt.Args
}

// DependOn is used for other resources to depend on [IdentityPlatformTenant].
func (ipt *IdentityPlatformTenant) DependOn() terra.Reference {
	return terra.ReferenceResource(ipt)
}

// Dependencies returns the list of resources [IdentityPlatformTenant] depends_on.
func (ipt *IdentityPlatformTenant) Dependencies() terra.Dependencies {
	return ipt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IdentityPlatformTenant].
func (ipt *IdentityPlatformTenant) LifecycleManagement() *terra.Lifecycle {
	return ipt.Lifecycle
}

// Attributes returns the attributes for [IdentityPlatformTenant].
func (ipt *IdentityPlatformTenant) Attributes() identityPlatformTenantAttributes {
	return identityPlatformTenantAttributes{ref: terra.ReferenceResource(ipt)}
}

// ImportState imports the given attribute values into [IdentityPlatformTenant]'s state.
func (ipt *IdentityPlatformTenant) ImportState(av io.Reader) error {
	ipt.state = &identityPlatformTenantState{}
	if err := json.NewDecoder(av).Decode(ipt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipt.Type(), ipt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IdentityPlatformTenant] has state.
func (ipt *IdentityPlatformTenant) State() (*identityPlatformTenantState, bool) {
	return ipt.state, ipt.state != nil
}

// StateMust returns the state for [IdentityPlatformTenant]. Panics if the state is nil.
func (ipt *IdentityPlatformTenant) StateMust() *identityPlatformTenantState {
	if ipt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipt.Type(), ipt.LocalName()))
	}
	return ipt.state
}

// IdentityPlatformTenantArgs contains the configurations for google_identity_platform_tenant.
type IdentityPlatformTenantArgs struct {
	// AllowPasswordSignup: bool, optional
	AllowPasswordSignup terra.BoolValue `hcl:"allow_password_signup,attr"`
	// DisableAuth: bool, optional
	DisableAuth terra.BoolValue `hcl:"disable_auth,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableEmailLinkSignin: bool, optional
	EnableEmailLinkSignin terra.BoolValue `hcl:"enable_email_link_signin,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *identityplatformtenant.Timeouts `hcl:"timeouts,block"`
}
type identityPlatformTenantAttributes struct {
	ref terra.Reference
}

// AllowPasswordSignup returns a reference to field allow_password_signup of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) AllowPasswordSignup() terra.BoolValue {
	return terra.ReferenceAsBool(ipt.ref.Append("allow_password_signup"))
}

// DisableAuth returns a reference to field disable_auth of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) DisableAuth() terra.BoolValue {
	return terra.ReferenceAsBool(ipt.ref.Append("disable_auth"))
}

// DisplayName returns a reference to field display_name of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("display_name"))
}

// EnableEmailLinkSignin returns a reference to field enable_email_link_signin of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) EnableEmailLinkSignin() terra.BoolValue {
	return terra.ReferenceAsBool(ipt.ref.Append("enable_email_link_signin"))
}

// Id returns a reference to field id of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("id"))
}

// Name returns a reference to field name of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("name"))
}

// Project returns a reference to field project of google_identity_platform_tenant.
func (ipt identityPlatformTenantAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("project"))
}

func (ipt identityPlatformTenantAttributes) Timeouts() identityplatformtenant.TimeoutsAttributes {
	return terra.ReferenceAsSingle[identityplatformtenant.TimeoutsAttributes](ipt.ref.Append("timeouts"))
}

type identityPlatformTenantState struct {
	AllowPasswordSignup   bool                                  `json:"allow_password_signup"`
	DisableAuth           bool                                  `json:"disable_auth"`
	DisplayName           string                                `json:"display_name"`
	EnableEmailLinkSignin bool                                  `json:"enable_email_link_signin"`
	Id                    string                                `json:"id"`
	Name                  string                                `json:"name"`
	Project               string                                `json:"project"`
	Timeouts              *identityplatformtenant.TimeoutsState `json:"timeouts"`
}
