package bmx055

import (
	"time"
	"tinygo.org/x/drivers"
)

type bmx055_data struct {
	// Acc
	XAcc float64
	YAcc float64
	ZAcc float64

	// Gyaro
	XGyro float64
	YGyro float64
	ZGyro float64

	// Mag
	XMag int32
	YMag int32
	ZMag int32
}

type Device struct {
	bus         drivers.I2C
	AccAddress  uint8
	GyroAddress uint8
	MagAddress  uint8
	Dat         bmx055_data
}

func New(bus drivers.I2C) Device {
	return Device{
		bus:         bus,
		AccAddress:  ACCAddress,
		GyroAddress: GyroAddress,
		MagAddress:  MagAddress,
	}
}

/* Init*/
func (d *Device) Configture() {
	/* Acc*/
	_ = d.bus.WriteRegister(d.AccAddress, AccPmuRangeReg, []byte{0x03})
	_ = d.bus.WriteRegister(d.AccAddress, AccPmuBwReg, []byte{0x08})
	_ = d.bus.WriteRegister(d.AccAddress, AccPmuLpwReg, []byte{0x00})
	time.Sleep(100 * time.Millisecond)

	/* Gyro*/
	_ = d.bus.WriteRegister(d.GyroAddress, GyroRangeReg, []byte{0x04})
	_ = d.bus.WriteRegister(d.GyroAddress, GyroBwReg, []byte{0x07})
	_ = d.bus.WriteRegister(d.GyroAddress, GyroLpm1Reg, []byte{0x00})
	time.Sleep(100 * time.Millisecond)

	/* Mag */
	_ = d.bus.WriteRegister(d.MagAddress, MagPowCtlReg, []byte{0x83})
	time.Sleep(100 * time.Millisecond)
	_ = d.bus.WriteRegister(d.MagAddress, MagPowCtlReg, []byte{0x01})
	time.Sleep(100 * time.Millisecond)

	_ = d.bus.WriteRegister(d.MagAddress, MagAdvOpOutputReg, []byte{0x00})
	_ = d.bus.WriteRegister(d.MagAddress, MagAxesReg, []byte{0x84})
	_ = d.bus.WriteRegister(d.MagAddress, MagRepXyReg, []byte{0x04})
	_ = d.bus.WriteRegister(d.MagAddress, MagRepZReg, []byte{0x16})

}

//isConnected
func (d *Device) IsConnected() bool {
	data := make([]byte, 1)
	_ = d.bus.ReadRegister(d.AccAddress, BGW_CHIPID, data)
	return data[0] == 0xfa
}

/* Get Acceralater Data Function */
func (d *Device) GetAcc() error {
	data := make([]byte, 6)
	err := d.bus.ReadRegister(d.AccAddress, AccDataStartReg, data)
	if err != nil {
		println(err)
		return err
	}
	// Get ACC Data

	x := int32(uint16(data[1])*256|uint16(data[0]&0xf0)) >> 4
	if x > 2047 {
		x -= 4096
	}
	y := int32(uint16(data[3])*256|uint16(data[2]&0xf0)) >> 4
	if y > 2047 {
		y -= 4096
	}
	z := int32(uint16(data[5])*256|uint16(data[4]&0xf0)) >> 4
	if z > 2047 {
		z -= 4096
	}

	d.Dat.XAcc = float64(x) * 0.0196
	d.Dat.YAcc = float64(y) * 0.0196
	d.Dat.ZAcc = float64(z) * 0.0196

	return nil
}

/* Get Gyro Data Function */
func (d *Device) GetGyro() error {
	data := make([]byte, 6)
	err := d.bus.ReadRegister(d.GyroAddress, GyroDataStartReg, data)
	if err != nil {
		println(err)
		return err
	}

	// Get Gyro Data
	x := int32(uint16(data[1])<<8 | uint16(data[0]))
	if x > 32767 {
		x -= 65536
	}
	y := int32(uint16(data[3])<<8 | uint16(data[2]))
	if y > 32767 {
		y -= 65536
	}
	z := int32(uint16(data[5])<<8 | uint16(data[4]))
	if z > 32767 {
		z -= 65536
	}

	//
	d.Dat.XGyro = float64(x) * 0.0038
	d.Dat.YGyro = float64(y) * 0.0038
	d.Dat.ZGyro = float64(z) * 0.0038

	return nil
}

/* Get Magnetic Filed Data Function */
func (d *Device) GetMag() error {

	data := make([]byte, 6)

	err := d.bus.ReadRegister(d.MagAddress, MagDataStartReg, data)
	if err != nil {
		println(err)
		return err
	}
	// Get ACC Data
	d.Dat.XMag = int32(uint16(data[1])<<8|uint16(data[0]&0xf8)) / 8
	if d.Dat.XMag > 4095 {
		d.Dat.XMag -= 8192
	}
	d.Dat.YMag = int32(uint16(data[3])<<8|uint16(data[2]&0xf8)) / 8
	if d.Dat.YMag > 4095 {
		d.Dat.YMag -= 8192
	}
	d.Dat.ZMag = int32(uint16(data[5])<<8|uint16(data[4]&0xfe)) / 2
	if d.Dat.ZMag > 16383 {
		d.Dat.ZMag -= 32768
	}
	return nil
}
