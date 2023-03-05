package creational

type AppType int

const (
	Cli AppType = iota
	Gui
)

type application struct {
	name    string
	version string
	author  string
	appType AppType
}

type Builder interface {
	SetName(string) Builder
	SetVersion(string) Builder
	SetAuthor(string) Builder
	SetAppType(AppType) Builder
	Build() application
}

type CliAppBuilder struct {
	app application
}

func (c *CliAppBuilder) SetName(name string) Builder {
	c.app.name = name
	return c
}

func (c *CliAppBuilder) SetVersion(version string) Builder {
	c.app.version = version
	return c
}

func (c *CliAppBuilder) SetAuthor(author string) Builder {
	c.app.author = author
	return c
}

func (c *CliAppBuilder) SetAppType(appType AppType) Builder {
	c.app.appType = appType
	return c
}

func (c *CliAppBuilder) Build() application {
	return c.app
}

type ApplicationBuilder struct {
	builder Builder
}

func (a *ApplicationBuilder) SetBuilder(b Builder) {
	a.builder = b
}

func (a *ApplicationBuilder) Build() application {
	return a.builder.Build()
}
