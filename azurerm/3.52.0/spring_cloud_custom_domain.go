// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudcustomdomain "github.com/golingon/terraproviders/azurerm/3.52.0/springcloudcustomdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudCustomDomain creates a new instance of [SpringCloudCustomDomain].
func NewSpringCloudCustomDomain(name string, args SpringCloudCustomDomainArgs) *SpringCloudCustomDomain {
	return &SpringCloudCustomDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudCustomDomain)(nil)

// SpringCloudCustomDomain represents the Terraform resource azurerm_spring_cloud_custom_domain.
type SpringCloudCustomDomain struct {
	Name      string
	Args      SpringCloudCustomDomainArgs
	state     *springCloudCustomDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudCustomDomain].
func (sccd *SpringCloudCustomDomain) Type() string {
	return "azurerm_spring_cloud_custom_domain"
}

// LocalName returns the local name for [SpringCloudCustomDomain].
func (sccd *SpringCloudCustomDomain) LocalName() string {
	return sccd.Name
}

// Configuration returns the configuration (args) for [SpringCloudCustomDomain].
func (sccd *SpringCloudCustomDomain) Configuration() interface{} {
	return sccd.Args
}

// DependOn is used for other resources to depend on [SpringCloudCustomDomain].
func (sccd *SpringCloudCustomDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(sccd)
}

// Dependencies returns the list of resources [SpringCloudCustomDomain] depends_on.
func (sccd *SpringCloudCustomDomain) Dependencies() terra.Dependencies {
	return sccd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudCustomDomain].
func (sccd *SpringCloudCustomDomain) LifecycleManagement() *terra.Lifecycle {
	return sccd.Lifecycle
}

// Attributes returns the attributes for [SpringCloudCustomDomain].
func (sccd *SpringCloudCustomDomain) Attributes() springCloudCustomDomainAttributes {
	return springCloudCustomDomainAttributes{ref: terra.ReferenceResource(sccd)}
}

// ImportState imports the given attribute values into [SpringCloudCustomDomain]'s state.
func (sccd *SpringCloudCustomDomain) ImportState(av io.Reader) error {
	sccd.state = &springCloudCustomDomainState{}
	if err := json.NewDecoder(av).Decode(sccd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sccd.Type(), sccd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudCustomDomain] has state.
func (sccd *SpringCloudCustomDomain) State() (*springCloudCustomDomainState, bool) {
	return sccd.state, sccd.state != nil
}

// StateMust returns the state for [SpringCloudCustomDomain]. Panics if the state is nil.
func (sccd *SpringCloudCustomDomain) StateMust() *springCloudCustomDomainState {
	if sccd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sccd.Type(), sccd.LocalName()))
	}
	return sccd.state
}

// SpringCloudCustomDomainArgs contains the configurations for azurerm_spring_cloud_custom_domain.
type SpringCloudCustomDomainArgs struct {
	// CertificateName: string, optional
	CertificateName terra.StringValue `hcl:"certificate_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudAppId: string, required
	SpringCloudAppId terra.StringValue `hcl:"spring_cloud_app_id,attr" validate:"required"`
	// Thumbprint: string, optional
	Thumbprint terra.StringValue `hcl:"thumbprint,attr"`
	// Timeouts: optional
	Timeouts *springcloudcustomdomain.Timeouts `hcl:"timeouts,block"`
}
type springCloudCustomDomainAttributes struct {
	ref terra.Reference
}

// CertificateName returns a reference to field certificate_name of azurerm_spring_cloud_custom_domain.
func (sccd springCloudCustomDomainAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("certificate_name"))
}

// Id returns a reference to field id of azurerm_spring_cloud_custom_domain.
func (sccd springCloudCustomDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_custom_domain.
func (sccd springCloudCustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("name"))
}

// SpringCloudAppId returns a reference to field spring_cloud_app_id of azurerm_spring_cloud_custom_domain.
func (sccd springCloudCustomDomainAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("spring_cloud_app_id"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_spring_cloud_custom_domain.
func (sccd springCloudCustomDomainAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("thumbprint"))
}

func (sccd springCloudCustomDomainAttributes) Timeouts() springcloudcustomdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudcustomdomain.TimeoutsAttributes](sccd.ref.Append("timeouts"))
}

type springCloudCustomDomainState struct {
	CertificateName  string                                 `json:"certificate_name"`
	Id               string                                 `json:"id"`
	Name             string                                 `json:"name"`
	SpringCloudAppId string                                 `json:"spring_cloud_app_id"`
	Thumbprint       string                                 `json:"thumbprint"`
	Timeouts         *springcloudcustomdomain.TimeoutsState `json:"timeouts"`
}
