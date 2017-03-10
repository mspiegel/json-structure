package jsonstructure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	multierror "github.com/mspiegel/go-multierror"
)

type shadowStructure JSONStructureDefinition

func (structure *JSONStructureDefinition) UnmarshalJSON(data []byte) error {
	var shell map[string]interface{}
	var shadow shadowStructure

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	err := decoder.Decode(&shell)
	if err != nil {
		return err
	}
	err = transformCompose(shell)
	if err != nil {
		return err
	}
	data, err = json.Marshal(shell)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &shadow)
	if err != nil {
		return err
	}
	*structure = JSONStructureDefinition(shadow)
	return nil
}

func elementOf(elem []string, candidate string) bool {
	for _, e := range elem {
		if e == candidate {
			return true
		}
	}
	return false
}

func mapMerge(dst map[string]interface{}, src map[string]interface{}, scope []string) error {
	var errs error
	for k, v := range src {
		newscope := append(scope, k)
		srcMap, srcOK := v.(map[string]interface{})
		dstValue := dst[k]
		if dstValue != nil {
			dstMap, dstOK := dstValue.(map[string]interface{})
			if (srcOK && !dstOK) || (!srcOK && dstOK) {
				err := errors.New("Attempted merge between map and non-map types")
				err = errorAt(err, newscope)
				errs = multierror.Append(errs, err)
			} else if srcOK && dstOK {
				err := mapMerge(dstMap, srcMap, newscope)
				errs = multierror.Append(errs, err)
			} else {
				dst[k] = v
			}
		} else {
			dst[k] = v
		}
	}
	return errs
}

func compose(target interface{},
	types map[string]interface{},
	fragments map[string]interface{},
	prev []string,
	scope []string) error {
	var err, errs error
	object, ok := target.(map[string]interface{})
	if !ok {
		err = errors.New("definition is not a JSON object")
		err = errorAt(err, scope)
		return err
	}
	for k, v := range object {
		if obj, ok2 := v.(map[string]interface{}); ok2 {
			newscope := append(scope, k)
			err = compose(obj, types, fragments, prev, newscope)
			errs = multierror.Append(err, errs)
		}
	}
	if errs != nil {
		return errs
	}
	c := object["compose"]
	if c == nil {
		return nil
	}
	defs, ok3 := c.([]interface{})
	if !ok3 {
		err = errors.New("'compose' property must be an string array")
		return errorAt(err, scope)
	}
	local := make(map[string]interface{})
	for k, v := range object {
		local[k] = v
		delete(object, k)
	}
	for _, iDef := range defs {
		defName, ok := iDef.(string)
		if !ok {
			err = errors.New("'compose' property must be an string array")
			return errorAt(err, scope)
		}
		if elementOf(prev, defName) {
			err = fmt.Errorf("cycle detected with definitions %v", prev)
			err = errorAt(err, scope)
			errs = multierror.Append(err, errs)
			continue
		}
		tDef := types[defName]
		fDef := fragments[defName]
		if (tDef == nil) && (fDef == nil) {
			err = fmt.Errorf("Unknown definition %s", defName)
			err = errorAt(err, scope)
			errs = multierror.Append(err, errs)
			continue
		}
		if (tDef != nil) && (fDef != nil) {
			// this should have been previously detected
			err = fmt.Errorf("Internal error duplicate definition %s", defName)
			err = errorAt(err, scope)
			errs = multierror.Append(err, errs)
			continue
		}
		var def map[string]interface{}
		if tDef != nil {
			next := append(prev, defName)
			newscope := []string{"types", defName}
			err = compose(tDef, types, fragments, next, newscope)
			errs = multierror.Append(err, errs)
			if err != nil {
				continue
			}
			def = tDef.(map[string]interface{})
		} else if fDef != nil {
			next := append(prev, defName)
			newscope := []string{"fragments", defName}
			err = compose(fDef, types, fragments, next, newscope)
			errs = multierror.Append(err, errs)
			if err != nil {
				continue
			}
			def = fDef.(map[string]interface{})
		}
		err = mapMerge(object, def, scope)
		errs = multierror.Append(err, errs)
	}
	err = mapMerge(object, local, scope)
	errs = multierror.Append(err, errs)
	if errs != nil {
		return errs
	}
	delete(object, "compose")
	return nil
}

func transformCompose(shell map[string]interface{}) error {
	var typ, frag map[string]interface{}
	var ok bool
	var err, errs error
	t := shell["types"]
	f := shell["fragments"]
	m := shell["main"]
	if m == nil {
		err = errors.New("JSON structure missing required 'main' property")
		err = errorAt(err, nil)
		return err
	}
	if t == nil {
		t = make(map[string]interface{})
	}
	if f == nil {
		f = make(map[string]interface{})
	}
	typ, ok = t.(map[string]interface{})
	if !ok {
		err = errors.New("'types' property must be a JSON object")
		err = errorAt(err, nil)
		errs = multierror.Append(errs, err)
	}
	frag, ok = f.(map[string]interface{})
	if !ok {
		err = errors.New("'fragments' property must be a JSON object")
		err = errorAt(err, nil)
		errs = multierror.Append(errs, err)
	}
	if errs != nil {
		return errs
	}
	dupls := intersection(typ, frag)
	if len(dupls) > 0 {
		err = fmt.Errorf("Duplicate keys across 'types' and 'fragments': %v", dupls)
		err = errorAt(err, nil)
		return err
	}
	for k, v := range frag {
		prev := []string{k}
		scope := []string{"fragments", k}
		if PrimitiveTypes[k] {
			err = errors.New("Cannot declare fragment with primitive type name")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else {
			err = compose(v, typ, frag, prev, scope)
			errs = multierror.Append(errs, err)
		}
	}
	for k, v := range typ {
		prev := []string{k}
		scope := []string{"types", k}
		if PrimitiveTypes[k] {
			err = errors.New("Cannot declare type with primitive type name")
			err = errorAt(err, scope)
			errs = multierror.Append(errs, err)
		} else {
			err = compose(v, typ, frag, prev, scope)
			errs = multierror.Append(errs, err)
		}
	}
	prev := []string{}
	scope := []string{"main"}
	err = compose(m, typ, frag, prev, scope)
	errs = multierror.Append(errs, err)
	return errs
}

func intersection(m1 map[string]interface{}, m2 map[string]interface{}) []string {
	var iSect []string
	if len(m1) > len(m2) {
		tmp := m1
		m1 = m2
		m2 = tmp
	}
	for k := range m1 {
		if _, ok := m2[k]; ok {
			iSect = append(iSect, k)
		}
	}
	return iSect
}
