package goz

import (
	"fmt"
	"log"

	"github.com/idoubi/goz"
)

func ExampleRequest_Get() {
	cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Get_withQuery_arr() {
	cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get-with-query", goz.Options{
		Query: map[string]interface{}{
			"key1": "value1",
			"key2": []string{"value21", "value22"},
			"key3": "333",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
	// Output: key1=value1&key2=value21&key2=value22&key3=333
}

func ExampleRequest_Get_withQuery_str() {
	cli := goz.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get-with-query?key0=value0", goz.Options{
		Query: "key1=value1&key2=value21&key2=value22&key3=333",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
	// Output: key1=value1&key2=value21&key2=value22&key3=333
}

func ExampleRequest_Get_withProxy() {
	cli := goz.NewClient()

	resp, err := cli.Get("https://www.fbisb.com/ip.php", goz.Options{
		Timeout: 5.0,
		Proxy:   "http://127.0.0.1:1087",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetStatusCode())
	// Output: 200
}

func ExampleRequest_Post() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goz.Response
}

func ExampleRequest_Post_withHeaders() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-headers", goz.Options{
		Headers: map[string]interface{}{
			"User-Agent": "testing/1.0",
			"Accept":     "application/json",
			"X-Foo":      []string{"Bar", "Baz"},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	headers := resp.GetRequest().Header["X-Foo"]
	fmt.Println(headers)
	// Output: [Bar Baz]
}

func ExampleRequest_Post_withFormParams() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-form-params", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: map[string]interface{}{
			"key1": "value1",
			"key2": []string{"value21", "value22"},
			"key3": "333",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := resp.GetBody()
	fmt.Println(body)
	// Output: form params:{"key1":["value1"],"key2":["value21","value22"],"key3":["333"]}
}

func ExampleRequest_Post_withJSON() {
	cli := goz.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-json", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		JSON: struct {
			Key1 string   `json:"key1"`
			Key2 []string `json:"key2"`
			Key3 int      `json:"key3"`
		}{"value1", []string{"value21", "value22"}, 333},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body, _ := resp.GetBody()
	fmt.Println(body)
	// Output: json:{"key1":"value1","key2":["value21","value22"],"key3":333}
}
