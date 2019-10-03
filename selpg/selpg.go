package main
import (
	"io"
	"os/exec"
	"bufio"
	"os"
	"fmt"
	flag "github.com/spf13/pflag"
)
type selpg_args struct {
	start_page int 
	end_page int
	in_filename string 
	print_dest string	
	page_len int 
	page_type string 
}
var name string
func process_args(sa * selpg_args) {
	flag.IntVarP(&sa.start_page,"start",  "s", -1, "start page")
	flag.IntVarP(&sa.end_page,"end", "e",  -1, "end page")
	flag.IntVarP(&sa.page_len,"len", "l", 72, "length is page")
	flag.StringVarP(&sa.print_dest,"dest", "d", "", "print dest")
	flag.StringVarP(&sa.page_type,"type", "f", "l", "'l' for lines-limited, 'f' for form-limited")
	flag.Lookup("type").NoOptDefVal = "f"
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,"USAGE: \n%s -s start_page -e end_page [ -f | -l lines_per_page ]\n[ -d dest ] [ in_filename ]\n", name)
		flag.PrintDefaults()
	}
	flag.Parse()
	max_page := 99999999
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "\n%s: not enough arguments\n", name)
		flag.Usage()
		os.Exit(1)
	}
	if(os.Args[1] != "-s"||os.Args[3] != "-e"){
		fmt.Fprintf(os.Stderr, "\n%s: the argument is wrong\n", name)
		flag.Usage()
		os.Exit(2)
	}
	if(sa.start_page < 1 || sa.start_page > max_page||sa.end_page < 1 || sa.end_page > max_page || sa.end_page < sa.start_page) {
		fmt.Fprintf(os.Stderr, "\nthe start_page page or end page is invalid\n")
		flag.Usage()
		os.Exit(3)
	}
	if ( sa.page_len < 1 || sa.page_len > (max_page - 1) ) {
		fmt.Fprintf(os.Stderr, "\n the page length is invalid\n")
		flag.Usage()
		os.Exit(4)
	}
	if len(flag.Args()) == 1 {
		_, err := os.Stat(flag.Args()[0])
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "\nthe file doesn't exist");
			os.Exit(5);
		}
		sa.in_filename = flag.Args()[0]
	}
}
func process_input(sa selpg_args) {
	var fin *os.File 
	var fout io.WriteCloser
	var err error
	var err1 error
	var err2 error
	if sa.in_filename!=""{
		fin, err = os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\ncould not open input file \n")
			os.Exit(6)
		}
		defer fin.Close()
	}else{
		fin = os.Stdin
	}
	bufFin := bufio.NewReader(fin)
	cmd := &exec.Cmd{}
	if sa.print_dest!=""{
		cmd = exec.Command("cat")
		cmd.Stdout, err1 = os.OpenFile(sa.print_dest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "\n%s: could not open file %s\n",
				name, sa.print_dest)
			os.Exit(7)
		}
		fout, err2 = cmd.StdinPipe()
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "\n%s: could not open pipe to \"lp -d%s\"\n",
				name, sa.print_dest)
			os.Exit(8)
		}
		cmd.Start()
		defer fout.Close()
	}else{
		fout = os.Stdout
	}
	var page_now int
	var line_now int
	if sa.page_type == "l" { 
		line_now = 0
		page_now = 1
		for {
			line,c := bufFin.ReadString('\n')
			if c != nil {
				break 
			}
			line_now++
			if line_now > sa.page_len {
				page_now++
				line_now = 1
			}
			if page_now >= sa.start_page && page_now <= sa.end_page {
				_, err3 := fout.Write([]byte(line))
				if err3 != nil {
					fmt.Println(err3)
					os.Exit(9)
				}
		 	}
		}  
	} else {			
		page_now = 1
		for {
			page, err4 := bufFin.ReadString('\n')
			if err4 != nil {
				break 
			}
			if page_now >= sa.start_page && page_now <= sa.end_page {
				_, err4 := fout.Write([]byte(page))
				if err4 != nil {
					os.Exit(10)
				}
			}
			page_now++
		}
	}
		if page_now < sa.start_page {
			fmt.Fprintf(os.Stderr,"\n%s: start_page (%d) greater than total pages (%d),no output written\n", name, sa.start_page, page_now)
		} else if page_now < sa.end_page {
			fmt.Fprintf(os.Stderr,"\n%s: end_page (%d) greater than total pages (%d),less output than expected\n", name, sa.end_page, page_now)
		}
}
func main() {
	sa := selpg_args{}
	name = os.Args[0]
	process_args(&sa)
	process_input(sa)
}