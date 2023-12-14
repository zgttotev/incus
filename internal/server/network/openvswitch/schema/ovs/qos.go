// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovsmodel

const QoSTable = "QoS"

// QoS defines an object in QoS table
type QoS struct {
	UUID        string            `ovsdb:"_uuid"`
	ExternalIDs map[string]string `ovsdb:"external_ids"`
	OtherConfig map[string]string `ovsdb:"other_config"`
	Queues      map[int]string    `ovsdb:"queues"`
	Type        string            `ovsdb:"type"`
}
