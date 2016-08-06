package app

type Resource map[int]string

var Resources Resource

const (
	GoogleGrantRefreshAccessUrl = iota
	NumberedRpcUrl
	PtcLoginUrl
	PtcLoginOauth
	RpcUrl
)

func init() {
	Resources = Resource{
		RpcUrl:                      "https://pgorelease.nianticlabs.com/plfe/rpc",
		NumberedRpcUrl:              "https://pgorelease.nianticlabs.com/plfe/{0}/rpc",
		PtcLoginUrl:                 "https://sso.pokemon.com/sso/login?service=https%3A%2F%2Fsso.pokemon.com%2Fsso%2Foauth2.0%2FcallbackAuthorize",
		PtcLoginOauth:               "https://sso.pokemon.com/sso/oauth2.0/accessToken",
		GoogleGrantRefreshAccessUrl: "https://android.clients.google.com/auth",
	}
}
