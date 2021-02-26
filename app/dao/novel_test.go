package dao

import (
	"github.com/stretchr/testify/assert"
	"log"
	"my_assistant_go/app/model"
	"testing"
)

func TestAddOrUpdateContent(t *testing.T) {
	content := &model.NovelContent{
		RawContent: "raw",
	}
	AddOrUpdateContent(content)
	assert.NotEqual(t, 0, content.Id)
	print(content.Id)
}

func TestAddOrUpdateNovel(t *testing.T) {
	novel := &model.Novel{RawId: "111"}
	AddOrUpdateNovel(novel)
}
func TestGetNovelById(t *testing.T) {
	novel, err := GetNovelById("11")
	if err != nil {
		t.Error(err)
	}
	log.Printf("%v", novel)
}
