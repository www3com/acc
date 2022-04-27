-- 模板账本
create table tpl_ledger
(
    id          bigint       not null constraint pk_tpl_ledger primary key,
    name        varchar(80)  not null,
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



-- 模板字典
create table tpl_dict
(
    id          bigint      not null  constraint pk_tpl_dict  primary key,
    ledger_id   int         not null,
    type        smallint    not null,
    name        varchar(80) not null,
    remark      varchar(200),
    sort_number int         not null,
    create_time bigint      not null
);

comment on table tpl_dict is '模板字典';

create index idx_tpl_dict_ledger_id
    on tpl_dict (ledger_id);

alter table tpl_dict owner to acc;

INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (1, 1, 1, '过年', '项目', 1, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (2, 1, 1, '装修', '项目', 2, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (3, 1, 1, '旅游', '项目', 3, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (4, 1, 1, '钓鱼', '项目', 4, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (5, 1, 2, '本人', '成员', 1, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (6, 1, 2, '老婆/老公', '成员', 2, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (7, 1, 2, '子女', '成员', 3, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (8, 1, 2, '父母', '成员', 4, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (9, 1, 2, '家庭公用', '成员', 5, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (10, 1, 3, '餐厅', '商家', 1, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (11, 1, 3, '商场', '商家', 2, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (12, 1, 3, '超市', '商家', 3, 1650986269537);
INSERT INTO public.tpl_dict (id, ledger_id, type, name, remark, sort_number, create_time) VALUES (13, 1, 3, '其他', '商家', 4, 1650986269537);



-- 账本
create table acc_ledger
(
    id             bigserial      not null constraint pk_acc_ledger primary key,
    tpl_ledger_id  int            not null,
    name           varchar(80)    not null,
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



-- 字典
create table acc_dict
(
    id          bigserial      not null  constraint pk_acc_dict  primary key,
    ledger_id   bigint         not null,
    type        smallint       not null,
    name        varchar(80)    not null,
    remark      varchar(200),
    hide_flag   smallint       not null,
    sort_number smallint       not null,
    create_time bigint         not null,
    update_time bigint         not null
);

comment on table acc_dict is '字典';
create index idx_acc_dict_ledger_id on acc_dict (ledger_id);

alter table acc_dict owner to acc;



-- 交易
create table acc_transaction
(
    id                      bigint         not null  constraint pk_acc_transaction  primary key,
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

comment on table acc_dict is '交易';
create index idx_acc_transaction_ledger_id on acc_transaction (ledger_id);

alter table acc_transaction owner to acc;