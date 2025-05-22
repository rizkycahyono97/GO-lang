package test

import (
	"github.com/stretchr/testify/assert"
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"
)

// =============================//
// Test Clean up Function
// =============================//
func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("nama_file")
	assert.NotNil(t, connection)

	cleanup()
}
