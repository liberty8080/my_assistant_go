package dao

import (
	"github.com/syndtr/goleveldb/leveldb/errors"
	"log"
	"my_assistant_go/app/model"
	"my_assistant_go/boot"
)

func AddOrUpdateChapter(chapter *model.NovelChapter) {
	has, err := boot.Engine.Get(chapter)
	if has {
		_, err = boot.Engine.Update(chapter, chapter)
	} else {
		_, err = boot.Engine.Insert(chapter)
	}
	if err != nil {
		log.Println(err)
	}
}

func AddOrUpdateContent(content *model.NovelContent) {
	has, err := boot.Engine.Where("raw_content=?", content.RawContent).Get(content)

	if has {
		_, err = boot.Engine.Update(content, content)
	} else {
		_, err = boot.Engine.Insert(content)
	}

	if err != nil {
		log.Println("novel content process failed", err)
	}
}

func AddOrUpdateNovel(novel *model.Novel) int {
	n, err := GetNovelById(novel.RawId)
	if err != nil {
		_, err = boot.Engine.Insert(novel)
	} else {
		novel.Id = n.Id
		_, err = boot.Engine.ID(n.Id).Update(&novel)
	}
	if err != nil {
		log.Print(err)
	}
	log.Printf("novelId:%d\n", novel.Id)
	return novel.Id
}

func GetNovelById(id string) (model.Novel, error) {
	var novel model.Novel
	has, err := boot.Engine.Where("raw_id=?", id).Get(&novel)
	if err != nil || !has {
		if !has {
			err = errors.New("没有该小说,id=" + id)
		}
		return model.Novel{}, err
	}
	return novel, err
}
