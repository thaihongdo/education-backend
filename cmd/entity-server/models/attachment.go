package models

import "github.com/jinzhu/gorm"

type Attachment struct {
	gorm.Model

	Title      string `json:"title"`
	EntityID   uint   `json:"entity_id"`
	EntityType string `json:"entity_type"`

	CreatedByUserID     uint `json:"created_by_user_id"`
	CreatedByUser       User
	UpdatedByUserID     uint `json:"updated_by_user_id"`
	UpdatedByUser       User
	StoragePrefix       string `json:"storage_prefix"`
	ClientFileName      string `json:"client_file_name"`
	AttachmentName      string `json:"attachment_name"`
	AttachmentExtention string `json:"attachment_extention"`
	Description         string `json:"description"`
	AttachmentSize      int64  `json:"attachment_size"`
}

func (obj *Attachment) Get(id uint) (*Attachment, error) {
	return nil, nil
}
