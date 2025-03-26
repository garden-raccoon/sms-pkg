package orders

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/garden-raccoon/sms-pkg/models"
	proto "github.com/garden-raccoon/sms-pkg/protocols/sms-pkg"
)

// TODO need to set timeout via lib initialisation
// timeOut is  hardcoded GRPC requests timeout value
const timeOut = 60

// Debug on/off
var Debug = true

type SmsPkgAPI interface {
	CreateOrUpdateSms(s *models.Sms) error
	DeleteSms(userUuid uuid.UUID) error
	SmsBy(userUuid uuid.UUID) (*models.Sms, error)
	HealthCheck() error
	// Close GRPC Api connection
	Close() error
}

// Api is profile-service GRPC Api
// structure with client Connection
type Api struct {
	addr    string
	timeout time.Duration
	*grpc.ClientConn
	mu sync.Mutex
	proto.SmsServiceClient
	grpc_health_v1.HealthClient
}

// New create new Battles Api instance
func New(addr string) (SmsPkgAPI, error) {
	api := &Api{timeout: timeOut * time.Second}

	if err := api.initConn(addr); err != nil {
		return nil, fmt.Errorf("create MealsApi:  %w", err)
	}
	api.HealthClient = grpc_health_v1.NewHealthClient(api.ClientConn)

	api.SmsServiceClient = proto.NewSmsServiceClient(api.ClientConn)
	return api, nil
}

func (api *Api) DeleteSms(userUuid uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	req := &proto.SmsByRequest{
		UserUuid: userUuid.Bytes(),
	}
	_, err := api.SmsServiceClient.DeleteSms(ctx, req)
	if err != nil {
		return fmt.Errorf("DeleteSms api request: %w", err)
	}
	return nil
}

func (api *Api) CreateOrUpdateSms(s *models.Sms) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	sms := models.ToProto(*s)
	_, err = api.SmsServiceClient.CreateOrUpdateSms(ctx, sms)
	if err != nil {
		return fmt.Errorf("create sms api request: %w", err)
	}
	return nil
}

// initConn initialize connection to Grpc servers
func (api *Api) initConn(addr string) (err error) {
	var kacp = keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}

	api.ClientConn, err = grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}
	return
}

func (api *Api) SmsBy(userUuid uuid.UUID) (*models.Sms, error) {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()
	req := &proto.SmsByRequest{
		UserUuid: userUuid.Bytes(),
	}
	resp, err := api.SmsServiceClient.SmsBy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("MealAPI getMeal request failed: %w", err)
	}

	sms := models.FromProto(resp)
	return sms, nil
}

func (api *Api) HealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), api.timeout)
	defer cancel()

	api.mu.Lock()
	defer api.mu.Unlock()

	resp, err := api.HealthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{Service: "smsapi"})
	if err != nil {
		return fmt.Errorf("healthcheck error: %w", err)
	}

	if resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
		return fmt.Errorf("node is %s", errors.New("service is unhealthy"))
	}
	//api.status = service.StatusHealthy
	return nil
}
