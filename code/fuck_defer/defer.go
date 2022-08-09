package fuck_defer

import "fmt"

type Entity struct {
	fieldOne string
}

//定义的Option方法
type Option func(e *Entity)

func FieldOne(fieldOne string) Option {
	return func(e *Entity) {
		e.fieldOne = fieldOne
	}
}

func (e *Entity) WithOptions(options ...Option) *Entity {
	for _, option := range options {
		option(e)
	}
	return e
}
func (e *Entity) Do() {
	fmt.Println(e.fieldOne)
}

//e := &Entity{}
//e.WithOptions(
//    FieldOne("field one's value"),
//)
