package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cfw/v1/model"
)

type CfwClient struct {
	HcClient *http_client.HcHttpClient
}

func NewCfwClient(hcClient *http_client.HcHttpClient) *CfwClient {
	return &CfwClient{HcClient: hcClient}
}

func CfwClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

// AddAddressItem 添加地址组成员
//
// 添加地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAddressItem(request *model.AddAddressItemRequest) (*model.AddAddressItemResponse, error) {
	requestDef := GenReqDefForAddAddressItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAddressItemResponse), nil
	}
}

// AddAddressItemInvoker 添加地址组成员
func (c *CfwClient) AddAddressItemInvoker(request *model.AddAddressItemRequest) *AddAddressItemInvoker {
	requestDef := GenReqDefForAddAddressItem()
	return &AddAddressItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAddressSet 添加地址组
//
// 添加地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAddressSet(request *model.AddAddressSetRequest) (*model.AddAddressSetResponse, error) {
	requestDef := GenReqDefForAddAddressSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAddressSetResponse), nil
	}
}

// AddAddressSetInvoker 添加地址组
func (c *CfwClient) AddAddressSetInvoker(request *model.AddAddressSetRequest) *AddAddressSetInvoker {
	requestDef := GenReqDefForAddAddressSet()
	return &AddAddressSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddBlackWhiteList 创建黑白名单规则
//
// 创建黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddBlackWhiteList(request *model.AddBlackWhiteListRequest) (*model.AddBlackWhiteListResponse, error) {
	requestDef := GenReqDefForAddBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddBlackWhiteListResponse), nil
	}
}

// AddBlackWhiteListInvoker 创建黑白名单规则
func (c *CfwClient) AddBlackWhiteListInvoker(request *model.AddBlackWhiteListRequest) *AddBlackWhiteListInvoker {
	requestDef := GenReqDefForAddBlackWhiteList()
	return &AddBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDomainSet 添加域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddDomainSet(request *model.AddDomainSetRequest) (*model.AddDomainSetResponse, error) {
	requestDef := GenReqDefForAddDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDomainSetResponse), nil
	}
}

// AddDomainSetInvoker 添加域名组
func (c *CfwClient) AddDomainSetInvoker(request *model.AddDomainSetRequest) *AddDomainSetInvoker {
	requestDef := GenReqDefForAddDomainSet()
	return &AddDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddDomains 添加域名列表
//
// 添加域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddDomains(request *model.AddDomainsRequest) (*model.AddDomainsResponse, error) {
	requestDef := GenReqDefForAddDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddDomainsResponse), nil
	}
}

// AddDomainsInvoker 添加域名列表
func (c *CfwClient) AddDomainsInvoker(request *model.AddDomainsRequest) *AddDomainsInvoker {
	requestDef := GenReqDefForAddDomains()
	return &AddDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddServiceItems 新建服务成员
//
// 批量添加服务组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddServiceItems(request *model.AddServiceItemsRequest) (*model.AddServiceItemsResponse, error) {
	requestDef := GenReqDefForAddServiceItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServiceItemsResponse), nil
	}
}

// AddServiceItemsInvoker 新建服务成员
func (c *CfwClient) AddServiceItemsInvoker(request *model.AddServiceItemsRequest) *AddServiceItemsInvoker {
	requestDef := GenReqDefForAddServiceItems()
	return &AddServiceItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddServiceSet 新建服务组
//
// 创建服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddServiceSet(request *model.AddServiceSetRequest) (*model.AddServiceSetResponse, error) {
	requestDef := GenReqDefForAddServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddServiceSetResponse), nil
	}
}

// AddServiceSetInvoker 新建服务组
func (c *CfwClient) AddServiceSetInvoker(request *model.AddServiceSetRequest) *AddServiceSetInvoker {
	requestDef := GenReqDefForAddServiceSet()
	return &AddServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAddressItems 批量删除地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteAddressItems(request *model.BatchDeleteAddressItemsRequest) (*model.BatchDeleteAddressItemsResponse, error) {
	requestDef := GenReqDefForBatchDeleteAddressItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAddressItemsResponse), nil
	}
}

// BatchDeleteAddressItemsInvoker 批量删除地址组成员
func (c *CfwClient) BatchDeleteAddressItemsInvoker(request *model.BatchDeleteAddressItemsRequest) *BatchDeleteAddressItemsInvoker {
	requestDef := GenReqDefForBatchDeleteAddressItems()
	return &BatchDeleteAddressItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteServiceItems 批量删除服务组成员信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteServiceItems(request *model.BatchDeleteServiceItemsRequest) (*model.BatchDeleteServiceItemsResponse, error) {
	requestDef := GenReqDefForBatchDeleteServiceItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteServiceItemsResponse), nil
	}
}

// BatchDeleteServiceItemsInvoker 批量删除服务组成员信息
func (c *CfwClient) BatchDeleteServiceItemsInvoker(request *model.BatchDeleteServiceItemsRequest) *BatchDeleteServiceItemsInvoker {
	requestDef := GenReqDefForBatchDeleteServiceItems()
	return &BatchDeleteServiceItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEastWestFirewallStatus 修改东西向防火墙防护状态
//
// 东西向防护资源防护开启/关闭
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeEastWestFirewallStatus(request *model.ChangeEastWestFirewallStatusRequest) (*model.ChangeEastWestFirewallStatusResponse, error) {
	requestDef := GenReqDefForChangeEastWestFirewallStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEastWestFirewallStatusResponse), nil
	}
}

// ChangeEastWestFirewallStatusInvoker 修改东西向防火墙防护状态
func (c *CfwClient) ChangeEastWestFirewallStatusInvoker(request *model.ChangeEastWestFirewallStatusRequest) *ChangeEastWestFirewallStatusInvoker {
	requestDef := GenReqDefForChangeEastWestFirewallStatus()
	return &ChangeEastWestFirewallStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddressItem 删除地址组成员
//
// 删除地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAddressItem(request *model.DeleteAddressItemRequest) (*model.DeleteAddressItemResponse, error) {
	requestDef := GenReqDefForDeleteAddressItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressItemResponse), nil
	}
}

// DeleteAddressItemInvoker 删除地址组成员
func (c *CfwClient) DeleteAddressItemInvoker(request *model.DeleteAddressItemRequest) *DeleteAddressItemInvoker {
	requestDef := GenReqDefForDeleteAddressItem()
	return &DeleteAddressItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAddressSet 删除地址组
//
// 删除地址组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAddressSet(request *model.DeleteAddressSetRequest) (*model.DeleteAddressSetResponse, error) {
	requestDef := GenReqDefForDeleteAddressSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAddressSetResponse), nil
	}
}

// DeleteAddressSetInvoker 删除地址组
func (c *CfwClient) DeleteAddressSetInvoker(request *model.DeleteAddressSetRequest) *DeleteAddressSetInvoker {
	requestDef := GenReqDefForDeleteAddressSet()
	return &DeleteAddressSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteBlackWhiteList 删除黑白名单规则
//
// 删除黑白名单规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteBlackWhiteList(request *model.DeleteBlackWhiteListRequest) (*model.DeleteBlackWhiteListResponse, error) {
	requestDef := GenReqDefForDeleteBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteBlackWhiteListResponse), nil
	}
}

// DeleteBlackWhiteListInvoker 删除黑白名单规则
func (c *CfwClient) DeleteBlackWhiteListInvoker(request *model.DeleteBlackWhiteListRequest) *DeleteBlackWhiteListInvoker {
	requestDef := GenReqDefForDeleteBlackWhiteList()
	return &DeleteBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomainSet 删除域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteDomainSet(request *model.DeleteDomainSetRequest) (*model.DeleteDomainSetResponse, error) {
	requestDef := GenReqDefForDeleteDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainSetResponse), nil
	}
}

// DeleteDomainSetInvoker 删除域名组
func (c *CfwClient) DeleteDomainSetInvoker(request *model.DeleteDomainSetRequest) *DeleteDomainSetInvoker {
	requestDef := GenReqDefForDeleteDomainSet()
	return &DeleteDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteDomains 删除域名列表
//
// 删除域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

// DeleteDomainsInvoker 删除域名列表
func (c *CfwClient) DeleteDomainsInvoker(request *model.DeleteDomainsRequest) *DeleteDomainsInvoker {
	requestDef := GenReqDefForDeleteDomains()
	return &DeleteDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceItem 删除服务成员
//
// 删除服务组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteServiceItem(request *model.DeleteServiceItemRequest) (*model.DeleteServiceItemResponse, error) {
	requestDef := GenReqDefForDeleteServiceItem()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceItemResponse), nil
	}
}

// DeleteServiceItemInvoker 删除服务成员
func (c *CfwClient) DeleteServiceItemInvoker(request *model.DeleteServiceItemRequest) *DeleteServiceItemInvoker {
	requestDef := GenReqDefForDeleteServiceItem()
	return &DeleteServiceItemInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteServiceSet 删除服务组
//
// 删除服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteServiceSet(request *model.DeleteServiceSetRequest) (*model.DeleteServiceSetResponse, error) {
	requestDef := GenReqDefForDeleteServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteServiceSetResponse), nil
	}
}

// DeleteServiceSetInvoker 删除服务组
func (c *CfwClient) DeleteServiceSetInvoker(request *model.DeleteServiceSetRequest) *DeleteServiceSetInvoker {
	requestDef := GenReqDefForDeleteServiceSet()
	return &DeleteServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAccessControlLogs 查询访问控制日志
//
// 查询访问控制日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAccessControlLogs(request *model.ListAccessControlLogsRequest) (*model.ListAccessControlLogsResponse, error) {
	requestDef := GenReqDefForListAccessControlLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAccessControlLogsResponse), nil
	}
}

// ListAccessControlLogsInvoker 查询访问控制日志
func (c *CfwClient) ListAccessControlLogsInvoker(request *model.ListAccessControlLogsRequest) *ListAccessControlLogsInvoker {
	requestDef := GenReqDefForListAccessControlLogs()
	return &ListAccessControlLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressItems 查询地址组成员
//
// 查询地址组成员
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressItems(request *model.ListAddressItemsRequest) (*model.ListAddressItemsResponse, error) {
	requestDef := GenReqDefForListAddressItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressItemsResponse), nil
	}
}

// ListAddressItemsInvoker 查询地址组成员
func (c *CfwClient) ListAddressItemsInvoker(request *model.ListAddressItemsRequest) *ListAddressItemsInvoker {
	requestDef := GenReqDefForListAddressItems()
	return &ListAddressItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressSetDetail 查询地址组详细信息
//
// 查询地址组详细
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressSetDetail(request *model.ListAddressSetDetailRequest) (*model.ListAddressSetDetailResponse, error) {
	requestDef := GenReqDefForListAddressSetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressSetDetailResponse), nil
	}
}

// ListAddressSetDetailInvoker 查询地址组详细信息
func (c *CfwClient) ListAddressSetDetailInvoker(request *model.ListAddressSetDetailRequest) *ListAddressSetDetailInvoker {
	requestDef := GenReqDefForListAddressSetDetail()
	return &ListAddressSetDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAddressSets 查询地址组列表
//
// 查询地址组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAddressSets(request *model.ListAddressSetsRequest) (*model.ListAddressSetsResponse, error) {
	requestDef := GenReqDefForListAddressSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAddressSetsResponse), nil
	}
}

// ListAddressSetsInvoker 查询地址组列表
func (c *CfwClient) ListAddressSetsInvoker(request *model.ListAddressSetsRequest) *ListAddressSetsInvoker {
	requestDef := GenReqDefForListAddressSets()
	return &ListAddressSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAttackLogs 查询攻击日志
//
// 查询攻击日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAttackLogs(request *model.ListAttackLogsRequest) (*model.ListAttackLogsResponse, error) {
	requestDef := GenReqDefForListAttackLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAttackLogsResponse), nil
	}
}

// ListAttackLogsInvoker 查询攻击日志
func (c *CfwClient) ListAttackLogsInvoker(request *model.ListAttackLogsRequest) *ListAttackLogsInvoker {
	requestDef := GenReqDefForListAttackLogs()
	return &ListAttackLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListBlackWhiteLists 查询黑白名单列表
//
// 查询黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListBlackWhiteLists(request *model.ListBlackWhiteListsRequest) (*model.ListBlackWhiteListsResponse, error) {
	requestDef := GenReqDefForListBlackWhiteLists()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListBlackWhiteListsResponse), nil
	}
}

// ListBlackWhiteListsInvoker 查询黑白名单列表
func (c *CfwClient) ListBlackWhiteListsInvoker(request *model.ListBlackWhiteListsRequest) *ListBlackWhiteListsInvoker {
	requestDef := GenReqDefForListBlackWhiteLists()
	return &ListBlackWhiteListsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDnsServers 查询dns服务器列表
//
// 查询dns服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDnsServers(request *model.ListDnsServersRequest) (*model.ListDnsServersResponse, error) {
	requestDef := GenReqDefForListDnsServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDnsServersResponse), nil
	}
}

// ListDnsServersInvoker 查询dns服务器列表
func (c *CfwClient) ListDnsServersInvoker(request *model.ListDnsServersRequest) *ListDnsServersInvoker {
	requestDef := GenReqDefForListDnsServers()
	return &ListDnsServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainParseDetail 查询域名解析ip地址
//
// 测试域名有效性
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomainParseDetail(request *model.ListDomainParseDetailRequest) (*model.ListDomainParseDetailResponse, error) {
	requestDef := GenReqDefForListDomainParseDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainParseDetailResponse), nil
	}
}

// ListDomainParseDetailInvoker 查询域名解析ip地址
func (c *CfwClient) ListDomainParseDetailInvoker(request *model.ListDomainParseDetailRequest) *ListDomainParseDetailInvoker {
	requestDef := GenReqDefForListDomainParseDetail()
	return &ListDomainParseDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomainSets 查询域名组列表
//
// 查询域名组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomainSets(request *model.ListDomainSetsRequest) (*model.ListDomainSetsResponse, error) {
	requestDef := GenReqDefForListDomainSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainSetsResponse), nil
	}
}

// ListDomainSetsInvoker 查询域名组列表
func (c *CfwClient) ListDomainSetsInvoker(request *model.ListDomainSetsRequest) *ListDomainSetsInvoker {
	requestDef := GenReqDefForListDomainSets()
	return &ListDomainSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDomains 获取域名组下域名列表
//
// 获取域名组下域名列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}

// ListDomainsInvoker 获取域名组下域名列表
func (c *CfwClient) ListDomainsInvoker(request *model.ListDomainsRequest) *ListDomainsInvoker {
	requestDef := GenReqDefForListDomains()
	return &ListDomainsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEastWestFirewall 获取东西向防火墙信息
//
// 获取东西向防火墙信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEastWestFirewall(request *model.ListEastWestFirewallRequest) (*model.ListEastWestFirewallResponse, error) {
	requestDef := GenReqDefForListEastWestFirewall()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEastWestFirewallResponse), nil
	}
}

// ListEastWestFirewallInvoker 获取东西向防火墙信息
func (c *CfwClient) ListEastWestFirewallInvoker(request *model.ListEastWestFirewallRequest) *ListEastWestFirewallInvoker {
	requestDef := GenReqDefForListEastWestFirewall()
	return &ListEastWestFirewallInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFirewallDetail 查询防火墙详细信息
//
// 查询防火墙实例
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFirewallDetail(request *model.ListFirewallDetailRequest) (*model.ListFirewallDetailResponse, error) {
	requestDef := GenReqDefForListFirewallDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallDetailResponse), nil
	}
}

// ListFirewallDetailInvoker 查询防火墙详细信息
func (c *CfwClient) ListFirewallDetailInvoker(request *model.ListFirewallDetailRequest) *ListFirewallDetailInvoker {
	requestDef := GenReqDefForListFirewallDetail()
	return &ListFirewallDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFirewallList 查询防火墙列表
//
// 查询防火墙列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFirewallList(request *model.ListFirewallListRequest) (*model.ListFirewallListResponse, error) {
	requestDef := GenReqDefForListFirewallList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFirewallListResponse), nil
	}
}

// ListFirewallListInvoker 查询防火墙列表
func (c *CfwClient) ListFirewallListInvoker(request *model.ListFirewallListRequest) *ListFirewallListInvoker {
	requestDef := GenReqDefForListFirewallList()
	return &ListFirewallListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListFlowLogs 查询流日志
//
// 查询流日志
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListFlowLogs(request *model.ListFlowLogsRequest) (*model.ListFlowLogsResponse, error) {
	requestDef := GenReqDefForListFlowLogs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListFlowLogsResponse), nil
	}
}

// ListFlowLogsInvoker 查询流日志
func (c *CfwClient) ListFlowLogsInvoker(request *model.ListFlowLogsRequest) *ListFlowLogsInvoker {
	requestDef := GenReqDefForListFlowLogs()
	return &ListFlowLogsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListProtectedVpcs 查询防护VPC数
//
// 查询防护vpc信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListProtectedVpcs(request *model.ListProtectedVpcsRequest) (*model.ListProtectedVpcsResponse, error) {
	requestDef := GenReqDefForListProtectedVpcs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListProtectedVpcsResponse), nil
	}
}

// ListProtectedVpcsInvoker 查询防护VPC数
func (c *CfwClient) ListProtectedVpcsInvoker(request *model.ListProtectedVpcsRequest) *ListProtectedVpcsInvoker {
	requestDef := GenReqDefForListProtectedVpcs()
	return &ListProtectedVpcsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceItems 查询服务成员列表
//
// 查询服务组成员列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceItems(request *model.ListServiceItemsRequest) (*model.ListServiceItemsResponse, error) {
	requestDef := GenReqDefForListServiceItems()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceItemsResponse), nil
	}
}

// ListServiceItemsInvoker 查询服务成员列表
func (c *CfwClient) ListServiceItemsInvoker(request *model.ListServiceItemsRequest) *ListServiceItemsInvoker {
	requestDef := GenReqDefForListServiceItems()
	return &ListServiceItemsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceSetDetail 查询服务组详情
//
// 查询服务组细节
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceSetDetail(request *model.ListServiceSetDetailRequest) (*model.ListServiceSetDetailResponse, error) {
	requestDef := GenReqDefForListServiceSetDetail()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceSetDetailResponse), nil
	}
}

// ListServiceSetDetailInvoker 查询服务组详情
func (c *CfwClient) ListServiceSetDetailInvoker(request *model.ListServiceSetDetailRequest) *ListServiceSetDetailInvoker {
	requestDef := GenReqDefForListServiceSetDetail()
	return &ListServiceSetDetailInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListServiceSets 获取服务组列表
//
// 获取服务组列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListServiceSets(request *model.ListServiceSetsRequest) (*model.ListServiceSetsResponse, error) {
	requestDef := GenReqDefForListServiceSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListServiceSetsResponse), nil
	}
}

// ListServiceSetsInvoker 获取服务组列表
func (c *CfwClient) ListServiceSetsInvoker(request *model.ListServiceSetsRequest) *ListServiceSetsInvoker {
	requestDef := GenReqDefForListServiceSets()
	return &ListServiceSetsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAddressSet 更新地址组信息
//
// 更新地址组信息
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAddressSet(request *model.UpdateAddressSetRequest) (*model.UpdateAddressSetResponse, error) {
	requestDef := GenReqDefForUpdateAddressSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAddressSetResponse), nil
	}
}

// UpdateAddressSetInvoker 更新地址组信息
func (c *CfwClient) UpdateAddressSetInvoker(request *model.UpdateAddressSetRequest) *UpdateAddressSetInvoker {
	requestDef := GenReqDefForUpdateAddressSet()
	return &UpdateAddressSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateBlackWhiteList 更新黑白名单列表
//
// 更新黑白名单列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateBlackWhiteList(request *model.UpdateBlackWhiteListRequest) (*model.UpdateBlackWhiteListResponse, error) {
	requestDef := GenReqDefForUpdateBlackWhiteList()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateBlackWhiteListResponse), nil
	}
}

// UpdateBlackWhiteListInvoker 更新黑白名单列表
func (c *CfwClient) UpdateBlackWhiteListInvoker(request *model.UpdateBlackWhiteListRequest) *UpdateBlackWhiteListInvoker {
	requestDef := GenReqDefForUpdateBlackWhiteList()
	return &UpdateBlackWhiteListInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDnsServers 更新dns服务器列表
//
// 更新dns服务器列表
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateDnsServers(request *model.UpdateDnsServersRequest) (*model.UpdateDnsServersResponse, error) {
	requestDef := GenReqDefForUpdateDnsServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDnsServersResponse), nil
	}
}

// UpdateDnsServersInvoker 更新dns服务器列表
func (c *CfwClient) UpdateDnsServersInvoker(request *model.UpdateDnsServersRequest) *UpdateDnsServersInvoker {
	requestDef := GenReqDefForUpdateDnsServers()
	return &UpdateDnsServersInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateDomainSet 更新域名组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateDomainSet(request *model.UpdateDomainSetRequest) (*model.UpdateDomainSetResponse, error) {
	requestDef := GenReqDefForUpdateDomainSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDomainSetResponse), nil
	}
}

// UpdateDomainSetInvoker 更新域名组
func (c *CfwClient) UpdateDomainSetInvoker(request *model.UpdateDomainSetRequest) *UpdateDomainSetInvoker {
	requestDef := GenReqDefForUpdateDomainSet()
	return &UpdateDomainSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateServiceSet 修改服务组
//
// 更新服务组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateServiceSet(request *model.UpdateServiceSetRequest) (*model.UpdateServiceSetResponse, error) {
	requestDef := GenReqDefForUpdateServiceSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateServiceSetResponse), nil
	}
}

// UpdateServiceSetInvoker 修改服务组
func (c *CfwClient) UpdateServiceSetInvoker(request *model.UpdateServiceSetRequest) *UpdateServiceSetInvoker {
	requestDef := GenReqDefForUpdateServiceSet()
	return &UpdateServiceSetInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// AddAclRule 创建ACL规则
//
// 创建ACL规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) AddAclRule(request *model.AddAclRuleRequest) (*model.AddAclRuleResponse, error) {
	requestDef := GenReqDefForAddAclRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AddAclRuleResponse), nil
	}
}

// AddAclRuleInvoker 创建ACL规则
func (c *CfwClient) AddAclRuleInvoker(request *model.AddAclRuleRequest) *AddAclRuleInvoker {
	requestDef := GenReqDefForAddAclRule()
	return &AddAclRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchDeleteAclRules 批量删除Acl规则
//
// 批量删除Acl规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchDeleteAclRules(request *model.BatchDeleteAclRulesRequest) (*model.BatchDeleteAclRulesResponse, error) {
	requestDef := GenReqDefForBatchDeleteAclRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchDeleteAclRulesResponse), nil
	}
}

// BatchDeleteAclRulesInvoker 批量删除Acl规则
func (c *CfwClient) BatchDeleteAclRulesInvoker(request *model.BatchDeleteAclRulesRequest) *BatchDeleteAclRulesInvoker {
	requestDef := GenReqDefForBatchDeleteAclRules()
	return &BatchDeleteAclRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// BatchUpdateAclRuleActions 批量更新规则动作
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) BatchUpdateAclRuleActions(request *model.BatchUpdateAclRuleActionsRequest) (*model.BatchUpdateAclRuleActionsResponse, error) {
	requestDef := GenReqDefForBatchUpdateAclRuleActions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchUpdateAclRuleActionsResponse), nil
	}
}

// BatchUpdateAclRuleActionsInvoker 批量更新规则动作
func (c *CfwClient) BatchUpdateAclRuleActionsInvoker(request *model.BatchUpdateAclRuleActionsRequest) *BatchUpdateAclRuleActionsInvoker {
	requestDef := GenReqDefForBatchUpdateAclRuleActions()
	return &BatchUpdateAclRuleActionsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclRule 删除ACL规则
//
// 删除ACL规则组
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAclRule(request *model.DeleteAclRuleRequest) (*model.DeleteAclRuleResponse, error) {
	requestDef := GenReqDefForDeleteAclRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclRuleResponse), nil
	}
}

// DeleteAclRuleInvoker 删除ACL规则
func (c *CfwClient) DeleteAclRuleInvoker(request *model.DeleteAclRuleRequest) *DeleteAclRuleInvoker {
	requestDef := GenReqDefForDeleteAclRule()
	return &DeleteAclRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// DeleteAclRuleHitCount 删除规则击中次数
//
// 清除规则击中次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) DeleteAclRuleHitCount(request *model.DeleteAclRuleHitCountRequest) (*model.DeleteAclRuleHitCountResponse, error) {
	requestDef := GenReqDefForDeleteAclRuleHitCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteAclRuleHitCountResponse), nil
	}
}

// DeleteAclRuleHitCountInvoker 删除规则击中次数
func (c *CfwClient) DeleteAclRuleHitCountInvoker(request *model.DeleteAclRuleHitCountRequest) *DeleteAclRuleHitCountInvoker {
	requestDef := GenReqDefForDeleteAclRuleHitCount()
	return &DeleteAclRuleHitCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclRuleHitCount 获取规则击中次数
//
// 获取规则击中次数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAclRuleHitCount(request *model.ListAclRuleHitCountRequest) (*model.ListAclRuleHitCountResponse, error) {
	requestDef := GenReqDefForListAclRuleHitCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclRuleHitCountResponse), nil
	}
}

// ListAclRuleHitCountInvoker 获取规则击中次数
func (c *CfwClient) ListAclRuleHitCountInvoker(request *model.ListAclRuleHitCountRequest) *ListAclRuleHitCountInvoker {
	requestDef := GenReqDefForListAclRuleHitCount()
	return &ListAclRuleHitCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListAclRules 查询防护规则
//
// 查询防护规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListAclRules(request *model.ListAclRulesRequest) (*model.ListAclRulesResponse, error) {
	requestDef := GenReqDefForListAclRules()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListAclRulesResponse), nil
	}
}

// ListAclRulesInvoker 查询防护规则
func (c *CfwClient) ListAclRulesInvoker(request *model.ListAclRulesRequest) *ListAclRulesInvoker {
	requestDef := GenReqDefForListAclRules()
	return &ListAclRulesInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListRuleAclTags 查询规则标签
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListRuleAclTags(request *model.ListRuleAclTagsRequest) (*model.ListRuleAclTagsResponse, error) {
	requestDef := GenReqDefForListRuleAclTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRuleAclTagsResponse), nil
	}
}

// ListRuleAclTagsInvoker 查询规则标签
func (c *CfwClient) ListRuleAclTagsInvoker(request *model.ListRuleAclTagsRequest) *ListRuleAclTagsInvoker {
	requestDef := GenReqDefForListRuleAclTags()
	return &ListRuleAclTagsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAclRule 更新ACL规则
//
// 更新ACL规则
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAclRule(request *model.UpdateAclRuleRequest) (*model.UpdateAclRuleResponse, error) {
	requestDef := GenReqDefForUpdateAclRule()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclRuleResponse), nil
	}
}

// UpdateAclRuleInvoker 更新ACL规则
func (c *CfwClient) UpdateAclRuleInvoker(request *model.UpdateAclRuleRequest) *UpdateAclRuleInvoker {
	requestDef := GenReqDefForUpdateAclRule()
	return &UpdateAclRuleInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAclRuleOrder ACL防护规则优先级设置
//
// # ACL防护规则优先级设置
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) UpdateAclRuleOrder(request *model.UpdateAclRuleOrderRequest) (*model.UpdateAclRuleOrderResponse, error) {
	requestDef := GenReqDefForUpdateAclRuleOrder()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAclRuleOrderResponse), nil
	}
}

// UpdateAclRuleOrderInvoker ACL防护规则优先级设置
func (c *CfwClient) UpdateAclRuleOrderInvoker(request *model.UpdateAclRuleOrderRequest) *UpdateAclRuleOrderInvoker {
	requestDef := GenReqDefForUpdateAclRuleOrder()
	return &UpdateAclRuleOrderInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeEipStatus 弹性IP开启关闭
//
// 开启关闭EIP,客户购买EIP后首次开启EIP防护前需使用ListEips同步EIP资产，sync字段设置为1。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeEipStatus(request *model.ChangeEipStatusRequest) (*model.ChangeEipStatusResponse, error) {
	requestDef := GenReqDefForChangeEipStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeEipStatusResponse), nil
	}
}

// ChangeEipStatusInvoker 弹性IP开启关闭
func (c *CfwClient) ChangeEipStatusInvoker(request *model.ChangeEipStatusRequest) *ChangeEipStatusInvoker {
	requestDef := GenReqDefForChangeEipStatus()
	return &ChangeEipStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEipCount 查询Eip个数
//
// 查询Eip个数
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEipCount(request *model.ListEipCountRequest) (*model.ListEipCountResponse, error) {
	requestDef := GenReqDefForListEipCount()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEipCountResponse), nil
	}
}

// ListEipCountInvoker 查询Eip个数
func (c *CfwClient) ListEipCountInvoker(request *model.ListEipCountRequest) *ListEipCountInvoker {
	requestDef := GenReqDefForListEipCount()
	return &ListEipCountInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListEips 弹性IP列表查询
//
// 弹性IP列表查询
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListEips(request *model.ListEipsRequest) (*model.ListEipsResponse, error) {
	requestDef := GenReqDefForListEips()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListEipsResponse), nil
	}
}

// ListEipsInvoker 弹性IP列表查询
func (c *CfwClient) ListEipsInvoker(request *model.ListEipsRequest) *ListEipsInvoker {
	requestDef := GenReqDefForListEips()
	return &ListEipsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsProtectMode 切换防护模式
//
// 切换防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsProtectMode(request *model.ChangeIpsProtectModeRequest) (*model.ChangeIpsProtectModeResponse, error) {
	requestDef := GenReqDefForChangeIpsProtectMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsProtectModeResponse), nil
	}
}

// ChangeIpsProtectModeInvoker 切换防护模式
func (c *CfwClient) ChangeIpsProtectModeInvoker(request *model.ChangeIpsProtectModeRequest) *ChangeIpsProtectModeInvoker {
	requestDef := GenReqDefForChangeIpsProtectMode()
	return &ChangeIpsProtectModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ChangeIpsSwitchStatus IPS特性开关操作
//
// 切换开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ChangeIpsSwitchStatus(request *model.ChangeIpsSwitchStatusRequest) (*model.ChangeIpsSwitchStatusResponse, error) {
	requestDef := GenReqDefForChangeIpsSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ChangeIpsSwitchStatusResponse), nil
	}
}

// ChangeIpsSwitchStatusInvoker IPS特性开关操作
func (c *CfwClient) ChangeIpsSwitchStatusInvoker(request *model.ChangeIpsSwitchStatusRequest) *ChangeIpsSwitchStatusInvoker {
	requestDef := GenReqDefForChangeIpsSwitchStatus()
	return &ChangeIpsSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsProtectMode 查询防护模式
//
// 查询防护模式
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsProtectMode(request *model.ListIpsProtectModeRequest) (*model.ListIpsProtectModeResponse, error) {
	requestDef := GenReqDefForListIpsProtectMode()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsProtectModeResponse), nil
	}
}

// ListIpsProtectModeInvoker 查询防护模式
func (c *CfwClient) ListIpsProtectModeInvoker(request *model.ListIpsProtectModeRequest) *ListIpsProtectModeInvoker {
	requestDef := GenReqDefForListIpsProtectMode()
	return &ListIpsProtectModeInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListIpsSwitchStatus 查询IPS特性开关状态
//
// 查询IPS特性开关状态
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *CfwClient) ListIpsSwitchStatus(request *model.ListIpsSwitchStatusRequest) (*model.ListIpsSwitchStatusResponse, error) {
	requestDef := GenReqDefForListIpsSwitchStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListIpsSwitchStatusResponse), nil
	}
}

// ListIpsSwitchStatusInvoker 查询IPS特性开关状态
func (c *CfwClient) ListIpsSwitchStatusInvoker(request *model.ListIpsSwitchStatusRequest) *ListIpsSwitchStatusInvoker {
	requestDef := GenReqDefForListIpsSwitchStatus()
	return &ListIpsSwitchStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
