package auth

import (
	"testing"
	"time"
)

var now = time.Now().UTC()

func TestIsValid(t *testing.T) {
	t.Parallel()
	for _, fixture := range []struct {
		name            string
		validUntil      time.Time
		expectedIsValid bool
	}{
		{
			name:            "now before valid until with buffer",
			validUntil:      now.Add(validityBuffer + time.Hour),
			expectedIsValid: true,
		},
		{
			name:            "now before valid until, but within buffer",
			validUntil:      now.Add(validityBuffer - 1*time.Second),
			expectedIsValid: false,
		},
		{
			name:            "now after valid until",
			validUntil:      now.Add(-1 * time.Second),
			expectedIsValid: false,
		},
		{
			name:            "now after valid until plus validity buffer",
			validUntil:      now.Add(-(validityBuffer + time.Second)),
			expectedIsValid: false,
		},
	} {
		t.Run(fixture.name, func(t *testing.T) {
			t.Parallel()
			jwt := JWT{
				ValidUntil: fixture.validUntil,
			}
			if jwt.IsValid() != fixture.expectedIsValid {
				t.Errorf("expected IsValid() to be %v when `now` is %v and valid-until is %v", fixture.expectedIsValid, now, jwt.ValidUntil)
			}
		})
	}

}
