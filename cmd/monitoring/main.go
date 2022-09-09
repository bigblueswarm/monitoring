package main

import "github.com/b3lb/monitoring/pkg/app"

func main() {
	run()
}

func run() error {
	err := app.NewServer().Run()
	if err != nil {
		return err
	}

	return nil
}
