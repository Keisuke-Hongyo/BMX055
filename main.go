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

	conected := d.IsConnected()

	// BMX055 is Conenect ?
	if !conected {
		// Not Connected
		println("Not Connected!")
		for {
		}
	}

	// Initilaize For BMX055
	d.Configture()

	for {

		/* Get Accelerometer Data */
		err = d.GetGyro()
		if err == nil {
			fmt.Printf("Gyro : x=%5.2f y=%5.2f z=%5.2f\n", d.Dat.XGyro, d.Dat.YGyro, d.Dat.ZGyro)
		}

		/* Get Gyro Data */
		err = d.GetAcc()
		if err == nil {
			fmt.Printf("Acc  : x=%5.2f y=%5.2f z=%5.2f\n", d.Dat.XAcc, d.Dat.YAcc, d.Dat.ZAcc)
		}

		/* Get Geomagnetic sensor Data*/
		err = d.GetMag()
		if err == nil {
			fmt.Printf("Mag  : x=%4d y=%4d z=%4d\n", d.Dat.XMag, d.Dat.YMag, d.Dat.ZMag)
		}

		fmt.Println()

		time.Sleep(100 * time.Millisecond)
	}
}
