// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/ports/input.go
//
// Generated by this command:
//
//	Cursor-0.47.9-x86_64.AppImage -source=internal/domain/ports/input.go -destination=internal/mocks/domain/ports/mock_input.go
//

// Package mock_ports is a generated GoMock package.
package mock_ports

import (
	context "context"
	entities "real-time-messaging/consumer/internal/domain/entities"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	websocket "github.com/gorilla/websocket"
	gomock "go.uber.org/mock/gomock"
)

// MockWebsocket is a mock of Websocket interface.
type MockWebsocket struct {
	ctrl     *gomock.Controller
	recorder *MockWebsocketMockRecorder
	isgomock struct{}
}

// MockWebsocketMockRecorder is the mock recorder for MockWebsocket.
type MockWebsocketMockRecorder struct {
	mock *MockWebsocket
}

// NewMockWebsocket creates a new mock instance.
func NewMockWebsocket(ctrl *gomock.Controller) *MockWebsocket {
	mock := &MockWebsocket{ctrl: ctrl}
	mock.recorder = &MockWebsocketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebsocket) EXPECT() *MockWebsocketMockRecorder {
	return m.recorder
}

// Receive mocks base method.
func (m *MockWebsocket) Receive(conn *websocket.Conn) (*entities.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive", conn)
	ret0, _ := ret[0].(*entities.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockWebsocketMockRecorder) Receive(conn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockWebsocket)(nil).Receive), conn)
}

// Upgrade mocks base method.
func (m *MockWebsocket) Upgrade(c *gin.Context) (*websocket.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", c)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockWebsocketMockRecorder) Upgrade(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockWebsocket)(nil).Upgrade), c)
}

// MockConsumer is a mock of Consumer interface.
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
	isgomock struct{}
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer.
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance.
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockConsumer) Consume(c *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Consume indicates an expected call of Consume.
func (mr *MockConsumerMockRecorder) Consume(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockConsumer)(nil).Consume), c)
}

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// GetUserByCredentials mocks base method.
func (m *MockUserService) GetUserByCredentials(ctx context.Context, email, password string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByCredentials", ctx, email, password)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByCredentials indicates an expected call of GetUserByCredentials.
func (mr *MockUserServiceMockRecorder) GetUserByCredentials(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByCredentials", reflect.TypeOf((*MockUserService)(nil).GetUserByCredentials), ctx, email, password)
}

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
	isgomock struct{}
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuthService) GenerateToken(ctx context.Context, user *entities.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthServiceMockRecorder) GenerateToken(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthService)(nil).GenerateToken), ctx, user)
}

// ValidateToken mocks base method.
func (m *MockAuthService) ValidateToken(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthServiceMockRecorder) ValidateToken(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthService)(nil).ValidateToken), ctx, token)
}

// Close mocks base method.
func (m *MockAuthService) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAuthServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAuthService)(nil).Close))
}
