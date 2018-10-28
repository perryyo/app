package app

// JavaScript is the interface that provide a limited support to deal
// with javascript.
type JavaScript interface {
	// Get get the value pointed by the pipeline and unmarshal it into the given
	// output.
	Get(pipeline string, out interface{})

	// Set set the value pointed by the pipeline with the given input.
	Set(pipeline string, v interface{})

	// Call calls the function pointed by the pipeline with the given
	// inputs as arguments.
	// If the function returns a value, it will be unmarshaled into the output.
	Call(pipeline string, out interface{}, in ...interface{})
}

// Node is the interface that describes a DOM node.
// Node have a limited support.
type Node interface {
	// The node identifier.
	ID() string

	// The parent node.
	Parent() Node

	// Get get the javascript node field and unmarshal it into the output.
	Get(field string, out interface{})

	// Set set the javascript field with the given input.
	Set(field string, in interface{})
}
