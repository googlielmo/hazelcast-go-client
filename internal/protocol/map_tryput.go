// Copyright (c) 2008-2017, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	. "github.com/hazelcast/hazelcast-go-client/internal/common"
	. "github.com/hazelcast/hazelcast-go-client/internal/serialization"
)

type MapTryPutResponseParameters struct {
	Response bool
}

func MapTryPutCalculateSize(name *string, key *Data, value *Data, threadId int64, timeout int64) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += StringCalculateSize(name)
	dataSize += DataCalculateSize(key)
	dataSize += DataCalculateSize(value)
	dataSize += INT64_SIZE_IN_BYTES
	dataSize += INT64_SIZE_IN_BYTES
	return dataSize
}

func MapTryPutEncodeRequest(name *string, key *Data, value *Data, threadId int64, timeout int64) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, MapTryPutCalculateSize(name, key, value, threadId, timeout))
	clientMessage.SetMessageType(MAP_TRYPUT)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendData(key)
	clientMessage.AppendData(value)
	clientMessage.AppendInt64(threadId)
	clientMessage.AppendInt64(timeout)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

func MapTryPutDecodeResponse(clientMessage *ClientMessage) *MapTryPutResponseParameters {
	// Decode response from client message
	parameters := new(MapTryPutResponseParameters)
	parameters.Response = clientMessage.ReadBool()
	return parameters
}
