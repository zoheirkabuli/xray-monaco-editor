package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func removeExtDatRules(obj interface{}) {
	switch v := obj.(type) {
	case map[string]interface{}:
		fmt.Printf("Processing map with keys: %v\n", getKeys(v))
		for key, value := range v {
			switch val := value.(type) {
			case []interface{}:
				fmt.Printf("Processing array in key '%s' with %d elements\n", key, len(val))
				filteredRules := make([]interface{}, 0)
				for _, rule := range val {
					switch r := rule.(type) {
					case string:
						if strings.Contains(r, "ext:") && strings.Contains(r, ".dat") {
							fmt.Printf("Removing string rule: %s\n", r)
							continue
						}
						filteredRules = append(filteredRules, r)
					case map[string]interface{}:
						if domains, ok := r["domain"].([]interface{}); ok {
							fmt.Printf("Found nested domain rules with %d elements\n", len(domains))
							filteredDomains := make([]interface{}, 0)
							for _, domain := range domains {
								if domainStr, ok := domain.(string); ok {
									if strings.Contains(domainStr, "ext:") && strings.Contains(domainStr, ".dat") {
										fmt.Printf("Removing nested rule: %s\n", domainStr)
										continue
									}
									filteredDomains = append(filteredDomains, domain)
								}
							}
							r["domain"] = filteredDomains
						}
						filteredRules = append(filteredRules, r)
					default:
						filteredRules = append(filteredRules, rule)
					}
				}
				v[key] = filteredRules
			case map[string]interface{}:
				removeExtDatRules(val)
			}
		}
	case []interface{}:
		for _, item := range v {
			removeExtDatRules(item)
		}
	}
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	testConfig := `{
		"routing": {
			"rules": [{
				"domain": [
					"ext:zapret.dat:zapret",
					"ext:zapret.dat:zapret-zapad",
					"geosite:github",
					"domain:example.com"
				]
			}]
		}
	}`
	
	var configObj interface{}
	json.Unmarshal([]byte(testConfig), &configObj)
	
	fmt.Println("Before removal:")
	beforeJson, _ := json.MarshalIndent(configObj, "", "  ")
	fmt.Println(string(beforeJson))
	
	removeExtDatRules(configObj)
	
	fmt.Println("\nAfter removal:")
	afterJson, _ := json.MarshalIndent(configObj, "", "  ")
	fmt.Println(string(afterJson))
} 