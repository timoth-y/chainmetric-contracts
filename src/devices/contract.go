package main

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
	"github.com/rs/xid"

	"github.com/timoth-y/chainmetric-core/models"

	"github.com/timoth-y/chainmetric-contracts/model/request"
	"github.com/timoth-y/chainmetric-contracts/shared"
)

type DevicesContract struct {
	contractapi.Contract
}

func NewDevicesContact() *DevicesContract {
	return &DevicesContract{}
}

func (c *DevicesContract) Retrieve(ctx contractapi.TransactionContextInterface, id string) (*models.Device, error) {
	data, err := ctx.GetStub().GetState(id); if err != nil {
		err = errors.Wrap(err, "failed to read from world state")
		shared.Logger.Error(err)
		return nil, err
	}

	if data == nil {
		return nil, fmt.Errorf("the device %s does not exist", id)
	}

	return models.Device{}.Decode(data)
}

func (c *DevicesContract) All(ctx contractapi.TransactionContextInterface) ([]*models.Device, error) {
	iterator, err := ctx.GetStub().GetStateByPartialCompositeKey("device", []string{})
	if err != nil {
		err = errors.Wrap(err, "failed to read from world state")
		shared.Logger.Error(err)
		return nil, err
	}

	var devices []*models.Device
	for iterator.HasNext() {
		result, err := iterator.Next(); if err != nil {
			shared.Logger.Error(err)
			continue
		}

		device, err := models.Device{}.Decode(result.Value); if err != nil {
			shared.Logger.Error(err)
			continue
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (c *DevicesContract) Register(ctx contractapi.TransactionContextInterface, data string) (string, error) {
	var (
		device = &models.Device{}
		err error
		event = "updated"
	)

	if device, err = device.Decode([]byte(data)); err != nil {
		err = errors.Wrap(err, "failed to deserialize input")
		shared.Logger.Error(err)
		return "", err
	}

	if len(device.ID) == 0 {
		event = "inserted"

		if device.ID, err = generateCompositeKey(ctx, device); err != nil {
			err = errors.Wrap(err, "failed to generate composite key")
			shared.Logger.Error(err)
			return "", err
		}
	}

	if err = device.Validate(); err != nil {
		return "", errors.Wrap(err, "device is not valid")
	}

	if err := c.save(ctx, device, event); err != nil {
		err = errors.Wrap(err, "failed saving device")
		shared.Logger.Error(err)
		return "", err
	}

	return device.ID, nil
}

func (c *DevicesContract) Update(ctx contractapi.TransactionContextInterface, id string, data string) (*models.Device, error) {
	if len(id) == 0 {
		return nil, errors.New("device id must be provided in order to update one")
	}

	device, err := c.Retrieve(ctx, id); if err != nil {
		return nil, err
	}

	req, err := request.DeviceUpdateRequest{}.Decode([]byte(data)); if err != nil {
		err = errors.Wrap(err, "failed to deserialize input")
		shared.Logger.Error(err)
		return nil, err
	}

	req.Update(device)

	if err = device.Validate(); err != nil {
		return nil, errors.Wrap(err, "device is not valid")
	}

	if err := c.save(ctx, device, "updated"); err != nil {
		err = errors.Wrap(err, "failed updating device")
		shared.Logger.Error(err)
		return nil, err
	}

	return device, nil
}

func (c *DevicesContract) Exists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	data, err := ctx.GetStub().GetState(id); if err != nil {
		return false, err
	}
	return data != nil, nil
}

func (c *DevicesContract) Unbind(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := c.Exists(ctx, id); if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("the device with ID %q does not exist", id)
	}

	if err = ctx.GetStub().DelState(id); err != nil {
		return err
	}

	ctx.GetStub().SetEvent("devices.removed", models.Device{ID: id}.Encode())

	return nil
}

func (c *DevicesContract) RemoveAll(ctx contractapi.TransactionContextInterface) error {
	iterator, err := ctx.GetStub().GetStateByPartialCompositeKey("device", []string{})
	if err != nil {
		err = errors.Wrap(err, "failed to read from world state")
		shared.Logger.Error(err)
		return err
	}

	for iterator.HasNext() {
		result, err := iterator.Next(); if err != nil {
			shared.Logger.Error(err)
			continue
		}

		if err = ctx.GetStub().DelState(result.Key); err != nil {
			shared.Logger.Error(err)
			continue
		}

		ctx.GetStub().SetEvent("devices.removed", models.Device{ID: result.Key}.Encode())
	}
	return nil
}

func (c *DevicesContract) save(ctx contractapi.TransactionContextInterface, device *models.Device, events ...string) error {
	if len(device.ID) == 0 {
		return errors.New("the unique id must be defined for device")
	}

	if err := ctx.GetStub().PutState(device.ID, device.Encode()); err != nil {
		return err
	}

	if len(events) != 0 {
		for _, event := range events {
			ctx.GetStub().SetEvent(fmt.Sprintf("devices.%s", event), device.Encode())
		}
	}

	return nil
}

func generateCompositeKey(ctx contractapi.TransactionContextInterface, dev *models.Device) (string, error) {
	return ctx.GetStub().CreateCompositeKey("device", []string{
		shared.Hash(dev.Hostname),
		xid.NewWithTime(time.Now()).String(),
	})
}
