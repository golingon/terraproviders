// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	certificatemanagerdnsauthorization "github.com/golingon/terraproviders/googlebeta/4.64.0/certificatemanagerdnsauthorization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCertificateManagerDnsAuthorization creates a new instance of [CertificateManagerDnsAuthorization].
func NewCertificateManagerDnsAuthorization(name string, args CertificateManagerDnsAuthorizationArgs) *CertificateManagerDnsAuthorization {
	return &CertificateManagerDnsAuthorization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CertificateManagerDnsAuthorization)(nil)

// CertificateManagerDnsAuthorization represents the Terraform resource google_certificate_manager_dns_authorization.
type CertificateManagerDnsAuthorization struct {
	Name      string
	Args      CertificateManagerDnsAuthorizationArgs
	state     *certificateManagerDnsAuthorizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CertificateManagerDnsAuthorization].
func (cmda *CertificateManagerDnsAuthorization) Type() string {
	return "google_certificate_manager_dns_authorization"
}

// LocalName returns the local name for [CertificateManagerDnsAuthorization].
func (cmda *CertificateManagerDnsAuthorization) LocalName() string {
	return cmda.Name
}

// Configuration returns the configuration (args) for [CertificateManagerDnsAuthorization].
func (cmda *CertificateManagerDnsAuthorization) Configuration() interface{} {
	return cmda.Args
}

// DependOn is used for other resources to depend on [CertificateManagerDnsAuthorization].
func (cmda *CertificateManagerDnsAuthorization) DependOn() terra.Reference {
	return terra.ReferenceResource(cmda)
}

// Dependencies returns the list of resources [CertificateManagerDnsAuthorization] depends_on.
func (cmda *CertificateManagerDnsAuthorization) Dependencies() terra.Dependencies {
	return cmda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CertificateManagerDnsAuthorization].
func (cmda *CertificateManagerDnsAuthorization) LifecycleManagement() *terra.Lifecycle {
	return cmda.Lifecycle
}

// Attributes returns the attributes for [CertificateManagerDnsAuthorization].
func (cmda *CertificateManagerDnsAuthorization) Attributes() certificateManagerDnsAuthorizationAttributes {
	return certificateManagerDnsAuthorizationAttributes{ref: terra.ReferenceResource(cmda)}
}

// ImportState imports the given attribute values into [CertificateManagerDnsAuthorization]'s state.
func (cmda *CertificateManagerDnsAuthorization) ImportState(av io.Reader) error {
	cmda.state = &certificateManagerDnsAuthorizationState{}
	if err := json.NewDecoder(av).Decode(cmda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cmda.Type(), cmda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CertificateManagerDnsAuthorization] has state.
func (cmda *CertificateManagerDnsAuthorization) State() (*certificateManagerDnsAuthorizationState, bool) {
	return cmda.state, cmda.state != nil
}

// StateMust returns the state for [CertificateManagerDnsAuthorization]. Panics if the state is nil.
func (cmda *CertificateManagerDnsAuthorization) StateMust() *certificateManagerDnsAuthorizationState {
	if cmda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cmda.Type(), cmda.LocalName()))
	}
	return cmda.state
}

// CertificateManagerDnsAuthorizationArgs contains the configurations for google_certificate_manager_dns_authorization.
type CertificateManagerDnsAuthorizationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Domain: string, required
	Domain terra.StringValue `hcl:"domain,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// DnsResourceRecord: min=0
	DnsResourceRecord []certificatemanagerdnsauthorization.DnsResourceRecord `hcl:"dns_resource_record,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *certificatemanagerdnsauthorization.Timeouts `hcl:"timeouts,block"`
}
type certificateManagerDnsAuthorizationAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_certificate_manager_dns_authorization.
func (cmda certificateManagerDnsAuthorizationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cmda.ref.Append("description"))
}

// Domain returns a reference to field domain of google_certificate_manager_dns_authorization.
func (cmda certificateManagerDnsAuthorizationAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(cmda.ref.Append("domain"))
}

// Id returns a reference to field id of google_certificate_manager_dns_authorization.
func (cmda certificateManagerDnsAuthorizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cmda.ref.Append("id"))
}

// Labels returns a reference to field labels of google_certificate_manager_dns_authorization.
func (cmda certificateManagerDnsAuthorizationAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cmda.ref.Append("labels"))
}

// Name returns a reference to field name of google_certificate_manager_dns_authorization.
func (cmda certificateManagerDnsAuthorizationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cmda.ref.Append("name"))
}

// Project returns a reference to field project of google_certificate_manager_dns_authorization.
func (cmda certificateManagerDnsAuthorizationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cmda.ref.Append("project"))
}

func (cmda certificateManagerDnsAuthorizationAttributes) DnsResourceRecord() terra.ListValue[certificatemanagerdnsauthorization.DnsResourceRecordAttributes] {
	return terra.ReferenceAsList[certificatemanagerdnsauthorization.DnsResourceRecordAttributes](cmda.ref.Append("dns_resource_record"))
}

func (cmda certificateManagerDnsAuthorizationAttributes) Timeouts() certificatemanagerdnsauthorization.TimeoutsAttributes {
	return terra.ReferenceAsSingle[certificatemanagerdnsauthorization.TimeoutsAttributes](cmda.ref.Append("timeouts"))
}

type certificateManagerDnsAuthorizationState struct {
	Description       string                                                      `json:"description"`
	Domain            string                                                      `json:"domain"`
	Id                string                                                      `json:"id"`
	Labels            map[string]string                                           `json:"labels"`
	Name              string                                                      `json:"name"`
	Project           string                                                      `json:"project"`
	DnsResourceRecord []certificatemanagerdnsauthorization.DnsResourceRecordState `json:"dns_resource_record"`
	Timeouts          *certificatemanagerdnsauthorization.TimeoutsState           `json:"timeouts"`
}
