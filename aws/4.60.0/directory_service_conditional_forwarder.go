// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceConditionalForwarder creates a new instance of [DirectoryServiceConditionalForwarder].
func NewDirectoryServiceConditionalForwarder(name string, args DirectoryServiceConditionalForwarderArgs) *DirectoryServiceConditionalForwarder {
	return &DirectoryServiceConditionalForwarder{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceConditionalForwarder)(nil)

// DirectoryServiceConditionalForwarder represents the Terraform resource aws_directory_service_conditional_forwarder.
type DirectoryServiceConditionalForwarder struct {
	Name      string
	Args      DirectoryServiceConditionalForwarderArgs
	state     *directoryServiceConditionalForwarderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceConditionalForwarder].
func (dscf *DirectoryServiceConditionalForwarder) Type() string {
	return "aws_directory_service_conditional_forwarder"
}

// LocalName returns the local name for [DirectoryServiceConditionalForwarder].
func (dscf *DirectoryServiceConditionalForwarder) LocalName() string {
	return dscf.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceConditionalForwarder].
func (dscf *DirectoryServiceConditionalForwarder) Configuration() interface{} {
	return dscf.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceConditionalForwarder].
func (dscf *DirectoryServiceConditionalForwarder) DependOn() terra.Reference {
	return terra.ReferenceResource(dscf)
}

// Dependencies returns the list of resources [DirectoryServiceConditionalForwarder] depends_on.
func (dscf *DirectoryServiceConditionalForwarder) Dependencies() terra.Dependencies {
	return dscf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceConditionalForwarder].
func (dscf *DirectoryServiceConditionalForwarder) LifecycleManagement() *terra.Lifecycle {
	return dscf.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceConditionalForwarder].
func (dscf *DirectoryServiceConditionalForwarder) Attributes() directoryServiceConditionalForwarderAttributes {
	return directoryServiceConditionalForwarderAttributes{ref: terra.ReferenceResource(dscf)}
}

// ImportState imports the given attribute values into [DirectoryServiceConditionalForwarder]'s state.
func (dscf *DirectoryServiceConditionalForwarder) ImportState(av io.Reader) error {
	dscf.state = &directoryServiceConditionalForwarderState{}
	if err := json.NewDecoder(av).Decode(dscf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dscf.Type(), dscf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceConditionalForwarder] has state.
func (dscf *DirectoryServiceConditionalForwarder) State() (*directoryServiceConditionalForwarderState, bool) {
	return dscf.state, dscf.state != nil
}

// StateMust returns the state for [DirectoryServiceConditionalForwarder]. Panics if the state is nil.
func (dscf *DirectoryServiceConditionalForwarder) StateMust() *directoryServiceConditionalForwarderState {
	if dscf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dscf.Type(), dscf.LocalName()))
	}
	return dscf.state
}

// DirectoryServiceConditionalForwarderArgs contains the configurations for aws_directory_service_conditional_forwarder.
type DirectoryServiceConditionalForwarderArgs struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// DnsIps: list of string, required
	DnsIps terra.ListValue[terra.StringValue] `hcl:"dns_ips,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RemoteDomainName: string, required
	RemoteDomainName terra.StringValue `hcl:"remote_domain_name,attr" validate:"required"`
}
type directoryServiceConditionalForwarderAttributes struct {
	ref terra.Reference
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_conditional_forwarder.
func (dscf directoryServiceConditionalForwarderAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dscf.ref.Append("directory_id"))
}

// DnsIps returns a reference to field dns_ips of aws_directory_service_conditional_forwarder.
func (dscf directoryServiceConditionalForwarderAttributes) DnsIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dscf.ref.Append("dns_ips"))
}

// Id returns a reference to field id of aws_directory_service_conditional_forwarder.
func (dscf directoryServiceConditionalForwarderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dscf.ref.Append("id"))
}

// RemoteDomainName returns a reference to field remote_domain_name of aws_directory_service_conditional_forwarder.
func (dscf directoryServiceConditionalForwarderAttributes) RemoteDomainName() terra.StringValue {
	return terra.ReferenceAsString(dscf.ref.Append("remote_domain_name"))
}

type directoryServiceConditionalForwarderState struct {
	DirectoryId      string   `json:"directory_id"`
	DnsIps           []string `json:"dns_ips"`
	Id               string   `json:"id"`
	RemoteDomainName string   `json:"remote_domain_name"`
}
