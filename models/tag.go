package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//GetTags
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&tags)
	return
}

//GetTagTotal
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//GetTagById
func GetTagById(id int) (tag Tag) {
	maps := make(map[string]interface{})
	maps["id"] = id
	db.Where(maps).Find(&tag)

	return
}

//AddTag
func AddTag(name string, state int, createdBy string) bool {

	db.Create(&Tag{
		Name:       name,
		State:      state,
		CreatedBy:  createdBy,
		ModifiedBy: createdBy,
	})

	return true
}

//EditTag
func EditTag(id int, name string, state int, updatedBy string) bool {

	data := make(map[string]interface{})

	data["name"] = name
	data["state"] = state
	data["updatedBy"] = updatedBy

	db.Model(&Tag{}).Where("id = ?", id).Update(data)

	return true
}

//CheckExistsById
func CheckExistsById(id int) bool {

	var tag Tag

	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	} else {
		return false
	}
}

//DeleteTagById
func DeleteTagById(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
