-- 模板账本
create table tpl_ledger
(
    id          bigint      not null,
    name        varchar(80) not null,
    icon        varchar(100),
    remark      varchar(200),
    create_time bigint      not null,
    CONSTRAINT pk_tpl_ledger PRIMARY KEY (id)
);

comment on table tpl_ledger is '模板账本';
alter table tpl_ledger
    owner to acc;

INSERT INTO tpl_ledger (id, name, remark, create_time)
VALUES (1, '标准账本', null, 1650955710013);


-- 模板账户
create table tpl_account
(
    id          bigserial   not null,
    ledger_id   int         not null,
    type        smallint    not null,
    name        varchar(80) not null,
    code        varchar(30) not null,
    level       int         not null,
    icon        varchar(50),
    currency_id int         not null,
    remark      varchar(200),
    CONSTRAINT pk_tpl_account PRIMARY KEY ("id"),
    CONSTRAINT uk_tpl_account_code UNIQUE ("ledger_id", "code")

);

comment on table tpl_account is '模板账户';
create index idx_tpl_account_ledger_id on tpl_account (ledger_id);

alter table tpl_account
    owner to acc;

INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '现金账户', '1001', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '现金', '1001.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '金融账户', '1002', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '银行卡', '1002.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '股票', '1002.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '基金', '1002.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '虚拟账户', '1003', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '支付宝', '1003.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '微信', '1003.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '公交卡', '1003.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '饭卡', '1003.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '债权账户', '1004', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '应收款项', '1004.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '公司报销', '1004.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 1, '投资账户', '1005', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 2, '信用账户', '1006', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 2, '信用卡', '1006.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 2, '负债账户', '1007', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 2, '应付款项', '1007.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '职业收入', '1009', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '工资', '1009.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '经营', '1009.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '利息', '1009.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '兼职', '1009.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '投资', '1009.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '奖金', '1009.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '加班', '1009.07', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '其他收入', '1010', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '礼金', '1010.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '抢红包', '1010.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '意外来钱', '1010.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '家里给钱', '1010.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '中奖', '1010.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '退税', '1010.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '餐饮', '1012', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '三餐', '1012.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '早餐', '1012.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '中餐', '1012.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '晚餐', '1012.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '买菜', '1012.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '水果', '1012.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '零食', '1012.07', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '加餐', '1012.08', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '下午茶', '1012.09', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '烟酒饮品', '1012.10', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '粮油调料', '1012.11', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '交通', '1013', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '公交地铁', '1013.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '打车', '1013.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '私家车', '1013.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '飞机火车', '1013.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '共享单车', '1013.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '购物', '1014', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '日用品', '1014.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '衣帽鞋包', '1014.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '护肤美妆', '1014.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '饰品', '1014.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '数码', '1014.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '电器', '1014.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '家装', '1014.07', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '居家', '1015', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '水电煤', '1015.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '话费', '1015.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '网费', '1015.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '房租', '1015.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '物业', '1015.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '维修', '1015.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '人情', '1016', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '送礼', '1016.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '请客', '1016.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '孝心', '1016.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '亲密付', '1016.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '发红包', '1016.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '借出', '1016.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '还钱', '1016.07', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '医疗', '1017', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '药品', '1017.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '保健', '1017.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '治疗', '1017.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '美容', '1017.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '娱乐', '1018', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '休闲', '1018.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '约会', '1018.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '聚会', '1018.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '游戏', '1018.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '健身', '1018.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '学习', '1019', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '书籍', '1019.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '培训', '1019.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '网课', '1019.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '金融', '1020', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '房贷', '1020.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '车贷', '1020.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '购物分期', '1020.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '手续费', '1020.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '保险', '1020.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '养卡', '1020.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '其他', '1021', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '旅游', '1021.01', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '装修', '1021.02', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '宝宝', '1021.03', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '生意', '1021.04', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '宠物', '1021.05', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '坏账', '1021.06', 2, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 3, '余额调整', '9998', 1, NULL, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, level, icon, currency_id, remark)
VALUES (1, 4, '余额调整', '9999', 1, NULL, 1, NULL);

-- 项目
create table tpl_project
(
    id          bigint      not null,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    create_time bigint      not null,
    CONSTRAINT pk_tpl_project PRIMARY KEY (id)
);

INSERT INTO public.tpl_project (id, ledger_id, name, remark, create_time)
VALUES (1, 1, '过年', '', 1650986269537);
INSERT INTO public.tpl_project (id, ledger_id, name, remark, create_time)
VALUES (2, 1, '装修', '', 1650986269537);
INSERT INTO public.tpl_project (id, ledger_id, name, remark, create_time)
VALUES (3, 1, '旅游', '', 1650986269537);
INSERT INTO public.tpl_project (id, ledger_id, name, remark, create_time)
VALUES (4, 1, '钓鱼', '', 1650986269537);


-- 家庭成员
create table tpl_member
(
    id          bigint      not null,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    create_time bigint      not null,
    CONSTRAINT pk_tpl_member PRIMARY KEY (id)
);

INSERT INTO public.tpl_member (id, ledger_id, name, remark, create_time)
VALUES (1, 1, '本人', '', 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, create_time)
VALUES (2, 1, '老婆/老公', '', 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, create_time)
VALUES (3, 1, '子女', '', 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, create_time)
VALUES (4, 1, '父母', '', 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, create_time)
VALUES (5, 1, '家庭公用', '', 1650986269537);


-- 供应商
create table tpl_supplier
(
    id          bigint      not null,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    create_time bigint      not null,
    CONSTRAINT pk_tpl_supplier PRIMARY KEY (id)
);

INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (1, 1, '餐厅', '', 1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (2, 1, '商场', '', 1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (3, 1, '超市', '', 1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (4, 1, '其他', '', 1650986269537);


-- 账本
create table acc_ledger
(
    id          bigserial    not null,
    t_ledger_id int          not null,
    name        varchar(80)  not null,
    icon        varchar(100) not null,
    owner_id    bigint       not null,
    remark      varchar(200),
    selected    int          not null,
    create_time bigint       not null,
    update_time bigint       not null,
    CONSTRAINT pk_acc_ledger PRIMARY KEY (id)
);

comment on table acc_ledger is '账本';
alter table acc_ledger
    owner to acc;


-- 账户
create table acc_account
(
    id          bigserial      not null,
    ledger_id   bigint         not null,
    type        smallint       not null,
    name        varchar(80)    not null,
    code        varchar(30)    not null,
    level       int            not null,
    debit       decimal(18, 2) not null,
    credit      decimal(18, 2) not null,
    balance     decimal(18, 2) not null,
    icon        varchar(50),
    currency_id int            not null,
    remark      varchar(200),
    create_time bigint         not null,
    update_time bigint         not null,
    CONSTRAINT pk_acc_account PRIMARY KEY (id),
    CONSTRAINT uk_acc_account_code UNIQUE (ledger_id, code)
);

comment on table acc_account is '账户';
comment on column acc_account.type is '账户类型: 1-资产, 2-负债, 3-收入, 4-支出';
create index idx_acc_account_ledger_id on acc_account (ledger_id);

alter table acc_account
    owner to acc;


-- 项目
create table acc_project
(
    id          bigserial   not null,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    is_show     smallint    not null,
    create_time bigint      not null,
    update_time bigint      not null,
    CONSTRAINT pk_acc_project PRIMARY KEY (id)
);

create index idx_acc_project_ledger_id on acc_project (ledger_id);

alter table acc_project
    owner to acc;


-- 家庭成员
create table acc_member
(
    id          bigserial   not null,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    is_show     smallint    not null,
    create_time bigint      not null,
    update_time bigint      not null,
    CONSTRAINT pk_acc_member PRIMARY KEY (id)
);

create index idx_acc_member_ledger_id on acc_project (ledger_id);

alter table acc_member
    owner to acc;

-- 供应商
create table acc_supplier
(
    id          bigserial   not null,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    is_show     smallint    not null,
    create_time bigint      not null,
    update_time bigint      not null,
    CONSTRAINT pk_acc_supplier PRIMARY KEY (id)
);

create index idx_acc_supplier_ledger_id on acc_supplier (ledger_id);

alter table acc_supplier
    owner to acc;


-- 交易
create table acc_transaction
(
    id                bigserial      not null,
    ledger_id         bigint         not null,
    type              smallint       not null,
    debit_account_id  bigint         not null,
    credit_account_id bigint         not null,
    amount            decimal(18, 2) not null,
    remark            varchar(200),
    project_id        bigint         not null,
    is_show           smallint       not null,
    create_time       bigint         not null,
    update_time       bigint         not null,
    CONSTRAINT pk_acc_transaction PRIMARY KEY (id)
);

comment on table acc_transaction is '交易';
create index idx_acc_transaction_ledger_id on acc_transaction (ledger_id);

alter table acc_transaction
    owner to acc;


create table upm_user
(
    id          bigserial    not null,
    nickname    varchar(80)  not null,
    username    varchar(59)  not null,
    email       varchar(200) not null,
    password    varchar(40)  not null,
    agreement   integer,
    state       smallint     not null,
    create_time bigint,
    update_time bigint,
    CONSTRAINT pk_upm_user PRIMARY KEY (id),
    CONSTRAINT uni_upm_user_username UNIQUE (username)
);

comment on table upm_user is '用户';

comment on column upm_user.id is '主键';

comment on column upm_user.nickname is '姓名';

comment on column upm_user.username is '用户名';

comment on column upm_user.email is '电子邮箱';

comment on column upm_user.password is '密码';

comment on column upm_user.agreement is '协议';

comment on column upm_user.create_time is '创建时间';

comment on column upm_user.update_time is '更新时间';

alter table upm_user
    owner to acc;
