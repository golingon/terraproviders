// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataClientOpenidUserinfo creates a new instance of [DataClientOpenidUserinfo].
func NewDataClientOpenidUserinfo(name string, args DataClientOpenidUserinfoArgs) *DataClientOpenidUserinfo {
	return &DataClientOpenidUserinfo{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataClientOpenidUserinfo)(nil)

// DataClientOpenidUserinfo represents the Terraform data resource google_client_openid_userinfo.
type DataClientOpenidUserinfo struct {
	Name string
	Args DataClientOpenidUserinfoArgs
}

// DataSource returns the Terraform object type for [DataClientOpenidUserinfo].
func (cou *DataClientOpenidUserinfo) DataSource() string {
	return "google_client_openid_userinfo"
}

// LocalName returns the local name for [DataClientOpenidUserinfo].
func (cou *DataClientOpenidUserinfo) LocalName() string {
	return cou.Name
}

// Configuration returns the configuration (args) for [DataClientOpenidUserinfo].
func (cou *DataClientOpenidUserinfo) Configuration() interface{} {
	return cou.Args
}

// Attributes returns the attributes for [DataClientOpenidUserinfo].
func (cou *DataClientOpenidUserinfo) Attributes() dataClientOpenidUserinfoAttributes {
	return dataClientOpenidUserinfoAttributes{ref: terra.ReferenceDataResource(cou)}
}

// DataClientOpenidUserinfoArgs contains the configurations for google_client_openid_userinfo.
type DataClientOpenidUserinfoArgs struct{}
type dataClientOpenidUserinfoAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of google_client_openid_userinfo.
func (cou dataClientOpenidUserinfoAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(cou.ref.Append("email"))
}

// Id returns a reference to field id of google_client_openid_userinfo.
func (cou dataClientOpenidUserinfoAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cou.ref.Append("id"))
}
