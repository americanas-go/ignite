package server

import (
	"github.com/americanas-go/errors"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Error(err error) error {

	if errors.IsNotFound(err) {
		return status.Errorf(codes.NotFound, err.Error())
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		return status.Errorf(codes.InvalidArgument, err.Error())
	} else if errors.IsServiceUnavailable(err) {
		return status.Errorf(codes.Unavailable, err.Error())
	} else if errors.IsConflict(err) || errors.IsAlreadyExists(err) {
		return status.Errorf(codes.AlreadyExists, err.Error())
	} else if errors.IsNotImplemented(err) || errors.IsNotProvisioned(err) {
		return status.Errorf(codes.Unimplemented, err.Error())
	} else if errors.IsUnauthorized(err) {
		return status.Errorf(codes.Unauthenticated, err.Error())
	} else if errors.IsForbidden(err) {
		return status.Errorf(codes.PermissionDenied, err.Error())
	} else {
		switch t := err.(type) {
		case validator.ValidationErrors:
			return status.Errorf(codes.InvalidArgument, t.Error())
		default:
			return status.Errorf(codes.Internal, t.Error())
		}
	}
}
