package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	u "net/url"
	"strconv"
	"strings"
)

type Protocol = string

const (
	pHttp  = "http"
	pHttps = "https"
)

type GravatarOptions struct {
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

type GravatarOptionsConfiuration func(*GravatarOptions)

func DefaultImage(url string) GravatarOptionsConfiuration {
	return func(opt *GravatarOptions) {
		opt.Parameters.Add("default", url)
	}
}

func Size(pixels int) GravatarOptionsConfiuration {
	return func(opt *GravatarOptions) {
		opt.Parameters.Add("size", strconv.Itoa(pixels))
	}
}
func ForceDefault() GravatarOptionsConfiuration {
	return func(opt *GravatarOptions) {
		opt.Parameters.Add("forcedefault", "y")
	}
}
func Rating(rating Rate) GravatarOptionsConfiuration {
	return func(opt *GravatarOptions) {
		opt.Parameters.Add("rating", rating)
	}
}

func UseHttp() GravatarOptionsConfiuration {
	return func(opt *GravatarOptions) {
		opt.Protocol = pHttp
	}
}

const gravatarDefaultUrl = "://gravatar.com/avatar/"

func newGravatarOptions(email string) *GravatarOptions {
	return &GravatarOptions{
		Email:      email,
		Parameters: u.Values{},
		Protocol:   pHttps,
	}
}

func (opt *GravatarOptions) Config(config GravatarOptionsConfiuration) {
	config(opt)
}
func (opt *GravatarOptions) BuildUrl() *u.URL {
	str := opt.Protocol + gravatarDefaultUrl + hashEmail(opt.Email)
	// log.Printf("%s\n", str)
	url, err := u.Parse(str)
	if err != nil {
		log.Fatalf("%x\n", &err)
	}
	url.RawQuery = opt.Parameters.Encode()

	return url
}

func NewGravatar(email string, configs ...GravatarOptionsConfiuration) string {
	opt := newGravatarOptions(email)
	for _, conf := range configs {
		opt.Config(conf)
	}
	return opt.BuildUrl().String()
}

func hashEmail(email string) string {
	email = strings.ToLower(email)

	inputData := []byte(email)
	outputData := sha256.Sum256(inputData)
	hash := hex.EncodeToString(outputData[:])
	// log.Printf("%s\n", hash)
	return hash

}
