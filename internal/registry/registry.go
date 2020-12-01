package registry

import (
	"context"
	"log"
	"os"

	"github.com/google/uuid"
	pb "github.com/selftechio/sandhog/internal/api/selftechio/naas"
	"google.golang.org/grpc"
)

// Register registers this sandhog.
func Register() error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("127.0.0.1:55555", opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewRegistryClient(conn)

	req := new(pb.RegisterRequest)
	req.Info = new(pb.SandhogInfo)
	// fixme 01/12/2020: uuid generation will create a panic if unsuccessful
	req.Info.Uuid = uuid.New().String()
	req.Info.Hostname, err = os.Hostname()
	if err != nil {
		return err
	}

	res, err := client.Register(context.Background(), req)
	if err != nil {
		return err
	}

	log.Println("register response:", res.Status.Result)

	return nil
}
