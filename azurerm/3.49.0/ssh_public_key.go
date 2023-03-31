// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sshpublickey "github.com/golingon/terraproviders/azurerm/3.49.0/sshpublickey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSshPublicKey(name string, args SshPublicKeyArgs) *SshPublicKey {
	return &SshPublicKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SshPublicKey)(nil)

type SshPublicKey struct {
	Name  string
	Args  SshPublicKeyArgs
	state *sshPublicKeyState
}

func (spk *SshPublicKey) Type() string {
	return "azurerm_ssh_public_key"
}

func (spk *SshPublicKey) LocalName() string {
	return spk.Name
}

func (spk *SshPublicKey) Configuration() interface{} {
	return spk.Args
}

func (spk *SshPublicKey) Attributes() sshPublicKeyAttributes {
	return sshPublicKeyAttributes{ref: terra.ReferenceResource(spk)}
}

func (spk *SshPublicKey) ImportState(av io.Reader) error {
	spk.state = &sshPublicKeyState{}
	if err := json.NewDecoder(av).Decode(spk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", spk.Type(), spk.LocalName(), err)
	}
	return nil
}

func (spk *SshPublicKey) State() (*sshPublicKeyState, bool) {
	return spk.state, spk.state != nil
}

func (spk *SshPublicKey) StateMust() *sshPublicKeyState {
	if spk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", spk.Type(), spk.LocalName()))
	}
	return spk.state
}

func (spk *SshPublicKey) DependOn() terra.Reference {
	return terra.ReferenceResource(spk)
}

type SshPublicKeyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicKey: string, required
	PublicKey terra.StringValue `hcl:"public_key,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *sshpublickey.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that SshPublicKey depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sshPublicKeyAttributes struct {
	ref terra.Reference
}

func (spk sshPublicKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(spk.ref.Append("id"))
}

func (spk sshPublicKeyAttributes) Location() terra.StringValue {
	return terra.ReferenceString(spk.ref.Append("location"))
}

func (spk sshPublicKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(spk.ref.Append("name"))
}

func (spk sshPublicKeyAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceString(spk.ref.Append("public_key"))
}

func (spk sshPublicKeyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(spk.ref.Append("resource_group_name"))
}

func (spk sshPublicKeyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](spk.ref.Append("tags"))
}

func (spk sshPublicKeyAttributes) Timeouts() sshpublickey.TimeoutsAttributes {
	return terra.ReferenceSingle[sshpublickey.TimeoutsAttributes](spk.ref.Append("timeouts"))
}

type sshPublicKeyState struct {
	Id                string                      `json:"id"`
	Location          string                      `json:"location"`
	Name              string                      `json:"name"`
	PublicKey         string                      `json:"public_key"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Timeouts          *sshpublickey.TimeoutsState `json:"timeouts"`
}
