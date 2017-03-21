package kubernetes

import (
	"fmt"
	"testing"
)

func TestIsInternalAnnotationKey(t *testing.T) {
	testCases := []struct {
		Key      string
		Expected bool
	}{
		{
			Key:      "%",
			Expected: false,
		},
		{
			Key:      "pv.kubernetes.io/bound-by-controller",
			Expected: true,
		},
		{
			Key:      "pv.kubernetes.io/bind-completed",
			Expected: true,
		},
		{
			Key:      "pv.customhost.io/bind-completed",
			Expected: false,
		},
		{
			Key:      "pv.customhost.io/bind-completed",
			Expected: false,
		},
		{
			Key:      "anything",
			Expected: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			isInternal := isInternalAnnotationKey(tc.Key)
			if tc.Expected && !isInternal {
				t.Fatalf("Expected %q to be internal", tc.Key)
			}
			if !tc.Expected && isInternal {
				t.Fatalf("Expected %q not to be internal", tc.Key)
			}
		})
	}
}
