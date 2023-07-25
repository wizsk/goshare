package main

import (
	"log"
	"net/http"
	"strings"
)

// cli
type CliUiData struct {
	Root string
	Dirs []Directory
}

const cliUiTemplate = `{{range .Dirs}}
name::{{.Name}} type::{{if .IsDir}}Dir{{else}}File{{end}} url::http://{{$.Root}}{{.Url}}{{end}}
`

func cliUi(w http.ResponseWriter, r *http.Request, pass string) {
	var err error
	var data CliUiData
	data.Root = r.Host

	data.Dirs, err = file(w, r)
	if err != nil {
		return
	}
	indexTemplate.ExecuteTemplate(w, "cli", data)
}

// form page
type FormPageDatas struct {
	RedirectURL string
}

func allQueries(r *http.Request) string {
	params := r.URL.Query()
	if len(params) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.WriteString("?")
	// Loop through all the parameters and print them.
	for key, values := range params {
		for _, value := range values {
			builder.WriteString(key + "=" + value + "&")
		}
	}

	str := builder.String()
	if len(str) > 1 {
		return str[:len(str)-1]
	}

	return ""
}

// web ui

type WebUiData struct {
	Dirtitle     string
	PreviousPage string
	Directories  []Directory
	ProgessPah   []ProgessPah
	SortOptions  []SortOption
}

type SortOption struct {
	Title    string
	Name     string
	Selected bool
}

type ProgessPah struct {
	Title    string
	Url      string
	SlashPre bool
}

func ServeWebUi(w http.ResponseWriter, r *http.Request) {
	var err error
	var datas WebUiData
	datas.Directories, err = file(w, r)
	if err != nil {
		return
	}

	// web ui part
	if split := strings.Split(r.URL.EscapedPath(), "/"); len(split) > 2 {
		datas.PreviousPage = strings.Join(split[:len(split)-1], "/")
	} else {
		datas.PreviousPage = "/"
	}

	datas.ProgessPah = possiblePahts(r, allQueries(r))
	datas.Dirtitle = datas.ProgessPah[len(datas.ProgessPah)-1].Title
	// "namedesc":
	// 		sort.Slice(d, func(i, j int) bool {
	// 			return d[i].Name > d[j].Name
	// 		})
	// 	case "sizeasc":
	// 		sort.Slice(d, func(i, j int) bool {
	// 			return d[i].SizeBytes < d[j].SizeBytes
	// 		})
	// 	case "sizedesc":
	// 		sort.Slice(d, func(i, j int) bool {
	// 			return d[i].SizeBytes > d[j].SizeBytes
	// 		})
	// 	default:
	// "nameasc"
	datas.SortOptions = []SortOption{
		{
			Title: "Letter asn",
			Name:  "nameasc",
		},
		{
			Title: "Letter desc",
			Name:  "namedesc",
		},
		{
			Title: "size soto to boto",
			Name:  "sizeasc",
		},
		{
			Title: "size boto to soto",
			Name:  "sizedesc",
		},
		{
			Title: "by dir",
			Name:  "bydir",
		},
		{
			Title: "by file",
			Name:  "byfile",
		},
	}
	if s := r.FormValue("sort"); s != "" {
		for i, options := range datas.SortOptions {
			if options.Name == s {
				datas.SortOptions[i].Selected = true
				break
			}
		}
	}

	w.Header().Set("Content-Type", "text/html")
	if indexTemplate.ExecuteTemplate(w, "main", datas) != nil {
		log.Println(err)
		http.Error(w, "someting went wrong", http.StatusInternalServerError)
		return
	}
}
