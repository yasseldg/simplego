package sTemp

import (
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/yasseldg/simplego/sLog"
	"go.mongodb.org/mongo-driver/bson"
)

type Paths []string
type Data bson.M

func Basics(dir_path string) Paths {
	return Paths{
		filepath.Join(dir_path, "templates/basics/header.html"), filepath.Join(dir_path, "templates/basics/message.html"),
		filepath.Join(dir_path, "templates/basics/footer.html"), filepath.Join(dir_path, "templates/basics/layout.html")}
}

func Layout(w http.ResponseWriter, dir_temp_paths, path_prefix string, msg FlashMessages, temp_paths Paths, data Data) error {
	return LayoutWithFunc(w, nil, dir_temp_paths, path_prefix, msg, temp_paths, data)
}

func LayoutWithFunc(w http.ResponseWriter, func_maps template.FuncMap, dir_temp_paths, path_prefix string, msg FlashMessages, temp_paths Paths, data Data) error {
	tmpl, err := template.New("layout.html").Funcs(mergeFunctions(func_maps)).ParseFiles(append(Basics(dir_temp_paths), temp_paths...)...)
	if err != nil {
		sLog.Error(err.Error())
		return err
	}
	sLog.Debug("templates: %v", tmpl.DefinedTemplates())

	data["Messages"] = msg
	data["PathPrefix"] = path_prefix

	err = tmpl.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		sLog.Error(err.Error())
		return err
	}
	return nil
}

func mergeFunctions(func_maps template.FuncMap) template.FuncMap {
	if func_maps == nil {
		return Functions()
	}

	all_functions := Functions()
	for k, v := range func_maps {
		if _, ok := all_functions[k]; !ok {
			all_functions[k] = v
		}
	}
	return all_functions
}
