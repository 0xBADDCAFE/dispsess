package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

const PROFILE_NAME = "profile"
const SESSION_NAME = "session"

var flagSave bool

type Display struct {
	DeviceName [32]uint16 `json:"display_Name"`
	DevMode    DEVMODEW   `json:"dev_mode"`
}

type WindowRect struct {
	Text string       `json:"text"`
	HWnd windows.HWND `json:"hwnd"`
	Rect RECT         `json:"rect"`
}

type Session struct {
	Displays []Display    `json:"displays"`
	WRects   []WindowRect `json:"window_rects"`
}

func main() {
	flag.BoolVar(&flagSave, "save", false, "Save current display settings")
	flag.Parse()

	// ddList := []DISPLAY_DEVICE{}
	dispList := []Display{}

	for dispNum := uint32(0); ; dispNum++ {
		dd := &DISPLAY_DEVICE{}
		dd.Cb = uint32(unsafe.Sizeof(*dd))

		result, _ := EnumDisplayDevices(nil, dispNum, dd, 0)
		if result == 0 {
			break
		}

		if dd.isAttached() {
			fmt.Printf("%s (PRIMARY: %v)\n", windows.UTF16PtrToString(&dd.DeviceName[0]), dd.isPrimary())

			dm := &DEVMODEW{}
			dm.Size = uint16(unsafe.Sizeof(*dm))

			_, err := EnumDisplaySettings(&dd.DeviceName[0], ENUM_REGISTRY_SETTINGS, dm)
			if err != nil {
				log.Fatal("EnumDisplaySettings: ", err)
			}

			// log.Println("dm.DeviceName:", windows.UTF16ToString(dm.DeviceName[:]))
			dispList = append(dispList, Display{DeviceName: dd.DeviceName, DevMode: *dm})
			// log.Printf("%#v\n", dmList)
		}
	}

	if flagSave {
		marshallToJsonFile(PROFILE_NAME, dispList)
		fmt.Println("Create profile.json")
	} else {
		if jsonExists(SESSION_NAME) {
			// Restore (Session exists)
			session := &Session{}
			unmarshallJsonFile(SESSION_NAME, session)
			applyDisplaySettings(&session.Displays)
			// Wait to fix display layout
			time.Sleep(time.Second * 5)
			applyWindowRect(&session.WRects)

			err := os.Remove(SESSION_NAME + ".json")
			if err != nil {
				log.Fatal(err)
			}
		} else {
			if !jsonExists(PROFILE_NAME) {
				log.Fatal("Execute with -save before to create profile.json")
			}

			// Store
			session := &Session{Displays: dispList, WRects: getAllWindowRect()}
			marshallToJsonFile(SESSION_NAME, session)
			fmt.Println("Create session.json")

			dispList := &[]Display{}
			unmarshallJsonFile(PROFILE_NAME, dispList)
			applyDisplaySettings(dispList)
		}
	}
}

func jsonExists(jsonPath string) bool {
	if _, err := os.Stat(jsonPath + ".json"); err != nil {
		return false
	}
	return true
}
func unmarshallJsonFile(jsonPath string, data interface{}) {
	jsonBytes, err := ioutil.ReadFile(jsonPath + ".json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		log.Fatal(err)
	}
}
func marshallToJsonFile(jsonPath string, data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(jsonPath+".json", jsonBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func applyDisplaySettings(dispList *[]Display) {
	for _, disp := range *dispList {
		ChangeDisplaySettingsEx(&disp.DeviceName[0], &disp.DevMode, 0, CDS_UPDATEREGISTRY, 0)
	}
}

func (dd *DISPLAY_DEVICE) isAttached() bool {
	return dd.StateFlags&DISPLAY_DEVICE_ATTACHED_TO_DESKTOP != 0
}

func (dd *DISPLAY_DEVICE) isPrimary() bool {
	return dd.StateFlags&DISPLAY_DEVICE_PRIMARY_DEVICE != 0
}

func getAllWindowRect() (wRectList []WindowRect) {
	// for hWnd := GetTopWindow(0); hWnd != 0; hWnd = GetWindow(hWnd, GW_HWNDNEXT) {
	for hWnd := GetForegroundWindow(); hWnd != 0; hWnd = GetWindow(hWnd, GW_HWNDNEXT) {
		if !IsWindowVisible(hWnd) {
			continue
		}
		buf := make([]uint16, GetWindowTextLength(hWnd)+1)
		GetWindowText(hWnd, &buf[0], len(buf))
		rect := &RECT{}
		GetWindowRect(hWnd, rect)
		// extendRect := &RECT{}
		// hResult := DwmGetWindowAttribute(hWnd, DWMWA_EXTENDED_FRAME_BOUNDS, unsafe.Pointer(extendRect), uint32(unsafe.Sizeof(*extendRect)))

		wRectList = append(wRectList,
			WindowRect{Text: syscall.UTF16ToString(buf), HWnd: hWnd, Rect: *rect})
		// fmt.Println(hWnd)
		// fmt.Printf("%s:\n%+v\n", syscall.UTF16ToString(buf), rect)
		// fmt.Printf("%s (result: %x):\n  %+v\n", syscall.UTF16ToString(buf), hResult, extendRect)
	}
	return
}

func applyWindowRect(wRectList *[]WindowRect) {
	winPosInfo := BeginDeferWindowPos(len(*wRectList))
	for _, wRect := range *wRectList {
		winPosInfo = DeferWindowPos(winPosInfo, wRect.HWnd, HWND_TOP,
			int(wRect.Rect.Left), int(wRect.Rect.Top), int(wRect.Rect.Right-wRect.Rect.Left), int(wRect.Rect.Bottom-wRect.Rect.Top),
			SWP_NOZORDER&SWP_NOACTIVATE)
		// SWP_NOZORDER&SWP_NOACTIVATE&SWP_FRAMECHANGED)
	}
	if !EndDeferWindowPos(winPosInfo) {
		log.Fatal("Failure to apply window rect")
	}
}
