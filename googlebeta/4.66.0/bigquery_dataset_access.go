// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	bigquerydatasetaccess "github.com/golingon/terraproviders/googlebeta/4.66.0/bigquerydatasetaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryDatasetAccess creates a new instance of [BigqueryDatasetAccess].
func NewBigqueryDatasetAccess(name string, args BigqueryDatasetAccessArgs) *BigqueryDatasetAccess {
	return &BigqueryDatasetAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryDatasetAccess)(nil)

// BigqueryDatasetAccess represents the Terraform resource google_bigquery_dataset_access.
type BigqueryDatasetAccess struct {
	Name      string
	Args      BigqueryDatasetAccessArgs
	state     *bigqueryDatasetAccessState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryDatasetAccess].
func (bda *BigqueryDatasetAccess) Type() string {
	return "google_bigquery_dataset_access"
}

// LocalName returns the local name for [BigqueryDatasetAccess].
func (bda *BigqueryDatasetAccess) LocalName() string {
	return bda.Name
}

// Configuration returns the configuration (args) for [BigqueryDatasetAccess].
func (bda *BigqueryDatasetAccess) Configuration() interface{} {
	return bda.Args
}

// DependOn is used for other resources to depend on [BigqueryDatasetAccess].
func (bda *BigqueryDatasetAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(bda)
}

// Dependencies returns the list of resources [BigqueryDatasetAccess] depends_on.
func (bda *BigqueryDatasetAccess) Dependencies() terra.Dependencies {
	return bda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryDatasetAccess].
func (bda *BigqueryDatasetAccess) LifecycleManagement() *terra.Lifecycle {
	return bda.Lifecycle
}

// Attributes returns the attributes for [BigqueryDatasetAccess].
func (bda *BigqueryDatasetAccess) Attributes() bigqueryDatasetAccessAttributes {
	return bigqueryDatasetAccessAttributes{ref: terra.ReferenceResource(bda)}
}

// ImportState imports the given attribute values into [BigqueryDatasetAccess]'s state.
func (bda *BigqueryDatasetAccess) ImportState(av io.Reader) error {
	bda.state = &bigqueryDatasetAccessState{}
	if err := json.NewDecoder(av).Decode(bda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bda.Type(), bda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryDatasetAccess] has state.
func (bda *BigqueryDatasetAccess) State() (*bigqueryDatasetAccessState, bool) {
	return bda.state, bda.state != nil
}

// StateMust returns the state for [BigqueryDatasetAccess]. Panics if the state is nil.
func (bda *BigqueryDatasetAccess) StateMust() *bigqueryDatasetAccessState {
	if bda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bda.Type(), bda.LocalName()))
	}
	return bda.state
}

// BigqueryDatasetAccessArgs contains the configurations for google_bigquery_dataset_access.
type BigqueryDatasetAccessArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Domain: string, optional
	Domain terra.StringValue `hcl:"domain,attr"`
	// GroupByEmail: string, optional
	GroupByEmail terra.StringValue `hcl:"group_by_email,attr"`
	// IamMember: string, optional
	IamMember terra.StringValue `hcl:"iam_member,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, optional
	Role terra.StringValue `hcl:"role,attr"`
	// SpecialGroup: string, optional
	SpecialGroup terra.StringValue `hcl:"special_group,attr"`
	// UserByEmail: string, optional
	UserByEmail terra.StringValue `hcl:"user_by_email,attr"`
	// Dataset: optional
	Dataset *bigquerydatasetaccess.Dataset `hcl:"dataset,block"`
	// Routine: optional
	Routine *bigquerydatasetaccess.Routine `hcl:"routine,block"`
	// Timeouts: optional
	Timeouts *bigquerydatasetaccess.Timeouts `hcl:"timeouts,block"`
	// View: optional
	View *bigquerydatasetaccess.View `hcl:"view,block"`
}
type bigqueryDatasetAccessAttributes struct {
	ref terra.Reference
}

// ApiUpdatedMember returns a reference to field api_updated_member of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) ApiUpdatedMember() terra.BoolValue {
	return terra.ReferenceAsBool(bda.ref.Append("api_updated_member"))
}

// DatasetId returns a reference to field dataset_id of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("dataset_id"))
}

// Domain returns a reference to field domain of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("domain"))
}

// GroupByEmail returns a reference to field group_by_email of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) GroupByEmail() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("group_by_email"))
}

// IamMember returns a reference to field iam_member of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) IamMember() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("iam_member"))
}

// Id returns a reference to field id of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("id"))
}

// Project returns a reference to field project of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("project"))
}

// Role returns a reference to field role of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("role"))
}

// SpecialGroup returns a reference to field special_group of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) SpecialGroup() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("special_group"))
}

// UserByEmail returns a reference to field user_by_email of google_bigquery_dataset_access.
func (bda bigqueryDatasetAccessAttributes) UserByEmail() terra.StringValue {
	return terra.ReferenceAsString(bda.ref.Append("user_by_email"))
}

func (bda bigqueryDatasetAccessAttributes) Dataset() terra.ListValue[bigquerydatasetaccess.DatasetAttributes] {
	return terra.ReferenceAsList[bigquerydatasetaccess.DatasetAttributes](bda.ref.Append("dataset"))
}

func (bda bigqueryDatasetAccessAttributes) Routine() terra.ListValue[bigquerydatasetaccess.RoutineAttributes] {
	return terra.ReferenceAsList[bigquerydatasetaccess.RoutineAttributes](bda.ref.Append("routine"))
}

func (bda bigqueryDatasetAccessAttributes) Timeouts() bigquerydatasetaccess.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigquerydatasetaccess.TimeoutsAttributes](bda.ref.Append("timeouts"))
}

func (bda bigqueryDatasetAccessAttributes) View() terra.ListValue[bigquerydatasetaccess.ViewAttributes] {
	return terra.ReferenceAsList[bigquerydatasetaccess.ViewAttributes](bda.ref.Append("view"))
}

type bigqueryDatasetAccessState struct {
	ApiUpdatedMember bool                                 `json:"api_updated_member"`
	DatasetId        string                               `json:"dataset_id"`
	Domain           string                               `json:"domain"`
	GroupByEmail     string                               `json:"group_by_email"`
	IamMember        string                               `json:"iam_member"`
	Id               string                               `json:"id"`
	Project          string                               `json:"project"`
	Role             string                               `json:"role"`
	SpecialGroup     string                               `json:"special_group"`
	UserByEmail      string                               `json:"user_by_email"`
	Dataset          []bigquerydatasetaccess.DatasetState `json:"dataset"`
	Routine          []bigquerydatasetaccess.RoutineState `json:"routine"`
	Timeouts         *bigquerydatasetaccess.TimeoutsState `json:"timeouts"`
	View             []bigquerydatasetaccess.ViewState    `json:"view"`
}
