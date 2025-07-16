package api

import (
	"github.com/gin-gonic/gin"
	"Berra/internal/dto/request/group"
	"Berra/internal/service/gorm"
	"Berra/pkg/constants"
	"Berra/pkg/zlog"
	"net/http"
)

// 1.CreateGroup 创建群聊
func CreateGroup(c *gin.Context) {
	var createGroupReq request.CreateGroupRequest
	if err := c.BindJSON(&createGroupReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.CreateGroup(createGroupReq)
	JsonBack(c, message, ret, nil)
}

// 2.LoadMyGroup 获取我创建的群聊
func LoadMyGroup(c *gin.Context) {
	var loadMyGroupReq request.OwnlistRequest
	if err := c.BindJSON(&loadMyGroupReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupList, ret := gorm.GroupInfoService.LoadMyGroup(loadMyGroupReq.OwnerId)
	JsonBack(c, message, ret, groupList)
}

// 3.CheckGroupAddMode 检查群聊加群方式
func CheckGroupAddMode(c *gin.Context) {
	var req request.CheckGroupAddModeRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, addMode, ret := gorm.GroupInfoService.CheckGroupAddMode(req.GroupId)
	JsonBack(c, message, ret, addMode)
}

// 4.EnterGroupDirectly 直接进群
func EnterGroupDirectly(c *gin.Context) {
	var req request.EnterGroupDirectlyRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.EnterGroupDirectly(req.OwnerId, req.ContactId)
	JsonBack(c, message, ret, nil)
}

// 5.LeaveGroup 退群
func LeaveGroup(c *gin.Context) {
	var req request.LeaveGroupRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.LeaveGroup(req.UserId, req.GroupId)
	JsonBack(c, message, ret, nil)
}

// 6.DismissGroup 解散群聊
func DismissGroup(c *gin.Context) {
	var req request.DismissGroupRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.DismissGroup(req.OwnerId, req.GroupId)
	JsonBack(c, message, ret, nil)
}

// 7.GetGroupInfo 获取群聊详情
func GetGroupInfo(c *gin.Context) {
	var req request.GetGroupInfoRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupInfo, ret := gorm.GroupInfoService.GetGroupInfo(req.GroupId)
	JsonBack(c, message, ret, groupInfo)
}

// 8.GetGroupInfoList 获取群聊列表 - 管理员
func GetGroupInfoList(c *gin.Context) {
	message, groupList, ret := gorm.GroupInfoService.GetGroupInfoList()
	JsonBack(c, message, ret, groupList)
}

// 9.DeleteGroups 删除列表中群聊 - 管理员
func DeleteGroups(c *gin.Context) {
	var req request.DeleteGroupsRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.DeleteGroups(req.UuidList)
	JsonBack(c, message, ret, nil)
}

// 10.SetGroupsStatus 设置群聊是否启用
func SetGroupsStatus(c *gin.Context) {
	var req request.SetGroupsStatusRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.SetGroupsStatus(req.UuidList, req.Status)
	JsonBack(c, message, ret, nil)
}

// 11.UpdateGroupInfo 更新群聊消息
func UpdateGroupInfo(c *gin.Context) {
	var req request.UpdateGroupInfoRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.UpdateGroupInfo(req)
	JsonBack(c, message, ret, nil)
}

// 12.GetGroupMemberList 获取群聊成员列表
func GetGroupMemberList(c *gin.Context) {
	var req request.GetGroupMemberListRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupMemberList, ret := gorm.GroupInfoService.GetGroupMemberList(req.GroupId)
	JsonBack(c, message, ret, groupMemberList)
}

// 13.RemoveGroupMembers 移除群聊成员
func RemoveGroupMembers(c *gin.Context) {
	var req request.RemoveGroupMembersRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.RemoveGroupMembers(req)
	JsonBack(c, message, ret, nil)
}
