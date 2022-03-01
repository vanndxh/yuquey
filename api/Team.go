package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"yuquey/database"
	"yuquey/model"
)

func CreateTeam(c *gin.Context) {
	// 获取数据
	teamName := c.PostForm("teamName")
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建新小组
	newTeam := model.Team{
		TeamName:   teamName,
		TeamLeader: userId,
		TeamNotice: "暂无~",
	}
	err2 := database.DB.Create(&newTeam).Error
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	// 小组成员中，加入组长
	tp, _ := time.ParseDuration("-24h")
	lastPunchTime := time.Now().Add(tp)
	newTeamUser := model.TeamUser{
		UserId:        userId,
		TeamId:        newTeam.TeamId,
		LastPunchTime: lastPunchTime,
	}
	database.DB.Create(&newTeamUser)
	// 返回结果
	c.JSON(200, gin.H{"msg": "小组创建成功！"})
}
func AddTeamUser(c *gin.Context) {
	// 获取数据
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
	// 判断是否已经存在
	var t model.TeamUser
	result := database.DB.Find(&t, "team_id=? AND user_id=?", teamId, newUserId)
	if result.RowsAffected != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": "小组成员已经存在！"})
		return
	}
	// 小组成员中，加入组员
	tp, _ := time.ParseDuration("-24h")
	lastPunchTime := time.Now().Add(tp)
	newTeamUser := model.TeamUser{
		UserId:        newUserId,
		TeamId:        teamId,
		Position:      1,
		LastPunchTime: lastPunchTime,
	}
	database.DB.Create(&newTeamUser)
	// 返回结果
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": "小组成员添加成功！"})
}

func DeleteTeam(c *gin.Context) {
	teamId := c.Query("teamId")
	// 删除操作
	var t model.Team
	result := database.DB.Delete(&t, "team_id=?", teamId)
	if result.Error != nil {
		fmt.Println(1)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 删除该小组相关组员信息
	var tu model.TeamUser
	res := database.DB.Delete(&tu, "team_id=?", teamId)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res.Error.Error()})
		return
	}
	// 返回结果
	c.JSON(200, gin.H{"msg": "小组删除成功！"})
}
func DeleteTeamUser(c *gin.Context) {
	teamId := c.Query("teamId")
	teamUser := c.Query("teamUser")

	var tu model.TeamUser
	res := database.DB.Delete(&tu, "team_id=? AND user_id=?", teamId, teamUser)
	if res.Error != nil {
		c.JSON(400, gin.H{"msg": res.Error.Error()})
	}

	c.JSON(200, gin.H{"msg": "ok！"})
}

func UpdateTeamInfo(c *gin.Context) {
	// 获取数据
	teamId, err := strconv.Atoi(c.PostForm("teamId"))
	if err != nil {
		fmt.Println(err)
		return
	}
	teamName := c.PostForm("teamName")
	teamNotice := c.PostForm("teamNotice")
	// 找到对应记录
	var t model.Team
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

	var tu model.TeamUser
	res := database.DB.Find(&tu, "team_id=? AND user_id=?", teamId, userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	// 对打卡日期进行判定，今日是否已经打卡
	if time.Now().Day() == tu.LastPunchTime.Day() && time.Now().Month() == tu.LastPunchTime.Month() && time.Now().Year() == tu.LastPunchTime.Year() {
		c.JSON(400, gin.H{"msg": "今天已经打卡！"})
		return
	} else {
		punchNow := tu.Punch
		database.DB.Model(&tu).Where("team_id=? AND user_id=?", teamId, userId).Update("punch", punchNow+1)
		database.DB.Model(&tu).Where("team_id=? AND user_id=?", teamId, userId).Update("last_punch_time", time.Now())

		c.JSON(200, gin.H{"msg": "ok"})
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
	var tu model.TeamUser
	database.DB.Find(&tu, "team_id=? AND user_id=?", teamId, userId)
	if tu.Position == 0 {
		// 删除操作
		var t model.Team
		result := database.DB.Delete(&t, "team_id=?", teamId)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
			return
		}
		// 删除该小组相关组员信息
		var tus model.TeamUser
		res := database.DB.Delete(&tus, "team_id=?", teamId)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res.Error.Error()})
			return
		}
	} else if tu.Position == 1 {
		var tuu model.TeamUser
		res2 := database.DB.Delete(&tuu, "team_id=? AND user_id=?", teamId, userId)
		if res2.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": res2.Error.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"msg": "ok"})
}

func GetTeamArticles(c *gin.Context) {
	teamId := c.Query("teamId")

	// 根据teamId找到有哪些用户(tus
	var tus []model.TeamUser
	res := database.DB.Find(&tus, "team_id=?", teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	// 从tus中截取出所有的用户id
	userIds := make([]int, len(tus))
	for i := range tus {
		userIds[i] = tus[i].UserId
	}
	// 根据用户找到所有文章
	var as []model.Article
	res2 := database.DB.Where("article_author IN (?)", userIds).Find(&as)
	if res2.Error != nil {
		fmt.Println(res2.Error)
		return
	}
	// 再给每个找到的a获取authorName
	for j := range as {
		var u model.User
		database.DB.Find(&u, "user_id=?", as[j].ArticleAuthor)
		as[j].AuthorName = u.Username
	}

	c.JSON(200, gin.H{"status": 200, "data": as})
}
func GetTeamMembers(c *gin.Context) {
	teamId := c.Query("teamId")

	var tus []model.TeamUser
	res := database.DB.Order("position").Find(&tus, "team_id=?", teamId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	for i := range tus {
		var u model.User
		database.DB.Find(&u, "user_id=?", tus[i].UserId)
		tus[i].Username = u.Username
		if tus[i].Position == 0 {
			tus[i].PositionName = "组长"
		} else {
			tus[i].PositionName = "组员"
		}
	}

	c.JSON(200, gin.H{"status": 200, "data": tus})
}
func GetAllTeams(c *gin.Context) {
	var ts []model.Team
	database.DB.Order("team_id").Find(&ts)

	for i := range ts {
		var u model.User
		database.DB.Find(&u, "user_id=?", ts[i].TeamLeader)
		ts[i].LeaderName = u.Username
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ts})
}
func GetTeams(c *gin.Context) { // 根据用户获取用户参与的小组
	userId := c.Query("userId")
	// 根据用户id获取所有参与的小组
	var tus []model.TeamUser
	result := database.DB.Find(&tus, "user_id=?", userId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	// 把所有得到的小组id放一起
	teamIds := make([]int, len(tus))
	for i := range tus {
		teamIds[i] = tus[i].TeamId
	}
	// 根据这些id去获取ts
	var ts []model.Team
	res := database.DB.Find(&ts, teamIds)
	if res.Error != nil {
		c.JSON(400, gin.H{"msg": res.Error.Error()})
		return
	}
	// 对每个t内的LeaderName进行赋值
	for j := range ts {
		var u model.User
		database.DB.Find(&u, "user_id=?", ts[j].TeamLeader)
		ts[j].LeaderName = u.Username
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": ts})
}
func GetTeamInfo(c *gin.Context) {
	teamId := c.Query("teamId")

	var t model.Team
	result := database.DB.Find(&t, "team_id=?", teamId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "msg": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": t})
}
