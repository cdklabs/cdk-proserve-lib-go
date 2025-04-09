package types


// Represents types of destructive operations that can be performed on resources.
//
// Destructive operations are actions that modify or remove existing resources,
// potentially resulting in data loss if not handled properly.
// Experimental.
type DestructiveOperation string

const (
	// Indicates an operation that modifies existing resources.
	// Experimental.
	DestructiveOperation_UPDATE DestructiveOperation = "UPDATE"
	// Indicates an operation that removes resources.
	// Experimental.
	DestructiveOperation_DELETE DestructiveOperation = "DELETE"
	// Represents all types of destructive operations.
	// Experimental.
	DestructiveOperation_ALL DestructiveOperation = "ALL"
)

