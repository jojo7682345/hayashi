package cli

import "github.com/crispybaccoon/hayashi/pkg"

func Clone(name string) error {
	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	if err = clone_pkg(p, false, false); err != nil {
		return err
	}
	return nil
}

func Build(name string) error {
	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	if err = install(p, false, false); err != nil {
		return err
	}
	return nil
}

func Pack(name string) error {
	p, err := pkg.GetPkg(name)
	if err != nil {
		return err
	}

	printf("packing pkg " + COLOR_MAGENTA + p.Name)

	if err = p.CreatePack(); err != nil {
		return err
	}
	if err = AddInstalled(p); err != nil {
		return err
	}
	return nil
}
