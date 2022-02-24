package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func HandleFollow(c *gin.Context) {
	up, err := strconv.Atoi(c.PostForm("up"))
	if err != nil {
		fmt.Println(err)
		return
	}
	follower, err2 := strconv.Atoi(c.PostForm("follower"))
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	handle, err3 := strconv.Atoi(c.PostForm("handle"))
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	if handle == 0 {
		newFollow := model.Follow{
			UpId:       up,
			FollowerId: follower,
		}
		err4 := database.DB.Create(&newFollow).Error
		if err4 != nil {
			fmt.Println(err4)
			return
		}
	} else {
		var f model.Follow
		res := database.DB.Delete(&f, "up_id=? AND follower_id=?", up, follower)
		if res.Error != nil {
			fmt.Println(res.Error)
			return
		}
	}

	var u1 model.User
	database.DB.Find(&u1, "user_id=?", up)
	followerNow := u1.FollowerAmount
	if handle == 0 {
		database.DB.Model(&u1).Update("follower_amount", followerNow+1)
	} else {
		database.DB.Model(&u1).Update("follower_amount", followerNow-1)
	}

	var u2 model.User
	database.DB.Find(&u2, "user_id=?", follower)
	followNow := u2.FollowAmount
	if handle == 0 {
		database.DB.Model(&u2).Update("follow_amount", followNow+1)
	} else {
		database.DB.Model(&u2).Update("follow_amount", followNow-1)
	}

	c.JSON(200, gin.H{
		"msg": "成功！",
	})
}

func GetUps(c *gin.Context) {
	userId := c.Query("userId")
	var fs []model.Follow
	res := database.DB.Find(&fs, "follower_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := range fs {
		var u model.User
		database.DB.Find(&u, "user_id=?", fs[i].UpId)
		fs[i].UpName = u.Username
	}
	c.JSON(200, gin.H{"data": fs})
}

func GetFollowers(c *gin.Context) {
	userId := c.Query("userId")
	var fs []model.Follow
	res := database.DB.Find(&fs, "up_id=?", userId)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}
	for i := range fs {
		var u model.User
		database.DB.Find(&u, "user_id=?", fs[i].FollowerId)
		fs[i].FollowerName = u.Username
	}
	c.JSON(200, gin.H{"data": fs})
}

func GetIsFollowed(c *gin.Context) {
	userId := c.Query("userId")
	upId := c.Query("upId")
	var f model.Follow
	res := database.DB.Find(&f, "up_id=? AND follower_id=?", upId, userId)
	if res.Error != nil {
		c.JSON(200, gin.H{"data": false, "msg": "no find => false"})
		return
	}
	c.JSON(200, gin.H{"data": true})
}
