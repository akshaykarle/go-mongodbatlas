package mongodb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errAPI = APIError{
	Detail:    "A cluster named test is already present in group null.",
	Code:      400,
	ErrorCode: "DUPLICATE_CLUSTER_NAME",
	Reason:    "Bad Request",
}

var errHTTP = fmt.Errorf("unknown host")

func TestAPIError_Error(t *testing.T) {
	err := APIError{}
	if assert.Error(t, err) {
		assert.Equal(t, "", err.Error())
	}
	if assert.Error(t, errAPI) {
		assert.Equal(t, "MongoDB Atlas: 400 A cluster named test is already present in group null.", errAPI.Error())
	}
}

func TestRelevantError(t *testing.T) {
	cases := []struct {
		httpError error
		apiError  APIError
		expected  error
	}{
		{nil, APIError{}, nil},
		{nil, errAPI, errAPI},
		{errHTTP, APIError{}, errHTTP},
		{errHTTP, errAPI, errHTTP},
	}
	for _, c := range cases {
		err := relevantError(c.httpError, c.apiError)
		assert.Equal(t, c.expected, err)
	}
}
