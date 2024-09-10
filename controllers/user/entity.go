package usercontroller

type InputRegister struct {
	Fullname string `json: "fullname"`
	Email    string `json: "Email"`
	Password string `json: "Password"`
}
