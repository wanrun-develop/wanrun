``` mermaid
---
title: DogRunner
---
erDiagram
	
	
  dog_owners {
      serial dog_owner_id PK
      varchar(64) name "ユーザー名"
      varchar(255) email UK "メアド"
      text image "プロフィール画像"
      char(1) sex "性別"
      timestamp reg_at
      timestamp upd_at
  }
  
  dogs o{--|| dog_owners : "1の飼い主は0以上の犬を所有"
  dogs ||--|| dog_type_mst :"1の犬は1の犬種を持つ"
  dogs {
      serial dog_id PK
      reference owner FK
      varchar(64) name "犬の名前"
      reference dog_type "犬種"
      int weight "体重"
      int sex "性別"
      text image "写真"
      timestamp created_at
      timestamp update_at
  }
  
  dog_type_mst {
	  serial dog_type_id 
	  varchar(32) name "犬種名"
  }
  
  injection_certifications o{--|| dogs : "1の犬は0以上の予防注射照明を持つ"
  injection_certifications {
	  serial injection_certification_id PK
	  reference dog FK
	  int type "予防注射タイプ"
	  text file "証明書ファイル"
	  timestamp created_at
      timestamp update_at
	}
	

  dogrun_managers {
      serial dogrun_manager_id PK
      varchar(128) name "ユーザー名"
      varchar(255) email UK "メアド"
      timestamp created_at
      timestamp update_at
  }

	dogruns o{--|| dogrun_managers : "1のドッグラン運営は1以上のコメントをドッグラン"
  dogruns {
      serial dogrun_id PK
      reference dogrun_manager FK 
      varchar name "ドッグラン名"
      varchar address "住所"
      varchar post_code "郵便番号"
      int business_day "営業日"
      int holiday "休業日"
      time open_time "営業開始時間"
      time close_time "営業終了時間" 
      text description "その他詳細説明"
      timestamp created_at
      timestamp update_at
  }
  
  dogrun_images o{--|| dogruns : "1のドッグランは0以上のドッグラン画像を持つ"
  dogrun_images {   
      serial dogrun_image_id PK
      eference dogrun FK
	  text image 
      int order
	  timestamp upload_at
  }
  
  dogrun_tags ||--||dogruns : "1のドッグランは0以上のドッグランタグを持つ"
  dogrun_tags {
	  serial dogrun_detail_id PK
	  reference dogrun FK
	  int tag 
  }
  
  tag_master o{--|| dogrun_tags : "1のドッグランタグは0以上のタグマスタを持つ"
  tag_master {
	  serial tag_id PK
	  varchar(64) name "タグ名"
	  text description "説明/定義"
  }
  
	auth_dog_owners ||--|| dog_owners : "1ユーザー1認証認証情報"
  auth_dog_owners{
	  serial auth_dog_owner_id PK
	  references dog_owner 
	  varchar(256) password
	  int grant_type
	  timestamp pass_regist_at
	  
  }
  
  auth_dogrun_managers ||--|| dogrun_managers:"1ユーザー1認証認証情報"
  auth_dogrun_managers{
    serial auth_dogrun_manager_id PK
	  references dogrun_manager 
	  varchar(256) password
	  int grant_type
	  timestamp pass_regist_at
  }
```