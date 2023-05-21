package network

import "net/http"

///
type MethodHandler map[string]http.Handler

// ルーティング毎にMethodが正しいかを確認する
func (m MethodHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {

	if h, ok := m[r.Method]; ok {
		h.ServeHTTP(w, r)
		return
	}

	http.Error(w, "このメソッドは使うことができません", http.StatusMethodNotAllowed)
}
