package learning



import "Go_Micro/Micro03/proto/learning"
import "context"
import "fmt"

type Handler struct {

}

func (h *Handler) Hi(ctx context.Context, req *learning.Request, res *learning.Response) error {
	res.Msg = fmt.Sprintf("Hello! %s. Welcome to Learning",req.Name)
	return nil
}
