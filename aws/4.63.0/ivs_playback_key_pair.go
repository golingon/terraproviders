// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ivsplaybackkeypair "github.com/golingon/terraproviders/aws/4.63.0/ivsplaybackkeypair"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIvsPlaybackKeyPair creates a new instance of [IvsPlaybackKeyPair].
func NewIvsPlaybackKeyPair(name string, args IvsPlaybackKeyPairArgs) *IvsPlaybackKeyPair {
	return &IvsPlaybackKeyPair{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IvsPlaybackKeyPair)(nil)

// IvsPlaybackKeyPair represents the Terraform resource aws_ivs_playback_key_pair.
type IvsPlaybackKeyPair struct {
	Name      string
	Args      IvsPlaybackKeyPairArgs
	state     *ivsPlaybackKeyPairState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IvsPlaybackKeyPair].
func (ipkp *IvsPlaybackKeyPair) Type() string {
	return "aws_ivs_playback_key_pair"
}

// LocalName returns the local name for [IvsPlaybackKeyPair].
func (ipkp *IvsPlaybackKeyPair) LocalName() string {
	return ipkp.Name
}

// Configuration returns the configuration (args) for [IvsPlaybackKeyPair].
func (ipkp *IvsPlaybackKeyPair) Configuration() interface{} {
	return ipkp.Args
}

// DependOn is used for other resources to depend on [IvsPlaybackKeyPair].
func (ipkp *IvsPlaybackKeyPair) DependOn() terra.Reference {
	return terra.ReferenceResource(ipkp)
}

// Dependencies returns the list of resources [IvsPlaybackKeyPair] depends_on.
func (ipkp *IvsPlaybackKeyPair) Dependencies() terra.Dependencies {
	return ipkp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IvsPlaybackKeyPair].
func (ipkp *IvsPlaybackKeyPair) LifecycleManagement() *terra.Lifecycle {
	return ipkp.Lifecycle
}

// Attributes returns the attributes for [IvsPlaybackKeyPair].
func (ipkp *IvsPlaybackKeyPair) Attributes() ivsPlaybackKeyPairAttributes {
	return ivsPlaybackKeyPairAttributes{ref: terra.ReferenceResource(ipkp)}
}

// ImportState imports the given attribute values into [IvsPlaybackKeyPair]'s state.
func (ipkp *IvsPlaybackKeyPair) ImportState(av io.Reader) error {
	ipkp.state = &ivsPlaybackKeyPairState{}
	if err := json.NewDecoder(av).Decode(ipkp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipkp.Type(), ipkp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IvsPlaybackKeyPair] has state.
func (ipkp *IvsPlaybackKeyPair) State() (*ivsPlaybackKeyPairState, bool) {
	return ipkp.state, ipkp.state != nil
}

// StateMust returns the state for [IvsPlaybackKeyPair]. Panics if the state is nil.
func (ipkp *IvsPlaybackKeyPair) StateMust() *ivsPlaybackKeyPairState {
	if ipkp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipkp.Type(), ipkp.LocalName()))
	}
	return ipkp.state
}

// IvsPlaybackKeyPairArgs contains the configurations for aws_ivs_playback_key_pair.
type IvsPlaybackKeyPairArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PublicKey: string, required
	PublicKey terra.StringValue `hcl:"public_key,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *ivsplaybackkeypair.Timeouts `hcl:"timeouts,block"`
}
type ivsPlaybackKeyPairAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ipkp.ref.Append("arn"))
}

// Fingerprint returns a reference to field fingerprint of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(ipkp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipkp.ref.Append("id"))
}

// Name returns a reference to field name of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipkp.ref.Append("name"))
}

// PublicKey returns a reference to field public_key of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceAsString(ipkp.ref.Append("public_key"))
}

// Tags returns a reference to field tags of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ipkp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ivs_playback_key_pair.
func (ipkp ivsPlaybackKeyPairAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ipkp.ref.Append("tags_all"))
}

func (ipkp ivsPlaybackKeyPairAttributes) Timeouts() ivsplaybackkeypair.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ivsplaybackkeypair.TimeoutsAttributes](ipkp.ref.Append("timeouts"))
}

type ivsPlaybackKeyPairState struct {
	Arn         string                            `json:"arn"`
	Fingerprint string                            `json:"fingerprint"`
	Id          string                            `json:"id"`
	Name        string                            `json:"name"`
	PublicKey   string                            `json:"public_key"`
	Tags        map[string]string                 `json:"tags"`
	TagsAll     map[string]string                 `json:"tags_all"`
	Timeouts    *ivsplaybackkeypair.TimeoutsState `json:"timeouts"`
}
