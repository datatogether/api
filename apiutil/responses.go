package apiutil

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WriteResponse(w http.ResponseWriter, data interface{}) error {
	env := map[string]interface{}{
		"meta": map[string]interface{}{
			"code": http.StatusOK,
		},
		"data": data,
	}
	return jsonResponse(w, env)
}

func WritePageResponse(w http.ResponseWriter, data interface{}, r *http.Request, p Page) error {
	env := map[string]interface{}{
		"meta": map[string]interface{}{
			"code": http.StatusOK,
		},
		"data": data,
		"pagination": map[string]interface{}{
			"nextUrl": nextPageUrl(r, p),
		},
	}
	return jsonResponse(w, env)
}

func nextPageUrl(r *http.Request, p Page) string {
	q := r.URL.Query()
	pageNum := 1
	if val, ok := q["page"]; ok {
		var err error
		if pageNum, err = strconv.Atoi(val[0]); err != nil {
			return r.URL.String()
		}
	}
	q.Set("page", strconv.Itoa(pageNum+1))
	r.URL.RawQuery = q.Encode()
	return r.URL.String()
}

func WriteMessageResponse(w http.ResponseWriter, message string, data interface{}) error {
	env := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":    http.StatusOK,
			"message": message,
		},
		"data": data,
	}

	return jsonResponse(w, env)
}

func WriteErrResponse(w http.ResponseWriter, code int, err error) error {
	env := map[string]interface{}{
		"meta": map[string]interface{}{
			"code":  code,
			"error": err.Error(),
		},
	}

	res, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(code)
	_, err = w.Write(res)
	return err
}

func jsonResponse(w http.ResponseWriter, env interface{}) error {
	res, err := json.Marshal(env)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	return err
}
