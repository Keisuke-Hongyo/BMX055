package bmx055

import (
	"fmt"
	"machine"
	"time"
)

func main() {

	var err error

	machine.I2C0.Configure(machine.I2CConfig{})
	d := New(machine.I2C0)

	conected := d.isConnected()

	// BMX055 is Conenect ?
	if !conected {
		// Not Connected
		println("Not Connected!")
		return
	}

	// Initilaize For BMX055
	d.Configture()

	for {
		/* Get Accelerometer Data */
		err = d.getGyro()
		if err == nil {
			fmt.Printf("Gyro : x=%5.2f y=%5.2f z=%5.2f\n", d.dt.xGyro, d.dt.yGyro, d.dt.zGyro)
		}

		/* Get Gyro Data */
		err = d.getAcc()
		if err == nil {
			fmt.Printf("Acc  : x=%5.2f y=%5.2f z=%5.2f\n", d.dt.xAcc, d.dt.yAcc, d.dt.zAcc)
		}

		/* Get Geomagnetic sensor Data*/
		err = d.getMag()
		if err == nil {
			fmt.Printf("Mag  : x=%4d y=%4d z=%4d\n", d.dt.xMag, d.dt.yMag, d.dt.zMag)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
