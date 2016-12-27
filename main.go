package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"

	"gopkg.in/urfave/cli.v1"
)

type (
	// Content struct
	Content struct {
		Groups   []Group
		Sections []Section
	}
	// Section struct
	Section struct {
		Title   string
		Content template.HTML
		Example template.HTML
		Link    string
	}
	// Group struct
	Group struct {
		Title string
		Urls  []URL
	}
	// URL struct
	URL struct {
		Href  string
		Title string
	}
)

func main() {

	app := cli.NewApp()
	app.Name = "docsgen"
	app.Usage = "Generate docs file from markdown"
	app.Version = "0.0.1"

	// CLI Flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "theme",
			Value: "default",
			Usage: "Theme name",
		},
		cli.StringFlag{
			Name:  "out",
			Value: "docs",
			Usage: "Output folder",
		},
		cli.StringFlag{
			Name:  "in",
			Value: "api_docs",
			Usage: "Input folder",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Generates the documents from markdown",
			Action: func(c *cli.Context) error {

				// Get variables
				theme := c.GlobalString("theme")
				outputFolder := c.GlobalString("out")
				inputFolder := c.GlobalString("in")

				// Clean path
				os.RemoveAll(outputFolder)
				os.MkdirAll(outputFolder, 0777)

				// Get template file
				data, err := Asset("themes/" + theme + "/index.html")
				if err != nil {
					return errors.New("Asset not found")
				}

				// Open template
				tmpl := template.New("template")
				tmpl, err = tmpl.Parse(string(data))
				if err != nil {
					return errors.New("Theme does not exist")
				}

				// Create file for storing the result
				fs, err := os.Create(outputFolder + "/index.html")
				if err != nil {
					return errors.New("Could not create docs file")
				}
				defer fs.Close()

				// Get groups & sections
				directories, _ := ioutil.ReadDir(inputFolder)
				groups := []Group{}
				sections := []Section{}
				for _, dir := range directories {
					var urls []URL
					outerPath := filepath.Join(inputFolder, dir.Name())
					innerDirs, _ := ioutil.ReadDir(outerPath)
					for _, innerDir := range innerDirs {
						// Url
						url := URL{
							Title: createNameFromFile(innerDir.Name(), " "),
							Href:  createNameFromFile(innerDir.Name(), "-"),
						}
						urls = append(urls, url)

						// Section
						var content template.HTML
						var example template.HTML

						path := filepath.Join(outerPath, innerDir.Name())
						files, _ := filepath.Glob(filepath.Join(path, "*.md"))
						for _, f := range files {
							b, err := ioutil.ReadFile(f)
							if err != nil {
								return errors.New("Could not read file")
							}

							// Parse markdown
							output := blackfriday.MarkdownCommon(b)

							if strings.Index(f, "example") > -1 {
								example = template.HTML(output)
							} else {
								content = template.HTML(output)
							}
						}

						section := Section{
							Title:   url.Title,
							Link:    url.Href,
							Content: content,
							Example: example,
						}
						sections = append(sections, section)
					}
					if len(urls) != 0 {
						var group = Group{
							Title: createNameFromFile(dir.Name(), " "),
							Urls:  urls,
						}
						groups = append(groups, group)
					}
				}

				// Create docs file content
				content := Content{
					Sections: sections,
					Groups:   groups,
				}

				// Write template to file
				err = tmpl.Execute(fs, content)
				if err != nil {
					return errors.New("Could not write template to file")
				}

				// Create theme file
				data, err = Asset("themes/" + theme + "/style.css")
				if err != nil {
					return errors.New("Asset not found")
				}

				fsTheme, err := os.Create(outputFolder + "/style.css")
				if err != nil {
					return errors.New("Could not create docs file")
				}
				defer fsTheme.Close()
				fsTheme.Write(data)

				// Return nil
				return nil
			},
		},
	}

	app.Run(os.Args)

}

func createNameFromFile(path string, join string) string {
	splittedName := strings.Split(path, "_")
	finalName := append(splittedName[:0], splittedName[0+1:]...)
	return strings.Join(finalName[:], join)
}
