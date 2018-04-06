package generator

type outputter interface {
	OutputModel(m *modelSchema) error
}
