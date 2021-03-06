package cbcluster

import (
	"fmt"
	"strconv"
	"strings"
)

func IsCommandEnabled(arguments map[string]interface{}, commandKey string) bool {
	val, ok := arguments[commandKey]
	if !ok {
		return false
	}
	boolVal, ok := val.(bool)
	if !ok {
		return false
	}
	return boolVal
}

// convert from comma separated list to a string slice
func ExtractEtcdServerList(docOptParsed map[string]interface{}) []string {

	rawServerList, found := docOptParsed["--etcd-servers"]
	if !found {
		return nil
	}

	rawServerListStr, ok := rawServerList.(string)
	if !ok {
		return nil
	}

	return strings.Split(rawServerListStr, ",")

}

func ExtractUserPass(docOptParsed map[string]interface{}) (string, error) {
	return ExtractStringArg(docOptParsed, "--userpass")
}

func ExtractDockerTagOrLatest(docOptParsed map[string]interface{}) string {
	dockerTag, err := ExtractStringArg(docOptParsed, "--docker-tag")
	if err != nil || dockerTag == "" {
		return "latest"
	}
	return dockerTag

}

func ExtractCbVersion(docOptParsed map[string]interface{}) (string, error) {
	return ExtractStringArg(docOptParsed, "--version")
}

func ExtractIntArg(docOptParsed map[string]interface{}, argToExtract string) (int, error) {

	rawVal, found := docOptParsed[argToExtract]
	if !found {
		return -1, fmt.Errorf("Did not find arg: %v", argToExtract)
	}

	stringVal, ok := rawVal.(string)
	if !ok {
		return -1, fmt.Errorf("Invalid type for %v", argToExtract)
	}

	intVal, err := strconv.ParseInt(stringVal, 10, 64)
	if err != nil {
		return -1, err
	}

	return int(intVal), nil

}

func ExtractStringArg(docOptParsed map[string]interface{}, argToExtract string) (string, error) {

	rawVal, found := docOptParsed[argToExtract]
	if !found {
		return "", fmt.Errorf("Did not find arg: %v", argToExtract)
	}

	stringVal, ok := rawVal.(string)
	if !ok {
		return "", fmt.Errorf("Invalid type for %v", argToExtract)
	}

	return stringVal, nil

}

func ExtractBoolArg(docOptParsed map[string]interface{}, argToExtract string) bool {

	rawVal, found := docOptParsed[argToExtract]
	if !found {
		return false
	}

	boolVal, ok := rawVal.(bool)
	if !ok {
		return false
	}

	return boolVal

}

func ExtractSkipCheckCleanState(docOptParsed map[string]interface{}) bool {

	return ExtractBoolArg(docOptParsed, "--skip-clean-slate-check")

}

func ExtractNumNodes(docOptParsed map[string]interface{}) (int, error) {

	return ExtractIntArg(docOptParsed, "--num-nodes")

}
