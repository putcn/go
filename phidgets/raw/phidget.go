package raw

// #cgo darwin CFLAGS: -I/Library/Frameworks/Phidget22.framework/Headers
// #cgo darwin LDFLAGS: -framework Phidget22
// #cgo linux CFLAGS: -lphidget22
// #cgo linux LDFLAGS: -lphidget22
// #include "phidget.h"
import "C"

import (
	"errors"
	"fmt"
	"time"
)

const (
	Any   = -1
	False = C.PFALSE
	True  = C.PTRUE

	channelSize = 50
)

type Status int

const (
	Attached    = true
	NotAttached = false
)

type Class int

const (
	NothingClass           = C.PHIDCLASS_NOTHING
	AccelerometerClass     = C.PHIDCLASS_ACCELEROMETER
	AdvancedServoClass     = C.PHIDCLASS_ADVANCEDSERVO
	AnalogClass            = C.PHIDCLASS_ANALOG
	BridgeClass            = C.PHIDCLASS_BRIDGE
	EncoderClass           = C.PHIDCLASS_ENCODER
	FrequencyCounterClass  = C.PHIDCLASS_FREQUENCYCOUNTER
	GPSClass               = C.PHIDCLASS_GPS
	HubClass               = C.PHIDCLASS_HUB
	InterfaceKitClass      = C.PHIDCLASS_INTERFACEKIT
	IRClass                = C.PHIDCLASS_IR
	LEDClass               = C.PHIDCLASS_LED
	MeshDongleClass        = C.PHIDCLASS_MESHDONGLE
	MotorControlClass      = C.PHIDCLASS_MOTORCONTROL
	PHSensorClass          = C.PHIDCLASS_PHSENSOR
	RFIDClass              = C.PHIDCLASS_RFID
	ServoClass             = C.PHIDCLASS_SERVO
	SpatialClass           = C.PHIDCLASS_SPATIAL
	StepperClass           = C.PHIDCLASS_STEPPER
	TemperatureSensorClass = C.PHIDCLASS_TEMPERATURESENSOR
	TextLCDClass           = C.PHIDCLASS_TEXTLCD
	VINTClass              = C.PHIDCLASS_VINT
	GenericClass           = C.PHIDCLASS_GENERIC
	FirmwareUpgradeClass   = C.PHIDCLASS_FIRMWAREUPGRADE
	DictionaryClass        = C.PHIDCLASS_DICTIONARY
)

type ID int

const (
	// Device        = C.PHIDID_INTERFACEKIT_4_8_8 // No idea
	// Device        = C.PHIDID_RFID               // No idea
	// Device        = C.PHIDID_TEXTLED_1x8        // No idea
	// Device        = C.PHIDID_TEXTLED_4x8        // No idea
	// Device1000    = C.PHIDID_SERVO_1MOTOR
	// Device1000Old = C.PHIDID_SERVO_1MOTOR_OLD
	// Device1001    = C.PHIDID_SERVO_4MOTOR
	// Device1001Old = C.PHIDID_SERVO_4MOTOR_OLD
	// Device1002    = C.PHIDID_ANALOG_4OUTPUT
	// Device1011    = C.PHIDID_INTERFACEKIT_2_2_2
	// Device1012    = C.PHIDID_INTERFACEKIT_0_16_16
	// Device1014    = C.PHIDID_INTERFACEKIT_0_0_4
	// Device1015    = C.PHIDID_LINEAR_TOUCH
	// Device1016    = C.PHIDID_ROTARY_TOUCH
	// Device1017    = C.PHIDID_INTERFACEKIT_0_0_8
	// Device1018    = C.PHIDID_INTERFACEKIT_8_8_8
	// Device1023    = C.PHIDID_RFID_2OUTPUT
	// Device1024    = C.PHIDID_RFID_2OUTPUT_READ_WRITE // Not Found
	// Device1030    = C.PHIDID_LED_64
	// Device1031    = C.PHIDID_LED_64_ADV
	// Device1040    = C.PHIDID_GPS
	// Device1043    = C.PHIDID_SPATIAL_ACCEL_3AXIS
	// Device1044    = C.PHIDID_SPATIAL_ACCEL_GYRO_COMPASS
	// Device1045    = C.PHIDID_TEMPERATURESENSOR_IR
	// Device1046    = C.PHIDID_BRIDGE_4INPUT
	// Device1047    = C.PHIDID_ENCODER_HS_4ENCODER_4INPUT
	// Device1048    = C.PHIDID_TEMPERATURESENSOR_4
	// Device1050    = C.PHIDID_WEIGHTSENSOR // Not Found
	// Device1051    = C.PHIDID_TEMPERATURESENSOR
	// Device1052    = C.PHIDID_ENCODER_1ENCODER_1INPUT
	// Device1054    = C.PHIDID_ACCELEROMETER_2AXIS // PhidgetFrequencyCounter
	// Device1054    = C.PHIDID_FREQUENCYCOUNTER_2INPUT
	IRID = C.PHIDID_1055
	// Device1057    = C.PHIDID_ENCODER_HS_1ENCODER
	// Device1058    = C.PHIDID_PHSENSOR
	// Device1059    = C.PHIDID_ACCELEROMETER_3AXIS
	// Device1060    = C.PHIDID_MOTORCONTROL_LV_2MOTOR_4INPUT
	// Device1061    = C.PHIDID_ADVANCEDSERVO_8MOTOR
	// Device1062    = C.PHIDID_UNIPOLAR_STEPPER_4MOTOR
	// Device1063    = C.PHIDID_BIPOLAR_STEPPER_1MOTOR
	// Device1064    = C.PHIDID_MOTORCONTROL_HC_2MOTOR
	// Device1065    = C.PHIDID_MOTORCONTROL_1MOTOR
	// Device1066    = C.PHIDID_ADVANCEDSERVO_1MOTOR
	// Device1203    = C.PHIDID_INTERFACEKIT_8_8_8_w_LCD
	// Device1203    = C.PHIDID_TEXTLCD_2x20_w_8_8_8
	// Device1204    = C.PHIDID_TEXTLCD_ADAPTER
	// Device1210    = C.PHIDID_TEXTLCD_2x20             // Not Found
	// Device1221    = C.PHIDID_INTERFACEKIT_0_8_8_w_LCD // Not Found
	// Device1221    = C.PHIDID_TEXTLCD_2x20_w_0_8_8     // Not Found
)

type eventType int

const (
	phidgetAttach     = C.phidgetAttach
	phidgetConnect    = C.phidgetConnect
	phidgetDetach     = C.phidgetDetach
	phidgetDisconnect = C.phidgetDisconnect
)

type Phidget struct {
	Attached     <-chan bool
	Connected    <-chan bool
	Detached     <-chan bool
	Disconnected <-chan bool
	Error        <-chan error

	attached            chan bool
	connected           chan bool
	detached            chan bool
	disconnected        chan bool
	error               chan error
	handle              *C.PhidgetHandle
	onAttachHandler     *C.handler
	onConnectHandler    *C.handler
	onDetachHandler     *C.handler
	onDisconnectHandler *C.handler
	onErrorHandler      *C.handler
}

func ErrorDescription(code int) string {
	str := new(*C.char)
	C.Phidget_getErrorDescription(C.PhidgetReturnCode(code), str)
	return C.GoString(*str)
}

func LibraryVersion() (string, error) {
	str := new(*C.char)
	err := resultWithReturnCode(C.Phidget_getLibraryVersion(str))
	if err != nil {
		return "", err
	}
	return C.GoString(*str), nil
}

func (p *Phidget) Close() error {
	return resultWithReturnCode(C.Phidget_close(p.handle))
}

func (p *Phidget) GetDeviceClass() (Class, error) {
	ptr := new(C.Phidget_DeviceClass)
	err := resultWithReturnCode(C.Phidget_getDeviceClass(p.handle, ptr))
	if err != nil {
		return 0, err
	}
	return Class(*ptr), nil
}

func (p *Phidget) GetDeviceID() (ID, error) {
	ptr := new(C.Phidget_DeviceID)
	err := resultWithReturnCode(C.Phidget_getDeviceID(p.handle, ptr))
	if err != nil {
		return 0, err
	}
	return ID(*ptr), nil
}

func (p *Phidget) GetDeviceLabel() (string, error) {
	return resultWithString(func(ptr **C.char) C.int { return C.int(C.Phidget_getDeviceLabel(p.handle, ptr)) })
}

func (p *Phidget) GetDeviceName() (string, error) {
	return resultWithString(func(ptr **C.char) C.int { return C.int(C.Phidget_getDeviceName(p.handle, ptr)) })
}

/* Deprecated
func (p *Phidget) GetDeviceStatus() (int, error) {
	return resultWithInt(func(ptr *C.int) C.int { return C.Phidget_getDeviceStatus(p.handle, ptr) })
}

func (p *Phidget) GetDeviceType() (string, error) {
	return resultWithString(func(ptr **C.char) C.int { return C.Phidget_getDeviceType(p.handle, ptr) })
}
*/
func (p *Phidget) GetDeviceVersion() (int, error) {
	return resultWithInt(func(ptr *C.int) C.int { return C.int(C.Phidget_getDeviceVersion(p.handle, ptr)) })
}

func (p *Phidget) GetSerialNumber() (int, error) {
	return resultWithInt(func(ptr *C.int) C.int { return C.int(C.Phidget_getDeviceSerialNumber(p.handle, ptr)) })
}

/*
func (p *Phidget) GetServerAddress() (string, int, error) {
	addr := new(*C.char)
	port := new(C.int)
	err := result(C.Phidget_getServerAddress(p.handle, addr, port))
	if err != nil {
		return "", 0, err
	}
	return C.GoString(*addr), int(*port), nil
}


func (p *Phidget) GetServerID() (string, error) {
	return resultWithString(func(ptr **C.char) C.int { return C.Phidget_getServerID(p.handle, ptr) })
}

func (p *Phidget) GetServerStatus() (int, error) {
	return resultWithInt(func(ptr *C.int) C.int { return C.Phidget_getServerStatus(p.handle, ptr) })
}

*/

func (p *Phidget) Open(serial int) error {
	return resultWithReturnCode(C.Phidget_open(p.handle, C.int(serial)))
}

/*

func (p *Phidget) OpenLabel(label string) error {
	return result(C.Phidget_openLabel(p.handle, convertString(label)))
}

func (p *Phidget) OpenLabelRemote(label, server, password string) error {
	return result(C.Phidget_openLabelRemote(p.handle, convertString(label), convertString(server), convertString(password)))
}

func (p *Phidget) OpenLabelRemoteIP(label string, address string, port int, password string) error {
	return result(C.Phidget_openLabelRemoteIP(p.handle, convertString(label), convertString(address), C.int(port), convertString(password)))
}

func (p *Phidget) OpenRemote(serial int, server, password string) error {
	return result(C.Phidget_openRemote(p.handle, C.int(serial), convertString(server), convertString(password)))
}

func (p *Phidget) OpenRemoteIP(serial int, address string, port int, password string) error {
	return result(C.Phidget_openRemoteIP(p.handle, C.int(serial), convertString(address), C.int(port), convertString(password)))
}
*/

func (p *Phidget) SetDeviceLabel(label string) error {
	return resultWithReturnCode(C.Phidget_setDeviceLabel(p.handle, convertString(label)))
}

func (p *Phidget) WaitForAttachment(timeout time.Duration) error {
	return resultWithReturnCode(C.Phidget_openWaitForAttachment(p.handle, C.uint32_t(timeout/time.Millisecond)))
}

func convertString(str string) *C.char {
	if str == "" {
		return nil
	}
	return C.CString(str)
}

func _result(code int) error {
	if code != 0 {
		return errors.New(fmt.Sprintf("%s (%d)", ErrorDescription(code), code))
	}
	return nil
}

func result(result C.int) error {
	code := int(result)
	return _result(code)
}

func resultWithReturnCode(result C.PhidgetReturnCode) error {
	code := int(result)
	return _result(code)
}

func resultWithInt(f func(*C.int) C.int) (int, error) {
	ptr := new(C.int)
	err := result(f(ptr))
	if err != nil {
		return 0, err
	}
	return int(*ptr), nil
}

func resultWithString(f func(**C.char) C.int) (string, error) {
	ptr := new(*C.char)
	err := result(f(ptr))
	if err != nil {
		return "", err
	}
	return C.GoString(*ptr), nil
}

func (p *Phidget) cleanup() {
	p.unsetOnErrorHandler()
	p.unsetOnEventHandler(phidgetDisconnect, &p.onDisconnectHandler)
	p.unsetOnEventHandler(phidgetDetach, &p.onDetachHandler)
	p.unsetOnEventHandler(phidgetConnect, &p.onConnectHandler)
	p.unsetOnEventHandler(phidgetAttach, &p.onAttachHandler)
	C.Phidget_delete(p.handle)
}

func (p *Phidget) initPhidget(h *C.PhidgetHandle) error {
	var err error

	p.handle = h

	p.attached = make(chan bool, channelSize)
	p.connected = make(chan bool, channelSize)
	p.detached = make(chan bool, channelSize)
	p.disconnected = make(chan bool, channelSize)
	p.error = make(chan error, channelSize)

	p.Attached = p.attached
	p.Connected = p.connected
	p.Detached = p.detached
	p.Disconnected = p.disconnected
	p.Error = p.error

	if p.onAttachHandler, err = p.setOnEventHandler(p.attached, phidgetAttach); err != nil {
		return err
	}

	if p.onConnectHandler, err = p.setOnEventHandler(p.connected, phidgetConnect); err != nil {
		return err
	}

	if p.onDetachHandler, err = p.setOnEventHandler(p.detached, phidgetDetach); err != nil {
		return err
	}

	if p.onDisconnectHandler, err = p.setOnEventHandler(p.disconnected, phidgetDisconnect); err != nil {
		return err
	}

	if err := p.setOnErrorHandler(); err != nil {
		return err
	}

	return nil
}

func (p *Phidget) setOnErrorHandler() error {
	var err error

	p.onErrorHandler, err = createHandler(func(h *C.handler) C.int {
		return C.setOnErrorHandler(p.handle, h)
	})
	if err != nil {
		return err
	}

	go func() {
		for {
			r := C.onErrorAwait(p.onErrorHandler)
			e := errors.New(C.GoString(r.string))

			C.onErrorResultFree(r)

			select {
			case p.error <- e:
			default:
			}
		}
	}()

	return nil
}

func (p *Phidget) setOnEventHandler(c chan bool, t eventType) (*C.handler, error) {
	h, err := createHandler(func(h *C.handler) C.int {
		return C.setOnEventHandler(p.handle, h, C.eventType(t))
	})
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			C.onEventAwait(h)
			select {
			case c <- true:
			default:
			}
		}
	}()

	return h, nil
}

func (p *Phidget) unsetOnErrorHandler() {
	C.unsetOnErrorHandler(p.handle)
	C.handlerFree(p.onErrorHandler)
	p.onErrorHandler = nil
}

func (p *Phidget) unsetOnEventHandler(t eventType, h **C.handler) {
	C.unsetOnEventHandler(p.handle, C.eventType(t))
	C.handlerFree(*h)
	*h = nil
}
