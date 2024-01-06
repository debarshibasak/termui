// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	ui "github.com/debarshibasak/termui/v3"
	"github.com/debarshibasak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	table1 := widgets.NewTableScrollable()
	table1.Rows = [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"你好吗", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"2016", "10", "11"},
				[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"Foundations", "Go-lang is so cool", "Im working on Ruby"},
	}

	table1.TopRow = 0

	table1.Header = []string{"test", "left", "right"}
	table1.ColWidths = []int{10,10,10}

	table1.SetRect(0, 0, 80, 10)

	ui.Render(table1)

	table3 := widgets.NewTableScrollable()
	table3.Rows = [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"AAA", "BBB", "CCC"},
		[]string{"DDD", "EEE", "FFF"},
		[]string{"GGG", "HHH", "III"},
	}
	table3.BorderStyle = ui.NewStyle(ui.ColorGreen)
	table3.SetRect(0, 30, 70, 20)

	ui.Render(table3)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		fmt.Println(e.ID)
		switch e.ID {
		case "q", "<C-c>":
			return
	     case "<Down>":
			fmt.Print("down detected")
				table1.ScrollDown()

		}
	}
}
