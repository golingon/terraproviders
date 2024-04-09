package terraproviders_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

//go:generate go install github.com/golingon/lingon/cmd/terragen@latest
//go:generate go test -v -run ^TestParseProv$
//go:generate go test -v -run ^TestGenerate$

type Provider struct {
	Name      string `json:"name"`
	Namespace string `json:"-"`
	CleanName string `json:"cleanname"`
	Source    string `json:"source"`
	Version   string `json:"version"`
}

type ProviderAttr struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		PublishedAt time.Time `json:"published-at"`
		Tag         string    `json:"tag"`
		Version     string    `json:"version"`
	} `json:"attributes"`
}

type Response struct {
	Data     map[string]interface{} `json:"data"`
	Included []ProviderAttr         `json:"included"`
}

const (
	TFRegistryURL   = "https://registry.terraform.io/v2/providers"
	ListProviders   = "providers.txt"
	LatestProviders = "latest.json"
)

func makeURL(name, ns string) (string, error) {
	baseURL := TFRegistryURL + "/" + ns + "/" + name
	queryParams := url.Values{
		"name":      []string{name},
		"namespace": []string{ns},
		"include":   []string{"provider-versions"},
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	u.RawQuery = queryParams.Encode()
	return u.String(), nil
}

func extractProvider(r io.Reader) ([]*Provider, error) {
	pp := make([]*Provider, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		txt := scanner.Text()
		ns, name, found := strings.Cut(txt, "/")
		if !found {
			return nil, fmt.Errorf("not found: %s", txt)
		}
		p := Provider{
			Name:      name,
			CleanName: strings.ReplaceAll(name, "-", ""),
			Namespace: ns,
			Source:    txt,
		}
		pp = append(pp, &p)
	}
	return pp, scanner.Err()
}

func TestParseProv(t *testing.T) {
	f, err := os.OpenFile(ListProviders, os.O_RDONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	pp, err := extractProvider(f)
	if err != nil {
		t.Fatal(err)
	}

	// populate version with latest version
	for i := 0; i < len(pp); i++ {
		lv, err := getLatest(pp[i])
		if err != nil {
			t.Fatal(err)
		}
		pp[i].Version = lv
	}

	f, err = os.OpenFile(LatestProviders, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err = enc.Encode(pp); err != nil {
		t.Fatal(err)
	}
}

func getLatest(p *Provider) (string, error) {
	u, err := makeURL(p.Name, p.Namespace)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	if len(data.Included) == 0 {
		return "", fmt.Errorf("no version found for %s", p.Name)
	}

	vv := data.Included
	sort.SliceStable(vv, func(i, j int) bool {
		return vv[i].Attributes.PublishedAt.After(vv[j].Attributes.PublishedAt)
	})

	return vv[0].Attributes.Version, nil
}

func TestGenerate(t *testing.T) {
	f, err := os.OpenFile(LatestProviders, os.O_RDONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	var res []Provider
	if err = json.NewDecoder(f).Decode(&res); err != nil {
		t.Fatal(err)
	}
	for _, p := range res {
		sp := p.CleanName + "=" + p.Source + ":" + p.Version
		outDir := filepath.Join(p.CleanName, p.Version)
		outPkg := "github.com/golingon/terraproviders/" + p.CleanName + "/" + p.Version
		t.Log(sp, outDir, outPkg)

		cmd := exec.Command("terragen",
			"-out", outDir,
			"-pkg", outPkg,
			"-provider", sp,
			"-force")

		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		if err = cmd.Run(); err != nil {
			t.Fatal(err)
		}

		// create go.mod
		gocmd := exec.Command("go", "mod", "init", outPkg)
		gocmd.Dir = outDir
		gocmd.Stderr = os.Stderr
		gocmd.Stdout = os.Stdout

		if err = gocmd.Run(); err != nil {
			t.Fatal(err)
		}

		tidycmd := exec.Command("go", "mod", "tidy")
		tidycmd.Dir = outDir
		tidycmd.Stderr = os.Stderr
		tidycmd.Stdout = os.Stdout

		if err = tidycmd.Run(); err != nil {
			t.Fatal(err)
		}

	}
}
