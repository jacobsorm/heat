package heat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL = "https://api.myuplink.com"
)

type ErrorUnauthorized struct{}

func (e *ErrorUnauthorized) Error() string {
	return "Unauthorized"
}

type ErrorResponse struct {
	code    int
	message string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}

type AcccessRefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type OAuthResponse struct {
	AcccessRefreshToken
	ExpiresInSeconds int    `json:"expires_in"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
}

type System struct {
	Id            string         `json:"systemId"`
	Name          string         `json:"name"`
	SecurityLevel string         `json:"securityLevel"`
	HasAlarm      bool           `json:"hasAlarm"`
	Country       string         `json:"country"`
	Devices       []SystemDevice `json:"devices"`
}

type SystemsResponse struct {
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	NumItems     int      `json:"numItems"`
	Systems      []System `json:"systems"`
}

type SystemDevice struct {
	Id               string `json:"id"`
	ConnectionState  string `json:"connectionState"`
	CurrentFwVersion string `json:"currentFwVersion"`
	DeviceType       string `json:"deviceType"`
	Product          struct {
		SerialNumber string `json:"serialNumber"`
		Name         string `json:"name"`
	} `json:"product"`
}

type NotificationStatusHistory struct {
	Status   string `json:"status"`
	Datetime int    `json:"datetime"`
}

type NotificationTimestamp struct {
	time.Time
}

func (t *NotificationTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		return nil
	}

	tz := s[len(s)-1]
	trimmed := s[:len(s)-2]
	parsed, err := time.Parse(time.RFC3339, fmt.Sprintf("%s+0%c:00", trimmed, tz))
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}

type Notifications struct {
	Id              string                      `json:"id"`
	AlarmNumber     int                         `json:"alarmNumber"`
	DeviceId        string                      `json:"deviceId"`
	Severity        int                         `json:"severity"`
	Status          string                      `json:"status"`
	CreatedDateTime NotificationTimestamp       `json:"createdDateTime"`
	StatusHistory   []NotificationStatusHistory `json:"statusHistory"`
	Header          string                      `json:"header"`
	Description     string                      `json:"description"`
	EquipName       string                      `json:"equipName"`
}
type NotificationsResponse struct {
	Page          int             `json:"page"`
	ItemsPerPage  int             `json:"itemsPerPage"`
	NumItems      int             `json:"numItems"`
	Notifications []Notifications `json:"notifications"`
}

func GetNotifications(systemId, accessToken string) (NotificationsResponse, error) {
	n := NotificationsResponse{}
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/systems/%s/notifications", baseURL, systemId), nil)
	if err != nil {
		return n, err
	}
	r.Header.Add("Authorization", "Bearer "+accessToken)
	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return n, err
	}
	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return n, err
	}
	if resp.StatusCode != http.StatusOK {
		return n, fmt.Errorf("non 200 status: %s", resp.Status)
	}
	if len(bod) == 0 {
		return n, nil
	}
	return n, json.Unmarshal(bod, &n)
}

func GetNotificationsReauth(systemId, clientId, clientSecret string, oauth OAuthResponse) (NotificationsResponse, OAuthResponse, error) {
	n, err := GetNotifications(systemId, oauth.AccessToken)
	if errors.Is(err, &ErrorUnauthorized{}) {
		oauth, err = OAuthRefreshToken(clientId, clientSecret, oauth.RefreshToken)
		if err != nil {
			return NotificationsResponse{}, OAuthResponse{}, err
		}
		n, err = GetNotifications(systemId, oauth.AccessToken)
	}
	return n, oauth, err
}

func GetSystemsReauth(clientId, clientSecret string, oauth OAuthResponse) (SystemsResponse, OAuthResponse, error) {
	r, err := GetSystems(oauth.AccessToken)
	if errors.Is(err, &ErrorUnauthorized{}) {
		if oauth, err = OAuthRefreshToken(clientId, clientSecret, oauth.RefreshToken); err != nil {
			return SystemsResponse{}, OAuthResponse{}, err
		} else {
			r, err = GetSystems(oauth.AccessToken)
		}
	}
	return r, oauth, err
}

func GetDevice(deviceId, accessToken string) (DeviceResponse, error) {
	d := DeviceResponse{}
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/devices/%s", baseURL, deviceId), nil)
	if err != nil {
		return d, err
	}
	r.Header.Add("Authorization", "Bearer "+accessToken)
	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return d, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return DeviceResponse{}, &ErrorUnauthorized{}
	}
	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}
	if resp.StatusCode != http.StatusOK {
		return DeviceResponse{}, &ErrorResponse{
			code:    resp.StatusCode,
			message: string(bod),
		}
	}
	fmt.Println(string(bod))
	return d, json.Unmarshal(bod, &d)
}

func GetSystems(accessToken string) (SystemsResponse, error) {
	s := SystemsResponse{}
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/systems/me", baseURL), nil)
	if err != nil {
		return s, err
	}
	r.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return SystemsResponse{}, &ErrorUnauthorized{}
	} else if resp.StatusCode != http.StatusOK {
		return SystemsResponse{}, fmt.Errorf("Error in systems request (%d): %q", resp.StatusCode, string(bod))
	}
	return s, json.Unmarshal(bod, &s)
}

func OAuth(grantType, clientId, clientSecret string, values url.Values) (OAuthResponse, error) {
	oauth := OAuthResponse{}
	postForm := url.Values{
		"grant_type":    []string{grantType}, // refresh_token, authorization_code
		"client_id":     []string{clientId},
		"client_secret": []string{clientSecret},
	}
	for k, v := range values {
		postForm[k] = v
	}
	resp, err := http.PostForm(fmt.Sprintf("%s/oauth/token", baseURL), postForm)
	if err != nil {
		return oauth, err
	}
	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body)
	if err != nil {
		return oauth, err
	}
	if resp.StatusCode != 200 {
		return oauth, fmt.Errorf("Status code %d for access token response: %q", resp.StatusCode, string(bod))
	}
	return oauth, json.Unmarshal(bod, &oauth)
}

func OAuthRefreshToken(clientId, clientSecret, refreshToken string) (OAuthResponse, error) {
	return OAuth("refresh_token", clientId, clientSecret, url.Values{
		"refresh_token": []string{refreshToken},
	})
}
