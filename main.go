package main

func main() {
	cfg := LoadConfig()
	db := ConnectDB(cfg)
	defer db.Close()
	r := SetupRouter(db)
	r.Run(":" + cfg.Port)
}
