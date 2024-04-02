package v1

import (
	"github.com/parquet-go/parquet-go"

	profilev1 "github.com/grafana/pyroscope/api/gen/proto/go/google/v1"
)

var mappingsSchema = parquet.SchemaOf(new(profilev1.Mapping))

type MappingPersister struct{}

func (MappingPersister) Name() string { return "mappings" }

func (MappingPersister) Schema() *parquet.Schema { return mappingsSchema }

func (MappingPersister) SortingColumns() parquet.SortingOption { return parquet.SortingColumns() }

func (MappingPersister) Deconstruct(row parquet.Row, _ uint64, m InMemoryMapping) parquet.Row {
	if cap(row) < 10 {
		row = make(parquet.Row, 0, 10)
	}
	row = row[:0]
	row = append(row, parquet.Int64Value(int64(m.Id)).Level(0, 0, 0))
	row = append(row, parquet.Int64Value(int64(m.MemoryStart)).Level(0, 0, 1))
	row = append(row, parquet.Int64Value(int64(m.MemoryLimit)).Level(0, 0, 2))
	row = append(row, parquet.Int64Value(int64(m.FileOffset)).Level(0, 0, 3))
	row = append(row, parquet.Int32Value(int32(m.Filename)).Level(0, 0, 4))
	row = append(row, parquet.Int32Value(int32(m.BuildId)).Level(0, 0, 5))
	row = append(row, parquet.BooleanValue(m.HasFunctions).Level(0, 0, 6))
	row = append(row, parquet.BooleanValue(m.HasFilenames).Level(0, 0, 7))
	row = append(row, parquet.BooleanValue(m.HasLineNumbers).Level(0, 0, 8))
	row = append(row, parquet.BooleanValue(m.HasInlineFrames).Level(0, 0, 9))
	return row
}

func (MappingPersister) Reconstruct(row parquet.Row) (uint64, InMemoryMapping, error) {
	mapping := InMemoryMapping{
		Id:              row[0].Uint64(),
		MemoryStart:     row[1].Uint64(),
		MemoryLimit:     row[2].Uint64(),
		FileOffset:      row[3].Uint64(),
		Filename:        row[4].Uint32(),
		BuildId:         row[5].Uint32(),
		HasFunctions:    row[6].Boolean(),
		HasFilenames:    row[7].Boolean(),
		HasLineNumbers:  row[8].Boolean(),
		HasInlineFrames: row[9].Boolean(),
	}
	return 0, mapping, nil
}

type InMemoryMapping struct {
	// Unique nonzero id for the mapping.
	Id uint64
	// Address at which the binary (or DLL) is loaded into memory.
	MemoryStart uint64
	// The limit of the address range occupied by this mapping.
	MemoryLimit uint64
	// Offset in the binary that corresponds to the first mapped address.
	FileOffset uint64
	// The object this entry is loaded from.  This can be a filename on
	// disk for the main binary and shared libraries, or virtual
	// abstractions like "[vdso]".
	Filename uint32
	// A string that uniquely identifies a particular program version
	// with high probability. E.g., for binaries generated by GNU tools,
	// it could be the contents of the .note.gnu.build-id field.
	BuildId uint32
	// The following fields indicate the resolution of symbolic info.
	HasFunctions    bool
	HasFilenames    bool
	HasLineNumbers  bool
	HasInlineFrames bool
}

func (m InMemoryMapping) Clone() InMemoryMapping {
	return m
}
