package manage_site

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageSiteModel = (*customManageSiteModel)(nil)

type (
	// ManageSiteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageSiteModel.
	ManageSiteModel interface {
		manageSiteModel
	}

	customManageSiteModel struct {
		*defaultManageSiteModel
	}
)

// NewManageSiteModel returns a model for the database table.
func NewManageSiteModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageSiteModel {
	return &customManageSiteModel{
		defaultManageSiteModel: newManageSiteModel(conn, op...),
	}
}
