package paginator

type PageInfo struct {
	Page     int
	LastPage int
	Offset   int
	Limit    int
}

func Info(elemsCount int, pageSize int, offset int) PageInfo {
	limit := offset + pageSize
	if limit > elemsCount {
		limit = elemsCount
	}

	return PageInfo{
		Page:     offset / pageSize,
		LastPage: (elemsCount - 1) / pageSize,
		Offset:   offset,
		Limit:    limit - 1,
	}
}

func Paginate[T any](
	elems []T,
	pageSize int,
	iterFunc func(page []T, info PageInfo) error,
) error {
	for offset := 0; offset < len(elems); offset += pageSize {
		info := Info(len(elems), pageSize, offset)
		page := elems[offset : info.Limit+1]
		if err := iterFunc(page, info); err != nil {
			return err
		}
	}

	return nil
}
