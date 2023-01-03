package cli

import (
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/pkg"
	"github.com/crispybaccoon/hayashi/util"
)

func Init() error {
	err := os.MkdirAll(util.HAYASHI_ROOT, os.ModeAppend.Perm())
	if err != nil {
		return err
	}
	err = os.MkdirAll(util.PKG_ROOT, os.ModeAppend.Perm())
	if err != nil {
		return err
	}
	err = os.MkdirAll(util.REPO_ROOT, os.ModeAppend.Perm())
	if err != nil {
		return err
	}
	err = os.MkdirAll(util.PathPkg("custom", ""), os.ModeAppend.Perm())
	if err != nil {
		return err
	}

	return nil
}

func Create() error {

	if util.PathExists(util.PathConfig()) {
		return fmt.Errorf(COLOR_RED + "config file " + COLOR_CYAN +
			util.PathConfig() + COLOR_RED + " already exists.")
	}
	printf("creating config file at " + COLOR_CYAN +
		util.PathConfig() + COLOR_RESET + " ...")
	err := pkg.SaveConfig(pkg.DefaultConfig())
	if err != nil {
		return err
	}

	return nil
}
