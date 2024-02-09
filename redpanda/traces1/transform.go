package main

import (
	"encoding/json"

	"github.com/itchyny/gojq"
	"github.com/redpanda-data/redpanda/src/transform-sdk/go/transform"
)

func main() {
	// Register your transform function.
	// This is a good place to perform other setup too.
	transform.OnRecordWritten(doTransform)
}

func doTransform(e transform.WriteEvent, w transform.RecordWriter) error {

	// Unmarshal the incoming message into a map
	var incomingMessage map[string]interface{}
	err := json.Unmarshal(e.Record().Value, &incomingMessage)
	if err != nil {
		return err
	}

	// Create a new jq query
	query, err := gojq.Parse("select( .process_kprobe != null  " +
		"and .process_kprobe.process.pod.namespace != \"jupyter\"   " +
		"and .process_kprobe.process.pod.namespace != \"cert-manager\" " +
		"and .process_kprobe.process.pod.namespace != \"redpanda\" " +
		"and .process_kprobe.process.pod.namespace != \"spark\" " +
		"and .process_kprobe.process.pod.namespace != \"parseable\" " +
		"and .process_kprobe.process.pod.namespace != \"vector\" )| " +
		"\"\\(.time) \\(.process_kprobe.policy_name) \\(.process_kprobe.function_name) \\(.process_kprobe.process.binary) \\(.process_kprobe.process.arguments) \\(.process_kprobe.process.pod.namespace) \\(.process_kprobe.args[] | select(.sock_arg.priority != null) | .sock_arg.priority)\"")

	if err != nil {
		return err
	}

	// Execute the jq query
	iter := query.Run(incomingMessage)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return err
		}
		// Marshal the result back to JSON
		jsonData, err := json.Marshal(v)
		if err != nil {
			return err
		}

		// Create a new record with the JSON data
		record := &transform.Record{
			Key:       e.Record().Key,
			Value:     jsonData,
			Offset:    e.Record().Offset,
			Timestamp: e.Record().Timestamp,
			Headers:   e.Record().Headers,
		}
		// Write the record to the destination topic
		err = w.Write(*record)
		if err != nil {
			return err
		}
	}
	return nil
}
