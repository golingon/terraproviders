// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeekeystoresaliaseskeycertfile "github.com/golingon/terraproviders/googlebeta/4.72.1/apigeekeystoresaliaseskeycertfile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeKeystoresAliasesKeyCertFile creates a new instance of [ApigeeKeystoresAliasesKeyCertFile].
func NewApigeeKeystoresAliasesKeyCertFile(name string, args ApigeeKeystoresAliasesKeyCertFileArgs) *ApigeeKeystoresAliasesKeyCertFile {
	return &ApigeeKeystoresAliasesKeyCertFile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeKeystoresAliasesKeyCertFile)(nil)

// ApigeeKeystoresAliasesKeyCertFile represents the Terraform resource google_apigee_keystores_aliases_key_cert_file.
type ApigeeKeystoresAliasesKeyCertFile struct {
	Name      string
	Args      ApigeeKeystoresAliasesKeyCertFileArgs
	state     *apigeeKeystoresAliasesKeyCertFileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeKeystoresAliasesKeyCertFile].
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) Type() string {
	return "google_apigee_keystores_aliases_key_cert_file"
}

// LocalName returns the local name for [ApigeeKeystoresAliasesKeyCertFile].
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) LocalName() string {
	return akakcf.Name
}

// Configuration returns the configuration (args) for [ApigeeKeystoresAliasesKeyCertFile].
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) Configuration() interface{} {
	return akakcf.Args
}

// DependOn is used for other resources to depend on [ApigeeKeystoresAliasesKeyCertFile].
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) DependOn() terra.Reference {
	return terra.ReferenceResource(akakcf)
}

// Dependencies returns the list of resources [ApigeeKeystoresAliasesKeyCertFile] depends_on.
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) Dependencies() terra.Dependencies {
	return akakcf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeKeystoresAliasesKeyCertFile].
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) LifecycleManagement() *terra.Lifecycle {
	return akakcf.Lifecycle
}

// Attributes returns the attributes for [ApigeeKeystoresAliasesKeyCertFile].
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) Attributes() apigeeKeystoresAliasesKeyCertFileAttributes {
	return apigeeKeystoresAliasesKeyCertFileAttributes{ref: terra.ReferenceResource(akakcf)}
}

// ImportState imports the given attribute values into [ApigeeKeystoresAliasesKeyCertFile]'s state.
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) ImportState(av io.Reader) error {
	akakcf.state = &apigeeKeystoresAliasesKeyCertFileState{}
	if err := json.NewDecoder(av).Decode(akakcf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", akakcf.Type(), akakcf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeKeystoresAliasesKeyCertFile] has state.
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) State() (*apigeeKeystoresAliasesKeyCertFileState, bool) {
	return akakcf.state, akakcf.state != nil
}

// StateMust returns the state for [ApigeeKeystoresAliasesKeyCertFile]. Panics if the state is nil.
func (akakcf *ApigeeKeystoresAliasesKeyCertFile) StateMust() *apigeeKeystoresAliasesKeyCertFileState {
	if akakcf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", akakcf.Type(), akakcf.LocalName()))
	}
	return akakcf.state
}

// ApigeeKeystoresAliasesKeyCertFileArgs contains the configurations for google_apigee_keystores_aliases_key_cert_file.
type ApigeeKeystoresAliasesKeyCertFileArgs struct {
	// Alias: string, required
	Alias terra.StringValue `hcl:"alias,attr" validate:"required"`
	// Cert: string, required
	Cert terra.StringValue `hcl:"cert,attr" validate:"required"`
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Keystore: string, required
	Keystore terra.StringValue `hcl:"keystore,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Password: string, optional
	Password terra.StringValue `hcl:"password,attr"`
	// CertsInfo: optional
	CertsInfo *apigeekeystoresaliaseskeycertfile.CertsInfo `hcl:"certs_info,block"`
	// Timeouts: optional
	Timeouts *apigeekeystoresaliaseskeycertfile.Timeouts `hcl:"timeouts,block"`
}
type apigeeKeystoresAliasesKeyCertFileAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("alias"))
}

// Cert returns a reference to field cert of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Cert() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("cert"))
}

// Environment returns a reference to field environment of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("environment"))
}

// Id returns a reference to field id of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("id"))
}

// Key returns a reference to field key of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("key"))
}

// Keystore returns a reference to field keystore of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Keystore() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("keystore"))
}

// OrgId returns a reference to field org_id of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("org_id"))
}

// Password returns a reference to field password of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("password"))
}

// Type returns a reference to field type of google_apigee_keystores_aliases_key_cert_file.
func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(akakcf.ref.Append("type"))
}

func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) CertsInfo() terra.ListValue[apigeekeystoresaliaseskeycertfile.CertsInfoAttributes] {
	return terra.ReferenceAsList[apigeekeystoresaliaseskeycertfile.CertsInfoAttributes](akakcf.ref.Append("certs_info"))
}

func (akakcf apigeeKeystoresAliasesKeyCertFileAttributes) Timeouts() apigeekeystoresaliaseskeycertfile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeekeystoresaliaseskeycertfile.TimeoutsAttributes](akakcf.ref.Append("timeouts"))
}

type apigeeKeystoresAliasesKeyCertFileState struct {
	Alias       string                                             `json:"alias"`
	Cert        string                                             `json:"cert"`
	Environment string                                             `json:"environment"`
	Id          string                                             `json:"id"`
	Key         string                                             `json:"key"`
	Keystore    string                                             `json:"keystore"`
	OrgId       string                                             `json:"org_id"`
	Password    string                                             `json:"password"`
	Type        string                                             `json:"type"`
	CertsInfo   []apigeekeystoresaliaseskeycertfile.CertsInfoState `json:"certs_info"`
	Timeouts    *apigeekeystoresaliaseskeycertfile.TimeoutsState   `json:"timeouts"`
}
