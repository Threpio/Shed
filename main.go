package main

import (
	"embed"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed schema.sql
var ddl string

//go:embed all:frontend/dist
var assets embed.FS

//func run() error {
//	ctx := context.Background()
//
//	db, err := sql.Open("sqlite3", ":memory:")
//	if err != nil {
//		return err
//	}
//
//	// create tables
//	if _, err := db.ExecContext(ctx, ddl); err != nil {
//		return err
//	}
//
//	_ = data.New(db)
//
//	return nil
//}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "shed",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
