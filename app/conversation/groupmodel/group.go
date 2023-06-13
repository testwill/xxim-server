package groupmodel

import (
	"context"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/cherish-chat/xxim-server/common/xcache"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"os"
)

// Group 群组 数据库模型
type Group struct {
	//GroupId 群id 主键
	GroupId int `json:"groupId" gorm:"column:groupId;type:bigint(20);not null;primaryKey"`
	//GroupName 群名称
	GroupName string `json:"groupName" gorm:"column:groupName;type:varchar(32);not null;index;"`
	//GroupAvatar 群头像
	GroupAvatar string `json:"groupAvatar" gorm:"column:groupAvatar;type:varchar(255);not null"`
	//GroupInfo 自定义群信息
	//如果没有查询需要，不要SELECT这个字段
	GroupInfo string `json:"groupInfo" gorm:"column:groupInfo;type:text;not null;"`

	//OwnerUserId 群主id
	OwnerUserId string `json:"ownerUserId" gorm:"column:ownerUserId;type:char(32);not null"`
	//ManagerUserIds 管理员id列表 逗号分隔
	//如果没有查询需要，不要SELECT这个字段，因为这个字段可能会很大，群管理员上限是1900人，因为65535/33=1985.90
	ManagerUserIds string `json:"managerUserIds" gorm:"column:managerUserIds;type:text;not null"`
	//CreatedAt 创建时间 13位时间戳
	CreateTime int64 `json:"createTime" gorm:"column:createTime;type:bigint;not null;index;"`
	//UpdatedAt 更新时间 13位时间戳
	UpdateTime int64 `json:"updateTime" gorm:"column:updateTime;type:bigint;not null"`
	//DismissTime 解散时间 13位时间戳
	DismissTime int64 `json:"dismissTime" gorm:"column:dismissTime;type:bigint;not null;default:0;index;"`
	//MemberCount 成员数量
	MemberCount int `json:"memberCount" gorm:"column:memberCount;type:int;not null;default:0;"`

	//RemarkForAdmin 管理员设置的备注
	RemarkForAdmin string `json:"remarkForAdmin" gorm:"column:remarkForAdmin;type:varchar(32);not null;default:'';"`
}

// TableName 表名
func (m *Group) TableName() string {
	return "group"
}

type xGroupModel struct {
	db *gorm.DB
	rc *redis.Redis
}

var GroupModel *xGroupModel

func InitGroupModel(db *gorm.DB, rc *redis.Redis, minGroupId int) {
	GroupModel = &xGroupModel{
		db: db,
		rc: rc,
	}
	err := db.AutoMigrate(&Group{})
	if err != nil {
		logx.Errorf("db.AutoMigrate(&Group{}) error: %v", err)
		os.Exit(1)
		return
	}
	// 查询最大的群组id
	var group Group
	err = db.Model(&Group{}).Order("groupId desc").First(&group).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 说明没有数据
			group.GroupId = minGroupId
		} else {
			logx.Errorf(`db.Model(&Group{}).Order("groupId desc").First(&group) error: %v`, err)
			os.Exit(1)
			return
		}
	}
	// 保存到redis
	err = rc.Set(xcache.IncrKeyGroupId, utils.AnyString(group.GroupId))
	if err != nil {
		logx.Errorf(`rc.Set(xcache.IncrKeyGroupId, utils.AnyString(group.GroupId)) error: %v`, err)
		os.Exit(1)
		return
	}
}

func (x *xGroupModel) GenerateGroupId() int {
	// 从redis中获取
	groupId, err := x.rc.Incr(xcache.IncrKeyGroupId)
	if err != nil {
		logx.Errorf("x.rc.Incr(xcache.IncrKeyGroupId) error: %v", err)
		return 0
	}
	return int(groupId)
}

func (x *xGroupModel) Insert(ctx context.Context, tx gorm.DB, group *Group) error {
	if group.GroupId == 0 {
		group.GroupId = x.GenerateGroupId()
	}
	return tx.WithContext(ctx).Create(group).Error
}