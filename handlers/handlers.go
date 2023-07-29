package handlers

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"html/template"
	"net/http"
	"net/mail"
	"recipio-admin/models"
)

func SendInvite(pb *pocketbase.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request models.SendInviteRequest
		bindErr := c.Bind(&request)
		if bindErr != nil {
			return c.JSON(http.StatusBadRequest, bindErr.Error())
		}

		tmpl, tmplErr := template.ParseFiles("templates/member_invitation.html")
		if tmplErr != nil {
			return c.JSON(http.StatusInternalServerError, tmplErr.Error())
		}

		var emailBodyBuffer bytes.Buffer
		bufErr := tmpl.Execute(&emailBodyBuffer, request)
		if bufErr != nil {
			return c.JSON(http.StatusInternalServerError, bufErr.Error())
		}

		email := &mailer.Message{
			From: mail.Address{
				Address: pb.Settings().Meta.SenderAddress,
				Name:    pb.Settings().Meta.SenderName,
			},
			To: []mail.Address{
				{Address: request.Email},
			},
			Subject: "You were invited to a Recipio family",
			HTML:    emailBodyBuffer.String(),
		}

		err := pb.NewMailClient().Send(email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("Invite sent to %s", request.Email),
		})
	}
}
