package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	g "gravatar/pkg/url"
)

type CliOptions struct {
}

type AvatarCommand struct {
	Http         bool   `short:"t" long:"http" help:"Use http over https" env:"USE_HTTP"`
	ForceDefault bool   `short:"f" long:"force-default" description:"Force use default avatar image" env:"FORCE_DEFAULT"`
	DefaultImage string `short:"d" long:"default-image" description:"Specify default avatar image (url or 404, mp, identicon, monsterid, wavatar, retro, robohash, blank)" value-name:"DEFAULT_IMAGE" env:"DEFAULT_IMAGE"`
	Size         int    `short:"s" long:"size" description:"Specifi avatar image size" value-name:"SIZE" env:"SIZE"`
	Rate         string `short:"r" long:"rate" description:"Specify avatar rating" value-name:"RATE" env:"RATE" choice:"g" choice:"pg" choice:"r" choice:"x"`
	Args         struct {
		Email string `positional-arg-name:"MAIL" description:"User email" env:"EMAIL"`
	} `positional-args:"yes" required:"yes"`
}

func (cmd *AvatarCommand) Execute(args []string) error {

	options := []g.AvatarOptionsConfiuration{}

	if cmd.Http {
		options = append(options, g.UseHttp())
	}
	if cmd.ForceDefault {
		options = append(options, g.ForceDefault())
	}
	if cmd.DefaultImage != "" {
		options = append(options, g.DefaultImage(cmd.DefaultImage))
	}
	if cmd.Size > 0 {
		options = append(options, g.Size(cmd.Size))
	}
	if cmd.Rate != "" {
		options = append(options, g.Rating(cmd.Rate))
	}

	link := g.NewAvatarUrl(cmd.Args.Email, options...)
	fmt.Printf("%s\n", link)
	return nil
}

type QrCommand struct {
	UseUsername bool   `short:"u" long:"username" description:"Use usernme instead of email hash" env:"USE_USERNAME"`
	Size        int    `short:"s" long:"size" description:"Specify size of QR image" value-name:"SIZE" env:"SIZE"`
	Type        string `short:"t" long:"type" description:"Specify what to display in the center of QR code" choice:"default" choice:"user" choice:"gravatar" default:"default" value-name:"TYPE" env:"TYPE"`
	Version     int    `short:"v" long:"version" description:"Specify style of QR code" choice:"1" choice:"3" default:"1" value-name:"VERSION" env:"VERSION"`
	Args        struct {
		Identifier string `positional-arg-name:"IDENTIFIER" description:"User email or username"`
	} `positional-args:"yes" required:"yes"`
}

func (cmd *QrCommand) Execute(args []string) error {
	options := []g.QrOptionsConfiguration{}

	if cmd.UseUsername {
		options = append(options, g.UseUsername())
	}
	if cmd.Size > 0 {
		options = append(options, g.QrSize(cmd.Size))
	}
	if cmd.Type != "" {
		options = append(options, g.QrDisplayType(cmd.Type))
	}
	if cmd.Version > 0 {
		options = append(options, g.Version(cmd.Version))
	}
	link := g.NewQrCodeUrl(cmd.Args.Identifier, options...)
	fmt.Println(link)
	return nil
}

type InfoCommand struct {
	Format       g.DataFormat `short:"f" long:"format" choice:"json" description:"Specify information format" choice:"xml" choice:"php" choice:"vcf" default:"json" value-name:"FORMAT" env:"FORMAT"`
	JsonCallback string       `short:"c" long:"callback" description:"Specify callback (only for JSON format)" value-name:"CALLBACK" env:"CALLBACK"`
	Args         struct {
		Email string `positional-arg-name:"EMAIL" description:"User email"`
	} `positional-args:"yes" required:"yes"`
}

func (cmd *InfoCommand) Execute(args []string) error {
	options := []g.InfoOptionsConfiguration{}

	if cmd.Format != "" {
		options = append(options, g.InfoFormat(cmd.Format))
	}
	if cmd.JsonCallback != "" {
		options = append(options, g.JsonCallback(cmd.JsonCallback))
	}
	link := g.NewInfoUrl(cmd.Args.Email, options...)
	fmt.Println(link)
	return nil
}

func NewParser() *flags.Parser {
	parser := flags.NewParser(&CliOptions{}, flags.Default)
	parser.AddCommand("avatar", "Generate avatar url", "", &AvatarCommand{})
	parser.AddCommand("info", "Generate user info url", "", &InfoCommand{})
	parser.AddCommand("qr", "Generate qr url", "", &QrCommand{})
	return parser
}
func main() {
	parser := NewParser()
	_, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}
}
