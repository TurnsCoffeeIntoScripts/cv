package concourse

import (
	"cv/pkg/models"
	"cv/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"strings"
)

func ApiCall(m string, body []byte) error {
	client := http.DefaultClient

	var baseUrl string
	if viper.Get("concourse.url") != nil {
		baseUrl = viper.Get("concourse.url").(string)
	}
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}

	baseUrl += "api/v1/"

	url := baseUrl + "teams"

	var req *http.Request
	var newRequestError error
	if body == nil {
		req, newRequestError = http.NewRequest(m, url, nil)
	}

	if newRequestError != nil {
		// TODO
	}

	if req != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "d5f7nhI9zZe5hDXHurSsqthnl2oFrTVkAAAAAA")
		//req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImJhN2NlMDU0ODdiNjFlMTZiMzcwZjNhZGNhMzM5YzBhM2QwNjhlMjIifQ.eyJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvc2t5L2lzc3VlciIsInN1YiI6IkNnUjBaWE4wRWdWc2IyTmhiQSIsImF1ZCI6ImZseSIsImV4cCI6MTY4MTIzOTIyOSwiaWF0IjoxNjgxMTUyODI5LCJhdF9oYXNoIjoiVkhXaHpReUFEaXlNRzZjZkw0YXdhUSIsImVtYWlsIjoidGVzdCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYW1lIjoidGVzdCIsImZlZGVyYXRlZF9jbGFpbXMiOnsiY29ubmVjdG9yX2lkIjoibG9jYWwiLCJ1c2VyX2lkIjoidGVzdCJ9fQ.Ahsw0hfCfDh2cq6TphO3kC8EKSsRFoXTLYH2os4HhVxYN3ay1H-vinaOlcJQVsddsRVCBsC-ofnhwqStgsu9sVkMXUB1efIKqqYLrwCzKr8aXeWXxhN2K5KV33XaDoFFAes9ln6xJzu_rB8NZuBFQ4WOmHt8uxu3T4_yDUycB0GQZOrh2ojpkDKhYuONxpGrIRbH9dwZqKz-jZS-DwUCeVk-gZkCAG3yWgWuavbOmR2UMXsowndNw0RGlnbHLb2DzvDFxkCn9S08K4PZmP4pb2vBgAEtHUa-zIP5HHgDqEF0jPq0OlT05ZeZE9K5WNnlCEXEZi85ZHUyVpZifOJexA")
	}

	fmt.Println(req.URL.String())
	resp, _ := client.Do(req)

	defer utils.CloseHttpReadCloser(resp)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var t []models.Team
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(len(t))

	return nil
}
