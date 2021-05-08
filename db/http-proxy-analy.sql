CREATE TABLE "hpa_api" (
  "application_id" bigint(20) unsigned NOT NULL DEFAULT '0',
  "url" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "get" tinyint(1) NOT NULL DEFAULT '0',
  "get_summary" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "get_allow_mirror" tinyint(1) NOT NULL DEFAULT '1',
  "post" tinyint(1) NOT NULL DEFAULT '0',
  "post_summary" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "post_allow_mirror" tinyint(1) NOT NULL DEFAULT '0',
  "put" tinyint(1) NOT NULL DEFAULT '0',
  "put_summary" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "put_allow_mirror" tinyint(1) NOT NULL DEFAULT '0',
  "patch" tinyint(1) NOT NULL DEFAULT '0',
  "patch_summary" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "patch_allow_mirror" tinyint(1) NOT NULL DEFAULT '0',
  "delete" tinyint(1) NOT NULL DEFAULT '0',
  "delete_summary" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "delete_allow_mirror" tinyint(1) NOT NULL DEFAULT '0',
  "status" tinyint(1) NOT NULL DEFAULT '0',
  "id" bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" datetime(3) DEFAULT NULL,
  "updated_at" datetime(3) DEFAULT NULL,
  PRIMARY KEY ("id"),
  KEY "idx_hpa_api_application_id" ("application_id"),
  KEY "idx_hpa_api_created_at" ("created_at"),
  KEY "idx_hpa_api_updated_at" ("updated_at"),
  CONSTRAINT "fk_hpa_api_application" FOREIGN KEY ("application_id") REFERENCES "hpa_application" ("id"),
  CONSTRAINT "fk_hpa_application_apis" FOREIGN KEY ("application_id") REFERENCES "hpa_application" ("id")
) ENGINE=InnoDB AUTO_INCREMENT=231 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE "hpa_application" (
  "name" varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "old_host" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "new_host" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "status" tinyint(1) NOT NULL DEFAULT '0',
  "id" bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" datetime(3) DEFAULT NULL,
  "updated_at" datetime(3) DEFAULT NULL,
  "host" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY ("id"),
  KEY "idx_hpa_application_created_at" ("created_at"),
  KEY "idx_hpa_application_updated_at" ("updated_at")
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE "hpa_diff_strategy" (
  "id" bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  "status" tinyint(1) NOT NULL DEFAULT '0',
  "created_at" datetime(3) DEFAULT NULL,
  "updated_at" datetime(3) DEFAULT NULL,
  "code" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "api_id" bigint(20) unsigned NOT NULL DEFAULT '0',
  "field" varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY ("id"),
  KEY "idx_hpa_diff_strategy_created_at" ("created_at"),
  KEY "idx_hpa_diff_strategy_updated_at" ("updated_at")
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE "hpa_proxy_log" (
  "application_id" bigint(20) unsigned NOT NULL DEFAULT '0',
  "api_id" bigint(20) unsigned NOT NULL DEFAULT '0',
  "old_request_method" varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  "old_request_url" text COLLATE utf8mb4_unicode_ci NOT NULL,
  "old_request_header" text COLLATE utf8mb4_unicode_ci NOT NULL,
  "old_request_body" mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  "old_response_header" text COLLATE utf8mb4_unicode_ci NOT NULL,
  "old_response_body" mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  "old_response_status" bigint(20) NOT NULL DEFAULT '0',
  "new_response_header" text COLLATE utf8mb4_unicode_ci NOT NULL,
  "new_response_body" mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  "new_response_status" bigint(20) NOT NULL DEFAULT '0',
  "analysis_result" text COLLATE utf8mb4_unicode_ci NOT NULL,
  "id" bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  "created_at" datetime(3) DEFAULT NULL,
  "updated_at" datetime(3) DEFAULT NULL,
  "new_duration" bigint(20) NOT NULL DEFAULT '0',
  "analysis_diff_count" bigint(20) NOT NULL DEFAULT '0',
  "old_duration" bigint(20) NOT NULL DEFAULT '0',
  "status" tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY ("id"),
  KEY "idx_hpa_proxy_log_new_response_status" ("new_response_status"),
  KEY "idx_hpa_proxy_log_created_at" ("created_at"),
  KEY "idx_hpa_proxy_log_updated_at" ("updated_at"),
  KEY "idx_hpa_proxy_log_application_id" ("application_id"),
  KEY "idx_hpa_proxy_log_api_id" ("api_id"),
  KEY "idx_hpa_proxy_log_old_response_status" ("old_response_status")
) ENGINE=InnoDB AUTO_INCREMENT=2754151 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;