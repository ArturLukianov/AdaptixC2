package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type GenerateConfig struct {
	Arch             string `json:"arch"`
	Format           string `json:"format"`
	ReconnectTimeout string `json:"reconn_timeout"`
	ReconnectCount   int    `json:"reconn_count"`
}

var (
	SrcPath = "src_cog"
)

func AgentGenerateBuild(agentConfig string, operatingSystem string, listenerMap map[string]any) ([]byte, string, error) {
	var (
		generateConfig GenerateConfig
		// GoArch         string
		Filename  string
		buildPath string
		stdout    bytes.Buffer
		stderr    bytes.Buffer
	)

	err := json.Unmarshal([]byte(agentConfig), &generateConfig)
	if err != nil {
		return nil, "", err
	}

	currentDir := ModuleDir
	tempDir, err := os.MkdirTemp("", "ax-*")
	if err != nil {
		return nil, "", err
	}

	if generateConfig.Arch == "amd64" {
		// GoArch = "amd64"
	} else {
		_ = os.RemoveAll(tempDir)
		return nil, "", errors.New("unknown architecture")
	}

	if operatingSystem == "windows" {
		Filename = "agent.exe"
	} else {
		_ = os.RemoveAll(tempDir)
		return nil, "", errors.New("operating system not supported")
	}
	buildPath = tempDir + "/" + Filename

	// config := fmt.Sprintf("package main\n\nvar encProfile = []byte(\"%s\")\n", string(agentProfile))
	// configPath := currentDir + "/" + SrcPath + "/config.go"
	// err = os.WriteFile(configPath, []byte(config), 0644)
	// if err != nil {
	// 	_ = os.RemoveAll(tempDir)
	// 	return nil, "", err
	// }

	cmdBuild := fmt.Sprintf("cargo build --release --target x86_64-pc-windows-gnu && cp ./target/release/release/src_cog.exe %s", buildPath)
	runnerCmdBuild := exec.Command("sh", "-c", cmdBuild)
	runnerCmdBuild.Dir = currentDir + "/" + SrcPath
	runnerCmdBuild.Stdout = &stdout
	runnerCmdBuild.Stderr = &stderr
	err = runnerCmdBuild.Run()
	if err != nil {
		_ = os.RemoveAll(tempDir)
		return nil, "", err
	}

	buildContent, err := os.ReadFile(buildPath)
	if err != nil {
		return nil, "", err
	}
	_ = os.RemoveAll(tempDir)

	return buildContent, Filename, nil
}
