# Dell-Poweredge-G13-Fan-Utility-in-Go

Dell PowerEdge Gen 13 fan control CLI Utility
***Not intended for PRODUCTION use***
***Not Affiliation is Dell***

## How

Ipmi Wrapper made with Go for Linux

## Why

My Dell r330s are a little too loud. Don't want to use Bash anymore. I like Go...

## ToDo

- [x] Collect hex codes for raw commands
- [x] Query current system temps and fan speeds
- [x] Add fan mode control
- [x] Add fan speed control
- [x] Add 3rd party card behavior control
- [ ] Add flag to view all sensors data
- [x] Add logic for user interation
- [x] Add check for dependencies (ipmi/ipmitool) (linux)
- [ ] Add steps to enable ipmi over Lan to READ.md
- [ ] Add help to READ.md
- [ ] Cross post with r/homelab and r/golang for feedback

## Prerequisites

- [Linux](https://www.linux.org/)
- [Git](https://git-scm.com/)
- [Go](https://golang.org/)
- [ipmitool](http://www.aslab.com/support/kb/224.html) installed and working
- [Dell PowerEdge with idrac]

## To Build

1. Clone Repo `Git clone https://github.com/DevMoFu/Dell-Poweredge-G13-Fan-Utility-in-Go/blob/master/main.go`
2. `Go build -o "Your desired name" main.go`

## Args

```go
Usage:
  -FanSpeed int
        10 < 'init' < 100 in increments of 5
        FanMode required to be enabled (default 888)
  -H string
        Hostname or IP
  -ManualFanMode string
        'enable' or 'disable' manual fan control (default "na")
  -P string
        Password
  -ThirdPartyCardBehavior string
        'enable' or 'disable' 3rd Party Fan Behavior
  -U string
        Username
```

## Sample Input

```bash
"Set fan speed to 20%"
./fanUtility -H <idrac hostname or ip> -U <username> -P <password> -FanSpeed 20

"Disable 3rd part card behavior (Fans running on high while a "non-supported" card is in the PCIe slot)"
./fanUtility -H <idrac hostname or ip> -U <username> -P <password> -ThirdPartyCardBehavior disable

"Disable Manual Fan Mode and allow system to manage fan speeds"
./fanUtility -H <idrac hostname or ip> -U <username> -P <password> -ManualFanMode disable
```

## Sample Output

```bash
[]# ./fanUtility -H <idrac hostname or ip> -U <username> -P <password> -FanSpeed 20

Current System Temps:
Temp             | 64
Inlet Temp       | 21

Current Fan Speeds:
Fan1A            | 4080
Fan1B            | 3720
Fan2A            | 4080
Fan2B            | 3720
Fan3A            | 3960
Fan3B            | 3720
Fan4A            | 3960
Fan4B            | 3720

Current 3rd Party Card Behavior:
 16 05 00 00 00 05 00 01 00 00

Automatically enabling Manual Fan Mode to set Fan Speed to 20%
setManualFanMode: Enabled

[]# ./fanUtility -H <idrac hostname or ip> -U <username> -P <password> -ThirdPartyCardBehavior disable

Current System Temps:
Temp             | 64
Inlet Temp       | 21

Current Fan Speeds:
Fan1A            | 4080
Fan1B            | 3720
Fan2A            | 4080
Fan2B            | 3720
Fan3A            | 3960
Fan3B            | 3720
Fan4A            | 3960
Fan4B            | 3720

Current 3rd Party Card Behavior:
 16 05 00 00 00 05 00 01 00 00

setThirdPartyCardBehavior: 'disable' option selected
 16 05 00 00 00
```

## Tested On

- Dell PowerEdge R330

## Sources/Inspiration

- https://github.com/ipmitool/ipmitool/issues/30
- https://www.reddit.com/r/homelab/comments/7xqb11/dell_fan_noise_control_silence_your_poweredge/
- https://www.spxlabs.com/blog/2019/3/16/silence-your-dell-poweredge-server