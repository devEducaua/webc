package main

import (
	"fmt"
	"os"
	"strconv"

	"webc/internal"
)

func main() {
	argv := os.Args[1:];
	if len(argv) == 0 {
		os.Exit(1);
	}
	
	if err := parseArgs(argv); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
	}
}

func parseArgs(argv []string) error {
	switch argv[0] {
	case "fetch":
		if len(argv) < 2 {
			return fmt.Errorf("ERROR: fetch command needs an url");
		}
		url := argv[1];
		dt, err := internal.Fetch(url);
		if err != nil {
			return err;
		}
		fmt.Println(dt.Body);
	case "parse":
		if len(argv) < 2 {
			return fmt.Errorf("parse command needs an url");
		}
		url := argv[1];
		dt, err := internal.Parse(url);
		if err != nil {
			return err;
		}
		for _,tk := range dt.Tokens {
			fmt.Println(tk);
		}
	case "tab":
		if len(argv) < 2 {
			return fmt.Errorf("tab command needs a subcommand");
		}
		err := parseTabSubcommands(argv);
		if err != nil {
			return err;
		}
	default:
		return fmt.Errorf("invalid command: %v", argv[0]);
	}
	return nil;
}

func parseTabSubcommands(argv []string) error {
	sub := argv[1];
	switch sub {
	case "get":
		tab, err := internal.TabGet();
		if err != nil {
			return err;
		}
		fmt.Println(tab.Content);
	case "new":
		if len(argv) < 3 {
			return fmt.Errorf("ERROR: `tab new` command needs an url");
		}
		url := argv[2];
		tab, err := internal.TabNew(url);
		if err != nil {
			panic(err);
		}
		fmt.Println(tab);
	case "all":
		list, err := internal.TabAll();
		if err != nil {
			return err;
		}
		for i,t := range list {
			fmt.Printf("%v : %v\n", i, t.Title);
		}
	case "sel":
		if len(argv) < 3 {
			return fmt.Errorf("ERROR: `tab sel` command needs the tab id");
		}
		id, err := strconv.Atoi(argv[2]);
		if err != nil {
			return err;
		}
		tab, err := internal.TabSel(id);
		if err != nil {
			return err;
		}
		fmt.Printf("changed to tab: %v\n", tab.Title);
	case "del":
		if len(argv) < 3 {
			return fmt.Errorf("ERROR: `tab del` command needs the tab id");
		}
		id, err := strconv.Atoi(argv[2]);
		if err != nil {
			return err;
		}
		err = internal.TabDel(id);
		if err != nil {
			return err;
		}
	case "put":
		if len(argv) < 3 {
			return fmt.Errorf("ERROR: `tab new` command needs an url");
		}
		url := argv[2];
		tab, err := internal.TabNew(url);
		if err != nil {
			panic(err);
		}
		fmt.Println(tab);
	default:
		return fmt.Errorf("invalid subcommand to tab: %v", argv[1]);
	}

	return nil;
}

