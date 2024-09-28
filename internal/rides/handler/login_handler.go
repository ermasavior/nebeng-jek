package handler

// type Credentials struct {
// 	MSISDN   string `json:"msisdn"`
// 	Password string `json:"password"`
// }

// func (h *ridesHandler) LoginHandler(c *gin.Context) {
// 	var creds Credentials
// 	if err := json.NewDecoder(c.Request.Body).Decode(&creds); err != nil {
// 		http.Error(c.Writer, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	// Example: Validate credentials (you should replace this with your own validation logic)
// 	if creds.MSISDN != "081" || creds.Password != "pass" {
// 		http.Error(c.Writer, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	token, err := h.jwt.GenerateJWT(map[string]interface{}{
// 		"msisdn": creds.MSISDN,
// 	})
// 	if err != nil {
// 		http.Error(c.Writer, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(c.Writer).Encode(map[string]string{"token": token})
// }
