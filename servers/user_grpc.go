package servers

import (
	"context"
	"fmt"
	"time"

	"github.com/wenealves10/golang-grpc/pb"
)

type User struct {
}

func (u *User) FnUser(ctx context.Context, in *pb.UserResquest) (*pb.UserResponse, error) {
	fmt.Println("userName:", in.User.GetName())
	fmt.Println("userEmail:", in.User.GetEmail())
	fmt.Println("userBirthDay:", in.User.GetBirthDate())

	return &pb.UserResponse{
		Age: uint64(time.Now().Year()) - in.User.BirthDate,
	}, nil
}
