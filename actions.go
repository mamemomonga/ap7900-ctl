package main

import (
	"fmt"
	expect "github.com/google/goexpect"
	"github.com/google/goterm/term"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func outputManagement(exp expect.Expecter, port int, action int) {

	//  1- Immediate On
	//  2- Immediate Off
	//  3- Immediate Reboot
	//  4- Delayed On
	//  5- Delayed Off
	//  6- Delayed Reboot

	res, err := exp.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: `User Name : `},
		&expect.BSnd{S: fmt.Sprintf("%s\r\n", Conf().Username)},
		&expect.BExp{R: `Password  : `},
		&expect.BSnd{S: fmt.Sprintf("%s\r\n", Conf().Password)},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Device Manager
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "2\r\n"}, // 2- Outlet Management
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Outlet Control/Configuration
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: fmt.Sprintf("%d\r\n", port)}, // Outlet
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Control Outlet
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: fmt.Sprintf("%d\r\n", action)}, // Action
		&expect.BExp{R: `YES`},
		&expect.BSnd{S: "YES\r\n"}, // YES
	}, timeout)
	if err != nil {
		fmt.Println(term.Redf("exp.ExpectBatch failed: %v , res: %v", err, res))
		os.Exit(1)
	}

	bufs := strings.Split(res[len(res)-1].Output, "\n")
	fmt.Println(term.Greenf(strings.Join(bufs[3:len(bufs)-2], "\n")))

	res, err = exp.ExpectBatch([]expect.Batcher{
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "4\r\n"},
	}, timeout)
	if err != nil {
		fmt.Println(term.Redf("exp.ExpectBatch failed: %v , res: %v", err, res))
		os.Exit(1)
	}

}

func outputStatus(exp expect.Expecter, port int) (ret string) {

	ret = ""
	res, err := exp.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: `User Name : `},
		&expect.BSnd{S: fmt.Sprintf("%s\r\n", Conf().Username)},
		&expect.BExp{R: `Password  : `},
		&expect.BSnd{S: fmt.Sprintf("%s\r\n", Conf().Password)},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Device Manager
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "2\r\n"}, // 2- Outlet Management
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Outlet Control/Configuration
		&expect.BExp{R: `<ESC>`},
	}, timeout)
	if err != nil {
		fmt.Println(term.Redf("exp.ExpectBatch failed: %v , res: %v", err, res))
		os.Exit(1)
	}

	result := res[len(res)-1].Output

	res, err = exp.ExpectBatch([]expect.Batcher{
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "4\r\n"},
	}, timeout)
	if err != nil {
		fmt.Println(term.Redf("exp.ExpectBatch failed: %v , res: %v", err, res))
		os.Exit(1)
	}

	bufs := strings.Split(result, "\n")
	r := regexp.MustCompile(`     ` + strconv.Itoa(port) + `.+(ON|OFF)`)
	for _, v := range bufs {
		rv := r.FindStringSubmatch(v)
		if rv != nil {
			ret = rv[1]
		}
	}

	return ret
}

func loadStatus(exp expect.Expecter) (ret string) {

	ret = ""
	res, err := exp.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: `User Name : `},
		&expect.BSnd{S: fmt.Sprintf("%s\r\n", Conf().Username)},
		&expect.BExp{R: `Password  : `},
		&expect.BSnd{S: fmt.Sprintf("%s\r\n", Conf().Password)},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Device Manager
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "1\r\n"}, // 1- Phase Management
		&expect.BExp{R: `<ESC>`},
	}, timeout)
	if err != nil {
		fmt.Println(term.Redf("exp.ExpectBatch failed: %v , res: %v", err, res))
		os.Exit(1)
	}

	result := res[len(res)-1].Output

	res, err = exp.ExpectBatch([]expect.Batcher{
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: string(27) + "\r\n"},
		&expect.BExp{R: `<ESC>`},
		&expect.BSnd{S: "4\r\n"},
	}, timeout)
	if err != nil {
		fmt.Println(term.Redf("exp.ExpectBatch failed: %v , res: %v", err, res))
		os.Exit(1)
	}

	r := regexp.MustCompile(`Phase Load :\s*([0-9\.]+)`)
	rv := r.FindStringSubmatch(result)
	if rv != nil {
		ret = rv[1]
	}
	return ret
}


