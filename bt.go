package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
)

func scanBluetooth() error {
	d, err := dev.NewDevice("default")
	if err != nil {
		return fmt.Errorf("can't create new device: %w", err)
	}
	ble.SetDefaultDevice(d)

	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 10*time.Second))

	fmt.Println("Scanning for Bluetooth devices...")
	err = ble.Scan(ctx, true, func(a ble.Advertisement) {
		fmt.Printf("Name: %s, MAC: %s, RSSI: %d dBm\n", a.LocalName(), a.Addr(), a.RSSI())
		err := discoverServicesAndCharacteristics(a.Addr().String())
		if err != nil {
			fmt.Printf("Failed to discover services and characteristics for %s: %v\n", a.Addr(), err)
		}
	}, nil)

	if err != nil && err != context.DeadlineExceeded {
		return fmt.Errorf("scan failed: %w", err)
	}

	return nil
}

func discoverServicesAndCharacteristics(macAddress string) error {
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 10*time.Second))

	client, err := ble.Dial(ctx, ble.NewAddr(macAddress))
	if err != nil {
		return fmt.Errorf("can't dial: %w", err)
	}
	defer client.CancelConnection()

	profile, err := client.DiscoverProfile(true)
	if err != nil {
		return fmt.Errorf("can't discover profile: %w", err)
	}

	for _, service := range profile.Services {
		fmt.Printf("Service: %s\n", service.UUID)
		for _, characteristic := range service.Characteristics {
			fmt.Printf("  Characteristic: %s\n", characteristic.UUID)
		}
	}

	return nil
}

func sendOverBluetooth(data string, macAddress string) error {
	d, err := dev.NewDevice("default")
	if err != nil {
		return fmt.Errorf("can't create new device: %w", err)
	}
	ble.SetDefaultDevice(d)

	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 10*time.Second))

	client, err := ble.Dial(ctx, ble.NewAddr(macAddress))
	if err != nil {
		return fmt.Errorf("can't dial: %w", err)
	}
	defer client.CancelConnection()

	profile, err := client.DiscoverProfile(true)
	if err != nil {
		return fmt.Errorf("can't discover profile: %w", err)
	}

	var serviceUUID, characteristicUUID ble.UUID
	if len(profile.Services) > 0 {
		serviceUUID = profile.Services[0].UUID
		if len(profile.Services[0].Characteristics) > 0 {
			characteristicUUID = profile.Services[0].Characteristics[0].UUID
		} else {
			return fmt.Errorf("no characteristics found in service: %s", serviceUUID)
		}
	} else {
		return fmt.Errorf("no services found")
	}

	var service *ble.Service
	for _, s := range profile.Services {
		if s.UUID.Equal(serviceUUID) {
			service = s
			break
		}
	}
	if service == nil {
		return fmt.Errorf("service not found: %s", serviceUUID)
	}

	var char *ble.Characteristic
	for _, c := range service.Characteristics {
		if c.UUID.Equal(characteristicUUID) {
			char = c
			break
		}
	}
	if char == nil {
		return fmt.Errorf("characteristic not found: %s", characteristicUUID)
	}

	err = client.WriteCharacteristic(char, []byte(data), true)
	if err != nil {
		return fmt.Errorf("failed to write characteristic: %w", err)
	}

	log.Printf("Sent data over Bluetooth: %s", data)
	return nil
}
