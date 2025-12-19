package middleware

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex" // or base64
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const SHARED_SECRET_KEY = "your-very-secret-and-long-key" // !!! in env
const requestValidityDuration = 200 * time.Second

// AuthGuardMiddlewareWithHMAC by HMAC
func AuthGuardMiddlewareWithHMAC() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientSign := r.Header.Get("X-Sign")
			requestTimeStr := r.Header.Get("X-Request-Time")

			if clientSign == "" || requestTimeStr == "" {
				log.Println("HMAC Auth: Missing X-Sign or X-Request-Time header")
				writeErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "Missing required signature headers")
				return
			}

			// 1. Kiểm tra Timestamp
			requestTime, err := strconv.ParseInt(requestTimeStr, 10, 64)
			if err != nil {
				log.Println("HMAC Auth: Invalid request time format")
				writeErrorResponse(w, http.StatusBadRequest, "Invalid request", "Invalid request time format")
				return
			}

			now := time.Now().Unix()
			if now-requestTime > int64(requestValidityDuration.Seconds()) || now-requestTime < -5 { // Chấp nhận trễ 5s
				log.Printf("HMAC Auth: Request timestamp out of bounds. Now: %d, ReqTime: %d\n", now, requestTime)
				writeErrorResponse(w, http.StatusBadRequest, "Invalid request", "Request timestamp out of bounds")
				return
			}

			// 2. Tái tạo String-to-Sign
			// Đọc body (nếu có) và chuẩn bị để đọc lại
			var bodyBytes []byte
			if r.Body != nil {
				bodyBytes, err = io.ReadAll(r.Body)
				if err != nil {
					log.Println("HMAC Auth: Error reading request body:", err)
					writeErrorResponse(w, http.StatusInternalServerError, "Server Error", "Could not read request body")
					return
				}
				// Tạo lại body để handler sau có thể đọc
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}

			stringToSign := buildStringToSign(r, requestTimeStr, bodyBytes)
			log.Printf("HMAC Auth: Server StringToSign: [%s]\n", strings.ReplaceAll(stringToSign, "\n", "\\n"))

			// 3. Tính toán HMAC phía server
			serverSign := calculateHMAC(stringToSign, SHARED_SECRET_KEY)
			log.Printf("HMAC Auth: ClientSign: %s, ServerSign: %s\n", clientSign, serverSign)

			// 4. So sánh chữ ký
			// Sử dụng hmac.Equal để so sánh an toàn, chống timing attacks
			if !hmac.Equal([]byte(clientSign), []byte(serverSign)) {
				log.Println("HMAC Auth: Invalid signature")
				writeErrorResponse(w, http.StatusUnauthorized, "Unauthorized", "Invalid signature")
				return
			}

			log.Println("HMAC Auth: Signature verified successfully")
			next.ServeHTTP(w, r)
		})
	}
}

func buildStringToSign(r *http.Request, requestTimeStr string, bodyBytes []byte) string {
	method := r.Method
	path := r.URL.Path // Chỉ path, không có query string

	// Sắp xếp query parameters
	queryParams := r.URL.Query()
	var sortedQueryKeys []string
	for k := range queryParams {
		sortedQueryKeys = append(sortedQueryKeys, k)
	}
	sort.Strings(sortedQueryKeys)

	var canonicalQueryParts []string
	for _, k := range sortedQueryKeys {
		if len(queryParams[k]) > 0 {
			canonicalQueryParts = append(canonicalQueryParts, fmt.Sprintf("%s=%s", k, url.QueryEscape(queryParams[k][0])))
		}
	}
	canonicalQueryString := strings.Join(canonicalQueryParts, "&")

	// Xử lý body:
	//  - Nếu body rỗng, có thể dùng chuỗi rỗng hoặc một hash cố định của chuỗi rỗng.
	//  - Nếu body không rỗng, nên hash body (ví dụ SHA256) rồi đưa hash đó vào stringToSign.
	//    Điều này an toàn hơn là đưa raw body (có thể rất lớn) vào stringToSign.
	//    Ở đây, để đơn giản, tôi sẽ đưa raw bodyBytes (đã được convert sang string) vào,
	//    nhưng HASHING BODY LÀ BEST PRACTICE.
	bodyString := ""
	if len(bodyBytes) > 0 {
		bodyString = string(bodyBytes) // CẢNH BÁO: Nếu body không phải UTF-8, sẽ có vấn đề. Nên hash body!
		// Ví dụ hash body:
		// bodyHasher := sha256.New()
		// bodyHasher.Write(bodyBytes)
		// bodyString = hex.EncodeToString(bodyHasher.Sum(nil))
	}

	// Thứ tự phải nhất quán giữa client và server
	// Ví dụ: METHOD\nPATH\nTIMESTAMP\nSORTED_QUERY_STRING\nBODY_STRING_OR_HASH
	parts := []string{
		method,
		path,
		requestTimeStr,
		canonicalQueryString,
		bodyString, // Hoặc hash của body
	}
	return strings.Join(parts, "\n")
}

func calculateHMAC(data string, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil)) // Hoặc base64.StdEncoding.EncodeToString
}

// Context helpers for storing request-scoped data
type contextKeyType string

const bodyBytesKey contextKeyType = "bodyBytes"

// WithBodyBytes stores the request body bytes in context
func WithBodyBytes(ctx context.Context, bodyBytes []byte) context.Context {
	return context.WithValue(ctx, bodyBytesKey, bodyBytes)
}

// GetBodyBytes retrieves the request body bytes from context
func GetBodyBytes(ctx context.Context) []byte {
	if val := ctx.Value(bodyBytesKey); val != nil {
		if bytes, ok := val.([]byte); ok {
			return bytes
		}
	}
	return nil
}

// writeErrorResponse writes a JSON error response
func writeErrorResponse(w http.ResponseWriter, code int, message string, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    code,
		"message": message,
		"error":   err,
	})
}
