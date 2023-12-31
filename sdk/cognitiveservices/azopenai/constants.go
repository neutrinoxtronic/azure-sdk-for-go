//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// azureOpenAIOperationState - The state of a job or item.
type azureOpenAIOperationState string

const (
	azureOpenAIOperationStateCanceled   azureOpenAIOperationState = "canceled"
	azureOpenAIOperationStateFailed     azureOpenAIOperationState = "failed"
	azureOpenAIOperationStateNotRunning azureOpenAIOperationState = "notRunning"
	azureOpenAIOperationStateRunning    azureOpenAIOperationState = "running"
	azureOpenAIOperationStateSucceeded  azureOpenAIOperationState = "succeeded"
)

// ChatRole - A description of the intended purpose of a message within a chat completions interaction.
type ChatRole string

const (
	ChatRoleAssistant ChatRole = "assistant"
	ChatRoleSystem    ChatRole = "system"
	ChatRoleUser      ChatRole = "user"
)

// PossibleChatRoleValues returns the possible values for the ChatRole const type.
func PossibleChatRoleValues() []ChatRole {
	return []ChatRole{
		ChatRoleAssistant,
		ChatRoleSystem,
		ChatRoleUser,
	}
}

// CompletionsFinishReason - Representation of the manner in which a completions response concluded.
type CompletionsFinishReason string

const (
	CompletionsFinishReasonContentFilter CompletionsFinishReason = "content_filter"
	CompletionsFinishReasonLength        CompletionsFinishReason = "length"
	CompletionsFinishReasonStop          CompletionsFinishReason = "stop"
)

// PossibleCompletionsFinishReasonValues returns the possible values for the CompletionsFinishReason const type.
func PossibleCompletionsFinishReasonValues() []CompletionsFinishReason {
	return []CompletionsFinishReason{
		CompletionsFinishReasonContentFilter,
		CompletionsFinishReasonLength,
		CompletionsFinishReasonStop,
	}
}

// ImageGenerationResponseFormat - The format in which the generated images are returned.
type ImageGenerationResponseFormat string

const (
	ImageGenerationResponseFormatB64JSON ImageGenerationResponseFormat = "b64_json"
	ImageGenerationResponseFormatURL     ImageGenerationResponseFormat = "url"
)

// PossibleImageGenerationResponseFormatValues returns the possible values for the ImageGenerationResponseFormat const type.
func PossibleImageGenerationResponseFormatValues() []ImageGenerationResponseFormat {
	return []ImageGenerationResponseFormat{
		ImageGenerationResponseFormatB64JSON,
		ImageGenerationResponseFormatURL,
	}
}

// ImageSize - The desired size of the generated images. Must be one of 256x256, 512x512, or 1024x1024.
type ImageSize string

const (
	ImageSize512x512   ImageSize = "512x512"
	ImageSize1024x1024 ImageSize = "1024x1024"
	ImageSize256x256   ImageSize = "256x256"
)

// PossibleImageSizeValues returns the possible values for the ImageSize const type.
func PossibleImageSizeValues() []ImageSize {
	return []ImageSize{
		ImageSize512x512,
		ImageSize1024x1024,
		ImageSize256x256,
	}
}
