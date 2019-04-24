package cce

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCreateCluster(t *testing.T) {
	clusterClient := NewTestClient()
	args := &CreateClusterArgs{
		ClusterName:       "test-1",
		Version:           "1.13.4",
		MainAvailableZone: "zoneC",
		ContainerNet:      "172.18.0.0/16",
		AdvancedOptions: &AdvancedOptions{
			KubeProxyMode:         "iptables",
			SecureContainerEnable: false,
		},
		//CDSPreMountInfo: &PreMountInfo{
		//	MountPath: "/data",
		//	CdsConfig: []DiskSizeConfig{
		//		{
		//			Size:       "50",
		//			VolumeType: "sata",
		//		},
		//	},
		//},
		OrderContent: &BaseCreateOrderRequestVo{
			Items: []Item{},
		},
	}
	bccRequest := &BCCConfig{
		LogicalZone:       "zoneC",
		InstanceType:      10,
		CPU:               4,
		Memory:            8,
		ImageType:         "common",
		OSType:            "linux",
		OSVersion:         "16.04 LTS amd64 (64bit)",
		SecurityGroupID:   "0cc217c3-fcde-403f-9a84-c066aa30fd1b",
		SecurityGroupName: "CCE默认安全组",
		AdminPass:         "1qaz!QAZ",
		AdminPassConfirm:  "1qaz!QAZ",
		EBSSize:           []int64{},
		ImageID:           "m-DpgNg8lO",
		SubnetUUID:        "a0919779-7cee-4e90-87ba-8e9f7b331fd9",
		IfBuyEIP:          false,
		EIPName:           "test",
		BandwidthInMbps:   1,
		AutoRenew:         false,
	}
	bccRequest.ProductType = PostPayProductType
	bccRequest.Region = os.Getenv("Region")
	bccRequest.PurchaseNum = 2
	bccRequest.ServiceType = InstanceServiceType
	bccRequest.SubProductType = "netraffic"

	args.OrderContent.Items = append(args.OrderContent.Items, Item{
		Config: bccRequest,
	})

	//eipRequest := &EIPConfig{
	//	Name: "test",
	//	BandwidthInMbps: 1,
	//}
	//eipRequest.ProductType = PostPayProductType
	//eipRequest.Region = os.Getenv("Region")
	//eipRequest.PurchaseNum = 2
	//eipRequest.ServiceType = EIPServiceType
	//eipRequest.SubProductType = "netraffic"
	//
	//args.OrderContent.Items = append(args.OrderContent.Items, Item{
	//	Config: eipRequest,
	//})

	//cdsRequest := &CDSConfig{
	//	LogicalZone: "zoneC",
	//	CDSDiskSize: []DiskSizeConfig{
	//		{
	//			Size:       "50",
	//			VolumeType: "sata",
	//		},
	//	},
	//}
	//cdsRequest.ProductType = PostPayProductType
	//cdsRequest.Region = os.Getenv("Region")
	//cdsRequest.PurchaseNum = 2
	//cdsRequest.ServiceType = CDSServiceType
	//cdsRequest.SubProductType = "netraffic"
	//
	//args.OrderContent.Items = append(args.OrderContent.Items, Item{
	//	Config: cdsRequest,
	//})

	s, err := json.Marshal(args)
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Printf("%v", string(s))

	clusterID, err := clusterClient.CreateCluster(args, nil)
	if err != nil {
		t.Fatalf("Failed to create cluster, err: %+v \n", err)
	}
	t.Logf("Created cluster: %+v \n", clusterID)
	time.Sleep(5 * time.Second)
	ins, err := clusterClient.DescribeCluster(clusterID, nil)
	if err != nil {
		t.Fatalf("Describe cluster err: %+v \n", err)
	}
	t.Logf("cluster info: %+v \n", ins)
}

func TestGetClusters(t *testing.T) {
	clusterClient := NewTestClient()
	listArgs := &GetClustersArgs{
		Marker:  "-1",
		MaxKeys: 1000,
	}
	l, err := clusterClient.GetClusters(listArgs, nil)
	if err != nil {
		t.Errorf("Failed to get cluster list, err: %+v", err)
	}
	t.Logf("Get cluster list : %+v \n", l)
}

func TestDescribeCluster(t *testing.T) {
	clusterClient := NewTestClient()
	clusterID := "c-qBH9sE28"
	c, err := clusterClient.DescribeCluster(clusterID, nil)
	if err != nil {
		t.Errorf("Failed to get cluster %v, err: %+v", clusterID, err)
	}
	t.Logf("Get cluster info : %+v \n", c)
}

func TestGetNodeList(t *testing.T) {
	clusterID := "c-ocHPQIne"
	clusterClient := NewTestClient()
	args := &GetClusterNodesArgs{
		Marker:      "-1",
		MaxKeys:     1000,
		ClusterUUID: clusterID,
	}
	nodeList, err := clusterClient.GetNodeList(args, nil)
	if err != nil {
		t.Errorf("Failed to get node list of cluster %v, err: %+v", clusterID, err)
	}
	t.Logf("Get cluster node list : %+v \n", nodeList)
}

func TestScalingDownCluster(t *testing.T) {
	clusterID := "c-ocHPQIne"
	client := NewTestClient()
	args := &ScalingDownClusterArgs{
		ClusterUUID: clusterID,
		NodeInfo: []NodeInfo{
			{
				InstanceID: "i-Kx6YEod8",
			},
		},
	}
	err := client.ScalingDownCluster(args, nil)
	if err != nil {
		t.Fatalf("Scale down cluster error %v", err)
	}
	c, err := client.DescribeCluster(clusterID, nil)
	if err != nil {
		t.Fatalf("get cluster %v error %v", clusterID, err)
	}
	t.Logf("Scale down cluster success %v", c)
}

func TestScalingUpCluster(t *testing.T) {
	client := NewTestClient()
	args := &ScalingUpClusterArgs{
		ClusterUUID: "c-ocHPQIne",
		OrderContent: &BaseCreateOrderRequestVo{
			Items: []Item{},
		},
	}
	bccRequest := &BCCConfig{
		LogicalZone:       "zoneC",
		InstanceType:      10,
		CPU:               4,
		Memory:            8,
		ImageType:         "common",
		OSType:            "linux",
		OSVersion:         "16.04 LTS amd64 (64bit)",
		SecurityGroupID:   "0cc217c3-fcde-403f-9a84-c066aa30fd1b",
		SecurityGroupName: "CCE默认安全组",
		AdminPass:         "1qaz!QAZ",
		AdminPassConfirm:  "1qaz!QAZ",
		ImageID:           "f82c1177-5a24-41a0-89eb-1c926f605025",
		SubnetUUID:        "a0919779-7cee-4e90-87ba-8e9f7b331fd9",
		IfBuyEIP:          false,
		AutoRenew:         false,
	}
	bccRequest.ProductType = PostPayProductType
	bccRequest.Region = os.Getenv("Region")
	bccRequest.PurchaseNum = 1
	bccRequest.ServiceType = InstanceServiceType
	bccRequest.SubProductType = "netraffic"

	args.OrderContent.Items = append(args.OrderContent.Items, Item{
		Config: bccRequest,
	})
	clusterID, err := client.ScalingUpCluster(args, nil)
	if err != nil {
		t.Fatalf("Failed to scalingUp cluster %v, err: %+v", clusterID, err)
	}
	time.Sleep(5 * time.Second)
	ins, err := client.DescribeCluster(clusterID, nil)
	if err != nil {
		t.Fatalf("Failed to get cluster %v, err: %+v", clusterID, err)
	}
	t.Logf("cluster info: %+v", ins)
	nodeArgs := &GetClusterNodesArgs{
		Marker:      "-1",
		MaxKeys:     1000,
		ClusterUUID: clusterID,
	}
	nodeList, err := client.GetNodeList(nodeArgs, nil)
	if err != nil {
		t.Fatalf("Failed to get node list of cluster %v, err: %+v", clusterID, err)
	}
	t.Logf("Get cluster node list : %+v \n", nodeList)
}

func TestDeleteCluster(t *testing.T) {
	client := NewTestClient()
	clusterID := "c-fHtQ8kzq"
	err := client.DeleteCluster(clusterID, nil)
	if err != nil {
		t.Errorf("Failed to delete cluster %v, err: %+v", clusterID, err)
	}
}

func TestGetKubeConfig(t *testing.T) {
	client := NewTestClient()
	clusterID := "c-Q8APKWjk"
	kubeconfig, err := client.GetKubeConfig(clusterID, nil)
	if err != nil {
		t.Fatalf("Failed to get kubeconfig file for cluster %v, err: %+v", clusterID, err)
	}
	fmt.Printf("config file for cluster %v is :: %v", clusterID, kubeconfig)
}

func TestClusterUpgrade(t *testing.T) {
	client := NewTestClient()
	clusterID := "c-AiTOMmrt"
	err := client.ClusterUpgrade(clusterID, "1.13.4", nil)
	if err != nil {
		t.Errorf("Failed to upgrade cluster, err: %+v", err)
	}
}
