package middlewares

import (
	"audit/src/di"
	"audit/src/sessions"
	"audit/src/utils"
	"audit/src/utils/res"
	"fmt"
	"net"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

type rateLimitData struct {
	Requests   int
	RetryAfter int
}

const rateLimitKeyPrefix = "RATE-LIMIT:"

// MdlwRateLimit sends 429 http error if req limit
func MdlwRateLimit(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cfg := di.GetAppConfig()
		storage := di.GetSessionStorage()

		ip := getIP(r)
		key := rateLimitKeyPrefix + ip
		data := getRateLimitData(storage, key)
		maxRequests := cfg.RateLimit.MaxRequests
		intervalSeconds := cfg.RateLimit.IntervalMs

		data.Requests++
		if data.Requests <= maxRequests {
			storage.SetJSON(key, data, intervalSeconds)
		} else {
			data.RetryAfter += intervalSeconds
			storage.SetJSON(key, data, data.RetryAfter)

			w.Header().Set("Retry-After", fmt.Sprint(data.RetryAfter))
			res.SendStatusError(w, http.StatusTooManyRequests, &utils.AppError{
				Status:  http.StatusTooManyRequests,
				Code:    "TOO_MANY_REQUESTS",
				Details: map[string]interface{}{"retryAfter": data.RetryAfter},
			})
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getRateLimitData(storage sessions.IStorage, key string) *rateLimitData {
	var res rateLimitData
	json, err := storage.GetJSON(key)
	if err != nil || json == nil {
		return &res
	}

	mapstructure.Decode(json, &res)
	return &res
}

func getIP(req *http.Request) string {
	forward := req.Header.Get("X-Forwarded-For")
	if forward == "" {
		forward = req.Header.Get("x-forwarded-for")
	}
	if forward == "" {
		forward = req.Header.Get("X-FORWARDED-FOR")
	}

	forwardIP := net.ParseIP(forward)
	if forwardIP != nil {
		return forwardIP.String()
	}

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return ""
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}

	return userIP.String()
}
