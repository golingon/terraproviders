// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	efsaccesspoint "github.com/golingon/terraproviders/aws/5.0.1/efsaccesspoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEfsAccessPoint creates a new instance of [EfsAccessPoint].
func NewEfsAccessPoint(name string, args EfsAccessPointArgs) *EfsAccessPoint {
	return &EfsAccessPoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsAccessPoint)(nil)

// EfsAccessPoint represents the Terraform resource aws_efs_access_point.
type EfsAccessPoint struct {
	Name      string
	Args      EfsAccessPointArgs
	state     *efsAccessPointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EfsAccessPoint].
func (eap *EfsAccessPoint) Type() string {
	return "aws_efs_access_point"
}

// LocalName returns the local name for [EfsAccessPoint].
func (eap *EfsAccessPoint) LocalName() string {
	return eap.Name
}

// Configuration returns the configuration (args) for [EfsAccessPoint].
func (eap *EfsAccessPoint) Configuration() interface{} {
	return eap.Args
}

// DependOn is used for other resources to depend on [EfsAccessPoint].
func (eap *EfsAccessPoint) DependOn() terra.Reference {
	return terra.ReferenceResource(eap)
}

// Dependencies returns the list of resources [EfsAccessPoint] depends_on.
func (eap *EfsAccessPoint) Dependencies() terra.Dependencies {
	return eap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EfsAccessPoint].
func (eap *EfsAccessPoint) LifecycleManagement() *terra.Lifecycle {
	return eap.Lifecycle
}

// Attributes returns the attributes for [EfsAccessPoint].
func (eap *EfsAccessPoint) Attributes() efsAccessPointAttributes {
	return efsAccessPointAttributes{ref: terra.ReferenceResource(eap)}
}

// ImportState imports the given attribute values into [EfsAccessPoint]'s state.
func (eap *EfsAccessPoint) ImportState(av io.Reader) error {
	eap.state = &efsAccessPointState{}
	if err := json.NewDecoder(av).Decode(eap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", eap.Type(), eap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EfsAccessPoint] has state.
func (eap *EfsAccessPoint) State() (*efsAccessPointState, bool) {
	return eap.state, eap.state != nil
}

// StateMust returns the state for [EfsAccessPoint]. Panics if the state is nil.
func (eap *EfsAccessPoint) StateMust() *efsAccessPointState {
	if eap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", eap.Type(), eap.LocalName()))
	}
	return eap.state
}

// EfsAccessPointArgs contains the configurations for aws_efs_access_point.
type EfsAccessPointArgs struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// PosixUser: optional
	PosixUser *efsaccesspoint.PosixUser `hcl:"posix_user,block"`
	// RootDirectory: optional
	RootDirectory *efsaccesspoint.RootDirectory `hcl:"root_directory,block"`
}
type efsAccessPointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_efs_access_point.
func (eap efsAccessPointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("arn"))
}

// FileSystemArn returns a reference to field file_system_arn of aws_efs_access_point.
func (eap efsAccessPointAttributes) FileSystemArn() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("file_system_arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_efs_access_point.
func (eap efsAccessPointAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_access_point.
func (eap efsAccessPointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_efs_access_point.
func (eap efsAccessPointAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_efs_access_point.
func (eap efsAccessPointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eap.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_efs_access_point.
func (eap efsAccessPointAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eap.ref.Append("tags_all"))
}

func (eap efsAccessPointAttributes) PosixUser() terra.ListValue[efsaccesspoint.PosixUserAttributes] {
	return terra.ReferenceAsList[efsaccesspoint.PosixUserAttributes](eap.ref.Append("posix_user"))
}

func (eap efsAccessPointAttributes) RootDirectory() terra.ListValue[efsaccesspoint.RootDirectoryAttributes] {
	return terra.ReferenceAsList[efsaccesspoint.RootDirectoryAttributes](eap.ref.Append("root_directory"))
}

type efsAccessPointState struct {
	Arn           string                              `json:"arn"`
	FileSystemArn string                              `json:"file_system_arn"`
	FileSystemId  string                              `json:"file_system_id"`
	Id            string                              `json:"id"`
	OwnerId       string                              `json:"owner_id"`
	Tags          map[string]string                   `json:"tags"`
	TagsAll       map[string]string                   `json:"tags_all"`
	PosixUser     []efsaccesspoint.PosixUserState     `json:"posix_user"`
	RootDirectory []efsaccesspoint.RootDirectoryState `json:"root_directory"`
}
