package pages

import "github.com/gofiber/fiber/v3"

type HomeHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	h.router.Get("/", h.home)
}

func (h *HomeHandler) home(c fiber.Ctx) error {
	data := struct {
		MenuItems []string
	}{
		MenuItems: []string{"#Еда", "#Животные", "#Машины", "#Спорт", "#Музыка", "#Технологии", "#Прочее"},
	}
	return c.Render("page", data)
}
