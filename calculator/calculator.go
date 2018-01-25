package calculator

import (
	"log"
)

type Service int

func (c *Service) Sum(rq *Request, rp *Response) error {

	log.Printf("Exectuting sum with data: %+v", rq)
	rp.Result = rq.A + rq.B
	return nil

}
