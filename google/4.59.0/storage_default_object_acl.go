// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewStorageDefaultObjectAcl(name string, args StorageDefaultObjectAclArgs) *StorageDefaultObjectAcl {
	return &StorageDefaultObjectAcl{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageDefaultObjectAcl)(nil)

type StorageDefaultObjectAcl struct {
	Name  string
	Args  StorageDefaultObjectAclArgs
	state *storageDefaultObjectAclState
}

func (sdoa *StorageDefaultObjectAcl) Type() string {
	return "google_storage_default_object_acl"
}

func (sdoa *StorageDefaultObjectAcl) LocalName() string {
	return sdoa.Name
}

func (sdoa *StorageDefaultObjectAcl) Configuration() interface{} {
	return sdoa.Args
}

func (sdoa *StorageDefaultObjectAcl) Attributes() storageDefaultObjectAclAttributes {
	return storageDefaultObjectAclAttributes{ref: terra.ReferenceResource(sdoa)}
}

func (sdoa *StorageDefaultObjectAcl) ImportState(av io.Reader) error {
	sdoa.state = &storageDefaultObjectAclState{}
	if err := json.NewDecoder(av).Decode(sdoa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdoa.Type(), sdoa.LocalName(), err)
	}
	return nil
}

func (sdoa *StorageDefaultObjectAcl) State() (*storageDefaultObjectAclState, bool) {
	return sdoa.state, sdoa.state != nil
}

func (sdoa *StorageDefaultObjectAcl) StateMust() *storageDefaultObjectAclState {
	if sdoa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdoa.Type(), sdoa.LocalName()))
	}
	return sdoa.state
}

func (sdoa *StorageDefaultObjectAcl) DependOn() terra.Reference {
	return terra.ReferenceResource(sdoa)
}

type StorageDefaultObjectAclArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleEntity: set of string, optional
	RoleEntity terra.SetValue[terra.StringValue] `hcl:"role_entity,attr"`
	// DependsOn contains resources that StorageDefaultObjectAcl depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type storageDefaultObjectAclAttributes struct {
	ref terra.Reference
}

func (sdoa storageDefaultObjectAclAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(sdoa.ref.Append("bucket"))
}

func (sdoa storageDefaultObjectAclAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sdoa.ref.Append("id"))
}

func (sdoa storageDefaultObjectAclAttributes) RoleEntity() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](sdoa.ref.Append("role_entity"))
}

type storageDefaultObjectAclState struct {
	Bucket     string   `json:"bucket"`
	Id         string   `json:"id"`
	RoleEntity []string `json:"role_entity"`
}
