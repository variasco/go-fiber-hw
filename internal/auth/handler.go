package auth

import (
	"variasco/go-fiber-hw/pkg/tadapter"
	"variasco/go-fiber-hw/pkg/validator"
	"variasco/go-fiber-hw/views/components"
	"variasco/go-fiber-hw/views/pages"

	v "github.com/gobuffalo/validate"
	vs "github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type AuthHandlerDeps struct {
	Router fiber.Router
	Logger *zerolog.Logger
}

type AuthHandler struct {
	router fiber.Router
	logger *zerolog.Logger
}

func NewHandler(deps AuthHandlerDeps) {
	handler := &AuthHandler{
		router: deps.Router,
		logger: deps.Logger,
	}

	handler.router.Get("/register", handler.register)
	handler.router.Post("/register", handler.registerByForm)
}

func (handler *AuthHandler) register(c *fiber.Ctx) error {
	return tadapter.Render(c, pages.Register())
}

func (handler *AuthHandler) registerByForm(c *fiber.Ctx) error {
	form := RegisterForm{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	errors := v.Validate(
		&vs.StringIsPresent{Name: "Name", Field: form.Name, Message: "Имя: Обязательное поле"},
		&vs.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email: Неверный или не указан"},
		&vs.StringIsPresent{Name: "Password", Field: form.Password, Message: "Пароль: Обязательное поле"},
	)

	if len(errors.Errors) != 0 {
		handler.logger.Warn().Msgf("Validation errors: %v", errors.Errors)

		return tadapter.Render(c, components.Notification(components.NotificationProps{
			Messages: validator.FormatErrors(errors),
			Type:     components.NotificationError,
		}))
	}

	// TODO
	return nil
}
