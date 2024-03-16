package routine

import (
	"fmt"
	"ms-lil-service/model"
	"sync"
)

type RoutineService struct {
	lil model.Structure /* :side_eye: */
	wg *sync.WaitGroup
}

func (rs *RoutineService) Init(wgg *sync.WaitGroup) {
   rs.lil = model.Lil{
		Name: "test",
		Nickname: "test",
		Age: 1,
   }
   rs.wg = wgg

   fmt.Println(rs.PrintSmth())
   fmt.Println(rs.Die())
}

func (rs RoutineService) PrintSmth() int {
	return rs.lil.AgePlus()
}

func (rs RoutineService) Die() int {
	rs.wg.Done()
	return 1
}