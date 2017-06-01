package string_hid

import (
	"fmt"

	pb "firmware/proto"
)

// StringToHID returns the HIDKeyboardKey representing the character b.
// The mapping is not intended to be exhaustive, since some HID defined keys do
// not have a related alphanumeric representation. This is just intended for the
// ones that do to make defining the mappings a little easier for those common
// cases.
func StringToHID(b byte) (pb.HIDKeyboardKey, error) {
	var m = map[byte]int32{
		'a':  pb.HIDKeyboardKey_value["KEY_a_A"],
		'b':  pb.HIDKeyboardKey_value["KEY_b_B"],
		'c':  pb.HIDKeyboardKey_value["KEY_c_C"],
		'd':  pb.HIDKeyboardKey_value["KEY_d_D"],
		'e':  pb.HIDKeyboardKey_value["KEY_e_E"],
		'f':  pb.HIDKeyboardKey_value["KEY_f_F"],
		'g':  pb.HIDKeyboardKey_value["KEY_g_G"],
		'h':  pb.HIDKeyboardKey_value["KEY_h_H"],
		'i':  pb.HIDKeyboardKey_value["KEY_i_I"],
		'j':  pb.HIDKeyboardKey_value["KEY_j_J"],
		'k':  pb.HIDKeyboardKey_value["KEY_k_K"],
		'l':  pb.HIDKeyboardKey_value["KEY_l_L"],
		'm':  pb.HIDKeyboardKey_value["KEY_m_M"],
		'n':  pb.HIDKeyboardKey_value["KEY_n_N"],
		'o':  pb.HIDKeyboardKey_value["KEY_o_O"],
		'p':  pb.HIDKeyboardKey_value["KEY_p_P"],
		'q':  pb.HIDKeyboardKey_value["KEY_q_Q"],
		'r':  pb.HIDKeyboardKey_value["KEY_r_R"],
		's':  pb.HIDKeyboardKey_value["KEY_s_S"],
		't':  pb.HIDKeyboardKey_value["KEY_t_T"],
		'u':  pb.HIDKeyboardKey_value["KEY_u_U"],
		'v':  pb.HIDKeyboardKey_value["KEY_v_V"],
		'w':  pb.HIDKeyboardKey_value["KEY_w_W"],
		'x':  pb.HIDKeyboardKey_value["KEY_x_X"],
		'y':  pb.HIDKeyboardKey_value["KEY_y_Y"],
		'z':  pb.HIDKeyboardKey_value["KEY_z_Z"],
		'1':  pb.HIDKeyboardKey_value["KEY_1_Exclamation"],
		'2':  pb.HIDKeyboardKey_value["KEY_2_At"],
		'3':  pb.HIDKeyboardKey_value["KEY_3_Pound"],
		'4':  pb.HIDKeyboardKey_value["KEY_4_Dollar"],
		'5':  pb.HIDKeyboardKey_value["KEY_5_Percent"],
		'6':  pb.HIDKeyboardKey_value["KEY_6_Caret"],
		'7':  pb.HIDKeyboardKey_value["KEY_7_Ampersand"],
		'8':  pb.HIDKeyboardKey_value["KEY_8_Asterisk"],
		'9':  pb.HIDKeyboardKey_value["KEY_9_LeftParenthesis"],
		'0':  pb.HIDKeyboardKey_value["KEY_0_RightParenthesis"],
		'-':  pb.HIDKeyboardKey_value["KEY_Dash_Underscore"],
		'=':  pb.HIDKeyboardKey_value["KEY_Equal_Plus"],
		'[':  pb.HIDKeyboardKey_value["KEY_LeftBracket_LeftBrace"],
		']':  pb.HIDKeyboardKey_value["KEY_RightBracket_RightBrace"],
		'\\': pb.HIDKeyboardKey_value["KEY_Backslash_Pipe"],
		';':  pb.HIDKeyboardKey_value["KEY_Semicolon_Colon"],
		'\'': pb.HIDKeyboardKey_value["KEY_SingleQuote_DoubleQuote"],
		'~':  pb.HIDKeyboardKey_value["KEY_GraveAccent_Tilde"],
		',':  pb.HIDKeyboardKey_value["KEY_Comma_LessThan"],
		'.':  pb.HIDKeyboardKey_value["KEY_Period_GreaterThan"],
		'/':  pb.HIDKeyboardKey_value["KEY_Slash_Question"],
	}
	if v, ok := m[b]; !ok {
		return 0, fmt.Errorf(
			`string did not have a valid HID key: %v; these must be lower-case,
			alphanumeric characters. use a HIDKeyboardKey directly instead?`, b)
	} else {
		return pb.HIDKeyboardKey(v), nil
	}
}
