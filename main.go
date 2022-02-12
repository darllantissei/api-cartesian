package main

import (
	"flag"
	"fmt"
	"strconv"

	dbfile "github.com/darllantissei/api-cartesian/adapters/db/file/points"
	webserver "github.com/darllantissei/api-cartesian/adapters/web"
	"github.com/darllantissei/api-cartesian/application"
	"github.com/darllantissei/api-cartesian/application/coordinate"
	"github.com/darllantissei/api-cartesian/application/utils"
)

const instructions = `
Service type undefined. Valid type:
http: Will launch the application in web mode, serving API resources
Example: When starting the application, set the flag service=<service-type>. So: -service=http
	`

const flagService = "service"
const flagPort = "port"
const flagDebug = "debug"

func init() {
	flag.String(flagService, "http", instructions)
	flag.Int(flagPort, int(9000), "port to server HTTP")
	flag.Bool(flagDebug, true, "Show debug")
}

func main() {

	port := getPortHTTP()
	isDebug := getDebug()
	serviceType := getServiceType()

	utilsService := utils.UtilsService{}

	sourcePoints := dbfile.NewPointsCache(dbfile.PointsSource{
		Utils:      &utilsService,
		FilePoints: "points.json",
	})

	_, err := sourcePoints.ListPoints()

	if err != nil {
		panic("unable load list of points")
	}

	coordinateService := coordinate.CoordinateService{
		PersisenceFile: sourcePoints,
	}

	application := application.Application{
		CoordinanteService: &coordinateService,
		Utils:              &utilsService,
	}

	switch serviceType {

	case "http":

		serverWeb := webserver.MakeNewWebServer(application)

		serverWeb.Serve(port, isDebug)

	default:
		panic(fmt.Sprintf("%s type service is invalid. The service type allowed is http only\n\n%s", serviceType, instructions))

	}

}

func getPortHTTP() (port int) {
	flag.Parse()
	flagPort := flag.Lookup("port")

	if flagPort == nil {
		return 9000
	}

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
		return "http"
	}

	serviceType = flagServiceType.Value.String()

	return

}

func getDebug() (debug bool) {
	flag.Parse()
	flagDebug := flag.Lookup("debug")

	if flagDebug == nil {
		return true
	}

	if flagDebug != nil {
		debugEnv, err := strconv.ParseBool(flagDebug.Value.String())

		if err != nil {
			debug = false
		}

		debug = debugEnv
	}

	return
}
