package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func GetTagById(id int) (tag Tag) {
	maps := make(map[string]interface{})
	maps["id"] = id
	db.Where(maps).Find(&tag)

	return
}

func AddTag(name string, state int, createdBy string) bool {

	db.Create(&Tag{
		Name:       name,
		State:      state,
		CreatedBy:  createdBy,
		ModifiedBy: createdBy,
	})

	return true
}

func EditTag(id int, name string, state int, updatedBy string) bool {

	data := make(map[string]interface{})

	data["name"] = name
	data["state"] = state
	data["updatedBy"] = updatedBy
	data["updated_at"] = ""

	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}

func CheckExistsById(id int) bool {

	var tag Tag

	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	} else {
		return false
	}
}

func DeleteTagById(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
