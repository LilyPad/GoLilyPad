package proxy

type Localizer interface {
	LocaleFull() string
	LocaleOffline() string
	LocaleLoggedIn() string
	LocaleLostConn() string
	LocaleShutdown() string
}