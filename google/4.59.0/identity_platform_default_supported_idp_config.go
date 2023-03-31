// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	identityplatformdefaultsupportedidpconfig "github.com/golingon/terraproviders/google/4.59.0/identityplatformdefaultsupportedidpconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIdentityPlatformDefaultSupportedIdpConfig(name string, args IdentityPlatformDefaultSupportedIdpConfigArgs) *IdentityPlatformDefaultSupportedIdpConfig {
	return &IdentityPlatformDefaultSupportedIdpConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IdentityPlatformDefaultSupportedIdpConfig)(nil)

type IdentityPlatformDefaultSupportedIdpConfig struct {
	Name  string
	Args  IdentityPlatformDefaultSupportedIdpConfigArgs
	state *identityPlatformDefaultSupportedIdpConfigState
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Type() string {
	return "google_identity_platform_default_supported_idp_config"
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) LocalName() string {
	return ipdsic.Name
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Configuration() interface{} {
	return ipdsic.Args
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) Attributes() identityPlatformDefaultSupportedIdpConfigAttributes {
	return identityPlatformDefaultSupportedIdpConfigAttributes{ref: terra.ReferenceResource(ipdsic)}
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) ImportState(av io.Reader) error {
	ipdsic.state = &identityPlatformDefaultSupportedIdpConfigState{}
	if err := json.NewDecoder(av).Decode(ipdsic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipdsic.Type(), ipdsic.LocalName(), err)
	}
	return nil
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) State() (*identityPlatformDefaultSupportedIdpConfigState, bool) {
	return ipdsic.state, ipdsic.state != nil
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) StateMust() *identityPlatformDefaultSupportedIdpConfigState {
	if ipdsic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipdsic.Type(), ipdsic.LocalName()))
	}
	return ipdsic.state
}

func (ipdsic *IdentityPlatformDefaultSupportedIdpConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(ipdsic)
}

type IdentityPlatformDefaultSupportedIdpConfigArgs struct {
	// ClientId: string, required
	ClientId terra.StringValue `hcl:"client_id,attr" validate:"required"`
	// ClientSecret: string, required
	ClientSecret terra.StringValue `hcl:"client_secret,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdpId: string, required
	IdpId terra.StringValue `hcl:"idp_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *identityplatformdefaultsupportedidpconfig.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that IdentityPlatformDefaultSupportedIdpConfig depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type identityPlatformDefaultSupportedIdpConfigAttributes struct {
	ref terra.Reference
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) ClientId() terra.StringValue {
	return terra.ReferenceString(ipdsic.ref.Append("client_id"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceString(ipdsic.ref.Append("client_secret"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(ipdsic.ref.Append("enabled"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ipdsic.ref.Append("id"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) IdpId() terra.StringValue {
	return terra.ReferenceString(ipdsic.ref.Append("idp_id"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ipdsic.ref.Append("name"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Project() terra.StringValue {
	return terra.ReferenceString(ipdsic.ref.Append("project"))
}

func (ipdsic identityPlatformDefaultSupportedIdpConfigAttributes) Timeouts() identityplatformdefaultsupportedidpconfig.TimeoutsAttributes {
	return terra.ReferenceSingle[identityplatformdefaultsupportedidpconfig.TimeoutsAttributes](ipdsic.ref.Append("timeouts"))
}

type identityPlatformDefaultSupportedIdpConfigState struct {
	ClientId     string                                                   `json:"client_id"`
	ClientSecret string                                                   `json:"client_secret"`
	Enabled      bool                                                     `json:"enabled"`
	Id           string                                                   `json:"id"`
	IdpId        string                                                   `json:"idp_id"`
	Name         string                                                   `json:"name"`
	Project      string                                                   `json:"project"`
	Timeouts     *identityplatformdefaultsupportedidpconfig.TimeoutsState `json:"timeouts"`
}
