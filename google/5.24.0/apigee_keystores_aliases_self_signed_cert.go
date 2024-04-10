// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	apigeekeystoresaliasesselfsignedcert "github.com/golingon/terraproviders/google/5.24.0/apigeekeystoresaliasesselfsignedcert"
	"io"
)

// NewApigeeKeystoresAliasesSelfSignedCert creates a new instance of [ApigeeKeystoresAliasesSelfSignedCert].
func NewApigeeKeystoresAliasesSelfSignedCert(name string, args ApigeeKeystoresAliasesSelfSignedCertArgs) *ApigeeKeystoresAliasesSelfSignedCert {
	return &ApigeeKeystoresAliasesSelfSignedCert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeKeystoresAliasesSelfSignedCert)(nil)

// ApigeeKeystoresAliasesSelfSignedCert represents the Terraform resource google_apigee_keystores_aliases_self_signed_cert.
type ApigeeKeystoresAliasesSelfSignedCert struct {
	Name      string
	Args      ApigeeKeystoresAliasesSelfSignedCertArgs
	state     *apigeeKeystoresAliasesSelfSignedCertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeKeystoresAliasesSelfSignedCert].
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) Type() string {
	return "google_apigee_keystores_aliases_self_signed_cert"
}

// LocalName returns the local name for [ApigeeKeystoresAliasesSelfSignedCert].
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) LocalName() string {
	return akassc.Name
}

// Configuration returns the configuration (args) for [ApigeeKeystoresAliasesSelfSignedCert].
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) Configuration() interface{} {
	return akassc.Args
}

// DependOn is used for other resources to depend on [ApigeeKeystoresAliasesSelfSignedCert].
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) DependOn() terra.Reference {
	return terra.ReferenceResource(akassc)
}

// Dependencies returns the list of resources [ApigeeKeystoresAliasesSelfSignedCert] depends_on.
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) Dependencies() terra.Dependencies {
	return akassc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeKeystoresAliasesSelfSignedCert].
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) LifecycleManagement() *terra.Lifecycle {
	return akassc.Lifecycle
}

// Attributes returns the attributes for [ApigeeKeystoresAliasesSelfSignedCert].
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) Attributes() apigeeKeystoresAliasesSelfSignedCertAttributes {
	return apigeeKeystoresAliasesSelfSignedCertAttributes{ref: terra.ReferenceResource(akassc)}
}

// ImportState imports the given attribute values into [ApigeeKeystoresAliasesSelfSignedCert]'s state.
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) ImportState(av io.Reader) error {
	akassc.state = &apigeeKeystoresAliasesSelfSignedCertState{}
	if err := json.NewDecoder(av).Decode(akassc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akassc.Type(), akassc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeKeystoresAliasesSelfSignedCert] has state.
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) State() (*apigeeKeystoresAliasesSelfSignedCertState, bool) {
	return akassc.state, akassc.state != nil
}

// StateMust returns the state for [ApigeeKeystoresAliasesSelfSignedCert]. Panics if the state is nil.
func (akassc *ApigeeKeystoresAliasesSelfSignedCert) StateMust() *apigeeKeystoresAliasesSelfSignedCertState {
	if akassc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akassc.Type(), akassc.LocalName()))
	}
	return akassc.state
}

// ApigeeKeystoresAliasesSelfSignedCertArgs contains the configurations for google_apigee_keystores_aliases_self_signed_cert.
type ApigeeKeystoresAliasesSelfSignedCertArgs struct {
	// Alias: string, required
	Alias terra.StringValue `hcl:"alias,attr" validate:"required"`
	// CertValidityInDays: number, optional
	CertValidityInDays terra.NumberValue `hcl:"cert_validity_in_days,attr"`
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeySize: string, optional
	KeySize terra.StringValue `hcl:"key_size,attr"`
	// Keystore: string, required
	Keystore terra.StringValue `hcl:"keystore,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// SigAlg: string, required
	SigAlg terra.StringValue `hcl:"sig_alg,attr" validate:"required"`
	// CertsInfo: min=0
	CertsInfo []apigeekeystoresaliasesselfsignedcert.CertsInfo `hcl:"certs_info,block" validate:"min=0"`
	// Subject: required
	Subject *apigeekeystoresaliasesselfsignedcert.Subject `hcl:"subject,block" validate:"required"`
	// SubjectAlternativeDnsNames: optional
	SubjectAlternativeDnsNames *apigeekeystoresaliasesselfsignedcert.SubjectAlternativeDnsNames `hcl:"subject_alternative_dns_names,block"`
	// Timeouts: optional
	Timeouts *apigeekeystoresaliasesselfsignedcert.Timeouts `hcl:"timeouts,block"`
}
type apigeeKeystoresAliasesSelfSignedCertAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("alias"))
}

// CertValidityInDays returns a reference to field cert_validity_in_days of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) CertValidityInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(akassc.ref.Append("cert_validity_in_days"))
}

// Environment returns a reference to field environment of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("environment"))
}

// Id returns a reference to field id of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("id"))
}

// KeySize returns a reference to field key_size of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) KeySize() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("key_size"))
}

// Keystore returns a reference to field keystore of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Keystore() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("keystore"))
}

// OrgId returns a reference to field org_id of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("org_id"))
}

// SigAlg returns a reference to field sig_alg of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) SigAlg() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("sig_alg"))
}

// Type returns a reference to field type of google_apigee_keystores_aliases_self_signed_cert.
func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(akassc.ref.Append("type"))
}

func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) CertsInfo() terra.ListValue[apigeekeystoresaliasesselfsignedcert.CertsInfoAttributes] {
	return terra.ReferenceAsList[apigeekeystoresaliasesselfsignedcert.CertsInfoAttributes](akassc.ref.Append("certs_info"))
}

func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Subject() terra.ListValue[apigeekeystoresaliasesselfsignedcert.SubjectAttributes] {
	return terra.ReferenceAsList[apigeekeystoresaliasesselfsignedcert.SubjectAttributes](akassc.ref.Append("subject"))
}

func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) SubjectAlternativeDnsNames() terra.ListValue[apigeekeystoresaliasesselfsignedcert.SubjectAlternativeDnsNamesAttributes] {
	return terra.ReferenceAsList[apigeekeystoresaliasesselfsignedcert.SubjectAlternativeDnsNamesAttributes](akassc.ref.Append("subject_alternative_dns_names"))
}

func (akassc apigeeKeystoresAliasesSelfSignedCertAttributes) Timeouts() apigeekeystoresaliasesselfsignedcert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeekeystoresaliasesselfsignedcert.TimeoutsAttributes](akassc.ref.Append("timeouts"))
}

type apigeeKeystoresAliasesSelfSignedCertState struct {
	Alias                      string                                                                 `json:"alias"`
	CertValidityInDays         float64                                                                `json:"cert_validity_in_days"`
	Environment                string                                                                 `json:"environment"`
	Id                         string                                                                 `json:"id"`
	KeySize                    string                                                                 `json:"key_size"`
	Keystore                   string                                                                 `json:"keystore"`
	OrgId                      string                                                                 `json:"org_id"`
	SigAlg                     string                                                                 `json:"sig_alg"`
	Type                       string                                                                 `json:"type"`
	CertsInfo                  []apigeekeystoresaliasesselfsignedcert.CertsInfoState                  `json:"certs_info"`
	Subject                    []apigeekeystoresaliasesselfsignedcert.SubjectState                    `json:"subject"`
	SubjectAlternativeDnsNames []apigeekeystoresaliasesselfsignedcert.SubjectAlternativeDnsNamesState `json:"subject_alternative_dns_names"`
	Timeouts                   *apigeekeystoresaliasesselfsignedcert.TimeoutsState                    `json:"timeouts"`
}
