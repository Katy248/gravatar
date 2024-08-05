package main

import (
	"fmt"
	"github.com/katy248/gravatar/pkg/info"
	"github.com/katy248/gravatar/pkg/url"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/dolmen-go/kittyimg"
)

func main() {
	email := "petrovanton247@gmail.com"
	profile := info.FetchData(email)
	kittyPrintImage(url.NewAvatarUrl(email, url.Size(256)))
	fmt.Printf("%s (%s)\n", profile.DisplayName, profile.PreferredUsername)
	fmt.Printf("%s, %s\n", profile.JobTitle, profile.Company)
	fmt.Printf("%s\n", profile.AboutMe)

	for _, acc := range profile.Accounts {
		verified := ""
		if !acc.IsVerified() {
			verified = ""
		}
		fmt.Printf("\t%s %s\n\t%s\n", acc.Name, verified, acc.Url)
	}
	// fmt.Println(profile)
}

func kittyPrintImage(imageSource string) {
	r, err := http.Get(imageSource)

	if err != nil {
		log.Fatal(err)
	}

	img, err := png.Decode(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" ")
	kittyimg.Fprintln(os.Stdout, img)
}
