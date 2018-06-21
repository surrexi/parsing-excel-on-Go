package main

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Database struct {
    Driver   	 string
    Port     	 int
    User     	 string
    Password 	 string
    DatabaseName string
}

type SeoInfo struct {
    Id          int
    Title       string
    Description string
}

func dbConn() (db *sql.DB) {
    dbConf := config.DB
    db, err := sql.Open(dbConf.Driver, dbConf.User+":"+dbConf.Password+"@/"+dbConf.DatabaseName)
    CheckError(err)
    return db
}

func GetMetaSeo(link string) SeoInfo {
    db := dbConn()

    var seoId int
    var title string
    var article string
    var documentation string
    var titleH1 string

    sqlStatement := `
        SELECT
		  p.seo_info_id,
          p.product_title,
          p.product_article,
          concat(doc_type.documentation_type_title, ' ', doc.documentation_number),
          IFNULL(p_h1.product_h1_title, '')
		FROM product p
          LEFT OUTER JOIN product_h1 AS p_h1 ON p_h1.product_h1_id = p.product_h1_id
          INNER JOIN product_documentation AS p_doc ON p_doc.product_id = p.product_id
          INNER JOIN documentation AS doc ON doc.documentation_id = p_doc.documentation_id
          INNER JOIN documentation_type AS doc_type  ON doc_type.documentation_type_id = doc.documentation_type_id
		WHERE p.product_link = ?;
    `

    row := db.QueryRow(sqlStatement, link)
    err := row.Scan(&seoId, &title, &article, &documentation, &titleH1)
    if err != nil {
        log.Print(err, " for link = ", link)
    }

    defer db.Close()

    return SeoInfo {
        Id: seoId,
        Title: title + " " + article + " по документации " + documentation + titleH1 + " – купить с доставкой по всей РФ",
        Description: documentation + " - " + title + " " + article + titleH1 + " Вы можете приобрести по хорошей цене с доставкой в любой регион РФ в компании СТАЛЬБЕТОН на нашем сайте или по телефону +7 (812) 660-55-05",
    }
}

func UpdateSeoInfo(s SeoInfo) {
    db := dbConn()
    if s.Id > 0 {
        updProd, err := db.Prepare("UPDATE seo_info SET seo_info_title = ?, seo_info_description = ? WHERE seo_info_id = ?")

        if err != nil {
            log.Print("Prepare: ", err, " for seo_info = ", s)
        }
        _, err = updProd.Exec(s.Title, s.Description, s.Id)

        if err != nil {
            log.Print("Exec: ", err, " for seo_info = ", s)
        }
    }

    defer db.Close()
}
