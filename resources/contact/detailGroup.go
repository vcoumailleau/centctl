/*
MIT License

Copyright (c)  2020-2021 YPSI SAS
Centctl is developped by : Mélissa Bertin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package contact

import (
	"centctl/resources"
	"encoding/json"
	"fmt"

	"github.com/jszwec/csvutil"
	"gopkg.in/yaml.v2"
)

//DetailGroup represents the caracteristics of a contact group
type DetailGroup struct {
	Name     string              `json:"name" yaml:"name"` //group Name
	ID       string              `json:"id" yaml:"id"`     //group ID
	Alias    string              `json:"alias" yaml:"alias"`
	Contacts DetailGroupContacts `json:"contacts" yaml:"contacts"`
}

type DetailGroupContacts []DetailGroupContact

func (t DetailGroupContacts) MarshalCSV() ([]byte, error) {
	var value string
	for i, parent := range t {
		value += parent.ID + "|" + parent.Name
		if i < len(t)-1 {
			value += ","
		}
	}
	return []byte(value), nil
}

//DetailGroupContact represents the caracteristics of a contact
type DetailGroupContact struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

//DetailResultGroup represents a contact group array
type DetailResultGroup struct {
	DetailGroups []DetailGroup `json:"result" yaml:"result"`
}

//DetailResultGroupContact represents a contact array
type DetailResultGroupContact struct {
	DetailGroupContacts []DetailGroupContact `json:"result" yaml:"result"`
}

//DetailGroupServer represents a server with informations
type DetailGroupServer struct {
	Server DetailGroupInformations `json:"server" yaml:"server"`
}

//DetailGroupInformations represents the informations of the server
type DetailGroupInformations struct {
	Name  string       `json:"name" yaml:"name"`
	Group *DetailGroup `json:"group" yaml:"group"`
}

//StringText permits to display the caracteristics of the contact group to text
func (s DetailGroupServer) StringText() string {
	var values string
	group := s.Server.Group
	if group != nil {
		elements := [][]string{{"0", "Contact group:"}, {"1", "ID: " + (*group).ID}, {"1", "Name: " + (*group).Name}, {"1", "Alias: " + (*group).Alias}}
		items := resources.GenerateListItems(elements, "")
		values = resources.BulletList(items)
	} else {
		values += "group: null\n"
	}

	return fmt.Sprintf(values)
}

//StringCSV permits to display the caracteristics of the contact group to csv
func (s DetailGroupServer) StringCSV() string {
	var p []DetailGroup
	if s.Server.Group != nil {
		p = append(p, *s.Server.Group)
	}
	b, _ := csvutil.Marshal(p)
	return string(b)
}

//StringJSON permits to display the caracteristics of the contact group to json
func (s DetailGroupServer) StringJSON() string {
	r, _ := json.MarshalIndent(s, "", " ")
	return string(r)
}

//StringYAML permits to display the caracteristics of the contact group to yaml
func (s DetailGroupServer) StringYAML() string {
	r, _ := yaml.Marshal(s)
	return string(r)
}
