package osin

// Client information
type Client interface {
	// Client id
	GetId() string

	// Client secret
	GetSecret() string

	// Base client uri
	GetRedirectUri() string

	// get client app name
	GetAppName() string

	// get image url
	GetImageUrl() string

	// External Client
	IsExternal() bool

	// Data to be passed to storage. Not used by the library.
	GetUserData() interface{}
}

// ClientSecretMatcher is an optional interface clients can implement
// which allows them to be the one to determine if a secret matches.
// If a Client implements ClientSecretMatcher, the framework will never call GetSecret
type ClientSecretMatcher interface {
	// SecretMatches returns true if the given secret matches
	ClientSecretMatches(secret string) bool
}

// DefaultClient stores all data in struct variables
type DefaultClient struct {
	Id          string
	Secret      string
	RedirectUri string
	AppName     string
	ImageUrl    string
	External    bool
	UserData    interface{}
}

func (d *DefaultClient) GetId() string {
	return d.Id
}

func (d *DefaultClient) GetSecret() string {
	return d.Secret
}

func (d *DefaultClient) GetRedirectUri() string {
	return d.RedirectUri
}

func (d *DefaultClient) GetAppName() string {
	return d.AppName
}

func (d *DefaultClient) GetImageUrl() string {
	return d.ImageUrl
}

func (d *DefaultClient) IsExternal() bool {
	return d.External
}

func (d *DefaultClient) GetUserData() interface{} {
	return d.UserData
}

// Implement the ClientSecretMatcher interface
func (d *DefaultClient) ClientSecretMatches(secret string) bool {
	return d.Secret == secret
}

func (d *DefaultClient) CopyFrom(client Client) {
	d.Id = client.GetId()
	d.Secret = client.GetSecret()
	d.RedirectUri = client.GetRedirectUri()
	d.AppName = client.GetAppName()
	d.ImageUrl = client.GetImageUrl()
	d.UserData = client.GetUserData()
}
