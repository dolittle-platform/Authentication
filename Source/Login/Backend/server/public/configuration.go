package public

type FrontendConfiguration struct {
	ShowDolittleHeadline bool
	AnimateBackground    bool
	ApplicationName      string
	SupportEmail         string
}

type Configuration interface {
	DevMode() bool
	Frontend() FrontendConfiguration
}
