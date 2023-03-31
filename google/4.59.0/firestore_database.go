// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	firestoredatabase "github.com/golingon/terraproviders/google/4.59.0/firestoredatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewFirestoreDatabase(name string, args FirestoreDatabaseArgs) *FirestoreDatabase {
	return &FirestoreDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirestoreDatabase)(nil)

type FirestoreDatabase struct {
	Name  string
	Args  FirestoreDatabaseArgs
	state *firestoreDatabaseState
}

func (fd *FirestoreDatabase) Type() string {
	return "google_firestore_database"
}

func (fd *FirestoreDatabase) LocalName() string {
	return fd.Name
}

func (fd *FirestoreDatabase) Configuration() interface{} {
	return fd.Args
}

func (fd *FirestoreDatabase) Attributes() firestoreDatabaseAttributes {
	return firestoreDatabaseAttributes{ref: terra.ReferenceResource(fd)}
}

func (fd *FirestoreDatabase) ImportState(av io.Reader) error {
	fd.state = &firestoreDatabaseState{}
	if err := json.NewDecoder(av).Decode(fd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fd.Type(), fd.LocalName(), err)
	}
	return nil
}

func (fd *FirestoreDatabase) State() (*firestoreDatabaseState, bool) {
	return fd.state, fd.state != nil
}

func (fd *FirestoreDatabase) StateMust() *firestoreDatabaseState {
	if fd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fd.Type(), fd.LocalName()))
	}
	return fd.state
}

func (fd *FirestoreDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(fd)
}

type FirestoreDatabaseArgs struct {
	// AppEngineIntegrationMode: string, optional
	AppEngineIntegrationMode terra.StringValue `hcl:"app_engine_integration_mode,attr"`
	// ConcurrencyMode: string, optional
	ConcurrencyMode terra.StringValue `hcl:"concurrency_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationId: string, required
	LocationId terra.StringValue `hcl:"location_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *firestoredatabase.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that FirestoreDatabase depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type firestoreDatabaseAttributes struct {
	ref terra.Reference
}

func (fd firestoreDatabaseAttributes) AppEngineIntegrationMode() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("app_engine_integration_mode"))
}

func (fd firestoreDatabaseAttributes) ConcurrencyMode() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("concurrency_mode"))
}

func (fd firestoreDatabaseAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("create_time"))
}

func (fd firestoreDatabaseAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("etag"))
}

func (fd firestoreDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("id"))
}

func (fd firestoreDatabaseAttributes) KeyPrefix() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("key_prefix"))
}

func (fd firestoreDatabaseAttributes) LocationId() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("location_id"))
}

func (fd firestoreDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("name"))
}

func (fd firestoreDatabaseAttributes) Project() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("project"))
}

func (fd firestoreDatabaseAttributes) Type() terra.StringValue {
	return terra.ReferenceString(fd.ref.Append("type"))
}

func (fd firestoreDatabaseAttributes) Timeouts() firestoredatabase.TimeoutsAttributes {
	return terra.ReferenceSingle[firestoredatabase.TimeoutsAttributes](fd.ref.Append("timeouts"))
}

type firestoreDatabaseState struct {
	AppEngineIntegrationMode string                           `json:"app_engine_integration_mode"`
	ConcurrencyMode          string                           `json:"concurrency_mode"`
	CreateTime               string                           `json:"create_time"`
	Etag                     string                           `json:"etag"`
	Id                       string                           `json:"id"`
	KeyPrefix                string                           `json:"key_prefix"`
	LocationId               string                           `json:"location_id"`
	Name                     string                           `json:"name"`
	Project                  string                           `json:"project"`
	Type                     string                           `json:"type"`
	Timeouts                 *firestoredatabase.TimeoutsState `json:"timeouts"`
}
