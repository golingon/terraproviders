// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewServiceAccountKey(name string, args ServiceAccountKeyArgs) *ServiceAccountKey {
	return &ServiceAccountKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceAccountKey)(nil)

type ServiceAccountKey struct {
	Name  string
	Args  ServiceAccountKeyArgs
	state *serviceAccountKeyState
}

func (sak *ServiceAccountKey) Type() string {
	return "google_service_account_key"
}

func (sak *ServiceAccountKey) LocalName() string {
	return sak.Name
}

func (sak *ServiceAccountKey) Configuration() interface{} {
	return sak.Args
}

func (sak *ServiceAccountKey) Attributes() serviceAccountKeyAttributes {
	return serviceAccountKeyAttributes{ref: terra.ReferenceResource(sak)}
}

func (sak *ServiceAccountKey) ImportState(av io.Reader) error {
	sak.state = &serviceAccountKeyState{}
	if err := json.NewDecoder(av).Decode(sak.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sak.Type(), sak.LocalName(), err)
	}
	return nil
}

func (sak *ServiceAccountKey) State() (*serviceAccountKeyState, bool) {
	return sak.state, sak.state != nil
}

func (sak *ServiceAccountKey) StateMust() *serviceAccountKeyState {
	if sak.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sak.Type(), sak.LocalName()))
	}
	return sak.state
}

func (sak *ServiceAccountKey) DependOn() terra.Reference {
	return terra.ReferenceResource(sak)
}

type ServiceAccountKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Keepers: map of string, optional
	Keepers terra.MapValue[terra.StringValue] `hcl:"keepers,attr"`
	// KeyAlgorithm: string, optional
	KeyAlgorithm terra.StringValue `hcl:"key_algorithm,attr"`
	// PrivateKeyType: string, optional
	PrivateKeyType terra.StringValue `hcl:"private_key_type,attr"`
	// PublicKeyData: string, optional
	PublicKeyData terra.StringValue `hcl:"public_key_data,attr"`
	// PublicKeyType: string, optional
	PublicKeyType terra.StringValue `hcl:"public_key_type,attr"`
	// ServiceAccountId: string, required
	ServiceAccountId terra.StringValue `hcl:"service_account_id,attr" validate:"required"`
	// DependsOn contains resources that ServiceAccountKey depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type serviceAccountKeyAttributes struct {
	ref terra.Reference
}

func (sak serviceAccountKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("id"))
}

func (sak serviceAccountKeyAttributes) Keepers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sak.ref.Append("keepers"))
}

func (sak serviceAccountKeyAttributes) KeyAlgorithm() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("key_algorithm"))
}

func (sak serviceAccountKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("name"))
}

func (sak serviceAccountKeyAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("private_key"))
}

func (sak serviceAccountKeyAttributes) PrivateKeyType() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("private_key_type"))
}

func (sak serviceAccountKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("public_key"))
}

func (sak serviceAccountKeyAttributes) PublicKeyData() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("public_key_data"))
}

func (sak serviceAccountKeyAttributes) PublicKeyType() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("public_key_type"))
}

func (sak serviceAccountKeyAttributes) ServiceAccountId() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("service_account_id"))
}

func (sak serviceAccountKeyAttributes) ValidAfter() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("valid_after"))
}

func (sak serviceAccountKeyAttributes) ValidBefore() terra.StringValue {
	return terra.ReferenceString(sak.ref.Append("valid_before"))
}

type serviceAccountKeyState struct {
	Id               string            `json:"id"`
	Keepers          map[string]string `json:"keepers"`
	KeyAlgorithm     string            `json:"key_algorithm"`
	Name             string            `json:"name"`
	PrivateKey       string            `json:"private_key"`
	PrivateKeyType   string            `json:"private_key_type"`
	PublicKey        string            `json:"public_key"`
	PublicKeyData    string            `json:"public_key_data"`
	PublicKeyType    string            `json:"public_key_type"`
	ServiceAccountId string            `json:"service_account_id"`
	ValidAfter       string            `json:"valid_after"`
	ValidBefore      string            `json:"valid_before"`
}
