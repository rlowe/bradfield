package err

const (
  ALREADY_THERE = 1
  NO_WAY        = 2
)

type Error struct {
  ErrCode int
  ErrMsg string
}

func (e *Error) Error() string {
  return e.ErrMsg
}
