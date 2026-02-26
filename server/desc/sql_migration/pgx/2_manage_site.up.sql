CREATE TABLE IF NOT EXISTS "manage_site" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    site_name varchar(100) NOT NULL,
    site_logo varchar(500) NOT NULL DEFAULT '',
    price_currency varchar(20) NOT NULL DEFAULT '',
    score_stat_type varchar(20) NOT NULL DEFAULT '',
    status varchar(1) NOT NULL DEFAULT '1',
    remark text NOT NULL DEFAULT '',
    db_host varchar(255) NOT NULL DEFAULT '',
    db_username varchar(100) NOT NULL DEFAULT '',
    db_password varchar(255) NOT NULL DEFAULT '',
    db_port bigint NOT NULL DEFAULT 3306,
    db_name varchar(100) NOT NULL DEFAULT '',
    jwt_secret varchar(255) NOT NULL DEFAULT '',
    admin_url varchar(500) NOT NULL DEFAULT '',
    admin_username varchar(50) NOT NULL DEFAULT '',
    contact_name varchar(50) NOT NULL DEFAULT '',
    contact_phone varchar(30) NOT NULL DEFAULT '',
    contact_email varchar(100) NOT NULL DEFAULT '',
    remaining_score decimal(20,2) NOT NULL DEFAULT 0,
    total_topup decimal(20,2) NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_site_update_time BEFORE UPDATE ON "manage_site"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

CREATE TABLE IF NOT EXISTS "manage_site_topup" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    site_uuid varchar(36) NOT NULL,
    score decimal(20,2) NOT NULL,
    topup_method varchar(20) NOT NULL DEFAULT '',
    topup_status varchar(20) NOT NULL DEFAULT '',
    remark text NOT NULL DEFAULT '',
    site_name varchar(100) NOT NULL DEFAULT '',
    price_currency varchar(20) NOT NULL DEFAULT '',
    operator_uuid varchar(36) NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_site_topup_update_time BEFORE UPDATE ON "manage_site_topup"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

CREATE TABLE IF NOT EXISTS "manage_site_deduction" (
    id serial NOT NULL,
    uuid varchar(36) NOT NULL UNIQUE,
    create_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    site_uuid varchar(36) NOT NULL,
    score decimal(20,2) NOT NULL,
    deduction_method varchar(20) NOT NULL DEFAULT '',
    deduction_status varchar(20) NOT NULL DEFAULT '',
    remark text NOT NULL DEFAULT '',
    site_name varchar(100) NOT NULL DEFAULT '',
    price_currency varchar(20) NOT NULL DEFAULT '',
    operator_uuid varchar(36) NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TRIGGER update_manage_site_deduction_update_time BEFORE UPDATE ON "manage_site_deduction"
FOR EACH ROW EXECUTE FUNCTION update_update_time_column();

-- Menu: 站点管理 (parent)
INSERT INTO "manage_menu" (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000001','2025-01-01 00:00:00','2025-01-01 00:00:00','1','','1','站点管理',0,'',3,'site','/site','layout.base','mdi:web','1','route.site',0,'',0,0,'[]','[]',0,'');

INSERT INTO "manage_menu" (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000002','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000001','2','站点列表',0,'',1,'site_list','/site/list','view.site_list','mdi:format-list-bulleted','1','route.site_list',0,'',0,0,'[]','[{"code":"v1:manage:site:list","desc":"站点列表"}]',0,'');

INSERT INTO "manage_menu" (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000003','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000001','2','上分记录',0,'',2,'site_topup','/site/topup','view.site_topup','mdi:arrow-up-bold-circle','1','route.site_topup',0,'',0,0,'[]','[{"code":"v1:manage:siteTopup:list","desc":"上分列表"}]',0,'');

INSERT INTO "manage_menu" (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000004','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000001','2','下分记录',0,'',3,'site_deduction','/site/deduction','view.site_deduction','mdi:arrow-down-bold-circle','1','route.site_deduction',0,'',0,0,'[]','[{"code":"v1:manage:siteDeduction:list","desc":"下分列表"}]',0,'');

INSERT INTO "manage_menu" (uuid, create_time, update_time, status, parent_uuid, menu_type, menu_name, hide_in_menu, active_menu, "order", route_name, route_path, component, icon, icon_type, i18n_key, keep_alive, href, multi_tab, fixed_index_in_tab, query, permissions, constant, button_code)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000005','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','新增站点',0,'',0,'','','','','1','button.site.add',0,'',0,0,'[]','[{"code":"v1:manage:site:add","desc":"新增站点"},{"code":"v1:manage:site:getAll","desc":"获取所有站点"}]',0,'v1:manage:site:add'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000006','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','编辑站点',0,'',1,'','','','','1','button.site.edit',0,'',0,0,'[]','[{"code":"v1:manage:site:edit","desc":"编辑站点"}]',0,'v1:manage:site:edit'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000007','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','删除站点',0,'',2,'','','','','1','button.site.delete',0,'',0,0,'[]','[{"code":"v1:manage:site:delete","desc":"删除站点"}]',0,'v1:manage:site:delete'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000008','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000003','3','添加上分',0,'',0,'','','','','1','button.siteTopup.add',0,'',0,0,'[]','[{"code":"v1:manage:siteTopup:add","desc":"添加上分"},{"code":"v1:manage:site:getAll","desc":"获取所有站点"}]',0,'v1:manage:siteTopup:add'),
    ('a0b1c2d3-e4f5-6789-abcd-000000000009','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000004','3','添加下分',0,'',0,'','','','','1','button.siteDeduction.add',0,'',0,0,'[]','[{"code":"v1:manage:siteDeduction:add","desc":"添加下分"},{"code":"v1:manage:site:getAll","desc":"获取所有站点"}]',0,'v1:manage:siteDeduction:add');

INSERT INTO "manage_role_menu" (uuid, create_time, update_time, role_uuid, menu_uuid, is_home)
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
