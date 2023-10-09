create table users
(
    id           serial primary key,
    account      varchar(50) unique not null,
    birthdate    date default null,
    gender       smallint default 1 check (gender >= 1 and gender <= 3),
    password     varchar(256) not null,
    nickname     varchar(50),
    country_code varchar(10),
    address      varchar(100),
    phone_number varchar(20),
    create_at    timestamp(0) with time zone default now() not null,
    update_at    timestamp(0) with time zone default now() not null
);

INSERT INTO users (account,birthdate,gender,"password",nickname,country_code,address,phone_number,create_at,update_at) values
('sean001',NULL,1,'test1234',NULL,NULL,NULL,NULL,'2023-10-07 23:37:22+08','2023-10-09 16:00:19+08'),
('sean002',NULL,1,'test1234','Sean02',NULL,NULL,NULL,'2023-10-08 10:03:07+08','2023-10-08 10:03:07+08'),
('sean003',NULL,1,'test1234','Sean03',NULL,NULL,NULL,'2023-10-08 11:23:03+08','2023-10-08 11:23:03+08'),
('sean004',NULL,1,'test1234','Sean04',NULL,NULL,NULL,'2023-10-08 11:23:41+08','2023-10-08 11:23:41+08'),
('sean005',NULL,1,'test1234','Sean05',NULL,NULL,NULL,'2023-10-08 11:42:34+08','2023-10-08 11:42:34+08'),
('sean006',NULL,1,'test1234','Sean06',NULL,NULL,NULL,'2023-10-08 11:43:38+08','2023-10-08 11:43:38+08'),
('sean007',NULL,1,'test1234','Sean07',NULL,NULL,NULL,'2023-10-08 13:23:02+08','2023-10-08 13:23:02+08');


create table user_hobbies
(
    id        serial primary key,
    user_id   int not null,
    hobby     varchar(50),
    create_at timestamp(0) with time zone default now() not null,
    update_at timestamp(0) with time zone default now() not null
);

create table user_jobs
(
    id        serial primary key,
    user_id   int not null,
    job       varchar(50),
    create_at timestamp(0) with time zone default now() not null,
    update_at timestamp(0) with time zone default now() not null
);

create or replace
function update_update_at()
returns trigger as $$
begin
  new.update_at = now();
  return new;
end;
$$ language plpgsql;

drop trigger if exists update_users_update_at on users;
create trigger update_users_update_at
    before update
    on users
    for each row execute procedure update_update_at();

drop trigger if exists update_user_hobbies_update_at on user_hobbies;

create trigger update_user_hobbies_update_at
    before update
    on user_hobbies
    for each row execute procedure update_update_at();

drop trigger if exists update_user_jobs_update_at on user_jobs;

create trigger update_user_jobs_update_at
    before update
    on user_jobs
    for each row execute procedure update_update_at();

create table if not exists countries (
    id serial primary key,
    code          VARCHAR(4) NOT NULL,
    calling_code  VARCHAR(10) NOT NULL,
    name          VARCHAR(50) NOT NULL,
    official_name VARCHAR(100) NOT NULL,
    name_ch       VARCHAR(50) NOT NULL
);

INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('af','+93','Afghanistan','Islamic Republic of Afghanistan','阿富汗');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sd','+249','Sudan','Republic of the Sudan','苏丹');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bi','+257','Burundi','Republic of Burundi','布隆迪');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mx','+52','Mexico','United Mexican States','墨西哥');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cu','+53','Cuba','Republic of Cuba','古巴');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('rs','+381','Serbia','Republic of Serbia','塞尔维亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mc','+377','Monaco','Principality of Monaco','摩纳哥');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bt','+975','Bhutan','Kingdom of Bhutan','不丹');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gy','+592','Guyana','Co-operative Republic of Guyana','圭亚那');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gs','+500','South Georgia','South Georgia and the South Sandwich Islands','南乔治亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ba','+387','Bosnia and Herzegovina','Bosnia and Herzegovina','波斯尼亚和黑塞哥维那');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bn','+673','Brunei','Nation of Brunei, Abode of Peace','文莱');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pk','+92','Pakistan','Islamic Republic of Pakistan','巴基斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ke','+254','Kenya','Republic of Kenya','肯尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pr','+1787','Puerto Rico','Commonwealth of Puerto Rico','波多黎各');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('so','+252','Somalia','Federal Republic of Somalia','索马里');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sj','+4779','Svalbard and Jan Mayen','Svalbard og Jan Mayen','斯瓦尔巴特');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sl','+232','Sierra Leone','Republic of Sierra Leone','塞拉利昂');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pf','+689','French Polynesia','French Polynesia','法属波利尼西亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('az','+994','Azerbaijan','Republic of Azerbaijan','阿塞拜疆');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ck','+682','Cook Islands','Cook Islands','库克群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pe','+51','Peru','Republic of Peru','秘鲁');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bv','+47','Bouvet Island','Bouvet Island','布维岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mp','+1670','Northern Mariana Islands','Commonwealth of the Northern Mariana Islands','北马里亚纳群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ao','+244','Angola','Republic of Angola','安哥拉');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cg','+242','Republic of the Congo','Republic of the Congo','刚果');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ss','+211','South Sudan','Republic of South Sudan','南苏丹');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mf','+590','Saint Martin','Saint Martin','圣马丁');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tr','+90','Turkey','Republic of Turkey','土耳其');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ai','+1264','Anguilla','Anguilla','安圭拉');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kn','+1869','Saint Kitts and Nevis','Federation of Saint Christopher and Nevis','圣基茨和尼维斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('aw','+297','Aruba','Aruba','阿鲁巴');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tc','+1649','Turks and Caicos Islands','Turks and Caicos Islands','特克斯和凯科斯群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tw','+886','Taiwan','Republic of China (Taiwan)','台灣');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('se','+46','Sweden','Kingdom of Sweden','瑞典');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lr','+231','Liberia','Republic of Liberia','利比里亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ve','+58','Venezuela','Bolivarian Republic of Venezuela','委内瑞拉');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('vi','+1340','United States Virgin Islands','Virgin Islands of the United States','美属维尔京群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bm','+1441','Bermuda','Bermuda','百慕大');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('al','+355','Albania','Republic of Albania','阿尔巴尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('hk','+852','Hong Kong','Hong Kong Special Administrative Region of the People''s Republic of China','香港');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lu','+352','Luxembourg','Grand Duchy of Luxembourg','卢森堡');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('er','+291','Eritrea','State of Eritrea','厄立特里亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('co','+57','Colombia','Republic of Colombia','哥伦比亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bq','+599','Caribbean Netherlands','Bonaire, Sint Eustatius and Saba','荷蘭加勒比區');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mn','+976','Mongolia','Mongolia','蒙古');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ye','+967','Yemen','Republic of Yemen','也门');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lb','+961','Lebanon','Lebanese Republic','黎巴嫩');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ag','+1268','Antigua and Barbuda','Antigua and Barbuda','安提瓜和巴布达');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('vn','+84','Vietnam','Socialist Republic of Vietnam','越南');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pw','+680','Palau','Republic of Palau','帕劳');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('je','+44','Jersey','Bailiwick of Jersey','泽西岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tt','+1868','Trinidad and Tobago','Republic of Trinidad and Tobago','特立尼达和多巴哥');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('il','+972','Israel','State of Israel','以色列');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bg','+359','Bulgaria','Republic of Bulgaria','保加利亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pt','+351','Portugal','Portuguese Republic','葡萄牙');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gi','+350','Gibraltar','Gibraltar','直布罗陀');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sm','+378','San Marino','Republic of San Marino','圣马力诺');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sg','+65','Singapore','Republic of Singapore','新加坡');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sx','+1721','Sint Maarten','Sint Maarten','圣马丁岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sa','+966','Saudi Arabia','Kingdom of Saudi Arabia','沙特阿拉伯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gh','+233','Ghana','Republic of Ghana','加纳');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('md','+373','Moldova','Republic of Moldova','摩尔多瓦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('td','+235','Chad','Republic of Chad','乍得');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('aq','N/A','Antarctica','Antarctica','南极洲');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gb','+44','United Kingdom','United Kingdom of Great Britain and Northern Ireland','英国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pg','+675','Papua New Guinea','Independent State of Papua New Guinea','巴布亚新几内亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tf','+262','French Southern and Antarctic Lands','Territory of the French Southern and Antarctic Lands','法国南部和南极土地');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('um','+268','United States Minor Outlying Islands','United States Minor Outlying Islands','美国本土外小岛屿');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bz','+501','Belize','Belize','伯利兹');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('km','+269','Comoros','Union of the Comoros','科摩罗');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bf','+226','Burkina Faso','Burkina Faso','布基纳法索');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('fo','+298','Faroe Islands','Faroe Islands','法罗群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gn','+224','Guinea','Republic of Guinea','几内亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('in','+91','India','Republic of India','印度');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cw','+599','Curaçao','Country of Curaçao','库拉索');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tg','+228','Togo','Togolese Republic','多哥');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tn','+216','Tunisia','Tunisian Republic','突尼斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bl','+590','Saint Barthélemy','Collectivity of Saint Barthélemy','圣巴泰勒米');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('id','+62','Indonesia','Republic of Indonesia','印度尼西亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('fm','+691','Micronesia','Federated States of Micronesia','密克罗尼西亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('at','+43','Austria','Republic of Austria','奥地利');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tj','+992','Tajikistan','Republic of Tajikistan','塔吉克斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cd','+243','DR Congo','Democratic Republic of the Congo','民主刚果');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('yt','+262','Mayotte','Department of Mayotte','马约特');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('re','+262','Réunion','Réunion Island','留尼旺岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ro','+40','Romania','Romania','罗马尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('qa','+974','Qatar','State of Qatar','卡塔尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lt','+370','Lithuania','Republic of Lithuania','立陶宛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cn','+86','China','People''s Republic of China','中国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('nz','+64','New Zealand','New Zealand','新西兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('nf','+672','Norfolk Island','Territory of Norfolk Island','诺福克岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mr','+222','Mauritania','Islamic Republic of Mauritania','毛里塔尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('uz','+998','Uzbekistan','Republic of Uzbekistan','乌兹别克斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('fi','+358','Finland','Republic of Finland','芬兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cm','+237','Cameroon','Republic of Cameroon','喀麦隆');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('hm','N/A','Heard Island and McDonald Islands','Heard Island and McDonald Islands','赫德岛和麦当劳群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('dm','+1767','Dominica','Commonwealth of Dominica','多米尼加');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('is','+354','Iceland','Iceland','冰岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('om','+968','Oman','Sultanate of Oman','阿曼');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mk','+389','North Macedonia','Republic of North Macedonia','北馬其頓');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('li','+423','Liechtenstein','Principality of Liechtenstein','列支敦士登');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('es','+34','Spain','Kingdom of Spain','西班牙');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gr','+30','Greece','Hellenic Republic','希腊');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('py','+595','Paraguay','Republic of Paraguay','巴拉圭');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bh','+973','Bahrain','Kingdom of Bahrain','巴林');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('nu','+683','Niue','Niue','纽埃');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sn','+221','Senegal','Republic of Senegal','塞内加尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ms','+1664','Montserrat','Montserrat','蒙特塞拉特');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lv','+371','Latvia','Republic of Latvia','拉脱维亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tk','+690','Tokelau','Tokelau','托克劳');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('jp','+81','Japan','Japan','日本');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cf','+236','Central African Republic','Central African Republic','中非共和国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ga','+241','Gabon','Gabonese Republic','加蓬');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('iq','+964','Iraq','Republic of Iraq','伊拉克');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('im','+44','Isle of Man','Isle of Man','马恩岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mm','+95','Myanmar','Republic of the Union of Myanmar','缅甸');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('me','+382','Montenegro','Montenegro','黑山');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('nr','+674','Nauru','Republic of Nauru','瑙鲁');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('vu','+678','Vanuatu','Republic of Vanuatu','瓦努阿图');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('fr','+33','France','French Republic','法国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('zw','+263','Zimbabwe','Republic of Zimbabwe','津巴布韦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ph','+63','Philippines','Republic of the Philippines','菲律宾');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sk','+421','Slovakia','Slovak Republic','斯洛伐克');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('au','+61','Australia','Commonwealth of Australia','澳大利亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ci','+225','Ivory Coast','Republic of Côte d''Ivoire','科特迪瓦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('io','+246','British Indian Ocean Territory','British Indian Ocean Territory','英属印度洋领地');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sz','+268','Eswatini','Kingdom of Eswatini','斯威士兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('rw','+250','Rwanda','Republic of Rwanda','卢旺达');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bj','+229','Benin','Republic of Benin','贝宁');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('vg','+1284','British Virgin Islands','Virgin Islands','英属维尔京群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ug','+256','Uganda','Republic of Uganda','乌干达');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ie','+353','Ireland','Republic of Ireland','爱尔兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ir','+98','Iran','Islamic Republic of Iran','伊朗');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('si','+386','Slovenia','Republic of Slovenia','斯洛文尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ml','+223','Mali','Republic of Mali','马里');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ch','+41','Switzerland','Swiss Confederation','瑞士');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('as','+1684','American Samoa','American Samoa','美属萨摩亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('uy','+598','Uruguay','Oriental Republic of Uruguay','乌拉圭');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gu','+1671','Guam','Guam','关岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sr','+597','Suriname','Republic of Suriname','苏里南');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ua','+380','Ukraine','Ukraine','乌克兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cz','+420','Czechia','Czech Republic','捷克');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('hn','+504','Honduras','Republic of Honduras','洪都拉斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ws','+685','Samoa','Independent State of Samoa','萨摩亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('la','+856','Laos','Lao People''s Democratic Republic','老挝');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cv','+238','Cape Verde','Republic of Cabo Verde','佛得角');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('et','+251','Ethiopia','Federal Democratic Republic of Ethiopia','埃塞俄比亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bd','+880','Bangladesh','People''s Republic of Bangladesh','孟加拉国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sh','+290','Saint Helena, Ascension and Tristan da Cunha','Saint Helena, Ascension and Tristan da Cunha','圣赫勒拿、阿森松和特里斯坦-达库尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('by','+375','Belarus','Republic of Belarus','白俄罗斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('hr','+385','Croatia','Republic of Croatia','克罗地亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kw','+965','Kuwait','State of Kuwait','科威特');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gf','+594','French Guiana','Guiana','法属圭亚那');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ma','+212','Morocco','Kingdom of Morocco','摩洛哥');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ru','+73','Russia','Russian Federation','俄罗斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ee','+372','Estonia','Republic of Estonia','爱沙尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lk','+94','Sri Lanka','Democratic Socialist Republic of Sri Lanka','斯里兰卡');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('nc','+687','New Caledonia','New Caledonia','新喀里多尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pl','+48','Poland','Republic of Poland','波兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mg','+261','Madagascar','Republic of Madagascar','马达加斯加');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cr','+506','Costa Rica','Republic of Costa Rica','哥斯达黎加');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sv','+503','El Salvador','Republic of El Salvador','萨尔瓦多');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mo','+853','Macau','Macao Special Administrative Region of the People''s Republic of China','澳门');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ad','+376','Andorra','Principality of Andorra','安道尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('it','+39','Italy','Italian Republic','意大利');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('na','+264','Namibia','Republic of Namibia','纳米比亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sc','+248','Seychelles','Republic of Seychelles','塞舌尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pa','+507','Panama','Republic of Panama','巴拿马');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ht','+509','Haiti','Republic of Haiti','海地');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ar','+54','Argentina','Argentine Republic','阿根廷');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ne','+227','Niger','Republic of Niger','尼日尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mw','+265','Malawi','Republic of Malawi','马拉维');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pn','+64','Pitcairn Islands','Pitcairn Group of Islands','皮特凯恩群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('de','+49','Germany','Federal Republic of Germany','德国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ki','+686','Kiribati','Independent and Sovereign Republic of Kiribati','基里巴斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sy','+963','Syria','Syrian Arab Republic','叙利亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mh','+692','Marshall Islands','Republic of the Marshall Islands','马绍尔群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('hu','+36','Hungary','Hungary','匈牙利');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ky','+1345','Cayman Islands','Cayman Islands','开曼群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('dk','+45','Denmark','Kingdom of Denmark','丹麦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('lc','+1758','Saint Lucia','Saint Lucia','圣卢西亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bo','+591','Bolivia','Plurinational State of Bolivia','玻利维亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('dj','+253','Djibouti','Republic of Djibouti','吉布提');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('za','+27','South Africa','Republic of South Africa','南非');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ng','+234','Nigeria','Federal Republic of Nigeria','尼日利亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('st','+239','São Tomé and Príncipe','Democratic Republic of São Tomé and Príncipe','圣多美和普林西比');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ni','+505','Nicaragua','Republic of Nicaragua','尼加拉瓜');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gp','+590','Guadeloupe','Guadeloupe','瓜德罗普岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('pm','+508','Saint Pierre and Miquelon','Saint Pierre and Miquelon','圣皮埃尔和密克隆');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ec','+593','Ecuador','Republic of Ecuador','厄瓜多尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gl','+299','Greenland','Greenland','格陵兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gd','+1473','Grenada','Grenada','格林纳达');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bs','+1242','Bahamas','Commonwealth of the Bahamas','巴哈马');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cl','+56','Chile','Republic of Chile','智利');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('my','+60','Malaysia','Malaysia','马来西亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tv','+688','Tuvalu','Tuvalu','图瓦卢');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cx','+61','Christmas Island','Territory of Christmas Island','圣诞岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('sb','+677','Solomon Islands','Solomon Islands','所罗门群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tz','+255','Tanzania','United Republic of Tanzania','坦桑尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kp','+850','North Korea','Democratic People''s Republic of Korea','朝鲜');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gw','+245','Guinea-Bissau','Republic of Guinea-Bissau','几内亚比绍');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('xk','+383','Kosovo','Republic of Kosovo','科索沃');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('va','+3906698','Vatican City','Vatican City State','梵蒂冈');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('no','+47','Norway','Kingdom of Norway','挪威');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ps','+970','Palestine','State of Palestine','巴勒斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cc','+61','Cocos (Keeling) Islands','Territory of the Cocos (Keeling) Islands','科科斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('jm','+1876','Jamaica','Jamaica','牙买加');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('eg','+20','Egypt','Arab Republic of Egypt','埃及');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kh','+855','Cambodia','Kingdom of Cambodia','柬埔寨');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mu','+230','Mauritius','Republic of Mauritius','毛里求斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gm','+220','Gambia','Republic of the Gambia','冈比亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gq','+240','Equatorial Guinea','Republic of Equatorial Guinea','赤道几内亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ls','+266','Lesotho','Kingdom of Lesotho','莱索托');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mq','+596','Martinique','Martinique','马提尼克');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('us','+1201','United States','United States of America','美国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('eh','+2125288','Western Sahara','Sahrawi Arab Democratic Republic','西撒哈拉');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ae','+971','United Arab Emirates','United Arab Emirates','阿拉伯联合酋长国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mz','+258','Mozambique','Republic of Mozambique','莫桑比克');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('dz','+213','Algeria','People''s Democratic Republic of Algeria','阿尔及利亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('zm','+260','Zambia','Republic of Zambia','赞比亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gt','+502','Guatemala','Republic of Guatemala','危地马拉');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kr','+82','South Korea','Republic of Korea','韩国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kg','+996','Kyrgyzstan','Kyrgyz Republic','吉尔吉斯斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tl','+670','Timor-Leste','Democratic Republic of Timor-Leste','东帝汶');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ax','+35818','Åland Islands','Åland Islands','奥兰群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('jo','+962','Jordan','Hashemite Kingdom of Jordan','约旦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mt','+356','Malta','Republic of Malta','马耳他');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('cy','+357','Cyprus','Republic of Cyprus','塞浦路斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('fk','+500','Falkland Islands','Falkland Islands','福克兰群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('kz','+76','Kazakhstan','Republic of Kazakhstan','哈萨克斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bw','+267','Botswana','Republic of Botswana','博茨瓦纳');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('vc','+1784','Saint Vincent and the Grenadines','Saint Vincent and the Grenadines','圣文森特和格林纳丁斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('bb','+1246','Barbados','Barbados','巴巴多斯');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('to','+676','Tonga','Kingdom of Tonga','汤加');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('th','+66','Thailand','Kingdom of Thailand','泰国');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('be','+32','Belgium','Kingdom of Belgium','比利时');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ca','+1','Canada','Canada','加拿大');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ge','+995','Georgia','Georgia','格鲁吉亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('wf','+681','Wallis and Futuna','Territory of the Wallis and Futuna Islands','瓦利斯和富图纳群岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('fj','+679','Fiji','Republic of Fiji','斐济');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('nl','+31','Netherlands','Kingdom of the Netherlands','荷兰');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('am','+374','Armenia','Republic of Armenia','亚美尼亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('do','+1809','Dominican Republic','Dominican Republic','多明尼加');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('gg','+44','Guernsey','Bailiwick of Guernsey','根西岛');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('tm','+993','Turkmenistan','Turkmenistan','土库曼斯坦');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('np','+977','Nepal','Federal Democratic Republic of Nepal','尼泊尔');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('mv','+960','Maldives','Republic of the Maldives','马尔代夫');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('ly','+218','Libya','State of Libya','利比亚');
INSERT INTO countries(code,calling_code,name,official_name,name_ch) VALUES ('br','+55','Brazil','Federative Republic of Brazil','巴西');
