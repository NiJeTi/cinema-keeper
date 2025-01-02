package paginator

func Paginate[T any](
	elems []T,
	pageSize int,
	iterFunc func(page []T) error,
) error {
	for offset := 0; offset < len(elems); offset += pageSize {
		limit := offset + pageSize
		if limit > len(elems) {
			limit = len(elems)
		}

		if err := iterFunc(elems[offset:limit]); err != nil {
			return err
		}
	}

	return nil
}
