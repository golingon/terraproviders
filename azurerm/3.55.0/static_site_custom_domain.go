// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	staticsitecustomdomain "github.com/golingon/terraproviders/azurerm/3.55.0/staticsitecustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStaticSiteCustomDomain creates a new instance of [StaticSiteCustomDomain].
func NewStaticSiteCustomDomain(name string, args StaticSiteCustomDomainArgs) *StaticSiteCustomDomain {
	return &StaticSiteCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StaticSiteCustomDomain)(nil)

// StaticSiteCustomDomain represents the Terraform resource azurerm_static_site_custom_domain.
type StaticSiteCustomDomain struct {
	Name      string
	Args      StaticSiteCustomDomainArgs
	state     *staticSiteCustomDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StaticSiteCustomDomain].
func (sscd *StaticSiteCustomDomain) Type() string {
	return "azurerm_static_site_custom_domain"
}

// LocalName returns the local name for [StaticSiteCustomDomain].
func (sscd *StaticSiteCustomDomain) LocalName() string {
	return sscd.Name
}

// Configuration returns the configuration (args) for [StaticSiteCustomDomain].
func (sscd *StaticSiteCustomDomain) Configuration() interface{} {
	return sscd.Args
}

// DependOn is used for other resources to depend on [StaticSiteCustomDomain].
func (sscd *StaticSiteCustomDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(sscd)
}

// Dependencies returns the list of resources [StaticSiteCustomDomain] depends_on.
func (sscd *StaticSiteCustomDomain) Dependencies() terra.Dependencies {
	return sscd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StaticSiteCustomDomain].
func (sscd *StaticSiteCustomDomain) LifecycleManagement() *terra.Lifecycle {
	return sscd.Lifecycle
}

// Attributes returns the attributes for [StaticSiteCustomDomain].
func (sscd *StaticSiteCustomDomain) Attributes() staticSiteCustomDomainAttributes {
	return staticSiteCustomDomainAttributes{ref: terra.ReferenceResource(sscd)}
}

// ImportState imports the given attribute values into [StaticSiteCustomDomain]'s state.
func (sscd *StaticSiteCustomDomain) ImportState(av io.Reader) error {
	sscd.state = &staticSiteCustomDomainState{}
	if err := json.NewDecoder(av).Decode(sscd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sscd.Type(), sscd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StaticSiteCustomDomain] has state.
func (sscd *StaticSiteCustomDomain) State() (*staticSiteCustomDomainState, bool) {
	return sscd.state, sscd.state != nil
}

// StateMust returns the state for [StaticSiteCustomDomain]. Panics if the state is nil.
func (sscd *StaticSiteCustomDomain) StateMust() *staticSiteCustomDomainState {
	if sscd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sscd.Type(), sscd.LocalName()))
	}
	return sscd.state
}

// StaticSiteCustomDomainArgs contains the configurations for azurerm_static_site_custom_domain.
type StaticSiteCustomDomainArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StaticSiteId: string, required
	StaticSiteId terra.StringValue `hcl:"static_site_id,attr" validate:"required"`
	// ValidationType: string, optional
	ValidationType terra.StringValue `hcl:"validation_type,attr"`
	// Timeouts: optional
	Timeouts *staticsitecustomdomain.Timeouts `hcl:"timeouts,block"`
}
type staticSiteCustomDomainAttributes struct {
	ref terra.Reference
}

// DomainName returns a reference to field domain_name of azurerm_static_site_custom_domain.
func (sscd staticSiteCustomDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(sscd.ref.Append("domain_name"))
}

// Id returns a reference to field id of azurerm_static_site_custom_domain.
func (sscd staticSiteCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sscd.ref.Append("id"))
}

// StaticSiteId returns a reference to field static_site_id of azurerm_static_site_custom_domain.
func (sscd staticSiteCustomDomainAttributes) StaticSiteId() terra.StringValue {
	return terra.ReferenceAsString(sscd.ref.Append("static_site_id"))
}

// ValidationToken returns a reference to field validation_token of azurerm_static_site_custom_domain.
func (sscd staticSiteCustomDomainAttributes) ValidationToken() terra.StringValue {
	return terra.ReferenceAsString(sscd.ref.Append("validation_token"))
}

// ValidationType returns a reference to field validation_type of azurerm_static_site_custom_domain.
func (sscd staticSiteCustomDomainAttributes) ValidationType() terra.StringValue {
	return terra.ReferenceAsString(sscd.ref.Append("validation_type"))
}

func (sscd staticSiteCustomDomainAttributes) Timeouts() staticsitecustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[staticsitecustomdomain.TimeoutsAttributes](sscd.ref.Append("timeouts"))
}

type staticSiteCustomDomainState struct {
	DomainName      string                                `json:"domain_name"`
	Id              string                                `json:"id"`
	StaticSiteId    string                                `json:"static_site_id"`
	ValidationToken string                                `json:"validation_token"`
	ValidationType  string                                `json:"validation_type"`
	Timeouts        *staticsitecustomdomain.TimeoutsState `json:"timeouts"`
}