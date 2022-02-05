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

func GetTeams(c *gin.Context) {
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
	// 返回数据
	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "没有参与小组~"})
	}
}

func GetTeamInfo(c *gin.Context) {
	// 获取数据
	var t model.Team
	teamId := c.PostForm("teamId")
	// 查找
	result := database.DB.Find(&t, "team_id=?", teamId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 返回数据
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

func UpdateTeamInfo(c *gin.Context) {
	// 获取数据
	var t model.Team
	teamName := c.PostForm("teamName")
	teamNotice := c.PostForm("teamNotice")
	// 找到对应记录
	result := database.DB.Find(&t, "team_id=?", 1)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// update
	database.DB.Model(&t).Update("team_name", teamName)
	database.DB.Model(&t).Update("team_notice", teamNotice)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "小组信息修改成功！"})
}

func AddTeamUser(c *gin.Context) {
	// 获取数据
	var t model.Team
	teamId := c.PostForm("teamId")
	newUserId := c.PostForm("newUserId")
	// 找到对应记录
	result := database.DB.Find(&t, "team_id=?", teamId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 把新用户id放入第一个空位
	if t.TeamMember1 == 0 {
		database.DB.Model(&t).Update("team_member1", newUserId)
	} else if t.TeamMember2 == 0 {
		database.DB.Model(&t).Update("team_member2", newUserId)
	} else if t.TeamMember3 == 0 {
		database.DB.Model(&t).Update("team_member3", newUserId)
	} else if t.TeamMember4 == 0 {
		database.DB.Model(&t).Update("team_member4", newUserId)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": "小组成员已满，无法加入！"})
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "小组信息修改成功！"})
}
