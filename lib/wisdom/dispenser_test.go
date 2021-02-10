package wisdom_test

import (
	"testing"

	. "github.com/pytyagi/wisdom/lib/wisdom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//-------------------------------------------------------------------------------------------------

func TestNewDispenser(t *testing.T) {

	q1 := Quote{Quote: "first", Author: "amy"}
	q2 := Quote{Quote: "second", Author: "bob"}
	q3 := Quote{Quote: "third", Author: "carol"}

	d := NewDispenser([]Quote{q1, q2, q3})

	require.NotNil(t, d)
	require.Equal(t, 3, d.Count())

	assert.Equal(t, "first\n - amy", d.Get(0).String())
	assert.Equal(t, "second\n - bob", d.Get(1).String())
	assert.Equal(t, "third\n - carol", d.Get(2).String())

}

//-------------------------------------------------------------------------------------------------

func TestDispenserFromFile(t *testing.T) {

	dispenser, err := FromFile("../../quotes.json")
	require.Nil(t, err)

	assert.Equal(t, 31, dispenser.Count())
	quote := dispenser.Get(0)

	assert.Equal(t, "The problem with object-oriented languages is", quote.Quote[0:45])
	assert.Equal(t, "Joe Armstrong, creator of Erlang", quote.Author)
}

//-------------------------------------------------------------------------------------------------

func TestDispenserFromUnknownFile(t *testing.T) {
	_, err := FromFile("unknown.json")
	require.NotNil(t, err)
	assert.Equal(t, "ioutil.ReadFile failed: open unknown.json: no such file or directory", err.Error())
}

//-------------------------------------------------------------------------------------------------

func TestDispenserFromFileInvalidJson(t *testing.T) {
	_, err := FromFile("dispenser_test.go")
	require.NotNil(t, err)
	assert.Equal(t, "json.Unmarshal failed: invalid character 'p' looking for beginning of value", err.Error())
}

//-------------------------------------------------------------------------------------------------

func TestDispenserRandomDistribution(t *testing.T) {

	q1 := Quote{Quote: "first", Author: "adam"}
	q2 := Quote{Quote: "second", Author: "brian"}

	d := NewDispenser([]Quote{q1, q2})
	require.NotNil(t, d)

	distribution := []int{0, 0}
	for i := 0; i < 100; i++ {
		q := d.Random()
		if q == q1 {
			distribution[0]++
		} else {
			distribution[1]++
		}
	}

	assert.True(t, distribution[0] > 40, "100 tests with 2 quotes should be close to 50 each, leaving a little wiggle room")
	assert.True(t, distribution[1] > 40, "100 tests with 2 quotes should be close to 50 each, leaving a little wiggle room")

}
