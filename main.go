package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// Create impitool boilerplate to enforce DRY
func ipmiBoilerplate(args string, algoName string) {
	argsArray := strings.Split(args, " ")

	// Fix for items that need white space in a single arg. Replace underscore.
	for i := range argsArray {
		matched, _ := regexp.MatchString(`_`, argsArray[i])
		if matched {
			argsArray[i] = strings.Replace(argsArray[i], "_", " ", 3)
		}
	}

	cmd := exec.Command("ipmitool", argsArray...)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%v failed with:\n%v\n", algoName, err)
		if stdout != nil {
			fmt.Printf("Stdout:\n%s", stdout)
		}
		os.Exit(1)
	}
	fmt.Printf("%s\n", stdout)
}

func checkSystemSensors(c string) {
	algoName := "checkSystemSensors"
	args := fmt.Sprintf("%vsensor", c)
	fmt.Println("System Sensors:")
	ipmiBoilerplate(args, algoName)
}

func checkSystemTemps(c string) {
	algoName := "checkSystemTemps"
	// Split in ipmiBoilerplate breaks 'Inlet Temp' arg.
	// Fix has been added to make 'Inlet_Temp' -> 'Intlet Temp'
	args := fmt.Sprintf("%vsensor reading Temp Inlet_Temp", c)
	fmt.Println("Current System Temps:")
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
	fmt.Println("Current 3rd Party Card Behavior:")
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

	// Notify user of fan mode change and speed selected
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

type userCredentials struct {
	hostnameIP string
	username   string
	password   string
}

type userInput struct {
	manualFanMode          string
	fanSpeed               int
	thirdPartyCardBehavior string
	time                   time.Time
}

func main() {

	// Credential
	hostnameIP := flag.String("H", "", "Idrac Hostname or IP")
	username := flag.String("U", "", "Idrac username")
	password := flag.String("P", "", "Idrac password")
	// Modifying args
	ManualFanMode := flag.String("ManualFanMode", "", "'enable' or 'disable' manual fan control")
	FanSpeed := flag.Int("FanSpeed", 888, "10 < 'int' < 100 in increments of 5\nFan mode will automatically set to 'enable' if speed is selected")
	thirdPartyCardBehavior := flag.String("ThirdPartyCardBehavior", "", "'enable' or 'disable' third party fan behavior")
	flag.Parse()

	u := userInput{*ManualFanMode, *FanSpeed, *thirdPartyCardBehavior, time.Now()}
	c := userCredentials{*hostnameIP, *username, *password}
	credString := fmt.Sprintf("-I lanplus -H %v -U %v -P %v ", c.hostnameIP, c.username, c.password)
	fmt.Println("")
	// Silence Go when no using associated flags
	fmt.Sprintln(ManualFanMode, FanSpeed, thirdPartyCardBehavior)

	// Verify OS and other dependencies
	if runtime.GOOS != "linux" {
		fmt.Println("This tool is only compatible with linux!")
		os.Exit(1)
	}

	// Verify fields are not == ""
	if c.hostnameIP == "" || c.password == "" || c.username == "" {
		fmt.Println("(Hostname|Ip), Password or Username Missing.")
		fmt.Println("Type 'Help' to see input options")
	}

	// Verify idrac connectivity while showing current status
	// Wrap in function
	checkSystemTemps(credString)
	checkCurrentFanSpeed(credString)
	checkThirdPartyCardBehavior(credString)

	// User logic starts here
	// Perform actions
	if u.fanSpeed != 888 {
		setFanSpeed(credString, *FanSpeed)
	} else if u.manualFanMode != "" {
		setManualFanMode(credString, *ManualFanMode)
	}

	if u.thirdPartyCardBehavior != "" {
		setThirdPartyCardBehavior(credString, *thirdPartyCardBehavior)
	}
}
