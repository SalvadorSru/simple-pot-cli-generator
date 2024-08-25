package action

type Parameters struct {
	Singular int
	Plural   int
	Context  int
}

type Type struct {
	Name       string
	Plural     bool
	Parameters Parameters
}

func New(name string, plural int, context int) Type {
	return Type{
		Name: name,
		Parameters: Parameters{
			Singular: 0,
			Plural:   plural,
			Context:  context,
		},
	}
}

var AilableActions = []Type{
	New("__", -1, -1),
	New("_e", -1, -1),
	New("_ex", -1, 1),
	New("_n", 1, -1),
	New("_n_noop", 1, -1),
	New("_nx", 1, 3),
	New("_nx_noop", 1, 2),
	New("_x", -1, 1),
	New("esc_attr__", -1, -1),
	New("esc_attr_e", -1, -1),
	New("esc_attr_x", -1, 1),
	New("esc_html__", -1, -1),
	New("esc_html_e", -1, -1),
	New("esc_html_x", -1, 1),
}
