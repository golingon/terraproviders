// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	nginxcertificate "github.com/golingon/terraproviders/azurerm/3.67.0/nginxcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNginxCertificate creates a new instance of [NginxCertificate].
func NewNginxCertificate(name string, args NginxCertificateArgs) *NginxCertificate {
	return &NginxCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NginxCertificate)(nil)

// NginxCertificate represents the Terraform resource azurerm_nginx_certificate.
type NginxCertificate struct {
	Name      string
	Args      NginxCertificateArgs
	state     *nginxCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NginxCertificate].
func (nc *NginxCertificate) Type() string {
	return "azurerm_nginx_certificate"
}

// LocalName returns the local name for [NginxCertificate].
func (nc *NginxCertificate) LocalName() string {
	return nc.Name
}

// Configuration returns the configuration (args) for [NginxCertificate].
func (nc *NginxCertificate) Configuration() interface{} {
	return nc.Args
}

// DependOn is used for other resources to depend on [NginxCertificate].
func (nc *NginxCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(nc)
}

// Dependencies returns the list of resources [NginxCertificate] depends_on.
func (nc *NginxCertificate) Dependencies() terra.Dependencies {
	return nc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NginxCertificate].
func (nc *NginxCertificate) LifecycleManagement() *terra.Lifecycle {
	return nc.Lifecycle
}

// Attributes returns the attributes for [NginxCertificate].
func (nc *NginxCertificate) Attributes() nginxCertificateAttributes {
	return nginxCertificateAttributes{ref: terra.ReferenceResource(nc)}
}

// ImportState imports the given attribute values into [NginxCertificate]'s state.
func (nc *NginxCertificate) ImportState(av io.Reader) error {
	nc.state = &nginxCertificateState{}
	if err := json.NewDecoder(av).Decode(nc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nc.Type(), nc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NginxCertificate] has state.
func (nc *NginxCertificate) State() (*nginxCertificateState, bool) {
	return nc.state, nc.state != nil
}

// StateMust returns the state for [NginxCertificate]. Panics if the state is nil.
func (nc *NginxCertificate) StateMust() *nginxCertificateState {
	if nc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nc.Type(), nc.LocalName()))
	}
	return nc.state
}

// NginxCertificateArgs contains the configurations for azurerm_nginx_certificate.
type NginxCertificateArgs struct {
	// CertificateVirtualPath: string, required
	CertificateVirtualPath terra.StringValue `hcl:"certificate_virtual_path,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyVaultSecretId: string, required
	KeyVaultSecretId terra.StringValue `hcl:"key_vault_secret_id,attr" validate:"required"`
	// KeyVirtualPath: string, required
	KeyVirtualPath terra.StringValue `hcl:"key_virtual_path,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NginxDeploymentId: string, required
	NginxDeploymentId terra.StringValue `hcl:"nginx_deployment_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *nginxcertificate.Timeouts `hcl:"timeouts,block"`
}
type nginxCertificateAttributes struct {
	ref terra.Reference
}

// CertificateVirtualPath returns a reference to field certificate_virtual_path of azurerm_nginx_certificate.
func (nc nginxCertificateAttributes) CertificateVirtualPath() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("certificate_virtual_path"))
}

// Id returns a reference to field id of azurerm_nginx_certificate.
func (nc nginxCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("id"))
}

// KeyVaultSecretId returns a reference to field key_vault_secret_id of azurerm_nginx_certificate.
func (nc nginxCertificateAttributes) KeyVaultSecretId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("key_vault_secret_id"))
}

// KeyVirtualPath returns a reference to field key_virtual_path of azurerm_nginx_certificate.
func (nc nginxCertificateAttributes) KeyVirtualPath() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("key_virtual_path"))
}

// Name returns a reference to field name of azurerm_nginx_certificate.
func (nc nginxCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("name"))
}

// NginxDeploymentId returns a reference to field nginx_deployment_id of azurerm_nginx_certificate.
func (nc nginxCertificateAttributes) NginxDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("nginx_deployment_id"))
}

func (nc nginxCertificateAttributes) Timeouts() nginxcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[nginxcertificate.TimeoutsAttributes](nc.ref.Append("timeouts"))
}

type nginxCertificateState struct {
	CertificateVirtualPath string                          `json:"certificate_virtual_path"`
	Id                     string                          `json:"id"`
	KeyVaultSecretId       string                          `json:"key_vault_secret_id"`
	KeyVirtualPath         string                          `json:"key_virtual_path"`
	Name                   string                          `json:"name"`
	NginxDeploymentId      string                          `json:"nginx_deployment_id"`
	Timeouts               *nginxcertificate.TimeoutsState `json:"timeouts"`
}
