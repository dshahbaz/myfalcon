package verification_test

import (
	"testing"

	"verification"

	"github.com/golang/protobuf/proto"

	pb "firmware/proto"
)

func TestVerifyButtonBinding(t *testing.T) {
	// Just shorthand to make the table below more clear.
	const ok =  true
	const not_ok = false

	var tests = []struct {
		// Contents of text proto to unmarshal.
		text      string
		// Whether or not parsing should pass or fail.
		expect_ok bool
	}{
		// String of length 1: OK.
		{`string: 'a'`, ok},

		// String of length != 1.
		{`string: 'aa'`, not_ok},

		// Invalid string (technically this is testing string_hid, oh well).
		{"string: '\t'", not_ok},

		// Valid key (one of HIDKeyboardKey; see proto/hid.proto).
		{"key: KEY_k_K", ok},

		// Identical to above.
		{"key: 0x0E", ok},

		// Not a valid HID key.
		{"key: 0xFF", not_ok},

		// Although weird, an empty program set can exist.
		{`program: {
			program_set {
			}
		}`, ok},

		// Although weird, an empty program can exist.
		{`program: { }`, ok},

		// Over 100 program sets.
		{`program: {
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }
			program_set { }  # 101st.
		}`, not_ok},

		// Unrecognized modifier.
		{`program: {
			program_set {
				modifier: 789
			}
		}`, not_ok},

		// Milliseconds out of range.
		{`program: {
			program_set {
				milliseconds_between_keys: 31
			}
		}`, not_ok},
		{`program: {
			program_set {
				milliseconds_between_keys: 30
			}
		}`, ok},
		{`program: {
			program_set {
				milliseconds_between_keys: 5
			}
		}`, ok},
		{`program: {
			program_set {
				milliseconds_between_keys: 0
			}
		}`, ok},

		// Key length.
		{`program: {
			program_set {
				keys: [KEY_a_A]
			}
		}`, ok},
		{`program: {
			program_set {
				keys: [KEY_a_A, KEY_b_B, KEY_c_C, KEY_d_D, KEY_e_E, KEY_f_F]
			}
		}`, ok},
		{`program: {
			program_set {
				keys: [KEY_a_A, KEY_b_B, KEY_c_C, KEY_d_D, KEY_e_E, KEY_f_F, KEY_g_G]
			}
		}`, not_ok},
	}

	var bb pb.ButtonBinding

	for _, test := range tests {
		bb.Reset()

		err := proto.UnmarshalText(test.text, &bb)
		if err != nil {
			t.Errorf("UnmarshalText(%v): %v", test.text, err)
		}

		err = verification.VerifyButtonBinding(&bb)

		if test.expect_ok && err != nil {
			t.Errorf("VerifyButtonBinding(%v) expected ok; was: %v", test.text, err)
		}

		if !test.expect_ok && err == nil {
			t.Errorf("VerifyButtonBinding(%v) expected error; was: %v", test.text, err)
		}
	}
}
