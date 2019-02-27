package aliyunimpl

import (
	"errors"
	"log"
	"reflect"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	cons "github.com/autorestart/constant"
)

func RebootInstance(instanceid, accesskey, accessecret, region string) (bool, error) {
	ProcessError := errors.New("ProcessRequestError")
	NewClientError := errors.New("NewClientError")
	var returnCode int
	var isOk bool
	var area string
	if region == "" {
		log.Printf("[INFO] Using Default Region cn-hangzhou")
		area = cons.DefaultRegion
	}
	area = region
	client, err := sdk.NewClientWithAccessKey(area, accesskey, accessecret)
	if err != nil {
		isOk = false
		log.Printf("[ERROR] error creating NewClientWithAccessKey: %s", err)
		return isOk, NewClientError

	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.ApiName = "RebootInstance"
	request.Domain = "ecs.aliyuncs.com"
	request.Scheme = "https"
	request.Version = "2014-05-26"
	request.QueryParams["InstanceId"] = instanceid

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		isOk = false
		log.Printf("[ERROR] error processing common request: %s", err)
		return isOk, ProcessError
	}
	returnCode = response.GetHttpStatus()
	log.Printf("[INFO] Request successful: %s Response body: %s", returnCode, response.GetHttpContentString())
	isOk = true
	return isOk, nil
}

type InstanceAttr struct {
	InnerIPaddr struct {
		Ipaddress []string `json:"IpAddress"`
	} `json:"InnerIpAddress"`
	ImageId    string `json:"ImageId"`
	VlanId     string `json:"VlanId"`
	InstanceId string `json:"InstanceId"`
	EipAddress struct {
		Ipaddress          string `json:"IpAddress"`
		AllocationId       string `json:"AllocationId"`
		InternetChargeType string `json:"InternetChargeType"`
	} `json:"EipAddress"`
	InternetMaxBandwidthIn  int    `json:"InternetMaxBandwidthIn"`
	ZoneId                  string `json:"ZoneId"`
	CreditSpecification     string `json:"CreditSpecification"`
	InternetChargeType      string `json:"InternetChargeType"`
	StoppedMode             string `json:"StoppedMode"`
	SerialNumber            string `json:"SerialNumber"`
	IoOptimized             string `json:"IoOptimized"`
	Cpu                     int    `json:"Cpu"`
	Memory                  int16  `json:"Memory"`
	InternetMaxBandwidthOut int    `json:"InternetMaxBandwidthOut"`
	VpcAttributes           struct {
		NatIpAddress     string `json:"NatIpAddress"`
		PrivateIpAddress struct {
			IpAddress []string `json:"IpAddress"`
		} `json:"PrivateIpAddress"`
		VSwitchId string `json:"VSwitchId"`
		VpcId     string `json:"VpcId"`
	} `json:"VpcAttributes"`
	SecurityGroupIds struct {
		SecurityGroupId []string `json:"SecurityGroupId"`
	} `json:"SecurityGroupIds"`
	InstanceName        string `json:"InstanceName"`
	Description         string `json:"Description"`
	InstanceNetworkType string `json:"InstanceNetworkType"`
	PublicIpAddress     struct {
		IpAddress []string `json:"IpAddress"`
	} `json:"PublicIpAddress"`
	HostName               string `json:"HostName"`
	InstanceType           string `json:"InstanceType"`
	CreationTime           string `json:"CreationTime"`
	Status                 string `json:"Status"`
	ClusterId              string `json:"ClusterId"`
	RegionId               string `json:"RegionId"`
	RequestId              string `json:"RequestId"`
	DedicatedHostAttribute struct {
		DedicatedHostId   string `json:"DedicatedHostId"`
		DedicatedHostName string `json:"DedicatedHostName"`
	} `json:"DedicatedHostAttribute"`
	OperationLocks struct {
		LockReason []string `json:"LockReason"`
	} `json:"OperationLocks"`
	InstanceChargeType string `json:"InstanceChargeType"`
	ExpiredTime        string `json:"ExpiredTime"`
}

func (i *InstanceAttr) GetEipAddress() (string, error) {
	getError := errors.New("EipAddress NotFound")
	eipaddress := i.EipAddress.Ipaddress
	if eipaddress == "" {
		log.Printf("[ERROR] error geting eip address")
		return eipaddress, getError
	}
	return eipaddress, nil
}

func (i *InstanceAttr) GetPubIpAddress() ([]string, error) {
	getError := errors.New("PubAddress NotFound")
	pubipaddress := i.PublicIpAddress.IpAddress
	if len(pubipaddress) == 0 {
		log.Printf("[ERROR] error geting public ipadress")
		return nil, getError
	}
	return pubipaddress, nil
}

func DescribeInstanceAttr(instanceid, accesskey, accessecret, region string) ([]byte, error) {
	ProcessError := errors.New("ProcessRequestError")
	NewClientError := errors.New("NewClientError")
	var area string
	var content []byte
	if region == "" {
		log.Printf("[INFO] Using Default Region cn-hangzhou")
		area = cons.DefaultRegion
	}
	area = region
	client, err := sdk.NewClientWithAccessKey(area, accesskey, accessecret)
	if err != nil {
		log.Printf("[ERROR] error creating NewClientWithAccessKey :%s", err)
		return content, NewClientError
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "ecs.aliyuncs.com"
	request.Version = "2014-05-26"
	request.ApiName = "DescribeInstanceAttribute"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["InstanceId"] = instanceid

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		log.Printf("[ERROR] error processing Common request :%s", err)
		return content, ProcessError
	}
	content = response.GetHttpContentBytes()
	return content, nil
}
