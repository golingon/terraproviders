// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudapiportalcustomdomain "github.com/golingon/terraproviders/azurerm/3.64.0/springcloudapiportalcustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudApiPortalCustomDomain creates a new instance of [SpringCloudApiPortalCustomDomain].
func NewSpringCloudApiPortalCustomDomain(name string, args SpringCloudApiPortalCustomDomainArgs) *SpringCloudApiPortalCustomDomain {
	return &SpringCloudApiPortalCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudApiPortalCustomDomain)(nil)

// SpringCloudApiPortalCustomDomain represents the Terraform resource azurerm_spring_cloud_api_portal_custom_domain.
type SpringCloudApiPortalCustomDomain struct {
	Name      string
	Args      SpringCloudApiPortalCustomDomainArgs
	state     *springCloudApiPortalCustomDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudApiPortalCustomDomain].
func (scapcd *SpringCloudApiPortalCustomDomain) Type() string {
	return "azurerm_spring_cloud_api_portal_custom_domain"
}

// LocalName returns the local name for [SpringCloudApiPortalCustomDomain].
func (scapcd *SpringCloudApiPortalCustomDomain) LocalName() string {
	return scapcd.Name
}

// Configuration returns the configuration (args) for [SpringCloudApiPortalCustomDomain].
func (scapcd *SpringCloudApiPortalCustomDomain) Configuration() interface{} {
	return scapcd.Args
}

// DependOn is used for other resources to depend on [SpringCloudApiPortalCustomDomain].
func (scapcd *SpringCloudApiPortalCustomDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(scapcd)
}

// Dependencies returns the list of resources [SpringCloudApiPortalCustomDomain] depends_on.
func (scapcd *SpringCloudApiPortalCustomDomain) Dependencies() terra.Dependencies {
	return scapcd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudApiPortalCustomDomain].
func (scapcd *SpringCloudApiPortalCustomDomain) LifecycleManagement() *terra.Lifecycle {
	return scapcd.Lifecycle
}

// Attributes returns the attributes for [SpringCloudApiPortalCustomDomain].
func (scapcd *SpringCloudApiPortalCustomDomain) Attributes() springCloudApiPortalCustomDomainAttributes {
	return springCloudApiPortalCustomDomainAttributes{ref: terra.ReferenceResource(scapcd)}
}

// ImportState imports the given attribute values into [SpringCloudApiPortalCustomDomain]'s state.
func (scapcd *SpringCloudApiPortalCustomDomain) ImportState(av io.Reader) error {
	scapcd.state = &springCloudApiPortalCustomDomainState{}
	if err := json.NewDecoder(av).Decode(scapcd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scapcd.Type(), scapcd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudApiPortalCustomDomain] has state.
func (scapcd *SpringCloudApiPortalCustomDomain) State() (*springCloudApiPortalCustomDomainState, bool) {
	return scapcd.state, scapcd.state != nil
}

// StateMust returns the state for [SpringCloudApiPortalCustomDomain]. Panics if the state is nil.
func (scapcd *SpringCloudApiPortalCustomDomain) StateMust() *springCloudApiPortalCustomDomainState {
	if scapcd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scapcd.Type(), scapcd.LocalName()))
	}
	return scapcd.state
}

// SpringCloudApiPortalCustomDomainArgs contains the configurations for azurerm_spring_cloud_api_portal_custom_domain.
type SpringCloudApiPortalCustomDomainArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudApiPortalId: string, required
	SpringCloudApiPortalId terra.StringValue `hcl:"spring_cloud_api_portal_id,attr" validate:"required"`
	// Thumbprint: string, optional
	Thumbprint terra.StringValue `hcl:"thumbprint,attr"`
	// Timeouts: optional
	Timeouts *springcloudapiportalcustomdomain.Timeouts `hcl:"timeouts,block"`
}
type springCloudApiPortalCustomDomainAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_api_portal_custom_domain.
func (scapcd springCloudApiPortalCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scapcd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_api_portal_custom_domain.
func (scapcd springCloudApiPortalCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scapcd.ref.Append("name"))
}

// SpringCloudApiPortalId returns a reference to field spring_cloud_api_portal_id of azurerm_spring_cloud_api_portal_custom_domain.
func (scapcd springCloudApiPortalCustomDomainAttributes) SpringCloudApiPortalId() terra.StringValue {
	return terra.ReferenceAsString(scapcd.ref.Append("spring_cloud_api_portal_id"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_spring_cloud_api_portal_custom_domain.
func (scapcd springCloudApiPortalCustomDomainAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(scapcd.ref.Append("thumbprint"))
}

func (scapcd springCloudApiPortalCustomDomainAttributes) Timeouts() springcloudapiportalcustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudapiportalcustomdomain.TimeoutsAttributes](scapcd.ref.Append("timeouts"))
}

type springCloudApiPortalCustomDomainState struct {
	Id                     string                                          `json:"id"`
	Name                   string                                          `json:"name"`
	SpringCloudApiPortalId string                                          `json:"spring_cloud_api_portal_id"`
	Thumbprint             string                                          `json:"thumbprint"`
	Timeouts               *springcloudapiportalcustomdomain.TimeoutsState `json:"timeouts"`
}
