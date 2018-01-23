package consul

import (
	"crypto/tls"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"net/http"
	"sort"
	"strings"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func getConsul() *consulapi.Client {
	server := viper.GetString("server")
	scheme := viper.GetString("scheme")
	user := viper.GetString("user")
	pass := viper.GetString("pass")
	token := viper.GetString("token")

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	config := consulapi.DefaultConfig()
	config.Address = server
	config.Scheme = scheme
	config.HttpClient = &http.Client{Transport: transport}
	config.HttpAuth = &consulapi.HttpBasicAuth{Username: user, Password: pass}
	config.Token = token

	consul, _ := consulapi.NewClient(config)
	return consul
}

func contains(s []string, v string) bool {
	for _, value := range s {
		if value == v {
			return true
		}
	}
	return false
}

func prefixIsUnique(s []string, prefix string) bool {
	for _, p := range s {
		if p != prefix && strings.HasPrefix(prefix, p) {
			return false
		}
	}
	return true
}

func prefixesToQuery(prefixes []string) []string {
	sort.Sort(ByLength(prefixes))

	var uniquePrefixes []string

	for _, prefix := range prefixes {
		prefix = strings.Trim(prefix, "/")
		if prefixIsUnique(prefixes, prefix) {
			uniquePrefixes = append(uniquePrefixes, prefix)
		}
	}

	return uniquePrefixes
}

func processEnv(envMap map[string]map[string]string, prefixes []string) map[string]string {
	sort.Sort(sort.Reverse(ByLength(prefixes)))

	env := make(map[string]string)

	for _, prefix := range prefixes {
		prefix = strings.Trim(prefix, "/")
		if _, ok := envMap[prefix]; ok {
			for k, v := range envMap[prefix] {
				if _, ok := env[k]; !ok {
					env[k] = v
				}
			}
		}
	}
	for k, v := range env {
		fmt.Printf("%s=%s\n", k, v)
	}
	return env
}

func Get() {
	consul := getConsul()

	prefixes := viper.GetStringSlice("prefix")

	uniquePrefixes := prefixesToQuery(prefixes)

	kv := consul.KV()

	envMap := map[string]map[string]string{}

	for _, p := range uniquePrefixes {
		kvPairs, qm, err := kv.List(p, nil)
		if err != nil {
			fmt.Println(err, qm)
		} else {
			for _, kvPair := range kvPairs {
				val := string(kvPair.Value)

				parts := strings.Split(kvPair.Key, "/")
				folder := strings.Join(parts[:len(parts)-1], "/")
				folder = strings.Trim(folder, "/")
				varName := parts[len(parts)-1]

				if val != "" {
					if _, ok := envMap[folder]; !ok {
						envMap[folder] = make(map[string]string)
					}
					envMap[folder][varName] = val
				}
			}
		}
	}

	processEnv(envMap, prefixes)
}
