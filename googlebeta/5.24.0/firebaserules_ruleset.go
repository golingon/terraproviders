// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	firebaserulesruleset "github.com/golingon/terraproviders/googlebeta/5.24.0/firebaserulesruleset"
	"io"
)

// NewFirebaserulesRuleset creates a new instance of [FirebaserulesRuleset].
func NewFirebaserulesRuleset(name string, args FirebaserulesRulesetArgs) *FirebaserulesRuleset {
	return &FirebaserulesRuleset{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaserulesRuleset)(nil)

// FirebaserulesRuleset represents the Terraform resource google_firebaserules_ruleset.
type FirebaserulesRuleset struct {
	Name      string
	Args      FirebaserulesRulesetArgs
	state     *firebaserulesRulesetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaserulesRuleset].
func (fr *FirebaserulesRuleset) Type() string {
	return "google_firebaserules_ruleset"
}

// LocalName returns the local name for [FirebaserulesRuleset].
func (fr *FirebaserulesRuleset) LocalName() string {
	return fr.Name
}

// Configuration returns the configuration (args) for [FirebaserulesRuleset].
func (fr *FirebaserulesRuleset) Configuration() interface{} {
	return fr.Args
}

// DependOn is used for other resources to depend on [FirebaserulesRuleset].
func (fr *FirebaserulesRuleset) DependOn() terra.Reference {
	return terra.ReferenceResource(fr)
}

// Dependencies returns the list of resources [FirebaserulesRuleset] depends_on.
func (fr *FirebaserulesRuleset) Dependencies() terra.Dependencies {
	return fr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaserulesRuleset].
func (fr *FirebaserulesRuleset) LifecycleManagement() *terra.Lifecycle {
	return fr.Lifecycle
}

// Attributes returns the attributes for [FirebaserulesRuleset].
func (fr *FirebaserulesRuleset) Attributes() firebaserulesRulesetAttributes {
	return firebaserulesRulesetAttributes{ref: terra.ReferenceResource(fr)}
}

// ImportState imports the given attribute values into [FirebaserulesRuleset]'s state.
func (fr *FirebaserulesRuleset) ImportState(av io.Reader) error {
	fr.state = &firebaserulesRulesetState{}
	if err := json.NewDecoder(av).Decode(fr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fr.Type(), fr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaserulesRuleset] has state.
func (fr *FirebaserulesRuleset) State() (*firebaserulesRulesetState, bool) {
	return fr.state, fr.state != nil
}

// StateMust returns the state for [FirebaserulesRuleset]. Panics if the state is nil.
func (fr *FirebaserulesRuleset) StateMust() *firebaserulesRulesetState {
	if fr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fr.Type(), fr.LocalName()))
	}
	return fr.state
}

// FirebaserulesRulesetArgs contains the configurations for google_firebaserules_ruleset.
type FirebaserulesRulesetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Metadata: min=0
	Metadata []firebaserulesruleset.Metadata `hcl:"metadata,block" validate:"min=0"`
	// Source: required
	Source *firebaserulesruleset.Source `hcl:"source,block" validate:"required"`
	// Timeouts: optional
	Timeouts *firebaserulesruleset.Timeouts `hcl:"timeouts,block"`
}
type firebaserulesRulesetAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_firebaserules_ruleset.
func (fr firebaserulesRulesetAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("create_time"))
}

// Id returns a reference to field id of google_firebaserules_ruleset.
func (fr firebaserulesRulesetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("id"))
}

// Name returns a reference to field name of google_firebaserules_ruleset.
func (fr firebaserulesRulesetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("name"))
}

// Project returns a reference to field project of google_firebaserules_ruleset.
func (fr firebaserulesRulesetAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("project"))
}

func (fr firebaserulesRulesetAttributes) Metadata() terra.ListValue[firebaserulesruleset.MetadataAttributes] {
	return terra.ReferenceAsList[firebaserulesruleset.MetadataAttributes](fr.ref.Append("metadata"))
}

func (fr firebaserulesRulesetAttributes) Source() terra.ListValue[firebaserulesruleset.SourceAttributes] {
	return terra.ReferenceAsList[firebaserulesruleset.SourceAttributes](fr.ref.Append("source"))
}

func (fr firebaserulesRulesetAttributes) Timeouts() firebaserulesruleset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaserulesruleset.TimeoutsAttributes](fr.ref.Append("timeouts"))
}

type firebaserulesRulesetState struct {
	CreateTime string                               `json:"create_time"`
	Id         string                               `json:"id"`
	Name       string                               `json:"name"`
	Project    string                               `json:"project"`
	Metadata   []firebaserulesruleset.MetadataState `json:"metadata"`
	Source     []firebaserulesruleset.SourceState   `json:"source"`
	Timeouts   *firebaserulesruleset.TimeoutsState  `json:"timeouts"`
}
