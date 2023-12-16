package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Message string `json:"Message"`
}

func main() {
	// Đăng ký một hàm xử lý cho đường dẫn "/api/data"
	http.HandleFunc("/api/get", getDataHandler)
	http.HandleFunc("/api/post", postDataHandler)

	// Khởi động máy chủ tại cổng 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// Hàm xử lý yêu cầu POST cho "/api/data"
// Hàm xử lý yêu cầu POST cho "/api/data"
func postDataHandler(w http.ResponseWriter, r *http.Request) {
	// Tạo một đối tượng Data từ dữ liệu POST
	var inputData Data
	err := json.NewDecoder(r.Body).Decode(&inputData)
	log.Default().Printf(inputData.Message)
	if err != nil {
		// Nếu có lỗi, trả về lỗi 400 (Bad Request)
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Ở đây, bạn có thể thực hiện các xử lý với dữ liệu được gửi đến (inputData)
	// ...

	// Phản hồi thành công
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))
}

// Hàm xử lý yêu cầu GET cho "/api/data"
func getDataHandler(writer http.ResponseWriter, request *http.Request) {
	// Tạo một đối tượng Data đơn giản
	data := Data{
		Message: "Hello, this is get api !",
	}

	// Chuyển đổi đối tượng Data thành chuỗi JSON
	response, err := json.Marshal(data)
	if err != nil {
		// Nếu có lỗi, trả về lỗi 500 (Internal Server Error)
		http.Error(writer, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Thiết lập tiêu đề cho phản hồi là JSON
	writer.Header().Set("Content-Type", "application/json")

	// Ghi dữ liệu JSON vào ResponseWriter
	writer.Write(response)
}
