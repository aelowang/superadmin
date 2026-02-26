CREATE TABLE IF NOT EXISTS manage_site (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    site_name TEXT NOT NULL,
    site_logo TEXT NOT NULL DEFAULT '',
    price_currency TEXT NOT NULL DEFAULT '',
    score_stat_type TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT '1',
    remark TEXT NOT NULL DEFAULT '',
    db_host TEXT NOT NULL DEFAULT '',
    db_username TEXT NOT NULL DEFAULT '',
    db_password TEXT NOT NULL DEFAULT '',
    db_port INTEGER NOT NULL DEFAULT 3306,
    db_name TEXT NOT NULL DEFAULT '',
    jwt_secret TEXT NOT NULL DEFAULT '',
    admin_url TEXT NOT NULL DEFAULT '',
    admin_username TEXT NOT NULL DEFAULT '',
    contact_name TEXT NOT NULL DEFAULT '',
    contact_phone TEXT NOT NULL DEFAULT '',
    contact_email TEXT NOT NULL DEFAULT '',
    remaining_score REAL NOT NULL DEFAULT 0,
    total_topup REAL NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS manage_site_topup (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    site_uuid TEXT NOT NULL,
    score REAL NOT NULL,
    topup_method TEXT NOT NULL DEFAULT '',
    topup_status TEXT NOT NULL DEFAULT '',
    remark TEXT NOT NULL DEFAULT '',
    site_name TEXT NOT NULL DEFAULT '',
    price_currency TEXT NOT NULL DEFAULT '',
    operator_uuid TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS manage_site_deduction (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    site_uuid TEXT NOT NULL,
    score REAL NOT NULL,
    deduction_method TEXT NOT NULL DEFAULT '',
    deduction_status TEXT NOT NULL DEFAULT '',
    remark TEXT NOT NULL DEFAULT '',
    site_name TEXT NOT NULL DEFAULT '',
    price_currency TEXT NOT NULL DEFAULT '',
    operator_uuid TEXT NOT NULL DEFAULT ''
);

-- Menu: 站点管理 (parent)
INSERT INTO manage_menu (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000001','2025-01-01 00:00:00','2025-01-01 00:00:00','1','','1','站点管理',0,'',3,'site','/site','layout.base','mdi:web','1','route.site',0,'',0,0,'[]','[]',0,'');

INSERT INTO manage_menu (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000002','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000001','2','站点列表',0,'',1,'site_list','/site/list','view.site_list','mdi:format-list-bulleted','1','route.site_list',0,'',0,0,'[]','[{"code":"v1:manage:site:list","desc":"站点列表"}]',0,'');

INSERT INTO manage_menu (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000003','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000001','2','上分记录',0,'',2,'site_topup','/site/topup','view.site_topup','mdi:arrow-up-bold-circle','1','route.site_topup',0,'',0,0,'[]','[{"code":"v1:manage:siteTopup:list","desc":"上分列表"}]',0,'');

INSERT INTO manage_menu (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000004','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000001','2','下分记录',0,'',3,'site_deduction','/site/deduction','view.site_deduction','mdi:arrow-down-bold-circle','1','route.site_deduction',0,'',0,0,'[]','[{"code":"v1:manage:siteDeduction:list","desc":"下分列表"}]',0,'');

INSERT INTO manage_menu (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000005','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','新增站点',0,'',0,'','','','','1','button.site.add',0,'',0,0,'[]','[{"code":"v1:manage:site:add","desc":"新增站点"},{"code":"v1:manage:site:getAll","desc":"获取所有站点"}]',0,'v1:manage:site:add'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000006','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','编辑站点',0,'',1,'','','','','1','button.site.edit',0,'',0,0,'[]','[{"code":"v1:manage:site:edit","desc":"编辑站点"}]',0,'v1:manage:site:edit'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000007','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','删除站点',0,'',2,'','','','','1','button.site.delete',0,'',0,0,'[]','[{"code":"v1:manage:site:delete","desc":"删除站点"}]',0,'v1:manage:site:delete'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000008','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000003','3','添加上分',0,'',0,'','','','','1','button.siteTopup.add',0,'',0,0,'[]','[{"code":"v1:manage:siteTopup:add","desc":"添加上分"},{"code":"v1:manage:site:getAll","desc":"获取所有站点"}]',0,'v1:manage:siteTopup:add'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000009','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000004','3','添加下分',0,'',0,'','','','','1','button.siteDeduction.add',0,'',0,0,'[]','[{"code":"v1:manage:siteDeduction:add","desc":"添加下分"},{"code":"v1:manage:site:getAll","desc":"获取所有站点"}]',0,'v1:manage:siteDeduction:add');

INSERT INTO manage_role_menu (uuid, create_time, update_time, role_uuid, menu_uuid, is_home)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-100000000001','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000001',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000002','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000002',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000003','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000003',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000004','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000004',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000005','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000005',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000006','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000006',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000007','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000007',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000008','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000008',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000009','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000009',0);
