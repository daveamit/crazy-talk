package crazytalk

// CrazyTalk holds the indended functionality.
// NewReflectionCrazyTalk will provide implementation via Reflection implementation
// NewProtoCrazyTalk will provide implementation via .proto file
type CrazyTalk interface {
	ListServices() ([]Service, error)
	InvokeRPC(rpc string, JSONPayload string) (JSONResponse string, err error)
}
