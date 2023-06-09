// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnendpointcustomdomain "github.com/golingon/terraproviders/azurerm/3.49.0/cdnendpointcustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnEndpointCustomDomain creates a new instance of [CdnEndpointCustomDomain].
func NewCdnEndpointCustomDomain(name string, args CdnEndpointCustomDomainArgs) *CdnEndpointCustomDomain {
	return &CdnEndpointCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnEndpointCustomDomain)(nil)

// CdnEndpointCustomDomain represents the Terraform resource azurerm_cdn_endpoint_custom_domain.
type CdnEndpointCustomDomain struct {
	Name      string
	Args      CdnEndpointCustomDomainArgs
	state     *cdnEndpointCustomDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnEndpointCustomDomain].
func (cecd *CdnEndpointCustomDomain) Type() string {
	return "azurerm_cdn_endpoint_custom_domain"
}

// LocalName returns the local name for [CdnEndpointCustomDomain].
func (cecd *CdnEndpointCustomDomain) LocalName() string {
	return cecd.Name
}

// Configuration returns the configuration (args) for [CdnEndpointCustomDomain].
func (cecd *CdnEndpointCustomDomain) Configuration() interface{} {
	return cecd.Args
}

// DependOn is used for other resources to depend on [CdnEndpointCustomDomain].
func (cecd *CdnEndpointCustomDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(cecd)
}

// Dependencies returns the list of resources [CdnEndpointCustomDomain] depends_on.
func (cecd *CdnEndpointCustomDomain) Dependencies() terra.Dependencies {
	return cecd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnEndpointCustomDomain].
func (cecd *CdnEndpointCustomDomain) LifecycleManagement() *terra.Lifecycle {
	return cecd.Lifecycle
}

// Attributes returns the attributes for [CdnEndpointCustomDomain].
func (cecd *CdnEndpointCustomDomain) Attributes() cdnEndpointCustomDomainAttributes {
	return cdnEndpointCustomDomainAttributes{ref: terra.ReferenceResource(cecd)}
}

// ImportState imports the given attribute values into [CdnEndpointCustomDomain]'s state.
func (cecd *CdnEndpointCustomDomain) ImportState(av io.Reader) error {
	cecd.state = &cdnEndpointCustomDomainState{}
	if err := json.NewDecoder(av).Decode(cecd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cecd.Type(), cecd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnEndpointCustomDomain] has state.
func (cecd *CdnEndpointCustomDomain) State() (*cdnEndpointCustomDomainState, bool) {
	return cecd.state, cecd.state != nil
}

// StateMust returns the state for [CdnEndpointCustomDomain]. Panics if the state is nil.
func (cecd *CdnEndpointCustomDomain) StateMust() *cdnEndpointCustomDomainState {
	if cecd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cecd.Type(), cecd.LocalName()))
	}
	return cecd.state
}

// CdnEndpointCustomDomainArgs contains the configurations for azurerm_cdn_endpoint_custom_domain.
type CdnEndpointCustomDomainArgs struct {
	// CdnEndpointId: string, required
	CdnEndpointId terra.StringValue `hcl:"cdn_endpoint_id,attr" validate:"required"`
	// HostName: string, required
	HostName terra.StringValue `hcl:"host_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// CdnManagedHttps: optional
	CdnManagedHttps *cdnendpointcustomdomain.CdnManagedHttps `hcl:"cdn_managed_https,block"`
	// Timeouts: optional
	Timeouts *cdnendpointcustomdomain.Timeouts `hcl:"timeouts,block"`
	// UserManagedHttps: optional
	UserManagedHttps *cdnendpointcustomdomain.UserManagedHttps `hcl:"user_managed_https,block"`
}
type cdnEndpointCustomDomainAttributes struct {
	ref terra.Reference
}

// CdnEndpointId returns a reference to field cdn_endpoint_id of azurerm_cdn_endpoint_custom_domain.
func (cecd cdnEndpointCustomDomainAttributes) CdnEndpointId() terra.StringValue {
	return terra.ReferenceAsString(cecd.ref.Append("cdn_endpoint_id"))
}

// HostName returns a reference to field host_name of azurerm_cdn_endpoint_custom_domain.
func (cecd cdnEndpointCustomDomainAttributes) HostName() terra.StringValue {
	return terra.ReferenceAsString(cecd.ref.Append("host_name"))
}

// Id returns a reference to field id of azurerm_cdn_endpoint_custom_domain.
func (cecd cdnEndpointCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cecd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_cdn_endpoint_custom_domain.
func (cecd cdnEndpointCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cecd.ref.Append("name"))
}

func (cecd cdnEndpointCustomDomainAttributes) CdnManagedHttps() terra.ListValue[cdnendpointcustomdomain.CdnManagedHttpsAttributes] {
	return terra.ReferenceAsList[cdnendpointcustomdomain.CdnManagedHttpsAttributes](cecd.ref.Append("cdn_managed_https"))
}

func (cecd cdnEndpointCustomDomainAttributes) Timeouts() cdnendpointcustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnendpointcustomdomain.TimeoutsAttributes](cecd.ref.Append("timeouts"))
}

func (cecd cdnEndpointCustomDomainAttributes) UserManagedHttps() terra.ListValue[cdnendpointcustomdomain.UserManagedHttpsAttributes] {
	return terra.ReferenceAsList[cdnendpointcustomdomain.UserManagedHttpsAttributes](cecd.ref.Append("user_managed_https"))
}

type cdnEndpointCustomDomainState struct {
	CdnEndpointId    string                                          `json:"cdn_endpoint_id"`
	HostName         string                                          `json:"host_name"`
	Id               string                                          `json:"id"`
	Name             string                                          `json:"name"`
	CdnManagedHttps  []cdnendpointcustomdomain.CdnManagedHttpsState  `json:"cdn_managed_https"`
	Timeouts         *cdnendpointcustomdomain.TimeoutsState          `json:"timeouts"`
	UserManagedHttps []cdnendpointcustomdomain.UserManagedHttpsState `json:"user_managed_https"`
}
