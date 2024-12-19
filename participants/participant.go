package participants

type Participant interface {
	GetName() string
	Run(distance int) bool
	Jump(height int) bool
}
