// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	netappkmsconfig "github.com/golingon/terraproviders/google/5.24.0/netappkmsconfig"
	"io"
)

// NewNetappKmsconfig creates a new instance of [NetappKmsconfig].
func NewNetappKmsconfig(name string, args NetappKmsconfigArgs) *NetappKmsconfig {
	return &NetappKmsconfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetappKmsconfig)(nil)

// NetappKmsconfig represents the Terraform resource google_netapp_kmsconfig.
type NetappKmsconfig struct {
	Name      string
	Args      NetappKmsconfigArgs
	state     *netappKmsconfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetappKmsconfig].
func (nk *NetappKmsconfig) Type() string {
	return "google_netapp_kmsconfig"
}

// LocalName returns the local name for [NetappKmsconfig].
func (nk *NetappKmsconfig) LocalName() string {
	return nk.Name
}

// Configuration returns the configuration (args) for [NetappKmsconfig].
func (nk *NetappKmsconfig) Configuration() interface{} {
	return nk.Args
}

// DependOn is used for other resources to depend on [NetappKmsconfig].
func (nk *NetappKmsconfig) DependOn() terra.Reference {
	return terra.ReferenceResource(nk)
}

// Dependencies returns the list of resources [NetappKmsconfig] depends_on.
func (nk *NetappKmsconfig) Dependencies() terra.Dependencies {
	return nk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetappKmsconfig].
func (nk *NetappKmsconfig) LifecycleManagement() *terra.Lifecycle {
	return nk.Lifecycle
}

// Attributes returns the attributes for [NetappKmsconfig].
func (nk *NetappKmsconfig) Attributes() netappKmsconfigAttributes {
	return netappKmsconfigAttributes{ref: terra.ReferenceResource(nk)}
}

// ImportState imports the given attribute values into [NetappKmsconfig]'s state.
func (nk *NetappKmsconfig) ImportState(av io.Reader) error {
	nk.state = &netappKmsconfigState{}
	if err := json.NewDecoder(av).Decode(nk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nk.Type(), nk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetappKmsconfig] has state.
func (nk *NetappKmsconfig) State() (*netappKmsconfigState, bool) {
	return nk.state, nk.state != nil
}

// StateMust returns the state for [NetappKmsconfig]. Panics if the state is nil.
func (nk *NetappKmsconfig) StateMust() *netappKmsconfigState {
	if nk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nk.Type(), nk.LocalName()))
	}
	return nk.state
}

// NetappKmsconfigArgs contains the configurations for google_netapp_kmsconfig.
type NetappKmsconfigArgs struct {
	// CryptoKeyName: string, required
	CryptoKeyName terra.StringValue `hcl:"crypto_key_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *netappkmsconfig.Timeouts `hcl:"timeouts,block"`
}
type netappKmsconfigAttributes struct {
	ref terra.Reference
}

// CryptoKeyName returns a reference to field crypto_key_name of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) CryptoKeyName() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("crypto_key_name"))
}

// Description returns a reference to field description of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nk.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("id"))
}

// Instructions returns a reference to field instructions of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Instructions() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("instructions"))
}

// Labels returns a reference to field labels of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nk.ref.Append("labels"))
}

// Location returns a reference to field location of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("location"))
}

// Name returns a reference to field name of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("name"))
}

// Project returns a reference to field project of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("project"))
}

// ServiceAccount returns a reference to field service_account of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(nk.ref.Append("service_account"))
}

// TerraformLabels returns a reference to field terraform_labels of google_netapp_kmsconfig.
func (nk netappKmsconfigAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nk.ref.Append("terraform_labels"))
}

func (nk netappKmsconfigAttributes) Timeouts() netappkmsconfig.TimeoutsAttributes {
	return terra.ReferenceAsSingle[netappkmsconfig.TimeoutsAttributes](nk.ref.Append("timeouts"))
}

type netappKmsconfigState struct {
	CryptoKeyName   string                         `json:"crypto_key_name"`
	Description     string                         `json:"description"`
	EffectiveLabels map[string]string              `json:"effective_labels"`
	Id              string                         `json:"id"`
	Instructions    string                         `json:"instructions"`
	Labels          map[string]string              `json:"labels"`
	Location        string                         `json:"location"`
	Name            string                         `json:"name"`
	Project         string                         `json:"project"`
	ServiceAccount  string                         `json:"service_account"`
	TerraformLabels map[string]string              `json:"terraform_labels"`
	Timeouts        *netappkmsconfig.TimeoutsState `json:"timeouts"`
}
