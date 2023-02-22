package main

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

//sys EnumDisplayDevices(lpDevice *uint16, iDevNum uint32, lpDisplayDevice *DISPLAY_DEVICE, dwFlags uint32) (numChars uint32, err error) = user32.EnumDisplayDevicesW
//sys EnumDisplaySettings(lpszDeviceName *uint16, iModeNum int32, lpDevMode *DEVMODEW) (numChars uint32, err error) = user32.EnumDisplaySettingsW
//sys ChangeDisplaySettingsEx(lpszDeviceName *uint16, lpDevMode *DEVMODEW, hWnd windows.HWND, dwFlags uint32, lParam uintptr) (dispChange int) = user32.ChangeDisplaySettingsExW
//sys GetForegroundWindow() (hWnd windows.HWND) = user32.GetForegroundWindow
//sys GetTopWindow(hWnd windows.HWND) (hWndTarget windows.HWND) = user32.GetTopWindow
//sys GetWindow(hWnd windows.HWND, cmd uint) (hWndTarget windows.HWND) = user32.GetWindow
//sys GetWindowText(hWnd windows.HWND, lpString *uint16 , maxCount int) (length int) = user32.GetWindowTextW
//sys GetWindowTextLength(hWnd windows.HWND) (length int) = user32.GetWindowTextLengthW
//sys GetWindowRect(hWnd windows.HWND, rect *RECT) (numChars uint32, err error) = user32.GetWindowRect
//sys DwmGetWindowAttribute(hWnd windows.HWND, dwAttribute uint32, pvAttribute unsafe.Pointer, cbAttribute uint32) (hResult int) = dwmapi.DwmGetWindowAttribute
//sys IsWindowVisible(hWnd windows.HWND) (visible bool) = user32.IsWindowVisible
//sys IsIconic(hWnd windows.HWND) (iconic bool) = user32.IsIconic
//sys GetShellWindow() (hWnd windows.HWND) = user32.GetShellWindow
//sys BeginDeferWindowPos(numWindows int) (winPosInfo windows.HWND) = user32.BeginDeferWindowPos
//sys DeferWindowPos(hdwp windows.HWND, hWnd windows.HWND, hWndInsertAfter windows.HWND, x int, y int, cx int, cy int, flags uint32) (winPosInfo windows.HWND) = user32.DeferWindowPos
//sys EndDeferWindowPos(hdwp windows.HWND) (result bool) = user32.EndDeferWindowPos

type DEVMODEW struct {
	// [System.Runtime.InteropServices.FieldOffset(0)]
	// public string dmDeviceName;
	DeviceName [32]uint16 `json:"device_name"`
	// [System.Runtime.InteropServices.FieldOffset(32)]
	// public Int16 dmSpecVersion;
	SpecVersion uint16 `json:"spec_version"`
	// [System.Runtime.InteropServices.FieldOffset(34)]
	// public Int16 dmDriverVersion;
	DriverVersion uint16 `json:"driver_version"`
	// [System.Runtime.InteropServices.FieldOffset(36)]
	// public Int16 dmSize;
	Size uint16 `json:"size"`
	// [System.Runtime.InteropServices.FieldOffset(38)]
	// public Int16 dmDriverExtra;
	DriverExtra uint16 `json:"driver_extra"`
	// [System.Runtime.InteropServices.FieldOffset(40)]
	// public DM dmFields;
	Fields uint32 `json:"fields"`

	// [System.Runtime.InteropServices.FieldOffset(44)]
	// Int16 dmOrientation;
	// dmOrientation int16
	// [System.Runtime.InteropServices.FieldOffset(46)]
	// Int16 dmPaperSize;
	// dmPaperSize int16
	// [System.Runtime.InteropServices.FieldOffset(48)]
	// Int16 dmPaperLength;
	// dmPaperLength int16
	// [System.Runtime.InteropServices.FieldOffset(50)]
	// Int16 dmPaperWidth;
	// dmPaperWidth int16
	// [System.Runtime.InteropServices.FieldOffset(52)]
	// Int16 dmScale;
	// dmScale int16
	// [System.Runtime.InteropServices.FieldOffset(54)]
	// Int16 dmCopies;
	// dmCopies int16
	// [System.Runtime.InteropServices.FieldOffset(56)]
	// Int16 dmDefaultSource;
	// dmDefaultSource int16
	// [System.Runtime.InteropServices.FieldOffset(58)]
	// Int16 dmPrintQuality;
	// dmPrintQuality int16

	// [System.Runtime.InteropServices.FieldOffset(44)]
	// public POINTL dmPosition;
	Position struct {
		X int32 `json:"x"`
		Y int32 `json:"y"`
	} `json:"position"`

	// [System.Runtime.InteropServices.FieldOffset(52)]
	// public Int32 dmDisplayOrientation;
	DisplayOrientation int32 `json:"display_orientation"`
	// [System.Runtime.InteropServices.FieldOffset(56)]
	// public Int32 dmDisplayFixedOutput;
	DisplayFixedOutput int32 `json:"display_fixed_output"`

	// [System.Runtime.InteropServices.FieldOffset(60)]
	// public short dmColor; // See note below!
	Color int16 `json:"color"`
	// [System.Runtime.InteropServices.FieldOffset(62)]
	// public short dmDuplex; // See note below!
	Duplex int16 `json:"duplex"`
	// [System.Runtime.InteropServices.FieldOffset(64)]
	// public short dmYResolution;
	YResolution int16 `json:"y_resolution"`
	// [System.Runtime.InteropServices.FieldOffset(66)]
	// public short dmTTOption;
	TTOption int16 `json:"tt_option"`
	// [System.Runtime.InteropServices.FieldOffset(68)]
	// public short dmCollate; // See note below!
	Collate int16 `json:"collate"`
	// [System.Runtime.InteropServices.FieldOffset(70)]
	// [MarshalAs(UnmanagedType.ByValTStr, SizeConst = CCHFORMNAME)]
	// public string dmFormName;
	FormName [32]uint16 `json:"form_name"`
	// [System.Runtime.InteropServices.FieldOffset(102)]
	// public Int16 dmLogPixels;
	LogPixels uint16 `json:"log_pixels"`
	// [System.Runtime.InteropServices.FieldOffset(104)]
	// public Int32 dmBitsPerPel;
	BitsPerPel uint32 `json:"bits_per_pel"`
	// [System.Runtime.InteropServices.FieldOffset(108)]
	// public Int32 dmPelsWidth;
	PelsWidth uint32 `json:"pels_width"`
	// [System.Runtime.InteropServices.FieldOffset(112)]
	// public Int32 dmPelsHeight;
	PelsHeight uint32 `json:"pels_height"`
	// [System.Runtime.InteropServices.FieldOffset(116)]
	// public Int32 dmDisplayFlags;
	DisplayFlags uint32 `json:"display_flags"`
	// [System.Runtime.InteropServices.FieldOffset(116)]
	// public Int32 dmNup;
	// dmNup int32
	// [System.Runtime.InteropServices.FieldOffset(120)]
	// public Int32 dmDisplayFrequency;
	DisplayFrequency int32 `json:"display_frequency"`

	// dmICMMethod     uint32
	// dmICMIntent     uint32
	// dmMediaType     uint32
	// dmDitherType    uint32
	// dmReserved1     uint32
	// dmReserved2     uint32
	// dmPanningWidth  uint32
	// dmPanningHeight uint32
}

type DISPLAY_DEVICE struct {
	Cb           uint32      `json:"cb"`
	DeviceName   [32]uint16  `json:"device_name"`
	DeviceString [128]uint16 `json:"device_string"`
	StateFlags   int32       `json:"state_flags"`
	DeviceID     [128]uint16 `json:"device_id"`
	DeviceKey    [128]uint16 `json:"device_key"`
}

const ENUM_CURRENT_SETTINGS = -1
const ENUM_REGISTRY_SETTINGS = -2

// DisplayDeviceStateFlags
// AttachedToDesktop = 0x1,
const DISPLAY_DEVICE_ATTACHED_TO_DESKTOP = 1

// MultiDriver = 0x2,
// /// <summary>The device is part of the desktop.</summary>
// PrimaryDevice = 0x4,
const DISPLAY_DEVICE_PRIMARY_DEVICE = 4

// /// <summary>Represents a pseudo device used to mirror application drawing for remoting or other purposes.</summary>
// MirroringDriver = 0x8,
// /// <summary>The device is VGA compatible.</summary>
// VGACompatible = 0x10,
// /// <summary>The device is removable; it cannot be the primary display.</summary>
// Removable = 0x20,
// /// <summary>The device has more display modes than its output devices support.</summary>
// ModesPruned = 0x8000000,
// Remote = 0x4000000,
// Disconnect = 0x2000000

// [Flags()]
// public enum ChangeDisplaySettingsFlags : uint
// {
//     CDS_NONE = 0,
//     CDS_UPDATEREGISTRY = 0x00000001,
const CDS_UPDATEREGISTRY = 0x00000001

//     CDS_TEST = 0x00000002,
//     CDS_FULLSCREEN = 0x00000004,
//     CDS_GLOBAL = 0x00000008,
//     CDS_SET_PRIMARY = 0x00000010,
//     CDS_VIDEOPARAMETERS = 0x00000020,
//     CDS_ENABLE_UNSAFE_MODES = 0x00000100,
//     CDS_DISABLE_UNSAFE_MODES = 0x00000200,
//     CDS_RESET = 0x40000000,
//     CDS_RESET_EX = 0x20000000,
//     CDS_NORESET = 0x10000000
// }

const GW_HWNDNEXT = 2

type RECT struct {
	Left   int32 `json:"left"`
	Top    int32 `json:"top"`
	Right  int32 `json:"right"`
	Bottom int32 `json:"bottom"`
}

const (
	DWMWA_NCRENDERING_ENABLED = iota + 1
	DWMWA_NCRENDERING_POLICY
	DWMWA_TRANSITIONS_FORCEDISABLED
	DWMWA_ALLOW_NCPAINT
	DWMWA_CAPTION_BUTTON_BOUNDS
	DWMWA_NONCLIENT_RTL_LAYOUT
	DWMWA_FORCE_ICONIC_REPRESENTATION
	DWMWA_FLIP3D_POLICY
	DWMWA_EXTENDED_FRAME_BOUNDS
	DWMWA_HAS_ICONIC_BITMAP
	DWMWA_DISALLOW_PEEK
	DWMWA_EXCLUDED_FROM_PEEK
	DWMWA_CLOAK
	DWMWA_CLOAKED
	DWMWA_FREEZE_REPRESENTATION
	DWMWA_PASSIVE_UPDATE_MODE
	DWMWA_USE_HOSTBACKDROPBRUSH
)
const DWMWA_USE_IMMERSIVE_DARK_MODE = 20
const (
	DWMWA_WINDOW_CORNER_PREFERENCE = iota + 33
	DWMWA_BORDER_COLOR
	DWMWA_CAPTION_COLOR
	DWMWA_TEXT_COLOR
	DWMWA_VISIBLE_FRAME_BORDER_THICKNESS
	DWMWA_SYSTEMBACKDROP_TYPE
	DWMWA_LAST
)

const (
	HWND_BOTTOM    = 1
	HWND_NOTOPMOST = -2
	HWND_TOP       = 0
	HWND_TOPMOST   = -1
)

const (
	SWP_DRAWFRAME      = 0x0020
	SWP_FRAMECHANGED   = 0x0020
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOACTIVATE     = 0x0010
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOMOVE         = 0x0002
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOREDRAW       = 0x0008
	SWP_NOREPOSITION   = 0x0200
	SWP_NOSENDCHANGING = 0x0400
	SWP_NOSIZE         = 0x0001
	SWP_NOZORDER       = 0x0004
	SWP_SHOWWINDOW     = 0x0040
)
