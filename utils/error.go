package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"runtime"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrMiddleware(err error, msgName string, uid string, logger *slog.Logger, otherMsgs map[string]string, callerId ...int) error {
	cId := 1
	if callerId != nil || len(callerId) != 0 {
		cId = callerId[0]
	}
	_, f, l, _ := runtime.Caller(cId)
	line := fmt.Sprintf("%s:%d", f, l)

	msgs := []any{}
	msgs = append(msgs, "source")
	msgs = append(msgs, line)
	msgs = append(msgs, "error")
	msgs = append(msgs, err.Error())
	msgs = append(msgs, "x-required-id")
	msgs = append(msgs, uid)
	for k, v := range otherMsgs {
		msgs = append(msgs, k)
		msgs = append(msgs, v)
	}
	s, ok := status.FromError(err)
	if !ok {
		logger.Error(msgName, msgs...)
		err = status.Error(codes.Internal, codes.Internal.String())
		return err
	}

	switch s.Code() {
	case codes.Canceled:
		logger.Error(msgName, msgs...)
		err = status.Error(s.Code(), s.Code().String())
		return err
	case codes.Unknown, codes.Internal:
		logger.Error(msgName, msgs...)
		err = status.Error(codes.Internal, codes.Internal.String())
		return err
	case codes.InvalidArgument:
		logger.Debug(msgName, msgs...)
		// err = status.Error(s.Code(), s.Message())
		return err
	case codes.NotFound:
		logger.Debug(msgName, msgs...)
		// err = status.Error(s.Code(), err.Error())
		return err
	default:
		logger.Debug(msgName, msgs...)
		err = status.Error(s.Code(), s.Code().String())
		return err
	}
}

func MongoErrMiddleware(err error, msgName, uid string, logger *slog.Logger, otherMsgs map[string]string) error {
	if mongo.IsDuplicateKeyError(err) {
		err = status.Error(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		err = status.Error(codes.NotFound, err.Error()+" - "+otherMsgs["item"])
	}

	if IsValidationError(err) {
		err = status.Error(codes.InvalidArgument, err.Error())
	}

	return ErrMiddleware(err, msgName, uid, logger, otherMsgs, 2)
}

// IsValidationError - returns true if mongo validation failed
func IsValidationError(err error) bool {
	if se := mongo.ServerError(nil); errors.As(err, &se) {
		return se.HasErrorCode(121)
	}
	return false
}
