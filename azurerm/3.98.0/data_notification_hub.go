// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datanotificationhub "github.com/golingon/terraproviders/azurerm/3.98.0/datanotificationhub"
)

// NewDataNotificationHub creates a new instance of [DataNotificationHub].
func NewDataNotificationHub(name string, args DataNotificationHubArgs) *DataNotificationHub {
	return &DataNotificationHub{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNotificationHub)(nil)

// DataNotificationHub represents the Terraform data resource azurerm_notification_hub.
type DataNotificationHub struct {
	Name string
	Args DataNotificationHubArgs
}

// DataSource returns the Terraform object type for [DataNotificationHub].
func (nh *DataNotificationHub) DataSource() string {
	return "azurerm_notification_hub"
}

// LocalName returns the local name for [DataNotificationHub].
func (nh *DataNotificationHub) LocalName() string {
	return nh.Name
}

// Configuration returns the configuration (args) for [DataNotificationHub].
func (nh *DataNotificationHub) Configuration() interface{} {
	return nh.Args
}

// Attributes returns the attributes for [DataNotificationHub].
func (nh *DataNotificationHub) Attributes() dataNotificationHubAttributes {
	return dataNotificationHubAttributes{ref: terra.ReferenceDataResource(nh)}
}

// DataNotificationHubArgs contains the configurations for azurerm_notification_hub.
type DataNotificationHubArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ApnsCredential: min=0
	ApnsCredential []datanotificationhub.ApnsCredential `hcl:"apns_credential,block" validate:"min=0"`
	// GcmCredential: min=0
	GcmCredential []datanotificationhub.GcmCredential `hcl:"gcm_credential,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanotificationhub.Timeouts `hcl:"timeouts,block"`
}
type dataNotificationHubAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_notification_hub.
func (nh dataNotificationHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_notification_hub.
func (nh dataNotificationHubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_notification_hub.
func (nh dataNotificationHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("name"))
}

// NamespaceName returns a reference to field namespace_name of azurerm_notification_hub.
func (nh dataNotificationHubAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("namespace_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_notification_hub.
func (nh dataNotificationHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nh.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_notification_hub.
func (nh dataNotificationHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nh.ref.Append("tags"))
}

func (nh dataNotificationHubAttributes) ApnsCredential() terra.ListValue[datanotificationhub.ApnsCredentialAttributes] {
	return terra.ReferenceAsList[datanotificationhub.ApnsCredentialAttributes](nh.ref.Append("apns_credential"))
}

func (nh dataNotificationHubAttributes) GcmCredential() terra.ListValue[datanotificationhub.GcmCredentialAttributes] {
	return terra.ReferenceAsList[datanotificationhub.GcmCredentialAttributes](nh.ref.Append("gcm_credential"))
}

func (nh dataNotificationHubAttributes) Timeouts() datanotificationhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanotificationhub.TimeoutsAttributes](nh.ref.Append("timeouts"))
}
