package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jacobsa/go-serial/serial"
)

type Args struct {
	Device string
	Power  string
	Music  int
	Mic    int
	Echo   int
}

func main() {
	args := getArgs()

	options := serial.OpenOptions{
		PortName:        args.Device,
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	defer port.Close()

	powerData := getPowerData()
	musicMicData := getMusicMicData()
	echoData := getEchoData()

	data := []byte{
		getStx(),
		getTextTop(),
		powerData[args.Power],
		musicMicData[args.Music],
		musicMicData[args.Mic],
		echoData[args.Echo],
		getEtx(),
	}

	data = append(data, getXor(data))

	fmt.Print(data)
	fmt.Print("\n")

	n, err := port.Write(data)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}

/**
 * get argument.
 */
func getArgs() Args {
	device := flag.String("device", "/dev/ttyUSB0", "device. example... \"/dev/ttyUSB0\"")
	power := flag.String("power", "on", "amp power. on or off")
	music := flag.Int("music", -20, "music volume. -62 ~ 0")
	mic := flag.Int("mic", -20, "mic volume. -62 ~ 0")
	echo := flag.Int("echo", 20, "echo volume. 0 ~ 63")

	flag.Parse()

	return Args{
		Device: *device,
		Power:  *power,
		Music:  *music,
		Mic:    *mic,
		Echo:   *echo,
	}
}

/**
 * get start of text.
 */
func getStx() byte {
	return 0x02
}

/**
 * get end of text.
 */
func getEtx() byte {
	return 0x03
}

/**
 * get text top.
 */
func getTextTop() byte {
	return 0x2F
}

/**
 * get exclusive or.
 */
func getXor(b []byte) byte {
	return b[1] ^ b[2] ^ b[3] ^ b[4] ^ b[5] ^ b[6]
}

/**
 * get power data.
 */
func getPowerData() map[string]byte {
	return map[string]byte{
		"off": 0x30,
		"on":  0x31,
	}
}

/**
 * get music mic up down data.
 */
func getMusicMicData() map[int]byte {
	return map[int]byte{
		0:   0x7F,
		-1:  0x7E,
		-2:  0x7D,
		-3:  0x7C,
		-4:  0x7B,
		-5:  0x7A,
		-6:  0x79,
		-7:  0x78,
		-8:  0x77,
		-9:  0x76,
		-10: 0x75,
		-11: 0x74,
		-12: 0x73,
		-13: 0x72,
		-14: 0x71,
		-15: 0x70,
		-16: 0x6F,
		-17: 0x6E,
		-18: 0x6D,
		-19: 0x6C,
		-20: 0x6B,
		-21: 0x6A,
		-22: 0x69,
		-23: 0x68,
		-24: 0x67,
		-25: 0x66,
		-26: 0x65,
		-27: 0x64,
		-28: 0x63,
		-29: 0x62,
		-30: 0x61,
		-31: 0x60,
		-32: 0x5F,
		-33: 0x5E,
		-34: 0x5D,
		-35: 0x5C,
		-36: 0x5B,
		-37: 0x5A,
		-38: 0x59,
		-39: 0x58,
		-40: 0x57,
		-41: 0x56,
		-42: 0x55,
		-43: 0x54,
		-44: 0x53,
		-45: 0x52,
		-46: 0x51,
		-47: 0x50,
		-48: 0x4F,
		-49: 0x4E,
		-50: 0x4D,
		-51: 0x4C,
		-52: 0x4B,
		-53: 0x4A,
		-54: 0x49,
		-55: 0x48,
		-56: 0x47,
		-57: 0x46,
		-58: 0x45,
		-59: 0x44,
		-60: 0x43,
		-61: 0x42,
		-62: 0x41,
		-63: 0x40,
	}
}

/**
 * get echo up down data.
 */
func getEchoData() map[int]byte {
	return map[int]byte{
		63: 0x7F,
		62: 0x7E,
		61: 0x7D,
		60: 0x7C,
		59: 0x7B,
		58: 0x7A,
		57: 0x79,
		56: 0x78,
		55: 0x77,
		54: 0x76,
		53: 0x75,
		52: 0x74,
		51: 0x73,
		50: 0x72,
		49: 0x71,
		48: 0x70,
		47: 0x6F,
		46: 0x6E,
		45: 0x6D,
		44: 0x6C,
		43: 0x6B,
		42: 0x6A,
		41: 0x69,
		40: 0x68,
		39: 0x67,
		38: 0x66,
		37: 0x65,
		36: 0x64,
		35: 0x63,
		34: 0x62,
		33: 0x61,
		32: 0x60,
		31: 0x5F,
		30: 0x5E,
		29: 0x5D,
		28: 0x5C,
		27: 0x5B,
		26: 0x5A,
		25: 0x59,
		24: 0x58,
		23: 0x57,
		22: 0x56,
		21: 0x55,
		20: 0x54,
		19: 0x53,
		18: 0x52,
		17: 0x51,
		16: 0x50,
		15: 0x4F,
		14: 0x4E,
		13: 0x4D,
		12: 0x4C,
		11: 0x4B,
		10: 0x4A,
		9:  0x49,
		8:  0x48,
		7:  0x47,
		6:  0x46,
		5:  0x45,
		4:  0x44,
		3:  0x43,
		2:  0x42,
		1:  0x41,
		0:  0x40,
	}
}
