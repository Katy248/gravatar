package url

import (
	"log"
	"net/url"
	"strconv"
)

type DataFormat = string

const (
	JsonFormat DataFormat = "json"
	XmlFormat  DataFormat = "xml"
	PhpFormat  DataFormat = "php"
	VcfFormat  DataFormat = "vcf"

	QrCodeFormat DataFormat = "qr"

	DefaultFormat = JsonFormat
)

type InfoOptions struct {
	Email      string
	DataFormat DataFormat
	Parameters url.Values
	// If true use email hash, otherwise use email as username
	UseEmail bool
}

func newInfoOptions(email string) *InfoOptions {
	return &InfoOptions{
		Email:      email,
		DataFormat: DefaultFormat,
		Parameters: url.Values{},
		UseEmail:   true,
	}
}

type InfoOptionsConfiguration func(*InfoOptions)

func (opt *InfoOptions) Config(conf InfoOptionsConfiguration) {
	conf(opt)
}

func (opt *InfoOptions) BuildUrl() *url.URL {
	var identifier string
	if opt.UseEmail {
		identifier = hashEmail(opt.Email)
	} else {
		identifier = opt.Email
	}

	u, err := url.Parse(InfoBaseUrl + identifier + "." + opt.DataFormat)

	if err != nil {
		log.Fatal(err)
	}

	u.RawQuery = opt.Parameters.Encode()

	return u
}

func InfoFormat(format DataFormat) InfoOptionsConfiguration {
	return func(opt *InfoOptions) {
		opt.DataFormat = format
	}
}

// Add callback only if specified format is JSON
func JsonCallback(callback string) InfoOptionsConfiguration {
	return func(opt *InfoOptions) {
		if opt.DataFormat != JsonFormat {
			return
		}

		opt.Parameters.Add("callback", callback)
	}
}

const InfoBaseUrl = "https://gravatar.com/"

func NewInfoUrl(email string, configs ...InfoOptionsConfiguration) string {
	options := newInfoOptions(email)

	for _, config := range configs {
		options.Config(config)
	}

	return options.BuildUrl().String()
}

type QrOptionsConfiguration = InfoOptionsConfiguration

func QrSize(size int) QrOptionsConfiguration {
	return func(opt *InfoOptions) {
		opt.Parameters.Add("size", strconv.Itoa(size))
	}
}

type QrType = string

const (
	QrTypeUserAvatar   QrType = "user"
	QrTypeGravatarLogo QrType = "gravatar"
	QrTypeDefault      QrType = "default"
)

type QrVersion = int

const (
	QrVersion1       QrVersion = 1
	QrVersion3       QrVersion = 3
	QrVersionDefault           = QrVersion1
)

func QrDisplayType(t QrType) QrOptionsConfiguration {
	return func(opt *InfoOptions) {
		opt.Parameters.Add("type", t)
	}
}

func Version(v QrVersion) QrOptionsConfiguration {
	return func(opt *InfoOptions) {
		opt.Parameters.Add("version", strconv.Itoa(v))
	}
}

func UseUsername() QrOptionsConfiguration {
	return func(opt *InfoOptions) {
		opt.UseEmail = false
	}
}

func NewQrCodeUrl(email string, configs ...QrOptionsConfiguration) string {
	options := newInfoOptions(email)

	options.Config(InfoFormat(QrCodeFormat))
	for _, config := range configs {
		options.Config(config)
	}

	return options.BuildUrl().String()

}
