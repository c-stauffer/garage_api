package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const microAddress string = "http://192.168.1.45:6262/"
const defaultRequestTimeoutSeconds int = 5

func Ping(w http.ResponseWriter, r *http.Request) {
	sendRequestAndRespond(w, "ping")
}

func Actuate(w http.ResponseWriter, r *http.Request) {
	sendRequestAndRespond(w, "actuate")
}

func GetState(w http.ResponseWriter, r *http.Request) {
	sendRequestAndRespond(w, "state")
}

func GetStateValues(w http.ResponseWriter, r *http.Request) {
	// example response from this function (assuming no error):
	// {"closed":"LOW","open":"LOW","last":"UNKNOWN"}
	resp, statusCode, err := sendRequest("state2")
	var bytes []byte
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
	} else {
		bytes = []byte(resp + "\n")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

func GetAllValues(w http.ResponseWriter, r *http.Request) {
	// example response from this function (assuming no error):
	// {"closed":"LOW","heatindex":"48.56","humidity":"48.50","last":"UNKNOWN","open":"LOW","state":"UNKNOWN","temperature":"10.80"}
	var bytes []byte
	w.Header().Add("Content-Type", "application/json")
	resp, statusCode, err := sendRequest("state")
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
		w.WriteHeader(statusCode)
		w.Write(bytes)
		return
	}
	mAll := map[string]interface{}{
		"state": resp,
	}

	resp, statusCode, err = sendRequest("temperature")
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
		w.WriteHeader(statusCode)
		w.Write(bytes)
		return
	}
	mAll["temperature"] = resp

	resp, statusCode, err = sendRequest("humidity")
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
		w.WriteHeader(statusCode)
		w.Write(bytes)
		return
	}
	mAll["humidity"] = resp

	resp, statusCode, err = sendRequest("heatindex")
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
		w.WriteHeader(statusCode)
		w.Write(bytes)
		return
	}
	mAll["heatindex"] = resp

	resp, statusCode, err = sendRequest("state2")
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
		w.WriteHeader(statusCode)
		w.Write(bytes)
		return
	} else {
		bytes = []byte(resp)
	}
	var mTmp map[string]interface{}
	json.Unmarshal(bytes, &mTmp)
	for k, v := range mTmp {
		mAll[k] = v
	}

	bytes, _ = json.Marshal(mAll)
	w.Write(bytes)
}

func GetTemperature(w http.ResponseWriter, r *http.Request) {
	sendRequestAndRespond(w, "temperature")
}

func GetHumidity(w http.ResponseWriter, r *http.Request) {
	sendRequestAndRespond(w, "humidity")
}

func GetHeatIndex(w http.ResponseWriter, r *http.Request) {
	sendRequestAndRespond(w, "heatindex")
}

func sendRequestAndRespond(w http.ResponseWriter, uri string) {
	microResponse, statusCode, err := sendGetRequestToMicrocontroller(uri, defaultRequestTimeoutSeconds)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	var bytes []byte
	if err != nil {
		bytes, _ = json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
	} else {
		bytes, _ = json.Marshal(map[string]interface{}{
			uri: microResponse,
		})
	}
	w.Write(bytes)
}

func sendRequest(uri string) (string, int, error) {
	microResponse, statusCode, err := sendGetRequestToMicrocontroller(uri, defaultRequestTimeoutSeconds)
	return microResponse, statusCode, err
}

func sendGetRequestToMicrocontroller(uri string, timeoutSeconds int) (string, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, microAddress+uri, nil)
	if err != nil {
		log.Println(err.Error())
		return err.Error(), http.StatusInternalServerError, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err.Error(), resp.StatusCode, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return err.Error(), http.StatusInternalServerError, err
	}

	return string(body), http.StatusOK, nil
}
