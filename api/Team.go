package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func CreateTeam(c *gin.Context) {
	// 获取数据
	teamName := c.PostForm("teamName")
	teamLeader, err := strconv.Atoi(c.PostForm("teamLeader"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新小组
	newTeam := model.Team{
		TeamName:   teamName,
		TeamLeader: teamLeader,
	}
	err2 := database.DB.Create(&newTeam).Error
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "小组创建成功！",
	})
}

func DeleteTeam(c *gin.Context) {
	// 获取小组id
	var t model.Team
	teamId := c.PostForm("teamId")
	// 删除操作
	result := database.DB.Delete(&t, "team_id=?", teamId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"msg": "小组删除成功！",
	})
}

func GetTeam(c *gin.Context) {
	// 获取数据
	var t []model.Team
	userId := c.PostForm("userId")
	// 查找所有参与的小组
	result := database.DB.Find(&t, "team_leader=? OR team_member1=? OR team_member2=? OR team_member3=? OR team_member4=?",
		userId, userId, userId, userId, userId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "没有参与小组~"})
	}
}
