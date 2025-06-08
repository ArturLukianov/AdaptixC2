package main

import (
	"encoding/json"

	adaptix "github.com/Adaptix-Framework/axc2"
)

const (
	OS_UNKNOWN = 0
	OS_WINDOWS = 1
	OS_LINUX   = 2
	OS_MAC     = 3

	TYPE_TASK       = 1
	TYPE_BROWSER    = 2
	TYPE_JOB        = 3
	TYPE_TUNNEL     = 4
	TYPE_PROXY_DATA = 5

	MESSAGE_INFO    = 5
	MESSAGE_ERROR   = 6
	MESSAGE_SUCCESS = 7
)

type Teamserver interface {
	TsListenerInteralHandler(watermark string, data []byte) (string, error)

	TsAgentProcessData(agentId string, bodyData []byte) error

	TsAgentUpdateData(newAgentData adaptix.AgentData) error
	TsAgentTerminate(agentId string, terminateTaskId string) error

	TsAgentConsoleOutput(agentId string, messageType int, message string, clearText string, store bool)
	TsAgentConsoleOutputClient(agentId string, client string, messageType int, message string, clearText string)

	TsTaskCreate(agentId string, cmdline string, client string, taskData adaptix.TaskData)
	TsTaskUpdate(agentId string, data adaptix.TaskData)
	TsTaskGetAvailableAll(agentId string, availableSize int) ([]adaptix.TaskData, error)

	TsDownloadAdd(agentId string, fileId string, fileName string, fileSize int) error
	TsDownloadUpdate(fileId string, state int, data []byte) error
	TsDownloadClose(fileId string, reason int) error

	TsClientGuiDisks(taskData adaptix.TaskData, jsonDrives string)
	TsClientGuiFiles(taskData adaptix.TaskData, path string, jsonFiles string)
	TsClientGuiFilesStatus(taskData adaptix.TaskData)
	TsClientGuiProcess(taskData adaptix.TaskData, jsonFiles string)
}

type ModuleExtender struct {
	ts Teamserver
}

var (
	ModuleObject   *ModuleExtender
	ModuleDir      string
	AgentWatermark string
)

func InitPlugin(ts any, moduleDir string, watermark string) any {
	ModuleDir = moduleDir
	AgentWatermark = watermark

	ModuleObject = &ModuleExtender{
		ts: ts.(Teamserver),
	}
	return ModuleObject
}

func (m *ModuleExtender) AgentGenerate(config string, operatingSystem string, listenerWM string, listenerProfile []byte) ([]byte, string, error) {
	var (
		listenerMap map[string]any
		// agentProfile []byte
		err error
	)

	err = json.Unmarshal(listenerProfile, &listenerMap)
	if err != nil {
		return nil, "", err
	}

	// agentProfile, err = AgentGenerateProfile(config, operatingSystem, listenerWM, listenerMap)
	// if err != nil {
	// 	return nil, "", err
	// }

	return AgentGenerateBuild(config, operatingSystem, listenerMap)
}

// func (m *ModuleExtender) AgentCreate(beat []byte) (adaptix.AgentData, error) {
// 	return CreateAgent(beat)
// }

// func (m *ModuleExtender) AgentCommand(client string, cmdline string, agentData adaptix.AgentData, args map[string]any) error {
// 	command, ok := args["command"].(string)
// 	if !ok {
// 		return errors.New("'command' must be set")
// 	}

// 	taskData, messageData, err := CreateTask(m.ts, agentData, command, args)
// 	if err != nil {
// 		return err
// 	}

// 	m.ts.TsTaskCreate(agentData.Id, cmdline, client, taskData)

// 	if len(messageData.Message) > 0 || len(messageData.Text) > 0 {
// 		m.ts.TsAgentConsoleOutput(agentData.Id, messageData.Status, messageData.Message, messageData.Text, false)
// 	}

// 	return nil
// }

// func (m *ModuleExtender) AgentPackData(agentData adaptix.AgentData, tasks []adaptix.TaskData) ([]byte, error) {
// 	packedData, err := PackTasks(agentData, tasks)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return AgentEncryptData(packedData, agentData.SessionKey)
// }

// func (m *ModuleExtender) AgentPivotPackData(pivotId string, data []byte) (adaptix.TaskData, error) {
// 	packData, err := PackPivotTasks(pivotId, data)
// 	if err != nil {
// 		return adaptix.TaskData{}, err
// 	}

// 	randomBytes := make([]byte, 16)
// 	rand.Read(randomBytes)
// 	uid := hex.EncodeToString(randomBytes)[:8]

// 	taskData := adaptix.TaskData{
// 		TaskId: uid,
// 		Type:   TYPE_PROXY_DATA,
// 		Data:   packData,
// 		Sync:   false,
// 	}

// 	return taskData, nil
// }

// func (m *ModuleExtender) AgentProcessData(agentData adaptix.AgentData, packedData []byte) ([]byte, error) {
// 	decryptData, err := AgentDecryptData(packedData, agentData.SessionKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	taskData := adaptix.TaskData{
// 		Type:        TYPE_TASK,
// 		AgentId:     agentData.Id,
// 		FinishDate:  time.Now().Unix(),
// 		MessageType: MESSAGE_SUCCESS,
// 		Completed:   true,
// 		Sync:        true,
// 	}

// 	resultTasks := ProcessTasksResult(m.ts, agentData, taskData, decryptData)

// 	for _, task := range resultTasks {
// 		m.ts.TsTaskUpdate(agentData.Id, task)
// 	}

// 	return nil, nil
// }
