package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/vinybergamo/cloud-manager/vars"
)

func ExecCommand(a ...string) (string, error) {
	if vars.IsServerMode {
		command := append([]string{"cloud"}, a...)
		cmd := exec.Command(command[0], command[1:]...)

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			errMsg := fmt.Sprintf("Command failed: %s \n%s", command, stderr.String())
			return "", fmt.Errorf(errMsg)

		}
		return stdout.String(), nil
	}

	command := append([]string{"ssh", vars.Server}, a...)
	cmd := exec.Command(command[0], command[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errMsg := fmt.Sprintf("Command failed: %s \n%s", command, stderr.String())
		return "", fmt.Errorf(errMsg)

	}

	return stdout.String(), nil
}

func Logger(c string, a ...string) {
	colorsCodes := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"reset":   "\033[0m",
	}

	code, ok := colorsCodes[c]
	if !ok {
		code = colorsCodes["reset"]
	}

	message := fmt.Sprintf("%s%s%s", code, strings.Join(a, " "), colorsCodes["reset"])

	fmt.Println(message)
}

func LoggerError(m ...string) {
	fmt.Println("\033[31m", "[ ERROR ]", "!!", strings.Join(m, " "), "\033[0m")
	if vars.IsShellMode {
		os.Exit(0)
	}
}

func CheckHTTPDomain(d string) (string, error) {
	Logger("magenta", "Checking domain", d)
	for i := 0; i < 3; i++ {
		res, err := http.Head("http://" + d)
		if err == nil && res.StatusCode == http.StatusOK {
			ips, err := net.LookupIP(d)
			if err == nil && len(ips) > 0 {
				ip := ips[0].String()
				return ip, nil
			}
		}
		Logger("cyan", "Domain is not available", "Retrying in 10 seconds")
		time.Sleep(10 * time.Second)
	}

	return "", fmt.Errorf("domain is not available")
}

func CheckHTTPSDomain(d string) (string, error) {
	Logger("magenta", "Checking domain", d)
	for i := 0; i < 3; i++ {
		res, err := http.Head("https://" + d)
		if err == nil && res.StatusCode == http.StatusOK {
			ips, err := net.LookupIP(d)
			if err == nil && len(ips) > 0 {
				ip := ips[0].String()
				return ip, nil
			}
		}
		Logger("cyan", "Domain is not available", "Retrying in 10 seconds")
		time.Sleep(10 * time.Second)
	}

	return "", fmt.Errorf("domain is not available")
}

func IsLink(t string) bool {
	pattern := regexp.MustCompile(`(?i)^(https|http|ftp)://[^\s/$.?#].[^\s]*$`)

	return pattern.MatchString(t)
}

func IsNumber(t string) bool {
	pattern := regexp.MustCompile(`^[0-9]+$`)

	return pattern.MatchString(t)
}

func ReadInput(t string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s:\n> ", t)
	scanner.Scan()
	input := scanner.Text()

	if IsLink(t) {
		input = strings.TrimRight(input, " ")
		return input
	}

	if IsNumber(t) {
		input = strings.TrimSpace(input)
		return input
	}

	input = strings.TrimLeft(input, "-")
	input = strings.Replace(input, " ", "-", -1)
	input = strings.ToLower(input)
	input = strings.TrimRight(input, "-")

	return input
}

func CreateJSONResponse(r interface{}) string {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return `{"error": "Failed to create JSON response"}`
	}

	return string(jsonData)
}
