// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewStoragegatewayUploadBuffer(name string, args StoragegatewayUploadBufferArgs) *StoragegatewayUploadBuffer {
	return &StoragegatewayUploadBuffer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StoragegatewayUploadBuffer)(nil)

type StoragegatewayUploadBuffer struct {
	Name  string
	Args  StoragegatewayUploadBufferArgs
	state *storagegatewayUploadBufferState
}

func (sub *StoragegatewayUploadBuffer) Type() string {
	return "aws_storagegateway_upload_buffer"
}

func (sub *StoragegatewayUploadBuffer) LocalName() string {
	return sub.Name
}

func (sub *StoragegatewayUploadBuffer) Configuration() interface{} {
	return sub.Args
}

func (sub *StoragegatewayUploadBuffer) Attributes() storagegatewayUploadBufferAttributes {
	return storagegatewayUploadBufferAttributes{ref: terra.ReferenceResource(sub)}
}

func (sub *StoragegatewayUploadBuffer) ImportState(av io.Reader) error {
	sub.state = &storagegatewayUploadBufferState{}
	if err := json.NewDecoder(av).Decode(sub.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sub.Type(), sub.LocalName(), err)
	}
	return nil
}

func (sub *StoragegatewayUploadBuffer) State() (*storagegatewayUploadBufferState, bool) {
	return sub.state, sub.state != nil
}

func (sub *StoragegatewayUploadBuffer) StateMust() *storagegatewayUploadBufferState {
	if sub.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sub.Type(), sub.LocalName()))
	}
	return sub.state
}

func (sub *StoragegatewayUploadBuffer) DependOn() terra.Reference {
	return terra.ReferenceResource(sub)
}

type StoragegatewayUploadBufferArgs struct {
	// DiskId: string, optional
	DiskId terra.StringValue `hcl:"disk_id,attr"`
	// DiskPath: string, optional
	DiskPath terra.StringValue `hcl:"disk_path,attr"`
	// GatewayArn: string, required
	GatewayArn terra.StringValue `hcl:"gateway_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// DependsOn contains resources that StoragegatewayUploadBuffer depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type storagegatewayUploadBufferAttributes struct {
	ref terra.Reference
}

func (sub storagegatewayUploadBufferAttributes) DiskId() terra.StringValue {
	return terra.ReferenceString(sub.ref.Append("disk_id"))
}

func (sub storagegatewayUploadBufferAttributes) DiskPath() terra.StringValue {
	return terra.ReferenceString(sub.ref.Append("disk_path"))
}

func (sub storagegatewayUploadBufferAttributes) GatewayArn() terra.StringValue {
	return terra.ReferenceString(sub.ref.Append("gateway_arn"))
}

func (sub storagegatewayUploadBufferAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sub.ref.Append("id"))
}

type storagegatewayUploadBufferState struct {
	DiskId     string `json:"disk_id"`
	DiskPath   string `json:"disk_path"`
	GatewayArn string `json:"gateway_arn"`
	Id         string `json:"id"`
}
