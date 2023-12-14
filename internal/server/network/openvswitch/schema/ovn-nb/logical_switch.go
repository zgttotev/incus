// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovsmodel

const LogicalSwitchTable = "Logical_Switch"

// LogicalSwitch defines an object in Logical_Switch table
type LogicalSwitch struct {
	UUID              string            `ovsdb:"_uuid"`
	ACLs              []string          `ovsdb:"acls"`
	Copp              *string           `ovsdb:"copp"`
	DNSRecords        []string          `ovsdb:"dns_records"`
	ExternalIDs       map[string]string `ovsdb:"external_ids"`
	ForwardingGroups  []string          `ovsdb:"forwarding_groups"`
	LoadBalancer      []string          `ovsdb:"load_balancer"`
	LoadBalancerGroup []string          `ovsdb:"load_balancer_group"`
	Name              string            `ovsdb:"name"`
	OtherConfig       map[string]string `ovsdb:"other_config"`
	Ports             []string          `ovsdb:"ports"`
	QOSRules          []string          `ovsdb:"qos_rules"`
}
