// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	beyondcorpappconnector "github.com/golingon/terraproviders/googlebeta/4.77.0/beyondcorpappconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBeyondcorpAppConnector creates a new instance of [BeyondcorpAppConnector].
func NewBeyondcorpAppConnector(name string, args BeyondcorpAppConnectorArgs) *BeyondcorpAppConnector {
	return &BeyondcorpAppConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BeyondcorpAppConnector)(nil)

// BeyondcorpAppConnector represents the Terraform resource google_beyondcorp_app_connector.
type BeyondcorpAppConnector struct {
	Name      string
	Args      BeyondcorpAppConnectorArgs
	state     *beyondcorpAppConnectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BeyondcorpAppConnector].
func (bac *BeyondcorpAppConnector) Type() string {
	return "google_beyondcorp_app_connector"
}

// LocalName returns the local name for [BeyondcorpAppConnector].
func (bac *BeyondcorpAppConnector) LocalName() string {
	return bac.Name
}

// Configuration returns the configuration (args) for [BeyondcorpAppConnector].
func (bac *BeyondcorpAppConnector) Configuration() interface{} {
	return bac.Args
}

// DependOn is used for other resources to depend on [BeyondcorpAppConnector].
func (bac *BeyondcorpAppConnector) DependOn() terra.Reference {
	return terra.ReferenceResource(bac)
}

// Dependencies returns the list of resources [BeyondcorpAppConnector] depends_on.
func (bac *BeyondcorpAppConnector) Dependencies() terra.Dependencies {
	return bac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BeyondcorpAppConnector].
func (bac *BeyondcorpAppConnector) LifecycleManagement() *terra.Lifecycle {
	return bac.Lifecycle
}

// Attributes returns the attributes for [BeyondcorpAppConnector].
func (bac *BeyondcorpAppConnector) Attributes() beyondcorpAppConnectorAttributes {
	return beyondcorpAppConnectorAttributes{ref: terra.ReferenceResource(bac)}
}

// ImportState imports the given attribute values into [BeyondcorpAppConnector]'s state.
func (bac *BeyondcorpAppConnector) ImportState(av io.Reader) error {
	bac.state = &beyondcorpAppConnectorState{}
	if err := json.NewDecoder(av).Decode(bac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bac.Type(), bac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BeyondcorpAppConnector] has state.
func (bac *BeyondcorpAppConnector) State() (*beyondcorpAppConnectorState, bool) {
	return bac.state, bac.state != nil
}

// StateMust returns the state for [BeyondcorpAppConnector]. Panics if the state is nil.
func (bac *BeyondcorpAppConnector) StateMust() *beyondcorpAppConnectorState {
	if bac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bac.Type(), bac.LocalName()))
	}
	return bac.state
}

// BeyondcorpAppConnectorArgs contains the configurations for google_beyondcorp_app_connector.
type BeyondcorpAppConnectorArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// PrincipalInfo: required
	PrincipalInfo *beyondcorpappconnector.PrincipalInfo `hcl:"principal_info,block" validate:"required"`
	// Timeouts: optional
	Timeouts *beyondcorpappconnector.Timeouts `hcl:"timeouts,block"`
}
type beyondcorpAppConnectorAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("display_name"))
}

// Id returns a reference to field id of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bac.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("region"))
}

// State returns a reference to field state of google_beyondcorp_app_connector.
func (bac beyondcorpAppConnectorAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("state"))
}

func (bac beyondcorpAppConnectorAttributes) PrincipalInfo() terra.ListValue[beyondcorpappconnector.PrincipalInfoAttributes] {
	return terra.ReferenceAsList[beyondcorpappconnector.PrincipalInfoAttributes](bac.ref.Append("principal_info"))
}

func (bac beyondcorpAppConnectorAttributes) Timeouts() beyondcorpappconnector.TimeoutsAttributes {
	return terra.ReferenceAsSingle[beyondcorpappconnector.TimeoutsAttributes](bac.ref.Append("timeouts"))
}

type beyondcorpAppConnectorState struct {
	DisplayName   string                                      `json:"display_name"`
	Id            string                                      `json:"id"`
	Labels        map[string]string                           `json:"labels"`
	Name          string                                      `json:"name"`
	Project       string                                      `json:"project"`
	Region        string                                      `json:"region"`
	State         string                                      `json:"state"`
	PrincipalInfo []beyondcorpappconnector.PrincipalInfoState `json:"principal_info"`
	Timeouts      *beyondcorpappconnector.TimeoutsState       `json:"timeouts"`
}
