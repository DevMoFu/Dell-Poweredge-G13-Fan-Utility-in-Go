package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//// create impitool boilerplate to enforce DRY
func ipmiBoilerplate(args string, algoName string) {
	argsArray := strings.Split(args, " ")

	//fix for items that need white space in a single arg. Replace underscore.
	for i := range argsArray {
		matched, _ := regexp.MatchString(`_`, argsArray[i])
		if matched {
			argsArray[i] = strings.Replace(argsArray[i], "_", " ", 3)
			println(argsArray[i])
		}
	}

	cmd := exec.Command("ipmitool", argsArray...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%v failed with: \n%v\n", algoName, err)
		//os.Exit(1)
	}
	fmt.Printf("\n%s\n\n", stdout)
}

func checkSystemSensors(c string) {
	algoName := "checkSystemSensors"
	args := fmt.Sprintf("%vsensor", c)
	fmt.Println("System Sensors:")
	ipmiBoilerplate(args, algoName)
}

func checkSystemTemps(c string) {
	//  sensor reading 'Inlet Temp' Temp
	algoName := "checkSystemTemps"
	// Split in ipmiBoilerplate breaks 'Inlet Temp' arg.
	// Fix has been added to make 'Inlet_Temp' -> 'Intlet Temp'
	args := fmt.Sprintf("%vsensor reading Temp Inlet_Temp", c)
	fmt.Println("System Temps:")
	ipmiBoilerplate(args, algoName)
}

func checkCurrentFanSpeed(c string) {
	algoName := "checkCurrentFanSpeed"
	args := fmt.Sprintf("%ssensor reading Fan1A Fan1B Fan2A Fan2B Fan3A Fan3B Fan4A Fan4B", c)
	fmt.Println("Current Fan Speeds:")
	ipmiBoilerplate(args, algoName)
}

func checkThirdPartyCardBehavior(c string) {
	algoName := "checkThirdPartyCardBehavior"
	args := fmt.Sprintf("%sraw 0x30 0xce 0x01 0x16 0x05 0x00 0x00 0x00", c)
	ipmiBoilerplate(args, algoName)
}

func setManualFanMode(c string, ManualFanMode string) {
	algoName := "setManualFanMode"
	var ManualFanModeHex string
	switch ManualFanMode {
	case "disable":
		fmt.Printf("%s: Disabled\n", algoName)
		ManualFanModeHex = "0x30 0x30 0x01 0x01"
	case "enable":
		fmt.Printf("%s: Enabled\n", algoName)
		ManualFanModeHex = "0x30 0x30 0x01 0x00"
	default:
		fmt.Printf("%s: '%s' is not a valid value. Try 'enable' or 'disable\n", algoName, ManualFanMode)
		os.Exit(1)
	}

	args := fmt.Sprintf("%s%s %s", c, "raw", ManualFanModeHex)
	ipmiBoilerplate(args, algoName)
}

func setFanSpeed(c string, FanSpeed int) {
	algoName := "setFanSpeed"
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
		fmt.Printf("%v: '%s' is not a valid value. Try 'enable' or 'disable\n", FanSpeed, algoName)
		os.Exit(1)
	}

	// TODO: Notify user of fan mod change.
	fmt.Printf("Automatically enabling Manual Fan Mode to set Fan Speed to %v%%\n", FanSpeed)
	setManualFanMode(c, "enable")

	args := fmt.Sprintf("%s%s %s", c, "raw", fanSpeedHex)

	ipmiBoilerplate(args, algoName)
}

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
	ipmiBoilerplate(args, algoName)
}

type creds struct {
	hostnameIP string
	username   string
	password   string
}

// TODO: Add logic for option selection
// Use ping node test and system info to verify connectivity
// Show system temp fan
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
	// Silence Go when no using associated flags
	fmt.Sprintln(ManualFanMode, FanSpeed, thirdPartyCardBehavior)

	c := creds{*hostnameIP, *username, *password}
	credString := fmt.Sprintf("-I lanplus -H %v -U %v -P %v ", c.hostnameIP, c.username, c.password)

	checkSystemTemps(credString)
	//checkCurrentFanSpeed(credString)
	//checkThirdPartyCardBehavior(credString)
	//setManualFanMode(credString, *ManualFanMode)
	//setThirdPartyCardBehavior(credString, *thirdPartyCardBehavior)
	setFanSpeed(credString, *FanSpeed)

}

/*
	Note: “TEMP” was the name of my sensor, You can try “Ambient Temperature” for your server if you want to see the temperature of the CPU.  Also, the fan names are case sensitive.  So save yourself a few moments of troubleshooting by using the name as reported in iDRAC.
	This command will print out a ton of information about the Fans, stats for nerds basically.

https://github.com/ipmitool/ipmitool/issues/30
https://www.reddit.com/r/homelab/comments/7xqb11/dell_fan_noise_control_silence_your_poweredge/
*/
