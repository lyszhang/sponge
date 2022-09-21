package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/zhufuyi/sponge/api/types"
	pb "github.com/zhufuyi/sponge/api/userExample/v1"
	"github.com/zhufuyi/sponge/config"
	"github.com/zhufuyi/sponge/pkg/grpc/benchmark"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initUserExampleServiceClient() pb.UserExampleServiceClient {
	err := config.Init(config.Path("conf.yml"))
	if err != nil {
		panic(err)
	}
	addr := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)

	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	return pb.NewUserExampleServiceClient(conn)
}

// 通过客户端测试userExample的各个方法
func Test_userExampleService_methods(t *testing.T) {
	cli := initUserExampleServiceClient()
	ctx := context.Background()

	tests := []struct {
		name    string
		fn      func() (interface{}, error)
		wantErr bool
	}{
		// todo generate the service struct code here
		// delete the templates code start
		{
			name: "Create",
			fn: func() (interface{}, error) {
				// todo test after filling in parameters
				return cli.Create(ctx, &pb.CreateUserExampleRequest{
					Name:     "宋九",
					Email:    "foo7@bar.com",
					Password: "f447b20a7fcbf53a5d5be013ea0b15af",
					Phone:    "+8618576552066",
					Avatar:   "http://internal.com/7.jpg",
					Age:      21,
					Gender:   2,
				})
			},
			wantErr: false,
		},

		{
			name: "UpdateByID",
			fn: func() (interface{}, error) {
				// todo test after filling in parameters
				return cli.UpdateByID(ctx, &pb.UpdateUserExampleByIDRequest{
					Id:    7,
					Phone: "18666666666",
					Age:   21,
				})
			},
			wantErr: false,
		},
		// delete the templates code end
		{
			name: "DeleteByID",
			fn: func() (interface{}, error) {
				// todo test after filling in parameters
				return cli.DeleteByID(ctx, &pb.DeleteUserExampleByIDRequest{
					Id: 3,
				})
			},
			wantErr: false,
		},

		{
			name: "GetByID",
			fn: func() (interface{}, error) {
				// todo test after filling in parameters
				return cli.GetByID(ctx, &pb.GetUserExampleByIDRequest{
					Id: 3,
				})
			},
			wantErr: false,
		},

		{
			name: "List",
			fn: func() (interface{}, error) {
				// todo test after filling in parameters
				return cli.List(ctx, &pb.ListUserExampleRequest{
					Params: &types.Params{
						Page:  0,
						Limit: 10,
						Sort:  "",
						Columns: []*types.Column{
							{
								Name:  "id",
								Exp:   "<",
								Value: "100",
								Logic: "",
							},
						},
					},
				})
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			t.Log("reply data: ", got)
		})
	}
}

// 压测userExample的各个方法，完成后复制报告路径到浏览器查看
func Test_userExampleService_benchmark(t *testing.T) {
	err := config.Init(config.Path("conf.yml"))
	if err != nil {
		panic(err)
	}
	host := fmt.Sprintf("127.0.0.1:%d", config.Get().Grpc.Port)
	protoFile := config.Path("../api/userExample/v1/userExample.proto")
	// 如果压测过程中缺少第三方依赖，复制到项目的third_party目录下(不包括import路径)
	importPaths := []string{
		config.Path("../third_party"), // third_party目录
		config.Path(".."),             // third_party的上一级目录
	}

	tests := []struct {
		name    string
		fn      func() error
		wantErr bool
	}{
		{
			name: "GetByID",
			fn: func() error {
				// todo test after filling in parameters
				message := &pb.GetUserExampleByIDRequest{
					Id: 3,
				}
				b, err := benchmark.New(host, protoFile, "GetByID", message, 100, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},

		{
			name: "List",
			fn: func() error {
				// todo test after filling in parameters
				message := &pb.ListUserExampleRequest{
					Params: &types.Params{
						Page:  0,
						Limit: 10,
						Sort:  "",
						Columns: []*types.Column{
							{
								Name:  "id",
								Exp:   "<",
								Value: "100",
								Logic: "",
							},
						},
					},
				}
				b, err := benchmark.New(host, protoFile, "List", message, 100, importPaths...)
				if err != nil {
					return err
				}
				return b.Run()
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			if (err != nil) != tt.wantErr {
				t.Errorf("test '%s' error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
		})
	}
}
