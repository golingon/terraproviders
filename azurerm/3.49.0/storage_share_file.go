// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagesharefile "github.com/golingon/terraproviders/azurerm/3.49.0/storagesharefile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewStorageShareFile(name string, args StorageShareFileArgs) *StorageShareFile {
	return &StorageShareFile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageShareFile)(nil)

type StorageShareFile struct {
	Name  string
	Args  StorageShareFileArgs
	state *storageShareFileState
}

func (ssf *StorageShareFile) Type() string {
	return "azurerm_storage_share_file"
}

func (ssf *StorageShareFile) LocalName() string {
	return ssf.Name
}

func (ssf *StorageShareFile) Configuration() interface{} {
	return ssf.Args
}

func (ssf *StorageShareFile) Attributes() storageShareFileAttributes {
	return storageShareFileAttributes{ref: terra.ReferenceResource(ssf)}
}

func (ssf *StorageShareFile) ImportState(av io.Reader) error {
	ssf.state = &storageShareFileState{}
	if err := json.NewDecoder(av).Decode(ssf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssf.Type(), ssf.LocalName(), err)
	}
	return nil
}

func (ssf *StorageShareFile) State() (*storageShareFileState, bool) {
	return ssf.state, ssf.state != nil
}

func (ssf *StorageShareFile) StateMust() *storageShareFileState {
	if ssf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssf.Type(), ssf.LocalName()))
	}
	return ssf.state
}

func (ssf *StorageShareFile) DependOn() terra.Reference {
	return terra.ReferenceResource(ssf)
}

type StorageShareFileArgs struct {
	// ContentDisposition: string, optional
	ContentDisposition terra.StringValue `hcl:"content_disposition,attr"`
	// ContentEncoding: string, optional
	ContentEncoding terra.StringValue `hcl:"content_encoding,attr"`
	// ContentMd5: string, optional
	ContentMd5 terra.StringValue `hcl:"content_md5,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
	// StorageShareId: string, required
	StorageShareId terra.StringValue `hcl:"storage_share_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagesharefile.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that StorageShareFile depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type storageShareFileAttributes struct {
	ref terra.Reference
}

func (ssf storageShareFileAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("content_disposition"))
}

func (ssf storageShareFileAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("content_encoding"))
}

func (ssf storageShareFileAttributes) ContentLength() terra.NumberValue {
	return terra.ReferenceNumber(ssf.ref.Append("content_length"))
}

func (ssf storageShareFileAttributes) ContentMd5() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("content_md5"))
}

func (ssf storageShareFileAttributes) ContentType() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("content_type"))
}

func (ssf storageShareFileAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("id"))
}

func (ssf storageShareFileAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ssf.ref.Append("metadata"))
}

func (ssf storageShareFileAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("name"))
}

func (ssf storageShareFileAttributes) Path() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("path"))
}

func (ssf storageShareFileAttributes) Source() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("source"))
}

func (ssf storageShareFileAttributes) StorageShareId() terra.StringValue {
	return terra.ReferenceString(ssf.ref.Append("storage_share_id"))
}

func (ssf storageShareFileAttributes) Timeouts() storagesharefile.TimeoutsAttributes {
	return terra.ReferenceSingle[storagesharefile.TimeoutsAttributes](ssf.ref.Append("timeouts"))
}

type storageShareFileState struct {
	ContentDisposition string                          `json:"content_disposition"`
	ContentEncoding    string                          `json:"content_encoding"`
	ContentLength      float64                         `json:"content_length"`
	ContentMd5         string                          `json:"content_md5"`
	ContentType        string                          `json:"content_type"`
	Id                 string                          `json:"id"`
	Metadata           map[string]string               `json:"metadata"`
	Name               string                          `json:"name"`
	Path               string                          `json:"path"`
	Source             string                          `json:"source"`
	StorageShareId     string                          `json:"storage_share_id"`
	Timeouts           *storagesharefile.TimeoutsState `json:"timeouts"`
}
