package internal

import "testing"

func TestSnakeCase(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input    string
		expected string
	}{
		{"ThisIsATest", "this_is_a_test"},
		{"ContactsDeliveredAl7dHow", "contacts_delivered_al7d_how"},
		{"SEGMENT_ID_HASH", "segment_id_hash"},
		{"accountId", "account_id"},
		{"account_id", "account_id"},
		{"foo", "foo"},
		{"FOO", "foo"},
		{"MLFoo", "ml_foo"},
		{"Fish_Paste", "fish_paste"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel()
			result := ChalkpySnakeCase(tc.input)
			if result != tc.expected {
				t.Errorf("chalkpySnakeCase(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}
