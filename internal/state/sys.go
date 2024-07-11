package state

import (
	"bytes"
	"fmt"
	"sort"
	"time"

	restate "github.com/restatedev/sdk-go"
	"github.com/restatedev/sdk-go/generated/proto/protocol"
	"github.com/restatedev/sdk-go/internal/errors"
	"github.com/restatedev/sdk-go/internal/futures"
	"github.com/restatedev/sdk-go/internal/wire"
	"google.golang.org/protobuf/proto"
)

var (
	errEntryMismatch = restate.WithErrorCode(fmt.Errorf("log entry mismatch"), 32)
)

func (m *Machine) set(key string, value []byte) error {
	_, err := replayOrNew(
		m,
		func(entry *wire.SetStateEntryMessage) (void restate.Void, err error) {
			if string(entry.Key) != key || !bytes.Equal(entry.Value, value) {
				return void, errEntryMismatch
			}

			return
		}, func() (void restate.Void, err error) {
			return void, m._set(key, value)
		})
	if err != nil {
		return err
	}

	m.current[key] = value

	return nil
}

func (m *Machine) _set(key string, value []byte) error {
	return m.Write(
		&wire.SetStateEntryMessage{
			SetStateEntryMessage: protocol.SetStateEntryMessage{
				Key:   []byte(key),
				Value: value,
			},
		})
}

func (m *Machine) clear(key string) error {
	_, err := replayOrNew(
		m,
		func(entry *wire.ClearStateEntryMessage) (void restate.Void, err error) {
			if string(entry.Key) != key {
				return void, errEntryMismatch
			}

			return void, nil
		}, func() (restate.Void, error) {
			return restate.Void{}, m._clear(key)
		},
	)

	if err != nil {
		return err
	}

	delete(m.current, key)

	return err
}

func (m *Machine) _clear(key string) error {
	return m.Write(
		&wire.ClearStateEntryMessage{
			ClearStateEntryMessage: protocol.ClearStateEntryMessage{
				Key: []byte(key),
			},
		},
	)
}

func (m *Machine) clearAll() error {
	_, err := replayOrNew(
		m,
		func(entry *wire.ClearAllStateEntryMessage) (void restate.Void, err error) {
			return
		}, func() (restate.Void, error) {
			return restate.Void{}, m._clearAll()
		},
	)
	if err != nil {
		return err
	}

	m.current = map[string][]byte{}
	m.partial = false

	return nil
}

// clearAll drops all associated keys
func (m *Machine) _clearAll() error {
	return m.Write(
		&wire.ClearAllStateEntryMessage{},
	)
}

func (m *Machine) get(key string) ([]byte, error) {
	entry, err := replayOrNew(
		m,
		func(entry *wire.GetStateEntryMessage) (*wire.GetStateEntryMessage, error) {
			if string(entry.Key) != key {
				return nil, errEntryMismatch
			}
			return entry, nil
		}, func() (*wire.GetStateEntryMessage, error) {
			return m._get(key)
		})
	if err != nil {
		return nil, err
	}

	if err := entry.Await(m.ctx); err != nil {
		return nil, err
	}

	switch value := entry.Result.(type) {
	case *protocol.GetStateEntryMessage_Empty:
		return nil, nil
	case *protocol.GetStateEntryMessage_Failure:
		// the get state entry message is not failable so this should
		// never happen
		// TODO terminal?
		return nil, fmt.Errorf("[%d] %s", value.Failure.Code, value.Failure.Message)
	case *protocol.GetStateEntryMessage_Value:
		m.current[key] = value.Value
		return value.Value, nil
	}

	return nil, restate.TerminalError(fmt.Errorf("get state had invalid result: %v", entry.Result), errors.ErrProtocolViolation)
}

func (m *Machine) _get(key string) (*wire.GetStateEntryMessage, error) {
	msg := &wire.GetStateEntryMessage{
		GetStateEntryMessage: protocol.GetStateEntryMessage{
			Key: []byte(key),
		},
	}

	value, ok := m.current[key]

	if ok {
		// value in map, we still send the current
		// value to the runtime
		msg.Complete(&protocol.CompletionMessage{Result: &protocol.CompletionMessage_Value{Value: value}})

		if err := m.Write(msg); err != nil {
			return nil, err
		}

		return msg, nil
	}

	// key is not in map! there are 2 cases.
	if !m.partial {
		// current is complete. we need to return nil to the user
		// but also send an empty get state entry message
		msg.Complete(&protocol.CompletionMessage{Result: &protocol.CompletionMessage_Empty{Empty: &protocol.Empty{}}})

		if err := m.Write(msg); err != nil {
			return nil, err
		}

		return msg, nil
	}

	// we didn't see the value and we don't know for sure there isn't one; ask the runtime for it

	if err := m.Write(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Machine) keys() ([]string, error) {
	entry, err := replayOrNew(
		m,
		func(entry *wire.GetStateKeysEntryMessage) (*wire.GetStateKeysEntryMessage, error) {
			return entry, nil
		},
		m._keys,
	)
	if err != nil {
		return nil, err
	}

	if err := entry.Await(m.ctx); err != nil {
		return nil, err
	}

	switch value := entry.Result.(type) {
	case *protocol.GetStateKeysEntryMessage_Failure:
		// the get state entry message is not failable so this should
		// never happen
		return nil, fmt.Errorf("[%d] %s", value.Failure.Code, value.Failure.Message)
	case *protocol.GetStateKeysEntryMessage_Value:
		values := make([]string, 0, len(value.Value.Keys))
		for _, key := range value.Value.Keys {
			values = append(values, string(key))
		}

		return values, nil
	}

	return nil, nil
}

func (m *Machine) _keys() (*wire.GetStateKeysEntryMessage, error) {
	msg := &wire.GetStateKeysEntryMessage{}
	if !m.partial {
		keys := make([]string, 0, len(m.current))
		for k := range m.current {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		byteKeys := make([][]byte, len(keys))
		for i := range keys {
			byteKeys[i] = []byte(keys[i])
		}

		stateKeys := &protocol.GetStateKeysEntryMessage_StateKeys{Keys: byteKeys}
		value, err := proto.Marshal(stateKeys)
		if err != nil {
			return nil, err
		}

		// we can return keys entirely from cache
		// current is complete. we need to return nil to the user
		// but also send an empty get state entry message
		msg.Complete(&protocol.CompletionMessage{Result: &protocol.CompletionMessage_Value{
			Value: value,
		}})

		if err := m.Write(msg); err != nil {
			return nil, err
		}

		return nil, nil
	}

	if err := m.Write(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Machine) after(d time.Duration) (restate.After, error) {
	entry, err := replayOrNew(
		m,
		func(entry *wire.SleepEntryMessage) (*wire.SleepEntryMessage, error) {
			// we shouldn't verify the time because this would be different every time
			return entry, nil
		}, func() (*wire.SleepEntryMessage, error) {
			return m._sleep(d)
		},
	)
	if err != nil {
		return nil, err
	}

	return futures.NewAfter(m.ctx, entry), nil
}

func (m *Machine) sleep(d time.Duration) error {
	after, err := m.after(d)
	if err != nil {
		return err
	}

	return after.Done()
}

// _sleep creating a new sleep entry.
func (m *Machine) _sleep(d time.Duration) (*wire.SleepEntryMessage, error) {
	msg := &wire.SleepEntryMessage{
		SleepEntryMessage: protocol.SleepEntryMessage{
			WakeUpTime: uint64(time.Now().Add(d).UnixMilli()),
		},
	}
	if err := m.Write(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (m *Machine) sideEffect(fn func() ([]byte, error)) ([]byte, error) {
	entry, err := replayOrNew(
		m,
		func(entry *wire.RunEntryMessage) (*wire.RunEntryMessage, error) {
			return entry, nil
		},
		func() (*wire.RunEntryMessage, error) {
			return m._sideEffect(fn)
		},
	)
	if err != nil {
		// either a transient error from the fn or from our sending of the result
		return nil, err
	}

	// side effect must be acknowledged before proceeding
	if err := entry.Await(m.ctx); err != nil {
		return nil, err
	}

	switch result := entry.Result.(type) {
	case *protocol.RunEntryMessage_Failure:
		return nil, errors.ErrorFromFailure(result.Failure)
	case *protocol.RunEntryMessage_Value:
		return result.Value, nil
	case nil:
		// Empty result is valid
		return nil, nil
	}

	return nil, restate.TerminalError(fmt.Errorf("side effect entry had invalid result: %v", entry.Result), errors.ErrProtocolViolation)
}

func (m *Machine) _sideEffect(fn func() ([]byte, error)) (*wire.RunEntryMessage, error) {
	bytes, err := fn()

	if err != nil {
		if restate.IsTerminalError(err) {
			msg := &wire.RunEntryMessage{
				RunEntryMessage: protocol.RunEntryMessage{
					Result: &protocol.RunEntryMessage_Failure{
						Failure: &protocol.Failure{
							Code:    uint32(restate.ErrorCode(err)),
							Message: err.Error(),
						},
					},
				},
			}
			if err := m.Write(msg); err != nil {
				return nil, err
			}

			// don't return the original error, we will turn the entry back into an error later
			// that way its not different replay vs non-replay
			return msg, nil
		} else {
			ty := uint32(wire.RunEntryMessageType)
			msg := wire.ErrorMessage{
				ErrorMessage: protocol.ErrorMessage{
					Code:             uint32(restate.ErrorCode(err)),
					Message:          err.Error(),
					RelatedEntryType: &ty,
				},
			}
			if err := m.protocol.Write(&msg); err != nil {
				return nil, err
			}

			return nil, err
		}
	} else {
		msg := &wire.RunEntryMessage{
			RunEntryMessage: protocol.RunEntryMessage{
				Result: &protocol.RunEntryMessage_Value{
					Value: bytes,
				},
			},
		}
		if err := m.Write(msg); err != nil {
			return nil, err
		}

		return msg, nil
	}
}
