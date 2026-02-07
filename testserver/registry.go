package testserver

import (
	"sync"

	"google.golang.org/protobuf/proto"
)

// BehaviorFunc is a custom function that can be used to implement complex
// response logic in tests. It receives the request and returns a response
// or error.
type BehaviorFunc func(req proto.Message) (proto.Message, error)

// ResponseRegistry manages mock responses, errors, and captured requests
// for RPC methods. It is thread-safe and can be safely used from multiple
// goroutines.
type ResponseRegistry struct {
	mu               sync.RWMutex
	responses        map[string]proto.Message
	errors           map[string]error
	capturedRequests map[string][]proto.Message
	behaviors        map[string]BehaviorFunc
}

// NewResponseRegistry creates a new ResponseRegistry.
func NewResponseRegistry() *ResponseRegistry {
	return &ResponseRegistry{
		responses:        make(map[string]proto.Message),
		errors:           make(map[string]error),
		capturedRequests: make(map[string][]proto.Message),
		behaviors:        make(map[string]BehaviorFunc),
	}
}

// SetResponse stores a canned response for the given method.
func (r *ResponseRegistry) SetResponse(method string, response proto.Message) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.responses[method] = response
}

// GetResponse retrieves the configured response for the given method.
func (r *ResponseRegistry) GetResponse(method string) proto.Message {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.responses[method]
}

// SetError configures an error to return for the given method.
func (r *ResponseRegistry) SetError(method string, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.errors[method] = err
}

// GetError retrieves the configured error for the given method.
func (r *ResponseRegistry) GetError(method string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.errors[method]
}

// SetBehavior sets a custom behavior function for the given method.
func (r *ResponseRegistry) SetBehavior(method string, fn BehaviorFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.behaviors[method] = fn
}

// GetBehavior retrieves the custom behavior function for the given method.
func (r *ResponseRegistry) GetBehavior(method string) BehaviorFunc {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.behaviors[method]
}

// CaptureRequest records an incoming request for the given method.
func (r *ResponseRegistry) CaptureRequest(method string, req proto.Message) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.capturedRequests[method] = append(r.capturedRequests[method], req)
}

// GetCapturedRequests returns all captured requests for the given method.
func (r *ResponseRegistry) GetCapturedRequests(method string) []proto.Message {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// Return a copy to avoid concurrent modification issues
	requests := r.capturedRequests[method]
	result := make([]proto.Message, len(requests))
	copy(result, requests)
	return result
}

// Reset clears all state from the registry.
func (r *ResponseRegistry) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.responses = make(map[string]proto.Message)
	r.errors = make(map[string]error)
	r.capturedRequests = make(map[string][]proto.Message)
	r.behaviors = make(map[string]BehaviorFunc)
}
