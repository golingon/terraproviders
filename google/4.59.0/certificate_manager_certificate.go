// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	certificatemanagercertificate "github.com/golingon/terraproviders/google/4.59.0/certificatemanagercertificate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCertificateManagerCertificate(name string, args CertificateManagerCertificateArgs) *CertificateManagerCertificate {
	return &CertificateManagerCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertificateManagerCertificate)(nil)

type CertificateManagerCertificate struct {
	Name  string
	Args  CertificateManagerCertificateArgs
	state *certificateManagerCertificateState
}

func (cmc *CertificateManagerCertificate) Type() string {
	return "google_certificate_manager_certificate"
}

func (cmc *CertificateManagerCertificate) LocalName() string {
	return cmc.Name
}

func (cmc *CertificateManagerCertificate) Configuration() interface{} {
	return cmc.Args
}

func (cmc *CertificateManagerCertificate) Attributes() certificateManagerCertificateAttributes {
	return certificateManagerCertificateAttributes{ref: terra.ReferenceResource(cmc)}
}

func (cmc *CertificateManagerCertificate) ImportState(av io.Reader) error {
	cmc.state = &certificateManagerCertificateState{}
	if err := json.NewDecoder(av).Decode(cmc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmc.Type(), cmc.LocalName(), err)
	}
	return nil
}

func (cmc *CertificateManagerCertificate) State() (*certificateManagerCertificateState, bool) {
	return cmc.state, cmc.state != nil
}

func (cmc *CertificateManagerCertificate) StateMust() *certificateManagerCertificateState {
	if cmc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmc.Type(), cmc.LocalName()))
	}
	return cmc.state
}

func (cmc *CertificateManagerCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(cmc)
}

type CertificateManagerCertificateArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// Managed: optional
	Managed *certificatemanagercertificate.Managed `hcl:"managed,block"`
	// SelfManaged: optional
	SelfManaged *certificatemanagercertificate.SelfManaged `hcl:"self_managed,block"`
	// Timeouts: optional
	Timeouts *certificatemanagercertificate.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that CertificateManagerCertificate depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type certificateManagerCertificateAttributes struct {
	ref terra.Reference
}

func (cmc certificateManagerCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("description"))
}

func (cmc certificateManagerCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("id"))
}

func (cmc certificateManagerCertificateAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cmc.ref.Append("labels"))
}

func (cmc certificateManagerCertificateAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("name"))
}

func (cmc certificateManagerCertificateAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("project"))
}

func (cmc certificateManagerCertificateAttributes) Scope() terra.StringValue {
	return terra.ReferenceString(cmc.ref.Append("scope"))
}

func (cmc certificateManagerCertificateAttributes) Managed() terra.ListValue[certificatemanagercertificate.ManagedAttributes] {
	return terra.ReferenceList[certificatemanagercertificate.ManagedAttributes](cmc.ref.Append("managed"))
}

func (cmc certificateManagerCertificateAttributes) SelfManaged() terra.ListValue[certificatemanagercertificate.SelfManagedAttributes] {
	return terra.ReferenceList[certificatemanagercertificate.SelfManagedAttributes](cmc.ref.Append("self_managed"))
}

func (cmc certificateManagerCertificateAttributes) Timeouts() certificatemanagercertificate.TimeoutsAttributes {
	return terra.ReferenceSingle[certificatemanagercertificate.TimeoutsAttributes](cmc.ref.Append("timeouts"))
}

type certificateManagerCertificateState struct {
	Description string                                           `json:"description"`
	Id          string                                           `json:"id"`
	Labels      map[string]string                                `json:"labels"`
	Name        string                                           `json:"name"`
	Project     string                                           `json:"project"`
	Scope       string                                           `json:"scope"`
	Managed     []certificatemanagercertificate.ManagedState     `json:"managed"`
	SelfManaged []certificatemanagercertificate.SelfManagedState `json:"self_managed"`
	Timeouts    *certificatemanagercertificate.TimeoutsState     `json:"timeouts"`
}
