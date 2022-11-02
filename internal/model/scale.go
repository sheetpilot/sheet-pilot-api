package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strings"

	pb "github.com/sheetpilot/sheet-pilot-proto/scaleservice"
)

// [{"colName":"Applications","value":"api"},{"colName":"request CPU(string)","value":9},{"colName":"request Memory(string)","value":"5Gi"},{"colName":"limit CPU(string)","value":2},{"colName":"limit Memory(string)","value":"2Gi"},{"colName":"replica_count(int)","value":"6sd"}]
// type Scale struct {
// 	ColName string      `json:"colName"`
// 	Value   interface{} `json:"value"`
// }

type ScaleObject struct {
	Obj []map[string]any `json:"obj"`

	UpdatedRow []*pb.ScaleRequest_Updatedrow
}

func TransformScaleObject(reqBody io.ReadCloser) (*ScaleObject, error) {
	scaleObject := make([]map[string]any, 0)

	decoder := json.NewDecoder(reqBody)
	err := decoder.Decode(&scaleObject)
	if err != nil {
		return nil, errors.New("failed to decode request body")
	}

	return &ScaleObject{
		Obj: scaleObject,
	}, nil
}

func (s *ScaleObject) Validate() error {
	if len(s.Obj) == 0 {
		return errors.New("empty request")
	}

	for _, field := range s.Obj {
		colName, ok := field["colName"].(string)
		if !ok {
			return errors.New("invalid request: colName is missing")
		}

		value, ok := field["value"]
		if !ok {
			return errors.New("invalid request: value is missing")
		}

		regexStr, dataType, err := s.getTypeRegex(colName)
		if err != nil {
			return err
		}

		if reflect.TypeOf(value).Kind() != dataType {
			return fmt.Errorf("value: %v on col %s is not a valid type", value, colName)
		}

		r := regexp.MustCompile(regexStr)

		valByte, _ := json.Marshal(value)
		if !r.Match(valByte) {
			return errors.New("validation error")
		}

		s.appendScaleProto(colName, string(valByte), dataType.String())
	}

	return nil
}

func (s *ScaleObject) appendScaleProto(colName string, value string, dataType string) {
	s.UpdatedRow = append(s.UpdatedRow, &pb.ScaleRequest_Updatedrow{
		ColName:  colName,
		Value:    value,
		DataType: dataType,
	})
}

func (s *ScaleObject) getTypeRegex(colName string) (string, reflect.Kind, error) {
	allowedTypes := map[string]reflect.Kind{
		"string": reflect.String,
		"int":    reflect.Int,
	}
	colNameSperator := strings.Split(colName, "|")
	if len(colNameSperator) != 2 {
		return "", reflect.Invalid, errors.New("colName is invalid")
	}

	valType, ok := allowedTypes[colNameSperator[1]]
	if !ok {
		return "", reflect.Invalid, fmt.Errorf("invalid column value type: %s", colNameSperator[1])
	}

	return colNameSperator[2], valType, nil
}
