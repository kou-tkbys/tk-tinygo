package keyboard

type hidReportDescriptor struct {
	f uint8
	s uint8
}

const (
	USAGE_PAGE uint16 = (0x05 << 8) | 0x01
)

// 	//  Keyboard
// 	  0x05, 0x01,                    // USAGE_PAGE (Generic Desktop)  // 47
// 	  0x09, 0x06,                    // USAGE (Keyboard)
// 	  0xa1, 0x01,                    // COLLECTION (Application)
// 	  0x85, 0x02,                    //   REPORT_ID (2)
// 	  0x05, 0x07,                    //   USAGE_PAGE (Keyboard)

// 	0x19, 0xe0,                    //   USAGE_MINIMUM (Keyboard LeftControl)
// 	  0x29, 0xe7,                    //   USAGE_MAXIMUM (Keyboard Right GUI)
// 	  0x15, 0x00,                    //   LOGICAL_MINIMUM (0)
// 	  0x25, 0x01,                    //   LOGICAL_MAXIMUM (1)
// 	  0x75, 0x01,                    //   REPORT_SIZE (1)

// 	0x95, 0x08,                    //   REPORT_COUNT (8)
// 	  0x81, 0x02,                    //   INPUT (Data,Var,Abs)
// 	  0x95, 0x01,                    //   REPORT_COUNT (1)
// 	  0x75, 0x08,                    //   REPORT_SIZE (8)
// 	  0x81, 0x03,                    //   INPUT (Cnst,Var,Abs)

// 	0x95, 0x06,                    //   REPORT_COUNT (6)
// 	  0x75, 0x08,                    //   REPORT_SIZE (8)
// 	  0x15, 0x00,                    //   LOGICAL_MINIMUM (0)
// 	  0x25, 0x73,                    //   LOGICAL_MAXIMUM (115)
// 	  0x05, 0x07,                    //   USAGE_PAGE (Keyboard)

// 	0x19, 0x00,                    //   USAGE_MINIMUM (Reserved (no event indicated))
// 	  // 0x29, 0x73,                    //   USAGE_MAXIMUM (Keyboard Application)
// 	  0x29, TKEY_LEFTCTRL - 1,//   USAGE_MAXIMUM (Keyboard Application)　読み取り範囲を拡張
// 	  0x81, 0x00,                    //   INPUT (Data,Ary,Abs)
// 	  0xc0,                          // END_COLLECTION
//   };
