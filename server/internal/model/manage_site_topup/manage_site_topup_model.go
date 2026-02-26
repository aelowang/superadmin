package manage_site_topup

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageSiteTopupModel = (*customManageSiteTopupModel)(nil)

type (
	// ManageSiteTopupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageSiteTopupModel.
	ManageSiteTopupModel interface {
		manageSiteTopupModel
	}

	customManageSiteTopupModel struct {
		*defaultManageSiteTopupModel
	}
)

// NewManageSiteTopupModel returns a model for the database table.
func NewManageSiteTopupModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageSiteTopupModel {
	return &customManageSiteTopupModel{
		defaultManageSiteTopupModel: newManageSiteTopupModel(conn, op...),
	}
}
