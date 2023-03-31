// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubcertificate "github.com/golingon/terraproviders/azurerm/3.49.0/iothubcertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewIothubCertificate(name string, args IothubCertificateArgs) *IothubCertificate {
	return &IothubCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubCertificate)(nil)

type IothubCertificate struct {
	Name  string
	Args  IothubCertificateArgs
	state *iothubCertificateState
}

func (ic *IothubCertificate) Type() string {
	return "azurerm_iothub_certificate"
}

func (ic *IothubCertificate) LocalName() string {
	return ic.Name
}

func (ic *IothubCertificate) Configuration() interface{} {
	return ic.Args
}

func (ic *IothubCertificate) Attributes() iothubCertificateAttributes {
	return iothubCertificateAttributes{ref: terra.ReferenceResource(ic)}
}

func (ic *IothubCertificate) ImportState(av io.Reader) error {
	ic.state = &iothubCertificateState{}
	if err := json.NewDecoder(av).Decode(ic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ic.Type(), ic.LocalName(), err)
	}
	return nil
}

func (ic *IothubCertificate) State() (*iothubCertificateState, bool) {
	return ic.state, ic.state != nil
}

func (ic *IothubCertificate) StateMust() *iothubCertificateState {
	if ic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ic.Type(), ic.LocalName()))
	}
	return ic.state
}

func (ic *IothubCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(ic)
}

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
	// DependsOn contains resources that IothubCertificate depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type iothubCertificateAttributes struct {
	ref terra.Reference
}

func (ic iothubCertificateAttributes) CertificateContent() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("certificate_content"))
}

func (ic iothubCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("id"))
}

func (ic iothubCertificateAttributes) IothubName() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("iothub_name"))
}

func (ic iothubCertificateAttributes) IsVerified() terra.BoolValue {
	return terra.ReferenceBool(ic.ref.Append("is_verified"))
}

func (ic iothubCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("name"))
}

func (ic iothubCertificateAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ic.ref.Append("resource_group_name"))
}

func (ic iothubCertificateAttributes) Timeouts() iothubcertificate.TimeoutsAttributes {
	return terra.ReferenceSingle[iothubcertificate.TimeoutsAttributes](ic.ref.Append("timeouts"))
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
