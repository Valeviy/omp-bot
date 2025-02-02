package pagination

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
	"log"
)

type Pagination interface {
	GetList(chatId int64) (msg tgbotapi.MessageConfig)
}

const ListPerPageDefault = 2

type ListPagination struct {
	equipmentRequestService equipment_request.EquipmentRequestService
	perPage                 uint64
	callbackListData        CallbackListData
}

func NewListPagination(
	equipmentRequestService equipment_request.EquipmentRequestService,
	perPage uint64,
	callbackListData CallbackListData,
) *ListPagination {
	return &ListPagination{
		equipmentRequestService: equipmentRequestService,
		perPage:                 perPage,
		callbackListData:        callbackListData,
	}
}

func (l *ListPagination) GetMessageWithButtons() (string, []tgbotapi.InlineKeyboardButton) {
	outputMsgText := "Here all the equipment requests: \n\n"

	currentPage := l.callbackListData.Page
	count := l.equipmentRequestService.Count()

	if count == 0 {
		return "List with equipment requests is empty", nil
	}

	equipmentRequests, err := l.equipmentRequestService.List(currentPage, l.perPage)
	if err != nil {
		log.Printf("failed to get list of equipment requests in page %d with limit %d: %v", currentPage, l.perPage, err)
		return "Unable to get list of equipment requests for selected page", nil
	}

	for _, eq := range equipmentRequests {
		outputMsgText += eq.String()
		outputMsgText += "\n"
	}

	var buttons []tgbotapi.InlineKeyboardButton

	if currentPage > 0 {
		pagePrevData, _ := json.Marshal(CallbackListData{
			Page: currentPage - 1,
		})

		pagePrev := path.CallbackPath{
			Domain:       "business",
			Subdomain:    "equipmentRequest",
			CallbackName: "list",
			CallbackData: string(pagePrevData),
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Previous page", pagePrev.String()))
	}

	if (l.perPage * (currentPage + 1)) < count {
		pageNextData, _ := json.Marshal(CallbackListData{
			Page: currentPage + 1,
		})

		pageNext := path.CallbackPath{
			Domain:       "business",
			Subdomain:    "equipmentRequest",
			CallbackName: "list",
			CallbackData: string(pageNextData),
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Next page", pageNext.String()))
	}

	return outputMsgText, buttons
}
