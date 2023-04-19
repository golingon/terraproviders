// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	transferserver "github.com/golingon/terraproviders/aws/4.63.0/transferserver"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTransferServer creates a new instance of [TransferServer].
func NewTransferServer(name string, args TransferServerArgs) *TransferServer {
	return &TransferServer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransferServer)(nil)

// TransferServer represents the Terraform resource aws_transfer_server.
type TransferServer struct {
	Name      string
	Args      TransferServerArgs
	state     *transferServerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransferServer].
func (ts *TransferServer) Type() string {
	return "aws_transfer_server"
}

// LocalName returns the local name for [TransferServer].
func (ts *TransferServer) LocalName() string {
	return ts.Name
}

// Configuration returns the configuration (args) for [TransferServer].
func (ts *TransferServer) Configuration() interface{} {
	return ts.Args
}

// DependOn is used for other resources to depend on [TransferServer].
func (ts *TransferServer) DependOn() terra.Reference {
	return terra.ReferenceResource(ts)
}

// Dependencies returns the list of resources [TransferServer] depends_on.
func (ts *TransferServer) Dependencies() terra.Dependencies {
	return ts.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransferServer].
func (ts *TransferServer) LifecycleManagement() *terra.Lifecycle {
	return ts.Lifecycle
}

// Attributes returns the attributes for [TransferServer].
func (ts *TransferServer) Attributes() transferServerAttributes {
	return transferServerAttributes{ref: terra.ReferenceResource(ts)}
}

// ImportState imports the given attribute values into [TransferServer]'s state.
func (ts *TransferServer) ImportState(av io.Reader) error {
	ts.state = &transferServerState{}
	if err := json.NewDecoder(av).Decode(ts.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ts.Type(), ts.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransferServer] has state.
func (ts *TransferServer) State() (*transferServerState, bool) {
	return ts.state, ts.state != nil
}

// StateMust returns the state for [TransferServer]. Panics if the state is nil.
func (ts *TransferServer) StateMust() *transferServerState {
	if ts.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ts.Type(), ts.LocalName()))
	}
	return ts.state
}

// TransferServerArgs contains the configurations for aws_transfer_server.
type TransferServerArgs struct {
	// Certificate: string, optional
	Certificate terra.StringValue `hcl:"certificate,attr"`
	// DirectoryId: string, optional
	DirectoryId terra.StringValue `hcl:"directory_id,attr"`
	// Domain: string, optional
	Domain terra.StringValue `hcl:"domain,attr"`
	// EndpointType: string, optional
	EndpointType terra.StringValue `hcl:"endpoint_type,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Function: string, optional
	Function terra.StringValue `hcl:"function,attr"`
	// HostKey: string, optional
	HostKey terra.StringValue `hcl:"host_key,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityProviderType: string, optional
	IdentityProviderType terra.StringValue `hcl:"identity_provider_type,attr"`
	// InvocationRole: string, optional
	InvocationRole terra.StringValue `hcl:"invocation_role,attr"`
	// LoggingRole: string, optional
	LoggingRole terra.StringValue `hcl:"logging_role,attr"`
	// PostAuthenticationLoginBanner: string, optional
	PostAuthenticationLoginBanner terra.StringValue `hcl:"post_authentication_login_banner,attr"`
	// PreAuthenticationLoginBanner: string, optional
	PreAuthenticationLoginBanner terra.StringValue `hcl:"pre_authentication_login_banner,attr"`
	// Protocols: set of string, optional
	Protocols terra.SetValue[terra.StringValue] `hcl:"protocols,attr"`
	// SecurityPolicyName: string, optional
	SecurityPolicyName terra.StringValue `hcl:"security_policy_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
	// EndpointDetails: optional
	EndpointDetails *transferserver.EndpointDetails `hcl:"endpoint_details,block"`
	// ProtocolDetails: optional
	ProtocolDetails *transferserver.ProtocolDetails `hcl:"protocol_details,block"`
	// WorkflowDetails: optional
	WorkflowDetails *transferserver.WorkflowDetails `hcl:"workflow_details,block"`
}
type transferServerAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transfer_server.
func (ts transferServerAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("arn"))
}

// Certificate returns a reference to field certificate of aws_transfer_server.
func (ts transferServerAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("certificate"))
}

// DirectoryId returns a reference to field directory_id of aws_transfer_server.
func (ts transferServerAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("directory_id"))
}

// Domain returns a reference to field domain of aws_transfer_server.
func (ts transferServerAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("domain"))
}

// Endpoint returns a reference to field endpoint of aws_transfer_server.
func (ts transferServerAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("endpoint"))
}

// EndpointType returns a reference to field endpoint_type of aws_transfer_server.
func (ts transferServerAttributes) EndpointType() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("endpoint_type"))
}

// ForceDestroy returns a reference to field force_destroy of aws_transfer_server.
func (ts transferServerAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(ts.ref.Append("force_destroy"))
}

// Function returns a reference to field function of aws_transfer_server.
func (ts transferServerAttributes) Function() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("function"))
}

// HostKey returns a reference to field host_key of aws_transfer_server.
func (ts transferServerAttributes) HostKey() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("host_key"))
}

// HostKeyFingerprint returns a reference to field host_key_fingerprint of aws_transfer_server.
func (ts transferServerAttributes) HostKeyFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("host_key_fingerprint"))
}

// Id returns a reference to field id of aws_transfer_server.
func (ts transferServerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("id"))
}

// IdentityProviderType returns a reference to field identity_provider_type of aws_transfer_server.
func (ts transferServerAttributes) IdentityProviderType() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("identity_provider_type"))
}

// InvocationRole returns a reference to field invocation_role of aws_transfer_server.
func (ts transferServerAttributes) InvocationRole() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("invocation_role"))
}

// LoggingRole returns a reference to field logging_role of aws_transfer_server.
func (ts transferServerAttributes) LoggingRole() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("logging_role"))
}

// PostAuthenticationLoginBanner returns a reference to field post_authentication_login_banner of aws_transfer_server.
func (ts transferServerAttributes) PostAuthenticationLoginBanner() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("post_authentication_login_banner"))
}

// PreAuthenticationLoginBanner returns a reference to field pre_authentication_login_banner of aws_transfer_server.
func (ts transferServerAttributes) PreAuthenticationLoginBanner() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("pre_authentication_login_banner"))
}

// Protocols returns a reference to field protocols of aws_transfer_server.
func (ts transferServerAttributes) Protocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ts.ref.Append("protocols"))
}

// SecurityPolicyName returns a reference to field security_policy_name of aws_transfer_server.
func (ts transferServerAttributes) SecurityPolicyName() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("security_policy_name"))
}

// Tags returns a reference to field tags of aws_transfer_server.
func (ts transferServerAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ts.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transfer_server.
func (ts transferServerAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ts.ref.Append("tags_all"))
}

// Url returns a reference to field url of aws_transfer_server.
func (ts transferServerAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(ts.ref.Append("url"))
}

func (ts transferServerAttributes) EndpointDetails() terra.ListValue[transferserver.EndpointDetailsAttributes] {
	return terra.ReferenceAsList[transferserver.EndpointDetailsAttributes](ts.ref.Append("endpoint_details"))
}

func (ts transferServerAttributes) ProtocolDetails() terra.ListValue[transferserver.ProtocolDetailsAttributes] {
	return terra.ReferenceAsList[transferserver.ProtocolDetailsAttributes](ts.ref.Append("protocol_details"))
}

func (ts transferServerAttributes) WorkflowDetails() terra.ListValue[transferserver.WorkflowDetailsAttributes] {
	return terra.ReferenceAsList[transferserver.WorkflowDetailsAttributes](ts.ref.Append("workflow_details"))
}

type transferServerState struct {
	Arn                           string                                `json:"arn"`
	Certificate                   string                                `json:"certificate"`
	DirectoryId                   string                                `json:"directory_id"`
	Domain                        string                                `json:"domain"`
	Endpoint                      string                                `json:"endpoint"`
	EndpointType                  string                                `json:"endpoint_type"`
	ForceDestroy                  bool                                  `json:"force_destroy"`
	Function                      string                                `json:"function"`
	HostKey                       string                                `json:"host_key"`
	HostKeyFingerprint            string                                `json:"host_key_fingerprint"`
	Id                            string                                `json:"id"`
	IdentityProviderType          string                                `json:"identity_provider_type"`
	InvocationRole                string                                `json:"invocation_role"`
	LoggingRole                   string                                `json:"logging_role"`
	PostAuthenticationLoginBanner string                                `json:"post_authentication_login_banner"`
	PreAuthenticationLoginBanner  string                                `json:"pre_authentication_login_banner"`
	Protocols                     []string                              `json:"protocols"`
	SecurityPolicyName            string                                `json:"security_policy_name"`
	Tags                          map[string]string                     `json:"tags"`
	TagsAll                       map[string]string                     `json:"tags_all"`
	Url                           string                                `json:"url"`
	EndpointDetails               []transferserver.EndpointDetailsState `json:"endpoint_details"`
	ProtocolDetails               []transferserver.ProtocolDetailsState `json:"protocol_details"`
	WorkflowDetails               []transferserver.WorkflowDetailsState `json:"workflow_details"`
}
