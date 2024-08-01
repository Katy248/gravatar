package info

import "log"

type Response struct {
	Entry []ProfileInfo `json:"entry"`
}

func (r *Response) GetInfo() ProfileInfo {
	if len(r.Entry) <= 0 {
		log.Fatal("Bad response, no entry")
	}

	return r.Entry[0]
}

type ProfileInfo struct {
	Hash              string            `json:"hash"`
	RequestHash       string            `json:"requestHash"`
	ProfileUrl        string            `json:"proileUrl"`
	PreferredUsername string            `json:"preferredUsername"`
	ThumbnailUrl      string            `json:"thumbnailUrl"`
	Photos            []Photo           `json:"photos"`
	DisplayName       string            `json:"displayName"`
	AboutMe           string            `json:"aboutMe"`
	JobTitle          string            `json:"job_title"`
	Company           string            `json:"company"`
	Emails            []Email           `json:"emails"`
	Accounts          []Account         `json:"accounts"`
	ProfileBackground ProfileBackground `json:"profileBackground"`
}

type Photo struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Email struct {
	Primary string `json:"primary"`
	Value   string `json:"value"`
}

func (e *Email) IsPrimary() bool {
	return e.Primary == "true"
}

type Account struct {
	Domain    string `json:"domain"`
	Display   string `json:"display"`
	Url       string `json:"url"`
	IconUrl   string `json:"iconUrl"`
	Username  string `json:"username"`
	Verified  string `json:"verified"`
	Name      string `json:"name"`
	Shortname string `json:"shortname"`
}

func (a *Account) IsVerified() bool {
	return a.Verified == "true"
}

type ProfileBackground struct {
	Color        string  `json:"color"`
	Opacity      float32 `json:"opacity"`
	Primarycolor string  `json:"primaryColor"`
}
