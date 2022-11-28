package main

import (
	"fmt"
	"github.com/linqcod/student-testing-app/app/pkg/response"
)

func main() {
	resp := response.StatusOKWithDataModel{}
	resp.Code = 200
	resp.Message = "All is ok"
	resp.Data = "wadfg"

	fmt.Println(resp)
}
