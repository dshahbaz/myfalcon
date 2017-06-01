package verification

import (
	"fmt"
	"string_hid"

	pb "firmware/proto"
)

const (
	// Individual keys must be exactly of length 1.
	required_key_length = 1

	// Length in a program set can be a maximum of 6 keys.
	max_program_set_key_length = 6

	// Max program sets in a program.
	max_program_set_length = 100

	// Min milliseconds between keys.
	min_program_set_ms = 0

	// Max milliseconds between keys.
	max_program_set_ms = 30
)

// VerifyButtonBinding checks binding for validity (in-range values).
func VerifyButtonBinding(binding *pb.ButtonBinding) error {
	if binding == nil {
		return fmt.Errorf("missing definition; all buttons are required")
	}

	// binding must be one of key, string or program.
	switch binding.Binding.(type) {
	case *pb.ButtonBinding_Key:
		if _, ok := pb.HIDKeyboardKey_name[int32(binding.GetKey())]; !ok {
			return fmt.Errorf(
				"'key' must be one of valid HID keys; not valid: %v", binding)
		}
	case *pb.ButtonBinding_String_:
		s := binding.GetString_()

		if len(s) != required_key_length {
			return fmt.Errorf(
				"button binding %v has string length != %v", binding,
				required_key_length)
		}

		// Check if string has a valid HIDKeyboardKey.
		if _, err := string_hid.StringToHID(s[0]); err != nil {
			return err
		}
	case *pb.ButtonBinding_Program:
		program := binding.GetProgram()

		pss := program.GetProgramSet()

		// Check program set length.
		if len(pss) > max_program_set_length {
			return fmt.Errorf("program length exceeded: %v", binding)
		}

		for _, ps := range pss {
			if _, ok := pb.Modifiers_name[int32(ps.GetModifier())]; !ok {
				return fmt.Errorf("unrecognized modifier: %v", ps)
			}

			if ps.GetMillisecondsBetweenKeys() < min_program_set_ms ||
				ps.GetMillisecondsBetweenKeys() > max_program_set_ms {
				return fmt.Errorf(
					"milliseconds_between_keys out of range, must be >= 0, <= 30: %v", ps)
			}

			if len(ps.GetKeys()) > max_program_set_key_length {
				return fmt.Errorf("program set keys length exceeded: %v", ps)
			}
		}
	case nil:
		// Neither GetKey(), GetString_(), nor GetProgram() were defined; this is an
		// error.
		return fmt.Errorf("need either a key or a program for this button")
	default:
		return fmt.Errorf("need either a key or a program for this button")
	}

	return nil
}

// VerifyButtonBindings checks all bindings for validity (in-range values).
func VerifyButtonBindings(bindings *pb.ButtonBindings) error {
	b := bindings.GetButton1()
	err := VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button1: %v", err)
	}

	b = bindings.GetButton2()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button2: %v", err)
	}

	b = bindings.GetButton3()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button3: %v", err)
	}

	b = bindings.GetButton4()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button4: %v", err)
	}

	b = bindings.GetButton5()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button5: %v", err)
	}

	b = bindings.GetButton6()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button6: %v", err)
	}

	b = bindings.GetButton7()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button7: %v", err)
	}

	b = bindings.GetButton8()
	err = VerifyButtonBinding(b)
	if err != nil {
		return fmt.Errorf("button8: %v", err)
	}

	return nil
}
