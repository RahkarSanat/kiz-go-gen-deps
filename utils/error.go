package utils

import (
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrMiddleware(err error, msgName string, uid string, logger *slog.Logger, otherMsgs map[string]string) error {
	msgs := []any{}
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
		err = status.Error(s.Code(), s.Message())
		return err
	default:
		logger.Debug(msgName, msgs...)
		err = status.Error(s.Code(), s.Code().String())
		return err
	}
}
