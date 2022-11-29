package main
import (
    "log"
    "time"
    "github.com/spy16/clipboard"
	"fmt"
	"os"
)

func main() {
	text, _ := clipboard.ReadAll()
	fmt.Println(text)
    changes := make(chan string, 10)
    stopCh := make(chan struct{})

    go clipboard.Monitor(time.Second, stopCh, changes)
	
    for {
        select {
        case <-stopCh:
            break
        default:
            change, ok := <-changes
            if ok {
                log.Printf("change received: '%s'", change)
				f, err := os.OpenFile("clipboard.txt", os.O_APPEND|os.O_WRONLY, 0644)
				f.WriteString(change+"\n");
			if err != nil {
    			panic(err)
			}
            } else {
                log.Printf("channel has ben closede. exiting..")
            }
        }
    }
}