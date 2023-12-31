//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

func unmarshalRunRequestClassification(rawMsg json.RawMessage) (armcontainerregistry.RunRequestClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armcontainerregistry.RunRequestClassification
	switch m["type"] {
	case "DockerBuildRequest":
		b = &armcontainerregistry.DockerBuildRequest{}
	case "EncodedTaskRunRequest":
		b = &armcontainerregistry.EncodedTaskRunRequest{}
	case "FileTaskRunRequest":
		b = &armcontainerregistry.FileTaskRunRequest{}
	case "TaskRunRequest":
		b = &armcontainerregistry.TaskRunRequest{}
	default:
		b = &armcontainerregistry.RunRequest{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTaskStepPropertiesClassification(rawMsg json.RawMessage) (armcontainerregistry.TaskStepPropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armcontainerregistry.TaskStepPropertiesClassification
	switch m["type"] {
	case string(armcontainerregistry.StepTypeDocker):
		b = &armcontainerregistry.DockerBuildStep{}
	case string(armcontainerregistry.StepTypeEncodedTask):
		b = &armcontainerregistry.EncodedTaskStep{}
	case string(armcontainerregistry.StepTypeFileTask):
		b = &armcontainerregistry.FileTaskStep{}
	default:
		b = &armcontainerregistry.TaskStepProperties{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalTaskStepUpdateParametersClassification(rawMsg json.RawMessage) (armcontainerregistry.TaskStepUpdateParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b armcontainerregistry.TaskStepUpdateParametersClassification
	switch m["type"] {
	case string(armcontainerregistry.StepTypeDocker):
		b = &armcontainerregistry.DockerBuildStepUpdateParameters{}
	case string(armcontainerregistry.StepTypeEncodedTask):
		b = &armcontainerregistry.EncodedTaskStepUpdateParameters{}
	case string(armcontainerregistry.StepTypeFileTask):
		b = &armcontainerregistry.FileTaskStepUpdateParameters{}
	default:
		b = &armcontainerregistry.TaskStepUpdateParameters{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
