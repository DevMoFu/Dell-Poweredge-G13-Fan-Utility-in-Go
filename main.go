package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//// create impitool boilerplate to enforce DRY
//works
func ipmiBoilerplate(args string, algoName string) {
	argsArray := strings.Split(args, " ")
	cmd := exec.Command("ipmitool", argsArray...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%v failed with: \n%v\n", algoName, err)
		//os.Exit(1)
	}
	fmt.Printf("\n%s\n\n", stdout)
}

// May no longer be needed
//func setIpmiBoilerplate(args string, hexcode string, algoName string) {
//	argsArray := strings.Split(args, " ")
//	argsSlice := argsArray[:]
//	argsSlice = append(argsSlice, hexcode)
//	fmt.Printf("%s\n\n", argsSlice) // test
//	cmd := exec.Command("ipmitool", argsSlice...)
//	stdout, err := cmd.CombinedOutput()
//	if err != nil {
//		fmt.Printf("%v failed with: \n%v\n", algoName, err)
//		//os.Exit(1)
//	}
//	fmt.Printf("\n%s\n\n", stdout)
//}

// works
func checkSystemTemps(c string) {
	algoName := "checkSystemTemps"
	args := fmt.Sprintf("%vsensor", c)
	ipmiBoilerplate(args, algoName)
}

// works
func checkCurrentFanSpeed(c string) {
	algoName := "checkSystemTemps"
	args := fmt.Sprintf("%ssensor reading Fan1A Fan1B Fan2A Fan2B Fan3A Fan3B Fan4A Fan4B", c)
	ipmiBoilerplate(args, algoName)
}

//works
// Add output based on stdout
func checkThirdPartyCardBehavior(c string) {
	algoName := "checkThirdPartyCardBehavior"
	args := fmt.Sprintf("%sraw 0x30 0xce 0x01 0x16 0x05 0x00 0x00 0x00", c)
	ipmiBoilerplate(args, algoName)
}

//WIP
func setManualFanMode(c string, ManualFanMode string) {
	algoName := "setManualFanMode"
	var ManualFanModeHex string
	switch ManualFanMode {
	case "disable":
		fmt.Printf("%s: Disabled\n", algoName)
		ManualFanModeHex = "0x30 0x30 0x01 0x01" // disable manual fan control
	case "enable":
		fmt.Printf("%s: Enabled\n", algoName)
		ManualFanModeHex = "0x30 0x30 0x01 0x00" // enable manual fan control
	default:
		fmt.Printf("%s: '%s' is not a valid value. Try 'enable' or 'disable\n", algoName, ManualFanMode)
		os.Exit(1)
	}

	// May not be needed
	//args := fmt.Sprintf("%s%s", c, "raw")
	//fmt.Printf("%s %s\n%s\n\n", args, ManualFanModeHex, algoName) //test //works
	//setIpmiBoilerplate(args, ManualFanModeHex, algoName)

	args := fmt.Sprintf("%s%s %s", c, "raw", ManualFanModeHex)
	println(args)
	ipmiBoilerplate(args, algoName) // s
}

//func setFanSpeed(FanSpeed int) {
//	algoName := "setFanSpeed"
//	// fan speed
//	var fanSpeedHex string
//	switch FanSpeed {
//	case 10:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x0a"
//	case 15:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x0f"
//	case 20:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x14"
//	case 25:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x19"
//	case 30:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x1e"
//	case 35:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x23"
//	case 40:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x28"
//	case 45:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x2d"
//	case 50:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x32"
//	case 55:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x37"
//	case 60:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x3c"
//	case 65:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x41"
//	case 70:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x46"
//	case 75:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x4b"
//	case 80:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x50"
//	case 85:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x55"
//	case 90:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x5a"
//	case 95:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x5f"
//	case 100:
//		fanSpeedHex = "0x30 0x30 0x02 0xff 0x5a"
//	default:
//		fmt.Printf("Null or no valid 3rd Party Behavior selected. No action taken", algoName)
//	}
//}

func setThirdPartyCardBehavior(c string, thirdPartyCardBehavior string) {
	algoName := "setThirdPartyCardBehavior"
	var thirdPartyCardBehaviorHex string
	switch thirdPartyCardBehavior {
	case "disable":
		fmt.Printf("%s: 'disable' option selected\n", algoName)
		thirdPartyCardBehaviorHex = "0x30 0xce 0x00 0x16 0x05 0x00 0x00 0x00 0x05 0x00 0x01 0x00 0x00"
	case "enable":
		fmt.Printf("%s: 'enabled' option selected\n", algoName)
		thirdPartyCardBehaviorHex = "0x30 0xce 0x00 0x16 0x05 0x00 0x00 0x00 0x05 0x00 0x00 0x00 0x00"
	default:
		fmt.Printf("%s: '%s' is not a valid value. Try 'enable' or 'disable\n", thirdPartyCardBehaviorHex, algoName)
	}

	args := fmt.Sprintf("%s%s %s", c, "raw", thirdPartyCardBehaviorHex)
	println(args)
	ipmiBoilerplate(args, algoName)
}

type creds struct {
	hostnameIP string
	username   string
	password   string
}

func main() {

	// credential
	hostnameIP := flag.String("H", "", "Hostname or IP")
	username := flag.String("U", "", "Username")
	password := flag.String("P", "", "Password")
	// modifying args
	ManualFanMode := flag.String("ManualFanMode", "na", "'enable' or 'disable' manual fan control")
	FanSpeed := flag.Int("FanSpeed", 888, "10 < 'init' < 100 in increments of 5\nFanMode required to be enabled")
	thirdPartyCardBehavior := flag.String("ThirdPartyCardBehavior", "", "'enable' or 'disable' 3rd Party Fan Behavior")
	flag.Parse()

	c := creds{*hostnameIP, *username, *password}
	credString := fmt.Sprintf("-I lanplus -H %v -U %v -P %v ", c.hostnameIP, c.username, c.password)
	//checkSystemTemps(credString)
	//checkCurrentFanSpeed(credString)
	checkThirdPartyCardBehavior(credString)
	//setManualFanMode(credString, *ManualFanMode)
	setThirdPartyCardBehavior(credString, *thirdPartyCardBehavior)

	fmt.Sprintln(ManualFanMode, FanSpeed, thirdPartyCardBehavior) // remove after testing

}

/*
		This command will print information about the System Temperature and FAN RPMs.

	    ipmitool -I lanplus -H yourIPAddress -U yourUsername -P yourPassword sensor reading "Temp" 'Fan1A' 'Fan1B' 'Fan2A' 'Fan2B' 'Fan3A' 'Fan3B' 'Fan4A' 'Fan4B'

	Note: “TEMP” was the name of my sensor, You can try “Ambient Temperature” for your server if you want to see the temperature of the CPU.  Also, the fan names are case sensitive.  So save yourself a few moments of troubleshooting by using the name as reported in iDRAC.
	This command will print out a ton of information about the Fans, stats for nerds basically.

		ipmitool -I lanplus -H yourIPAddress -U yourUsername -P yourPassword sdr get 'Fan1A' 'Fan1B' 'Fan2A' 'Fan2B' 'Fan3A' 'Fan3B' 'Fan4A' 'Fan4B'

		ipmitool -I lanplus -H yourIPAddress -U yourUsername -P yourPassword raw

*/

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
