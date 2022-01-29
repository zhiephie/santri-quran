package quran

import (
	"github.com/santri-quran/database"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Quran struct {
	Number        int    `json:"number"`
	Asma          string `json:"asma"`
	Name          string `json:"name"`
	TranslationId string `json:"translation_id"`
	TranslationEn string `json:"translation_en"`
	NumberOfAyahs int    `json:"number_of_ayahs"`
	TypeId        string `json:"type_id"`
	TypeEn        string `json:"type_en"`
	OrderNumber   int    `json:"order_number"`
}

type Surah struct {
	Quran
	Ayat []struct {
		Id          int    `json:"id"`
		SurahId     int    `json:"surah_id"`
		VerseId     int    `json:"verse_id"`
		AyahText    string `json:"ayah_text"`
		IndoText    string `json:"indo_text"`
		EnText      string `json:"en_text"`
		ReadText    string `json:"read_text"`
		Juz         int    `json:"juz"`
		Manzil      int    `json:"manzil"`
		Page        int    `json:"page"`
		Ruku        int    `json:"ruku"`
		HizbQuarter int    `json:"hizb_quarter"`
		Sajda       int    `json:"sajda"`
	} `json:"ayat"`

	Prev Prev `json:"prev"`
	Next Next `json:"next"`
}

type Next struct {
	Quran
}

type Prev struct {
	Quran
}

func GetSurahs(c *fiber.Ctx) error {
	db := database.DBConn
	var result []Quran
	db.Raw("SELECT * FROM info_surah").Scan(&result)

	return c.JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
}

func GetSurah(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var result = Surah{}
	db.Raw("SELECT * FROM info_surah WHERE number = ?", id).Scan(&result.Quran)
	db.Raw("SELECT * FROM quran_id WHERE surah_id = ?", id).Scan(&result.Ayat)
	if id == "1" {
		result.Prev.Quran = result.Quran
	} else {
		prev := result.Quran.Number - 1
		db.Raw("SELECT * FROM info_surah WHERE number = ?", prev).Scan(&result.Prev.Quran)
	}
	next := result.Quran.Number + 1
	db.Raw("SELECT * FROM info_surah WHERE number = ?", next).Scan(&result.Next.Quran)

	return c.JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
}

func SearchSurah(c *fiber.Ctx) error {
	q := c.Query("q")
	db := database.DBConn
	var result []Quran
	db.Raw("SELECT * FROM info_surah WHERE name LIKE ? ", "%"+q+"%").Scan(&result)

	return c.JSON(&fiber.Map{
		"success": true,
		"data":    result,
	})
}
