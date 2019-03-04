package glo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jackmcguire1/go-glo/internal/utils"
)

// Attachment contains the information related to an attachment.
type Attachment struct {
	ID          string       `json:"id"`
	Filename    string       `json:"filename"`
	MimeType    string       `json:"mime_type"`
	CreatedDate string       `json:"created_date"`
	CreatedBy   *PartialUser `json:"created_by"`
}

// AttachmentsResp contains a list of attachments and
// if further pagination calls are available.
type AttachmentsResp struct {
	HasMore     bool          `json:"has_more"`
	Attachments []*Attachment `json:"attachments"`
}

var attachmentFields = []string{
	"created_date",
	"created_by",
	"filename",
	"mime_type",
}

// CreateAttachment After making this call, you must make another call to either
// update the card's description or add/update one of the card's comments
// and include the attachment url (in markdown format) in the text.
// FORMAT: [ANY_TEXT](ATTACHMENT_URL).
//
// If the attachment url is not in the card description or
// one of its comments within 1 hour of the attachment being created,
// it will be deleted.
// https://gloapi.gitkraken.com/v1/docs/#/Attachments/post_boards__board_id__cards__card_id__attachments
func (a *Glo) CreateAttachment(
	boardID string,
	cardID string,
	r io.Reader,
) (
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/attachments", a.BaseURI, boardID, cardID)

	_, _, err = a.multiPartReq(http.MethodPost, addr, r, nil)

	return
}

// GetAttachments get attachments related to a card
// https://gloapi.gitkraken.com/v1/docs/#/Attachments/get_boards__board_id__cards__card_id__attachments
func (a *Glo) GetAttachments(
	boardID string,
	cardID string,
	page int,
	limit int,
	sortDesc bool,
) (
	attachmentsResp *AttachmentsResp,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/attachments", a.BaseURI, boardID, cardID)

	q := utils.AddFields(attachmentFields)
	q.Set("page", fmt.Sprint(page))
	q.Set("per_page", fmt.Sprint(limit))

	if sortDesc {
		q.Set("sort", "desc")
	}

	resp, headers, err := a.jsonReq(http.MethodGet, addr, nil, q)
	if err != nil {
		return
	}

	attachmentsResp = &AttachmentsResp{}
	err = json.Unmarshal(resp, &attachmentsResp.Attachments)
	if err != nil {
		return
	}

	hasMore, err := strconv.ParseBool(headers.Get("has-more"))
	if err != nil {
		return
	}
	attachmentsResp.HasMore = hasMore

	return
}
