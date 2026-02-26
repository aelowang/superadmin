package manage_site_deduction

import (
	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero/core/stores/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ManageSiteDeductionModel = (*customManageSiteDeductionModel)(nil)

type (
	// ManageSiteDeductionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customManageSiteDeductionModel.
	ManageSiteDeductionModel interface {
		manageSiteDeductionModel
	}

	customManageSiteDeductionModel struct {
		*defaultManageSiteDeductionModel
	}
)

// NewManageSiteDeductionModel returns a model for the database table.
func NewManageSiteDeductionModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) ManageSiteDeductionModel {
	return &customManageSiteDeductionModel{
		defaultManageSiteDeductionModel: newManageSiteDeductionModel(conn, op...),
	}
}
