package xlambda_test

import (
	"strconv"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"

	"github.com/strongishllama/xlambda"
)

func TestUnmarshalDynamoDBAttributeValueMap(t *testing.T) {
	type childTestData struct {
		Bool              bool          `json:"bool"`
		ByteSlice         []byte        `json:"byteSlice"`
		SliceOfByteSlices [][]byte      `json:"sliceOfByteSlices"`
		Slice             []interface{} `json:"slice"`
		Int               int           `json:"int"`
		IntSlice          []int         `json:"intSlice"`
		Float             float64       `json:"float"`
		FloatSlice        []float64     `json:"floatSlice"`
		String            string        `json:"string"`
		StringSlice       []string      `json:"stringSlice"`
		Nil               *string       `json:"nil"`
	}
	type testData struct {
		Bool              bool          `json:"bool"`
		ByteSlice         []byte        `json:"byteSlice"`
		SliceOfByteSlices [][]byte      `json:"sliceOfByteSlices"`
		Slice             []interface{} `json:"slice"`
		Struct            childTestData `json:"struct"`
		Int               int           `json:"int"`
		IntSlice          []int         `json:"intSlice"`
		Float             float64       `json:"float"`
		FloatSlice        []float64     `json:"floatSlice"`
		String            string        `json:"string"`
		StringSlice       []string      `json:"stringSlice"`
		Nil               *string       `json:"nil"`
	}

	want := &testData{
		Bool:              true,
		ByteSlice:         []byte("byteSlice"),
		SliceOfByteSlices: stringSliceToSliceOfByteSlices("sliceOfByteSlices1", "sliceOfByteSlices2"),
		Slice: []interface{}{
			true,
			"string",
			// 1, // Doesn't work, an int will come back as a float64.
			1.5,
		},
		Struct: childTestData{
			Bool:              true,
			ByteSlice:         []byte("byteSlice"),
			SliceOfByteSlices: stringSliceToSliceOfByteSlices("sliceOfByteSlices1", "sliceOfByteSlices2"),
			Slice: []interface{}{
				true,
				"string",
				// 1, // Doesn't work, an int will come back as a float64.
				1.5,
			},
			Int:         1,
			IntSlice:    []int{1, 2, 3},
			Float:       1.5,
			FloatSlice:  []float64{1.5, 2.5, 3.5},
			String:      "string",
			StringSlice: []string{"string1", "string2", "string3"},
			Nil:         nil,
		},
		Int:         1,
		IntSlice:    []int{1, 2, 3},
		Float:       1.5,
		FloatSlice:  []float64{1.5, 2.5, 3.5},
		String:      "string",
		StringSlice: []string{"string1", "string2", "string3"},
		Nil:         nil,
	}

	td := &testData{}
	require.NoError(t, xlambda.UnmarshalDynamoDBEventAttributeValues(map[string]events.DynamoDBAttributeValue{
		"bool":              events.NewBooleanAttribute(want.Bool),
		"byteSlice":         events.NewBinaryAttribute(want.ByteSlice),
		"sliceOfByteSlices": events.NewBinarySetAttribute(want.SliceOfByteSlices),
		"slice": events.NewListAttribute([]events.DynamoDBAttributeValue{
			events.NewBooleanAttribute(want.Slice[0].(bool)),
			events.NewStringAttribute(want.Slice[1].(string)),
			events.NewNumberAttribute(float64ToString(want.Slice[2].(float64))),
		}),
		"struct": events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{
			"bool":              events.NewBooleanAttribute(want.Bool),
			"byteSlice":         events.NewBinaryAttribute(want.ByteSlice),
			"sliceOfByteSlices": events.NewBinarySetAttribute(want.SliceOfByteSlices),
			"slice": events.NewListAttribute([]events.DynamoDBAttributeValue{
				events.NewBooleanAttribute(want.Slice[0].(bool)),
				events.NewStringAttribute(want.Slice[1].(string)),
				events.NewNumberAttribute(float64ToString(want.Slice[2].(float64))),
			}),
			"int":         events.NewNumberAttribute(strconv.Itoa(want.Int)),
			"intSlice":    events.NewNumberSetAttribute([]string{strconv.Itoa(want.IntSlice[0]), strconv.Itoa(want.IntSlice[1]), strconv.Itoa(want.IntSlice[2])}),
			"float":       events.NewNumberAttribute(float64ToString(want.Float)),
			"floatSlice":  events.NewNumberSetAttribute([]string{float64ToString(want.FloatSlice[0]), float64ToString(want.FloatSlice[1]), float64ToString(want.FloatSlice[2])}),
			"string":      events.NewStringAttribute(want.String),
			"stringSlice": events.NewStringSetAttribute(want.StringSlice),
			"nil":         events.NewNullAttribute(),
		}),
		"int":         events.NewNumberAttribute(strconv.Itoa(want.Int)),
		"intSlice":    events.NewNumberSetAttribute([]string{strconv.Itoa(want.IntSlice[0]), strconv.Itoa(want.IntSlice[1]), strconv.Itoa(want.IntSlice[2])}),
		"float":       events.NewNumberAttribute(float64ToString(want.Float)),
		"floatSlice":  events.NewNumberSetAttribute([]string{float64ToString(want.FloatSlice[0]), float64ToString(want.FloatSlice[1]), float64ToString(want.FloatSlice[2])}),
		"string":      events.NewStringAttribute(want.String),
		"stringSlice": events.NewStringSetAttribute(want.StringSlice),
		"nil":         events.NewNullAttribute(),
	}, td))

	require.NotNil(t, td)
	require.Equal(t, want, td)
}

func stringSliceToSliceOfByteSlices(stringSlice ...string) [][]byte {
	sliceOfByteSlices := make([][]byte, len(stringSlice))

	for i := 0; i < len(stringSlice); i++ {
		sliceOfByteSlices[i] = []byte(stringSlice[i])
	}

	return sliceOfByteSlices
}

func float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}
