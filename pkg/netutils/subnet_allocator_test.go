package netutils

import (
	"fmt"
	"net"
	"testing"
)

func TestAllocateSubnet(t *testing.T) {
	sna, err := NewSubnetAllocator("10.1.0.0/16", 8, nil)
	if err != nil {
		t.Fatal("Failed to initialize subnet allocator: ", err)
	}

	sn, err := sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.0.0/24" {
		t.Fatalf("Did not get expected subnet (n=%d, sn=%s)", 0, sn.String())
	}
	sn, err = sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.1.0/24" {
		t.Fatalf("Did not get expected subnet (n=%d, sn=%s)", 1, sn.String())
	}
	sn, err = sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.2.0/24" {
		t.Fatalf("Did not get expected subnet (n=%d, sn=%s)", 2, sn.String())
	}
}

func TestAllocateSubnetInUse(t *testing.T) {
	inUse := []string{"10.1.0.0/24", "10.1.2.0/24", "10.2.2.2/24", "Invalid"}
	sna, err := NewSubnetAllocator("10.1.0.0/16", 8, inUse)
	if err != nil {
		t.Fatal("Failed to initialize IP allocator: ", err)
	}

	sn, err := sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.1.0/24" {
		t.Fatalf("Did not get expected subnet (sn=%s)", sn.String())
	}
	sn, err = sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.3.0/24" {
		t.Fatalf("Did not get expected subnet (sn=%s)", sn.String())
	}
}

func TestAllocateReleaseSubnet(t *testing.T) {
	sna, err := NewSubnetAllocator("10.1.0.0/16", 14, nil)
	if err != nil {
		t.Fatal("Failed to initialize IP allocator: ", err)
	}

	var releaseSn *net.IPNet

	for i := 0; i < 4; i++ {
		sn, err := sna.GetNetwork()
		if err != nil {
			t.Fatal("Failed to get network: ", err)
		}
		if sn.String() != fmt.Sprintf("10.1.%d.0/18", i*64) {
			t.Fatalf("Did not get expected subnet (i=%d, sn=%s)", i, sn.String())
		}
		if i == 2 {
			releaseSn = sn
		}
	}

	sn, err := sna.GetNetwork()
	if err == nil {
		t.Fatalf("Unexpectedly succeeded in getting network (sn=%s)", sn.String())
	}

	if err := sna.ReleaseNetwork(releaseSn); err != nil {
		t.Fatal("Failed to release the subnet: ", err)
	}

	sn, err = sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != releaseSn.String() {
		t.Fatalf("Did not get expected subnet (sn=%s)", sn.String())
	}

	sn, err = sna.GetNetwork()
	if err == nil {
		t.Fatalf("Unexpectedly succeeded in getting network (sn=%s)", sn.String())
	}
}

func TestGenerateGateway(t *testing.T) {
	sna, err := NewSubnetAllocator("10.1.0.0/16", 8, nil)
	if err != nil {
		t.Fatal("Failed to initialize IP allocator: ", err)
	}

	sn, err := sna.GetNetwork()
	if err != nil {
		t.Fatal("Failed to get network: ", err)
	}
	if sn.String() != "10.1.0.0/24" {
		t.Fatalf("Did not get expected subnet (sn=%s)", sn.String())
	}

	gatewayIP := GenerateDefaultGateway(sn)
	if gatewayIP.String() != "10.1.0.1" {
		t.Fatalf("Did not get expected gateway IP Address (gatewayIP=%s)", gatewayIP.String())
	}
}
