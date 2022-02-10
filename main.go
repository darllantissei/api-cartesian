package main

import (
	"flag"
	"strconv"

	webserver "github.com/darllantissei/api-cartesian/adapters/web"
	"github.com/darllantissei/api-cartesian/application"
	"github.com/darllantissei/api-cartesian/application/authentication"
	"github.com/darllantissei/api-cartesian/application/coordinate"
	"github.com/darllantissei/api-cartesian/application/utils"
	"github.com/eucatur/go-toolbox/slice"
)

const instructions = `
	Tipo do serviço não definido. Tipos válidos: 
	http: Irá iniciar a aplicação no modo web, servindo recursos de API
	Exemplo: Ao inicia a aplicação defina a flag service=<tipo-servico>. Assim: -service=http
	`

const flagService = "service"

func init() {
	flag.String(flagService, "http", instructions)
}

func main() {

	port := getPortHTTP()
	serviceType := getServiceType()

	utilsService := utils.UtilsService{}

	authenticationService := authentication.AuthenticationService{
		PersistenceDB:    nil,
		PersistenceCache: nil,
		Utils:            &utilsService,
	}

	coordinateService := coordinate.CoordinateService{}

	application := application.Application{
		AuthenticationService: &authenticationService,
		CoordinanteService:    &coordinateService,
		Utils:                 &utilsService,
	}

	switch serviceType {

	case "http":

		serverWeb := webserver.MakeNewWebServer(application)

		serverWeb.Serve(port)

	default:
		panic("é necessário definir o tipo de serviço que a aplicação irá executar")

	}

}

func getPortHTTP() (port int) {
	flag.Parse()
	flagPort := flag.Lookup("port")

	if flagPort != nil {
		portEnv, err := strconv.Atoi(flagPort.Value.String())

		if err != nil {
			port = 9000
		}

		port = portEnv
	}

	return
}

func getServiceType() (serviceType string) {

	flag.Parse()

	flagServiceType := flag.Lookup(flagService)

	if flagServiceType == nil {
		panic(instructions)
	}

	flagAllowed := []string{
		"http",
		"agent",
	}

	if !slice.SliceExists(flagAllowed, flagServiceType.Value.String()) {
		panic(instructions)
	}

	serviceType = flagServiceType.Value.String()

	return

}
