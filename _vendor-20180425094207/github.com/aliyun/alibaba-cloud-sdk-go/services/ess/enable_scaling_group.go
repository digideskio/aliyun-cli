package ess

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

// EnableScalingGroup invokes the ess.EnableScalingGroup API synchronously
// api document: https://help.aliyun.com/api/ess/enablescalinggroup.html
func (client *Client) EnableScalingGroup(request *EnableScalingGroupRequest) (response *EnableScalingGroupResponse, err error) {
	response = CreateEnableScalingGroupResponse()
	err = client.DoAction(request, response)
	return
}

// EnableScalingGroupWithChan invokes the ess.EnableScalingGroup API asynchronously
// api document: https://help.aliyun.com/api/ess/enablescalinggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableScalingGroupWithChan(request *EnableScalingGroupRequest) (<-chan *EnableScalingGroupResponse, <-chan error) {
	responseChan := make(chan *EnableScalingGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableScalingGroup(request)
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

// EnableScalingGroupWithCallback invokes the ess.EnableScalingGroup API asynchronously
// api document: https://help.aliyun.com/api/ess/enablescalinggroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableScalingGroupWithCallback(request *EnableScalingGroupRequest, callback func(response *EnableScalingGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableScalingGroupResponse
		var err error
		defer close(result)
		response, err = client.EnableScalingGroup(request)
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

// EnableScalingGroupRequest is the request struct for api EnableScalingGroup
type EnableScalingGroupRequest struct {
	*requests.RpcRequest
	OwnerId                      requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount         string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId              requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ScalingGroupId               string           `position:"Query" name:"ScalingGroupId"`
	ActiveScalingConfigurationId string           `position:"Query" name:"ActiveScalingConfigurationId"`
	InstanceId1                  string           `position:"Query" name:"InstanceId.1"`
	InstanceId2                  string           `position:"Query" name:"InstanceId.2"`
	InstanceId3                  string           `position:"Query" name:"InstanceId.3"`
	InstanceId4                  string           `position:"Query" name:"InstanceId.4"`
	InstanceId5                  string           `position:"Query" name:"InstanceId.5"`
	InstanceId6                  string           `position:"Query" name:"InstanceId.6"`
	InstanceId7                  string           `position:"Query" name:"InstanceId.7"`
	InstanceId8                  string           `position:"Query" name:"InstanceId.8"`
	InstanceId9                  string           `position:"Query" name:"InstanceId.9"`
	InstanceId10                 string           `position:"Query" name:"InstanceId.10"`
	InstanceId11                 string           `position:"Query" name:"InstanceId.11"`
	InstanceId12                 string           `position:"Query" name:"InstanceId.12"`
	InstanceId13                 string           `position:"Query" name:"InstanceId.13"`
	InstanceId14                 string           `position:"Query" name:"InstanceId.14"`
	InstanceId15                 string           `position:"Query" name:"InstanceId.15"`
	InstanceId16                 string           `position:"Query" name:"InstanceId.16"`
	InstanceId17                 string           `position:"Query" name:"InstanceId.17"`
	InstanceId18                 string           `position:"Query" name:"InstanceId.18"`
	InstanceId19                 string           `position:"Query" name:"InstanceId.19"`
	InstanceId20                 string           `position:"Query" name:"InstanceId.20"`
	LoadBalancerWeight1          requests.Integer `position:"Query" name:"LoadBalancerWeight.1"`
	LoadBalancerWeight2          requests.Integer `position:"Query" name:"LoadBalancerWeight.2"`
	LoadBalancerWeight3          requests.Integer `position:"Query" name:"LoadBalancerWeight.3"`
	LoadBalancerWeight4          requests.Integer `position:"Query" name:"LoadBalancerWeight.4"`
	LoadBalancerWeight5          requests.Integer `position:"Query" name:"LoadBalancerWeight.5"`
	LoadBalancerWeight6          requests.Integer `position:"Query" name:"LoadBalancerWeight.6"`
	LoadBalancerWeight7          requests.Integer `position:"Query" name:"LoadBalancerWeight.7"`
	LoadBalancerWeight8          requests.Integer `position:"Query" name:"LoadBalancerWeight.8"`
	LoadBalancerWeight9          requests.Integer `position:"Query" name:"LoadBalancerWeight.9"`
	LoadBalancerWeight10         requests.Integer `position:"Query" name:"LoadBalancerWeight.10"`
	LoadBalancerWeight11         requests.Integer `position:"Query" name:"LoadBalancerWeight.11"`
	LoadBalancerWeight12         requests.Integer `position:"Query" name:"LoadBalancerWeight.12"`
	LoadBalancerWeight13         requests.Integer `position:"Query" name:"LoadBalancerWeight.13"`
	LoadBalancerWeight14         requests.Integer `position:"Query" name:"LoadBalancerWeight.14"`
	LoadBalancerWeight15         requests.Integer `position:"Query" name:"LoadBalancerWeight.15"`
	LoadBalancerWeight16         requests.Integer `position:"Query" name:"LoadBalancerWeight.16"`
	LoadBalancerWeight17         requests.Integer `position:"Query" name:"LoadBalancerWeight.17"`
	LoadBalancerWeight18         requests.Integer `position:"Query" name:"LoadBalancerWeight.18"`
	LoadBalancerWeight19         requests.Integer `position:"Query" name:"LoadBalancerWeight.19"`
	LoadBalancerWeight20         requests.Integer `position:"Query" name:"LoadBalancerWeight.20"`
	OwnerAccount                 string           `position:"Query" name:"OwnerAccount"`
}

// EnableScalingGroupResponse is the response struct for api EnableScalingGroup
type EnableScalingGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableScalingGroupRequest creates a request to invoke EnableScalingGroup API
func CreateEnableScalingGroupRequest() (request *EnableScalingGroupRequest) {
	request = &EnableScalingGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "EnableScalingGroup", "ess", "openAPI")
	return
}

// CreateEnableScalingGroupResponse creates a response to parse from EnableScalingGroup response
func CreateEnableScalingGroupResponse() (response *EnableScalingGroupResponse) {
	response = &EnableScalingGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}