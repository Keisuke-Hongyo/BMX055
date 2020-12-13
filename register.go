package bmx055

// The I2C address which this device listens to.
const ACCAddress = 0x19  // ACC
const GyroAddress = 0x69 // Gyro
const MagAddress = 0x13  // Mag

// Device Register
const (
	BGW_CHIPID         = 0x00 // Who am I
	AccDevice          = 0xFA // Device check register
	ResetReg           = 0x14 // Reset register
	InitiatedSoftReset = 0xB6 // Soft reset parameter
	AccDataStartReg    = 0x02
	AccPmuRangeReg     = 0x0F
	AccPmuBwReg        = 0x10
	AccPmuLpwReg       = 0x11
	GyroDataStartReg   = 0x02
	GyroRangeReg       = 0x0F
	GyroBwReg          = 0x10
	GyroLpm1Reg        = 0x11
	MagChipID          = 0x40
	MagDataStartReg    = 0x42
	MagPowCtlReg       = 0x4B
	MagAdvOpOutputReg  = 0x4C
	MagAxesReg         = 0x4E
	MagRepXyReg        = 0x51
	MagRepZReg         = 0x52
)
