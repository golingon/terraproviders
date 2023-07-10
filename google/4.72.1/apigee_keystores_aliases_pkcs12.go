// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeekeystoresaliasespkcs12 "github.com/golingon/terraproviders/google/4.72.1/apigeekeystoresaliasespkcs12"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeKeystoresAliasesPkcs12 creates a new instance of [ApigeeKeystoresAliasesPkcs12].
func NewApigeeKeystoresAliasesPkcs12(name string, args ApigeeKeystoresAliasesPkcs12Args) *ApigeeKeystoresAliasesPkcs12 {
	return &ApigeeKeystoresAliasesPkcs12{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeKeystoresAliasesPkcs12)(nil)

// ApigeeKeystoresAliasesPkcs12 represents the Terraform resource google_apigee_keystores_aliases_pkcs12.
type ApigeeKeystoresAliasesPkcs12 struct {
	Name      string
	Args      ApigeeKeystoresAliasesPkcs12Args
	state     *apigeeKeystoresAliasesPkcs12State
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeKeystoresAliasesPkcs12].
func (akap *ApigeeKeystoresAliasesPkcs12) Type() string {
	return "google_apigee_keystores_aliases_pkcs12"
}

// LocalName returns the local name for [ApigeeKeystoresAliasesPkcs12].
func (akap *ApigeeKeystoresAliasesPkcs12) LocalName() string {
	return akap.Name
}

// Configuration returns the configuration (args) for [ApigeeKeystoresAliasesPkcs12].
func (akap *ApigeeKeystoresAliasesPkcs12) Configuration() interface{} {
	return akap.Args
}

// DependOn is used for other resources to depend on [ApigeeKeystoresAliasesPkcs12].
func (akap *ApigeeKeystoresAliasesPkcs12) DependOn() terra.Reference {
	return terra.ReferenceResource(akap)
}

// Dependencies returns the list of resources [ApigeeKeystoresAliasesPkcs12] depends_on.
func (akap *ApigeeKeystoresAliasesPkcs12) Dependencies() terra.Dependencies {
	return akap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeKeystoresAliasesPkcs12].
func (akap *ApigeeKeystoresAliasesPkcs12) LifecycleManagement() *terra.Lifecycle {
	return akap.Lifecycle
}

// Attributes returns the attributes for [ApigeeKeystoresAliasesPkcs12].
func (akap *ApigeeKeystoresAliasesPkcs12) Attributes() apigeeKeystoresAliasesPkcs12Attributes {
	return apigeeKeystoresAliasesPkcs12Attributes{ref: terra.ReferenceResource(akap)}
}

// ImportState imports the given attribute values into [ApigeeKeystoresAliasesPkcs12]'s state.
func (akap *ApigeeKeystoresAliasesPkcs12) ImportState(av io.Reader) error {
	akap.state = &apigeeKeystoresAliasesPkcs12State{}
	if err := json.NewDecoder(av).Decode(akap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akap.Type(), akap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeKeystoresAliasesPkcs12] has state.
func (akap *ApigeeKeystoresAliasesPkcs12) State() (*apigeeKeystoresAliasesPkcs12State, bool) {
	return akap.state, akap.state != nil
}

// StateMust returns the state for [ApigeeKeystoresAliasesPkcs12]. Panics if the state is nil.
func (akap *ApigeeKeystoresAliasesPkcs12) StateMust() *apigeeKeystoresAliasesPkcs12State {
	if akap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akap.Type(), akap.LocalName()))
	}
	return akap.state
}

// ApigeeKeystoresAliasesPkcs12Args contains the configurations for google_apigee_keystores_aliases_pkcs12.
type ApigeeKeystoresAliasesPkcs12Args struct {
	// Alias: string, required
	Alias terra.StringValue `hcl:"alias,attr" validate:"required"`
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// File: string, required
	File terra.StringValue `hcl:"file,attr" validate:"required"`
	// Filehash: string, required
	Filehash terra.StringValue `hcl:"filehash,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Keystore: string, required
	Keystore terra.StringValue `hcl:"keystore,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// CertsInfo: min=0
	CertsInfo []apigeekeystoresaliasespkcs12.CertsInfo `hcl:"certs_info,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *apigeekeystoresaliasespkcs12.Timeouts `hcl:"timeouts,block"`
}
type apigeeKeystoresAliasesPkcs12Attributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("alias"))
}

// Environment returns a reference to field environment of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("environment"))
}

// File returns a reference to field file of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) File() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("file"))
}

// Filehash returns a reference to field filehash of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Filehash() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("filehash"))
}

// Id returns a reference to field id of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("id"))
}

// Keystore returns a reference to field keystore of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Keystore() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("keystore"))
}

// OrgId returns a reference to field org_id of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("org_id"))
}

// Password returns a reference to field password of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Password() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("password"))
}

// Type returns a reference to field type of google_apigee_keystores_aliases_pkcs12.
func (akap apigeeKeystoresAliasesPkcs12Attributes) Type() terra.StringValue {
	return terra.ReferenceAsString(akap.ref.Append("type"))
}

func (akap apigeeKeystoresAliasesPkcs12Attributes) CertsInfo() terra.ListValue[apigeekeystoresaliasespkcs12.CertsInfoAttributes] {
	return terra.ReferenceAsList[apigeekeystoresaliasespkcs12.CertsInfoAttributes](akap.ref.Append("certs_info"))
}

func (akap apigeeKeystoresAliasesPkcs12Attributes) Timeouts() apigeekeystoresaliasespkcs12.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeekeystoresaliasespkcs12.TimeoutsAttributes](akap.ref.Append("timeouts"))
}

type apigeeKeystoresAliasesPkcs12State struct {
	Alias       string                                        `json:"alias"`
	Environment string                                        `json:"environment"`
	File        string                                        `json:"file"`
	Filehash    string                                        `json:"filehash"`
	Id          string                                        `json:"id"`
	Keystore    string                                        `json:"keystore"`
	OrgId       string                                        `json:"org_id"`
	Password    string                                        `json:"password"`
	Type        string                                        `json:"type"`
	CertsInfo   []apigeekeystoresaliasespkcs12.CertsInfoState `json:"certs_info"`
	Timeouts    *apigeekeystoresaliasespkcs12.TimeoutsState   `json:"timeouts"`
}
