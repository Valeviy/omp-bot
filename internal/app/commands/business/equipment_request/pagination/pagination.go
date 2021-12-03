package pagination

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/logger"
	"github.com/ozonmp/omp-bot/internal/service/business/equipment_request"
)

const listPaginationLogTag = "ListPagination"

//Pagination is an interface for messages with pagination
type Pagination interface {
	GetMessageWithButtons(ctx context.Context) (string, []tgbotapi.InlineKeyboardButton)
}

//ListPagination is a list with pagination buttons
type ListPagination struct {
	equipmentRequestService equipment_request.EquipmentRequestService
	perPage                 uint64
	callbackListData        CallbackListData
}

//NewListPagination returns a new ListPagination
func NewListPagination(
	equipmentRequestService equipment_request.EquipmentRequestService,
	perPage uint64,
	callbackListData CallbackListData,
) Pagination {
	return &ListPagination{
		equipmentRequestService: equipmentRequestService,
		perPage:                 perPage,
		callbackListData:        callbackListData,
	}
}

//GetMessageWithButtons get a message with list of items and pagination buttons
func (l *ListPagination) GetMessageWithButtons(ctx context.Context) (string, []tgbotapi.InlineKeyboardButton) {
	outputMsgText := "Here all the equipment requests: \n\n"

	currentPage := l.callbackListData.Page
	equipmentRequests, total, err := l.equipmentRequestService.List(ctx, currentPage, l.perPage)

	if err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: equipmentRequestService.List failed", listPaginationLogTag),
			"err", err,
			"page", currentPage,
			"limit", l.perPage,
		)

		return "Unable to get list of equipment requests for selected page", nil
	}

	if total == 0 {
		return "List with equipment requests is empty", nil
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

	if (l.perPage * (currentPage + 1)) < total {
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
