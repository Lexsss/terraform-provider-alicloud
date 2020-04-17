package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// StartInstances invokes the ecs.StartInstances API synchronously
// api document: https://help.aliyun.com/api/ecs/startinstances.html
func (client *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
	response = CreateStartInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// StartInstancesWithChan invokes the ecs.StartInstances API asynchronously
// api document: https://help.aliyun.com/api/ecs/startinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartInstancesWithChan(request *StartInstancesRequest) (<-chan *StartInstancesResponse, <-chan error) {
	responseChan := make(chan *StartInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartInstances(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// StartInstancesWithCallback invokes the ecs.StartInstances API asynchronously
// api document: https://help.aliyun.com/api/ecs/startinstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartInstancesWithCallback(request *StartInstancesRequest, callback func(response *StartInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartInstancesResponse
		var err error
		defer close(result)
		response, err = client.StartInstances(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// StartInstancesRequest is the request struct for api StartInstances
type StartInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BatchOptimization    string           `position:"Query" name:"BatchOptimization"`
	DryRun               requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
}

// StartInstancesResponse is the response struct for api StartInstances
type StartInstancesResponse struct {
	*responses.BaseResponse
	RequestId         string                            `json:"RequestId" xml:"RequestId"`
	InstanceResponses InstanceResponsesInStartInstances `json:"InstanceResponses" xml:"InstanceResponses"`
}

// CreateStartInstancesRequest creates a request to invoke StartInstances API
func CreateStartInstancesRequest() (request *StartInstancesRequest) {
	request = &StartInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "StartInstances", "ecs", "openAPI")
	return
}

// CreateStartInstancesResponse creates a response to parse from StartInstances response
func CreateStartInstancesResponse() (response *StartInstancesResponse) {
	response = &StartInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
