package iris_extend_view

import (
	helper "github.com/AlexZ33/iris-extend-helper"
	server "github.com/AlexZ33/iris-extend-server"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
)

var (
	Template view.Engine
	Viewpath string
)

func init() {
	config := server.Config
	enable := helper.GetBool(config, "view.enable")
	if enable {
		path := getViewPath("")
		mode := server.IsLocal()
		engine := helper.GetString(config, "view.engine", "pug")
		switch engine {
		case "pug":
			Template = iris.Pug(path, ".pug").Reload(mode)
		case "jade":
			Template = iris.Pug(path, ".jade").Reload(mode)
		case "amber":
			Template = iris.Amber(path, ".amber").Reload(mode)
		case "django":
			Template = iris.Django(path, ".html").Reload(mode)
		case "handlebars":
			Template = iris.Handlebars(path, ".html").Reload(mode)
		case "jet":
			Template = iris.Jet(path, ".jet").Reload(mode)
		default:
			Template = iris.HTML(path, ".html").Reload(mode)
		}
	}
}

func getViewPath(filename string) string {
	if filename != "" {
		Viewpath = filename
	} else {
		Viewpath = "./app/view"
	}
	return Viewpath
}
