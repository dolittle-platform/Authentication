package public

type FrontendConfiguration struct {
	ShowDolittleHeadline bool
	ApplicationName      string
	SupportEmail         string
	StartPath            string
	LogoutPath           string
}

type Configuration interface {
	DevMode() bool
	Frontend() FrontendConfiguration
}
