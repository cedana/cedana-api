package task

// cts encapsulates client functions to interact with the services

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	DEFAULT_PROCESS_DEADLINE    = 20 * time.Minute
	DEFAULT_CONTAINERD_DEADLINE = 10 * time.Minute
	DEFAULT_RUNC_DEADLINE       = 10 * time.Minute
)

type ServiceClient struct {
	TaskService TaskServiceClient
	TaskConn    *grpc.ClientConn
}

func NewClient(address string) (*ServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	taskConn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	taskClient := NewTaskServiceClient(taskConn)

	client := &ServiceClient{
		TaskService: taskClient,
		TaskConn:    taskConn,
	}
	return client, err
}

func (c *ServiceClient) Close() {
	c.TaskConn.Close()
}

/////////////////////////////
//      Health Check       //
/////////////////////////////

func (c *ServiceClient) HealthCheck(ctx context.Context) (bool, error) {
	healthClient := grpc_health_v1.NewHealthClient(c.TaskConn)
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()

	opts := getDefaultCallOptions()

	// Health check
	resp, err := healthClient.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: "TaskService",
	}, opts...)
	if err != nil {
		return false, err
	}

	if resp.Status == grpc_health_v1.HealthCheckResponse_SERVING {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *ServiceClient) DetailedHealthCheck(ctx context.Context, args *DetailedHealthCheckRequest) (*DetailedHealthCheckResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.DetailedHealthCheck(ctx, args, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

///////////////////////////
// Process Service Calls //
///////////////////////////

func (c *ServiceClient) Start(ctx context.Context, args *StartArgs) (*StartResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.Start(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) StartAttach(ctx context.Context, args *StartAttachArgs) (TaskService_StartAttachClient, error) {
	opts := getDefaultCallOptions()
	stream, err := c.TaskService.StartAttach(ctx, opts...)
	if err != nil {
		return nil, err
	}
	// Send the first start request
	if err := stream.Send(args); err != nil {
		return nil, err
	}
	return stream, nil
}

func (c *ServiceClient) Dump(ctx context.Context, args *DumpArgs) (*DumpResp, error) {
	// TODO NR - timeouts here need to be fixed
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.Dump(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) Restore(ctx context.Context, args *RestoreArgs) (*RestoreResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.Restore(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) RestoreAttach(ctx context.Context, args *RestoreAttachArgs) (TaskService_RestoreAttachClient, error) {
	opts := getDefaultCallOptions()
	stream, err := c.TaskService.RestoreAttach(ctx, opts...)
	if err != nil {
		return nil, err
	}
	// Send the first restore request
	if err := stream.Send(args); err != nil {
		return nil, err
	}
	return stream, nil
}

func (c *ServiceClient) Query(ctx context.Context, args *QueryArgs) (*QueryResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.Query(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//////////////////////////////
////// CRIO Rootfs Dump //////
//////////////////////////////

func (c *ServiceClient) CRIORootfsDump(ctx context.Context, args *CRIORootfsDumpArgs) (*CRIORootfsDumpResp, error) {
	// TODO NR - timeouts here need to be fixed
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.CRIORootfsDump(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) CRIOImagePush(ctx context.Context, args *CRIOImagePushArgs) (*CRIOImagePushResp, error) {
	// TODO NR - timeouts here need to be fixed
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.CRIOImagePush(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//////////////////////////////
// Containerd Service Calls //
//////////////////////////////

func (c *ServiceClient) ContainerdDump(ctx context.Context, args *ContainerdDumpArgs) (*ContainerdDumpResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_CONTAINERD_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.ContainerdDump(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) ContainerdRestore(ctx context.Context, args *ContainerdRestoreArgs) (*ContainerdRestoreResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_CONTAINERD_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.ContainerdRestore(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) ContainerdQuery(ctx context.Context, args *ContainerdQueryArgs) (*ContainerdQueryResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_CONTAINERD_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.ContainerdQuery(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) ContainerdRootfsDump(ctx context.Context, args *ContainerdRootfsDumpArgs) (*ContainerdRootfsDumpResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_CONTAINERD_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.ContainerdRootfsDump(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) ContainerdRootfsRestore(ctx context.Context, args *ContainerdRootfsRestoreArgs) (*ContainerdRootfsRestoreResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_CONTAINERD_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.ContainerdRootfsRestore(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

////////////////////////
// Runc Service Calls //
////////////////////////

func (c *ServiceClient) RuncGetPausePid(ctx context.Context, args *RuncGetPausePidArgs) (*RuncGetPausePidResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.RuncGetPausePid(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) RuncDump(ctx context.Context, args *RuncDumpArgs) (*RuncDumpResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_RUNC_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.RuncDump(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) RuncRestore(ctx context.Context, args *RuncRestoreArgs) (*RuncRestoreResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_RUNC_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.RuncRestore(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) RuncQuery(ctx context.Context, args *RuncQueryArgs) (*RuncQueryResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_RUNC_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.RuncQuery(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

///////////////////////////
// Kata Service Calls //
///////////////////////////

func (c *ServiceClient) KataDump(ctx context.Context, args *DumpArgs) (*DumpResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	resp, err := c.TaskService.KataDump(ctx, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ServiceClient) KataRestore(ctx context.Context, args *RestoreArgs) (*RestoreResp, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	resp, err := c.TaskService.KataRestore(ctx, args)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

////////////////////////////
/// Config Service Calls ///
////////////////////////////

func (c *ServiceClient) GetConfig(ctx context.Context, args *GetConfigRequest) (*GetConfigResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, DEFAULT_PROCESS_DEADLINE)
	defer cancel()
	opts := getDefaultCallOptions()
	resp, err := c.TaskService.GetConfig(ctx, args, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

///////////////////
//    Helpers    //
///////////////////

func getDefaultCallOptions() []grpc.CallOption {
	opts := []grpc.CallOption{}
	opts = append(opts, grpc.WaitForReady(true))
	return opts
}
