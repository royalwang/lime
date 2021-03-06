// Copyright 2014 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

// This file was generated as part of a build step and shouldn't be manually modified
package sublime

import (
	"fmt"
	"github.com/limetext/gopy/lib"
	"github.com/limetext/lime/backend"
	"github.com/quarnster/util/text"
)

var (
	_ = backend.View{}
	_ = text.Region{}
	_ = fmt.Errorf
)

func (o *View) Py_change_count() (py.Object, error) {
	ret0 := o.data.Buffer().ChangeCount()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_file_name() (py.Object, error) {
	ret0 := o.data.Buffer().FileName()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) fullline(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 int
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for text.buffer.FullLine() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().FullLine(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) fullliner(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.buffer.FullLineR() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().FullLineR(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_buffer_id() (py.Object, error) {
	ret0 := o.data.Buffer().Id()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) line(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 int
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for text.buffer.Line() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().Line(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) liner(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.buffer.LineR() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().LineR(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_lines(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.buffer.Lines() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().Lines(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_name() (py.Object, error) {
	ret0 := o.data.Buffer().Name()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_rowcol(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 int
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for text.buffer.RowCol() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0, ret1 := o.data.Buffer().RowCol(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	var pyret1 py.Object

	pyret1, err = toPython(ret1)
	if err != nil {
		pyret0.Decref()
		return nil, err
	}
	return py.PackTuple(pyret0, pyret1)
}

func (o *View) Py_set_name(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 string
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(string); !ok {
				return nil, fmt.Errorf("Expected type string for text.buffer.SetName() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.Buffer().SetName(arg1)
	return toPython(nil)
}

func (o *View) Py_size() (py.Object, error) {
	ret0 := o.data.Buffer().Size()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_text_point(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 int
		arg2 int
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for text.buffer.TextPoint() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	if v, err := tu.GetItem(1); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for text.buffer.TextPoint() arg2, not %s", v.Type())
			} else {
				arg2 = v2
			}
		}
	}
	ret0 := o.data.Buffer().TextPoint(arg1, arg2)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) word(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 int
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for text.buffer.Word() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().Word(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) wordr(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for text.buffer.WordR() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.Buffer().WordR(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}
