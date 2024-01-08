// gopackagecloud/main.go

package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
)

var (
	selfpackage string = "unknown"
	selfversion string = "unknown"
	this        string = "unknown"
	here        string = "unknown"
	whatami     string = "unknown"
	// Options
	opts struct {
		Config  string `long:"config" description:"Specify a path to config file containing your API token and URL." default:"~/.packagecloud"`
		Url     string `long:"url" description:"Specify the website URL to use." default:"https://packagecloud.io"`
		Verbose bool   `short:"v" long:"verbose" description:"Enable verbose mode."`

		// package_cloud distro SUBCMD ...ARGS
		//
		// manage repositories
		Distro struct {
			// package_cloud distro list package_type
			//
			// list available distros and versions for package_type
			List struct {
			} `command:"list" description:"list available distros and versions for package_type"`
		} `command:"distro" description:"manage repositories"`

		// package_cloud gpg_key SUBCMD ...ARGS
		//
		// manage GPG keys
		GpgKey struct {
		} `command:"gpg_key" description:"manage GPG keys"`

		// package_cloud help [COMMAND]
		//
		// Describe available commands or one specific command
		Help struct {
		} `command:"help" description:"Describe available commands or one specific command."`

		// package_cloud master_token SUBCMD ...ARGS
		//
		// manage master tokens
		MasterToken struct {
		} `command:"master_token" description:"manage master tokens"`

		// package_cloud promote user/repo[/distro/version]
		// [@scope/]package_name user/destination_repo
		//
		// promotes a package from user/repo [in dist/version] to
		// user/destination_repo [also in dist/version]
		Promote struct {
		} `command:"promote" description:"promotes a package from user/repo [in dist/version] to user/destination_repo [also in dist/version]"`

		// package_cloud push user/repo[/distro/version] /path/to/packages
		//
		// Push package(s) to repository (in distro/version, if
		// required). Optional settings shown above.
		Push struct {
		} `command:"push" description:"Push package(s) to repository (in distro/version, if required). Optional settings shown above."`

		// package_cloud read_token SUBCMD ...ARGS
		//
		// manage read tokens
		ReadToken struct {
		} `command:"read_token" description:"manage read tokens"`

		// package_cloud repository SUBCMD ...ARGS
		//
		// manage repositories
		Repository struct {
		} `command:"repository" description:"manage repositories"`

		// // package_cloud version
		// //
		// // print version information
		// Version struct {
		// } `command:"version" description:"Print version information."`

		// package_cloud yank user/repo[/distro/version] [@scope/]package_name
		//
		// yank package from user/repo [in dist/version]
		Yank struct {
		} `command:"yank" description:"yank package from user/repo [in dist/version]"`
	}
	parser *flags.Parser
)

// CmdVersion

type CmdVersion struct{}

func (x *CmdVersion) Execute(args []string) error {
	if 0 < len(args) {
		return errors.New("bad args")
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return errors.New("debug.ReadBuildInfo()")
	}
	if opts.Verbose {
		fmt.Println(info)
	} else {
		fmt.Printf("%v %v %v\n", info.Main.Path, info.Main.Version, info.Main.Sum)
	}
	return nil
}

func _info(format string, v ...interface{}) {
	if opts.Verbose {
		log.Printf("INFO: "+format, v...)
	}
}

func _error(format string, v ...interface{}) {
	log.Printf("ERROR: "+format, v...)
}

func _die(format string, v ...interface{}) {
	_error(format, v...)
	if nil != parser {
		parser.WriteHelp(os.Stderr)
	}
	os.Exit(1)
}

func main() {
	this, err := os.Executable()
	if nil != err {
		_die("os.Executable(): %v", err)
	}
	whatami := filepath.Base(this)
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%v[%v]: ", whatami, os.Getpid()))
	here := filepath.Dir(this)

	parser = flags.NewParser(&opts, flags.Default)
	parser.AddCommand("version", "print version information", "Print version information.", new(CmdVersion))
	args, err := parser.Parse()
	if nil != err {
		// generic parse error
		e, ok := err.(*flags.Error)
		if ok && e.Type == flags.ErrHelp {
			// caller passed '-h' and/or '--help'; parser has already
			// implicitly printed usage to stdout
			os.Exit(0)
		}
		// some other parse error; parser has already implicitly printed error
		// message to stderr
		parser.WriteHelp(os.Stderr)
		os.Exit(1)
	}
	// switch parser.Active.Name {
	// case "distro":
	// 	fmt.Println("distro")
	// 	fmt.Println("parser.Active.Active.Name:", parser.Active.Active.Name)
	// // case "version":
	// // 	fmt.Printf("%v (%v) %v\n", whatami, selfpackage, selfversion)
	// case "help":
	// 	// Clear the active sub command to force WriteHelp to print the generic
	// 	// help message; otherwise, WriteHelp prints the 'help'-specific help
	// 	// message.
	// 	parser.Active = nil
	// 	parser.WriteHelp(os.Stdout)
	// 	// fmt.Printf(parser.Usage)
	// default:
	// 	_die("parser.Active.Name: %v", parser.Active.Name)
	// }
	os.Exit(0)
	// if 0 != len(args) {
	// 	_die("too many arguments")
	// }
	_info("args: %v", args)
	_info("this: %v", this)
	_info("here: %v", here)
	_info("whatami: %v", whatami)
	_info("opts.Help: %v", opts.Help)
	_info("opts.Verbose: %v", opts.Verbose)
	_info("parser.Active: %v", parser.Active)
	_info("parser.Active.Name: %v", parser.Active.Name)
}
