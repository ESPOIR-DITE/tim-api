package domain

import (
	"errors"
	"net/http"
)

type Video struct {
	Id           string  `json:"id" gorm:"primaryKey"`
	Title        string  `json:"title"`
	Date         string  `json:"date"`
	DateUploaded string  `json:"dateUploaded"`
	Description  string  `json:"description"`
	IsPrivate    bool    `json:"isPrivate"`
	Price        float64 `json:"price"`
	URL          string  `json:"url"`
}

func (v Video) Bind(r *http.Request) error {
	if v.Title == "" && v.Description == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoData struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Picture  []byte `json:"picture"`
	Video    []byte `json:"video" gorm:"-:all"`
	FileType string `json:"fileType"`
	FileSize string `json:"fileSize"`
}

func (v VideoData) Bind(r *http.Request) error {
	if v.Id == "" && len(v.Picture) != 0 && len(v.Video) != 0 {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoCategory struct {
	Id         string `json:"id" gorm:"primaryKey"`
	VideoId    string `json:"videoId"`
	CategoryId string `json:"categoryId"`
}

func (v VideoCategory) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CategoryId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoComment struct {
	Id        string `json:"id" gorm:"primaryKey"`
	VideoId   string `json:"videoId"`
	CommentId string `json:"categoryId"`
}

func (v VideoComment) Bind(r *http.Request) error {
	if v.VideoId == "" && v.CommentId == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type Category struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c Category) Bind(r *http.Request) error {
	if c.Name == "" {
		return errors.New("missing required fields")
	}
	return nil
}

type VideoRelated struct {
	CurrentVideoId string `json:"currentVideo" gorm:"primaryKey"`
	RelatedVideoId string `json:"relatedVideoId"`
}
