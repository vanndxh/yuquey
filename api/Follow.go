package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"yuquey/database"
	"yuquey/model"
)

func AddFollow(c *gin.Context) {
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
	newFollow := model.Follow{
		Up:       up,
		Follower: follower,
	}
	err3 := database.DB.Create(&newFollow).Error
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	var u1 model.User
	database.DB.Find(&u1, "user_id=?", up)
	followerNow := u1.FollowerAmount
	database.DB.Model(&u1).Update("follower_amount", followerNow+1)

	var u2 model.User
	database.DB.Find(&u2, "user_id=?", follower)
	followNow := u2.FollowAmount
	database.DB.Model(&u2).Update("follow_amount", followNow+1)

	c.JSON(200, gin.H{
		"msg": "关注成功！",
	})
}

func UnFollow(c *gin.Context) {
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
	var f model.Follow
	res := database.DB.Delete(&f, "up=? AND follower=?", up, follower)
	if res.Error != nil {
		fmt.Println(res.Error)
		return
	}

	var u1 model.User
	database.DB.Find(&u1, "user_id=?", up)
	followerNow := u1.FollowerAmount
	database.DB.Model(&u1).Update("follower_amount", followerNow-1)

	var u2 model.User
	database.DB.Find(&u2, "user_id=?", follower)
	followNow := u2.FollowAmount
	database.DB.Model(&u2).Update("follow_amount", followNow-1)

	c.JSON(200, gin.H{
		"msg": "取消关注成功！",
	})
}
