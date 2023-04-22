package api

import (
	"github.com/aagolovanov/awesomeCodeService/util"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	conf = &util.Config{
		Port:   0,
		DBAddr: "",
		DBPass: "",
		TTL:    30,
	}
	dom  = GetMockDomain()
	logg = log.New(os.Stdout, "TESTSERVER ", log.Lmsgprefix)
	serv = NewServer(conf, dom, logg)
)

func TestBadMethod(t *testing.T) {
	t.Run("/send bad method", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/api/v1/send", nil)
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Fatalf("wanted %v, got %v", http.StatusMethodNotAllowed, recorder.Code)
		}
	})

	t.Run("/verify bad method", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/api/v1/verify", nil)
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusMethodNotAllowed {
			t.Fatalf("wanted %v, got %v", http.StatusMethodNotAllowed, recorder.Code)
		}
	})
}

func TestBadBody(t *testing.T) {
	t.Run("/send bad body", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"asd:123}`

		req, err := http.NewRequest("POST", "/api/v1/send", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("wanted %v, got %v", http.StatusBadRequest, recorder.Code)
		}
	})

	t.Run("/verify bad body", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"asd:123}`

		req, err := http.NewRequest("POST", "/api/v1/verify", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("wanted %v, got %v", http.StatusBadRequest, recorder.Code)
		}
	})
}

func TestInternalError(t *testing.T) {
	t.Run("/send internal", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"number": "internal"}`

		req, err := http.NewRequest("POST", "/api/v1/send", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusInternalServerError {
			t.Fatalf("wanted %v, got %v", http.StatusInternalServerError, recorder.Code)
		}
	})

	t.Run("/verify internal", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"requestId": "internal", "code": 1}`

		req, err := http.NewRequest("POST", "/api/v1/verify", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusInternalServerError {
			t.Fatalf("wanted %v, got %v", http.StatusInternalServerError, recorder.Code)
		}
	})
}

func TestCustomError(t *testing.T) {
	t.Run("/send custom error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"number": "custom"}`
		expected := `{"error":"customError"}`

		req, err := http.NewRequest("POST", "/api/v1/send", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("wanted %v, got %v", http.StatusBadRequest, recorder.Code)
		}

		responseBody, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("Error while reading body: %v", err)
		}

		responseString := string(responseBody)
		if responseString != expected {
			t.Fatalf("wanted %v, got %v", expected, responseString)
		}
	})

	t.Run("/verify custom error", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"requestId": "custom", "code": 1}`
		expected := `{"error":"customError"}`

		req, err := http.NewRequest("POST", "/api/v1/verify", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusBadRequest {
			t.Fatalf("wanted %v, got %v", http.StatusBadRequest, recorder.Code)
		}

		responseBody, err := io.ReadAll(recorder.Body)
		if err != nil {
			t.Fatalf("Error while reading body: %v", err)
		}

		responseString := string(responseBody)
		if responseString != expected {
			t.Fatalf("wanted %v, got %v", expected, responseString)
		}
	})
}

func TestSuccess(t *testing.T) {
	// todo test response body
	t.Run("/send OK", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"number": "123"}`

		req, err := http.NewRequest("POST", "/api/v1/send", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Fatalf("wanted %v, got %v", http.StatusOK, recorder.Code)
		}
	})

	t.Run("/verify OK", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		body := `{"requestId": "123", "code": 1}`

		req, err := http.NewRequest("POST", "/api/v1/verify", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Error while creating request: %v", err)
		}

		serv.router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Fatalf("wanted %v, got %v", http.StatusOK, recorder.Code)
		}
	})
}
