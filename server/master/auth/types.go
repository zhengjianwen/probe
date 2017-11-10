package auth

type IsMemberRequest struct {
	UserID int64
	OrgID  int64
}

type IsMemberResponse struct {
	Is      bool
	Message string
}

type Node struct {
	ID        uint64 `json:"id" xorm:"pk autoincr 'id'"`
	Name      string `json:"name" xorm:"name"`
	Tips      string `json:"tips" xorm:"tips"`
	CompanyID uint64 `json:"orgId" xorm:"company_id"`
	LevelID   int    `json:"levelId" xorm:"level_id"`
	ParentID  uint64 `json:"parentId" xorm:"parent_id"`
	Path      string `json:"path" xorm:"path"`
	Type      int    `json:"type" xorm:"type"`
}

type CookieData struct {
	UserId   int64
	Username string
}

type HasPermissionRequest struct {
	UserId    int64
	NodeId    int64
	Operation string
}

type HasPermissionResponse struct {
	Message string
	Has     bool
}
