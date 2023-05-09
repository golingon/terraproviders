// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vault

import (
	"encoding/json"
	"fmt"
	rabbitmqsecretbackendrole "github.com/golingon/terraproviders/vault/3.15.2/rabbitmqsecretbackendrole"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRabbitmqSecretBackendRole creates a new instance of [RabbitmqSecretBackendRole].
func NewRabbitmqSecretBackendRole(name string, args RabbitmqSecretBackendRoleArgs) *RabbitmqSecretBackendRole {
	return &RabbitmqSecretBackendRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RabbitmqSecretBackendRole)(nil)

// RabbitmqSecretBackendRole represents the Terraform resource vault_rabbitmq_secret_backend_role.
type RabbitmqSecretBackendRole struct {
	Name      string
	Args      RabbitmqSecretBackendRoleArgs
	state     *rabbitmqSecretBackendRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RabbitmqSecretBackendRole].
func (rsbr *RabbitmqSecretBackendRole) Type() string {
	return "vault_rabbitmq_secret_backend_role"
}

// LocalName returns the local name for [RabbitmqSecretBackendRole].
func (rsbr *RabbitmqSecretBackendRole) LocalName() string {
	return rsbr.Name
}

// Configuration returns the configuration (args) for [RabbitmqSecretBackendRole].
func (rsbr *RabbitmqSecretBackendRole) Configuration() interface{} {
	return rsbr.Args
}

// DependOn is used for other resources to depend on [RabbitmqSecretBackendRole].
func (rsbr *RabbitmqSecretBackendRole) DependOn() terra.Reference {
	return terra.ReferenceResource(rsbr)
}

// Dependencies returns the list of resources [RabbitmqSecretBackendRole] depends_on.
func (rsbr *RabbitmqSecretBackendRole) Dependencies() terra.Dependencies {
	return rsbr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RabbitmqSecretBackendRole].
func (rsbr *RabbitmqSecretBackendRole) LifecycleManagement() *terra.Lifecycle {
	return rsbr.Lifecycle
}

// Attributes returns the attributes for [RabbitmqSecretBackendRole].
func (rsbr *RabbitmqSecretBackendRole) Attributes() rabbitmqSecretBackendRoleAttributes {
	return rabbitmqSecretBackendRoleAttributes{ref: terra.ReferenceResource(rsbr)}
}

// ImportState imports the given attribute values into [RabbitmqSecretBackendRole]'s state.
func (rsbr *RabbitmqSecretBackendRole) ImportState(av io.Reader) error {
	rsbr.state = &rabbitmqSecretBackendRoleState{}
	if err := json.NewDecoder(av).Decode(rsbr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rsbr.Type(), rsbr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RabbitmqSecretBackendRole] has state.
func (rsbr *RabbitmqSecretBackendRole) State() (*rabbitmqSecretBackendRoleState, bool) {
	return rsbr.state, rsbr.state != nil
}

// StateMust returns the state for [RabbitmqSecretBackendRole]. Panics if the state is nil.
func (rsbr *RabbitmqSecretBackendRole) StateMust() *rabbitmqSecretBackendRoleState {
	if rsbr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rsbr.Type(), rsbr.LocalName()))
	}
	return rsbr.state
}

// RabbitmqSecretBackendRoleArgs contains the configurations for vault_rabbitmq_secret_backend_role.
type RabbitmqSecretBackendRoleArgs struct {
	// Backend: string, required
	Backend terra.StringValue `hcl:"backend,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// Tags: string, optional
	Tags terra.StringValue `hcl:"tags,attr"`
	// Vhost: min=0
	Vhost []rabbitmqsecretbackendrole.Vhost `hcl:"vhost,block" validate:"min=0"`
	// VhostTopic: min=0
	VhostTopic []rabbitmqsecretbackendrole.VhostTopic `hcl:"vhost_topic,block" validate:"min=0"`
}
type rabbitmqSecretBackendRoleAttributes struct {
	ref terra.Reference
}

// Backend returns a reference to field backend of vault_rabbitmq_secret_backend_role.
func (rsbr rabbitmqSecretBackendRoleAttributes) Backend() terra.StringValue {
	return terra.ReferenceAsString(rsbr.ref.Append("backend"))
}

// Id returns a reference to field id of vault_rabbitmq_secret_backend_role.
func (rsbr rabbitmqSecretBackendRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsbr.ref.Append("id"))
}

// Name returns a reference to field name of vault_rabbitmq_secret_backend_role.
func (rsbr rabbitmqSecretBackendRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsbr.ref.Append("name"))
}

// Namespace returns a reference to field namespace of vault_rabbitmq_secret_backend_role.
func (rsbr rabbitmqSecretBackendRoleAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(rsbr.ref.Append("namespace"))
}

// Tags returns a reference to field tags of vault_rabbitmq_secret_backend_role.
func (rsbr rabbitmqSecretBackendRoleAttributes) Tags() terra.StringValue {
	return terra.ReferenceAsString(rsbr.ref.Append("tags"))
}

func (rsbr rabbitmqSecretBackendRoleAttributes) Vhost() terra.ListValue[rabbitmqsecretbackendrole.VhostAttributes] {
	return terra.ReferenceAsList[rabbitmqsecretbackendrole.VhostAttributes](rsbr.ref.Append("vhost"))
}

func (rsbr rabbitmqSecretBackendRoleAttributes) VhostTopic() terra.ListValue[rabbitmqsecretbackendrole.VhostTopicAttributes] {
	return terra.ReferenceAsList[rabbitmqsecretbackendrole.VhostTopicAttributes](rsbr.ref.Append("vhost_topic"))
}

type rabbitmqSecretBackendRoleState struct {
	Backend    string                                      `json:"backend"`
	Id         string                                      `json:"id"`
	Name       string                                      `json:"name"`
	Namespace  string                                      `json:"namespace"`
	Tags       string                                      `json:"tags"`
	Vhost      []rabbitmqsecretbackendrole.VhostState      `json:"vhost"`
	VhostTopic []rabbitmqsecretbackendrole.VhostTopicState `json:"vhost_topic"`
}
