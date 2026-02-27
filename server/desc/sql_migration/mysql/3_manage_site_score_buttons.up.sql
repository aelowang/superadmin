-- Button: 站点上分 (under 站点列表)
INSERT INTO `manage_menu` (`uuid`, `create_time`, `update_time`, `status`, `parent_uuid`, `menu_type`, `menu_name`, `hide_in_menu`, `active_menu`, `order`, `route_name`, `route_path`, `component`, `icon`, `icon_type`, `i18n_key`, `keep_alive`, `href`, `multi_tab`, `fixed_index_in_tab`, `query`, `permissions`, `constant`, `button_code`)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000010','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','站点上分',0,'',3,'','','','','1','button.site.topup',0,'',0,0,'[]','[{"code":"v1:manage:site:topup","desc":"站点上分"}]',0,'v1:manage:site:topup');

-- Button: 站点下分 (under 站点列表)
INSERT INTO `manage_menu` (`uuid`, `create_time`, `update_time`, `status`, `parent_uuid`, `menu_type`, `menu_name`, `hide_in_menu`, `active_menu`, `order`, `route_name`, `route_path`, `component`, `icon`, `icon_type`, `i18n_key`, `keep_alive`, `href`, `multi_tab`, `fixed_index_in_tab`, `query`, `permissions`, `constant`, `button_code`)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000011','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','站点下分',0,'',4,'','','','','1','button.site.deduction',0,'',0,0,'[]','[{"code":"v1:manage:site:deduction","desc":"站点下分"}]',0,'v1:manage:site:deduction');

-- Button: 同步远程余额 (under 站点列表)
INSERT INTO `manage_menu` (`uuid`, `create_time`, `update_time`, `status`, `parent_uuid`, `menu_type`, `menu_name`, `hide_in_menu`, `active_menu`, `order`, `route_name`, `route_path`, `component`, `icon`, `icon_type`, `i18n_key`, `keep_alive`, `href`, `multi_tab`, `fixed_index_in_tab`, `query`, `permissions`, `constant`, `button_code`)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-000000000012','2025-01-01 00:00:00','2025-01-01 00:00:00','1','a0b1c2d3-e4f5-6789-abcd-000000000002','3','同步远程余额',0,'',5,'','','','','1','button.site.syncRemoteScore',0,'',0,0,'[]','[{"code":"v1:manage:site:syncRemoteScore","desc":"同步远程余额"}]',0,'v1:manage:site:syncRemoteScore');

-- Assign to super admin role
INSERT INTO `manage_role_menu` (`uuid`, `create_time`, `update_time`, `role_uuid`, `menu_uuid`, `is_home`)
VALUES
    ('a0b1c2d3-e4f5-6789-abcd-100000000010','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000010',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000011','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000011',0),
    ('a0b1c2d3-e4f5-6789-abcd-100000000012','2025-01-01 00:00:00','2025-01-01 00:00:00','1a2b3c4d-5e6f-7a8b-9c0d-1e2f3a4b5c6d','a0b1c2d3-e4f5-6789-abcd-000000000012',0);
