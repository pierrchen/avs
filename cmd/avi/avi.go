package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pierrchen/avs/images"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "avi"
	app.Version = "0.0.1"
	app.Usage = "Andoid Image Tool"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Please use avi subcommands, see avi -h")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "ramdisk",
			Aliases: []string{"r"},
			Usage:   "regerate the device config",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "image", Value: "", Usage: "ramdisk.img image for unpack"},
				cli.StringFlag{Name: "extract", Value: "avs_ramdisk", Usage: "extracted out dir"},
			},
			Action: func(c *cli.Context) error {
				r := images.Ramdisk{ImagePath: c.String("image")}
				err := r.Unpack(c.String("extract"))
				if err != nil {
					log.Fatalln(err)
				}
				return nil
			},
		},

		{
			Name:    "kernel",
			Aliases: []string{"r"},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "image", Value: "", Usage: "kernel image file"},
			},
			Action: func(c *cli.Context) error {
				r := images.Kernel{ImagePath: c.String("image")}
				v, err := r.Version()
				if err != nil {
					log.Fatalln(err)
				}

				fmt.Println(v)
				// check if it is arm64 image
				arm64 := images.Arm64Image{ImagePath: c.String("image")}
				hdr, err := arm64.Hdr()

				if err == nil {
					fmt.Println("Arm 64 Linux Kernel Image, info")
					fmt.Println(hdr)
				}

				// check if it has something appended
				if arm64.IsSomethingAppended() {
					fmt.Println("something is appended after the kernel")
					ks, _ := arm64.ActualKernelSize()
					fmt.Println("Actualy Kernel Size", ks)

					arm64.Split()

				} else {
					fmt.Println("Pure and simple kernel image")
				}
				return nil
			},
		},

		{
			Name:  "bootimg",
			Usage: "parse the bootimg",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "image", Value: "", Usage: "boot image file"},
				cli.BoolFlag{Name: "extract", Usage: "extract file"},
			},
			Action: func(c *cli.Context) error {
				b := images.Bootimg{ImagePath: c.String("image")}

				// dump the header
				v, err := b.Hdr()
				if err != nil {
					log.Fatalln(err)
				}

				fmt.Println(v)

				if c.Bool("extract") {
					b.Unpack()
				}
				return nil
			},
		},

		{
			Name: "dtb",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "image", Value: "", Usage: "dtb image file"},
				cli.BoolFlag{Name: "hdr", Usage: "dump hdr info"},
				cli.BoolFlag{Name: "dump", Usage: "dump dtb"},
			},
			Action: func(c *cli.Context) error {
				dtb := images.Dtb{ImagePath: c.String("image")}

				// dump the header
				if !dtb.IsDtb() {
					fmt.Println("Not a DTB file")
					return nil
				}

				fmt.Println("Is a DTB file")

				if c.Bool("hdr") {
					fmt.Printf("Header Info:\n %#v\n", dtb.Hdr())
				}

				if c.Bool("dump") {
					dtb.ToDts()
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
