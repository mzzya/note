CREATE TABLE "i2h_history_image_cut" (
  "id" int(11) unsigned NOT NULL AUTO_INCREMENT,
  "hash" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件hash值',
  "width" int(11) NOT NULL DEFAULT '0' COMMENT '宽度',
  "height" int(11) NOT NULL DEFAULT '0' COMMENT '高度',
  "data" mediumtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '数据',
  "status" varchar(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'A' COMMENT '状态,APX',
  "create_user_id" int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  "create_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  "update_user_id" int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  "update_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY ("id") USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文件切图表';


CREATE TABLE "i2h_record" (
  "id" int(11) unsigned NOT NULL AUTO_INCREMENT,
  "hash" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件hash值',
  "data" mediumtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '数据',
  "status" varchar(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'A' COMMENT '状态,APX',
  "create_user_id" int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  "create_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  "update_user_id" int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  "update_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY ("id") USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=144 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='记录表';


CREATE TABLE "i2h_resource" (
  "id" int(11) unsigned NOT NULL AUTO_INCREMENT,
  "hash" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件hash值',
  "size" bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小，字节',
  "type" varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '资源类型，image',
  "suffix" varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件后缀名',
  "width" int(11) NOT NULL DEFAULT '0' COMMENT '宽度',
  "height" int(11) NOT NULL DEFAULT '0' COMMENT '高度',
  "oss_path" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '上传到OSS地址，不包含域名部分',
  "original_file_name" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '原始文件名',
  "content_type" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '上传时资源类型',
  "status" varchar(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'A' COMMENT '状态,APX',
  "create_user_id" int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  "create_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  "update_user_id" int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  "update_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY ("id") USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='资源表';