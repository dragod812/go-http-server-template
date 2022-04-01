package register

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dragod812/go-http-server-template/entity"
	"github.com/dragod812/go-http-server-template/lib/templates"
)

func BenchmarkReadPageHttpHandler(b *testing.B) {
	r, _ := http.NewRequest("GET", ReadPath+"paddy", nil)

	templates.NewTemplateCache("../" + templates.TemplateFile)

	for i := 0; i < b.N; i++ {
		ReadPageHttpHandler(httptest.NewRecorder(), r)
	}
}

func BenchmarkWritePageHttpHandler(b *testing.B) {
	data := entity.Page{
		PageName: "paddy",
		PageData: []byte("paddy is cool"),
	}

	dataJson, _ := json.Marshal(data)

	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("POST", WritePath+"paddy", strings.NewReader(string(dataJson)))
		req.Header.Add("Content-Type", "application/json; charset=UTF-8")
		WritePageHttpHandler(httptest.NewRecorder(), req)
	}
}

func BenchmarkWritePageHttpHandlerInputSize(b *testing.B) {
	sizes := []uint64{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824, 2147483648, 4294967296, 8589934592}
	pageDatas := make([]string, len(sizes))
	for idx, i := range sizes {
		data := entity.Page{
			PageName: "paddy",
			PageData: []byte(randSeq(i)),
		}

		dataJson, _ := json.Marshal(data)
		pageDatas[idx] = randSeq(i)
		b.Run(fmt.Sprintf("WritePageHandlerForLength-%d", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				req, _ := http.NewRequest("POST", WritePath+"paddy", strings.NewReader(string(dataJson)))
				req.Header.Add("Content-Type", "application/json; charset=UTF-8")
				WritePageHttpHandler(httptest.NewRecorder(), req)
			}
		})
	}

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n uint64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Test_getPageNameFromURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"response check",
			args{
				"/page/write/paddy",
			},
			"paddy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPageNameFromURL(tt.args.url); got != tt.want {
				t.Errorf("getPageNameFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
