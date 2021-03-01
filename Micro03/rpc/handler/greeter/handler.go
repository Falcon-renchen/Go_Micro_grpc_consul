package greeter

import (
	"Go_Micro/Micro03/proto/learning"
	"context"
	"fmt"
)

type Handler struct {

}

func (h *Handler) Hi(ctx context.Context, req *learning.Request, res *learning.Response) error {
	res.Msg = fmt.Sprintf("Hello! %s. Welcome to Greeter",req.Name)
	return nil
}

