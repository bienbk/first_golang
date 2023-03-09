package controllers

import (
	"encoding/json"
	"net/http"

	"bien.com/hocgo/model"
)

// type Hello struct {
// 	web.Controller
// }

// // @router / [get]
// func (c *Hello) Get() {
// 	c.Ctx.WriteString("tai sao")
// }

type TestController struct {
	baseController
}

// @router /:b1/:b2 [get]
// func (c *TestController) Xinchao() {
// 	result := c.Ctx.Input.Param(":b1")
// 	value, err := strconv.ParseInt(result, 10, 64)
// 	if err != nil {
// 		return
// 	}

// 	result1 := c.Ctx.Input.Param(":b2")
// 	value1, err1 := strconv.ParseInt(result1, 10, 64)
// 	if err1 != nil {
// 		return
// 	}
// 	sum := value + value1

// 	c.OK(http.StatusOK, model.CalResult{
// 		Re: int32(sum),
// 	})
// }

// @router /add [get]
func (c *TestController) Xinchao() {
	value, err := c.GetInt("p1")
	if err != nil {
		return
	}

	value1, err1 := c.GetInt("p2")
	if err1 != nil {
		return
	}
	sum := value + value1

	c.OK(http.StatusOK, model.AddResult{
		Calc: int32(sum),
	})
}

// @router /sub [get]
func (c *TestController) SubFunction() {
	value, err := c.GetInt("p1")
	if err != nil {
		return
	}

	value1, err1 := c.GetInt("p2")
	if err1 != nil {
		return
	}
	sub := value - value1
	c.OK(http.StatusOK, model.SubResult{
		Calc: int32(sub),
	})
}

// @router /div [post]
func (c *TestController) DivFunction() {
	body := model.PostDivParam{}
	err := json.NewDecoder(c.Ctx.Request.Body).Decode(&body)
	if err != nil {
		return
	}
	if body.Param2 == 0 {
		return
	}
	div := body.Param / body.Param2

	c.OK(http.StatusOK, model.DivResult{
		Calc: float64(div),
	})
}

// @router /multiplication [post]
func (c *TestController) Multipli() {
	body := model.PostMultParam{}
	err := json.NewDecoder(c.Ctx.Request.Body).Decode(&body)
	if err != nil {
		return
	}
	mult := body.Param * body.Param2
	c.OK(http.StatusOK, model.MultResult{
		Calc: mult,
	})
}
