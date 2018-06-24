package templator

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

// Templator produces html
// tPath - template path
// jData - json data
func Templator(tPath string, jData map[string]interface{}) (string, error) {
	//fmt.Println("tPath: " + tPath + " jData: " + jData

	htmlFile, err := ioutil.ReadFile(tPath)
	if err != nil {
		return "Cannot read the html template file", err
	}

	s := string(htmlFile)
	t := template.Must(template.New("").Parse(s))
	var b bytes.Buffer
	t.Execute(&b, jData)
	//r := string(m)
	//r := StringIsh(m)
	return b.String(), nil
}
