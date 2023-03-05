package creational

import (
	"testing"
)

func TestApplicationBuilder(t *testing.T) {
	cliBuilder := CliAppBuilder{}
	cliBuilder.SetName("myApp")
	cliBuilder.SetVersion("1.0.0")
	cliBuilder.SetAuthor("gilberto.plaza")
	cliBuilder.SetAppType(Cli)

	appBuilder := ApplicationBuilder{}
	appBuilder.SetBuilder(&cliBuilder)

	app1 := appBuilder.Build()
	app2 := appBuilder.Build()

	if &app1 == &app2 {
		t.Errorf("app objects should be different")
	}

}
