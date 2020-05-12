package resource

type Resources struct {
	Icons []Icon `json:"icons,omitempty" validate:"required"`
}

func GetList() Resources {
	i := Icon{
		Src:  "/src",
		Name: "name",
	}

	r := Resources{
		Icons: []Icon{i},
	}

	return r
}

func SetList(r *Resources) error {
	// todo

	return nil
}
