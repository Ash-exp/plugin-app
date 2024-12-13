package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for key, value := range GetSystemEnvVariables() {
		fmt.Println(fmt.Sprintf("%s=%s", key, value))
	}
}

func GetSystemEnvVariables() map[string]string {
	envs := make(map[string]string)
	//get all environment variables
	envVars := os.Environ()
	for _, envVar := range envVars {
		subs := strings.SplitN(envVar, "=", 2)
		envs[subs[0]] = subs[1]
	}
	return envs
}
