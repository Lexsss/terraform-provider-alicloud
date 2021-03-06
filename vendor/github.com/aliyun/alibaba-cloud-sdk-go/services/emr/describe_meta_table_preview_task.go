package emr

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

// DescribeMetaTablePreviewTask invokes the emr.DescribeMetaTablePreviewTask API synchronously
func (client *Client) DescribeMetaTablePreviewTask(request *DescribeMetaTablePreviewTaskRequest) (response *DescribeMetaTablePreviewTaskResponse, err error) {
	response = CreateDescribeMetaTablePreviewTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMetaTablePreviewTaskWithChan invokes the emr.DescribeMetaTablePreviewTask API asynchronously
func (client *Client) DescribeMetaTablePreviewTaskWithChan(request *DescribeMetaTablePreviewTaskRequest) (<-chan *DescribeMetaTablePreviewTaskResponse, <-chan error) {
	responseChan := make(chan *DescribeMetaTablePreviewTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMetaTablePreviewTask(request)
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

// DescribeMetaTablePreviewTaskWithCallback invokes the emr.DescribeMetaTablePreviewTask API asynchronously
func (client *Client) DescribeMetaTablePreviewTaskWithCallback(request *DescribeMetaTablePreviewTaskRequest, callback func(response *DescribeMetaTablePreviewTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMetaTablePreviewTaskResponse
		var err error
		defer close(result)
		response, err = client.DescribeMetaTablePreviewTask(request)
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

// DescribeMetaTablePreviewTaskRequest is the request struct for api DescribeMetaTablePreviewTask
type DescribeMetaTablePreviewTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	TaskId          string           `position:"Query" name:"TaskId"`
}

// DescribeMetaTablePreviewTaskResponse is the response struct for api DescribeMetaTablePreviewTask
type DescribeMetaTablePreviewTaskResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TaskId      string `json:"TaskId" xml:"TaskId"`
	TaskStatus  string `json:"TaskStatus" xml:"TaskStatus"`
	StartTime   int64  `json:"StartTime" xml:"StartTime"`
	EndTime     int64  `json:"EndTime" xml:"EndTime"`
	ExecuteTime int64  `json:"ExecuteTime" xml:"ExecuteTime"`
	TaskProcess int    `json:"TaskProcess" xml:"TaskProcess"`
	Data        Data   `json:"Data" xml:"Data"`
}

// CreateDescribeMetaTablePreviewTaskRequest creates a request to invoke DescribeMetaTablePreviewTask API
func CreateDescribeMetaTablePreviewTaskRequest() (request *DescribeMetaTablePreviewTaskRequest) {
	request = &DescribeMetaTablePreviewTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeMetaTablePreviewTask", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMetaTablePreviewTaskResponse creates a response to parse from DescribeMetaTablePreviewTask response
func CreateDescribeMetaTablePreviewTaskResponse() (response *DescribeMetaTablePreviewTaskResponse) {
	response = &DescribeMetaTablePreviewTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
