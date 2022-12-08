package gowinkey

// VirtualKey represents a virtual key.
type VirtualKey int

const (
	VK_LBUTTON             VirtualKey = 0x01
	VK_RBUTTON                        = 0x02
	VK_CANCEL                         = 0x03
	VK_MBUTTON                        = 0x04
	VK_XBUTTON1                       = 0x05
	VK_XBUTTON2                       = 0x06
	VK_BACK                           = 0x08
	VK_TAB                            = 0x09
	VK_CLEAR                          = 0x0C
	VK_RETURN                         = 0x0D
	VK_SHIFT                          = 0x10
	VK_CONTROL                        = 0x11
	VK_MENU                           = 0x12
	VK_PAUSE                          = 0x13
	VK_CAPITAL                        = 0x14
	VK_KANA                           = 0x15
	VK_HANGUEL                        = 0x15
	VK_HANGUL                         = 0x15
	VK_JUNJA                          = 0x17
	VK_FINAL                          = 0x18
	VK_HANJA                          = 0x19
	VK_KANJI                          = 0x19
	VK_ESCAPE                         = 0x1B
	VK_CONVERT                        = 0x1C
	VK_NONCONVERT                     = 0x1D
	VK_ACCEPT                         = 0x1E
	VK_MODECHANGE                     = 0x1F
	VK_SPACE                          = 0x20
	VK_PRIOR                          = 0x21
	VK_NEXT                           = 0x22
	VK_END                            = 0x23
	VK_HOME                           = 0x24
	VK_LEFT                           = 0x25
	VK_UP                             = 0x26
	VK_RIGHT                          = 0x27
	VK_DOWN                           = 0x28
	VK_SELECT                         = 0x29
	VK_PRINT                          = 0x2A
	VK_EXECUTE                        = 0x2B
	VK_SNAPSHOT                       = 0x2C
	VK_INSERT                         = 0x2D
	VK_DELETE                         = 0x2E
	VK_HELP                           = 0x2F
	VK_0                              = 0x30
	VK_1                              = 0x31
	VK_2                              = 0x32
	VK_3                              = 0x33
	VK_4                              = 0x34
	VK_5                              = 0x35
	VK_6                              = 0x36
	VK_7                              = 0x37
	VK_8                              = 0x38
	VK_9                              = 0x39
	VK_A                              = 0x41
	VK_B                              = 0x42
	VK_C                              = 0x43
	VK_D                              = 0x44
	VK_E                              = 0x45
	VK_F                              = 0x46
	VK_G                              = 0x47
	VK_H                              = 0x48
	VK_I                              = 0x49
	VK_J                              = 0x4A
	VK_K                              = 0x4B
	VK_L                              = 0x4C
	VK_M                              = 0x4D
	VK_N                              = 0x4E
	VK_O                              = 0x4F
	VK_P                              = 0x50
	VK_Q                              = 0x51
	VK_R                              = 0x52
	VK_S                              = 0x53
	VK_T                              = 0x54
	VK_U                              = 0x55
	VK_V                              = 0x56
	VK_W                              = 0x57
	VK_X                              = 0x58
	VK_Y                              = 0x59
	VK_Z                              = 0x5A
	VK_LWIN                           = 0x5B
	VK_RWIN                           = 0x5C
	VK_APPS                           = 0x5D
	VK_SLEEP                          = 0x5F
	VK_NUMPAD0                        = 0x60
	VK_NUMPAD1                        = 0x61
	VK_NUMPAD2                        = 0x62
	VK_NUMPAD3                        = 0x63
	VK_NUMPAD4                        = 0x64
	VK_NUMPAD5                        = 0x65
	VK_NUMPAD6                        = 0x66
	VK_NUMPAD7                        = 0x67
	VK_NUMPAD8                        = 0x68
	VK_NUMPAD9                        = 0x69
	VK_MULTIPLY                       = 0x6A
	VK_ADD                            = 0x6B
	VK_SEPARATOR                      = 0x6C
	VK_SUBTRACT                       = 0x6D
	VK_DECIMAL                        = 0x6E
	VK_DIVIDE                         = 0x6F
	VK_F1                             = 0x70
	VK_F2                             = 0x71
	VK_F3                             = 0x72
	VK_F4                             = 0x73
	VK_F5                             = 0x74
	VK_F6                             = 0x75
	VK_F7                             = 0x76
	VK_F8                             = 0x77
	VK_F9                             = 0x78
	VK_F10                            = 0x79
	VK_F11                            = 0x7A
	VK_F12                            = 0x7B
	VK_F13                            = 0x7C
	VK_F14                            = 0x7D
	VK_F15                            = 0x7E
	VK_F16                            = 0x7F
	VK_F17                            = 0x80
	VK_F18                            = 0x81
	VK_F19                            = 0x82
	VK_F20                            = 0x83
	VK_F21                            = 0x84
	VK_F22                            = 0x85
	VK_F23                            = 0x86
	VK_F24                            = 0x87
	VK_NUMLOCK                        = 0x90
	VK_SCROLL                         = 0x91
	VK_LSHIFT                         = 0xA0
	VK_RSHIFT                         = 0xA1
	VK_LCONTROL                       = 0xA2
	VK_RCONTROL                       = 0xA3
	VK_LMENU                          = 0xA4
	VK_RMENU                          = 0xA5
	VK_BROWSER_BACK                   = 0xA6
	VK_BROWSER_FORWARD                = 0xA7
	VK_BROWSER_REFRESH                = 0xA8
	VK_BROWSER_STOP                   = 0xA9
	VK_BROWSER_SEARCH                 = 0xAA
	VK_BROWSER_FAVORITES              = 0xAB
	VK_BROWSER_HOME                   = 0xAC
	VK_VOLUME_MUTE                    = 0xAD
	VK_VOLUME_DOWN                    = 0xAE
	VK_VOLUME_UP                      = 0xAF
	VK_MEDIA_NEXT_TRACK               = 0xB0
	VK_MEDIA_PREV_TRACK               = 0xB1
	VK_MEDIA_STOP                     = 0xB2
	VK_MEDIA_PLAY_PAUSE               = 0xB3
	VK_LAUNCH_MAIL                    = 0xB4
	VK_LAUNCH_MEDIA_SELECT            = 0xB5
	VK_LAUNCH_APP1                    = 0xB6
	VK_LAUNCH_APP2                    = 0xB7
	VK_OEM_1                          = 0xBA
	VK_OEM_PLUS                       = 0xBB
	VK_OEM_COMMA                      = 0xBC
	VK_OEM_MINUS                      = 0xBD
	VK_OEM_PERIOD                     = 0xBE
	VK_OEM_2                          = 0xBF
	VK_OEM_3                          = 0xC0
	VK_OEM_4                          = 0xDB
	VK_OEM_5                          = 0xDC
	VK_OEM_6                          = 0xDD
	VK_OEM_7                          = 0xDE
	VK_OEM_8                          = 0xDF
	VK_OEM_102                        = 0xE2
	VK_PROCESSKEY                     = 0xE5
	VK_PACKET                         = 0xE7
	VK_ATTN                           = 0xF6
	VK_CRSEL                          = 0xF7
	VK_EXSEL                          = 0xF8
	VK_EREOF                          = 0xF9
	VK_PLAY                           = 0xFA
	VK_ZOOM                           = 0xFB
	VK_NONAME                         = 0xFC
	VK_PA1                            = 0xFD
	VK_OEM_CLEAR                      = 0xFE
)

var keyToStr = map[VirtualKey]string{
	VK_LBUTTON:             "left mouse",
	VK_RBUTTON:             "right mouse",
	VK_CANCEL:              "control-break",
	VK_MBUTTON:             "middle mouse",
	VK_XBUTTON1:            "X1 mouse (Windows 2000)",
	VK_XBUTTON2:            "X2 mouse (Windows 2000)",
	VK_BACK:                "BACKSPACE",
	VK_TAB:                 "TAB",
	VK_CLEAR:               "CLEAR",
	VK_RETURN:              "ENTER",
	VK_SHIFT:               "SHIFT",
	VK_CONTROL:             "CTRL",
	VK_MENU:                "ALT",
	VK_PAUSE:               "PAUSE",
	VK_CAPITAL:             "CAPS LOCK",
	VK_ESCAPE:              "ESC",
	VK_SPACE:               "SPACE",
	VK_PRIOR:               "PAGE UP",
	VK_NEXT:                "PAGE DOWN",
	VK_END:                 "END",
	VK_HOME:                "HOME",
	VK_LEFT:                "LEFT ARROW",
	VK_UP:                  "UP ARROW",
	VK_RIGHT:               "RIGHT ARROW",
	VK_DOWN:                "DOWN ARROW",
	VK_SELECT:              "SELECT",
	VK_PRINT:               "PRINT",
	VK_EXECUTE:             "EXECUTE",
	VK_SNAPSHOT:            "PRINT SCREEN",
	VK_INSERT:              "INS",
	VK_DELETE:              "DEL",
	VK_HELP:                "HELP",
	VK_0:                   "0",
	VK_1:                   "1",
	VK_2:                   "2",
	VK_3:                   "3",
	VK_4:                   "4",
	VK_5:                   "5",
	VK_6:                   "6",
	VK_7:                   "7",
	VK_8:                   "8",
	VK_9:                   "9",
	VK_A:                   "A",
	VK_B:                   "B",
	VK_C:                   "C",
	VK_D:                   "D",
	VK_E:                   "E",
	VK_F:                   "F",
	VK_G:                   "G",
	VK_H:                   "H",
	VK_I:                   "I",
	VK_J:                   "J",
	VK_K:                   "K",
	VK_L:                   "L",
	VK_M:                   "M",
	VK_N:                   "N",
	VK_O:                   "O",
	VK_P:                   "P",
	VK_Q:                   "Q",
	VK_R:                   "R",
	VK_S:                   "S",
	VK_T:                   "T",
	VK_U:                   "U",
	VK_V:                   "V",
	VK_W:                   "W",
	VK_X:                   "X",
	VK_Y:                   "Y",
	VK_Z:                   "Z",
	VK_LWIN:                "Left Windows",
	VK_RWIN:                "Right Windows",
	VK_APPS:                "Applications",
	VK_SLEEP:               "Computer Sleep",
	VK_NUMPAD0:             "Numpad 0",
	VK_NUMPAD1:             "Numpad 1",
	VK_NUMPAD2:             "Numpad 2",
	VK_NUMPAD3:             "Numpad 3",
	VK_NUMPAD4:             "Numpad 4",
	VK_NUMPAD5:             "Numpad 5",
	VK_NUMPAD6:             "Numpad 6",
	VK_NUMPAD7:             "Numpad 7",
	VK_NUMPAD8:             "Numpad 8",
	VK_NUMPAD9:             "Numpad 9",
	VK_MULTIPLY:            "Multiply",
	VK_ADD:                 "Add",
	VK_SEPARATOR:           "Separator",
	VK_SUBTRACT:            "Subtract",
	VK_DECIMAL:             "Decimal",
	VK_DIVIDE:              "Divide",
	VK_F1:                  "F1",
	VK_F2:                  "F2",
	VK_F3:                  "F3",
	VK_F4:                  "F4",
	VK_F5:                  "F5",
	VK_F6:                  "F6",
	VK_F7:                  "F7",
	VK_F8:                  "F8",
	VK_F9:                  "F9",
	VK_F10:                 "F11",
	VK_F11:                 "F12",
	VK_F12:                 "F13",
	VK_F13:                 "F14",
	VK_F14:                 "F15",
	VK_F15:                 "F16",
	VK_F16:                 "F17",
	VK_F17:                 "F18",
	VK_F18:                 "F19",
	VK_F19:                 "F20",
	VK_F20:                 "F21",
	VK_F21:                 "F22",
	VK_F22:                 "F23",
	VK_F23:                 "F24",
	VK_F24:                 "F25",
	VK_NUMLOCK:             "NUM LOCK",
	VK_SCROLL:              "SCROLL LOCK",
	VK_LSHIFT:              "Left SHIFT",
	VK_RSHIFT:              "Right SHIFT",
	VK_LCONTROL:            "Left CONTROL",
	VK_RCONTROL:            "Right CONTROL",
	VK_LMENU:               "Left MENU",
	VK_RMENU:               "Right MENU",
	VK_BROWSER_BACK:        "Windows 2000/XP/2003/Vista/2008: Browser Back",
	VK_BROWSER_FORWARD:     "Windows 2000/XP/2003/Vista/2008: Browser Forward",
	VK_BROWSER_REFRESH:     "Windows 2000/XP/2003/Vista/2008: Browser Refresh",
	VK_BROWSER_STOP:        "Windows 2000/XP/2003/Vista/2008: Browser Stop",
	VK_BROWSER_SEARCH:      "Windows 2000/XP/2003/Vista/2008: Browser Search",
	VK_BROWSER_FAVORITES:   "Windows 2000/XP/2003/Vista/2008: Browser Favorites",
	VK_BROWSER_HOME:        "Windows 2000/XP/2003/Vista/2008: Browser Start and Home",
	VK_VOLUME_MUTE:         "Windows 2000/XP/2003/Vista/2008: Volume Mute",
	VK_VOLUME_DOWN:         "Windows 2000/XP/2003/Vista/2008: Volume Down",
	VK_VOLUME_UP:           "Windows 2000/XP/2003/Vista/2008: Volume Up",
	VK_MEDIA_NEXT_TRACK:    "Windows 2000/XP/2003/Vista/2008: Next Track",
	VK_MEDIA_PREV_TRACK:    "Windows 2000/XP/2003/Vista/2008: Previous Track",
	VK_MEDIA_STOP:          "Windows 2000/XP/2003/Vista/2008: Stop Media",
	VK_MEDIA_PLAY_PAUSE:    "Windows 2000/XP/2003/Vista/2008: Play/Pause Media",
	VK_LAUNCH_MAIL:         "Windows 2000/XP/2003/Vista/2008: Start Mail",
	VK_LAUNCH_MEDIA_SELECT: "Windows 2000/XP/2003/Vista/2008: Select Media",
	VK_LAUNCH_APP1:         "Windows 2000/XP/2003/Vista/2008: Start Application 1",
	VK_LAUNCH_APP2:         "Windows 2000/XP/2003/Vista/2008: Start Application 2",
	VK_OEM_1:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: ';:'",
	VK_OEM_PLUS:            "Windows 2000/XP/2003/Vista/2008: '+'",
	VK_OEM_COMMA:           "Windows 2000/XP/2003/Vista/2008: ','",
	VK_OEM_MINUS:           "Windows 2000/XP/2003/Vista/2008: '-'",
	VK_OEM_PERIOD:          "Windows 2000/XP/2003/Vista/2008: '.'",
	VK_OEM_2:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: '/?'",
	VK_OEM_3:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: '`~'",
	VK_OEM_4:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: '[{'",
	VK_OEM_5:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: '\\|'",
	VK_OEM_6:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: ']}'",
	VK_OEM_7:               "Windows 2000/XP/2003/Vista/2008: US standard keyboard: 'single-quote/double-quote'",
	VK_OEM_8:               "miscellaneous character, may depend on keyboard",
	VK_OEM_102:             "Windows 2000/XP/2003/Vista/2008: either 'angle bracket' or 'backslash'",
	VK_PROCESSKEY:          "Windows 95/98/Me, Windows NT/2000/XP/2003/Vista/2008: IME PROCESS",
	VK_PACKET:              "Windows 2000/XP/2003/Vista/2008: used to pass Unicode characters as if they were keystrokes",
	VK_ATTN:                "attn",
	VK_CRSEL:               "crsel",
	VK_EXSEL:               "exsel",
	VK_EREOF:               "erase eof",
	VK_PLAY:                "play",
	VK_ZOOM:                "zoom",
	VK_NONAME:              "reserved for future use",
	VK_PA1:                 "pa1",
	VK_OEM_CLEAR:           "clear",
}

// String returns the string representation of vk, according
// to a predefined table. If that table does not contain such
// a representation, String falls back to the "MapVirtualKeyA"
// function exposed by Windows' user32.dll.
func (vk VirtualKey) String() string {
	s := keyToStr[vk]
	if s != "" {
		return s
	}
	return virtualKeyToString(vk)
}
