package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type creds struct {
	hostnameIP string
	username   string
	password   string
}

func checkSystemTemps(c string) string {
	args := fmt.Sprintf("%v reading 'Temp' 'Fan1A' 'Fan1B' 'Fan2A' 'Fan2B' 'Fan3A' 'Fan3B' 'Fan4A' 'Fan4B'", c)
	cmd := exec.Command("ipmitool", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run
	if err != nil {
		log.Fatalf("Comand failed with %v\n", err)
		os.Exit(1)
	}

	return cmd.Stdout
}

func checkCurrentFanSpeed(c creds) {
	args := fmt.Sprintf("-I lanplus -H %v -U %v -P %v sdr get 'Fan1A' 'Fan1B' 'Fan2A' 'Fan2B' 'Fan3A' 'Fan3B' 'Fan4A' 'Fan4B'", c.hostnameIP, c.username, c.password)
	exec.Command("ipmitool", args)
}

func checkThirdPartyCardBehavior(c creds) {
	args := fmt.Sprintf("-I lanplus -H %v -U %v -P %v raw 0x30 0xce 0x01 0x16 0x05 0x00 0x00 0x00", c.hostnameIP, c.username, c.password)
	exec.Command("ipmitool", args)
}

func setFanSpeed() {

}

func main() {

	/*
			This command will print information about the System Temperature and FAN RPMs.

		    ipmitool -I lanplus -H yourIPAddress -U yourUsername -P yourPassword sensor reading "Temp" 'Fan1A' 'Fan1B' 'Fan2A' 'Fan2B' 'Fan3A' 'Fan3B' 'Fan4A' 'Fan4B'

		Note: “TEMP” was the name of my sensor, You can try “Ambient Temperature” for your server if you want to see the temperature of the CPU.  Also, the fan names are case sensitive.  So save yourself a few moments of troubleshooting by using the name as reported in iDRAC.
		This command will print out a ton of information about the Fans, stats for nerds basically.

			ipmitool -I lanplus -H yourIPAddress -U yourUsername -P yourPassword sdr get 'Fan1A' 'Fan1B' 'Fan2A' 'Fan2B' 'Fan3A' 'Fan3B' 'Fan4A' 'Fan4B'

			ipmitool -I lanplus -H yourIPAddress -U yourUsername -P yourPassword raw

	*/

	// credential
	HostnameIP := flag.String("Hostname", "", "Hostname or IP")
	Username := flag.String("Username", "", "Username")
	Password := flag.String("Password", "", "Password")
	c := creds{*HostnameIP, *Username, *Password}
	credString := fmt.Sprintf("-I lanplus -H yourIPAddress -U yourUsername -P yourPassword ", c.hostnameIP, c.username, c.password)

	// modifying args
	ManualFanMode := flag.String("ManualFanMode", "na", "'enable' or 'disable' manual fan control")
	FanSpeed := flag.Int("FanSpeed", 888, "10 < 'init' < 100 in increments of 5\nFanMode required to be enabled")
	thirdPartyCardBehavior := flag.String("ThirdPartyCardBehavior", "", "'enable' or 'disable' 3rd Party Fan Behavior")
	flag.Parse()

	// 3rd Party Card fan behavoir
	switch *thirdPartyCardBehavior {
	case "disable":
		thirdPartyCardBehaviorHex := "0x30 0xce 0x00 0x16 0x05 0x00 0x00 0x00 0x05 0x00 0x01 0x00 0x00" // disable
	case "enable":
		thirdPartyCardBehaviorHex := "0x30 0xce 0x00 0x16 0x05 0x00 0x00 0x00 0x05 0x00 0x00 0x00 0x00" // enable
	default:
		fmt.Println(" Not or no valid 3rd Party Behavior selected")
	}

	// fan mode
	var ManualFanModeHex string
	switch FanMode {
	case "auto":
		ManualFanModeHex = "0x30 0x30 0x01 0x01" // disable manual fan control
	case "manual":
		ManualFanModeHex = "0x30 0x30 0x01 0x00" // enable manual fan control
	default:
		fmt.Println("")
	}

	// fan speed
	var fanSpeedHex string
	switch FanSpeed {
	case 10:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x0a"
	case 15:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x0f"
	case 20:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x14"
	case 25:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x19"
	case 30:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x1e"
	case 35:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x23"
	case 40:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x28"
	case 45:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x2d"
	case 50:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x32"
	case 55:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x37"
	case 60:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x3c"
	case 65:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x41"
	case 70:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x46"
	case 75:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x4b"
	case 80:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x50"
	case 85:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x55"
	case 90:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x5a"
	case 95:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x5f"
	case 100:
		fanSpeedHex = "0x30 0x30 0x02 0xff 0x5a"
	default:
		println("")
	}
}

/*
https://github.com/ipmitool/ipmitool/issues/30
https://www.reddit.com/r/homelab/comments/7xqb11/dell_fan_noise_control_silence_your_poweredge/

# print temps and fans rpms
ipmitool -I lanplus -H <iDRAC-IP> -U <iDRAC-USER> -P <iDRAC-PASSWORD> sensor reading "Ambient Temp" "FAN 1 RPM" "FAN 2 RPM" "FAN 3 RPM"
#
# print fan info
ipmitool -I lanplus -H <iDRAC-IP> -U <iDRAC-USER> -P <iDRAC-PASSWORD> sdr get "FAN 1 RPM" "FAN 2 RPM" "FAN 3 RPM"
#
# enable manual/static fan control
ipmitool -I lanplus -H <iDRAC-IP> -U <iDRAC-USER> -P <iDRAC-PASSWORD> raw 0x30 0x30 0x01 0x00
#
# disable manual/static fan control
ipmitool -I lanplus -H <iDRAC-IP> -U <iDRAC-USER> -P <iDRAC-PASSWORD> raw 0x30 0x30 0x01 0x01
#

Function 	ipmitool raw command
Disable 3rd Party Card fan behavior 	ipmitool raw 0x30 0xce 0x00 0x16 0x05 0x00 0x00 0x00 0x05 0x00 0x01 0x00 0x00
Enable 3rd Party Card fan behavior 	ipmitool raw 0x30 0xce 0x00 0x16 0x05 0x00 0x00 0x00 0x05 0x00 0x00 0x00 0x00
Check 3rd Party Card fan behavior 	ipmitool raw 0x30 0xce 0x01 0x16 0x05 0x00 0x00 0x00
*/
