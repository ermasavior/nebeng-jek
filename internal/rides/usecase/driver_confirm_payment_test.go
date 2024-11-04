package usecase

import (
	"context"
	"errors"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	pkgError "nebeng-jek/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_DriverConfirmPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	paymentRepoMock := mockRepo.NewMockPaymentRepository(ctrl)
	usecaseMock := NewUsecase(nil, ridesRepoMock, ridesPubsubMock, paymentRepoMock)

	var (
		driverID     = int64(1111)
		riderID      = int64(2222)
		driverMSISDN = "0811111"
		riderMSISDN  = "0811222"
		fare         = float64(20000)
		distance     = float64(3)
		customPrice  = float64(13000)
		netPrice     = customPrice * (1 - model.RideFeeDiscount)
		commission   = customPrice * model.RideFeeDiscount

		rideData = model.RideData{
			RideID:    111,
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideEnded,
			PickupLocation: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Fare:     &fare,
			Distance: &distance,
		}
		req = model.DriverConfirmPaymentRequest{
			RideID:      111,
			CustomPrice: customPrice,
		}

		expectedErr = errors.New("error from repo")
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should confirm ride payment and broadcast to rider", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return(riderMSISDN, nil)

		paymentRepoMock.EXPECT().DeductCredit(ctx, model.DeductCreditRequest{
			MSISDN: riderMSISDN,
			Value:  netPrice,
		}).Return(nil)
		paymentRepoMock.EXPECT().AddCredit(ctx, model.AddCreditRequest{
			MSISDN: driverMSISDN,
			Value:  netPrice,
		}).Return(nil)
		ridesRepoMock.EXPECT().StoreRideCommission(ctx, model.StoreRideCommissionRequest{
			RideID:     req.RideID,
			Commission: commission,
		}).Return(nil)

		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:     req.RideID,
			Status:     model.StatusNumRidePaid,
			FinalPrice: &customPrice,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRidePaid, model.RidePaidMessage{
			RideID:     rideData.RideID,
			RiderID:    rideData.RiderID,
			Distance:   *rideData.Distance,
			FinalPrice: customPrice,
		}).Return(nil)

		rideData, err := usecaseMock.DriverConfirmPayment(ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, customPrice, *rideData.FinalPrice)
		assert.Equal(t, model.StatusNumRidePaid, rideData.StatusNum)
	})

	t.Run("failed - invalid ride data", func(t *testing.T) {
		invalidRide := rideData
		invalidRide.StatusNum = model.StatusNumRideCancelled
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(invalidRide, nil)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})

	t.Run("failed - failed get ride data", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - ride not found", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, constants.ErrorDataNotFound)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrResourceNotFoundCode, err.GetCode())
	})

	t.Run("failed - failed get driver msisdn", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return("", expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - failed get rider msisdn", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return("", expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - failed deduct credit", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return(riderMSISDN, nil)

		paymentRepoMock.EXPECT().DeductCredit(ctx, model.DeductCreditRequest{
			MSISDN: riderMSISDN,
			Value:  netPrice,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - failed add credit", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return(riderMSISDN, nil)

		paymentRepoMock.EXPECT().DeductCredit(ctx, model.DeductCreditRequest{
			MSISDN: riderMSISDN,
			Value:  netPrice,
		}).Return(nil)
		paymentRepoMock.EXPECT().AddCredit(ctx, model.AddCreditRequest{
			MSISDN: driverMSISDN,
			Value:  netPrice,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - failed store commission", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return(riderMSISDN, nil)

		paymentRepoMock.EXPECT().DeductCredit(ctx, model.DeductCreditRequest{
			MSISDN: riderMSISDN,
			Value:  netPrice,
		}).Return(nil)
		paymentRepoMock.EXPECT().AddCredit(ctx, model.AddCreditRequest{
			MSISDN: driverMSISDN,
			Value:  netPrice,
		}).Return(nil)
		ridesRepoMock.EXPECT().StoreRideCommission(ctx, model.StoreRideCommissionRequest{
			RideID:     req.RideID,
			Commission: commission,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - failed update ride data", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return(riderMSISDN, nil)

		paymentRepoMock.EXPECT().DeductCredit(ctx, model.DeductCreditRequest{
			MSISDN: riderMSISDN,
			Value:  netPrice,
		}).Return(nil)
		paymentRepoMock.EXPECT().AddCredit(ctx, model.AddCreditRequest{
			MSISDN: driverMSISDN,
			Value:  netPrice,
		}).Return(nil)
		ridesRepoMock.EXPECT().StoreRideCommission(ctx, model.StoreRideCommissionRequest{
			RideID:     req.RideID,
			Commission: commission,
		}).Return(nil)

		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:     req.RideID,
			Status:     model.StatusNumRidePaid,
			FinalPrice: &customPrice,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - failed broadcast message", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)
		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, riderID).Return(riderMSISDN, nil)

		paymentRepoMock.EXPECT().DeductCredit(ctx, model.DeductCreditRequest{
			MSISDN: riderMSISDN,
			Value:  netPrice,
		}).Return(nil)
		paymentRepoMock.EXPECT().AddCredit(ctx, model.AddCreditRequest{
			MSISDN: driverMSISDN,
			Value:  netPrice,
		}).Return(nil)
		ridesRepoMock.EXPECT().StoreRideCommission(ctx, model.StoreRideCommissionRequest{
			RideID:     req.RideID,
			Commission: commission,
		}).Return(nil)

		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:     req.RideID,
			Status:     model.StatusNumRidePaid,
			FinalPrice: &customPrice,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRidePaid, model.RidePaidMessage{
			RideID:     rideData.RideID,
			RiderID:    rideData.RiderID,
			Distance:   *rideData.Distance,
			FinalPrice: customPrice,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverConfirmPayment(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})
}
