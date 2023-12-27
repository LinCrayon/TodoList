package service

import (
	"time"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

type ShowTaskService struct {
}

type DeleteTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

// Create 创建一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := e.SUCCESS
	model.DB.First(&user, id) //未查到，user量将保持零值（该模型类型的零值），并且不会返回错误。如需检查是否找到记录，用 gorm.IsRecordNotFoundError(err) 来检查错误。
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

// Show 查询一条备忘录 	tid是请求中获取的参数 id
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

// List 返回该用户的备忘录 （分页查询） uid:用户id
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// Update 更新备忘录（根据id）
func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新数据成功",
	}
}

// Search 查询备忘录 （模糊查询） uid:用户id
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	//Preload("User")：预加载相关联的关系，预加载Task模型关联的User模型。
	//LIKE操作符:筛选标题或内容中包含 service.Info 的记录。% 是通配符，表示匹配任意字符
	//Limit指定返回的记录数目，Offset指定从第几条记录开始返回。
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// Delete 删除备忘录 tid为备忘录id
func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, tid).Error
	err = model.DB.Delete(&task, tid).Error
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   "删除成功",
	}
}
