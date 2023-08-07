// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	firebaserulesrelease "github.com/golingon/terraproviders/google/4.76.0/firebaserulesrelease"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaserulesRelease creates a new instance of [FirebaserulesRelease].
func NewFirebaserulesRelease(name string, args FirebaserulesReleaseArgs) *FirebaserulesRelease {
	return &FirebaserulesRelease{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaserulesRelease)(nil)

// FirebaserulesRelease represents the Terraform resource google_firebaserules_release.
type FirebaserulesRelease struct {
	Name      string
	Args      FirebaserulesReleaseArgs
	state     *firebaserulesReleaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaserulesRelease].
func (fr *FirebaserulesRelease) Type() string {
	return "google_firebaserules_release"
}

// LocalName returns the local name for [FirebaserulesRelease].
func (fr *FirebaserulesRelease) LocalName() string {
	return fr.Name
}

// Configuration returns the configuration (args) for [FirebaserulesRelease].
func (fr *FirebaserulesRelease) Configuration() interface{} {
	return fr.Args
}

// DependOn is used for other resources to depend on [FirebaserulesRelease].
func (fr *FirebaserulesRelease) DependOn() terra.Reference {
	return terra.ReferenceResource(fr)
}

// Dependencies returns the list of resources [FirebaserulesRelease] depends_on.
func (fr *FirebaserulesRelease) Dependencies() terra.Dependencies {
	return fr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaserulesRelease].
func (fr *FirebaserulesRelease) LifecycleManagement() *terra.Lifecycle {
	return fr.Lifecycle
}

// Attributes returns the attributes for [FirebaserulesRelease].
func (fr *FirebaserulesRelease) Attributes() firebaserulesReleaseAttributes {
	return firebaserulesReleaseAttributes{ref: terra.ReferenceResource(fr)}
}

// ImportState imports the given attribute values into [FirebaserulesRelease]'s state.
func (fr *FirebaserulesRelease) ImportState(av io.Reader) error {
	fr.state = &firebaserulesReleaseState{}
	if err := json.NewDecoder(av).Decode(fr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fr.Type(), fr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaserulesRelease] has state.
func (fr *FirebaserulesRelease) State() (*firebaserulesReleaseState, bool) {
	return fr.state, fr.state != nil
}

// StateMust returns the state for [FirebaserulesRelease]. Panics if the state is nil.
func (fr *FirebaserulesRelease) StateMust() *firebaserulesReleaseState {
	if fr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fr.Type(), fr.LocalName()))
	}
	return fr.state
}

// FirebaserulesReleaseArgs contains the configurations for google_firebaserules_release.
type FirebaserulesReleaseArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RulesetName: string, required
	RulesetName terra.StringValue `hcl:"ruleset_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *firebaserulesrelease.Timeouts `hcl:"timeouts,block"`
}
type firebaserulesReleaseAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("create_time"))
}

// Disabled returns a reference to field disabled of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(fr.ref.Append("disabled"))
}

// Id returns a reference to field id of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("id"))
}

// Name returns a reference to field name of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("name"))
}

// Project returns a reference to field project of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("project"))
}

// RulesetName returns a reference to field ruleset_name of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) RulesetName() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("ruleset_name"))
}

// UpdateTime returns a reference to field update_time of google_firebaserules_release.
func (fr firebaserulesReleaseAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(fr.ref.Append("update_time"))
}

func (fr firebaserulesReleaseAttributes) Timeouts() firebaserulesrelease.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebaserulesrelease.TimeoutsAttributes](fr.ref.Append("timeouts"))
}

type firebaserulesReleaseState struct {
	CreateTime  string                              `json:"create_time"`
	Disabled    bool                                `json:"disabled"`
	Id          string                              `json:"id"`
	Name        string                              `json:"name"`
	Project     string                              `json:"project"`
	RulesetName string                              `json:"ruleset_name"`
	UpdateTime  string                              `json:"update_time"`
	Timeouts    *firebaserulesrelease.TimeoutsState `json:"timeouts"`
}
