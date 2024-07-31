package main

import (
	"fmt"
	g "gravatar/pkg"
	"os"

	"github.com/jessevdk/go-flags"
)

var (
	options = []g.GravatarOptionsConfiuration{}
)

type CliOptions struct {
	Http         bool   `short:"t" long:"http" help:"Use http over https" env:"USE_HTTP"`
	ForceDefault bool   `short:"f" long:"force-default" description:"Force use default avatar image" env:"FORCE_DEFAULT"`
	DefaultImage string `short:"d" long:"default-image" description:"Specify default avatar image (url or 404, mp, identicon, monsterid, wavatar, retro, robohash, blank)" value-name:"DEFAULT_IMAGE" env:"DEFAULT_IMAGE"`
	Size         int    `short:"s" long:"size" description:"Specifi avatar image size" value-name:"SIZE" env:"SIZE"`
	Rate         string `short:"r" long:"rate" description:"Specify avatar rating" value-name:"RATE" env:"RATE" choice:"g" choice:"pg" choice:"r" choice:"x"`
	Args         struct {
		Email string `positional-arg-name:"MAIL" description:"User email" env:"EMAIL"`
	} `positional-args:"yes" required:"yes"`
}

func main() {
	var cli CliOptions

	parser := flags.NewParser(&cli, flags.Default)
	// parser.Usage = "[OPTIONS] EMAIL"
	_, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	if cli.Http {
		options = append(options, g.UseHttp())
	}
	if cli.ForceDefault {
		options = append(options, g.ForceDefault())
	}
	if cli.DefaultImage != "" {
		options = append(options, g.DefaultImage(cli.DefaultImage))
	}
	if cli.Size > 0 {
		options = append(options, g.Size(cli.Size))
	}
	if cli.Rate != "" {
		options = append(options, g.Rating(cli.Rate))
	}

	link := g.NewGravatar(cli.Args.Email, options...)
	fmt.Printf("%s\n", link)
}
