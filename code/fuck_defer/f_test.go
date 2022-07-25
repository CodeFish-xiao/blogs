package fuck_defer

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {

	func() {
		defer func() {
			if p := recover(); p != nil {
				log.Println(p)
			}
		}()
		func() {
			go func() {
				panic("asd")
			}()
			time.Sleep(1 * time.Second)
		}()
	}()
}
