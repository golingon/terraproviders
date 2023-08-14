// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudcertificate "github.com/golingon/terraproviders/azurerm/3.69.0/springcloudcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudCertificate creates a new instance of [SpringCloudCertificate].
func NewSpringCloudCertificate(name string, args SpringCloudCertificateArgs) *SpringCloudCertificate {
	return &SpringCloudCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudCertificate)(nil)

// SpringCloudCertificate represents the Terraform resource azurerm_spring_cloud_certificate.
type SpringCloudCertificate struct {
	Name      string
	Args      SpringCloudCertificateArgs
	state     *springCloudCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudCertificate].
func (scc *SpringCloudCertificate) Type() string {
	return "azurerm_spring_cloud_certificate"
}

// LocalName returns the local name for [SpringCloudCertificate].
func (scc *SpringCloudCertificate) LocalName() string {
	return scc.Name
}

// Configuration returns the configuration (args) for [SpringCloudCertificate].
func (scc *SpringCloudCertificate) Configuration() interface{} {
	return scc.Args
}

// DependOn is used for other resources to depend on [SpringCloudCertificate].
func (scc *SpringCloudCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(scc)
}

// Dependencies returns the list of resources [SpringCloudCertificate] depends_on.
func (scc *SpringCloudCertificate) Dependencies() terra.Dependencies {
	return scc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudCertificate].
func (scc *SpringCloudCertificate) LifecycleManagement() *terra.Lifecycle {
	return scc.Lifecycle
}

// Attributes returns the attributes for [SpringCloudCertificate].
func (scc *SpringCloudCertificate) Attributes() springCloudCertificateAttributes {
	return springCloudCertificateAttributes{ref: terra.ReferenceResource(scc)}
}

// ImportState imports the given attribute values into [SpringCloudCertificate]'s state.
func (scc *SpringCloudCertificate) ImportState(av io.Reader) error {
	scc.state = &springCloudCertificateState{}
	if err := json.NewDecoder(av).Decode(scc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scc.Type(), scc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudCertificate] has state.
func (scc *SpringCloudCertificate) State() (*springCloudCertificateState, bool) {
	return scc.state, scc.state != nil
}

// StateMust returns the state for [SpringCloudCertificate]. Panics if the state is nil.
func (scc *SpringCloudCertificate) StateMust() *springCloudCertificateState {
	if scc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scc.Type(), scc.LocalName()))
	}
	return scc.state
}

// SpringCloudCertificateArgs contains the configurations for azurerm_spring_cloud_certificate.
type SpringCloudCertificateArgs struct {
	// CertificateContent: string, optional
	CertificateContent terra.StringValue `hcl:"certificate_content,attr"`
	// ExcludePrivateKey: bool, optional
	ExcludePrivateKey terra.BoolValue `hcl:"exclude_private_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultCertificateId: string, optional
	KeyVaultCertificateId terra.StringValue `hcl:"key_vault_certificate_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudcertificate.Timeouts `hcl:"timeouts,block"`
}
type springCloudCertificateAttributes struct {
	ref terra.Reference
}

// CertificateContent returns a reference to field certificate_content of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) CertificateContent() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("certificate_content"))
}

// ExcludePrivateKey returns a reference to field exclude_private_key of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) ExcludePrivateKey() terra.BoolValue {
	return terra.ReferenceAsBool(scc.ref.Append("exclude_private_key"))
}

// Id returns a reference to field id of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("id"))
}

// KeyVaultCertificateId returns a reference to field key_vault_certificate_id of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) KeyVaultCertificateId() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("key_vault_certificate_id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("resource_group_name"))
}

// ServiceName returns a reference to field service_name of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("service_name"))
}

// Thumbprint returns a reference to field thumbprint of azurerm_spring_cloud_certificate.
func (scc springCloudCertificateAttributes) Thumbprint() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("thumbprint"))
}

func (scc springCloudCertificateAttributes) Timeouts() springcloudcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudcertificate.TimeoutsAttributes](scc.ref.Append("timeouts"))
}

type springCloudCertificateState struct {
	CertificateContent    string                                `json:"certificate_content"`
	ExcludePrivateKey     bool                                  `json:"exclude_private_key"`
	Id                    string                                `json:"id"`
	KeyVaultCertificateId string                                `json:"key_vault_certificate_id"`
	Name                  string                                `json:"name"`
	ResourceGroupName     string                                `json:"resource_group_name"`
	ServiceName           string                                `json:"service_name"`
	Thumbprint            string                                `json:"thumbprint"`
	Timeouts              *springcloudcertificate.TimeoutsState `json:"timeouts"`
}
