package types

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusNoContent
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusRequestTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusForbidden
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}

	grpclog.Printf("Unknown gRPC error code: %v", code)
	return http.StatusInternalServerError
}

type errorBody struct {
	Error string `protobuf:"bytes,1,name=error" json:"error"`
	Code  int    `protobuf:"varint,2,name=code" json:"code"`
}

//Make this also conform to proto.Message for builtin JSONPb Marshaler
func (e *errorBody) Reset() { *e = errorBody{} }
func (e *errorBody) String() string {
	bytes, _ := json.Marshal(e)
	return string(bytes)
}
func (*errorBody) ProtoMessage() {}

// It simply writes a string representation of the given error into "w".
func DefaultErrorHandler(w http.ResponseWriter, status *status.Status) {
	httpCode := HTTPStatusFromCode(status.Code())
	eb := errorBody{Error: status.Message(), Code: httpCode}
	http.Error(w, eb.String(), httpCode)
}
