-- 模板账本
create table tpl_ledger
(
    id          bigint       not null constraint pk_tpl_ledger primary key,
    name        varchar(80)  not null,
    icon        varchar(100) not null,
    remark      varchar(200),
    create_time bigint       not null
);

comment on table tpl_ledger is '模板账本';
alter table tpl_ledger owner to acc;

INSERT INTO tpl_ledger (id, name, remark, create_time) VALUES (1, '标准账本', null, 1650955710013);



-- 模板账户
create table tpl_account
(
    id          bigint             not null constraint pk_tpl_account primary key,
    ledger_id   int                not null,
    type        smallint           not null,
    name        varchar(80)        not null,
    parent_id   int                not null,
    icon        varchar(50),
    currency_id int                not null,
    in_asset    smallint default 1 not null,
    sort_number int                not null,
    remark      varchar(200)

);

comment on table tpl_account is '模板账户';
create index idx_tpl_account_ledger_id on tpl_account (ledger_id);

alter table tpl_account owner to acc;


INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (1, 1, 1, '现金账户', 0, null, 1, 1, 1, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (15, 1, 2, '应付款项', 14, null, 1, 1, 1, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (13, 1, 2, '信用卡', 12, null, 1, 1, 1, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (6, 1, 1, '基金', 3, null, 1, 1, 3, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (7, 1, 1, '虚拟账户', 0, null, 1, 1, 3, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (9, 1, 1, '微信', 7, null, 1, 1, 2, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (3, 1, 1, '金融账户', 0, null, 1, 1, 2, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (8, 1, 1, '支付宝', 7, null, 1, 1, 1, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (19, 1, 1, '投资账户', 0, null, 1, 1, 7, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (17, 1, 1, '应收款项', 16, null, 1, 1, 1, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (12, 1, 2, '信用账户', 0, null, 1, 1, 4, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (18, 1, 1, '公司报销', 16, null, 1, 1, 2, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (11, 1, 1, '饭卡', 7, null, 1, 1, 4, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (2, 1, 1, '现金', 1, null, 1, 1, 1, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (10, 1, 1, '公交卡', 7, null, 1, 1, 3, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (5, 1, 1, '股票', 3, null, 1, 1, 2, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (14, 1, 2, '负债账户', 0, null, 1, 1, 5, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (16, 1, 1, '债权账户', 0, null, 1, 1, 6, null);
INSERT INTO public.tpl_account (id, ledger_id, type, name, parent_id, icon, currency_id, in_asset, sort_number, remark) VALUES (4, 1, 1, '银行卡', 3, null, 1, 1, 1, null);

-- 项目
create table tpl_project
(
    id          bigint      not null  constraint pk_tpl_project  primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    sort_number int         not null,
    create_time bigint      not null
);

INSERT INTO public.tpl_project (id, ledger_id, name, remark, sort_number, create_time) VALUES (1, 1, '过年', '', 1, 1650986269537);
INSERT INTO public.tpl_project (id, ledger_id, name, remark, sort_number, create_time) VALUES (2, 1, '装修', '', 2, 1650986269537);
INSERT INTO public.tpl_project (id, ledger_id, name, remark, sort_number, create_time) VALUES (3, 1, '旅游', '', 3, 1650986269537);
INSERT INTO public.tpl_project (id, ledger_id, name, remark, sort_number, create_time) VALUES (4, 1, '钓鱼', '', 4, 1650986269537);


-- 家庭成员
create table tpl_member
(
    id          bigint      not null  constraint pk_tpl_member  primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    sort_number int         not null,
    create_time bigint      not null
);

INSERT INTO public.tpl_member (id, ledger_id, name, remark, sort_number, create_time) VALUES (1, 1, '本人', '', 1, 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, sort_number, create_time) VALUES (2, 1, '老婆/老公', '', 2, 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, sort_number, create_time) VALUES (3, 1, '子女', '', 3, 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, sort_number, create_time) VALUES (4, 1, '父母', '', 4, 1650986269537);
INSERT INTO public.tpl_member (id, ledger_id, name, remark, sort_number, create_time) VALUES (5, 1, '家庭公用', '', 5, 1650986269537);


-- 供应商
create table tpl_supplier
(
    id          bigint      not null  constraint pk_tpl_supplier  primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    sort_number int         not null,
    create_time bigint      not null
);

INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, sort_number, create_time) VALUES (1, 1, '餐厅', '', 1, 1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, sort_number, create_time) VALUES (2, 1, '商场', '', 2, 1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, sort_number, create_time) VALUES (3, 1, '超市', '', 3, 1650986269537);
INSERT INTO public.tpl_supplier (id, ledger_id, name, remark, sort_number, create_time) VALUES (4, 1, '其他', '', 4, 1650986269537);



-- 账本
create table acc_ledger
(
    id             bigserial      not null constraint pk_acc_ledger primary key,
    tpl_ledger_id  int            not null,
    name           varchar(80)    not null,
    icon           varchar(100)   not null ,
    owner_id       bigint         not null,
    remark         varchar(200),
    sort_number    int            not null,
    create_time    bigint         not null,
    update_time    bigint         not null
);

comment on table acc_ledger is '账本';
alter table acc_ledger owner to acc;



-- 账户
create table acc_account
(
    id            bigserial          not null constraint pk_acc_account primary key,
    ledger_id     bigint             not null,
    type          smallint           not null,
    name          varchar(80)        not null,
    debit         decimal(18,2)      not null,
    credit        decimal(18,2)      not null,
    balance       decimal(18,2)      not null,
    parent_id     bigint             not null,
    icon          varchar(50),
    currency_id   int                not null,
    in_asset      smallint default 1 not null,
    sort_number   int                not null,
    remark        varchar(200),
    create_time    bigint         not null,
    update_time    bigint         not null
);

comment on table acc_account is '账户';
comment on column acc_account.type is '账户类型: 1-资产, 2-负债, 3-权益, 4-收入, 5-支出';
create index idx_acc_account_ledger_id on acc_account (ledger_id);

alter table acc_account owner to acc;


-- 项目
create table acc_project
(
    id          bigserial   not null  constraint pk_acc_project  primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    hide_flag   smallint    not null,
    sort_number int         not null,
    create_time bigint      not null,
    update_time bigint      not null
);

create index idx_acc_project_ledger_id on acc_project (ledger_id);

alter table acc_project owner to acc;


-- 家庭成员
create table acc_member
(
    id          bigserial   not null  constraint pk_acc_member  primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    hide_flag   smallint    not null,
    sort_number int         not null,
    create_time bigint      not null,
    update_time bigint      not null
);

create index idx_acc_member_ledger_id on acc_project (ledger_id);

alter table acc_member owner to acc;

-- 供应商
create table acc_supplier
(
    id          bigserial   not null  constraint pk_acc_supplier  primary key,
    ledger_id   int         not null,
    name        varchar(80) not null,
    remark      varchar(200),
    hide_flag   smallint    not null,
    sort_number int         not null,
    create_time bigint      not null,
    update_time bigint      not null
);

create index idx_acc_supplier_ledger_id on acc_supplier (ledger_id);

alter table acc_supplier owner to acc;


-- 交易
create table acc_transaction
(
    id                      bigserial      not null  constraint pk_acc_transaction  primary key,
    ledger_id               bigint         not null,
    type                    smallint       not null,
    debit_account_id        bigint         not null,
    credit_account_id       bigint         not null,
    amount                  decimal(18,2)  not null,
    remark                  varchar(200),
    project_id              bigint         not null,
    hide_flag               smallint       not null,
    create_time             bigint         not null,
    update_time             bigint         not null
);

comment on table acc_transaction is '交易';
create index idx_acc_transaction_ledger_id on acc_transaction (ledger_id);

alter table acc_transaction owner to acc;


create table upm_user
(
    id          bigserial    not null  constraint pk_upm_user  primary key,
    nickname    varchar(80)  not null,
    username    varchar(59)  not null constraint uni_upm_user_username unique,
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

alter table upm_user owner to acc;
