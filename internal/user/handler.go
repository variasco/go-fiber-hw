package user

import (
	"fmt"
	"variasco/go-fiber-hw/pkg/tadapter"
	"variasco/go-fiber-hw/pkg/validator"
	"variasco/go-fiber-hw/views/components"
	"variasco/go-fiber-hw/views/pages"

	v "github.com/gobuffalo/validate"
	vs "github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type UserHandlerDeps struct {
	Router     fiber.Router
	Logger     *zerolog.Logger
	Repository *UserRepository
}

type UserHandler struct {
	router     fiber.Router
	logger     *zerolog.Logger
	repository *UserRepository
}

func NewHandler(deps UserHandlerDeps) {
	handler := &UserHandler{
		router:     deps.Router,
		logger:     deps.Logger,
		repository: deps.Repository,
	}

	handler.router.Get("/register", handler.register)
	handler.router.Post("/register", handler.registerByForm)
}

func (handler *UserHandler) register(c *fiber.Ctx) error {
	return tadapter.Render(c, pages.Register())
}

func (handler *UserHandler) registerByForm(c *fiber.Ctx) error {
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

	user, err := handler.repository.create(form)
	if err != nil {
		handler.logger.Error().Err(err).Send()

		return tadapter.Render(c, components.Notification(components.NotificationProps{
			Messages: []string{"Ошибка при создании пользователя. Попробуйте позднее"},
			Type:     components.NotificationError,
		}))
	}

	return tadapter.Render(c, components.Notification(components.NotificationProps{
		Messages: []string{fmt.Sprintf("Пользователь: %s успешно создан", user.Email)},
		Type:     components.NotificationSuccess,
	}))
}
