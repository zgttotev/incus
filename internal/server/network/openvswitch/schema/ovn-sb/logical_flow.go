// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovsmodel

const LogicalFlowTable = "Logical_Flow"

type (
	LogicalFlowPipeline = string
)

var (
	LogicalFlowPipelineIngress LogicalFlowPipeline = "ingress"
	LogicalFlowPipelineEgress  LogicalFlowPipeline = "egress"
)

// LogicalFlow defines an object in Logical_Flow table
type LogicalFlow struct {
	UUID            string              `ovsdb:"_uuid"`
	Actions         string              `ovsdb:"actions"`
	ControllerMeter *string             `ovsdb:"controller_meter"`
	ExternalIDs     map[string]string   `ovsdb:"external_ids"`
	LogicalDatapath *string             `ovsdb:"logical_datapath"`
	LogicalDpGroup  *string             `ovsdb:"logical_dp_group"`
	Match           string              `ovsdb:"match"`
	Pipeline        LogicalFlowPipeline `ovsdb:"pipeline"`
	Priority        int                 `ovsdb:"priority"`
	TableID         int                 `ovsdb:"table_id"`
	Tags            map[string]string   `ovsdb:"tags"`
}
