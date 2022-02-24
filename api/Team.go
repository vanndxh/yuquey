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
	teamId := c.Query("teamId")
	var t model.Team
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
	userId := c.DefaultQuery("userId", "")

	var t []model.Team
	result := database.DB.Find(&t, "team_leader=? OR team_member1=? OR team_member2=? OR team_member3=? OR team_member4=?",
		userId, userId, userId, userId, userId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}

	if result.RowsAffected != 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": t})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "msg": "没有参与小组~"})
	}
}

func GetTeamInfo(c *gin.Context) {
	teamId := c.DefaultQuery("teamId", "")

	var t model.Team
	result := database.DB.Find(&t, "team_id=?", teamId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": t})
}

func UpdateTeamInfo(c *gin.Context) {
	// 获取数据
	var t model.Team
	teamId, err := strconv.Atoi(c.PostForm("teamId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	teamName := c.PostForm("teamName")
	teamNotice := c.PostForm("teamNotice")
	// 找到对应记录
	result := database.DB.Find(&t, "team_id=?", teamId)
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
	teamId, err := strconv.Atoi(c.PostForm("teamId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	newUserId, err2 := strconv.Atoi(c.PostForm("newUserId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 找到对应记录
	result := database.DB.Find(&t, "team_id=?", teamId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 判断是否已经存在
	if newUserId == t.TeamLeader || newUserId == t.TeamMember1 || newUserId == t.TeamMember2 || newUserId == t.TeamMember3 || newUserId == t.TeamMember4 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": "小组成员已经存在！"})
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
		return
	}
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "小组成员添加成功！"})
}

func Punch(c *gin.Context) {
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	teamId, err2 := strconv.Atoi(c.PostForm("teamId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	var t model.Team
	res := database.DB.Find(&t, "team_id=?", teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	if t.TeamLeader == userId {
		countNow := t.LeaderCount
		database.DB.Model(&t).Update("leader_count", countNow+1)
	} else if t.TeamMember1 == userId {
		countNow := t.Member1Count
		database.DB.Model(&t).Update("member1_count", countNow+1)
	} else if t.TeamMember2 == userId {
		countNow := t.Member2Count
		database.DB.Model(&t).Update("member2_count", countNow+1)
	} else if t.TeamMember3 == userId {
		countNow := t.Member3Count
		database.DB.Model(&t).Update("member3_count", countNow+1)
	} else if t.TeamMember4 == userId {
		countNow := t.Member4Count
		database.DB.Model(&t).Update("member4_count", countNow+1)
	}
}

func QuitTeam(c *gin.Context) {
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	teamId, err2 := strconv.Atoi(c.PostForm("teamId"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	var t model.Team
	res := database.DB.Find(&t, "team_id=?", teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	if t.TeamLeader == userId {
		database.DB.Delete(&t)
	} else if t.TeamMember1 == userId {
		database.DB.Model(&t).Update("team_member1", 0)
	} else if t.TeamMember2 == userId {
		database.DB.Model(&t).Update("team_member2", 0)
	} else if t.TeamMember3 == userId {
		database.DB.Model(&t).Update("team_member3", 0)
	} else if t.TeamMember4 == userId {
		database.DB.Model(&t).Update("team_member4", 0)
	}
}

func GetTeamArticles(c *gin.Context) {
	teamId := c.DefaultQuery("teamId", "")

	var t model.Team
	res := database.DB.Find(&t, "team_id=?", teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	var a []model.Article
	res2 := database.DB.Find(&a, "article_author=? OR article_author=? OR article_author=? OR article_author=? OR article_author=?",
		t.TeamLeader, t.TeamMember1, t.TeamMember2, t.TeamMember3, t.TeamMember4)
	if res2.Error != nil {
		fmt.Println(res2.Error)
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": a})
}

func GetTeamMembers(c *gin.Context) {
	teamId := c.DefaultQuery("teamId", "")
	var t model.Team
	res := database.DB.Find(&t, "team_id=?", teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	var u []model.User
	res2 := database.DB.Find(&u, "user_id=? OR user_id=? OR user_id=? OR user_id=? OR user_id=?",
		t.TeamLeader, t.TeamMember1, t.TeamMember2, t.TeamMember3, t.TeamMember4)
	if res2.Error != nil {
		fmt.Println(res2.Error)
		return
	}
	c.JSON(200, gin.H{"status": 200, "data": u})
}

func GetAllTeams(c *gin.Context) {
	var ts []model.Team
	database.DB.Order("team_id").Find(&ts)
	for i := range ts {
		var u1 model.User
		database.DB.Find(&u1, "user_id=?", ts[i].TeamLeader)
		ts[i].LeaderName = u1.Username
		var u2 model.User
		database.DB.Find(&u2, "user_id=?", ts[i].TeamMember1)
		ts[i].Member1Name = u2.Username
		var u3 model.User
		database.DB.Find(&u3, "user_id=?", ts[i].TeamMember2)
		ts[i].Member2Name = u3.Username
		var u4 model.User
		database.DB.Find(&u4, "user_id=?", ts[i].TeamMember3)
		ts[i].Member3Name = u4.Username
		var u5 model.User
		database.DB.Find(&u5, "user_id=?", ts[i].TeamMember4)
		ts[i].Member4Name = u5.Username
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ts})
}
