// Simulate black box testing
package wisdom_test

import (
	"encoding/json"
	"testing"

	. "github.com/pytyagi/wisdom/lib/wisdom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuote(t *testing.T) {
	q := NewQuote("Learn Go!", "piyush")
	assert.Equal(t, "Learn Go!", q.Quote)
	assert.Equal(t, "piyush", q.Author)
}

func TestQuoteAsString(t *testing.T) {
	q := NewQuote("Learn Go!", "piyush")
	assert.Equal(t, "Learn Go!\n - piyush", q.String())
}

func TestQuoteAsJSON(t *testing.T) {
	q := NewQuote("Learn Go!", "piyush")
	expected := `{"quote":"Learn Go!","author":"piyush"}`

	byte, err := json.Marshal(q)
	// use require if test should niot acrry on else use assert to log the rror
	require.Nil(t, err)
	assert.JSONEq(t, expected, string(byte))

}
