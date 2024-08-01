package url

import (
	"log"
	u "net/url"
	"strconv"
)

type Protocol = string

const (
	pHttp  = "http"
	pHttps = "https"
)

type AvatarOptions struct {
	Email      string
	Parameters u.Values
	Protocol   Protocol
}

const (
	Default404 = "404"
	// Mystery person
	DefaultMP         = "mp"
	DefaultIdentyIcon = "identicon"
	DefaultMonsterId  = "monsterid"
	DefaultWavater    = "wavatar"
	DefaultRetro      = "retro"
	DefaultRoboHash   = "robohash"
	DefaultBlank      = "blank"
)

type Rate = string

const (
	RatingG  Rate = "g"
	RatingPG Rate = "pg"
	RatingR  Rate = "r"
	RatingX  Rate = "x"
)

type AvatarOptionsConfiuration func(*AvatarOptions)

func DefaultImage(url string) AvatarOptionsConfiuration {
	return func(opt *AvatarOptions) {
		opt.Parameters.Add("default", url)
	}
}

func Size(pixels int) AvatarOptionsConfiuration {
	return func(opt *AvatarOptions) {
		opt.Parameters.Add("size", strconv.Itoa(pixels))
	}
}
func ForceDefault() AvatarOptionsConfiuration {
	return func(opt *AvatarOptions) {
		opt.Parameters.Add("forcedefault", "y")
	}
}
func Rating(rating Rate) AvatarOptionsConfiuration {
	return func(opt *AvatarOptions) {
		opt.Parameters.Add("rating", rating)
	}
}

func UseHttp() AvatarOptionsConfiuration {
	return func(opt *AvatarOptions) {
		opt.Protocol = pHttp
	}
}

const gravatarDefaultUrl = "://gravatar.com/avatar/"

func newGravatarOptions(email string) *AvatarOptions {
	return &AvatarOptions{
		Email:      email,
		Parameters: u.Values{},
		Protocol:   pHttps,
	}
}

func (opt *AvatarOptions) Config(config AvatarOptionsConfiuration) {
	config(opt)
}
func (opt *AvatarOptions) BuildUrl() *u.URL {
	str := opt.Protocol + gravatarDefaultUrl + hashEmail(opt.Email)
	// log.Printf("%s\n", str)
	url, err := u.Parse(str)
	if err != nil {
		log.Fatalf("%x\n", &err)
	}
	url.RawQuery = opt.Parameters.Encode()

	return url
}

func NewAvatarUrl(email string, configs ...AvatarOptionsConfiuration) string {
	opt := newGravatarOptions(email)
	for _, conf := range configs {
		opt.Config(conf)
	}
	return opt.BuildUrl().String()
}
