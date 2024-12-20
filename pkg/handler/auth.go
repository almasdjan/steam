package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"orden/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary SignUp
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.SignUp true "account info"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /auth/signup [post]
func (h *Handler) signup(c *gin.Context) {
	var input models.SignUp

	if err := c.BindJSON(&input); err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  "unknown",
			"error": err.Error(),
		}).Error("Failed to bind JSON in signup")
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	err := input.Validate()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  input.Email,
			"error": "Invalid email format",
		}).Error("Validation error in signup")
		NewErrorResponce(c, http.StatusBadRequest, "формат почты некорректный")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  input.Email,
			"error": err.Error(),
		}).Error("Failed to create user")
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	/*
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	*/

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user":  input.Email,
			"error": err.Error(),
		}).Error("Failed to generate token")
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"user": input.Email,
		"id":   id,
	}).Info("User successfully signed up")

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"id":    id,
	})
}

// @Summary go to steam
// @Security ApiKeyAuth
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /api/profile/steam [get]
func (h *Handler) signupSteam(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "user not found")
		return
	}

	steamRedirectURL := "https://steamcommunity.com/openid/login" +
		"?openid.ns=http://specs.openid.net/auth/2.0" +
		"&openid.mode=checkid_setup" +
		"&openid.return_to=https://gamepal.kz/auth/steam/callback?user_id=" + strconv.Itoa(userId) +
		"&openid.realm=https://gamepal.kz" +
		"&openid.claimed_id=http://specs.openid.net/auth/2.0/identifier_select" +
		"&openid.identity=http://specs.openid.net/auth/2.0/identifier_select"

	c.Redirect(http.StatusFound, steamRedirectURL)

}

// @Summary go to just steam
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /auth/juststeam [get]
func (h *Handler) signupJustSteam(c *gin.Context) {

	steamRedirectURL := "https://steamcommunity.com/openid/login" +
		"?openid.ns=http://specs.openid.net/auth/2.0" +
		"&openid.mode=checkid_setup" +
		"&openid.return_to=https://gamepal.kz/auth/juststeam/callback" +
		"&openid.realm=https://gamepal.kz" +
		"&openid.claimed_id=http://specs.openid.net/auth/2.0/identifier_select" +
		"&openid.identity=http://specs.openid.net/auth/2.0/identifier_select"

	c.Redirect(http.StatusFound, steamRedirectURL)

}

// @Summary go to just steamm
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /auth/juststeamm [get]
func (h *Handler) signupJustSteamm(c *gin.Context) {

	steamRedirectURL := "https://steamcommunity.com/openid/login" +
		"?openid.ns=http://specs.openid.net/auth/2.0" +
		"&openid.mode=checkid_setup" +
		"&openid.return_to=https://gamepal.kz/auth/juststeam/callback" +
		"&openid.realm=https://gamepal.kz" +
		"&openid.claimed_id=http://specs.openid.net/auth/2.0/identifier_select" +
		"&openid.identity=http://specs.openid.net/auth/2.0/identifier_select"

	c.Redirect(http.StatusFound, steamRedirectURL)

}

// @Summary get from steam
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /auth/juststeam/callback [get]
func (h *Handler) callbackJustSteam(c *gin.Context) {

	params := c.Request.URL.Query()

	fmt.Println("Steam OpenID callback parameters:", params)

	claimedID := params.Get("openid.claimed_id")
	if claimedID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing claimed_id"})
		return
	}

	openidSig := params.Get("openid.sig")
	if openidSig == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing signature"})
		return
	}

	verified, err := verifySteamOpenID(params)
	if err != nil || !verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to verify Steam OpenID"})
		return
	}

	steamID := extractSteamID(claimedID)
	if steamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Steam ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Steam login successful", "steam_id": steamID})
}

// @Summary get from steam
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Router /auth/steam/callback [get]
func (h *Handler) callbackSteam(c *gin.Context) {

	userID := c.Query("user_id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
		return
	}

	params := c.Request.URL.Query()

	fmt.Println("Steam OpenID callback parameters:", params)

	claimedID := params.Get("openid.claimed_id")
	if claimedID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing claimed_id"})
		return
	}

	openidSig := params.Get("openid.sig")
	if openidSig == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing signature"})
		return
	}

	verified, err := verifySteamOpenID(params)
	if err != nil || !verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to verify Steam OpenID"})
		return
	}

	steamID := extractSteamID(claimedID)
	if steamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Steam ID"})
		return
	}

	userId, err := strconv.Atoi(userID)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "incorrect user")
		return
	}

	err = h.services.Authorization.AddSteamId(userId, steamID)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "произашла ошибка")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Steam login successful", "steam_id": steamID, "user_id": userID})
}

// Вспомогательная функция для извлечения Steam ID
func extractSteamID(claimedID string) string {
	const steamPrefix = "https://steamcommunity.com/openid/id/"
	if len(claimedID) > len(steamPrefix) && claimedID[:len(steamPrefix)] == steamPrefix {
		return claimedID[len(steamPrefix):]
	}
	return ""
}

func verifySteamOpenID(params url.Values) (bool, error) {
	params.Add("openid.mode", "check_authentication")

	// Отправляем POST-запрос на OpenID сервер Steam
	resp, err := http.PostForm("https://steamcommunity.com/openid/login", params)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Steam возвращает строку "is_valid:true" для успешной проверки
	if strings.Contains(string(body), "is_valid:true") {
		return true, nil
	}

	return false, nil
}

// @Summary LogIn
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.Login true "email password"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input models.Login

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "password or email is not correct")
		return
	}

	user, err := h.services.Authorization.GetUserByEmail(input.Email)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "cannot find user by email")
		return
	}

	err = h.services.Authorization.UpdateDeviceToken(input.DeviceToken, user.Id)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "failed to update device token")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// @Summary LogIn for admin
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.Login true "email password"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /auth/adminlogin [post]
func (h *Handler) loginForAdmin(c *gin.Context) {
	var input models.Login

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	roleId, err := h.services.Authorization.GetRoleId(input.Email)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "failed to get role_id")
		return
	}
	logrus.Printf("role id %d", roleId)

	if roleId != 2 {
		NewErrorResponce(c, http.StatusBadRequest, "not admin")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "password or email is not correct")
		return
	}

	user, err := h.services.Authorization.GetUserByEmail(input.Email)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "cannot find user by email")
		return
	}

	err = h.services.Authorization.UpdateDeviceToken(input.DeviceToken, user.Id)
	if err != nil {
		NewErrorResponce(c, http.StatusInternalServerError, "failed to update device token")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// @Summary Reset Password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body models.ResetPasswd true "password"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /auth/resetpasswd [patch]
func (h *Handler) resetpasswd(c *gin.Context) {

	var input models.ResetPasswd

	if err := c.BindJSON(&input); err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusBadRequest, "неправильные данные")
		return
	}
	if input.Password != input.Password2 {
		NewErrorResponce(c, http.StatusBadRequest, "пароль не совпадает")
		return
	}
	user, err := h.services.Authorization.GetUserByEmail(input.Email)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "user not found")
		return
	}
	logrus.Println(user.Id)

	logrus.Println(user.Id)

	err = h.services.Authorization.ResetPasswd(user.Id, input.Password)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "произашла ошибка")
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

// @Summary Change Password
// @Security ApiKeyAuth
// @Tags profile
// @Accept json
// @Produce json
// @Param input body models.ChangePasswd true "password"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/profile/changepasswd [patch]
func (h *Handler) changePasswd(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "user not found")
		return
	}
	var input models.ChangePasswd

	if err := c.BindJSON(&input); err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusBadRequest, "incorrect input")
		return
	}

	if input.NewPassword != input.NewPassword2 {
		NewErrorResponce(c, http.StatusBadRequest, "new passwords do not match")
		return
	}

	err = h.services.Authorization.CheckPasswd(userId, input.Password)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Authorization.ResetPasswd(userId, input.NewPassword)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "failed to change password")
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

// @Summary Change Username
// @Security ApiKeyAuth
// @Tags profile
// @Accept json
// @Produce json
// @Param input body models.ChangeUserInfo true "user info"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/profile/changeusername [patch]
func (h *Handler) updateUsername(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, "user not found")
		return
	}
	var input models.ChangeUserInfo

	if err := c.BindJSON(&input); err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusBadRequest, "incorrect input")
		return
	}

	err = h.services.Authorization.UpdateUseranme(userId, input.Name)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "failed to change username")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

// @Summary get profile
// @Security ApiKeyAuth
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/profile/ [get]
func (h *Handler) getProfile(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusBadRequest, "user not found")
		return
	}

	user, err := h.services.Authorization.GetProfile(userId)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "failed to get user info")
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary delete profile
// @Security ApiKeyAuth
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/profile/ [delete]
func (h *Handler) deleteUser(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusBadRequest, "user not found")
		return
	}

	err = h.services.Authorization.DeleteUser(userId)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "failed to get user info")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

// @Summary get all users
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.Authorization.GetUsers()
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	admins, err := h.services.Authorization.GetAdmins()
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"admins": admins,
		"users":  users,
	})

}

// @Summary give admin rights
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/adminrights/{id} [patch]
func (h *Handler) giveAdminRight(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Authorization.MakeAdmin(userId)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "failed to give admin rights")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}

// @Summary remove admin rights
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} map[string]any
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/admin/adminrights/{id} [delete]
func (h *Handler) removeAdminRight(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Authorization.RemoveAdmin(userId)
	if err != nil {
		logrus.Println(err.Error())
		NewErrorResponce(c, http.StatusInternalServerError, "failed to remove admin rights")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Ok",
	})
}
