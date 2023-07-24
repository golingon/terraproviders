// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	quicksightnamespace "github.com/golingon/terraproviders/aws/5.9.0/quicksightnamespace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQuicksightNamespace creates a new instance of [QuicksightNamespace].
func NewQuicksightNamespace(name string, args QuicksightNamespaceArgs) *QuicksightNamespace {
	return &QuicksightNamespace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QuicksightNamespace)(nil)

// QuicksightNamespace represents the Terraform resource aws_quicksight_namespace.
type QuicksightNamespace struct {
	Name      string
	Args      QuicksightNamespaceArgs
	state     *quicksightNamespaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QuicksightNamespace].
func (qn *QuicksightNamespace) Type() string {
	return "aws_quicksight_namespace"
}

// LocalName returns the local name for [QuicksightNamespace].
func (qn *QuicksightNamespace) LocalName() string {
	return qn.Name
}

// Configuration returns the configuration (args) for [QuicksightNamespace].
func (qn *QuicksightNamespace) Configuration() interface{} {
	return qn.Args
}

// DependOn is used for other resources to depend on [QuicksightNamespace].
func (qn *QuicksightNamespace) DependOn() terra.Reference {
	return terra.ReferenceResource(qn)
}

// Dependencies returns the list of resources [QuicksightNamespace] depends_on.
func (qn *QuicksightNamespace) Dependencies() terra.Dependencies {
	return qn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QuicksightNamespace].
func (qn *QuicksightNamespace) LifecycleManagement() *terra.Lifecycle {
	return qn.Lifecycle
}

// Attributes returns the attributes for [QuicksightNamespace].
func (qn *QuicksightNamespace) Attributes() quicksightNamespaceAttributes {
	return quicksightNamespaceAttributes{ref: terra.ReferenceResource(qn)}
}

// ImportState imports the given attribute values into [QuicksightNamespace]'s state.
func (qn *QuicksightNamespace) ImportState(av io.Reader) error {
	qn.state = &quicksightNamespaceState{}
	if err := json.NewDecoder(av).Decode(qn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qn.Type(), qn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QuicksightNamespace] has state.
func (qn *QuicksightNamespace) State() (*quicksightNamespaceState, bool) {
	return qn.state, qn.state != nil
}

// StateMust returns the state for [QuicksightNamespace]. Panics if the state is nil.
func (qn *QuicksightNamespace) StateMust() *quicksightNamespaceState {
	if qn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qn.Type(), qn.LocalName()))
	}
	return qn.state
}

// QuicksightNamespaceArgs contains the configurations for aws_quicksight_namespace.
type QuicksightNamespaceArgs struct {
	// AwsAccountId: string, optional
	AwsAccountId terra.StringValue `hcl:"aws_account_id,attr"`
	// IdentityStore: string, optional
	IdentityStore terra.StringValue `hcl:"identity_store,attr"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *quicksightnamespace.Timeouts `hcl:"timeouts,block"`
}
type quicksightNamespaceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("arn"))
}

// AwsAccountId returns a reference to field aws_account_id of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) AwsAccountId() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("aws_account_id"))
}

// CapacityRegion returns a reference to field capacity_region of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) CapacityRegion() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("capacity_region"))
}

// CreationStatus returns a reference to field creation_status of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) CreationStatus() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("creation_status"))
}

// Id returns a reference to field id of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("id"))
}

// IdentityStore returns a reference to field identity_store of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) IdentityStore() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("identity_store"))
}

// Namespace returns a reference to field namespace of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(qn.ref.Append("namespace"))
}

// Tags returns a reference to field tags of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_quicksight_namespace.
func (qn quicksightNamespaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qn.ref.Append("tags_all"))
}

func (qn quicksightNamespaceAttributes) Timeouts() quicksightnamespace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[quicksightnamespace.TimeoutsAttributes](qn.ref.Append("timeouts"))
}

type quicksightNamespaceState struct {
	Arn            string                             `json:"arn"`
	AwsAccountId   string                             `json:"aws_account_id"`
	CapacityRegion string                             `json:"capacity_region"`
	CreationStatus string                             `json:"creation_status"`
	Id             string                             `json:"id"`
	IdentityStore  string                             `json:"identity_store"`
	Namespace      string                             `json:"namespace"`
	Tags           map[string]string                  `json:"tags"`
	TagsAll        map[string]string                  `json:"tags_all"`
	Timeouts       *quicksightnamespace.TimeoutsState `json:"timeouts"`
}
