package main

//func Test_microservice(t *testing.T) {
//	req := httptest.NewRequest("GET", "/api/v1/ping", nil)
//	w := httptest.NewRecorder()
//
//	server := setUpEngine()
//	server.ServeHTTP(w, req)
//
//	if w.Code != http.StatusOK {
//		t.Fail()
//		return
//	}
//
//	resp := types.HealthCheckResp{}
//	_ = json.Unmarshal(w.Body.Bytes(), &resp)
//
//	logx.Infof("%+v ", resp)
//}
