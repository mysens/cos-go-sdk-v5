package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"net/http"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func log_status(err error) {
	if err == nil {
		return
	}
	if cos.IsNotFoundError(err) {
		// WARN
		fmt.Println("WARN: Resource is not existed")
	} else if e, ok := cos.IsCOSError(err); ok {
		fmt.Printf("ERROR: Code: %v\n", e.Code)
		fmt.Printf("ERROR: Message: %v\n", e.Message)
		fmt.Printf("ERROR: Resource: %v\n", e.Resource)
		fmt.Printf("ERROR: RequestId: %v\n", e.RequestID)
		// ERROR
	} else {
		fmt.Printf("ERROR: %v\n", err)
		// ERROR
	}
}

func main() {
	u, _ := url.Parse("http://test-1259654469.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
			Listener:    &cos.DefaultProgressListener{},
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			XCosACL: "private",
		},
	}

	name := "append" + time.Now().Format(time.RFC3339)
	str1 := "test append object 1"
	pos, _, err := c.Object.Append(context.Background(), name, 0, strings.NewReader(str1), opt)
	log_status(err)
	fmt.Printf("pos: %d\n", pos)

	str2 := "test append object 2"
	pos, _, err = c.Object.Append(context.Background(), name, pos, strings.NewReader(str2), opt)
	log_status(err)
	fmt.Printf("pos: %d\n", pos)

}