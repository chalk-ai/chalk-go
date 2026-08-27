package fixtures

import "github.com/apache/arrow-go/v18/arrow/memory"

var TestAllocator = memory.NewCheckedAllocator(memory.DefaultAllocator)
