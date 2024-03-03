package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/heydp/WeatherReport/internal/weather"
)

func main() {
	os.Exit(mainerr())
}

func mainerr() int {
	errChan := make(chan error)
	go func(errchan chan error) {
		err := Run()
		if err != nil {
			errChan <- err
		}
	}(errChan)

	select {
	case err := <-errChan:
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}
}

func Run() error {
	var err error
	fmt.Println("starting the weather app")

	srv := NewWebServer()

	env := GetAppEnv()
	configPath := fmt.Sprintf("configs/%s.json", env)

	file, err := os.ReadFile(configPath)
	if err != nil {
		errMsg := fmt.Sprintf("error in reading configFile, err - %v", err.Error())
		fmt.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	var appConfig AppConfig
	err = json.Unmarshal(file, &appConfig)
	if err != nil {
		errMsg := fmt.Sprintf("error in unmarshalling configFile, err - %v", err.Error())
		fmt.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	ctrls, ictrls := setUpRoutes()

	var serverResources ServerResources
	serverResources.ctrls = ctrls
	serverResources.itcrls = ictrls

	srv.InitRouter(&serverResources)
	fmt.Println("initializing servers")

	srv.initialized = true
	err = srv.Start(appConfig)
	if err != nil {
		return err
	}

	return nil
}

func setUpRoutes() ([]Controllers, []Controllers) {

	var controller []Controllers
	var icontroller []Controllers

	weatherManager := weather.NewManager()
	weatherController := weather.NewController(weatherManager)
	controller = append(controller, weatherController)

	return controller, icontroller
}
