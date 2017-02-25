package types

type PrecisionSet []*MatchedMethod

func (pq *PrecisionSet) Max() *MatchedMethod {
	if pq == nil || *pq == nil {
		return nil
	}
	max := (*pq)[0]
	for i := 1; i < len(*pq); i++ {
		current := (*pq)[i]
		if max.precision < current.precision {
			max = current
		}
	}
	return max
}
