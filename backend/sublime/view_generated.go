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

var _viewClass = py.Class{
	Name:    "sublime.View",
	Pointer: (*View)(nil),
}

type View struct {
	py.BaseObject
	data *backend.View
}

func (o *View) PyInit(args *py.Tuple, kwds *py.Dict) error {
	return fmt.Errorf("Can't initialize type View")
}
func (o *View) Py_begin_edit() (py.Object, error) {
	ret0 := o.data.BeginEdit()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_end_edit(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 *backend.Edit
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(*backend.Edit); !ok {
				return nil, fmt.Errorf("Expected type *backend.Edit for backend.View.EndEdit() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.EndEdit(arg1)
	return toPython(nil)
}

func (o *View) Py_erase(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 *backend.Edit
		arg2 text.Region
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(*backend.Edit); !ok {
				return nil, fmt.Errorf("Expected type *backend.Edit for backend.View.Erase() arg1, not %s", v.Type())
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
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for backend.View.Erase() arg2, not %s", v.Type())
			} else {
				arg2 = v2
			}
		}
	}
	o.data.Erase(arg1, arg2)
	return toPython(nil)
}

func (o *View) Py_erase_regions(tu *py.Tuple) (py.Object, error) {
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
				return nil, fmt.Errorf("Expected type string for backend.View.EraseRegions() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.EraseRegions(arg1)
	return toPython(nil)
}

func (o *View) Py_extract_scope(tu *py.Tuple) (py.Object, error) {
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
				return nil, fmt.Errorf("Expected type int for backend.View.ExtractScope() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.ExtractScope(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_get_regions(tu *py.Tuple) (py.Object, error) {
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
				return nil, fmt.Errorf("Expected type string for backend.View.GetRegions() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.GetRegions(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_id() (py.Object, error) {
	ret0 := o.data.Id()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_insert(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 *backend.Edit
		arg2 int
		arg3 string
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(*backend.Edit); !ok {
				return nil, fmt.Errorf("Expected type *backend.Edit for backend.View.Insert() arg1, not %s", v.Type())
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
				return nil, fmt.Errorf("Expected type int for backend.View.Insert() arg2, not %s", v.Type())
			} else {
				arg2 = v2
			}
		}
	}
	if v, err := tu.GetItem(2); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(string); !ok {
				return nil, fmt.Errorf("Expected type string for backend.View.Insert() arg3, not %s", v.Type())
			} else {
				arg3 = v2
			}
		}
	}
	ret0 := o.data.Insert(arg1, arg2, arg3)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_is_dirty() (py.Object, error) {
	ret0 := o.data.IsDirty()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_is_scratch() (py.Object, error) {
	ret0 := o.data.IsScratch()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_overwrite_status() (py.Object, error) {
	ret0 := o.data.OverwriteStatus()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_replace(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 *backend.Edit
		arg2 text.Region
		arg3 string
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(*backend.Edit); !ok {
				return nil, fmt.Errorf("Expected type *backend.Edit for backend.View.Replace() arg1, not %s", v.Type())
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
			if v2, ok := v3.(text.Region); !ok {
				return nil, fmt.Errorf("Expected type text.Region for backend.View.Replace() arg2, not %s", v.Type())
			} else {
				arg2 = v2
			}
		}
	}
	if v, err := tu.GetItem(2); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(string); !ok {
				return nil, fmt.Errorf("Expected type string for backend.View.Replace() arg3, not %s", v.Type())
			} else {
				arg3 = v2
			}
		}
	}
	o.data.Replace(arg1, arg2, arg3)
	return toPython(nil)
}

func (o *View) Py_scope_name(tu *py.Tuple) (py.Object, error) {
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
				return nil, fmt.Errorf("Expected type int for backend.View.ScopeName() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	ret0 := o.data.ScopeName(arg1)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_score_selector(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 int
		arg2 string
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(int); !ok {
				return nil, fmt.Errorf("Expected type int for backend.View.ScoreSelector() arg1, not %s", v.Type())
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
			if v2, ok := v3.(string); !ok {
				return nil, fmt.Errorf("Expected type string for backend.View.ScoreSelector() arg2, not %s", v.Type())
			} else {
				arg2 = v2
			}
		}
	}
	ret0 := o.data.ScoreSelector(arg1, arg2)
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_sel() (py.Object, error) {
	ret0 := o.data.Sel()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_set_overwrite_status(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 bool
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(bool); !ok {
				return nil, fmt.Errorf("Expected type bool for backend.View.SetOverwriteStatus() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.SetOverwriteStatus(arg1)
	return toPython(nil)
}

func (o *View) Py_set_scratch(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 bool
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v3, err2 := fromPython(v); err2 != nil {
			return nil, err2
		} else {
			if v2, ok := v3.(bool); !ok {
				return nil, fmt.Errorf("Expected type bool for backend.View.SetScratch() arg1, not %s", v.Type())
			} else {
				arg1 = v2
			}
		}
	}
	o.data.SetScratch(arg1)
	return toPython(nil)
}

func (o *View) Py_settings() (py.Object, error) {
	ret0 := o.data.Settings()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *View) Py_window() (py.Object, error) {
	ret0 := o.data.Window()
	var err error
	var pyret0 py.Object

	pyret0, err = toPython(ret0)
	if err != nil {
		return nil, err
	}
	return pyret0, err
}
