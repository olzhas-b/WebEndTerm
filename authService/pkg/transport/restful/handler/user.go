package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/olzhas-b/PetService/authService/consts"
	"github.com/olzhas-b/PetService/authService/modules/logger"
	"github.com/olzhas-b/PetService/authService/pkg/models"
	"github.com/olzhas-b/PetService/authService/pkg/transport/restful/common"
	"github.com/olzhas-b/PetService/authService/tools"
	"log"
	"net/http"
)

func (h *Handler) CtlGetUser(ctx *fiber.Ctx) error {
	logger.WithContext(ctx.Context()).Info("information")
	id := ctx.Params("id")
	user, err := h.services.IUserService.ServiceGetUserByID(tools.StrToInt64(id))
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func (h *Handler) CtlCreateUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}
	result, err := h.services.IUserService.ServiceCreateUser(user)
	if err != nil {
		return err
	}
	return ctx.JSON(result)
}

func (h *Handler) CtlSignIn(ctx *fiber.Ctx) error {
	var userCred models.UserCredential
	if err := ctx.BodyParser(&userCred); err != nil {
		return fmt.Errorf("couldn't parse body err: %v", err)
	}

	accessToken, refreshToken, err := h.services.ServiceSignIn(userCred)
	if err != nil {
		return common.GenShortResponse(ctx, consts.DBSelectErr, "", "")
	}

	resp := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	return common.GenShortResponse(ctx, consts.Success, resp, "")
}

func (h *Handler) CtlSignOut(ctx *fiber.Ctx) error {
	token := tools.GetToken(ctx)
	err := h.services.IUserService.ServiceLogOut(token)
	return err
}

func (h *Handler) CtlRefreshToken(ctx *fiber.Ctx) error {
	token := tools.GetToken(ctx)
	accessToken, refreshToken, err := h.services.IUserService.ServiceUpdateToken(token)
	if err != nil {
		return common.GenShortResponse(ctx, consts.DBSelectErr, "", "")
	}

	resp := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	return common.GenShortResponse(ctx, consts.Success, resp, "")
}

func (h *Handler) CtlCheck(ctx *fiber.Ctx) error {
	log.Println(tools.GetToken(ctx))
	claims, err := h.services.IUserService.ServiceCheckToken(tools.GetToken(ctx))
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusUnauthorized)
		return ctx.JSON(err)
	}

	return common.GenShortResponse(ctx, consts.Success,
		map[string]interface{}{
			"id":     claims.ID,
			"scopes": []string{claims.UserType},
		}, "")
}
