// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubcertificate "github.com/golingon/terraproviders/azurerm/3.65.0/iothubcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubCertificate creates a new instance of [IothubCertificate].
func NewIothubCertificate(name string, args IothubCertificateArgs) *IothubCertificate {
	return &IothubCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubCertificate)(nil)

// IothubCertificate represents the Terraform resource azurerm_iothub_certificate.
type IothubCertificate struct {
	Name      string
	Args      IothubCertificateArgs
	state     *iothubCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubCertificate].
func (ic *IothubCertificate) Type() string {
	return "azurerm_iothub_certificate"
}

// LocalName returns the local name for [IothubCertificate].
func (ic *IothubCertificate) LocalName() string {
	return ic.Name
}

// Configuration returns the configuration (args) for [IothubCertificate].
func (ic *IothubCertificate) Configuration() interface{} {
	return ic.Args
}

// DependOn is used for other resources to depend on [IothubCertificate].
func (ic *IothubCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(ic)
}

// Dependencies returns the list of resources [IothubCertificate] depends_on.
func (ic *IothubCertificate) Dependencies() terra.Dependencies {
	return ic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubCertificate].
func (ic *IothubCertificate) LifecycleManagement() *terra.Lifecycle {
	return ic.Lifecycle
}

// Attributes returns the attributes for [IothubCertificate].
func (ic *IothubCertificate) Attributes() iothubCertificateAttributes {
	return iothubCertificateAttributes{ref: terra.ReferenceResource(ic)}
}

// ImportState imports the given attribute values into [IothubCertificate]'s state.
func (ic *IothubCertificate) ImportState(av io.Reader) error {
	ic.state = &iothubCertificateState{}
	if err := json.NewDecoder(av).Decode(ic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ic.Type(), ic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubCertificate] has state.
func (ic *IothubCertificate) State() (*iothubCertificateState, bool) {
	return ic.state, ic.state != nil
}

// StateMust returns the state for [IothubCertificate]. Panics if the state is nil.
func (ic *IothubCertificate) StateMust() *iothubCertificateState {
	if ic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ic.Type(), ic.LocalName()))
	}
	return ic.state
}

// IothubCertificateArgs contains the configurations for azurerm_iothub_certificate.
type IothubCertificateArgs struct {
	// CertificateContent: string, required
	CertificateContent terra.StringValue `hcl:"certificate_content,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IothubName: string, required
	IothubName terra.StringValue `hcl:"iothub_name,attr" validate:"required"`
	// IsVerified: bool, optional
	IsVerified terra.BoolValue `hcl:"is_verified,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iothubcertificate.Timeouts `hcl:"timeouts,block"`
}
type iothubCertificateAttributes struct {
	ref terra.Reference
}

// CertificateContent returns a reference to field certificate_content of azurerm_iothub_certificate.
func (ic iothubCertificateAttributes) CertificateContent() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("certificate_content"))
}

// Id returns a reference to field id of azurerm_iothub_certificate.
func (ic iothubCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("id"))
}

// IothubName returns a reference to field iothub_name of azurerm_iothub_certificate.
func (ic iothubCertificateAttributes) IothubName() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("iothub_name"))
}

// IsVerified returns a reference to field is_verified of azurerm_iothub_certificate.
func (ic iothubCertificateAttributes) IsVerified() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("is_verified"))
}

// Name returns a reference to field name of azurerm_iothub_certificate.
func (ic iothubCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_certificate.
func (ic iothubCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("resource_group_name"))
}

func (ic iothubCertificateAttributes) Timeouts() iothubcertificate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubcertificate.TimeoutsAttributes](ic.ref.Append("timeouts"))
}

type iothubCertificateState struct {
	CertificateContent string                           `json:"certificate_content"`
	Id                 string                           `json:"id"`
	IothubName         string                           `json:"iothub_name"`
	IsVerified         bool                             `json:"is_verified"`
	Name               string                           `json:"name"`
	ResourceGroupName  string                           `json:"resource_group_name"`
	Timeouts           *iothubcertificate.TimeoutsState `json:"timeouts"`
}
