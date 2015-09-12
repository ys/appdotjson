package appjson

import (
	"encoding/json"
	"io"
	"os"
)

type AppJson struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Website     string   `json:"website"`
	Repository  string   `json:"repository"`
	Logo        string   `json:"logo"`
	SuccessUrl  string   `json:"success_url"`
	Scripts     Scripts  `json:"scripts"`
	RawEnv      RawEnv   `json:"env"`
	Env         Env
	Image       string      `json:"image"`
	Addons      []Addon     `json:"addons"`
	Buildpacks  []Buildpack `json:"buildpacks"`
}

type Scripts struct {
	Postdeploy string `json:"postdeploy"`
}

type EnvKey string

type EnvVar struct {
	Description string `json:"description"`
	Value       string `json:"value"`
	Required    bool   `json:"required"`
	Generator   string `json:"generator"`
}

type RawEnv map[EnvKey]json.RawMessage
type Env map[EnvKey]EnvVar

type Addon string

type Buildpack struct {
	Url string `json:"url"`
}

func Decode(r io.Reader) (*AppJson, error) {
	x := new(AppJson)
	err := json.NewDecoder(r).Decode(x)
	if err != nil {
		return nil, err
	}
	x.Env = make(Env)
	for key, value := range x.RawEnv {
		curvar := new(EnvVar)
		err := json.Unmarshal(value, &curvar)
		if err != nil {
			curvar.Value = string(value)
		}
		x.Env[key] = *curvar
	}
	return x, nil
}

func FromFile(path string) (*AppJson, error) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return nil, err
	}
	return Decode(file)
}
