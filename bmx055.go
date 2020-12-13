package bmx055

import (
	"time"
	"tinygo.org/x/drivers"
)

type bmx055_data struct {
	// Acc
	xAcc float64
	yAcc float64
	zAcc float64

	// Gyaro
	xGyro float64
	yGyro float64
	zGyro float64

	// Mag
	xMag int32
	yMag int32
	zMag int32
}

type Device struct {
	bus         drivers.I2C
	AccAddress  uint8
	GyroAddress uint8
	MagAddress  uint8
	dt          bmx055_data
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
	d.bus.WriteRegister(d.AccAddress, AccPmuRangeReg, []byte{0x03})
	d.bus.WriteRegister(d.AccAddress, AccPmuBwReg, []byte{0x08})
	d.bus.WriteRegister(d.AccAddress, AccPmuLpwReg, []byte{0x00})
	time.Sleep(100 * time.Millisecond)

	/* Gyro*/
	d.bus.WriteRegister(d.GyroAddress, GyroRangeReg, []byte{0x04})
	d.bus.WriteRegister(d.GyroAddress, GyroBwReg, []byte{0x07})
	d.bus.WriteRegister(d.GyroAddress, GyroLpm1Reg, []byte{0x00})
	time.Sleep(100 * time.Millisecond)

	/* Mag */
	d.bus.WriteRegister(d.MagAddress, MagPowCtlReg, []byte{0x83})
	time.Sleep(100 * time.Millisecond)
	d.bus.WriteRegister(d.MagAddress, MagPowCtlReg, []byte{0x01})
	time.Sleep(100 * time.Millisecond)

	d.bus.WriteRegister(d.MagAddress, MagAdvOpOutputReg, []byte{0x00})
	d.bus.WriteRegister(d.MagAddress, MagAxesReg, []byte{0x84})
	d.bus.WriteRegister(d.MagAddress, MagRepXyReg, []byte{0x04})
	d.bus.WriteRegister(d.MagAddress, MagRepZReg, []byte{0x16})

}

//isConnected
func (d *Device) isConnected() bool {
	data := make([]byte, 1)
	d.bus.ReadRegister(d.AccAddress, BGW_CHIPID, data)
	return data[0] == 0xfa
}

/* Get Acceralater Data Function */
func (d *Device) getAcc() error {
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

	d.dt.xAcc = float64(x) * 0.0196
	d.dt.yAcc = float64(y) * 0.0196
	d.dt.zAcc = float64(z) * 0.0196

	return nil
}

/* Get Gyro Data Function */
func (d *Device) getGyro() error {
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
	d.dt.xGyro = float64(x) * 0.0038
	d.dt.yGyro = float64(y) * 0.0038
	d.dt.zGyro = float64(z) * 0.0038

	return nil
}

/* Get Magnetic Filed Data Function */
func (d *Device) getMag() error {

	data := make([]byte, 6)

	err := d.bus.ReadRegister(d.MagAddress, MagDataStartReg, data)
	if err != nil {
		println(err)
		return err
	}
	// Get ACC Data
	d.dt.xMag = int32(uint16(data[1])<<8|uint16(data[0]&0xf8)) / 8
	if d.dt.xMag > 4095 {
		d.dt.xMag -= 8192
	}
	d.dt.yMag = int32(uint16(data[3])<<8|uint16(data[2]&0xf8)) / 8
	if d.dt.yMag > 4095 {
		d.dt.yMag -= 8192
	}
	d.dt.zMag = int32(uint16(data[5])<<8|uint16(data[4]&0xfe)) / 2
	if d.dt.zMag > 16383 {
		d.dt.zMag -= 32768
	}
	return nil
}
