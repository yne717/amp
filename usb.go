package main

import (
	"fmt"
	"os"

	"github.com/kylelemons/gousb/usb"
)

func main() {
	ctx := usb.NewContext()
	defer ctx.Close()

	devs, err := ctx.ListDevices(func(desc *usb.Descriptor) bool {
		return true
	})

	fmt.Print(devs)

	defer func() {
		for _, dev := range devs {
			dev.Close()
		}
	}()

	if err != nil {
		die("エラーが発生しました")
	}

	for _, dev := range devs {
		i, err := dev.ActiveConfig()
		if err != nil {
			die("エラーが発生しました")
		}
		fmt.Println(i)
	}

}

func die(message string) {
	fmt.Println(message)
	os.Exit(0)
}
