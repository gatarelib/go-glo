package glo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jackmcguire1/go-glo/internal/utils"
)

// BaseAttachment contains generic attachment details
type BaseAttachment struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	MimeType string `json:"mime_type"`
}

// Attachment contains the information related to an attachment.
type Attachment struct {
	BaseAttachment
	CreatedDate string       `json:"created_date"`
	CreatedBy   *PartialUser `json:"created_by"`
}

// AttachmentsResp contains a list of attachments and
// if further pagination calls are available.
type AttachmentsResp struct {
	HasMore     bool          `json:"has_more"`
	Attachments []*Attachment `json:"attachments"`
}

// NewAttachment contains the information related to a new attachment.
type NewAttachment struct {
	BaseAttachment
	URL string `json:"url"`
}

// GeneratedAttachment contains the information related
// to a new attachment and the comment generated to prevent
// the new attachment from having a TTL
type GeneratedAttachment struct {
	Comment    *Comment       `json:"comment"`
	Attachment *NewAttachment `json:"attachment"`
}

var attachmentFields = []string{
	"created_date",
	"created_by",
	"filename",
	"mime_type",
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

// CreateAttachment Will create an attachment and create
// a new comment on the provided card so that the attachment
// does not have a Time To Live.
//
//https://gloapi.gitkraken.com/v1/docs/#/Attachments/post_boards__board_id__cards__card_id__attachments
func (a *Glo) CreateAttachment(
	boardID string,
	cardID string,
	description string,
	r io.Reader,
) (
	generated *GeneratedAttachment,
	err error,
) {
	attachment, err := a.uploadAttachment(boardID, cardID, r)
	if err != nil {
		return
	}

	commentInput := &CommentInput{
		Text: fmt.Sprintf("[%s](%s)", description, attachment.URL),
	}
	comment, err := a.CreateComment(boardID, cardID, commentInput)
	if err != nil {
		return
	}

	generated = &GeneratedAttachment{
		Comment:    comment,
		Attachment: attachment,
	}

	return
}

func (a *Glo) uploadAttachment(
	boardID string,
	cardID string,
	r io.Reader,
) (
	attachment *NewAttachment,
	err error,
) {
	addr := fmt.Sprintf("%s/boards/%s/cards/%s/attachments", a.BaseURI, boardID, cardID)

	resp, _, err := a.multiPartReq(http.MethodPost, addr, r, nil)
	if err != nil {
		return
	}

	attachment = &NewAttachment{}
	err = json.Unmarshal(resp, &attachment)

	return
}
