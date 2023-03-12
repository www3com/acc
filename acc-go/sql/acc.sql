-- 模板账本
create table tpl_ledger
(
    id          bigint       not null
        constraint pk_tpl_ledger primary key,
    name        varchar(80)  not null,
    icon        varchar(100) ,
    remark      varchar(200),
    create_time bigint       not null
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
    code        varchar(15) not null,
    icon        varchar(50),
    currency_id int         not null,
    leaf        int2        not null,
    remark      varchar(200),
    CONSTRAINT "pk_tpl_account" PRIMARY KEY ("id"),
    CONSTRAINT "uk_tpl_account_code" UNIQUE ("ledger_id", "code")

);

comment on table tpl_account is '模板账户';
create index idx_tpl_account_ledger_id on tpl_account (ledger_id);

alter table tpl_account
    owner to acc;

INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '现金账户', '001', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '现金', '001001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '金融账户', '002', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '银行卡', '002001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '股票', '002002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '基金', '002003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '虚拟账户', '003', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '支付宝', '003001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '微信', '003002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '公交卡', '003003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '饭卡', '003004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '债权账户', '004', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '应收款项', '004001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '公司报销', '004002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 1, '投资账户', '005', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 2, '信用账户', '006', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 2, '信用卡', '006001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 2, '负债账户', '007', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 2, '应付款项', '007001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '余额调整', '008', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '职业收入', '009', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '工资', '009001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '经营', '009002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '利息', '009003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '兼职', '009004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '投资', '009005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '奖金', '009006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '加班', '009007', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '其他收入', '010', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '礼金', '010001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '抢红包', '010002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '意外来钱', '010003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '家里给钱', '010004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '中奖', '010005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 3, '退税', '010006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '余额调整', '011', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '餐饮', '012', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '三餐', '012001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '早餐', '012002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '中餐', '012003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '晚餐', '012004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '买菜', '012005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '水果', '012006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '零食', '012007', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '加餐', '012008', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '下午茶', '012009', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '烟酒饮品', '012010', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '粮油调料', '012011', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '交通', '013', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '公交地铁', '013001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '打车', '013002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '私家车', '013003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '飞机火车', '013004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '共享单车', '013005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '购物', '014', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '日用品', '014001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '衣帽鞋包', '014002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '护肤美妆', '014003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '饰品', '014004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '数码', '014005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '电器', '014006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '家装', '014007', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '居家', '015', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '水电煤', '015001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '话费', '015002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '网费', '015003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '房租', '015004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '物业', '015005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '维修', '015006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '人情', '016', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '送礼', '016001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '请客', '016002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '孝心', '016003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '亲密付', '016004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '发红包', '016005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '借出', '016006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '还钱', '016007', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '医疗', '017', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '药品', '017001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '保健', '017002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '治疗', '017003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '美容', '017004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '娱乐', '018', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '休闲', '018001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '约会', '018002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '聚会', '018003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '游戏', '018004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '健身', '018005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '学习', '019', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '书籍', '019001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '培训', '019002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '网课', '019003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '金融', '020', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '房贷', '020001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '车贷', '020002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '购物分期', '020003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '手续费', '020004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '保险', '020005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '养卡', '020006', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '其他', '021', NULL, 1, 0, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '旅游', '021001', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '装修', '021002', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '宝宝', '021003', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '生意', '021004', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '宠物', '021005', NULL, 1, 1, NULL);
INSERT INTO public.tpl_account (ledger_id, type, name, code, icon, currency_id, leaf, remark)
VALUES (1, 4, '坏账', '021006', '', 1, 1, NULL);


-- 项目
create table tpl_project
(
    id          bigint      not null
        constraint pk_tpl_project primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    create_time bigint      not null
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
    id          bigint      not null
        constraint pk_tpl_member primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    create_time bigint      not null
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
    id          bigint      not null
        constraint pk_tpl_supplier primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    create_time bigint      not null
);

INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (1, 1, '餐厅', '',  1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (2, 1, '商场', '',  1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (3, 1, '超市', '',  1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, create_time)
VALUES (4, 1, '其他', '',  1650986269537);


-- 账本
create table acc_ledger
(
    id            bigserial    not null
        constraint pk_acc_ledger primary key,
    tpl_ledger_id int          not null,
    name          varchar(80)  not null,
    icon          varchar(100) not null,
    owner_id      bigint       not null,
    remark        varchar(200),
    create_time   bigint       not null,
    update_time   bigint       not null
);

comment on table acc_ledger is '账本';
alter table acc_ledger
    owner to acc;


-- 账户
create table acc_account
(
    id          bigserial      not null
        constraint pk_acc_account primary key,
    ledger_id   bigint         not null,
    type        smallint       not null,
    name        varchar(80)    not null,
    debit       decimal(18, 2) not null,
    credit      decimal(18, 2) not null,
    balance     decimal(18, 2) not null,
    icon        varchar(50),
    currency_id int            not null,
    leaf        int2           not null,
    remark      varchar(200),
    create_time bigint         not null,
    update_time bigint         not null
);

comment on table acc_account is '账户';
comment on column acc_account.type is '账户类型: 1-资产, 2-负债, 3-收入, 4-支出';
create index idx_acc_account_ledger_id on acc_account (ledger_id);

alter table acc_account
    owner to acc;


-- 项目
create table acc_project
(
    id          bigserial   not null
        constraint pk_acc_project primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    hide_flag   smallint    not null,
    create_time bigint      not null,
    update_time bigint      not null
);

create index idx_acc_project_ledger_id on acc_project (ledger_id);

alter table acc_project
    owner to acc;


-- 家庭成员
create table acc_member
(
    id          bigserial   not null
        constraint pk_acc_member primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    hide_flag   smallint    not null,
    create_time bigint      not null,
    update_time bigint      not null
);

create index idx_acc_member_ledger_id on acc_project (ledger_id);

alter table acc_member
    owner to acc;

-- 供应商
create table acc_supplier
(
    id          bigserial   not null
        constraint pk_acc_supplier primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    hide_flag   smallint    not null,
    create_time bigint      not null,
    update_time bigint      not null
);

create index idx_acc_supplier_ledger_id on acc_supplier (ledger_id);

alter table acc_supplier
    owner to acc;


-- 交易
create table acc_transaction
(
    id                bigserial      not null
        constraint pk_acc_transaction primary key,
    ledger_id         bigint         not null,
    type              smallint       not null,
    debit_account_id  bigint         not null,
    credit_account_id bigint         not null,
    amount            decimal(18, 2) not null,
    remark            varchar(200),
    project_id        bigint         not null,
    hide_flag         smallint       not null,
    create_time       bigint         not null,
    update_time       bigint         not null
);

comment on table acc_transaction is '交易';
create index idx_acc_transaction_ledger_id on acc_transaction (ledger_id);

alter table acc_transaction
    owner to acc;


create table upm_user
(
    id          bigserial    not null
        constraint pk_upm_user primary key,
    nickname    varchar(80)  not null,
    username    varchar(59)  not null
        constraint uni_upm_user_username unique,
    email       varchar(200) not null,
    password    varchar(40)  not null,
    agreement   integer,
    state       smallint     not null,
    create_time bigint,
    update_time bigint
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
