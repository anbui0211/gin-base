package service

import (
	"bufio"
	"context"
	"encoding/csv"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	ColumnTypeString         ColumnType = 0
	ColumnTypeDigit          ColumnType = 1
	ColumnTypeDecimal        ColumnType = 2
	ColumnTypeTypes          ColumnType = 3
	ColumnTypeDate           ColumnType = 4
	ColumnTypeDecimalPercent ColumnType = 5
)

const (
	patternRequireTypesDigit = `^[0-9]+$`
	patternTypesDigit        = `^[0-9]*$`

	patternTypeAlpha          = `^[A-Za-z]*$`
	patternRequireTypeAlpha   = `^[A-Za-z]+$`
	patternTypeAlphaVn        = `^[A-Za-z\s]+$` // support vietnamese
	patternRequireTypeAlphaVn = `^[A-Za-z\s]+$`

	patternStringDigit          = `.*`
	patternRequireAlphaNumDigit = `^[0-9a-zA-Z]+$`
	patternAlphaNumDigit        = `^[0-9a-zA-Z]*$`

	patternRequireDate = `^(20\d{2}|[HR]([1-9]|\d{2}))(\/0?[1-9]|0[1-9]|\/?1[0-2])(\/0?[1-9]|0[1-9]|\/?[12][0-9]|\/?3[0-1])$`

	patternDecimal = `^\d+(\.\d{2})?$` // 50.25
)

var (
	regexRequireAlpha   = ColumnDefinition{ColType: ColumnTypeString, regexp: regexp.MustCompile(patternRequireTypeAlpha)}
	regexRequireAlphaVn = ColumnDefinition{ColType: ColumnTypeString, regexp: regexp.MustCompile(patternRequireTypeAlphaVn)}
	regexRequireDigit   = ColumnDefinition{ColType: ColumnTypeDigit, regexp: regexp.MustCompile(patternRequireTypesDigit)}
	regexDecimal        = ColumnDefinition{ColType: ColumnTypeDecimal, regexp: regexp.MustCompile(patternDecimal)}
	regexStringDigit    = ColumnDefinition{ColType: ColumnTypeString, regexp: regexp.MustCompile(patternStringDigit)}
)

// ImportColumDef ...
type ImportColumDef struct {
	Name string
	Def  ColumnDefinition
}

// ColumnDefinition ...
type ColumnDefinition struct {
	ColType ColumnType
	regexp  *regexp.Regexp
}

type ColumnType uint16

// ImportRowSet ...
type ImportRowSet struct {
	def  map[int]string
	rows []*ImportRow
}

// ImportRow ...
type ImportRow struct {
	idx  int
	def  *map[int]string
	cols map[int]ImportColumn
}

// ImportColumn ...
type ImportColumn struct {
	Index int
	Value string
}

// readAndCheckCsv ...
func (s *importDataImpl) readAndCheckCsv(ctx context.Context, data io.ReadCloser, importColumnMap map[int]ImportColumDef, rs *ImportRowSet) error {
	icm := importColumnMap
	defLen := len(icm)

	rs.def = make(map[int]string, defLen)
	for idx := range icm {
		rs.def[idx] = icm[idx].Name
	}

	//fmt.Println("[readAndCheckCsv - rs.def]: ", rs.def) // -> map[0:name 1:searchString 2:categoryId 3:quantity 4:price 5:status]

	reader := newCsvReader(data)
	ridx := 0
	for {
		record, err := reader.Read()

		// No more data are available
		if err == io.EOF {
			break
		}

		// Must check nil error after check EOF error (otherwise you will panic)
		if err != nil {
			log.Fatal(err)
		}

		// If first character of the first column is "#", go to the next line without adding
		//fmt.Printf("first charactor of the first column is: %s\n", record[0])
		if strings.HasPrefix(record[0], "#") {
			continue
		}
		ridx++

		row := ImportRow{
			idx:  ridx,
			def:  &rs.def,
			cols: make(map[int]ImportColumn, len(record)),
		}

		for cidx := 0; cidx < defLen; cidx++ {
			if len(record) < cidx {
				continue
			}

			row.addColumn(ctx, cidx, icm[cidx].Def, record[cidx])
		}

		rs.addRow(&row)
	}

	return nil
}

// newCsvReader ...
func newCsvReader(r io.Reader) *csv.Reader {
	br := bufio.NewReader(r)
	bs, err := br.Peek(3)
	if err != nil {
		return csv.NewReader(br)
	}
	if bs[0] == 0xEF && bs[1] == 0xBB && bs[2] == 0xBF {
		_, _ = br.Discard(3)
	}

	return csv.NewReader(br)
}

// addColumn ...
func (ir *ImportRow) addColumn(ctx context.Context, cidx int, def ColumnDefinition, value string) {
	ir.cols[cidx] = ImportColumn{
		Index: cidx,
		Value: value,
	}

	// Check regex matchString -> làm sau

	if value == "" {
		return
	}

	// Switch case check  type column -> làm sau
}

// addRow ...
func (irs *ImportRowSet) addRow(row ...*ImportRow) {
	irs.rows = append(irs.rows, row...)
}

// String ...
func (ir *ImportRow) String(idx int) string {
	return ir.cols[idx].Value
}

// Float64 ...
func (ir *ImportRow) Float64(idx int) float64 {
	result, _ := strconv.ParseFloat(ir.cols[idx].Value, 64)
	return result
}

// Int64 ...
func (ir *ImportRow) Int64(idx int) int64 {
	result, _ := strconv.ParseInt(ir.cols[idx].Value, 0, 64)
	return result
}
