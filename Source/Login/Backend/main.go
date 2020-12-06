package main

import "dolittle.io/login/cmd"

// func getConfiguration() *configuration.Configuration {
// 	defaultConfig := configuration.GetDefaults()
// 	err := configuration.Setup(&defaultConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	config, err := configuration.Read()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return config
// }

func main() {
	cmd.Execute()
	// config := getConfiguration()
	// hydraAdminURL, _ := url.Parse(config.HydraAdminURL)
	// hydraAdmin := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{
	// 	Schemes:  []string{hydraAdminURL.Scheme},
	// 	Host:     hydraAdminURL.Host,
	// 	BasePath: hydraAdminURL.Path,
	// })

	// kratosPublicURL, _ := url.Parse(config.KratosPublicURL)
	// kratosPublic := kratos.NewHTTPClientWithConfig(nil, &kratos.TransportConfig{
	// 	Schemes:  []string{kratosPublicURL.Scheme},
	// 	Host:     kratosPublicURL.Host,
	// 	BasePath: kratosPublicURL.Path,
	// })

	// base := &handlers.Base{
	// 	HydraClient:  hydraAdmin,
	// 	KratosClient: kratosPublic,
	// }

	// selectedHandler := handlers.SelectedHandler{Base: base}
	// consentHandler := handlers.ConsentHandler{Base: base}

	// http.Handle("/selected-tenant/", &selectedHandler)
	// http.Handle("/consent/", &consentHandler)

	// http.Handle("/", http.FileServer(http.Dir("spa")))

	// log.Println(fmt.Sprintf("Listening on port %d", config.Port))
	// err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
